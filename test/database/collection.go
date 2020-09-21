package database

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"

	"github.com/Azure/ARO-RP/pkg/database/cosmosdb"
)

type collectionClient struct {
}

// NewCollectionClient returns a new collection client
func newCollectionClient(c cosmosdb.DatabaseClient, dbid string) cosmosdb.CollectionClient {
	return &collectionClient{}
}

func (c *collectionClient) Create(ctx context.Context, newcoll *cosmosdb.Collection) (coll *cosmosdb.Collection, err error) {
	return nil, nil
}

func (c *collectionClient) List() cosmosdb.CollectionIterator {
	return nil
}

func (c *collectionClient) ListAll(ctx context.Context) (*cosmosdb.Collections, error) {
	return nil, nil
}

func (c *collectionClient) Get(ctx context.Context, collid string) (coll *cosmosdb.Collection, err error) {
	return nil, nil
}

func (c *collectionClient) Delete(ctx context.Context, coll *cosmosdb.Collection) error {
	return nil
}

func (c *collectionClient) Replace(ctx context.Context, newcoll *cosmosdb.Collection) (coll *cosmosdb.Collection, err error) {
	return nil, nil
}

func (c *collectionClient) PartitionKeyRanges(ctx context.Context, collid string) (*cosmosdb.PartitionKeyRanges, error) {
	return &cosmosdb.PartitionKeyRanges{
		Count:      1,
		ResourceID: collid,
		PartitionKeyRanges: []cosmosdb.PartitionKeyRange{
			{
				ID: "singular",
			},
		},
	}, nil
}
