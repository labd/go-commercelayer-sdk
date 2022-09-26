# \ShippingZonesApi

All URIs are relative to *https://}.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEShippingZonesShippingZoneId**](ShippingZonesApi.md#DELETEShippingZonesShippingZoneId) | **Delete** /shipping_zones/{shippingZoneId} | Delete a shipping zone
[**GETShippingMethodIdShippingZone**](ShippingZonesApi.md#GETShippingMethodIdShippingZone) | **Get** /shipping_methods/{shippingMethodId}/shipping_zone | Retrieve the shipping zone associated to the shipping method
[**GETShippingZones**](ShippingZonesApi.md#GETShippingZones) | **Get** /shipping_zones | List all shipping zones
[**GETShippingZonesShippingZoneId**](ShippingZonesApi.md#GETShippingZonesShippingZoneId) | **Get** /shipping_zones/{shippingZoneId} | Retrieve a shipping zone
[**PATCHShippingZonesShippingZoneId**](ShippingZonesApi.md#PATCHShippingZonesShippingZoneId) | **Patch** /shipping_zones/{shippingZoneId} | Update a shipping zone
[**POSTShippingZones**](ShippingZonesApi.md#POSTShippingZones) | **Post** /shipping_zones | Create a shipping zone



## DELETEShippingZonesShippingZoneId

> DELETEShippingZonesShippingZoneId(ctx, shippingZoneId).Execute()

Delete a shipping zone



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
    shippingZoneId := "shippingZoneId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingZonesApi.DELETEShippingZonesShippingZoneId(context.Background(), shippingZoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingZonesApi.DELETEShippingZonesShippingZoneId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shippingZoneId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEShippingZonesShippingZoneIdRequest struct via the builder pattern


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


## GETShippingMethodIdShippingZone

> GETShippingMethodIdShippingZone(ctx, shippingMethodId).Execute()

Retrieve the shipping zone associated to the shipping method



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
    shippingMethodId := "shippingMethodId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingZonesApi.GETShippingMethodIdShippingZone(context.Background(), shippingMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingZonesApi.GETShippingMethodIdShippingZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shippingMethodId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETShippingMethodIdShippingZoneRequest struct via the builder pattern


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


## GETShippingZones

> ShippingZoneResponseList GETShippingZones(ctx).Execute()

List all shipping zones



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
    resp, r, err := apiClient.ShippingZonesApi.GETShippingZones(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingZonesApi.GETShippingZones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETShippingZones`: ShippingZoneResponseList
    fmt.Fprintf(os.Stdout, "Response from `ShippingZonesApi.GETShippingZones`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETShippingZonesRequest struct via the builder pattern


### Return type

[**ShippingZoneResponseList**](ShippingZoneResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETShippingZonesShippingZoneId

> ShippingZoneResponse GETShippingZonesShippingZoneId(ctx, shippingZoneId).Execute()

Retrieve a shipping zone



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
    shippingZoneId := "shippingZoneId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingZonesApi.GETShippingZonesShippingZoneId(context.Background(), shippingZoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingZonesApi.GETShippingZonesShippingZoneId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETShippingZonesShippingZoneId`: ShippingZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `ShippingZonesApi.GETShippingZonesShippingZoneId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shippingZoneId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETShippingZonesShippingZoneIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ShippingZoneResponse**](ShippingZoneResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHShippingZonesShippingZoneId

> ShippingZoneResponse PATCHShippingZonesShippingZoneId(ctx, shippingZoneId).ShippingZoneUpdate(shippingZoneUpdate).Execute()

Update a shipping zone



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
    shippingZoneUpdate := *openapiclient.NewShippingZoneUpdate(*openapiclient.NewShippingZoneUpdateData("Type_example", "XGZwpOSrWL", *openapiclient.NewShippingZoneUpdateDataAttributes())) // ShippingZoneUpdate | 
    shippingZoneId := "shippingZoneId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingZonesApi.PATCHShippingZonesShippingZoneId(context.Background(), shippingZoneId).ShippingZoneUpdate(shippingZoneUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingZonesApi.PATCHShippingZonesShippingZoneId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHShippingZonesShippingZoneId`: ShippingZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `ShippingZonesApi.PATCHShippingZonesShippingZoneId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shippingZoneId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHShippingZonesShippingZoneIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shippingZoneUpdate** | [**ShippingZoneUpdate**](ShippingZoneUpdate.md) |  | 


### Return type

[**ShippingZoneResponse**](ShippingZoneResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTShippingZones

> ShippingZoneResponse POSTShippingZones(ctx).ShippingZoneCreate(shippingZoneCreate).Execute()

Create a shipping zone



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
    shippingZoneCreate := *openapiclient.NewShippingZoneCreate(*openapiclient.NewShippingZoneCreateData("Type_example", *openapiclient.NewShippingZoneCreateDataAttributes("Europe (main countries)"))) // ShippingZoneCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingZonesApi.POSTShippingZones(context.Background()).ShippingZoneCreate(shippingZoneCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingZonesApi.POSTShippingZones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTShippingZones`: ShippingZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `ShippingZonesApi.POSTShippingZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTShippingZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shippingZoneCreate** | [**ShippingZoneCreate**](ShippingZoneCreate.md) |  | 

### Return type

[**ShippingZoneResponse**](ShippingZoneResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

