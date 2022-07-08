# \PaypalPaymentsApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEPaypalPaymentsPaypalPaymentId**](PaypalPaymentsApi.md#DELETEPaypalPaymentsPaypalPaymentId) | **Delete** /paypal_payments/{paypalPaymentId} | Delete a paypal payment
[**GETPaypalGatewayIdPaypalPayments**](PaypalPaymentsApi.md#GETPaypalGatewayIdPaypalPayments) | **Get** /paypal_gateways/{paypalGatewayId}/paypal_payments | Retrieve the paypal payments associated to the paypal gateway
[**GETPaypalPayments**](PaypalPaymentsApi.md#GETPaypalPayments) | **Get** /paypal_payments | List all paypal payments
[**GETPaypalPaymentsPaypalPaymentId**](PaypalPaymentsApi.md#GETPaypalPaymentsPaypalPaymentId) | **Get** /paypal_payments/{paypalPaymentId} | Retrieve a paypal payment
[**PATCHPaypalPaymentsPaypalPaymentId**](PaypalPaymentsApi.md#PATCHPaypalPaymentsPaypalPaymentId) | **Patch** /paypal_payments/{paypalPaymentId} | Update a paypal payment
[**POSTPaypalPayments**](PaypalPaymentsApi.md#POSTPaypalPayments) | **Post** /paypal_payments | Create a paypal payment



## DELETEPaypalPaymentsPaypalPaymentId

> DELETEPaypalPaymentsPaypalPaymentId(ctx, paypalPaymentId).Execute()

Delete a paypal payment



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
    paypalPaymentId := "paypalPaymentId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaypalPaymentsApi.DELETEPaypalPaymentsPaypalPaymentId(context.Background(), paypalPaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaypalPaymentsApi.DELETEPaypalPaymentsPaypalPaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paypalPaymentId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEPaypalPaymentsPaypalPaymentIdRequest struct via the builder pattern


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


## GETPaypalGatewayIdPaypalPayments

> GETPaypalGatewayIdPaypalPayments(ctx, paypalGatewayId).Execute()

Retrieve the paypal payments associated to the paypal gateway



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
    paypalGatewayId := "paypalGatewayId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaypalPaymentsApi.GETPaypalGatewayIdPaypalPayments(context.Background(), paypalGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaypalPaymentsApi.GETPaypalGatewayIdPaypalPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paypalGatewayId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPaypalGatewayIdPaypalPaymentsRequest struct via the builder pattern


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


## GETPaypalPayments

> GETPaypalPayments(ctx).Execute()

List all paypal payments



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
    resp, r, err := apiClient.PaypalPaymentsApi.GETPaypalPayments(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaypalPaymentsApi.GETPaypalPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETPaypalPaymentsRequest struct via the builder pattern


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


## GETPaypalPaymentsPaypalPaymentId

> PaypalPayment GETPaypalPaymentsPaypalPaymentId(ctx, paypalPaymentId).Execute()

Retrieve a paypal payment



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
    paypalPaymentId := "paypalPaymentId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaypalPaymentsApi.GETPaypalPaymentsPaypalPaymentId(context.Background(), paypalPaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaypalPaymentsApi.GETPaypalPaymentsPaypalPaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPaypalPaymentsPaypalPaymentId`: PaypalPayment
    fmt.Fprintf(os.Stdout, "Response from `PaypalPaymentsApi.GETPaypalPaymentsPaypalPaymentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paypalPaymentId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPaypalPaymentsPaypalPaymentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaypalPayment**](PaypalPayment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHPaypalPaymentsPaypalPaymentId

> PATCHPaypalPaymentsPaypalPaymentId(ctx, paypalPaymentId).PaypalPaymentUpdate(paypalPaymentUpdate).Execute()

Update a paypal payment



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
    paypalPaymentUpdate := *openapiclient.NewPaypalPaymentUpdate(*openapiclient.NewPaypalPaymentUpdateData("paypal_payments", "XGZwpOSrWL", *openapiclient.NewPaypalPaymentUpdateDataAttributes())) // PaypalPaymentUpdate | 
    paypalPaymentId := "paypalPaymentId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaypalPaymentsApi.PATCHPaypalPaymentsPaypalPaymentId(context.Background(), paypalPaymentId).PaypalPaymentUpdate(paypalPaymentUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaypalPaymentsApi.PATCHPaypalPaymentsPaypalPaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paypalPaymentId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHPaypalPaymentsPaypalPaymentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paypalPaymentUpdate** | [**PaypalPaymentUpdate**](PaypalPaymentUpdate.md) |  | 


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


## POSTPaypalPayments

> POSTPaypalPayments(ctx).PaypalPaymentCreate(paypalPaymentCreate).Execute()

Create a paypal payment



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
    paypalPaymentCreate := *openapiclient.NewPaypalPaymentCreate(*openapiclient.NewPaypalPaymentCreateData("paypal_payments", *openapiclient.NewPaypalPaymentCreateDataAttributes("https://yourdomain.com/thankyou", "https://yourdomain.com/checkout/payment"))) // PaypalPaymentCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaypalPaymentsApi.POSTPaypalPayments(context.Background()).PaypalPaymentCreate(paypalPaymentCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaypalPaymentsApi.POSTPaypalPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTPaypalPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paypalPaymentCreate** | [**PaypalPaymentCreate**](PaypalPaymentCreate.md) |  | 

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

