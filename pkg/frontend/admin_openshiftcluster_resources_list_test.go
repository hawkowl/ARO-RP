package frontend

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/env"
	"github.com/Azure/ARO-RP/pkg/frontend/adminactions"
	"github.com/Azure/ARO-RP/pkg/metrics/noop"
	mock_adminactions "github.com/Azure/ARO-RP/pkg/util/mocks/adminactions"
	testdatabase "github.com/Azure/ARO-RP/test/database"
)

func TestAdminListResourcesList(t *testing.T) {
	mockSubID := "00000000-0000-0000-0000-000000000000"
	mockTenantID := "00000000-0000-0000-0000-000000000000"
	ctx := context.Background()

	type test struct {
		name           string
		resourceID     string
		fixture        func(f *testdatabase.Fixture)
		mocks          func(*test, *mock_adminactions.MockInterface)
		wantStatusCode int
		wantResponse   []byte
		wantError      string
	}

	for _, tt := range []*test{
		{
			name:       "basic coverage",
			resourceID: getResourcePath(mockSubID, "resourceName"),
			fixture: func(f *testdatabase.Fixture) {
				f.AddOpenShiftClusterDocument(&api.OpenShiftClusterDocument{
					Key: strings.ToLower(getResourcePath(mockSubID, "resourceName")),
					OpenShiftCluster: &api.OpenShiftCluster{
						ID: getResourcePath(mockSubID, "resourceName"),
						Properties: api.OpenShiftClusterProperties{
							ClusterProfile: api.ClusterProfile{
								ResourceGroupID: fmt.Sprintf("/subscriptions/%s/resourceGroups/test-cluster", mockSubID),
							},
							MasterProfile: api.MasterProfile{
								SubnetID: fmt.Sprintf("/subscriptions/%s/resourceGroups/test-cluster/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/master", mockSubID),
							},
						},
					},
				})
				f.AddSubscriptionDocument(&api.SubscriptionDocument{
					ID: mockSubID,
					Subscription: &api.Subscription{
						State: api.SubscriptionStateRegistered,
						Properties: &api.SubscriptionProperties{
							TenantID: mockTenantID,
						},
					},
				})
			},
			mocks: func(tt *test, a *mock_adminactions.MockInterface) {

				a.EXPECT().
					ResourcesList(gomock.Any()).
					Return([]byte(`[{"properties":{"dhcpOptions":{"dnsServers":[]}},"id":"/subscriptions/id","type":"Microsoft.Network/virtualNetworks"},{"properties":{"provisioningState":"Succeeded"},"id":"/subscriptions/id","type":"Microsoft.Compute/virtualMachines"},{"id":"/subscriptions/id","name":"storage","type":"Microsoft.Storage/storageAccounts","location":"eastus"}]`), nil)

			},
			wantStatusCode: http.StatusOK,
			wantResponse:   []byte(`[{"properties":{"dhcpOptions":{"dnsServers":[]}},"id":"/subscriptions/id","type":"Microsoft.Network/virtualNetworks"},{"properties":{"provisioningState":"Succeeded"},"id":"/subscriptions/id","type":"Microsoft.Compute/virtualMachines"},{"id":"/subscriptions/id","name":"storage","type":"Microsoft.Storage/storageAccounts","location":"eastus"}]` + "\n"),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ti, err := newTestInfra(t)
			if err != nil {
				t.Fatal(err)
			}
			defer ti.done()

			a := mock_adminactions.NewMockInterface(ti.controller)
			tt.mocks(tt, a)

			err = ti.buildFixtures(tt.fixture)
			if err != nil {
				t.Fatal(err)
			}

			f, err := NewFrontend(ctx, ti.log, ti.env, ti.db, api.APIs, &noop.Noop{}, nil, func(*logrus.Entry, env.Interface, *api.OpenShiftCluster,
				*api.SubscriptionDocument) (adminactions.Interface, error) {
				return a, nil
			})

			if err != nil {
				t.Fatal(err)
			}

			go f.Run(ctx, nil, nil)

			resp, b, err := ti.request(http.MethodGet,
				fmt.Sprintf("https://server/admin/%s/resources", tt.resourceID),
				nil, nil)
			if err != nil {
				t.Fatal(err)
			}

			err = validateResponse(resp, b, tt.wantStatusCode, tt.wantError, tt.wantResponse)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
