/*
Commerce Layer API

Testing ShippingMethodsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api

import (
	"context"
	openapiclient "github.com/incentro-dc/go-commercelayer-sdk/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_api_ShippingMethodsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ShippingMethodsApiService DELETEShippingMethodsShippingMethodId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shippingMethodId interface{}

		httpRes, err := apiClient.ShippingMethodsApi.DELETEShippingMethodsShippingMethodId(context.Background(), shippingMethodId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShippingMethodsApiService GETDeliveryLeadTimeIdShippingMethod", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var deliveryLeadTimeId interface{}

		httpRes, err := apiClient.ShippingMethodsApi.GETDeliveryLeadTimeIdShippingMethod(context.Background(), deliveryLeadTimeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShippingMethodsApiService GETShipmentIdAvailableShippingMethods", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipmentId interface{}

		httpRes, err := apiClient.ShippingMethodsApi.GETShipmentIdAvailableShippingMethods(context.Background(), shipmentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShippingMethodsApiService GETShipmentIdShippingMethod", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipmentId interface{}

		httpRes, err := apiClient.ShippingMethodsApi.GETShipmentIdShippingMethod(context.Background(), shipmentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShippingMethodsApiService GETShippingMethodTierIdShippingMethod", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shippingMethodTierId interface{}

		httpRes, err := apiClient.ShippingMethodsApi.GETShippingMethodTierIdShippingMethod(context.Background(), shippingMethodTierId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShippingMethodsApiService GETShippingMethods", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ShippingMethodsApi.GETShippingMethods(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShippingMethodsApiService GETShippingMethodsShippingMethodId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shippingMethodId interface{}

		resp, httpRes, err := apiClient.ShippingMethodsApi.GETShippingMethodsShippingMethodId(context.Background(), shippingMethodId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShippingMethodsApiService GETShippingWeightTierIdShippingMethod", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shippingWeightTierId interface{}

		httpRes, err := apiClient.ShippingMethodsApi.GETShippingWeightTierIdShippingMethod(context.Background(), shippingWeightTierId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShippingMethodsApiService PATCHShippingMethodsShippingMethodId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shippingMethodId interface{}

		resp, httpRes, err := apiClient.ShippingMethodsApi.PATCHShippingMethodsShippingMethodId(context.Background(), shippingMethodId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShippingMethodsApiService POSTShippingMethods", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ShippingMethodsApi.POSTShippingMethods(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}