# \PercentageDiscountPromotionsApi

All URIs are relative to *https://}.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEPercentageDiscountPromotionsPercentageDiscountPromotionId**](PercentageDiscountPromotionsApi.md#DELETEPercentageDiscountPromotionsPercentageDiscountPromotionId) | **Delete** /percentage_discount_promotions/{percentageDiscountPromotionId} | Delete a percentage discount promotion
[**GETPercentageDiscountPromotions**](PercentageDiscountPromotionsApi.md#GETPercentageDiscountPromotions) | **Get** /percentage_discount_promotions | List all percentage discount promotions
[**GETPercentageDiscountPromotionsPercentageDiscountPromotionId**](PercentageDiscountPromotionsApi.md#GETPercentageDiscountPromotionsPercentageDiscountPromotionId) | **Get** /percentage_discount_promotions/{percentageDiscountPromotionId} | Retrieve a percentage discount promotion
[**PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId**](PercentageDiscountPromotionsApi.md#PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId) | **Patch** /percentage_discount_promotions/{percentageDiscountPromotionId} | Update a percentage discount promotion
[**POSTPercentageDiscountPromotions**](PercentageDiscountPromotionsApi.md#POSTPercentageDiscountPromotions) | **Post** /percentage_discount_promotions | Create a percentage discount promotion



## DELETEPercentageDiscountPromotionsPercentageDiscountPromotionId

> DELETEPercentageDiscountPromotionsPercentageDiscountPromotionId(ctx, percentageDiscountPromotionId).Execute()

Delete a percentage discount promotion



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
    resp, r, err := apiClient.PercentageDiscountPromotionsApi.DELETEPercentageDiscountPromotionsPercentageDiscountPromotionId(context.Background(), percentageDiscountPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PercentageDiscountPromotionsApi.DELETEPercentageDiscountPromotionsPercentageDiscountPromotionId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDELETEPercentageDiscountPromotionsPercentageDiscountPromotionIdRequest struct via the builder pattern


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


## GETPercentageDiscountPromotions

> PercentageDiscountPromotionResponseList GETPercentageDiscountPromotions(ctx).Execute()

List all percentage discount promotions



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
    resp, r, err := apiClient.PercentageDiscountPromotionsApi.GETPercentageDiscountPromotions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PercentageDiscountPromotionsApi.GETPercentageDiscountPromotions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPercentageDiscountPromotions`: PercentageDiscountPromotionResponseList
    fmt.Fprintf(os.Stdout, "Response from `PercentageDiscountPromotionsApi.GETPercentageDiscountPromotions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETPercentageDiscountPromotionsRequest struct via the builder pattern


### Return type

[**PercentageDiscountPromotionResponseList**](PercentageDiscountPromotionResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPercentageDiscountPromotionsPercentageDiscountPromotionId

> PercentageDiscountPromotionResponse GETPercentageDiscountPromotionsPercentageDiscountPromotionId(ctx, percentageDiscountPromotionId).Execute()

Retrieve a percentage discount promotion



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
    resp, r, err := apiClient.PercentageDiscountPromotionsApi.GETPercentageDiscountPromotionsPercentageDiscountPromotionId(context.Background(), percentageDiscountPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PercentageDiscountPromotionsApi.GETPercentageDiscountPromotionsPercentageDiscountPromotionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPercentageDiscountPromotionsPercentageDiscountPromotionId`: PercentageDiscountPromotionResponse
    fmt.Fprintf(os.Stdout, "Response from `PercentageDiscountPromotionsApi.GETPercentageDiscountPromotionsPercentageDiscountPromotionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**percentageDiscountPromotionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPercentageDiscountPromotionsPercentageDiscountPromotionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PercentageDiscountPromotionResponse**](PercentageDiscountPromotionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId

> PercentageDiscountPromotionResponse PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId(ctx, percentageDiscountPromotionId).PercentageDiscountPromotionUpdate(percentageDiscountPromotionUpdate).Execute()

Update a percentage discount promotion



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
    percentageDiscountPromotionUpdate := *openapiclient.NewPercentageDiscountPromotionUpdate(*openapiclient.NewPercentageDiscountPromotionUpdateData("Type_example", "XGZwpOSrWL", *openapiclient.NewPercentageDiscountPromotionUpdateDataAttributes())) // PercentageDiscountPromotionUpdate | 
    percentageDiscountPromotionId := "percentageDiscountPromotionId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PercentageDiscountPromotionsApi.PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId(context.Background(), percentageDiscountPromotionId).PercentageDiscountPromotionUpdate(percentageDiscountPromotionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PercentageDiscountPromotionsApi.PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId`: PercentageDiscountPromotionResponse
    fmt.Fprintf(os.Stdout, "Response from `PercentageDiscountPromotionsApi.PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**percentageDiscountPromotionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHPercentageDiscountPromotionsPercentageDiscountPromotionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **percentageDiscountPromotionUpdate** | [**PercentageDiscountPromotionUpdate**](PercentageDiscountPromotionUpdate.md) |  | 


### Return type

[**PercentageDiscountPromotionResponse**](PercentageDiscountPromotionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTPercentageDiscountPromotions

> PercentageDiscountPromotionResponse POSTPercentageDiscountPromotions(ctx).PercentageDiscountPromotionCreate(percentageDiscountPromotionCreate).Execute()

Create a percentage discount promotion



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
    percentageDiscountPromotionCreate := *openapiclient.NewPercentageDiscountPromotionCreate(*openapiclient.NewPercentageDiscountPromotionCreateData("Type_example", *openapiclient.NewPercentageDiscountPromotionCreateDataAttributes("Personal promotion", "2018-01-01T12:00:00.000Z", "2018-01-02T12:00:00.000Z", int32(5), int32(10)))) // PercentageDiscountPromotionCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PercentageDiscountPromotionsApi.POSTPercentageDiscountPromotions(context.Background()).PercentageDiscountPromotionCreate(percentageDiscountPromotionCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PercentageDiscountPromotionsApi.POSTPercentageDiscountPromotions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTPercentageDiscountPromotions`: PercentageDiscountPromotionResponse
    fmt.Fprintf(os.Stdout, "Response from `PercentageDiscountPromotionsApi.POSTPercentageDiscountPromotions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTPercentageDiscountPromotionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **percentageDiscountPromotionCreate** | [**PercentageDiscountPromotionCreate**](PercentageDiscountPromotionCreate.md) |  | 

### Return type

[**PercentageDiscountPromotionResponse**](PercentageDiscountPromotionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

