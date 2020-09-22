package database

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"

	"github.com/Azure/ARO-RP/pkg/database"
	"github.com/Azure/ARO-RP/pkg/database/cosmosdb"
)

var fakeCode []byte = []byte{'F', 'A', 'K', 'E'}

type fakeCipher struct {
}

func (c fakeCipher) Decrypt(in []byte) ([]byte, error) {
	return in[4:], nil
}

func (c fakeCipher) Encrypt(in []byte) ([]byte, error) {
	out := make([]byte, 4+len(in))
	_ = copy(out, fakeCode)
	_ = copy(out[4:], in)
	return out, nil
}

func NewFakeCipher() *fakeCipher {
	return &fakeCipher{}
}

type FakeClients struct {
	OpenShiftClusters *cosmosdb.FakeOpenShiftClusterDocumentClient
	Subscriptions     *cosmosdb.FakeSubscriptionDocumentClient
	Billing           *cosmosdb.FakeBillingDocumentClient
	AsyncOperations   *cosmosdb.FakeAsyncOperationDocumentClient
}

func (c *FakeClients) MakeUnavailable(err error) {
	c.OpenShiftClusters.MakeUnavailable(err)
	c.Subscriptions.MakeUnavailable(err)
	c.Billing.MakeUnavailable(err)
	c.AsyncOperations.MakeUnavailable(err)
}

func NewDatabase(ctx context.Context, log *logrus.Entry) (*database.Database, *FakeClients, string, error) {
	cipher := &fakeCipher{}
	uuid := uuid.NewV4().String()
	h := database.NewJSONHandle(cipher)

	coll := &collectionClient{}
	osc := cosmosdb.NewFakeOpenShiftClusterDocumentClient(h)
	sub := cosmosdb.NewFakeSubscriptionDocumentClient(h)
	bil := cosmosdb.NewFakeBillingDocumentClient(h)
	asy := cosmosdb.NewFakeAsyncOperationDocumentClient(h)

	injectOpenShiftClusters(osc)
	injectSubscriptions(sub)
	injectBilling(bil)

	db := &database.Database{
		OpenShiftClusters: database.NewOpenShiftClustersWithProvidedClient(uuid, osc, coll),
		Subscriptions:     database.NewSubscriptionsWithProvidedClient(uuid, sub),
		Billing:           database.NewBillingWithProvidedClient(uuid, bil),
		AsyncOperations:   database.NewAsyncOperationsWithProvidedClient(uuid, asy),
	}

	clients := &FakeClients{
		OpenShiftClusters: osc,
		Subscriptions:     sub,
		Billing:           bil,
		AsyncOperations:   asy,
	}

	return db, clients, uuid, nil
}
