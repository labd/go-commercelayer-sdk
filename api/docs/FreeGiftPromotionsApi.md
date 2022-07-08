# \FreeGiftPromotionsApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEFreeGiftPromotionsFreeGiftPromotionId**](FreeGiftPromotionsApi.md#DELETEFreeGiftPromotionsFreeGiftPromotionId) | **Delete** /free_gift_promotions/{freeGiftPromotionId} | Delete a free gift promotion
[**GETFreeGiftPromotions**](FreeGiftPromotionsApi.md#GETFreeGiftPromotions) | **Get** /free_gift_promotions | List all free gift promotions
[**GETFreeGiftPromotionsFreeGiftPromotionId**](FreeGiftPromotionsApi.md#GETFreeGiftPromotionsFreeGiftPromotionId) | **Get** /free_gift_promotions/{freeGiftPromotionId} | Retrieve a free gift promotion
[**PATCHFreeGiftPromotionsFreeGiftPromotionId**](FreeGiftPromotionsApi.md#PATCHFreeGiftPromotionsFreeGiftPromotionId) | **Patch** /free_gift_promotions/{freeGiftPromotionId} | Update a free gift promotion
[**POSTFreeGiftPromotions**](FreeGiftPromotionsApi.md#POSTFreeGiftPromotions) | **Post** /free_gift_promotions | Create a free gift promotion



## DELETEFreeGiftPromotionsFreeGiftPromotionId

> DELETEFreeGiftPromotionsFreeGiftPromotionId(ctx, freeGiftPromotionId).Execute()

Delete a free gift promotion



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
    resp, r, err := apiClient.FreeGiftPromotionsApi.DELETEFreeGiftPromotionsFreeGiftPromotionId(context.Background(), freeGiftPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FreeGiftPromotionsApi.DELETEFreeGiftPromotionsFreeGiftPromotionId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDELETEFreeGiftPromotionsFreeGiftPromotionIdRequest struct via the builder pattern


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


## GETFreeGiftPromotions

> GETFreeGiftPromotions(ctx).Execute()

List all free gift promotions



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
    resp, r, err := apiClient.FreeGiftPromotionsApi.GETFreeGiftPromotions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FreeGiftPromotionsApi.GETFreeGiftPromotions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETFreeGiftPromotionsRequest struct via the builder pattern


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


## GETFreeGiftPromotionsFreeGiftPromotionId

> FreeGiftPromotion GETFreeGiftPromotionsFreeGiftPromotionId(ctx, freeGiftPromotionId).Execute()

Retrieve a free gift promotion



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
    resp, r, err := apiClient.FreeGiftPromotionsApi.GETFreeGiftPromotionsFreeGiftPromotionId(context.Background(), freeGiftPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FreeGiftPromotionsApi.GETFreeGiftPromotionsFreeGiftPromotionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETFreeGiftPromotionsFreeGiftPromotionId`: FreeGiftPromotion
    fmt.Fprintf(os.Stdout, "Response from `FreeGiftPromotionsApi.GETFreeGiftPromotionsFreeGiftPromotionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**freeGiftPromotionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETFreeGiftPromotionsFreeGiftPromotionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FreeGiftPromotion**](FreeGiftPromotion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHFreeGiftPromotionsFreeGiftPromotionId

> PATCHFreeGiftPromotionsFreeGiftPromotionId(ctx, freeGiftPromotionId).FreeGiftPromotionUpdate(freeGiftPromotionUpdate).Execute()

Update a free gift promotion



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
    freeGiftPromotionUpdate := *openapiclient.NewFreeGiftPromotionUpdate(*openapiclient.NewFreeGiftPromotionUpdateData("free_gift_promotions", "XGZwpOSrWL", *openapiclient.NewFreeGiftPromotionUpdateDataAttributes())) // FreeGiftPromotionUpdate | 
    freeGiftPromotionId := "freeGiftPromotionId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FreeGiftPromotionsApi.PATCHFreeGiftPromotionsFreeGiftPromotionId(context.Background(), freeGiftPromotionId).FreeGiftPromotionUpdate(freeGiftPromotionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FreeGiftPromotionsApi.PATCHFreeGiftPromotionsFreeGiftPromotionId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPATCHFreeGiftPromotionsFreeGiftPromotionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **freeGiftPromotionUpdate** | [**FreeGiftPromotionUpdate**](FreeGiftPromotionUpdate.md) |  | 


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


## POSTFreeGiftPromotions

> POSTFreeGiftPromotions(ctx).FreeGiftPromotionCreate(freeGiftPromotionCreate).Execute()

Create a free gift promotion



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
    freeGiftPromotionCreate := *openapiclient.NewFreeGiftPromotionCreate(*openapiclient.NewFreeGiftPromotionCreateData("free_gift_promotions", *openapiclient.NewFreeGiftPromotionCreateDataAttributes("Personal promotion", "2018-01-01T12:00:00.000Z", "2018-01-02T12:00:00.000Z", int32(5)))) // FreeGiftPromotionCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FreeGiftPromotionsApi.POSTFreeGiftPromotions(context.Background()).FreeGiftPromotionCreate(freeGiftPromotionCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FreeGiftPromotionsApi.POSTFreeGiftPromotions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTFreeGiftPromotionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **freeGiftPromotionCreate** | [**FreeGiftPromotionCreate**](FreeGiftPromotionCreate.md) |  | 

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

