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
	v20200430 "github.com/Azure/ARO-RP/pkg/api/v20200430"
	"github.com/Azure/ARO-RP/pkg/database/cosmosdb"
	"github.com/Azure/ARO-RP/pkg/metrics/noop"
	testdatabase "github.com/Azure/ARO-RP/test/database"
)

func TestGetAsyncOperationResult(t *testing.T) {
	ctx := context.Background()

	mockSubID := "00000000-0000-0000-0000-000000000000"
	mockClusterDocKey := "22222222-2222-2222-2222-222222222222"
	mockOpID := "11111111-1111-1111-1111-111111111111"

	type test struct {
		name           string
		fixture        func(f *testdatabase.Fixture)
		dbError        error
		wantStatusCode int
		wantAsync      bool
		wantResponse   *v20200430.OpenShiftCluster
		wantError      string
	}

	for _, tt := range []*test{
		{
			name: "operation and cluster exist in db - final result is available with content",
			fixture: func(f *testdatabase.Fixture) {
				clusterDoc := &api.OpenShiftClusterDocument{
					ID:  mockClusterDocKey,
					Key: strings.ToLower(getResourcePath(mockSubID, "fakeClusterID")),
					OpenShiftCluster: &api.OpenShiftCluster{
						ID:   getResourcePath(mockSubID, "fakeClusterID"),
						Name: "resourceName",
						Type: "Microsoft.RedHatOpenShift/openshiftClusters",
						Properties: api.OpenShiftClusterProperties{
							ClusterProfile: api.ClusterProfile{
								PullSecret: "{}",
							},
							ServicePrincipalProfile: api.ServicePrincipalProfile{
								ClientSecret: "clientSecret",
							},
						},
					},
				}
				asyncDoc := &api.AsyncOperationDocument{
					ID:                  mockOpID,
					OpenShiftClusterKey: strings.ToLower(getResourcePath(mockSubID, "fakeClusterID")),
					OpenShiftCluster:    clusterDoc.OpenShiftCluster,
				}

				f.AddOpenShiftClusterDocument(clusterDoc)
				f.AddAsyncOperationDocument(asyncDoc)
			},
			wantStatusCode: http.StatusOK,
			wantResponse: &v20200430.OpenShiftCluster{
				ID:   getResourcePath(mockSubID, "fakeClusterID"),
				Name: "resourceName",
				Type: "Microsoft.RedHatOpenShift/openshiftClusters",
			},
		},
		{
			name: "operation exists in db, but no cluster - final result is available with no content",
			fixture: func(f *testdatabase.Fixture) {
				f.AddAsyncOperationDocument(&api.AsyncOperationDocument{
					ID:                  mockOpID,
					OpenShiftClusterKey: strings.ToLower(getResourcePath(mockSubID, "fakeClusterID")),
				})
			},
			wantStatusCode: http.StatusNoContent,
		},
		{
			name: "operation and cluster exist in db - final result is not yet available",
			fixture: func(f *testdatabase.Fixture) {
				f.AddAsyncOperationDocument(&api.AsyncOperationDocument{
					ID:                  mockOpID,
					OpenShiftClusterKey: strings.ToLower(getResourcePath(mockSubID, "fakeClusterID")),
				})
				f.AddOpenShiftClusterDocument(&api.OpenShiftClusterDocument{
					Key:              strings.ToLower(getResourcePath(mockSubID, "fakeClusterID")),
					AsyncOperationID: mockOpID,
				})
			},
			wantAsync:      true,
			wantStatusCode: http.StatusAccepted,
		},
		{
			name:           "operation not found in db",
			wantStatusCode: http.StatusNotFound,
			wantError:      `404: NotFound: : The entity was not found.`,
		},
		{
			name:           "internal error",
			dbError:        &cosmosdb.Error{Code: "500", Message: "oops"},
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

			if tt.dbError != nil {
				ti.dbclients.MakeUnavailable(tt.dbError)
			}

			err = ti.buildFixtures(tt.fixture)
			if err != nil {
				t.Fatal(err)
			}

			f, err := NewFrontend(ctx, ti.log, ti.env, ti.db, api.APIs, &noop.Noop{}, nil, nil)
			if err != nil {
				t.Fatal(err)
			}

			go f.Run(ctx, nil, nil)

			referer := fmt.Sprintf("/subscriptions/%s/providers/microsoft.redhatopenshift/locations/%s/operationresults/%s", mockSubID, ti.env.Location(), mockOpID)

			resp, b, err := ti.request(http.MethodGet,
				fmt.Sprintf("https://server/subscriptions/%s/providers/Microsoft.RedHatOpenShift/locations/%s/operationresults/%s?api-version=2020-04-30", mockSubID, ti.env.Location(), mockOpID),
				http.Header{
					"Content-Type": []string{"application/json"},
					"Referer":      []string{referer},
				}, nil)
			if err != nil {
				t.Fatal(err)
			}

			location := resp.Header.Get("Location")
			if tt.wantAsync {
				if location != referer {
					t.Error(location)
				}
			} else {
				if location != "" {
					t.Error(location)
				}
			}

			err = validateResponse(resp, b, tt.wantStatusCode, tt.wantError, tt.wantResponse)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
