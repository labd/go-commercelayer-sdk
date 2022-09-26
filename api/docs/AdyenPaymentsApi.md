# \AdyenPaymentsApi

All URIs are relative to *https://}.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEAdyenPaymentsAdyenPaymentId**](AdyenPaymentsApi.md#DELETEAdyenPaymentsAdyenPaymentId) | **Delete** /adyen_payments/{adyenPaymentId} | Delete an adyen payment
[**GETAdyenGatewayIdAdyenPayments**](AdyenPaymentsApi.md#GETAdyenGatewayIdAdyenPayments) | **Get** /adyen_gateways/{adyenGatewayId}/adyen_payments | Retrieve the adyen payments associated to the adyen gateway
[**GETAdyenPayments**](AdyenPaymentsApi.md#GETAdyenPayments) | **Get** /adyen_payments | List all adyen payments
[**GETAdyenPaymentsAdyenPaymentId**](AdyenPaymentsApi.md#GETAdyenPaymentsAdyenPaymentId) | **Get** /adyen_payments/{adyenPaymentId} | Retrieve an adyen payment
[**PATCHAdyenPaymentsAdyenPaymentId**](AdyenPaymentsApi.md#PATCHAdyenPaymentsAdyenPaymentId) | **Patch** /adyen_payments/{adyenPaymentId} | Update an adyen payment
[**POSTAdyenPayments**](AdyenPaymentsApi.md#POSTAdyenPayments) | **Post** /adyen_payments | Create an adyen payment



## DELETEAdyenPaymentsAdyenPaymentId

> DELETEAdyenPaymentsAdyenPaymentId(ctx, adyenPaymentId).Execute()

Delete an adyen payment



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
    adyenPaymentId := "adyenPaymentId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdyenPaymentsApi.DELETEAdyenPaymentsAdyenPaymentId(context.Background(), adyenPaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdyenPaymentsApi.DELETEAdyenPaymentsAdyenPaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adyenPaymentId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEAdyenPaymentsAdyenPaymentIdRequest struct via the builder pattern


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


## GETAdyenGatewayIdAdyenPayments

> GETAdyenGatewayIdAdyenPayments(ctx, adyenGatewayId).Execute()

Retrieve the adyen payments associated to the adyen gateway



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
    adyenGatewayId := "adyenGatewayId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdyenPaymentsApi.GETAdyenGatewayIdAdyenPayments(context.Background(), adyenGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdyenPaymentsApi.GETAdyenGatewayIdAdyenPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adyenGatewayId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETAdyenGatewayIdAdyenPaymentsRequest struct via the builder pattern


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


## GETAdyenPayments

> AdyenPaymentResponseList GETAdyenPayments(ctx).Execute()

List all adyen payments



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
    resp, r, err := apiClient.AdyenPaymentsApi.GETAdyenPayments(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdyenPaymentsApi.GETAdyenPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETAdyenPayments`: AdyenPaymentResponseList
    fmt.Fprintf(os.Stdout, "Response from `AdyenPaymentsApi.GETAdyenPayments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETAdyenPaymentsRequest struct via the builder pattern


### Return type

[**AdyenPaymentResponseList**](AdyenPaymentResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETAdyenPaymentsAdyenPaymentId

> AdyenPaymentResponse GETAdyenPaymentsAdyenPaymentId(ctx, adyenPaymentId).Execute()

Retrieve an adyen payment



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
    adyenPaymentId := "adyenPaymentId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdyenPaymentsApi.GETAdyenPaymentsAdyenPaymentId(context.Background(), adyenPaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdyenPaymentsApi.GETAdyenPaymentsAdyenPaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETAdyenPaymentsAdyenPaymentId`: AdyenPaymentResponse
    fmt.Fprintf(os.Stdout, "Response from `AdyenPaymentsApi.GETAdyenPaymentsAdyenPaymentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adyenPaymentId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETAdyenPaymentsAdyenPaymentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdyenPaymentResponse**](AdyenPaymentResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHAdyenPaymentsAdyenPaymentId

> AdyenPaymentResponse PATCHAdyenPaymentsAdyenPaymentId(ctx, adyenPaymentId).AdyenPaymentUpdate(adyenPaymentUpdate).Execute()

Update an adyen payment



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
    adyenPaymentUpdate := *openapiclient.NewAdyenPaymentUpdate(*openapiclient.NewAdyenPaymentUpdateData("Type_example", "XGZwpOSrWL", *openapiclient.NewAdyenPaymentUpdateDataAttributes())) // AdyenPaymentUpdate | 
    adyenPaymentId := "adyenPaymentId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdyenPaymentsApi.PATCHAdyenPaymentsAdyenPaymentId(context.Background(), adyenPaymentId).AdyenPaymentUpdate(adyenPaymentUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdyenPaymentsApi.PATCHAdyenPaymentsAdyenPaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHAdyenPaymentsAdyenPaymentId`: AdyenPaymentResponse
    fmt.Fprintf(os.Stdout, "Response from `AdyenPaymentsApi.PATCHAdyenPaymentsAdyenPaymentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adyenPaymentId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHAdyenPaymentsAdyenPaymentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adyenPaymentUpdate** | [**AdyenPaymentUpdate**](AdyenPaymentUpdate.md) |  | 


### Return type

[**AdyenPaymentResponse**](AdyenPaymentResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTAdyenPayments

> AdyenPaymentResponse POSTAdyenPayments(ctx).AdyenPaymentCreate(adyenPaymentCreate).Execute()

Create an adyen payment



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
    adyenPaymentCreate := *openapiclient.NewAdyenPaymentCreate(*openapiclient.NewAdyenPaymentCreateData("Type_example", *openapiclient.NewAdyenPaymentCreateDataAttributes())) // AdyenPaymentCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdyenPaymentsApi.POSTAdyenPayments(context.Background()).AdyenPaymentCreate(adyenPaymentCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdyenPaymentsApi.POSTAdyenPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTAdyenPayments`: AdyenPaymentResponse
    fmt.Fprintf(os.Stdout, "Response from `AdyenPaymentsApi.POSTAdyenPayments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTAdyenPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adyenPaymentCreate** | [**AdyenPaymentCreate**](AdyenPaymentCreate.md) |  | 

### Return type

[**AdyenPaymentResponse**](AdyenPaymentResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

