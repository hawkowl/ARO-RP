package frontend

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/database/cosmosdb"
	"github.com/Azure/ARO-RP/pkg/metrics/noop"
	testdatabase "github.com/Azure/ARO-RP/test/database"
)

func TestDeleteOpenShiftCluster(t *testing.T) {
	ctx := context.Background()

	mockSubID := "00000000-0000-0000-0000-000000000000"

	type test struct {
		name           string
		resourceID     string
		fixture        func(*testdatabase.Fixture)
		dbError        error
		wantDocuments  func(*testdatabase.Checker)
		wantStatusCode int
		wantAsync      bool
		wantError      string
	}

	for _, tt := range []*test{
		{
			name:       "cluster exists in db",
			resourceID: testdatabase.GetResourcePath(mockSubID, "resourceName"),
			fixture: func(f *testdatabase.Fixture) {
				f.AddSubscriptionDocument(&api.SubscriptionDocument{
					ID: mockSubID,
					Subscription: &api.Subscription{
						State: api.SubscriptionStateRegistered,
						Properties: &api.SubscriptionProperties{
							TenantID: "11111111-1111-1111-1111-111111111111",
						},
					},
				})
				f.AddOpenShiftClusterDocument(&api.OpenShiftClusterDocument{
					Key:      strings.ToLower(testdatabase.GetResourcePath(mockSubID, "resourceName")),
					Dequeues: 1,
					OpenShiftCluster: &api.OpenShiftCluster{
						ID:   testdatabase.GetResourcePath(mockSubID, "resourceName"),
						Name: "resourceName",
						Type: "Microsoft.RedHatOpenShift/openshiftClusters",
						Properties: api.OpenShiftClusterProperties{
							ProvisioningState: api.ProvisioningStateSucceeded,
						},
					},
				})
			},
			wantDocuments: func(c *testdatabase.Checker) {
				c.AddAsyncOperationDocument(&api.AsyncOperationDocument{
					OpenShiftClusterKey: strings.ToLower(testdatabase.GetResourcePath(mockSubID, "resourceName")),
					AsyncOperation: &api.AsyncOperation{
						InitialProvisioningState: api.ProvisioningStateDeleting,
						ProvisioningState:        api.ProvisioningStateDeleting,
					},
				})

				c.AddOpenShiftClusterDocument(&api.OpenShiftClusterDocument{
					Key: strings.ToLower(testdatabase.GetResourcePath(mockSubID, "resourceName")),
					OpenShiftCluster: &api.OpenShiftCluster{
						ID:   testdatabase.GetResourcePath(mockSubID, "resourceName"),
						Name: "resourceName",
						Type: "Microsoft.RedHatOpenShift/openshiftClusters",
						Properties: api.OpenShiftClusterProperties{
							ProvisioningState:     api.ProvisioningStateDeleting,
							LastProvisioningState: api.ProvisioningStateSucceeded,
						},
					},
				})
			},
			wantStatusCode: http.StatusAccepted,
			wantAsync:      true,
		},
		{
			name:           "cluster not found in db",
			resourceID:     testdatabase.GetResourcePath(mockSubID, "resourceName"),
			wantStatusCode: http.StatusNoContent,
		},
		{
			name:           "internal error",
			resourceID:     testdatabase.GetResourcePath(mockSubID, "resourceName"),
			dbError:        &cosmosdb.Error{Code: "500", Message: "blah"},
			wantStatusCode: http.StatusInternalServerError,
			wantError:      `500: InternalServerError: : Internal server error.`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ti, err := newTestInfra(t)
			if err != nil {
				t.Fatal(err)
			}
			defer ti.done()

			err = ti.buildFixtures(tt.fixture)
			if err != nil {
				t.Fatal(err)
			}

			if tt.dbError != nil {
				ti.dbclients.SetError(tt.dbError)
			}

			f, err := NewFrontend(ctx, ti.log, ti.env, ti.db, api.APIs, &noop.Noop{}, nil, nil)
			if err != nil {
				t.Fatal(err)
			}

			go f.Run(ctx, nil, nil)

			resp, b, err := ti.request(http.MethodDelete,
				"https://server"+tt.resourceID+"?api-version=2020-04-30",
				nil, nil)
			if err != nil {
				t.Error(err)
			}

			location := resp.Header.Get("Location")
			azureAsyncOperation := resp.Header.Get("Azure-AsyncOperation")
			if tt.wantAsync {
				if !strings.HasPrefix(location, fmt.Sprintf("/subscriptions/%s/providers/microsoft.redhatopenshift/locations/%s/operationresults/", mockSubID, ti.env.Location())) {
					t.Error(location)
				}
				if !strings.HasPrefix(azureAsyncOperation, fmt.Sprintf("/subscriptions/%s/providers/microsoft.redhatopenshift/locations/%s/operationsstatus/", mockSubID, ti.env.Location())) {
					t.Error(azureAsyncOperation)
				}
			} else {
				if location != "" {
					t.Error(location)
				}
				if azureAsyncOperation != "" {
					t.Error(azureAsyncOperation)
				}
			}

			err = validateResponse(resp, b, tt.wantStatusCode, tt.wantError, nil)
			if err != nil {
				t.Error(err)
			}

			ti.dbclients.SetError(nil)
			if tt.wantDocuments != nil {
				tt.wantDocuments(ti.checker)
			}
			errs := ti.checker.Check()
			for _, i := range errs {
				t.Error(i)
			}
		})
	}
}
