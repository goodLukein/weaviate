//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package test

import (
	"fmt"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/weaviate/weaviate/entities/models"
	"github.com/weaviate/weaviate/test/helper"
	graphqlhelper "github.com/weaviate/weaviate/test/helper/graphql"
)

func TestGQLGetTenantObjects(t *testing.T) {
	tenantKey := "tenantName"
	testClass := models.Class{
		Class: "MultiTenantClass",
		MultiTenancyConfig: &models.MultiTenancyConfig{
			Enabled: true,
		},
		Properties: []*models.Property{
			{
				Name:     tenantKey,
				DataType: []string{"string"},
			},
		},
	}
	tenantName := "Tenant1"
	tenantObjects := []*models.Object{
		{
			ID:    "0927a1e0-398e-4e76-91fb-04a7a8f0405c",
			Class: testClass.Class,
			Properties: map[string]interface{}{
				tenantKey: tenantName,
			},
			TenantName: tenantName,
		},
		{
			ID:    "831ae1d0-f441-44b1-bb2a-46548048e26f",
			Class: testClass.Class,
			Properties: map[string]interface{}{
				tenantKey: tenantName,
			},
			TenantName: tenantName,
		},
		{
			ID:    "6f3363e0-c0a0-4618-bf1f-b6cad9cdff59",
			Class: testClass.Class,
			Properties: map[string]interface{}{
				tenantKey: tenantName,
			},
			TenantName: tenantName,
		},
	}

	defer func() {
		helper.DeleteClass(t, testClass.Class)
	}()

	t.Run("setup test data", func(t *testing.T) {
		t.Run("create class with multi-tenancy enabled", func(t *testing.T) {
			helper.CreateClass(t, &testClass)
		})

		t.Run("create tenants", func(t *testing.T) {
			tenants := make([]*models.Tenant, len(tenantObjects))
			for i := range tenants {
				tenants[i] = &models.Tenant{tenantName}
			}
			helper.CreateTenants(t, testClass.Class, tenants)
		})

		t.Run("add tenant objects", func(t *testing.T) {
			helper.CreateObjectsBatch(t, tenantObjects)
		})

		t.Run("get tenant objects", func(t *testing.T) {
			for _, obj := range tenantObjects {
				resp, err := helper.TenantObject(t, obj.Class, obj.ID, tenantName)
				require.Nil(t, err)
				assert.Equal(t, obj.ID, resp.ID)
				assert.Equal(t, obj.Class, resp.Class)
				assert.Equal(t, obj.Properties, resp.Properties)
			}
		})
	})

	t.Run("GQL Get tenant objects", func(t *testing.T) {
		expectedIDs := map[strfmt.UUID]bool{}
		for _, obj := range tenantObjects {
			expectedIDs[obj.ID] = false
		}

		query := fmt.Sprintf(`{Get{%s(tenantKey:%q){_additional{id}}}}`, testClass.Class, tenantName)
		result := graphqlhelper.AssertGraphQL(t, helper.RootAuth, query)
		for _, obj := range result.Get("Get", testClass.Class).AsSlice() {
			id := obj.(map[string]any)["_additional"].(map[string]any)["id"].(string)
			if _, ok := expectedIDs[strfmt.UUID(id)]; ok {
				expectedIDs[strfmt.UUID(id)] = true
			} else {
				t.Fatalf("found unexpected id %q", id)
			}
		}

		for id, found := range expectedIDs {
			if !found {
				t.Fatalf("expected to find id %q, but didn't", id)
			}
		}
	})
}
