# \StripeGatewaysApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEStripeGatewaysStripeGatewayId**](StripeGatewaysApi.md#DELETEStripeGatewaysStripeGatewayId) | **Delete** /stripe_gateways/{stripeGatewayId} | Delete a stripe gateway
[**GETStripeGateways**](StripeGatewaysApi.md#GETStripeGateways) | **Get** /stripe_gateways | List all stripe gateways
[**GETStripeGatewaysStripeGatewayId**](StripeGatewaysApi.md#GETStripeGatewaysStripeGatewayId) | **Get** /stripe_gateways/{stripeGatewayId} | Retrieve a stripe gateway
[**PATCHStripeGatewaysStripeGatewayId**](StripeGatewaysApi.md#PATCHStripeGatewaysStripeGatewayId) | **Patch** /stripe_gateways/{stripeGatewayId} | Update a stripe gateway
[**POSTStripeGateways**](StripeGatewaysApi.md#POSTStripeGateways) | **Post** /stripe_gateways | Create a stripe gateway



## DELETEStripeGatewaysStripeGatewayId

> DELETEStripeGatewaysStripeGatewayId(ctx, stripeGatewayId).Execute()

Delete a stripe gateway



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
    stripeGatewayId := "stripeGatewayId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StripeGatewaysApi.DELETEStripeGatewaysStripeGatewayId(context.Background(), stripeGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StripeGatewaysApi.DELETEStripeGatewaysStripeGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stripeGatewayId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEStripeGatewaysStripeGatewayIdRequest struct via the builder pattern


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


## GETStripeGateways

> GETStripeGateways(ctx).Execute()

List all stripe gateways



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
    resp, r, err := apiClient.StripeGatewaysApi.GETStripeGateways(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StripeGatewaysApi.GETStripeGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETStripeGatewaysRequest struct via the builder pattern


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


## GETStripeGatewaysStripeGatewayId

> StripeGateway GETStripeGatewaysStripeGatewayId(ctx, stripeGatewayId).Execute()

Retrieve a stripe gateway



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
    stripeGatewayId := "stripeGatewayId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StripeGatewaysApi.GETStripeGatewaysStripeGatewayId(context.Background(), stripeGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StripeGatewaysApi.GETStripeGatewaysStripeGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETStripeGatewaysStripeGatewayId`: StripeGateway
    fmt.Fprintf(os.Stdout, "Response from `StripeGatewaysApi.GETStripeGatewaysStripeGatewayId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stripeGatewayId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETStripeGatewaysStripeGatewayIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StripeGateway**](StripeGateway.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHStripeGatewaysStripeGatewayId

> PATCHStripeGatewaysStripeGatewayId(ctx, stripeGatewayId).StripeGatewayUpdate(stripeGatewayUpdate).Execute()

Update a stripe gateway



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
    stripeGatewayId := "stripeGatewayId_example" // string | The resource's id
    stripeGatewayUpdate := *openapiclient.NewStripeGatewayUpdate(*openapiclient.NewStripeGatewayUpdateData("stripe_gateways", "XGZwpOSrWL", *openapiclient.NewStripeGatewayUpdateDataAttributes())) // StripeGatewayUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StripeGatewaysApi.PATCHStripeGatewaysStripeGatewayId(context.Background(), stripeGatewayId).StripeGatewayUpdate(stripeGatewayUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StripeGatewaysApi.PATCHStripeGatewaysStripeGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stripeGatewayId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHStripeGatewaysStripeGatewayIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stripeGatewayUpdate** | [**StripeGatewayUpdate**](StripeGatewayUpdate.md) |  | 

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


## POSTStripeGateways

> POSTStripeGateways(ctx).StripeGatewayCreate(stripeGatewayCreate).Execute()

Create a stripe gateway



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
    stripeGatewayCreate := *openapiclient.NewStripeGatewayCreate(*openapiclient.NewStripeGatewayCreateData("stripe_gateways", *openapiclient.NewStripeGatewayCreateDataAttributes("US payment gateway", "sk_live_xxxx-yyyy-zzzz"))) // StripeGatewayCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StripeGatewaysApi.POSTStripeGateways(context.Background()).StripeGatewayCreate(stripeGatewayCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StripeGatewaysApi.POSTStripeGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTStripeGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stripeGatewayCreate** | [**StripeGatewayCreate**](StripeGatewayCreate.md) |  | 

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

