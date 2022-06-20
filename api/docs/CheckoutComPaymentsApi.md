# \CheckoutComPaymentsApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETECheckoutComPaymentsCheckoutComPaymentId**](CheckoutComPaymentsApi.md#DELETECheckoutComPaymentsCheckoutComPaymentId) | **Delete** /checkout_com_payments/{checkoutComPaymentId} | Delete a checkout.com payment
[**GETCheckoutComGatewayIdCheckoutComPayments**](CheckoutComPaymentsApi.md#GETCheckoutComGatewayIdCheckoutComPayments) | **Get** /checkout_com_gateways/{checkoutComGatewayId}/checkout_com_payments | Retrieve the checkout com payments associated to the checkout.com gateway
[**GETCheckoutComPayments**](CheckoutComPaymentsApi.md#GETCheckoutComPayments) | **Get** /checkout_com_payments | List all checkout.com payments
[**GETCheckoutComPaymentsCheckoutComPaymentId**](CheckoutComPaymentsApi.md#GETCheckoutComPaymentsCheckoutComPaymentId) | **Get** /checkout_com_payments/{checkoutComPaymentId} | Retrieve a checkout.com payment
[**PATCHCheckoutComPaymentsCheckoutComPaymentId**](CheckoutComPaymentsApi.md#PATCHCheckoutComPaymentsCheckoutComPaymentId) | **Patch** /checkout_com_payments/{checkoutComPaymentId} | Update a checkout.com payment
[**POSTCheckoutComPayments**](CheckoutComPaymentsApi.md#POSTCheckoutComPayments) | **Post** /checkout_com_payments | Create a checkout.com payment



## DELETECheckoutComPaymentsCheckoutComPaymentId

> DELETECheckoutComPaymentsCheckoutComPaymentId(ctx, checkoutComPaymentId).Execute()

Delete a checkout.com payment



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
    checkoutComPaymentId := "checkoutComPaymentId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckoutComPaymentsApi.DELETECheckoutComPaymentsCheckoutComPaymentId(context.Background(), checkoutComPaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckoutComPaymentsApi.DELETECheckoutComPaymentsCheckoutComPaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkoutComPaymentId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETECheckoutComPaymentsCheckoutComPaymentIdRequest struct via the builder pattern


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


## GETCheckoutComGatewayIdCheckoutComPayments

> GETCheckoutComGatewayIdCheckoutComPayments(ctx, checkoutComGatewayId).Execute()

Retrieve the checkout com payments associated to the checkout.com gateway



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
    resp, r, err := apiClient.CheckoutComPaymentsApi.GETCheckoutComGatewayIdCheckoutComPayments(context.Background(), checkoutComGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckoutComPaymentsApi.GETCheckoutComGatewayIdCheckoutComPayments``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETCheckoutComGatewayIdCheckoutComPaymentsRequest struct via the builder pattern


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


## GETCheckoutComPayments

> GETCheckoutComPayments(ctx).Execute()

List all checkout.com payments



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
    resp, r, err := apiClient.CheckoutComPaymentsApi.GETCheckoutComPayments(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckoutComPaymentsApi.GETCheckoutComPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETCheckoutComPaymentsRequest struct via the builder pattern


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


## GETCheckoutComPaymentsCheckoutComPaymentId

> CheckoutComPayment GETCheckoutComPaymentsCheckoutComPaymentId(ctx, checkoutComPaymentId).Execute()

Retrieve a checkout.com payment



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
    checkoutComPaymentId := "checkoutComPaymentId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckoutComPaymentsApi.GETCheckoutComPaymentsCheckoutComPaymentId(context.Background(), checkoutComPaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckoutComPaymentsApi.GETCheckoutComPaymentsCheckoutComPaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCheckoutComPaymentsCheckoutComPaymentId`: CheckoutComPayment
    fmt.Fprintf(os.Stdout, "Response from `CheckoutComPaymentsApi.GETCheckoutComPaymentsCheckoutComPaymentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkoutComPaymentId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCheckoutComPaymentsCheckoutComPaymentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CheckoutComPayment**](CheckoutComPayment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHCheckoutComPaymentsCheckoutComPaymentId

> PATCHCheckoutComPaymentsCheckoutComPaymentId(ctx, checkoutComPaymentId).CheckoutComPaymentUpdate(checkoutComPaymentUpdate).Execute()

Update a checkout.com payment



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
    checkoutComPaymentId := "checkoutComPaymentId_example" // string | The resource's id
    checkoutComPaymentUpdate := *openapiclient.NewCheckoutComPaymentUpdate(*openapiclient.NewCheckoutComPaymentUpdateData("checkout_com_payments", "XGZwpOSrWL", *openapiclient.NewCheckoutComPaymentUpdateDataAttributes())) // CheckoutComPaymentUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckoutComPaymentsApi.PATCHCheckoutComPaymentsCheckoutComPaymentId(context.Background(), checkoutComPaymentId).CheckoutComPaymentUpdate(checkoutComPaymentUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckoutComPaymentsApi.PATCHCheckoutComPaymentsCheckoutComPaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkoutComPaymentId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHCheckoutComPaymentsCheckoutComPaymentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkoutComPaymentUpdate** | [**CheckoutComPaymentUpdate**](CheckoutComPaymentUpdate.md) |  | 

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


## POSTCheckoutComPayments

> POSTCheckoutComPayments(ctx).CheckoutComPaymentCreate(checkoutComPaymentCreate).Execute()

Create a checkout.com payment



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
    checkoutComPaymentCreate := *openapiclient.NewCheckoutComPaymentCreate(*openapiclient.NewCheckoutComPaymentCreateData("checkout_com_payments", *openapiclient.NewCheckoutComPaymentCreateDataAttributes("token", "tok_4gzeau5o2uqubbk6fufs3m7p54"))) // CheckoutComPaymentCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckoutComPaymentsApi.POSTCheckoutComPayments(context.Background()).CheckoutComPaymentCreate(checkoutComPaymentCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckoutComPaymentsApi.POSTCheckoutComPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTCheckoutComPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkoutComPaymentCreate** | [**CheckoutComPaymentCreate**](CheckoutComPaymentCreate.md) |  | 

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

