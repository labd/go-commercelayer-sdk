# \SatispayGatewaysApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETESatispayGatewaysSatispayGatewayId**](SatispayGatewaysApi.md#DELETESatispayGatewaysSatispayGatewayId) | **Delete** /satispay_gateways/{satispayGatewayId} | Delete a satispay gateway
[**GETSatispayGateways**](SatispayGatewaysApi.md#GETSatispayGateways) | **Get** /satispay_gateways | List all satispay gateways
[**GETSatispayGatewaysSatispayGatewayId**](SatispayGatewaysApi.md#GETSatispayGatewaysSatispayGatewayId) | **Get** /satispay_gateways/{satispayGatewayId} | Retrieve a satispay gateway
[**PATCHSatispayGatewaysSatispayGatewayId**](SatispayGatewaysApi.md#PATCHSatispayGatewaysSatispayGatewayId) | **Patch** /satispay_gateways/{satispayGatewayId} | Update a satispay gateway
[**POSTSatispayGateways**](SatispayGatewaysApi.md#POSTSatispayGateways) | **Post** /satispay_gateways | Create a satispay gateway



## DELETESatispayGatewaysSatispayGatewayId

> DELETESatispayGatewaysSatispayGatewayId(ctx, satispayGatewayId).Execute()

Delete a satispay gateway



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/incentro-dc/go-commercelayer-sdk/api"
)

func main() {
    satispayGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SatispayGatewaysApi.DELETESatispayGatewaysSatispayGatewayId(context.Background(), satispayGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SatispayGatewaysApi.DELETESatispayGatewaysSatispayGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**satispayGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETESatispayGatewaysSatispayGatewayIdRequest struct via the builder pattern


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


## GETSatispayGateways

> GETSatispayGateways200Response GETSatispayGateways(ctx).Execute()

List all satispay gateways



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/incentro-dc/go-commercelayer-sdk/api"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SatispayGatewaysApi.GETSatispayGateways(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SatispayGatewaysApi.GETSatispayGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETSatispayGateways`: GETSatispayGateways200Response
    fmt.Fprintf(os.Stdout, "Response from `SatispayGatewaysApi.GETSatispayGateways`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETSatispayGatewaysRequest struct via the builder pattern


### Return type

[**GETSatispayGateways200Response**](GETSatispayGateways200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETSatispayGatewaysSatispayGatewayId

> GETSatispayGatewaysSatispayGatewayId200Response GETSatispayGatewaysSatispayGatewayId(ctx, satispayGatewayId).Execute()

Retrieve a satispay gateway



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/incentro-dc/go-commercelayer-sdk/api"
)

func main() {
    satispayGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SatispayGatewaysApi.GETSatispayGatewaysSatispayGatewayId(context.Background(), satispayGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SatispayGatewaysApi.GETSatispayGatewaysSatispayGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETSatispayGatewaysSatispayGatewayId`: GETSatispayGatewaysSatispayGatewayId200Response
    fmt.Fprintf(os.Stdout, "Response from `SatispayGatewaysApi.GETSatispayGatewaysSatispayGatewayId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**satispayGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETSatispayGatewaysSatispayGatewayIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETSatispayGatewaysSatispayGatewayId200Response**](GETSatispayGatewaysSatispayGatewayId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHSatispayGatewaysSatispayGatewayId

> PATCHSatispayGatewaysSatispayGatewayId200Response PATCHSatispayGatewaysSatispayGatewayId(ctx, satispayGatewayId).PATCHSatispayGatewaysSatispayGatewayIdRequest(pATCHSatispayGatewaysSatispayGatewayIdRequest).Execute()

Update a satispay gateway



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/incentro-dc/go-commercelayer-sdk/api"
)

func main() {
    pATCHSatispayGatewaysSatispayGatewayIdRequest := *openapiclient.NewPATCHSatispayGatewaysSatispayGatewayIdRequest(*openapiclient.NewPATCHSatispayGatewaysSatispayGatewayIdRequestData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHManualGatewaysManualGatewayIdRequestDataAttributes())) // PATCHSatispayGatewaysSatispayGatewayIdRequest | 
    satispayGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SatispayGatewaysApi.PATCHSatispayGatewaysSatispayGatewayId(context.Background(), satispayGatewayId).PATCHSatispayGatewaysSatispayGatewayIdRequest(pATCHSatispayGatewaysSatispayGatewayIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SatispayGatewaysApi.PATCHSatispayGatewaysSatispayGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHSatispayGatewaysSatispayGatewayId`: PATCHSatispayGatewaysSatispayGatewayId200Response
    fmt.Fprintf(os.Stdout, "Response from `SatispayGatewaysApi.PATCHSatispayGatewaysSatispayGatewayId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**satispayGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHSatispayGatewaysSatispayGatewayIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pATCHSatispayGatewaysSatispayGatewayIdRequest** | [**PATCHSatispayGatewaysSatispayGatewayIdRequest**](PATCHSatispayGatewaysSatispayGatewayIdRequest.md) |  | 


### Return type

[**PATCHSatispayGatewaysSatispayGatewayId200Response**](PATCHSatispayGatewaysSatispayGatewayId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTSatispayGateways

> POSTSatispayGateways201Response POSTSatispayGateways(ctx).POSTSatispayGatewaysRequest(pOSTSatispayGatewaysRequest).Execute()

Create a satispay gateway



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/incentro-dc/go-commercelayer-sdk/api"
)

func main() {
    pOSTSatispayGatewaysRequest := *openapiclient.NewPOSTSatispayGatewaysRequest(*openapiclient.NewPOSTSatispayGatewaysRequestData(interface{}(123), *openapiclient.NewPOSTManualGatewaysRequestDataAttributes(interface{}(US payment gateway)))) // POSTSatispayGatewaysRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SatispayGatewaysApi.POSTSatispayGateways(context.Background()).POSTSatispayGatewaysRequest(pOSTSatispayGatewaysRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SatispayGatewaysApi.POSTSatispayGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTSatispayGateways`: POSTSatispayGateways201Response
    fmt.Fprintf(os.Stdout, "Response from `SatispayGatewaysApi.POSTSatispayGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTSatispayGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pOSTSatispayGatewaysRequest** | [**POSTSatispayGatewaysRequest**](POSTSatispayGatewaysRequest.md) |  | 

### Return type

[**POSTSatispayGateways201Response**](POSTSatispayGateways201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

