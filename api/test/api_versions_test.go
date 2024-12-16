/*
Commerce Layer API

Testing VersionsApiService

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

func Test_api_VersionsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test VersionsApiService GETAddressIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var addressId interface{}

		httpRes, err := apiClient.VersionsApi.GETAddressIdVersions(context.Background(), addressId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETAdjustmentIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var adjustmentId interface{}

		httpRes, err := apiClient.VersionsApi.GETAdjustmentIdVersions(context.Background(), adjustmentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETAdyenGatewayIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var adyenGatewayId interface{}

		httpRes, err := apiClient.VersionsApi.GETAdyenGatewayIdVersions(context.Background(), adyenGatewayId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETAdyenPaymentIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var adyenPaymentId interface{}

		httpRes, err := apiClient.VersionsApi.GETAdyenPaymentIdVersions(context.Background(), adyenPaymentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETAuthorizationIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var authorizationId interface{}

		httpRes, err := apiClient.VersionsApi.GETAuthorizationIdVersions(context.Background(), authorizationId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETAvalaraAccountIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var avalaraAccountId interface{}

		httpRes, err := apiClient.VersionsApi.GETAvalaraAccountIdVersions(context.Background(), avalaraAccountId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETAxerveGatewayIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var axerveGatewayId interface{}

		httpRes, err := apiClient.VersionsApi.GETAxerveGatewayIdVersions(context.Background(), axerveGatewayId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETAxervePaymentIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var axervePaymentId interface{}

		httpRes, err := apiClient.VersionsApi.GETAxervePaymentIdVersions(context.Background(), axervePaymentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETBraintreeGatewayIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var braintreeGatewayId interface{}

		httpRes, err := apiClient.VersionsApi.GETBraintreeGatewayIdVersions(context.Background(), braintreeGatewayId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETBraintreePaymentIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var braintreePaymentId interface{}

		httpRes, err := apiClient.VersionsApi.GETBraintreePaymentIdVersions(context.Background(), braintreePaymentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETBundleIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var bundleId interface{}

		httpRes, err := apiClient.VersionsApi.GETBundleIdVersions(context.Background(), bundleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETBuyXPayYPromotionIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var buyXPayYPromotionId interface{}

		httpRes, err := apiClient.VersionsApi.GETBuyXPayYPromotionIdVersions(context.Background(), buyXPayYPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETCaptureIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var captureId interface{}

		httpRes, err := apiClient.VersionsApi.GETCaptureIdVersions(context.Background(), captureId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETCarrierAccountIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var carrierAccountId interface{}

		httpRes, err := apiClient.VersionsApi.GETCarrierAccountIdVersions(context.Background(), carrierAccountId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETCheckoutComGatewayIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var checkoutComGatewayId interface{}

		httpRes, err := apiClient.VersionsApi.GETCheckoutComGatewayIdVersions(context.Background(), checkoutComGatewayId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETCheckoutComPaymentIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var checkoutComPaymentId interface{}

		httpRes, err := apiClient.VersionsApi.GETCheckoutComPaymentIdVersions(context.Background(), checkoutComPaymentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETCleanupIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cleanupId interface{}

		httpRes, err := apiClient.VersionsApi.GETCleanupIdVersions(context.Background(), cleanupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETCouponCodesPromotionRuleIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var couponCodesPromotionRuleId interface{}

		httpRes, err := apiClient.VersionsApi.GETCouponCodesPromotionRuleIdVersions(context.Background(), couponCodesPromotionRuleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETCouponIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var couponId interface{}

		httpRes, err := apiClient.VersionsApi.GETCouponIdVersions(context.Background(), couponId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETCouponRecipientIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var couponRecipientId interface{}

		httpRes, err := apiClient.VersionsApi.GETCouponRecipientIdVersions(context.Background(), couponRecipientId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETCustomPromotionRuleIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customPromotionRuleId interface{}

		httpRes, err := apiClient.VersionsApi.GETCustomPromotionRuleIdVersions(context.Background(), customPromotionRuleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETCustomerAddressIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customerAddressId interface{}

		httpRes, err := apiClient.VersionsApi.GETCustomerAddressIdVersions(context.Background(), customerAddressId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETCustomerGroupIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customerGroupId interface{}

		httpRes, err := apiClient.VersionsApi.GETCustomerGroupIdVersions(context.Background(), customerGroupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETCustomerPaymentSourceIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customerPaymentSourceId interface{}

		httpRes, err := apiClient.VersionsApi.GETCustomerPaymentSourceIdVersions(context.Background(), customerPaymentSourceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETCustomerSubscriptionIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customerSubscriptionId interface{}

		httpRes, err := apiClient.VersionsApi.GETCustomerSubscriptionIdVersions(context.Background(), customerSubscriptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETDeliveryLeadTimeIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var deliveryLeadTimeId interface{}

		httpRes, err := apiClient.VersionsApi.GETDeliveryLeadTimeIdVersions(context.Background(), deliveryLeadTimeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETExternalGatewayIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var externalGatewayId interface{}

		httpRes, err := apiClient.VersionsApi.GETExternalGatewayIdVersions(context.Background(), externalGatewayId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETExternalPaymentIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var externalPaymentId interface{}

		httpRes, err := apiClient.VersionsApi.GETExternalPaymentIdVersions(context.Background(), externalPaymentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETExternalPromotionIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var externalPromotionId interface{}

		httpRes, err := apiClient.VersionsApi.GETExternalPromotionIdVersions(context.Background(), externalPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETExternalTaxCalculatorIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var externalTaxCalculatorId interface{}

		httpRes, err := apiClient.VersionsApi.GETExternalTaxCalculatorIdVersions(context.Background(), externalTaxCalculatorId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETFixedAmountPromotionIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fixedAmountPromotionId interface{}

		httpRes, err := apiClient.VersionsApi.GETFixedAmountPromotionIdVersions(context.Background(), fixedAmountPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETFixedPricePromotionIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fixedPricePromotionId interface{}

		httpRes, err := apiClient.VersionsApi.GETFixedPricePromotionIdVersions(context.Background(), fixedPricePromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETFlexPromotionIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var flexPromotionId interface{}

		httpRes, err := apiClient.VersionsApi.GETFlexPromotionIdVersions(context.Background(), flexPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETFreeGiftPromotionIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var freeGiftPromotionId interface{}

		httpRes, err := apiClient.VersionsApi.GETFreeGiftPromotionIdVersions(context.Background(), freeGiftPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETFreeShippingPromotionIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var freeShippingPromotionId interface{}

		httpRes, err := apiClient.VersionsApi.GETFreeShippingPromotionIdVersions(context.Background(), freeShippingPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETGiftCardIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var giftCardId interface{}

		httpRes, err := apiClient.VersionsApi.GETGiftCardIdVersions(context.Background(), giftCardId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETGiftCardRecipientIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var giftCardRecipientId interface{}

		httpRes, err := apiClient.VersionsApi.GETGiftCardRecipientIdVersions(context.Background(), giftCardRecipientId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETInStockSubscriptionIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var inStockSubscriptionId interface{}

		httpRes, err := apiClient.VersionsApi.GETInStockSubscriptionIdVersions(context.Background(), inStockSubscriptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETInventoryModelIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var inventoryModelId interface{}

		httpRes, err := apiClient.VersionsApi.GETInventoryModelIdVersions(context.Background(), inventoryModelId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETInventoryReturnLocationIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var inventoryReturnLocationId interface{}

		httpRes, err := apiClient.VersionsApi.GETInventoryReturnLocationIdVersions(context.Background(), inventoryReturnLocationId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETInventoryStockLocationIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var inventoryStockLocationId interface{}

		httpRes, err := apiClient.VersionsApi.GETInventoryStockLocationIdVersions(context.Background(), inventoryStockLocationId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETKlarnaGatewayIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var klarnaGatewayId interface{}

		httpRes, err := apiClient.VersionsApi.GETKlarnaGatewayIdVersions(context.Background(), klarnaGatewayId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETKlarnaPaymentIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var klarnaPaymentId interface{}

		httpRes, err := apiClient.VersionsApi.GETKlarnaPaymentIdVersions(context.Background(), klarnaPaymentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETManualGatewayIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var manualGatewayId interface{}

		httpRes, err := apiClient.VersionsApi.GETManualGatewayIdVersions(context.Background(), manualGatewayId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETManualTaxCalculatorIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var manualTaxCalculatorId interface{}

		httpRes, err := apiClient.VersionsApi.GETManualTaxCalculatorIdVersions(context.Background(), manualTaxCalculatorId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETMarketIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var marketId interface{}

		httpRes, err := apiClient.VersionsApi.GETMarketIdVersions(context.Background(), marketId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETMerchantIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var merchantId interface{}

		httpRes, err := apiClient.VersionsApi.GETMerchantIdVersions(context.Background(), merchantId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETOrderAmountPromotionRuleIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orderAmountPromotionRuleId interface{}

		httpRes, err := apiClient.VersionsApi.GETOrderAmountPromotionRuleIdVersions(context.Background(), orderAmountPromotionRuleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETOrderIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orderId interface{}

		httpRes, err := apiClient.VersionsApi.GETOrderIdVersions(context.Background(), orderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETOrderSubscriptionIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orderSubscriptionId interface{}

		httpRes, err := apiClient.VersionsApi.GETOrderSubscriptionIdVersions(context.Background(), orderSubscriptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETPackageIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var packageId interface{}

		httpRes, err := apiClient.VersionsApi.GETPackageIdVersions(context.Background(), packageId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETParcelIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var parcelId interface{}

		httpRes, err := apiClient.VersionsApi.GETParcelIdVersions(context.Background(), parcelId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETParcelLineItemIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var parcelLineItemId interface{}

		httpRes, err := apiClient.VersionsApi.GETParcelLineItemIdVersions(context.Background(), parcelLineItemId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETPaymentGatewayIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var paymentGatewayId interface{}

		httpRes, err := apiClient.VersionsApi.GETPaymentGatewayIdVersions(context.Background(), paymentGatewayId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETPaymentMethodIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var paymentMethodId interface{}

		httpRes, err := apiClient.VersionsApi.GETPaymentMethodIdVersions(context.Background(), paymentMethodId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETPaypalGatewayIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var paypalGatewayId interface{}

		httpRes, err := apiClient.VersionsApi.GETPaypalGatewayIdVersions(context.Background(), paypalGatewayId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETPaypalPaymentIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var paypalPaymentId interface{}

		httpRes, err := apiClient.VersionsApi.GETPaypalPaymentIdVersions(context.Background(), paypalPaymentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETPercentageDiscountPromotionIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var percentageDiscountPromotionId interface{}

		httpRes, err := apiClient.VersionsApi.GETPercentageDiscountPromotionIdVersions(context.Background(), percentageDiscountPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETPriceFrequencyTierIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var priceFrequencyTierId interface{}

		httpRes, err := apiClient.VersionsApi.GETPriceFrequencyTierIdVersions(context.Background(), priceFrequencyTierId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETPriceIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var priceId interface{}

		httpRes, err := apiClient.VersionsApi.GETPriceIdVersions(context.Background(), priceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETPriceListIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var priceListId interface{}

		httpRes, err := apiClient.VersionsApi.GETPriceListIdVersions(context.Background(), priceListId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETPriceListSchedulerIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var priceListSchedulerId interface{}

		httpRes, err := apiClient.VersionsApi.GETPriceListSchedulerIdVersions(context.Background(), priceListSchedulerId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETPriceTierIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var priceTierId interface{}

		httpRes, err := apiClient.VersionsApi.GETPriceTierIdVersions(context.Background(), priceTierId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETPriceVolumeTierIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var priceVolumeTierId interface{}

		httpRes, err := apiClient.VersionsApi.GETPriceVolumeTierIdVersions(context.Background(), priceVolumeTierId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETPromotionIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var promotionId interface{}

		httpRes, err := apiClient.VersionsApi.GETPromotionIdVersions(context.Background(), promotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETPromotionRuleIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var promotionRuleId interface{}

		httpRes, err := apiClient.VersionsApi.GETPromotionRuleIdVersions(context.Background(), promotionRuleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETRefundIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var refundId interface{}

		httpRes, err := apiClient.VersionsApi.GETRefundIdVersions(context.Background(), refundId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETReservedStockIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var reservedStockId interface{}

		httpRes, err := apiClient.VersionsApi.GETReservedStockIdVersions(context.Background(), reservedStockId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETReturnIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var returnId interface{}

		httpRes, err := apiClient.VersionsApi.GETReturnIdVersions(context.Background(), returnId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETSatispayGatewayIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var satispayGatewayId interface{}

		httpRes, err := apiClient.VersionsApi.GETSatispayGatewayIdVersions(context.Background(), satispayGatewayId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETSatispayPaymentIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var satispayPaymentId interface{}

		httpRes, err := apiClient.VersionsApi.GETSatispayPaymentIdVersions(context.Background(), satispayPaymentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETShipmentIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipmentId interface{}

		httpRes, err := apiClient.VersionsApi.GETShipmentIdVersions(context.Background(), shipmentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETShippingCategoryIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shippingCategoryId interface{}

		httpRes, err := apiClient.VersionsApi.GETShippingCategoryIdVersions(context.Background(), shippingCategoryId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETShippingMethodIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shippingMethodId interface{}

		httpRes, err := apiClient.VersionsApi.GETShippingMethodIdVersions(context.Background(), shippingMethodId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETShippingMethodTierIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shippingMethodTierId interface{}

		httpRes, err := apiClient.VersionsApi.GETShippingMethodTierIdVersions(context.Background(), shippingMethodTierId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETShippingWeightTierIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shippingWeightTierId interface{}

		httpRes, err := apiClient.VersionsApi.GETShippingWeightTierIdVersions(context.Background(), shippingWeightTierId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETShippingZoneIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shippingZoneId interface{}

		httpRes, err := apiClient.VersionsApi.GETShippingZoneIdVersions(context.Background(), shippingZoneId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETSkuIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var skuId interface{}

		httpRes, err := apiClient.VersionsApi.GETSkuIdVersions(context.Background(), skuId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETSkuListIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var skuListId interface{}

		httpRes, err := apiClient.VersionsApi.GETSkuListIdVersions(context.Background(), skuListId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETSkuListItemIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var skuListItemId interface{}

		httpRes, err := apiClient.VersionsApi.GETSkuListItemIdVersions(context.Background(), skuListItemId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETSkuListPromotionRuleIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var skuListPromotionRuleId interface{}

		httpRes, err := apiClient.VersionsApi.GETSkuListPromotionRuleIdVersions(context.Background(), skuListPromotionRuleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETSkuOptionIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var skuOptionId interface{}

		httpRes, err := apiClient.VersionsApi.GETSkuOptionIdVersions(context.Background(), skuOptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETStockItemIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var stockItemId interface{}

		httpRes, err := apiClient.VersionsApi.GETStockItemIdVersions(context.Background(), stockItemId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETStockLineItemIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var stockLineItemId interface{}

		httpRes, err := apiClient.VersionsApi.GETStockLineItemIdVersions(context.Background(), stockLineItemId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETStockLocationIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var stockLocationId interface{}

		httpRes, err := apiClient.VersionsApi.GETStockLocationIdVersions(context.Background(), stockLocationId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETStockTransferIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var stockTransferId interface{}

		httpRes, err := apiClient.VersionsApi.GETStockTransferIdVersions(context.Background(), stockTransferId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETStoreIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var storeId interface{}

		httpRes, err := apiClient.VersionsApi.GETStoreIdVersions(context.Background(), storeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETStripeGatewayIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var stripeGatewayId interface{}

		httpRes, err := apiClient.VersionsApi.GETStripeGatewayIdVersions(context.Background(), stripeGatewayId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETStripePaymentIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var stripePaymentId interface{}

		httpRes, err := apiClient.VersionsApi.GETStripePaymentIdVersions(context.Background(), stripePaymentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETTaxCalculatorIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taxCalculatorId interface{}

		httpRes, err := apiClient.VersionsApi.GETTaxCalculatorIdVersions(context.Background(), taxCalculatorId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETTaxCategoryIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taxCategoryId interface{}

		httpRes, err := apiClient.VersionsApi.GETTaxCategoryIdVersions(context.Background(), taxCategoryId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETTaxRuleIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taxRuleId interface{}

		httpRes, err := apiClient.VersionsApi.GETTaxRuleIdVersions(context.Background(), taxRuleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETTaxjarAccountIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taxjarAccountId interface{}

		httpRes, err := apiClient.VersionsApi.GETTaxjarAccountIdVersions(context.Background(), taxjarAccountId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETTransactionIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var transactionId interface{}

		httpRes, err := apiClient.VersionsApi.GETTransactionIdVersions(context.Background(), transactionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.VersionsApi.GETVersions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETVersionsVersionId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var versionId interface{}

		resp, httpRes, err := apiClient.VersionsApi.GETVersionsVersionId(context.Background(), versionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETVoidIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var voidId interface{}

		httpRes, err := apiClient.VersionsApi.GETVoidIdVersions(context.Background(), voidId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETWebhookIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var webhookId interface{}

		httpRes, err := apiClient.VersionsApi.GETWebhookIdVersions(context.Background(), webhookId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GETWireTransferIdVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var wireTransferId interface{}

		httpRes, err := apiClient.VersionsApi.GETWireTransferIdVersions(context.Background(), wireTransferId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
