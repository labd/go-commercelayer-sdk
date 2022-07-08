# \KlarnaPaymentsApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEKlarnaPaymentsKlarnaPaymentId**](KlarnaPaymentsApi.md#DELETEKlarnaPaymentsKlarnaPaymentId) | **Delete** /klarna_payments/{klarnaPaymentId} | Delete a klarna payment
[**GETKlarnaGatewayIdKlarnaPayments**](KlarnaPaymentsApi.md#GETKlarnaGatewayIdKlarnaPayments) | **Get** /klarna_gateways/{klarnaGatewayId}/klarna_payments | Retrieve the klarna payments associated to the klarna gateway
[**GETKlarnaPayments**](KlarnaPaymentsApi.md#GETKlarnaPayments) | **Get** /klarna_payments | List all klarna payments
[**GETKlarnaPaymentsKlarnaPaymentId**](KlarnaPaymentsApi.md#GETKlarnaPaymentsKlarnaPaymentId) | **Get** /klarna_payments/{klarnaPaymentId} | Retrieve a klarna payment
[**PATCHKlarnaPaymentsKlarnaPaymentId**](KlarnaPaymentsApi.md#PATCHKlarnaPaymentsKlarnaPaymentId) | **Patch** /klarna_payments/{klarnaPaymentId} | Update a klarna payment
[**POSTKlarnaPayments**](KlarnaPaymentsApi.md#POSTKlarnaPayments) | **Post** /klarna_payments | Create a klarna payment



## DELETEKlarnaPaymentsKlarnaPaymentId

> DELETEKlarnaPaymentsKlarnaPaymentId(ctx, klarnaPaymentId).Execute()

Delete a klarna payment



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
    klarnaPaymentId := "klarnaPaymentId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KlarnaPaymentsApi.DELETEKlarnaPaymentsKlarnaPaymentId(context.Background(), klarnaPaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KlarnaPaymentsApi.DELETEKlarnaPaymentsKlarnaPaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**klarnaPaymentId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEKlarnaPaymentsKlarnaPaymentIdRequest struct via the builder pattern


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


## GETKlarnaGatewayIdKlarnaPayments

> GETKlarnaGatewayIdKlarnaPayments(ctx, klarnaGatewayId).Execute()

Retrieve the klarna payments associated to the klarna gateway



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
    klarnaGatewayId := "klarnaGatewayId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KlarnaPaymentsApi.GETKlarnaGatewayIdKlarnaPayments(context.Background(), klarnaGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KlarnaPaymentsApi.GETKlarnaGatewayIdKlarnaPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**klarnaGatewayId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETKlarnaGatewayIdKlarnaPaymentsRequest struct via the builder pattern


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


## GETKlarnaPayments

> GETKlarnaPayments(ctx).Execute()

List all klarna payments



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
    resp, r, err := apiClient.KlarnaPaymentsApi.GETKlarnaPayments(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KlarnaPaymentsApi.GETKlarnaPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETKlarnaPaymentsRequest struct via the builder pattern


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


## GETKlarnaPaymentsKlarnaPaymentId

> KlarnaPayment GETKlarnaPaymentsKlarnaPaymentId(ctx, klarnaPaymentId).Execute()

Retrieve a klarna payment



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
    klarnaPaymentId := "klarnaPaymentId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KlarnaPaymentsApi.GETKlarnaPaymentsKlarnaPaymentId(context.Background(), klarnaPaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KlarnaPaymentsApi.GETKlarnaPaymentsKlarnaPaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETKlarnaPaymentsKlarnaPaymentId`: KlarnaPayment
    fmt.Fprintf(os.Stdout, "Response from `KlarnaPaymentsApi.GETKlarnaPaymentsKlarnaPaymentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**klarnaPaymentId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETKlarnaPaymentsKlarnaPaymentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KlarnaPayment**](KlarnaPayment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHKlarnaPaymentsKlarnaPaymentId

> PATCHKlarnaPaymentsKlarnaPaymentId(ctx, klarnaPaymentId).KlarnaPaymentUpdate(klarnaPaymentUpdate).Execute()

Update a klarna payment



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
    klarnaPaymentUpdate := *openapiclient.NewKlarnaPaymentUpdate(*openapiclient.NewKlarnaPaymentUpdateData("klarna_payments", "XGZwpOSrWL", *openapiclient.NewKlarnaPaymentUpdateDataAttributes())) // KlarnaPaymentUpdate | 
    klarnaPaymentId := "klarnaPaymentId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KlarnaPaymentsApi.PATCHKlarnaPaymentsKlarnaPaymentId(context.Background(), klarnaPaymentId).KlarnaPaymentUpdate(klarnaPaymentUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KlarnaPaymentsApi.PATCHKlarnaPaymentsKlarnaPaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**klarnaPaymentId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHKlarnaPaymentsKlarnaPaymentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **klarnaPaymentUpdate** | [**KlarnaPaymentUpdate**](KlarnaPaymentUpdate.md) |  | 


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


## POSTKlarnaPayments

> POSTKlarnaPayments(ctx).KlarnaPaymentCreate(klarnaPaymentCreate).Execute()

Create a klarna payment



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
    klarnaPaymentCreate := *openapiclient.NewKlarnaPaymentCreate(*openapiclient.NewKlarnaPaymentCreateData("klarna_payments", *openapiclient.NewAdyenPaymentCreateDataAttributes())) // KlarnaPaymentCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KlarnaPaymentsApi.POSTKlarnaPayments(context.Background()).KlarnaPaymentCreate(klarnaPaymentCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KlarnaPaymentsApi.POSTKlarnaPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTKlarnaPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **klarnaPaymentCreate** | [**KlarnaPaymentCreate**](KlarnaPaymentCreate.md) |  | 

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

