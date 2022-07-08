# \AdjustmentsApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEAdjustmentsAdjustmentId**](AdjustmentsApi.md#DELETEAdjustmentsAdjustmentId) | **Delete** /adjustments/{adjustmentId} | Delete an adjustment
[**GETAdjustments**](AdjustmentsApi.md#GETAdjustments) | **Get** /adjustments | List all adjustments
[**GETAdjustmentsAdjustmentId**](AdjustmentsApi.md#GETAdjustmentsAdjustmentId) | **Get** /adjustments/{adjustmentId} | Retrieve an adjustment
[**PATCHAdjustmentsAdjustmentId**](AdjustmentsApi.md#PATCHAdjustmentsAdjustmentId) | **Patch** /adjustments/{adjustmentId} | Update an adjustment
[**POSTAdjustments**](AdjustmentsApi.md#POSTAdjustments) | **Post** /adjustments | Create an adjustment



## DELETEAdjustmentsAdjustmentId

> DELETEAdjustmentsAdjustmentId(ctx, adjustmentId).Execute()

Delete an adjustment



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
    adjustmentId := "adjustmentId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdjustmentsApi.DELETEAdjustmentsAdjustmentId(context.Background(), adjustmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdjustmentsApi.DELETEAdjustmentsAdjustmentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adjustmentId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEAdjustmentsAdjustmentIdRequest struct via the builder pattern


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


## GETAdjustments

> GETAdjustments(ctx).Execute()

List all adjustments



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
    resp, r, err := apiClient.AdjustmentsApi.GETAdjustments(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdjustmentsApi.GETAdjustments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETAdjustmentsRequest struct via the builder pattern


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


## GETAdjustmentsAdjustmentId

> Adjustment GETAdjustmentsAdjustmentId(ctx, adjustmentId).Execute()

Retrieve an adjustment



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
    adjustmentId := "adjustmentId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdjustmentsApi.GETAdjustmentsAdjustmentId(context.Background(), adjustmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdjustmentsApi.GETAdjustmentsAdjustmentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETAdjustmentsAdjustmentId`: Adjustment
    fmt.Fprintf(os.Stdout, "Response from `AdjustmentsApi.GETAdjustmentsAdjustmentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adjustmentId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETAdjustmentsAdjustmentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Adjustment**](Adjustment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHAdjustmentsAdjustmentId

> PATCHAdjustmentsAdjustmentId(ctx, adjustmentId).AdjustmentUpdate(adjustmentUpdate).Execute()

Update an adjustment



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
    adjustmentUpdate := *openapiclient.NewAdjustmentUpdate(*openapiclient.NewAdjustmentUpdateData("adjustments", "XGZwpOSrWL", *openapiclient.NewAdjustmentUpdateDataAttributes())) // AdjustmentUpdate | 
    adjustmentId := "adjustmentId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdjustmentsApi.PATCHAdjustmentsAdjustmentId(context.Background(), adjustmentId).AdjustmentUpdate(adjustmentUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdjustmentsApi.PATCHAdjustmentsAdjustmentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adjustmentId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHAdjustmentsAdjustmentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adjustmentUpdate** | [**AdjustmentUpdate**](AdjustmentUpdate.md) |  | 


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


## POSTAdjustments

> POSTAdjustments(ctx).AdjustmentCreate(adjustmentCreate).Execute()

Create an adjustment



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
    adjustmentCreate := *openapiclient.NewAdjustmentCreate(*openapiclient.NewAdjustmentCreateData("adjustments", *openapiclient.NewAdjustmentCreateDataAttributes("Additional service", "EUR", int32(1500)))) // AdjustmentCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdjustmentsApi.POSTAdjustments(context.Background()).AdjustmentCreate(adjustmentCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdjustmentsApi.POSTAdjustments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTAdjustmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adjustmentCreate** | [**AdjustmentCreate**](AdjustmentCreate.md) |  | 

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

