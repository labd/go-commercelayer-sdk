/*
Commerce Layer API

Testing BillingInfoValidationRulesApiService

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

func Test_api_BillingInfoValidationRulesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BillingInfoValidationRulesApiService DELETEBillingInfoValidationRulesBillingInfoValidationRuleId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var billingInfoValidationRuleId interface{}

		httpRes, err := apiClient.BillingInfoValidationRulesApi.DELETEBillingInfoValidationRulesBillingInfoValidationRuleId(context.Background(), billingInfoValidationRuleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BillingInfoValidationRulesApiService GETBillingInfoValidationRules", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BillingInfoValidationRulesApi.GETBillingInfoValidationRules(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BillingInfoValidationRulesApiService GETBillingInfoValidationRulesBillingInfoValidationRuleId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var billingInfoValidationRuleId interface{}

		resp, httpRes, err := apiClient.BillingInfoValidationRulesApi.GETBillingInfoValidationRulesBillingInfoValidationRuleId(context.Background(), billingInfoValidationRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BillingInfoValidationRulesApiService PATCHBillingInfoValidationRulesBillingInfoValidationRuleId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var billingInfoValidationRuleId interface{}

		resp, httpRes, err := apiClient.BillingInfoValidationRulesApi.PATCHBillingInfoValidationRulesBillingInfoValidationRuleId(context.Background(), billingInfoValidationRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BillingInfoValidationRulesApiService POSTBillingInfoValidationRules", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BillingInfoValidationRulesApi.POSTBillingInfoValidationRules(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}