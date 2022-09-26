# \WireTransfersApi

All URIs are relative to *https://}.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEWireTransfersWireTransferId**](WireTransfersApi.md#DELETEWireTransfersWireTransferId) | **Delete** /wire_transfers/{wireTransferId} | Delete a wire transfer
[**GETWireTransfers**](WireTransfersApi.md#GETWireTransfers) | **Get** /wire_transfers | List all wire transfers
[**GETWireTransfersWireTransferId**](WireTransfersApi.md#GETWireTransfersWireTransferId) | **Get** /wire_transfers/{wireTransferId} | Retrieve a wire transfer
[**PATCHWireTransfersWireTransferId**](WireTransfersApi.md#PATCHWireTransfersWireTransferId) | **Patch** /wire_transfers/{wireTransferId} | Update a wire transfer
[**POSTWireTransfers**](WireTransfersApi.md#POSTWireTransfers) | **Post** /wire_transfers | Create a wire transfer



## DELETEWireTransfersWireTransferId

> DELETEWireTransfersWireTransferId(ctx, wireTransferId).Execute()

Delete a wire transfer



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
    wireTransferId := "wireTransferId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WireTransfersApi.DELETEWireTransfersWireTransferId(context.Background(), wireTransferId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WireTransfersApi.DELETEWireTransfersWireTransferId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wireTransferId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEWireTransfersWireTransferIdRequest struct via the builder pattern


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


## GETWireTransfers

> WireTransferResponseList GETWireTransfers(ctx).Execute()

List all wire transfers



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
    resp, r, err := apiClient.WireTransfersApi.GETWireTransfers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WireTransfersApi.GETWireTransfers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETWireTransfers`: WireTransferResponseList
    fmt.Fprintf(os.Stdout, "Response from `WireTransfersApi.GETWireTransfers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETWireTransfersRequest struct via the builder pattern


### Return type

[**WireTransferResponseList**](WireTransferResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETWireTransfersWireTransferId

> WireTransferResponse GETWireTransfersWireTransferId(ctx, wireTransferId).Execute()

Retrieve a wire transfer



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
    wireTransferId := "wireTransferId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WireTransfersApi.GETWireTransfersWireTransferId(context.Background(), wireTransferId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WireTransfersApi.GETWireTransfersWireTransferId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETWireTransfersWireTransferId`: WireTransferResponse
    fmt.Fprintf(os.Stdout, "Response from `WireTransfersApi.GETWireTransfersWireTransferId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wireTransferId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETWireTransfersWireTransferIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WireTransferResponse**](WireTransferResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHWireTransfersWireTransferId

> WireTransferResponse PATCHWireTransfersWireTransferId(ctx, wireTransferId).WireTransferUpdate(wireTransferUpdate).Execute()

Update a wire transfer



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
    wireTransferUpdate := *openapiclient.NewWireTransferUpdate(*openapiclient.NewWireTransferUpdateData("Type_example", "XGZwpOSrWL", *openapiclient.NewAdyenPaymentCreateDataAttributes())) // WireTransferUpdate | 
    wireTransferId := "wireTransferId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WireTransfersApi.PATCHWireTransfersWireTransferId(context.Background(), wireTransferId).WireTransferUpdate(wireTransferUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WireTransfersApi.PATCHWireTransfersWireTransferId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHWireTransfersWireTransferId`: WireTransferResponse
    fmt.Fprintf(os.Stdout, "Response from `WireTransfersApi.PATCHWireTransfersWireTransferId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wireTransferId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHWireTransfersWireTransferIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wireTransferUpdate** | [**WireTransferUpdate**](WireTransferUpdate.md) |  | 


### Return type

[**WireTransferResponse**](WireTransferResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTWireTransfers

> WireTransferResponse POSTWireTransfers(ctx).WireTransferCreate(wireTransferCreate).Execute()

Create a wire transfer



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
    wireTransferCreate := *openapiclient.NewWireTransferCreate(*openapiclient.NewWireTransferCreateData("Type_example", *openapiclient.NewAdyenPaymentCreateDataAttributes())) // WireTransferCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WireTransfersApi.POSTWireTransfers(context.Background()).WireTransferCreate(wireTransferCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WireTransfersApi.POSTWireTransfers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTWireTransfers`: WireTransferResponse
    fmt.Fprintf(os.Stdout, "Response from `WireTransfersApi.POSTWireTransfers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTWireTransfersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wireTransferCreate** | [**WireTransferCreate**](WireTransferCreate.md) |  | 

### Return type

[**WireTransferResponse**](WireTransferResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

