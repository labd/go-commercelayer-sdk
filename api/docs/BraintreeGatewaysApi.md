# \BraintreeGatewaysApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEBraintreeGatewaysBraintreeGatewayId**](BraintreeGatewaysApi.md#DELETEBraintreeGatewaysBraintreeGatewayId) | **Delete** /braintree_gateways/{braintreeGatewayId} | Delete a braintree gateway
[**GETBraintreeGateways**](BraintreeGatewaysApi.md#GETBraintreeGateways) | **Get** /braintree_gateways | List all braintree gateways
[**GETBraintreeGatewaysBraintreeGatewayId**](BraintreeGatewaysApi.md#GETBraintreeGatewaysBraintreeGatewayId) | **Get** /braintree_gateways/{braintreeGatewayId} | Retrieve a braintree gateway
[**PATCHBraintreeGatewaysBraintreeGatewayId**](BraintreeGatewaysApi.md#PATCHBraintreeGatewaysBraintreeGatewayId) | **Patch** /braintree_gateways/{braintreeGatewayId} | Update a braintree gateway
[**POSTBraintreeGateways**](BraintreeGatewaysApi.md#POSTBraintreeGateways) | **Post** /braintree_gateways | Create a braintree gateway



## DELETEBraintreeGatewaysBraintreeGatewayId

> DELETEBraintreeGatewaysBraintreeGatewayId(ctx, braintreeGatewayId).Execute()

Delete a braintree gateway



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
    braintreeGatewayId := "braintreeGatewayId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BraintreeGatewaysApi.DELETEBraintreeGatewaysBraintreeGatewayId(context.Background(), braintreeGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BraintreeGatewaysApi.DELETEBraintreeGatewaysBraintreeGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**braintreeGatewayId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEBraintreeGatewaysBraintreeGatewayIdRequest struct via the builder pattern


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


## GETBraintreeGateways

> GETBraintreeGateways(ctx).Execute()

List all braintree gateways



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
    resp, r, err := apiClient.BraintreeGatewaysApi.GETBraintreeGateways(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BraintreeGatewaysApi.GETBraintreeGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETBraintreeGatewaysRequest struct via the builder pattern


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


## GETBraintreeGatewaysBraintreeGatewayId

> BraintreeGateway GETBraintreeGatewaysBraintreeGatewayId(ctx, braintreeGatewayId).Execute()

Retrieve a braintree gateway



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
    braintreeGatewayId := "braintreeGatewayId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BraintreeGatewaysApi.GETBraintreeGatewaysBraintreeGatewayId(context.Background(), braintreeGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BraintreeGatewaysApi.GETBraintreeGatewaysBraintreeGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETBraintreeGatewaysBraintreeGatewayId`: BraintreeGateway
    fmt.Fprintf(os.Stdout, "Response from `BraintreeGatewaysApi.GETBraintreeGatewaysBraintreeGatewayId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**braintreeGatewayId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETBraintreeGatewaysBraintreeGatewayIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BraintreeGateway**](BraintreeGateway.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHBraintreeGatewaysBraintreeGatewayId

> PATCHBraintreeGatewaysBraintreeGatewayId(ctx, braintreeGatewayId).BraintreeGatewayUpdate(braintreeGatewayUpdate).Execute()

Update a braintree gateway



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
    braintreeGatewayUpdate := *openapiclient.NewBraintreeGatewayUpdate(*openapiclient.NewBraintreeGatewayUpdateData("braintree_gateways", "XGZwpOSrWL", *openapiclient.NewBraintreeGatewayUpdateDataAttributes())) // BraintreeGatewayUpdate | 
    braintreeGatewayId := "braintreeGatewayId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BraintreeGatewaysApi.PATCHBraintreeGatewaysBraintreeGatewayId(context.Background(), braintreeGatewayId).BraintreeGatewayUpdate(braintreeGatewayUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BraintreeGatewaysApi.PATCHBraintreeGatewaysBraintreeGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**braintreeGatewayId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHBraintreeGatewaysBraintreeGatewayIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **braintreeGatewayUpdate** | [**BraintreeGatewayUpdate**](BraintreeGatewayUpdate.md) |  | 


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


## POSTBraintreeGateways

> POSTBraintreeGateways(ctx).BraintreeGatewayCreate(braintreeGatewayCreate).Execute()

Create a braintree gateway



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
    braintreeGatewayCreate := *openapiclient.NewBraintreeGatewayCreate(*openapiclient.NewBraintreeGatewayCreateData("braintree_gateways", *openapiclient.NewBraintreeGatewayCreateDataAttributes("US payment gateway", "xxxx-yyyy-zzzz", "xxxx-yyyy-zzzz", "xxxx-yyyy-zzzz", "xxxx-yyyy-zzzz"))) // BraintreeGatewayCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BraintreeGatewaysApi.POSTBraintreeGateways(context.Background()).BraintreeGatewayCreate(braintreeGatewayCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BraintreeGatewaysApi.POSTBraintreeGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTBraintreeGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **braintreeGatewayCreate** | [**BraintreeGatewayCreate**](BraintreeGatewayCreate.md) |  | 

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

