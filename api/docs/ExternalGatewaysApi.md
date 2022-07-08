# \ExternalGatewaysApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEExternalGatewaysExternalGatewayId**](ExternalGatewaysApi.md#DELETEExternalGatewaysExternalGatewayId) | **Delete** /external_gateways/{externalGatewayId} | Delete an external gateway
[**GETExternalGateways**](ExternalGatewaysApi.md#GETExternalGateways) | **Get** /external_gateways | List all external gateways
[**GETExternalGatewaysExternalGatewayId**](ExternalGatewaysApi.md#GETExternalGatewaysExternalGatewayId) | **Get** /external_gateways/{externalGatewayId} | Retrieve an external gateway
[**PATCHExternalGatewaysExternalGatewayId**](ExternalGatewaysApi.md#PATCHExternalGatewaysExternalGatewayId) | **Patch** /external_gateways/{externalGatewayId} | Update an external gateway
[**POSTExternalGateways**](ExternalGatewaysApi.md#POSTExternalGateways) | **Post** /external_gateways | Create an external gateway



## DELETEExternalGatewaysExternalGatewayId

> DELETEExternalGatewaysExternalGatewayId(ctx, externalGatewayId).Execute()

Delete an external gateway



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
    externalGatewayId := "externalGatewayId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalGatewaysApi.DELETEExternalGatewaysExternalGatewayId(context.Background(), externalGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalGatewaysApi.DELETEExternalGatewaysExternalGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalGatewayId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEExternalGatewaysExternalGatewayIdRequest struct via the builder pattern


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


## GETExternalGateways

> GETExternalGateways(ctx).Execute()

List all external gateways



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
    resp, r, err := apiClient.ExternalGatewaysApi.GETExternalGateways(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalGatewaysApi.GETExternalGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETExternalGatewaysRequest struct via the builder pattern


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


## GETExternalGatewaysExternalGatewayId

> ExternalGateway GETExternalGatewaysExternalGatewayId(ctx, externalGatewayId).Execute()

Retrieve an external gateway



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
    externalGatewayId := "externalGatewayId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalGatewaysApi.GETExternalGatewaysExternalGatewayId(context.Background(), externalGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalGatewaysApi.GETExternalGatewaysExternalGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETExternalGatewaysExternalGatewayId`: ExternalGateway
    fmt.Fprintf(os.Stdout, "Response from `ExternalGatewaysApi.GETExternalGatewaysExternalGatewayId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalGatewayId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETExternalGatewaysExternalGatewayIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalGateway**](ExternalGateway.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHExternalGatewaysExternalGatewayId

> PATCHExternalGatewaysExternalGatewayId(ctx, externalGatewayId).ExternalGatewayUpdate(externalGatewayUpdate).Execute()

Update an external gateway



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
    externalGatewayUpdate := *openapiclient.NewExternalGatewayUpdate(*openapiclient.NewExternalGatewayUpdateData("external_gateways", "XGZwpOSrWL", *openapiclient.NewExternalGatewayUpdateDataAttributes())) // ExternalGatewayUpdate | 
    externalGatewayId := "externalGatewayId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalGatewaysApi.PATCHExternalGatewaysExternalGatewayId(context.Background(), externalGatewayId).ExternalGatewayUpdate(externalGatewayUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalGatewaysApi.PATCHExternalGatewaysExternalGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalGatewayId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHExternalGatewaysExternalGatewayIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalGatewayUpdate** | [**ExternalGatewayUpdate**](ExternalGatewayUpdate.md) |  | 


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


## POSTExternalGateways

> POSTExternalGateways(ctx).ExternalGatewayCreate(externalGatewayCreate).Execute()

Create an external gateway



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
    externalGatewayCreate := *openapiclient.NewExternalGatewayCreate(*openapiclient.NewExternalGatewayCreateData("external_gateways", *openapiclient.NewExternalGatewayCreateDataAttributes("US payment gateway"))) // ExternalGatewayCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalGatewaysApi.POSTExternalGateways(context.Background()).ExternalGatewayCreate(externalGatewayCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalGatewaysApi.POSTExternalGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTExternalGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalGatewayCreate** | [**ExternalGatewayCreate**](ExternalGatewayCreate.md) |  | 

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

