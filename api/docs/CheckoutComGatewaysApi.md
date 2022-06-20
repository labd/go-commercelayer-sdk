# \CheckoutComGatewaysApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETECheckoutComGatewaysCheckoutComGatewayId**](CheckoutComGatewaysApi.md#DELETECheckoutComGatewaysCheckoutComGatewayId) | **Delete** /checkout_com_gateways/{checkoutComGatewayId} | Delete a checkout.com gateway
[**GETCheckoutComGateways**](CheckoutComGatewaysApi.md#GETCheckoutComGateways) | **Get** /checkout_com_gateways | List all checkout.com gateways
[**GETCheckoutComGatewaysCheckoutComGatewayId**](CheckoutComGatewaysApi.md#GETCheckoutComGatewaysCheckoutComGatewayId) | **Get** /checkout_com_gateways/{checkoutComGatewayId} | Retrieve a checkout.com gateway
[**PATCHCheckoutComGatewaysCheckoutComGatewayId**](CheckoutComGatewaysApi.md#PATCHCheckoutComGatewaysCheckoutComGatewayId) | **Patch** /checkout_com_gateways/{checkoutComGatewayId} | Update a checkout.com gateway
[**POSTCheckoutComGateways**](CheckoutComGatewaysApi.md#POSTCheckoutComGateways) | **Post** /checkout_com_gateways | Create a checkout.com gateway



## DELETECheckoutComGatewaysCheckoutComGatewayId

> DELETECheckoutComGatewaysCheckoutComGatewayId(ctx, checkoutComGatewayId).Execute()

Delete a checkout.com gateway



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
    checkoutComGatewayId := "checkoutComGatewayId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckoutComGatewaysApi.DELETECheckoutComGatewaysCheckoutComGatewayId(context.Background(), checkoutComGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckoutComGatewaysApi.DELETECheckoutComGatewaysCheckoutComGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkoutComGatewayId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETECheckoutComGatewaysCheckoutComGatewayIdRequest struct via the builder pattern


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


## GETCheckoutComGateways

> GETCheckoutComGateways(ctx).Execute()

List all checkout.com gateways



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
    resp, r, err := apiClient.CheckoutComGatewaysApi.GETCheckoutComGateways(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckoutComGatewaysApi.GETCheckoutComGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETCheckoutComGatewaysRequest struct via the builder pattern


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


## GETCheckoutComGatewaysCheckoutComGatewayId

> CheckoutComGateway GETCheckoutComGatewaysCheckoutComGatewayId(ctx, checkoutComGatewayId).Execute()

Retrieve a checkout.com gateway



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
    checkoutComGatewayId := "checkoutComGatewayId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckoutComGatewaysApi.GETCheckoutComGatewaysCheckoutComGatewayId(context.Background(), checkoutComGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckoutComGatewaysApi.GETCheckoutComGatewaysCheckoutComGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCheckoutComGatewaysCheckoutComGatewayId`: CheckoutComGateway
    fmt.Fprintf(os.Stdout, "Response from `CheckoutComGatewaysApi.GETCheckoutComGatewaysCheckoutComGatewayId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkoutComGatewayId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCheckoutComGatewaysCheckoutComGatewayIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CheckoutComGateway**](CheckoutComGateway.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHCheckoutComGatewaysCheckoutComGatewayId

> PATCHCheckoutComGatewaysCheckoutComGatewayId(ctx, checkoutComGatewayId).CheckoutComGatewayUpdate(checkoutComGatewayUpdate).Execute()

Update a checkout.com gateway



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
    checkoutComGatewayId := "checkoutComGatewayId_example" // string | The resource's id
    checkoutComGatewayUpdate := *openapiclient.NewCheckoutComGatewayUpdate(*openapiclient.NewCheckoutComGatewayUpdateData("checkout_com_gateways", "XGZwpOSrWL", *openapiclient.NewCheckoutComGatewayUpdateDataAttributes())) // CheckoutComGatewayUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckoutComGatewaysApi.PATCHCheckoutComGatewaysCheckoutComGatewayId(context.Background(), checkoutComGatewayId).CheckoutComGatewayUpdate(checkoutComGatewayUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckoutComGatewaysApi.PATCHCheckoutComGatewaysCheckoutComGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkoutComGatewayId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHCheckoutComGatewaysCheckoutComGatewayIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkoutComGatewayUpdate** | [**CheckoutComGatewayUpdate**](CheckoutComGatewayUpdate.md) |  | 

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


## POSTCheckoutComGateways

> POSTCheckoutComGateways(ctx).CheckoutComGatewayCreate(checkoutComGatewayCreate).Execute()

Create a checkout.com gateway



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
    checkoutComGatewayCreate := *openapiclient.NewCheckoutComGatewayCreate(*openapiclient.NewCheckoutComGatewayCreateData("checkout_com_gateways", *openapiclient.NewCheckoutComGatewayCreateDataAttributes("US payment gateway", "sk_test_xxxx-yyyy-zzzz", "pk_test_xxxx-yyyy-zzzz"))) // CheckoutComGatewayCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckoutComGatewaysApi.POSTCheckoutComGateways(context.Background()).CheckoutComGatewayCreate(checkoutComGatewayCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckoutComGatewaysApi.POSTCheckoutComGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTCheckoutComGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkoutComGatewayCreate** | [**CheckoutComGatewayCreate**](CheckoutComGatewayCreate.md) |  | 

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

