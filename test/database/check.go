package database

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-test/deep"

	"github.com/Azure/ARO-RP/pkg/api"
)

type Checker struct {
	openshiftClusterDocuments []*api.OpenShiftClusterDocument
	subscriptionDocuments     []*api.SubscriptionDocument
	billingDocuments          []*api.BillingDocument
	asyncOperationDocuments   []*api.AsyncOperationDocument

	clients *FakeClients
}

func NewChecker(clients *FakeClients) *Checker {
	return &Checker{clients: clients}
}

func (f *Checker) AddOpenShiftClusterDocuments(docs []*api.OpenShiftClusterDocument) {
	f.openshiftClusterDocuments = append(f.openshiftClusterDocuments, docs...)
}

func (f *Checker) AddOpenShiftClusterDocument(doc *api.OpenShiftClusterDocument) {
	f.openshiftClusterDocuments = append(f.openshiftClusterDocuments, doc)
}

func (f *Checker) AddSubscriptionDocuments(docs []*api.SubscriptionDocument) {
	f.subscriptionDocuments = append(f.subscriptionDocuments, docs...)
}

func (f *Checker) AddSubscriptionDocument(doc *api.SubscriptionDocument) {
	f.subscriptionDocuments = append(f.subscriptionDocuments, doc)
}

func (f *Checker) AddBillingDocuments(docs []*api.BillingDocument) {
	f.billingDocuments = append(f.billingDocuments, docs...)
}

func (f *Checker) AddBillingDocument(doc *api.BillingDocument) {
	f.billingDocuments = append(f.billingDocuments, doc)
}

func (f *Checker) AddAsyncOperationDocuments(docs []*api.AsyncOperationDocument) {
	f.asyncOperationDocuments = append(f.asyncOperationDocuments, docs...)
}

func (f *Checker) AddAsyncOperationDocument(doc *api.AsyncOperationDocument) {
	f.asyncOperationDocuments = append(f.asyncOperationDocuments, doc)
}

func (f *Checker) Check() []error {
	var errs []error
	ctx := context.Background()

	allAsyncDocs, err := f.clients.AsyncOperations.ListAll(ctx, nil)
	if err != nil {
		return []error{err}
	}

	if len(f.asyncOperationDocuments) != 0 && len(allAsyncDocs.AsyncOperationDocuments) == len(f.asyncOperationDocuments) {
		diff := deep.Equal(allAsyncDocs.AsyncOperationDocuments, f.asyncOperationDocuments)
		if diff != nil {
			for _, i := range diff {
				errs = append(errs, errors.New(i))
			}
		}
	} else if len(allAsyncDocs.AsyncOperationDocuments) != 0 || len(f.asyncOperationDocuments) != 0 {
		errs = append(errs, fmt.Errorf("async docs length different, %d vs %d", len(allAsyncDocs.AsyncOperationDocuments), len(f.asyncOperationDocuments)))
	}

	allOpenShiftDocs, err := f.clients.OpenShiftClusters.ListAll(ctx, nil)
	if err != nil {
		return []error{err}
	}

	if len(f.openshiftClusterDocuments) != 0 && len(allOpenShiftDocs.OpenShiftClusterDocuments) == len(f.openshiftClusterDocuments) {
		diff := deep.Equal(allOpenShiftDocs.OpenShiftClusterDocuments, f.openshiftClusterDocuments)
		if diff != nil {
			for _, i := range diff {
				errs = append(errs, errors.New(i))
			}
		}
	} else if len(allOpenShiftDocs.OpenShiftClusterDocuments) != 0 || len(f.openshiftClusterDocuments) != 0 {
		errs = append(errs, fmt.Errorf("async docs length different, %d vs %d", len(allOpenShiftDocs.OpenShiftClusterDocuments), len(f.openshiftClusterDocuments)))
	}
	return errs
}
