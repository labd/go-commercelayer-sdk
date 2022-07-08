# \CouponCodesPromotionRulesApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETECouponCodesPromotionRulesCouponCodesPromotionRuleId**](CouponCodesPromotionRulesApi.md#DELETECouponCodesPromotionRulesCouponCodesPromotionRuleId) | **Delete** /coupon_codes_promotion_rules/{couponCodesPromotionRuleId} | Delete a coupon codes promotion rule
[**GETCouponCodesPromotionRules**](CouponCodesPromotionRulesApi.md#GETCouponCodesPromotionRules) | **Get** /coupon_codes_promotion_rules | List all coupon codes promotion rules
[**GETCouponCodesPromotionRulesCouponCodesPromotionRuleId**](CouponCodesPromotionRulesApi.md#GETCouponCodesPromotionRulesCouponCodesPromotionRuleId) | **Get** /coupon_codes_promotion_rules/{couponCodesPromotionRuleId} | Retrieve a coupon codes promotion rule
[**GETCouponIdPromotionRule**](CouponCodesPromotionRulesApi.md#GETCouponIdPromotionRule) | **Get** /coupons/{couponId}/promotion_rule | Retrieve the promotion rule associated to the coupon
[**GETExternalPromotionIdCouponCodesPromotionRule**](CouponCodesPromotionRulesApi.md#GETExternalPromotionIdCouponCodesPromotionRule) | **Get** /external_promotions/{externalPromotionId}/coupon_codes_promotion_rule | Retrieve the coupon codes promotion rule associated to the external promotion
[**GETFixedAmountPromotionIdCouponCodesPromotionRule**](CouponCodesPromotionRulesApi.md#GETFixedAmountPromotionIdCouponCodesPromotionRule) | **Get** /fixed_amount_promotions/{fixedAmountPromotionId}/coupon_codes_promotion_rule | Retrieve the coupon codes promotion rule associated to the fixed amount promotion
[**GETFixedPricePromotionIdCouponCodesPromotionRule**](CouponCodesPromotionRulesApi.md#GETFixedPricePromotionIdCouponCodesPromotionRule) | **Get** /fixed_price_promotions/{fixedPricePromotionId}/coupon_codes_promotion_rule | Retrieve the coupon codes promotion rule associated to the fixed price promotion
[**GETFreeGiftPromotionIdCouponCodesPromotionRule**](CouponCodesPromotionRulesApi.md#GETFreeGiftPromotionIdCouponCodesPromotionRule) | **Get** /free_gift_promotions/{freeGiftPromotionId}/coupon_codes_promotion_rule | Retrieve the coupon codes promotion rule associated to the free gift promotion
[**GETFreeShippingPromotionIdCouponCodesPromotionRule**](CouponCodesPromotionRulesApi.md#GETFreeShippingPromotionIdCouponCodesPromotionRule) | **Get** /free_shipping_promotions/{freeShippingPromotionId}/coupon_codes_promotion_rule | Retrieve the coupon codes promotion rule associated to the free shipping promotion
[**GETPercentageDiscountPromotionIdCouponCodesPromotionRule**](CouponCodesPromotionRulesApi.md#GETPercentageDiscountPromotionIdCouponCodesPromotionRule) | **Get** /percentage_discount_promotions/{percentageDiscountPromotionId}/coupon_codes_promotion_rule | Retrieve the coupon codes promotion rule associated to the percentage discount promotion
[**GETPromotionIdCouponCodesPromotionRule**](CouponCodesPromotionRulesApi.md#GETPromotionIdCouponCodesPromotionRule) | **Get** /promotions/{promotionId}/coupon_codes_promotion_rule | Retrieve the coupon codes promotion rule associated to the promotion
[**PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId**](CouponCodesPromotionRulesApi.md#PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId) | **Patch** /coupon_codes_promotion_rules/{couponCodesPromotionRuleId} | Update a coupon codes promotion rule
[**POSTCouponCodesPromotionRules**](CouponCodesPromotionRulesApi.md#POSTCouponCodesPromotionRules) | **Post** /coupon_codes_promotion_rules | Create a coupon codes promotion rule



## DELETECouponCodesPromotionRulesCouponCodesPromotionRuleId

> DELETECouponCodesPromotionRulesCouponCodesPromotionRuleId(ctx, couponCodesPromotionRuleId).Execute()

Delete a coupon codes promotion rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    couponCodesPromotionRuleId := "couponCodesPromotionRuleId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponCodesPromotionRulesApi.DELETECouponCodesPromotionRulesCouponCodesPromotionRuleId(context.Background(), couponCodesPromotionRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponCodesPromotionRulesApi.DELETECouponCodesPromotionRulesCouponCodesPromotionRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponCodesPromotionRuleId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETECouponCodesPromotionRulesCouponCodesPromotionRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCouponCodesPromotionRules

> GETCouponCodesPromotionRules(ctx).Execute()

List all coupon codes promotion rules



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponCodesPromotionRulesApi.GETCouponCodesPromotionRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponCodesPromotionRulesApi.GETCouponCodesPromotionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETCouponCodesPromotionRulesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCouponCodesPromotionRulesCouponCodesPromotionRuleId

> CouponCodesPromotionRule GETCouponCodesPromotionRulesCouponCodesPromotionRuleId(ctx, couponCodesPromotionRuleId).Execute()

Retrieve a coupon codes promotion rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    couponCodesPromotionRuleId := "couponCodesPromotionRuleId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponCodesPromotionRulesApi.GETCouponCodesPromotionRulesCouponCodesPromotionRuleId(context.Background(), couponCodesPromotionRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponCodesPromotionRulesApi.GETCouponCodesPromotionRulesCouponCodesPromotionRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCouponCodesPromotionRulesCouponCodesPromotionRuleId`: CouponCodesPromotionRule
    fmt.Fprintf(os.Stdout, "Response from `CouponCodesPromotionRulesApi.GETCouponCodesPromotionRulesCouponCodesPromotionRuleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponCodesPromotionRuleId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCouponCodesPromotionRulesCouponCodesPromotionRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CouponCodesPromotionRule**](CouponCodesPromotionRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCouponIdPromotionRule

> GETCouponIdPromotionRule(ctx, couponId).Execute()

Retrieve the promotion rule associated to the coupon



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    couponId := "couponId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponCodesPromotionRulesApi.GETCouponIdPromotionRule(context.Background(), couponId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponCodesPromotionRulesApi.GETCouponIdPromotionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCouponIdPromotionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETExternalPromotionIdCouponCodesPromotionRule

> GETExternalPromotionIdCouponCodesPromotionRule(ctx, externalPromotionId).Execute()

Retrieve the coupon codes promotion rule associated to the external promotion



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    externalPromotionId := "externalPromotionId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponCodesPromotionRulesApi.GETExternalPromotionIdCouponCodesPromotionRule(context.Background(), externalPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponCodesPromotionRulesApi.GETExternalPromotionIdCouponCodesPromotionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalPromotionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETExternalPromotionIdCouponCodesPromotionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETFixedAmountPromotionIdCouponCodesPromotionRule

> GETFixedAmountPromotionIdCouponCodesPromotionRule(ctx, fixedAmountPromotionId).Execute()

Retrieve the coupon codes promotion rule associated to the fixed amount promotion



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    fixedAmountPromotionId := "fixedAmountPromotionId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponCodesPromotionRulesApi.GETFixedAmountPromotionIdCouponCodesPromotionRule(context.Background(), fixedAmountPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponCodesPromotionRulesApi.GETFixedAmountPromotionIdCouponCodesPromotionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fixedAmountPromotionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETFixedAmountPromotionIdCouponCodesPromotionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETFixedPricePromotionIdCouponCodesPromotionRule

> GETFixedPricePromotionIdCouponCodesPromotionRule(ctx, fixedPricePromotionId).Execute()

Retrieve the coupon codes promotion rule associated to the fixed price promotion



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    fixedPricePromotionId := "fixedPricePromotionId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponCodesPromotionRulesApi.GETFixedPricePromotionIdCouponCodesPromotionRule(context.Background(), fixedPricePromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponCodesPromotionRulesApi.GETFixedPricePromotionIdCouponCodesPromotionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fixedPricePromotionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETFixedPricePromotionIdCouponCodesPromotionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETFreeGiftPromotionIdCouponCodesPromotionRule

> GETFreeGiftPromotionIdCouponCodesPromotionRule(ctx, freeGiftPromotionId).Execute()

Retrieve the coupon codes promotion rule associated to the free gift promotion



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    freeGiftPromotionId := "freeGiftPromotionId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponCodesPromotionRulesApi.GETFreeGiftPromotionIdCouponCodesPromotionRule(context.Background(), freeGiftPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponCodesPromotionRulesApi.GETFreeGiftPromotionIdCouponCodesPromotionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**freeGiftPromotionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETFreeGiftPromotionIdCouponCodesPromotionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETFreeShippingPromotionIdCouponCodesPromotionRule

> GETFreeShippingPromotionIdCouponCodesPromotionRule(ctx, freeShippingPromotionId).Execute()

Retrieve the coupon codes promotion rule associated to the free shipping promotion



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    freeShippingPromotionId := "freeShippingPromotionId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponCodesPromotionRulesApi.GETFreeShippingPromotionIdCouponCodesPromotionRule(context.Background(), freeShippingPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponCodesPromotionRulesApi.GETFreeShippingPromotionIdCouponCodesPromotionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**freeShippingPromotionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETFreeShippingPromotionIdCouponCodesPromotionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPercentageDiscountPromotionIdCouponCodesPromotionRule

> GETPercentageDiscountPromotionIdCouponCodesPromotionRule(ctx, percentageDiscountPromotionId).Execute()

Retrieve the coupon codes promotion rule associated to the percentage discount promotion



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    percentageDiscountPromotionId := "percentageDiscountPromotionId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponCodesPromotionRulesApi.GETPercentageDiscountPromotionIdCouponCodesPromotionRule(context.Background(), percentageDiscountPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponCodesPromotionRulesApi.GETPercentageDiscountPromotionIdCouponCodesPromotionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**percentageDiscountPromotionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPercentageDiscountPromotionIdCouponCodesPromotionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPromotionIdCouponCodesPromotionRule

> GETPromotionIdCouponCodesPromotionRule(ctx, promotionId).Execute()

Retrieve the coupon codes promotion rule associated to the promotion



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    promotionId := "promotionId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponCodesPromotionRulesApi.GETPromotionIdCouponCodesPromotionRule(context.Background(), promotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponCodesPromotionRulesApi.GETPromotionIdCouponCodesPromotionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**promotionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPromotionIdCouponCodesPromotionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId

> PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId(ctx, couponCodesPromotionRuleId).CouponCodesPromotionRuleUpdate(couponCodesPromotionRuleUpdate).Execute()

Update a coupon codes promotion rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    couponCodesPromotionRuleUpdate := *openapiclient.NewCouponCodesPromotionRuleUpdate(*openapiclient.NewCouponCodesPromotionRuleUpdateData("coupon_codes_promotion_rules", "XGZwpOSrWL", *openapiclient.NewAdyenPaymentCreateDataAttributes())) // CouponCodesPromotionRuleUpdate | 
    couponCodesPromotionRuleId := "couponCodesPromotionRuleId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponCodesPromotionRulesApi.PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId(context.Background(), couponCodesPromotionRuleId).CouponCodesPromotionRuleUpdate(couponCodesPromotionRuleUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponCodesPromotionRulesApi.PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponCodesPromotionRuleId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHCouponCodesPromotionRulesCouponCodesPromotionRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **couponCodesPromotionRuleUpdate** | [**CouponCodesPromotionRuleUpdate**](CouponCodesPromotionRuleUpdate.md) |  | 


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTCouponCodesPromotionRules

> POSTCouponCodesPromotionRules(ctx).CouponCodesPromotionRuleCreate(couponCodesPromotionRuleCreate).Execute()

Create a coupon codes promotion rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    couponCodesPromotionRuleCreate := *openapiclient.NewCouponCodesPromotionRuleCreate(*openapiclient.NewCouponCodesPromotionRuleCreateData("coupon_codes_promotion_rules", *openapiclient.NewAdyenPaymentCreateDataAttributes())) // CouponCodesPromotionRuleCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponCodesPromotionRulesApi.POSTCouponCodesPromotionRules(context.Background()).CouponCodesPromotionRuleCreate(couponCodesPromotionRuleCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponCodesPromotionRulesApi.POSTCouponCodesPromotionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTCouponCodesPromotionRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **couponCodesPromotionRuleCreate** | [**CouponCodesPromotionRuleCreate**](CouponCodesPromotionRuleCreate.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

