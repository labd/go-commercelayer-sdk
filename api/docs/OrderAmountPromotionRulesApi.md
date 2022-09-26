# \OrderAmountPromotionRulesApi

All URIs are relative to *https://}.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEOrderAmountPromotionRulesOrderAmountPromotionRuleId**](OrderAmountPromotionRulesApi.md#DELETEOrderAmountPromotionRulesOrderAmountPromotionRuleId) | **Delete** /order_amount_promotion_rules/{orderAmountPromotionRuleId} | Delete an order amount promotion rule
[**GETExternalPromotionIdOrderAmountPromotionRule**](OrderAmountPromotionRulesApi.md#GETExternalPromotionIdOrderAmountPromotionRule) | **Get** /external_promotions/{externalPromotionId}/order_amount_promotion_rule | Retrieve the order amount promotion rule associated to the external promotion
[**GETFixedAmountPromotionIdOrderAmountPromotionRule**](OrderAmountPromotionRulesApi.md#GETFixedAmountPromotionIdOrderAmountPromotionRule) | **Get** /fixed_amount_promotions/{fixedAmountPromotionId}/order_amount_promotion_rule | Retrieve the order amount promotion rule associated to the fixed amount promotion
[**GETFixedPricePromotionIdOrderAmountPromotionRule**](OrderAmountPromotionRulesApi.md#GETFixedPricePromotionIdOrderAmountPromotionRule) | **Get** /fixed_price_promotions/{fixedPricePromotionId}/order_amount_promotion_rule | Retrieve the order amount promotion rule associated to the fixed price promotion
[**GETFreeGiftPromotionIdOrderAmountPromotionRule**](OrderAmountPromotionRulesApi.md#GETFreeGiftPromotionIdOrderAmountPromotionRule) | **Get** /free_gift_promotions/{freeGiftPromotionId}/order_amount_promotion_rule | Retrieve the order amount promotion rule associated to the free gift promotion
[**GETFreeShippingPromotionIdOrderAmountPromotionRule**](OrderAmountPromotionRulesApi.md#GETFreeShippingPromotionIdOrderAmountPromotionRule) | **Get** /free_shipping_promotions/{freeShippingPromotionId}/order_amount_promotion_rule | Retrieve the order amount promotion rule associated to the free shipping promotion
[**GETOrderAmountPromotionRules**](OrderAmountPromotionRulesApi.md#GETOrderAmountPromotionRules) | **Get** /order_amount_promotion_rules | List all order amount promotion rules
[**GETOrderAmountPromotionRulesOrderAmountPromotionRuleId**](OrderAmountPromotionRulesApi.md#GETOrderAmountPromotionRulesOrderAmountPromotionRuleId) | **Get** /order_amount_promotion_rules/{orderAmountPromotionRuleId} | Retrieve an order amount promotion rule
[**GETPercentageDiscountPromotionIdOrderAmountPromotionRule**](OrderAmountPromotionRulesApi.md#GETPercentageDiscountPromotionIdOrderAmountPromotionRule) | **Get** /percentage_discount_promotions/{percentageDiscountPromotionId}/order_amount_promotion_rule | Retrieve the order amount promotion rule associated to the percentage discount promotion
[**GETPromotionIdOrderAmountPromotionRule**](OrderAmountPromotionRulesApi.md#GETPromotionIdOrderAmountPromotionRule) | **Get** /promotions/{promotionId}/order_amount_promotion_rule | Retrieve the order amount promotion rule associated to the promotion
[**PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleId**](OrderAmountPromotionRulesApi.md#PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleId) | **Patch** /order_amount_promotion_rules/{orderAmountPromotionRuleId} | Update an order amount promotion rule
[**POSTOrderAmountPromotionRules**](OrderAmountPromotionRulesApi.md#POSTOrderAmountPromotionRules) | **Post** /order_amount_promotion_rules | Create an order amount promotion rule



## DELETEOrderAmountPromotionRulesOrderAmountPromotionRuleId

> DELETEOrderAmountPromotionRulesOrderAmountPromotionRuleId(ctx, orderAmountPromotionRuleId).Execute()

Delete an order amount promotion rule



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
    orderAmountPromotionRuleId := "orderAmountPromotionRuleId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAmountPromotionRulesApi.DELETEOrderAmountPromotionRulesOrderAmountPromotionRuleId(context.Background(), orderAmountPromotionRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAmountPromotionRulesApi.DELETEOrderAmountPromotionRulesOrderAmountPromotionRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderAmountPromotionRuleId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETExternalPromotionIdOrderAmountPromotionRule

> GETExternalPromotionIdOrderAmountPromotionRule(ctx, externalPromotionId).Execute()

Retrieve the order amount promotion rule associated to the external promotion



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
    resp, r, err := apiClient.OrderAmountPromotionRulesApi.GETExternalPromotionIdOrderAmountPromotionRule(context.Background(), externalPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAmountPromotionRulesApi.GETExternalPromotionIdOrderAmountPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETExternalPromotionIdOrderAmountPromotionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETFixedAmountPromotionIdOrderAmountPromotionRule

> GETFixedAmountPromotionIdOrderAmountPromotionRule(ctx, fixedAmountPromotionId).Execute()

Retrieve the order amount promotion rule associated to the fixed amount promotion



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
    resp, r, err := apiClient.OrderAmountPromotionRulesApi.GETFixedAmountPromotionIdOrderAmountPromotionRule(context.Background(), fixedAmountPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAmountPromotionRulesApi.GETFixedAmountPromotionIdOrderAmountPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETFixedAmountPromotionIdOrderAmountPromotionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETFixedPricePromotionIdOrderAmountPromotionRule

> GETFixedPricePromotionIdOrderAmountPromotionRule(ctx, fixedPricePromotionId).Execute()

Retrieve the order amount promotion rule associated to the fixed price promotion



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
    resp, r, err := apiClient.OrderAmountPromotionRulesApi.GETFixedPricePromotionIdOrderAmountPromotionRule(context.Background(), fixedPricePromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAmountPromotionRulesApi.GETFixedPricePromotionIdOrderAmountPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETFixedPricePromotionIdOrderAmountPromotionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETFreeGiftPromotionIdOrderAmountPromotionRule

> GETFreeGiftPromotionIdOrderAmountPromotionRule(ctx, freeGiftPromotionId).Execute()

Retrieve the order amount promotion rule associated to the free gift promotion



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
    resp, r, err := apiClient.OrderAmountPromotionRulesApi.GETFreeGiftPromotionIdOrderAmountPromotionRule(context.Background(), freeGiftPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAmountPromotionRulesApi.GETFreeGiftPromotionIdOrderAmountPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETFreeGiftPromotionIdOrderAmountPromotionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETFreeShippingPromotionIdOrderAmountPromotionRule

> GETFreeShippingPromotionIdOrderAmountPromotionRule(ctx, freeShippingPromotionId).Execute()

Retrieve the order amount promotion rule associated to the free shipping promotion



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
    resp, r, err := apiClient.OrderAmountPromotionRulesApi.GETFreeShippingPromotionIdOrderAmountPromotionRule(context.Background(), freeShippingPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAmountPromotionRulesApi.GETFreeShippingPromotionIdOrderAmountPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETFreeShippingPromotionIdOrderAmountPromotionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETOrderAmountPromotionRules

> OrderAmountPromotionRuleResponseList GETOrderAmountPromotionRules(ctx).Execute()

List all order amount promotion rules



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
    resp, r, err := apiClient.OrderAmountPromotionRulesApi.GETOrderAmountPromotionRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAmountPromotionRulesApi.GETOrderAmountPromotionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETOrderAmountPromotionRules`: OrderAmountPromotionRuleResponseList
    fmt.Fprintf(os.Stdout, "Response from `OrderAmountPromotionRulesApi.GETOrderAmountPromotionRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETOrderAmountPromotionRulesRequest struct via the builder pattern


### Return type

[**OrderAmountPromotionRuleResponseList**](OrderAmountPromotionRuleResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETOrderAmountPromotionRulesOrderAmountPromotionRuleId

> OrderAmountPromotionRuleResponse GETOrderAmountPromotionRulesOrderAmountPromotionRuleId(ctx, orderAmountPromotionRuleId).Execute()

Retrieve an order amount promotion rule



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
    orderAmountPromotionRuleId := "orderAmountPromotionRuleId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAmountPromotionRulesApi.GETOrderAmountPromotionRulesOrderAmountPromotionRuleId(context.Background(), orderAmountPromotionRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAmountPromotionRulesApi.GETOrderAmountPromotionRulesOrderAmountPromotionRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETOrderAmountPromotionRulesOrderAmountPromotionRuleId`: OrderAmountPromotionRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `OrderAmountPromotionRulesApi.GETOrderAmountPromotionRulesOrderAmountPromotionRuleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderAmountPromotionRuleId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrderAmountPromotionRuleResponse**](OrderAmountPromotionRuleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPercentageDiscountPromotionIdOrderAmountPromotionRule

> GETPercentageDiscountPromotionIdOrderAmountPromotionRule(ctx, percentageDiscountPromotionId).Execute()

Retrieve the order amount promotion rule associated to the percentage discount promotion



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
    resp, r, err := apiClient.OrderAmountPromotionRulesApi.GETPercentageDiscountPromotionIdOrderAmountPromotionRule(context.Background(), percentageDiscountPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAmountPromotionRulesApi.GETPercentageDiscountPromotionIdOrderAmountPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETPercentageDiscountPromotionIdOrderAmountPromotionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPromotionIdOrderAmountPromotionRule

> GETPromotionIdOrderAmountPromotionRule(ctx, promotionId).Execute()

Retrieve the order amount promotion rule associated to the promotion



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
    resp, r, err := apiClient.OrderAmountPromotionRulesApi.GETPromotionIdOrderAmountPromotionRule(context.Background(), promotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAmountPromotionRulesApi.GETPromotionIdOrderAmountPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETPromotionIdOrderAmountPromotionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleId

> OrderAmountPromotionRuleResponse PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleId(ctx, orderAmountPromotionRuleId).OrderAmountPromotionRuleUpdate(orderAmountPromotionRuleUpdate).Execute()

Update an order amount promotion rule



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
    orderAmountPromotionRuleUpdate := *openapiclient.NewOrderAmountPromotionRuleUpdate(*openapiclient.NewOrderAmountPromotionRuleUpdateData("Type_example", "XGZwpOSrWL", *openapiclient.NewOrderAmountPromotionRuleCreateDataAttributes())) // OrderAmountPromotionRuleUpdate | 
    orderAmountPromotionRuleId := "orderAmountPromotionRuleId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAmountPromotionRulesApi.PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleId(context.Background(), orderAmountPromotionRuleId).OrderAmountPromotionRuleUpdate(orderAmountPromotionRuleUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAmountPromotionRulesApi.PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleId`: OrderAmountPromotionRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `OrderAmountPromotionRulesApi.PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderAmountPromotionRuleId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderAmountPromotionRuleUpdate** | [**OrderAmountPromotionRuleUpdate**](OrderAmountPromotionRuleUpdate.md) |  | 


### Return type

[**OrderAmountPromotionRuleResponse**](OrderAmountPromotionRuleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTOrderAmountPromotionRules

> OrderAmountPromotionRuleResponse POSTOrderAmountPromotionRules(ctx).OrderAmountPromotionRuleCreate(orderAmountPromotionRuleCreate).Execute()

Create an order amount promotion rule



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
    orderAmountPromotionRuleCreate := *openapiclient.NewOrderAmountPromotionRuleCreate(*openapiclient.NewOrderAmountPromotionRuleCreateData("Type_example", *openapiclient.NewOrderAmountPromotionRuleCreateDataAttributes())) // OrderAmountPromotionRuleCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAmountPromotionRulesApi.POSTOrderAmountPromotionRules(context.Background()).OrderAmountPromotionRuleCreate(orderAmountPromotionRuleCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAmountPromotionRulesApi.POSTOrderAmountPromotionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTOrderAmountPromotionRules`: OrderAmountPromotionRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `OrderAmountPromotionRulesApi.POSTOrderAmountPromotionRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTOrderAmountPromotionRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderAmountPromotionRuleCreate** | [**OrderAmountPromotionRuleCreate**](OrderAmountPromotionRuleCreate.md) |  | 

### Return type

[**OrderAmountPromotionRuleResponse**](OrderAmountPromotionRuleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

