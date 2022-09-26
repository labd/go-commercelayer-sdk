# \KlarnaGatewaysApi

All URIs are relative to *https://}.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEKlarnaGatewaysKlarnaGatewayId**](KlarnaGatewaysApi.md#DELETEKlarnaGatewaysKlarnaGatewayId) | **Delete** /klarna_gateways/{klarnaGatewayId} | Delete a klarna gateway
[**GETKlarnaGateways**](KlarnaGatewaysApi.md#GETKlarnaGateways) | **Get** /klarna_gateways | List all klarna gateways
[**GETKlarnaGatewaysKlarnaGatewayId**](KlarnaGatewaysApi.md#GETKlarnaGatewaysKlarnaGatewayId) | **Get** /klarna_gateways/{klarnaGatewayId} | Retrieve a klarna gateway
[**PATCHKlarnaGatewaysKlarnaGatewayId**](KlarnaGatewaysApi.md#PATCHKlarnaGatewaysKlarnaGatewayId) | **Patch** /klarna_gateways/{klarnaGatewayId} | Update a klarna gateway
[**POSTKlarnaGateways**](KlarnaGatewaysApi.md#POSTKlarnaGateways) | **Post** /klarna_gateways | Create a klarna gateway



## DELETEKlarnaGatewaysKlarnaGatewayId

> DELETEKlarnaGatewaysKlarnaGatewayId(ctx, klarnaGatewayId).Execute()

Delete a klarna gateway



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
    resp, r, err := apiClient.KlarnaGatewaysApi.DELETEKlarnaGatewaysKlarnaGatewayId(context.Background(), klarnaGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KlarnaGatewaysApi.DELETEKlarnaGatewaysKlarnaGatewayId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDELETEKlarnaGatewaysKlarnaGatewayIdRequest struct via the builder pattern


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


## GETKlarnaGateways

> KlarnaGatewayResponseList GETKlarnaGateways(ctx).Execute()

List all klarna gateways



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
    resp, r, err := apiClient.KlarnaGatewaysApi.GETKlarnaGateways(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KlarnaGatewaysApi.GETKlarnaGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETKlarnaGateways`: KlarnaGatewayResponseList
    fmt.Fprintf(os.Stdout, "Response from `KlarnaGatewaysApi.GETKlarnaGateways`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETKlarnaGatewaysRequest struct via the builder pattern


### Return type

[**KlarnaGatewayResponseList**](KlarnaGatewayResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETKlarnaGatewaysKlarnaGatewayId

> KlarnaGatewayResponse GETKlarnaGatewaysKlarnaGatewayId(ctx, klarnaGatewayId).Execute()

Retrieve a klarna gateway



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
    resp, r, err := apiClient.KlarnaGatewaysApi.GETKlarnaGatewaysKlarnaGatewayId(context.Background(), klarnaGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KlarnaGatewaysApi.GETKlarnaGatewaysKlarnaGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETKlarnaGatewaysKlarnaGatewayId`: KlarnaGatewayResponse
    fmt.Fprintf(os.Stdout, "Response from `KlarnaGatewaysApi.GETKlarnaGatewaysKlarnaGatewayId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**klarnaGatewayId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETKlarnaGatewaysKlarnaGatewayIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KlarnaGatewayResponse**](KlarnaGatewayResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHKlarnaGatewaysKlarnaGatewayId

> KlarnaGatewayResponse PATCHKlarnaGatewaysKlarnaGatewayId(ctx, klarnaGatewayId).KlarnaGatewayUpdate(klarnaGatewayUpdate).Execute()

Update a klarna gateway



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
    klarnaGatewayUpdate := *openapiclient.NewKlarnaGatewayUpdate(*openapiclient.NewKlarnaGatewayUpdateData("Type_example", "XGZwpOSrWL", *openapiclient.NewKlarnaGatewayUpdateDataAttributes())) // KlarnaGatewayUpdate | 
    klarnaGatewayId := "klarnaGatewayId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KlarnaGatewaysApi.PATCHKlarnaGatewaysKlarnaGatewayId(context.Background(), klarnaGatewayId).KlarnaGatewayUpdate(klarnaGatewayUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KlarnaGatewaysApi.PATCHKlarnaGatewaysKlarnaGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHKlarnaGatewaysKlarnaGatewayId`: KlarnaGatewayResponse
    fmt.Fprintf(os.Stdout, "Response from `KlarnaGatewaysApi.PATCHKlarnaGatewaysKlarnaGatewayId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**klarnaGatewayId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHKlarnaGatewaysKlarnaGatewayIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **klarnaGatewayUpdate** | [**KlarnaGatewayUpdate**](KlarnaGatewayUpdate.md) |  | 


### Return type

[**KlarnaGatewayResponse**](KlarnaGatewayResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTKlarnaGateways

> KlarnaGatewayResponse POSTKlarnaGateways(ctx).KlarnaGatewayCreate(klarnaGatewayCreate).Execute()

Create a klarna gateway



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
    klarnaGatewayCreate := *openapiclient.NewKlarnaGatewayCreate(*openapiclient.NewKlarnaGatewayCreateData("Type_example", *openapiclient.NewKlarnaGatewayCreateDataAttributes("US payment gateway", "EU", "xxxx-yyyy-zzzz", "xxxx-yyyy-zzzz"))) // KlarnaGatewayCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KlarnaGatewaysApi.POSTKlarnaGateways(context.Background()).KlarnaGatewayCreate(klarnaGatewayCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KlarnaGatewaysApi.POSTKlarnaGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTKlarnaGateways`: KlarnaGatewayResponse
    fmt.Fprintf(os.Stdout, "Response from `KlarnaGatewaysApi.POSTKlarnaGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTKlarnaGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **klarnaGatewayCreate** | [**KlarnaGatewayCreate**](KlarnaGatewayCreate.md) |  | 

### Return type

[**KlarnaGatewayResponse**](KlarnaGatewayResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

