# \PriceVolumeTiersApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEPriceVolumeTiersPriceVolumeTierId**](PriceVolumeTiersApi.md#DELETEPriceVolumeTiersPriceVolumeTierId) | **Delete** /price_volume_tiers/{priceVolumeTierId} | Delete a price volume tier
[**GETPriceIdPriceVolumeTiers**](PriceVolumeTiersApi.md#GETPriceIdPriceVolumeTiers) | **Get** /prices/{priceId}/price_volume_tiers | Retrieve the price volume tiers associated to the price
[**GETPriceVolumeTiers**](PriceVolumeTiersApi.md#GETPriceVolumeTiers) | **Get** /price_volume_tiers | List all price volume tiers
[**GETPriceVolumeTiersPriceVolumeTierId**](PriceVolumeTiersApi.md#GETPriceVolumeTiersPriceVolumeTierId) | **Get** /price_volume_tiers/{priceVolumeTierId} | Retrieve a price volume tier
[**PATCHPriceVolumeTiersPriceVolumeTierId**](PriceVolumeTiersApi.md#PATCHPriceVolumeTiersPriceVolumeTierId) | **Patch** /price_volume_tiers/{priceVolumeTierId} | Update a price volume tier
[**POSTPriceVolumeTiers**](PriceVolumeTiersApi.md#POSTPriceVolumeTiers) | **Post** /price_volume_tiers | Create a price volume tier



## DELETEPriceVolumeTiersPriceVolumeTierId

> DELETEPriceVolumeTiersPriceVolumeTierId(ctx, priceVolumeTierId).Execute()

Delete a price volume tier



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
    priceVolumeTierId := "priceVolumeTierId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriceVolumeTiersApi.DELETEPriceVolumeTiersPriceVolumeTierId(context.Background(), priceVolumeTierId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceVolumeTiersApi.DELETEPriceVolumeTiersPriceVolumeTierId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceVolumeTierId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEPriceVolumeTiersPriceVolumeTierIdRequest struct via the builder pattern


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


## GETPriceIdPriceVolumeTiers

> GETPriceIdPriceVolumeTiers(ctx, priceId).Execute()

Retrieve the price volume tiers associated to the price



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
    priceId := "priceId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriceVolumeTiersApi.GETPriceIdPriceVolumeTiers(context.Background(), priceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceVolumeTiersApi.GETPriceIdPriceVolumeTiers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPriceIdPriceVolumeTiersRequest struct via the builder pattern


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


## GETPriceVolumeTiers

> GETPriceVolumeTiers(ctx).Execute()

List all price volume tiers



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
    resp, r, err := apiClient.PriceVolumeTiersApi.GETPriceVolumeTiers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceVolumeTiersApi.GETPriceVolumeTiers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETPriceVolumeTiersRequest struct via the builder pattern


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


## GETPriceVolumeTiersPriceVolumeTierId

> PriceVolumeTier GETPriceVolumeTiersPriceVolumeTierId(ctx, priceVolumeTierId).Execute()

Retrieve a price volume tier



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
    priceVolumeTierId := "priceVolumeTierId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriceVolumeTiersApi.GETPriceVolumeTiersPriceVolumeTierId(context.Background(), priceVolumeTierId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceVolumeTiersApi.GETPriceVolumeTiersPriceVolumeTierId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPriceVolumeTiersPriceVolumeTierId`: PriceVolumeTier
    fmt.Fprintf(os.Stdout, "Response from `PriceVolumeTiersApi.GETPriceVolumeTiersPriceVolumeTierId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceVolumeTierId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPriceVolumeTiersPriceVolumeTierIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PriceVolumeTier**](PriceVolumeTier.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHPriceVolumeTiersPriceVolumeTierId

> PATCHPriceVolumeTiersPriceVolumeTierId(ctx, priceVolumeTierId).PriceVolumeTierUpdate(priceVolumeTierUpdate).Execute()

Update a price volume tier



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
    priceVolumeTierUpdate := *openapiclient.NewPriceVolumeTierUpdate(*openapiclient.NewPriceVolumeTierUpdateData("price_volume_tiers", "XGZwpOSrWL", *openapiclient.NewPriceVolumeTierUpdateDataAttributes())) // PriceVolumeTierUpdate | 
    priceVolumeTierId := "priceVolumeTierId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriceVolumeTiersApi.PATCHPriceVolumeTiersPriceVolumeTierId(context.Background(), priceVolumeTierId).PriceVolumeTierUpdate(priceVolumeTierUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceVolumeTiersApi.PATCHPriceVolumeTiersPriceVolumeTierId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceVolumeTierId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHPriceVolumeTiersPriceVolumeTierIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **priceVolumeTierUpdate** | [**PriceVolumeTierUpdate**](PriceVolumeTierUpdate.md) |  | 


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


## POSTPriceVolumeTiers

> POSTPriceVolumeTiers(ctx).PriceVolumeTierCreate(priceVolumeTierCreate).Execute()

Create a price volume tier



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
    priceVolumeTierCreate := *openapiclient.NewPriceVolumeTierCreate(*openapiclient.NewPriceVolumeTierCreateData("price_volume_tiers", *openapiclient.NewPriceVolumeTierCreateDataAttributes("six pack", int32(1000)))) // PriceVolumeTierCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriceVolumeTiersApi.POSTPriceVolumeTiers(context.Background()).PriceVolumeTierCreate(priceVolumeTierCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceVolumeTiersApi.POSTPriceVolumeTiers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTPriceVolumeTiersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **priceVolumeTierCreate** | [**PriceVolumeTierCreate**](PriceVolumeTierCreate.md) |  | 

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

