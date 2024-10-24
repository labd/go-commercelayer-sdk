/*
Commerce Layer API

Testing AttachmentsApiService

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

func Test_api_AttachmentsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AttachmentsApiService DELETEAttachmentsAttachmentId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var attachmentId interface{}

		httpRes, err := apiClient.AttachmentsApi.DELETEAttachmentsAttachmentId(context.Background(), attachmentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AttachmentsApi.GETAttachments(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETAttachmentsAttachmentId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var attachmentId interface{}

		resp, httpRes, err := apiClient.AttachmentsApi.GETAttachmentsAttachmentId(context.Background(), attachmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETAvalaraAccountIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var avalaraAccountId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETAvalaraAccountIdAttachments(context.Background(), avalaraAccountId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETBillingInfoValidationRuleIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var billingInfoValidationRuleId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETBillingInfoValidationRuleIdAttachments(context.Background(), billingInfoValidationRuleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETBingGeocoderIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var bingGeocoderId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETBingGeocoderIdAttachments(context.Background(), bingGeocoderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETBundleIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var bundleId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETBundleIdAttachments(context.Background(), bundleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETCarrierAccountIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var carrierAccountId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETCarrierAccountIdAttachments(context.Background(), carrierAccountId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETCouponRecipientIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var couponRecipientId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETCouponRecipientIdAttachments(context.Background(), couponRecipientId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETCustomerGroupIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customerGroupId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETCustomerGroupIdAttachments(context.Background(), customerGroupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETCustomerIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customerId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETCustomerIdAttachments(context.Background(), customerId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETDeliveryLeadTimeIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var deliveryLeadTimeId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETDeliveryLeadTimeIdAttachments(context.Background(), deliveryLeadTimeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETExternalPromotionIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var externalPromotionId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETExternalPromotionIdAttachments(context.Background(), externalPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETExternalTaxCalculatorIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var externalTaxCalculatorId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETExternalTaxCalculatorIdAttachments(context.Background(), externalTaxCalculatorId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETFixedAmountPromotionIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fixedAmountPromotionId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETFixedAmountPromotionIdAttachments(context.Background(), fixedAmountPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETFixedPricePromotionIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fixedPricePromotionId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETFixedPricePromotionIdAttachments(context.Background(), fixedPricePromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETFreeGiftPromotionIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var freeGiftPromotionId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETFreeGiftPromotionIdAttachments(context.Background(), freeGiftPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETFreeShippingPromotionIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var freeShippingPromotionId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETFreeShippingPromotionIdAttachments(context.Background(), freeShippingPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETGeocoderIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var geocoderId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETGeocoderIdAttachments(context.Background(), geocoderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETGiftCardIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var giftCardId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETGiftCardIdAttachments(context.Background(), giftCardId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETGiftCardRecipientIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var giftCardRecipientId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETGiftCardRecipientIdAttachments(context.Background(), giftCardRecipientId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETGoogleGeocoderIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var googleGeocoderId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETGoogleGeocoderIdAttachments(context.Background(), googleGeocoderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETInventoryModelIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var inventoryModelId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETInventoryModelIdAttachments(context.Background(), inventoryModelId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETManualTaxCalculatorIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var manualTaxCalculatorId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETManualTaxCalculatorIdAttachments(context.Background(), manualTaxCalculatorId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETMarketIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var marketId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETMarketIdAttachments(context.Background(), marketId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETMerchantIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var merchantId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETMerchantIdAttachments(context.Background(), merchantId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETOrderIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orderId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETOrderIdAttachments(context.Background(), orderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETOrderValidationRuleIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orderValidationRuleId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETOrderValidationRuleIdAttachments(context.Background(), orderValidationRuleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETPackageIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var packageId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETPackageIdAttachments(context.Background(), packageId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETParcelIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var parcelId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETParcelIdAttachments(context.Background(), parcelId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETPaymentMethodIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var paymentMethodId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETPaymentMethodIdAttachments(context.Background(), paymentMethodId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETPercentageDiscountPromotionIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var percentageDiscountPromotionId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETPercentageDiscountPromotionIdAttachments(context.Background(), percentageDiscountPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETPriceFrequencyTierIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var priceFrequencyTierId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETPriceFrequencyTierIdAttachments(context.Background(), priceFrequencyTierId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETPriceIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var priceId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETPriceIdAttachments(context.Background(), priceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETPriceListIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var priceListId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETPriceListIdAttachments(context.Background(), priceListId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETPriceTierIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var priceTierId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETPriceTierIdAttachments(context.Background(), priceTierId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETPriceVolumeTierIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var priceVolumeTierId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETPriceVolumeTierIdAttachments(context.Background(), priceVolumeTierId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETPromotionIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var promotionId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETPromotionIdAttachments(context.Background(), promotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETReturnIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var returnId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETReturnIdAttachments(context.Background(), returnId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETShipmentIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipmentId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETShipmentIdAttachments(context.Background(), shipmentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETShippingCategoryIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shippingCategoryId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETShippingCategoryIdAttachments(context.Background(), shippingCategoryId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETShippingMethodIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shippingMethodId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETShippingMethodIdAttachments(context.Background(), shippingMethodId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETShippingMethodTierIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shippingMethodTierId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETShippingMethodTierIdAttachments(context.Background(), shippingMethodTierId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETShippingWeightTierIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shippingWeightTierId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETShippingWeightTierIdAttachments(context.Background(), shippingWeightTierId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETShippingZoneIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shippingZoneId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETShippingZoneIdAttachments(context.Background(), shippingZoneId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETSkuIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var skuId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETSkuIdAttachments(context.Background(), skuId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETSkuListIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var skuListId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETSkuListIdAttachments(context.Background(), skuListId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETSkuOptionIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var skuOptionId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETSkuOptionIdAttachments(context.Background(), skuOptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETStockItemIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var stockItemId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETStockItemIdAttachments(context.Background(), stockItemId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETStockLocationIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var stockLocationId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETStockLocationIdAttachments(context.Background(), stockLocationId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETSubscriptionModelIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var subscriptionModelId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETSubscriptionModelIdAttachments(context.Background(), subscriptionModelId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETTaxCalculatorIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taxCalculatorId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETTaxCalculatorIdAttachments(context.Background(), taxCalculatorId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETTaxCategoryIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taxCategoryId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETTaxCategoryIdAttachments(context.Background(), taxCategoryId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService GETTaxjarAccountIdAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taxjarAccountId interface{}

		httpRes, err := apiClient.AttachmentsApi.GETTaxjarAccountIdAttachments(context.Background(), taxjarAccountId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService PATCHAttachmentsAttachmentId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var attachmentId interface{}

		resp, httpRes, err := apiClient.AttachmentsApi.PATCHAttachmentsAttachmentId(context.Background(), attachmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsApiService POSTAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AttachmentsApi.POSTAttachments(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}