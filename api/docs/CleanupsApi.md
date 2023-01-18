# \CleanupsApi

All URIs are relative to *https://}.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETECleanupsCleanupId**](CleanupsApi.md#DELETECleanupsCleanupId) | **Delete** /cleanups/{cleanupId} | Delete a cleanup
[**GETCleanups**](CleanupsApi.md#GETCleanups) | **Get** /cleanups | List all cleanups
[**GETCleanupsCleanupId**](CleanupsApi.md#GETCleanupsCleanupId) | **Get** /cleanups/{cleanupId} | Retrieve a cleanup
[**POSTCleanups**](CleanupsApi.md#POSTCleanups) | **Post** /cleanups | Create a cleanup



## DELETECleanupsCleanupId

> DELETECleanupsCleanupId(ctx, cleanupId).Execute()

Delete a cleanup



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
    cleanupId := "cleanupId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CleanupsApi.DELETECleanupsCleanupId(context.Background(), cleanupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CleanupsApi.DELETECleanupsCleanupId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cleanupId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETECleanupsCleanupIdRequest struct via the builder pattern


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


## GETCleanups

> GETCleanups200Response GETCleanups(ctx).Execute()

List all cleanups



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
    resp, r, err := apiClient.CleanupsApi.GETCleanups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CleanupsApi.GETCleanups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCleanups`: GETCleanups200Response
    fmt.Fprintf(os.Stdout, "Response from `CleanupsApi.GETCleanups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETCleanupsRequest struct via the builder pattern


### Return type

[**GETCleanups200Response**](GETCleanups200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCleanupsCleanupId

> GETCleanupsCleanupId200Response GETCleanupsCleanupId(ctx, cleanupId).Execute()

Retrieve a cleanup



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
    cleanupId := "cleanupId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CleanupsApi.GETCleanupsCleanupId(context.Background(), cleanupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CleanupsApi.GETCleanupsCleanupId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCleanupsCleanupId`: GETCleanupsCleanupId200Response
    fmt.Fprintf(os.Stdout, "Response from `CleanupsApi.GETCleanupsCleanupId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cleanupId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCleanupsCleanupIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETCleanupsCleanupId200Response**](GETCleanupsCleanupId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTCleanups

> POSTCleanups201Response POSTCleanups(ctx).CleanupCreate(cleanupCreate).Execute()

Create a cleanup



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
    cleanupCreate := *openapiclient.NewCleanupCreate(*openapiclient.NewCleanupCreateData("Type_example", *openapiclient.NewPOSTCleanups201ResponseDataAttributes("skus"))) // CleanupCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CleanupsApi.POSTCleanups(context.Background()).CleanupCreate(cleanupCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CleanupsApi.POSTCleanups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTCleanups`: POSTCleanups201Response
    fmt.Fprintf(os.Stdout, "Response from `CleanupsApi.POSTCleanups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTCleanupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cleanupCreate** | [**CleanupCreate**](CleanupCreate.md) |  | 

### Return type

[**POSTCleanups201Response**](POSTCleanups201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

