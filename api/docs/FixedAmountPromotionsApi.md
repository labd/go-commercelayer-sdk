# \FixedAmountPromotionsApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEFixedAmountPromotionsFixedAmountPromotionId**](FixedAmountPromotionsApi.md#DELETEFixedAmountPromotionsFixedAmountPromotionId) | **Delete** /fixed_amount_promotions/{fixedAmountPromotionId} | Delete a fixed amount promotion
[**GETFixedAmountPromotions**](FixedAmountPromotionsApi.md#GETFixedAmountPromotions) | **Get** /fixed_amount_promotions | List all fixed amount promotions
[**GETFixedAmountPromotionsFixedAmountPromotionId**](FixedAmountPromotionsApi.md#GETFixedAmountPromotionsFixedAmountPromotionId) | **Get** /fixed_amount_promotions/{fixedAmountPromotionId} | Retrieve a fixed amount promotion
[**PATCHFixedAmountPromotionsFixedAmountPromotionId**](FixedAmountPromotionsApi.md#PATCHFixedAmountPromotionsFixedAmountPromotionId) | **Patch** /fixed_amount_promotions/{fixedAmountPromotionId} | Update a fixed amount promotion
[**POSTFixedAmountPromotions**](FixedAmountPromotionsApi.md#POSTFixedAmountPromotions) | **Post** /fixed_amount_promotions | Create a fixed amount promotion



## DELETEFixedAmountPromotionsFixedAmountPromotionId

> DELETEFixedAmountPromotionsFixedAmountPromotionId(ctx, fixedAmountPromotionId).Execute()

Delete a fixed amount promotion



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
    resp, r, err := apiClient.FixedAmountPromotionsApi.DELETEFixedAmountPromotionsFixedAmountPromotionId(context.Background(), fixedAmountPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FixedAmountPromotionsApi.DELETEFixedAmountPromotionsFixedAmountPromotionId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDELETEFixedAmountPromotionsFixedAmountPromotionIdRequest struct via the builder pattern


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


## GETFixedAmountPromotions

> GETFixedAmountPromotions(ctx).Execute()

List all fixed amount promotions



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
    resp, r, err := apiClient.FixedAmountPromotionsApi.GETFixedAmountPromotions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FixedAmountPromotionsApi.GETFixedAmountPromotions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETFixedAmountPromotionsRequest struct via the builder pattern


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


## GETFixedAmountPromotionsFixedAmountPromotionId

> FixedAmountPromotion GETFixedAmountPromotionsFixedAmountPromotionId(ctx, fixedAmountPromotionId).Execute()

Retrieve a fixed amount promotion



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
    resp, r, err := apiClient.FixedAmountPromotionsApi.GETFixedAmountPromotionsFixedAmountPromotionId(context.Background(), fixedAmountPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FixedAmountPromotionsApi.GETFixedAmountPromotionsFixedAmountPromotionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETFixedAmountPromotionsFixedAmountPromotionId`: FixedAmountPromotion
    fmt.Fprintf(os.Stdout, "Response from `FixedAmountPromotionsApi.GETFixedAmountPromotionsFixedAmountPromotionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fixedAmountPromotionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETFixedAmountPromotionsFixedAmountPromotionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FixedAmountPromotion**](FixedAmountPromotion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHFixedAmountPromotionsFixedAmountPromotionId

> PATCHFixedAmountPromotionsFixedAmountPromotionId(ctx, fixedAmountPromotionId).FixedAmountPromotionUpdate(fixedAmountPromotionUpdate).Execute()

Update a fixed amount promotion



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
    fixedAmountPromotionUpdate := *openapiclient.NewFixedAmountPromotionUpdate(*openapiclient.NewFixedAmountPromotionUpdateData("fixed_amount_promotions", "XGZwpOSrWL", *openapiclient.NewFixedAmountPromotionUpdateDataAttributes())) // FixedAmountPromotionUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FixedAmountPromotionsApi.PATCHFixedAmountPromotionsFixedAmountPromotionId(context.Background(), fixedAmountPromotionId).FixedAmountPromotionUpdate(fixedAmountPromotionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FixedAmountPromotionsApi.PATCHFixedAmountPromotionsFixedAmountPromotionId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPATCHFixedAmountPromotionsFixedAmountPromotionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fixedAmountPromotionUpdate** | [**FixedAmountPromotionUpdate**](FixedAmountPromotionUpdate.md) |  | 

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


## POSTFixedAmountPromotions

> POSTFixedAmountPromotions(ctx).FixedAmountPromotionCreate(fixedAmountPromotionCreate).Execute()

Create a fixed amount promotion



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
    fixedAmountPromotionCreate := *openapiclient.NewFixedAmountPromotionCreate(*openapiclient.NewFixedAmountPromotionCreateData("fixed_amount_promotions", *openapiclient.NewFixedAmountPromotionCreateDataAttributes("Personal promotion", "2018-01-01T12:00:00.000Z", "2018-01-02T12:00:00.000Z", int32(5), int32(1000)))) // FixedAmountPromotionCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FixedAmountPromotionsApi.POSTFixedAmountPromotions(context.Background()).FixedAmountPromotionCreate(fixedAmountPromotionCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FixedAmountPromotionsApi.POSTFixedAmountPromotions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTFixedAmountPromotionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fixedAmountPromotionCreate** | [**FixedAmountPromotionCreate**](FixedAmountPromotionCreate.md) |  | 

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

