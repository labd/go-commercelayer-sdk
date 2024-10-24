/*
Commerce Layer API

Testing FreeShippingPromotionsApiService

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

func Test_api_FreeShippingPromotionsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FreeShippingPromotionsApiService DELETEFreeShippingPromotionsFreeShippingPromotionId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var freeShippingPromotionId interface{}

		httpRes, err := apiClient.FreeShippingPromotionsApi.DELETEFreeShippingPromotionsFreeShippingPromotionId(context.Background(), freeShippingPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FreeShippingPromotionsApiService GETFreeShippingPromotions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FreeShippingPromotionsApi.GETFreeShippingPromotions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FreeShippingPromotionsApiService GETFreeShippingPromotionsFreeShippingPromotionId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var freeShippingPromotionId interface{}

		resp, httpRes, err := apiClient.FreeShippingPromotionsApi.GETFreeShippingPromotionsFreeShippingPromotionId(context.Background(), freeShippingPromotionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FreeShippingPromotionsApiService PATCHFreeShippingPromotionsFreeShippingPromotionId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var freeShippingPromotionId interface{}

		resp, httpRes, err := apiClient.FreeShippingPromotionsApi.PATCHFreeShippingPromotionsFreeShippingPromotionId(context.Background(), freeShippingPromotionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FreeShippingPromotionsApiService POSTFreeShippingPromotions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FreeShippingPromotionsApi.POSTFreeShippingPromotions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}