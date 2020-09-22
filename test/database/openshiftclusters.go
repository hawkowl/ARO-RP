package database

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/fvbommel/sortorder"

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/database"
	"github.com/Azure/ARO-RP/pkg/database/cosmosdb"
)

type ByKey []*api.OpenShiftClusterDocument

func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return sortorder.NaturalLess(a[i].Key, a[j].Key) }

func getQueuedOpenShiftDocuments(client cosmosdb.OpenShiftClusterDocumentClient) (res []*api.OpenShiftClusterDocument) {
	docs := fakeOpenShiftClustersGetAllDocuments(client)

	for _, r := range docs {
		var include bool

		switch r.OpenShiftCluster.Properties.ProvisioningState {
		case
			api.ProvisioningStateCreating,
			api.ProvisioningStateUpdating,
			api.ProvisioningStateAdminUpdating,
			api.ProvisioningStateDeleting:
			include = true
		}

		if include && (r.LeaseExpires > 0 && int64(r.LeaseExpires) < time.Now().Unix()) {
			include = false
		}

		if include {
			res = append(res, r)
		}
	}
	return
}

func fakeOpenShiftClustersQueueLengthQuery(client cosmosdb.OpenShiftClusterDocumentClient, query *cosmosdb.Query, options *cosmosdb.Options) cosmosdb.OpenShiftClusterDocumentRawIterator {
	results := getQueuedOpenShiftDocuments(client)
	return &fakeOpenShiftClustersQueueLengthIterator{resultCount: len(results)}
}

func fakeOpenShiftClustersDequeueQuery(client cosmosdb.OpenShiftClusterDocumentClient, query *cosmosdb.Query, options *cosmosdb.Options) cosmosdb.OpenShiftClusterDocumentRawIterator {
	docs := getQueuedOpenShiftDocuments(client)
	return cosmosdb.NewFakeOpenShiftClusterDocumentClientRawIterator(docs, 0)
}

func fakeOpenshiftClustersMatchQuery(client cosmosdb.OpenShiftClusterDocumentClient, query *cosmosdb.Query, options *cosmosdb.Options) cosmosdb.OpenShiftClusterDocumentRawIterator {
	var results []*api.OpenShiftClusterDocument

	startingIndex, err := fakeOpenShiftClustersGetContinuation(options)
	if err != nil {
		return cosmosdb.NewFakeOpenShiftClusterDocumentClientErroringRawIterator(err)
	}

	docs := fakeOpenShiftClustersGetAllDocuments(client)
	for _, r := range docs {
		var key string
		switch query.Parameters[0].Name {
		case "@key":
			key = r.Key
		case "@clientID":
			key = r.ClientIDKey
		case "@resourceGroupID":
			key = r.ClusterResourceGroupIDKey
		default:
			return cosmosdb.NewFakeOpenShiftClusterDocumentClientErroringRawIterator(cosmosdb.ErrNotImplemented)
		}
		if key == query.Parameters[0].Value {
			results = append(results, r)
		}
	}
	return cosmosdb.NewFakeOpenShiftClusterDocumentClientRawIterator(results, int(startingIndex))
}

func fakeOpenShiftClustersGetAllDocuments(client cosmosdb.OpenShiftClusterDocumentClient) []*api.OpenShiftClusterDocument {
	input, err := client.ListAll(context.Background(), nil)
	if err != nil {
		// TODO: should this never happen?
		panic(err)
	}
	docs := input.OpenShiftClusterDocuments
	sort.Sort(ByKey(docs))
	return docs
}

func fakeOpenShiftClustersGetContinuation(options *cosmosdb.Options) (startingIndex int64, err error) {
	if options != nil && options.Continuation != "" {
		startingIndex, err = strconv.ParseInt(options.Continuation, 10, 64)
	}
	return
}

func fakeOpenshiftClustersPrefixQuery(client cosmosdb.OpenShiftClusterDocumentClient, query *cosmosdb.Query, options *cosmosdb.Options) cosmosdb.OpenShiftClusterDocumentRawIterator {
	startingIndex, err := fakeOpenShiftClustersGetContinuation(options)
	if err != nil {
		return cosmosdb.NewFakeOpenShiftClusterDocumentClientErroringRawIterator(err)
	}

	docs := fakeOpenShiftClustersGetAllDocuments(client)
	var results []*api.OpenShiftClusterDocument
	for _, r := range docs {
		if strings.Index(r.Key, query.Parameters[0].Value) == 0 {
			results = append(results, r)
		}
	}
	return cosmosdb.NewFakeOpenShiftClusterDocumentClientRawIterator(results, int(startingIndex))
}

func fakeOpenShiftClustersRenewLeaseTrigger(ctx context.Context, doc *api.OpenShiftClusterDocument) error {
	doc.LeaseExpires = int(time.Now().Unix()) + 60
	return nil
}

func openShiftClusterConflictChecker(one *api.OpenShiftClusterDocument, two *api.OpenShiftClusterDocument) bool {
	if one.ClusterResourceGroupIDKey != "" && two.ClusterResourceGroupIDKey != "" && one.ClusterResourceGroupIDKey == two.ClusterResourceGroupIDKey {
		return true
	}
	if one.ClientIDKey != "" && two.ClientIDKey != "" && one.ClientIDKey == two.ClientIDKey {
		return true
	}
	return false
}

func injectOpenShiftClusters(c *cosmosdb.FakeOpenShiftClusterDocumentClient) {
	c.InjectQuery(database.OpenShiftClustersDequeueQuery, fakeOpenShiftClustersDequeueQuery)
	c.InjectQuery(database.OpenShiftClustersQueueLengthQuery, fakeOpenShiftClustersQueueLengthQuery)
	c.InjectQuery(database.OpenShiftClustersGetQuery, fakeOpenshiftClustersMatchQuery)
	c.InjectQuery(database.OpenshiftClustersClientIdQuery, fakeOpenshiftClustersMatchQuery)
	c.InjectQuery(database.OpenshiftClustersResourceGroupQuery, fakeOpenshiftClustersMatchQuery)
	c.InjectQuery(database.OpenshiftClustersPrefixQuery, fakeOpenshiftClustersPrefixQuery)

	c.InjectTrigger("renewLease", fakeOpenShiftClustersRenewLeaseTrigger)

	c.UseSorter(func(in []*api.OpenShiftClusterDocument) { sort.Sort(ByKey(in)) })
	c.UseDocumentConflictChecker(openShiftClusterConflictChecker)
}

// fakeOpenShiftClustersQueueLengthIterator is a RawIterator that will produce a
// document containing a list of a single integer when NextRaw is called.
type fakeOpenShiftClustersQueueLengthIterator struct {
	called      bool
	resultCount int
}

func (i *fakeOpenShiftClustersQueueLengthIterator) Next(ctx context.Context, maxItemCount int) (*api.OpenShiftClusterDocuments, error) {
	return nil, cosmosdb.ErrNotImplemented
}

func (i *fakeOpenShiftClustersQueueLengthIterator) NextRaw(ctx context.Context, continuation int, out interface{}) error {
	if i.called {
		return errors.New("Can't call twice!")
	}
	i.called = true

	res := fmt.Sprintf(`{"Count": 1, "Documents": [%d]}`, i.resultCount)
	return json.NewDecoder(bytes.NewBufferString(res)).Decode(out)
}

func (i *fakeOpenShiftClustersQueueLengthIterator) Continuation() string {
	return ""
}

func GetResourcePath(subscriptionID string, resourceID string) string {
	return fmt.Sprintf("/subscriptions/%s/resourcegroups/resourceGroup/providers/Microsoft.RedHatOpenShift/openShiftClusters/%s", subscriptionID, resourceID)
}
