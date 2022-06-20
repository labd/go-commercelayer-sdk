# \FreeShippingPromotionsApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEFreeShippingPromotionsFreeShippingPromotionId**](FreeShippingPromotionsApi.md#DELETEFreeShippingPromotionsFreeShippingPromotionId) | **Delete** /free_shipping_promotions/{freeShippingPromotionId} | Delete a free shipping promotion
[**GETFreeShippingPromotions**](FreeShippingPromotionsApi.md#GETFreeShippingPromotions) | **Get** /free_shipping_promotions | List all free shipping promotions
[**GETFreeShippingPromotionsFreeShippingPromotionId**](FreeShippingPromotionsApi.md#GETFreeShippingPromotionsFreeShippingPromotionId) | **Get** /free_shipping_promotions/{freeShippingPromotionId} | Retrieve a free shipping promotion
[**PATCHFreeShippingPromotionsFreeShippingPromotionId**](FreeShippingPromotionsApi.md#PATCHFreeShippingPromotionsFreeShippingPromotionId) | **Patch** /free_shipping_promotions/{freeShippingPromotionId} | Update a free shipping promotion
[**POSTFreeShippingPromotions**](FreeShippingPromotionsApi.md#POSTFreeShippingPromotions) | **Post** /free_shipping_promotions | Create a free shipping promotion



## DELETEFreeShippingPromotionsFreeShippingPromotionId

> DELETEFreeShippingPromotionsFreeShippingPromotionId(ctx, freeShippingPromotionId).Execute()

Delete a free shipping promotion



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
    resp, r, err := apiClient.FreeShippingPromotionsApi.DELETEFreeShippingPromotionsFreeShippingPromotionId(context.Background(), freeShippingPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FreeShippingPromotionsApi.DELETEFreeShippingPromotionsFreeShippingPromotionId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDELETEFreeShippingPromotionsFreeShippingPromotionIdRequest struct via the builder pattern


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


## GETFreeShippingPromotions

> GETFreeShippingPromotions(ctx).Execute()

List all free shipping promotions



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
    resp, r, err := apiClient.FreeShippingPromotionsApi.GETFreeShippingPromotions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FreeShippingPromotionsApi.GETFreeShippingPromotions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETFreeShippingPromotionsRequest struct via the builder pattern


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


## GETFreeShippingPromotionsFreeShippingPromotionId

> FreeShippingPromotion GETFreeShippingPromotionsFreeShippingPromotionId(ctx, freeShippingPromotionId).Execute()

Retrieve a free shipping promotion



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
    resp, r, err := apiClient.FreeShippingPromotionsApi.GETFreeShippingPromotionsFreeShippingPromotionId(context.Background(), freeShippingPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FreeShippingPromotionsApi.GETFreeShippingPromotionsFreeShippingPromotionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETFreeShippingPromotionsFreeShippingPromotionId`: FreeShippingPromotion
    fmt.Fprintf(os.Stdout, "Response from `FreeShippingPromotionsApi.GETFreeShippingPromotionsFreeShippingPromotionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**freeShippingPromotionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETFreeShippingPromotionsFreeShippingPromotionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FreeShippingPromotion**](FreeShippingPromotion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHFreeShippingPromotionsFreeShippingPromotionId

> PATCHFreeShippingPromotionsFreeShippingPromotionId(ctx, freeShippingPromotionId).FreeShippingPromotionUpdate(freeShippingPromotionUpdate).Execute()

Update a free shipping promotion



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
    freeShippingPromotionUpdate := *openapiclient.NewFreeShippingPromotionUpdate(*openapiclient.NewFreeShippingPromotionUpdateData("free_shipping_promotions", "XGZwpOSrWL", *openapiclient.NewFreeShippingPromotionUpdateDataAttributes())) // FreeShippingPromotionUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FreeShippingPromotionsApi.PATCHFreeShippingPromotionsFreeShippingPromotionId(context.Background(), freeShippingPromotionId).FreeShippingPromotionUpdate(freeShippingPromotionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FreeShippingPromotionsApi.PATCHFreeShippingPromotionsFreeShippingPromotionId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPATCHFreeShippingPromotionsFreeShippingPromotionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **freeShippingPromotionUpdate** | [**FreeShippingPromotionUpdate**](FreeShippingPromotionUpdate.md) |  | 

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


## POSTFreeShippingPromotions

> POSTFreeShippingPromotions(ctx).FreeShippingPromotionCreate(freeShippingPromotionCreate).Execute()

Create a free shipping promotion



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
    freeShippingPromotionCreate := *openapiclient.NewFreeShippingPromotionCreate(*openapiclient.NewFreeShippingPromotionCreateData("free_shipping_promotions", *openapiclient.NewFreeShippingPromotionCreateDataAttributes("Personal promotion", "2018-01-01T12:00:00.000Z", "2018-01-02T12:00:00.000Z", int32(5)))) // FreeShippingPromotionCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FreeShippingPromotionsApi.POSTFreeShippingPromotions(context.Background()).FreeShippingPromotionCreate(freeShippingPromotionCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FreeShippingPromotionsApi.POSTFreeShippingPromotions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTFreeShippingPromotionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **freeShippingPromotionCreate** | [**FreeShippingPromotionCreate**](FreeShippingPromotionCreate.md) |  | 

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

