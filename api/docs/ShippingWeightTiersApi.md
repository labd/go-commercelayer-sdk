# \ShippingWeightTiersApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEShippingWeightTiersShippingWeightTierId**](ShippingWeightTiersApi.md#DELETEShippingWeightTiersShippingWeightTierId) | **Delete** /shipping_weight_tiers/{shippingWeightTierId} | Delete a shipping weight tier
[**GETShippingMethodIdShippingWeightTiers**](ShippingWeightTiersApi.md#GETShippingMethodIdShippingWeightTiers) | **Get** /shipping_methods/{shippingMethodId}/shipping_weight_tiers | Retrieve the shipping weight tiers associated to the shipping method
[**GETShippingWeightTiers**](ShippingWeightTiersApi.md#GETShippingWeightTiers) | **Get** /shipping_weight_tiers | List all shipping weight tiers
[**GETShippingWeightTiersShippingWeightTierId**](ShippingWeightTiersApi.md#GETShippingWeightTiersShippingWeightTierId) | **Get** /shipping_weight_tiers/{shippingWeightTierId} | Retrieve a shipping weight tier
[**PATCHShippingWeightTiersShippingWeightTierId**](ShippingWeightTiersApi.md#PATCHShippingWeightTiersShippingWeightTierId) | **Patch** /shipping_weight_tiers/{shippingWeightTierId} | Update a shipping weight tier
[**POSTShippingWeightTiers**](ShippingWeightTiersApi.md#POSTShippingWeightTiers) | **Post** /shipping_weight_tiers | Create a shipping weight tier



## DELETEShippingWeightTiersShippingWeightTierId

> DELETEShippingWeightTiersShippingWeightTierId(ctx, shippingWeightTierId).Execute()

Delete a shipping weight tier



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
    shippingWeightTierId := "shippingWeightTierId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingWeightTiersApi.DELETEShippingWeightTiersShippingWeightTierId(context.Background(), shippingWeightTierId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingWeightTiersApi.DELETEShippingWeightTiersShippingWeightTierId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shippingWeightTierId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEShippingWeightTiersShippingWeightTierIdRequest struct via the builder pattern


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


## GETShippingMethodIdShippingWeightTiers

> GETShippingMethodIdShippingWeightTiers(ctx, shippingMethodId).Execute()

Retrieve the shipping weight tiers associated to the shipping method



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
    resp, r, err := apiClient.ShippingWeightTiersApi.GETShippingMethodIdShippingWeightTiers(context.Background(), shippingMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingWeightTiersApi.GETShippingMethodIdShippingWeightTiers``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETShippingMethodIdShippingWeightTiersRequest struct via the builder pattern


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


## GETShippingWeightTiers

> GETShippingWeightTiers(ctx).Execute()

List all shipping weight tiers



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
    resp, r, err := apiClient.ShippingWeightTiersApi.GETShippingWeightTiers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingWeightTiersApi.GETShippingWeightTiers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETShippingWeightTiersRequest struct via the builder pattern


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


## GETShippingWeightTiersShippingWeightTierId

> ShippingWeightTier GETShippingWeightTiersShippingWeightTierId(ctx, shippingWeightTierId).Execute()

Retrieve a shipping weight tier



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
    shippingWeightTierId := "shippingWeightTierId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingWeightTiersApi.GETShippingWeightTiersShippingWeightTierId(context.Background(), shippingWeightTierId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingWeightTiersApi.GETShippingWeightTiersShippingWeightTierId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETShippingWeightTiersShippingWeightTierId`: ShippingWeightTier
    fmt.Fprintf(os.Stdout, "Response from `ShippingWeightTiersApi.GETShippingWeightTiersShippingWeightTierId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shippingWeightTierId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETShippingWeightTiersShippingWeightTierIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ShippingWeightTier**](ShippingWeightTier.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHShippingWeightTiersShippingWeightTierId

> PATCHShippingWeightTiersShippingWeightTierId(ctx, shippingWeightTierId).ShippingWeightTierUpdate(shippingWeightTierUpdate).Execute()

Update a shipping weight tier



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
    shippingWeightTierId := "shippingWeightTierId_example" // string | The resource's id
    shippingWeightTierUpdate := *openapiclient.NewShippingWeightTierUpdate(*openapiclient.NewShippingWeightTierUpdateData("shipping_weight_tiers", "XGZwpOSrWL", *openapiclient.NewShippingWeightTierUpdateDataAttributes())) // ShippingWeightTierUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingWeightTiersApi.PATCHShippingWeightTiersShippingWeightTierId(context.Background(), shippingWeightTierId).ShippingWeightTierUpdate(shippingWeightTierUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingWeightTiersApi.PATCHShippingWeightTiersShippingWeightTierId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shippingWeightTierId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHShippingWeightTiersShippingWeightTierIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shippingWeightTierUpdate** | [**ShippingWeightTierUpdate**](ShippingWeightTierUpdate.md) |  | 

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


## POSTShippingWeightTiers

> POSTShippingWeightTiers(ctx).ShippingWeightTierCreate(shippingWeightTierCreate).Execute()

Create a shipping weight tier



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
    shippingWeightTierCreate := *openapiclient.NewShippingWeightTierCreate(*openapiclient.NewShippingWeightTierCreateData("shipping_weight_tiers", *openapiclient.NewShippingWeightTierCreateDataAttributes("Light shipping under 3kg", int32(1000)))) // ShippingWeightTierCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingWeightTiersApi.POSTShippingWeightTiers(context.Background()).ShippingWeightTierCreate(shippingWeightTierCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingWeightTiersApi.POSTShippingWeightTiers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTShippingWeightTiersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shippingWeightTierCreate** | [**ShippingWeightTierCreate**](ShippingWeightTierCreate.md) |  | 

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

