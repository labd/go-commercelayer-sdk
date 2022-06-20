# \StripePaymentsApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEStripePaymentsStripePaymentId**](StripePaymentsApi.md#DELETEStripePaymentsStripePaymentId) | **Delete** /stripe_payments/{stripePaymentId} | Delete a stripe payment
[**GETStripeGatewayIdStripePayments**](StripePaymentsApi.md#GETStripeGatewayIdStripePayments) | **Get** /stripe_gateways/{stripeGatewayId}/stripe_payments | Retrieve the stripe payments associated to the stripe gateway
[**GETStripePayments**](StripePaymentsApi.md#GETStripePayments) | **Get** /stripe_payments | List all stripe payments
[**GETStripePaymentsStripePaymentId**](StripePaymentsApi.md#GETStripePaymentsStripePaymentId) | **Get** /stripe_payments/{stripePaymentId} | Retrieve a stripe payment
[**PATCHStripePaymentsStripePaymentId**](StripePaymentsApi.md#PATCHStripePaymentsStripePaymentId) | **Patch** /stripe_payments/{stripePaymentId} | Update a stripe payment
[**POSTStripePayments**](StripePaymentsApi.md#POSTStripePayments) | **Post** /stripe_payments | Create a stripe payment



## DELETEStripePaymentsStripePaymentId

> DELETEStripePaymentsStripePaymentId(ctx, stripePaymentId).Execute()

Delete a stripe payment



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
    stripePaymentId := "stripePaymentId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StripePaymentsApi.DELETEStripePaymentsStripePaymentId(context.Background(), stripePaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StripePaymentsApi.DELETEStripePaymentsStripePaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stripePaymentId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEStripePaymentsStripePaymentIdRequest struct via the builder pattern


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


## GETStripeGatewayIdStripePayments

> GETStripeGatewayIdStripePayments(ctx, stripeGatewayId).Execute()

Retrieve the stripe payments associated to the stripe gateway



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
    resp, r, err := apiClient.StripePaymentsApi.GETStripeGatewayIdStripePayments(context.Background(), stripeGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StripePaymentsApi.GETStripeGatewayIdStripePayments``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETStripeGatewayIdStripePaymentsRequest struct via the builder pattern


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


## GETStripePayments

> GETStripePayments(ctx).Execute()

List all stripe payments



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
    resp, r, err := apiClient.StripePaymentsApi.GETStripePayments(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StripePaymentsApi.GETStripePayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETStripePaymentsRequest struct via the builder pattern


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


## GETStripePaymentsStripePaymentId

> StripePayment GETStripePaymentsStripePaymentId(ctx, stripePaymentId).Execute()

Retrieve a stripe payment



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
    stripePaymentId := "stripePaymentId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StripePaymentsApi.GETStripePaymentsStripePaymentId(context.Background(), stripePaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StripePaymentsApi.GETStripePaymentsStripePaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETStripePaymentsStripePaymentId`: StripePayment
    fmt.Fprintf(os.Stdout, "Response from `StripePaymentsApi.GETStripePaymentsStripePaymentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stripePaymentId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETStripePaymentsStripePaymentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StripePayment**](StripePayment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHStripePaymentsStripePaymentId

> PATCHStripePaymentsStripePaymentId(ctx, stripePaymentId).StripePaymentUpdate(stripePaymentUpdate).Execute()

Update a stripe payment



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
    stripePaymentId := "stripePaymentId_example" // string | The resource's id
    stripePaymentUpdate := *openapiclient.NewStripePaymentUpdate(*openapiclient.NewStripePaymentUpdateData("stripe_payments", "XGZwpOSrWL", *openapiclient.NewStripePaymentUpdateDataAttributes())) // StripePaymentUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StripePaymentsApi.PATCHStripePaymentsStripePaymentId(context.Background(), stripePaymentId).StripePaymentUpdate(stripePaymentUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StripePaymentsApi.PATCHStripePaymentsStripePaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stripePaymentId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHStripePaymentsStripePaymentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stripePaymentUpdate** | [**StripePaymentUpdate**](StripePaymentUpdate.md) |  | 

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


## POSTStripePayments

> POSTStripePayments(ctx).StripePaymentCreate(stripePaymentCreate).Execute()

Create a stripe payment



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
    stripePaymentCreate := *openapiclient.NewStripePaymentCreate(*openapiclient.NewStripePaymentCreateData("stripe_payments", *openapiclient.NewStripePaymentCreateDataAttributes())) // StripePaymentCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StripePaymentsApi.POSTStripePayments(context.Background()).StripePaymentCreate(stripePaymentCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StripePaymentsApi.POSTStripePayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTStripePaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stripePaymentCreate** | [**StripePaymentCreate**](StripePaymentCreate.md) |  | 

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

