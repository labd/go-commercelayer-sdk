# \PriceListsApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEPriceListsPriceListId**](PriceListsApi.md#DELETEPriceListsPriceListId) | **Delete** /price_lists/{priceListId} | Delete a price list
[**GETMarketIdPriceList**](PriceListsApi.md#GETMarketIdPriceList) | **Get** /markets/{marketId}/price_list | Retrieve the price list associated to the market
[**GETPriceIdPriceList**](PriceListsApi.md#GETPriceIdPriceList) | **Get** /prices/{priceId}/price_list | Retrieve the price list associated to the price
[**GETPriceLists**](PriceListsApi.md#GETPriceLists) | **Get** /price_lists | List all price lists
[**GETPriceListsPriceListId**](PriceListsApi.md#GETPriceListsPriceListId) | **Get** /price_lists/{priceListId} | Retrieve a price list
[**PATCHPriceListsPriceListId**](PriceListsApi.md#PATCHPriceListsPriceListId) | **Patch** /price_lists/{priceListId} | Update a price list
[**POSTPriceLists**](PriceListsApi.md#POSTPriceLists) | **Post** /price_lists | Create a price list



## DELETEPriceListsPriceListId

> DELETEPriceListsPriceListId(ctx, priceListId).Execute()

Delete a price list



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
    priceListId := "priceListId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriceListsApi.DELETEPriceListsPriceListId(context.Background(), priceListId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceListsApi.DELETEPriceListsPriceListId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceListId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEPriceListsPriceListIdRequest struct via the builder pattern


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


## GETMarketIdPriceList

> GETMarketIdPriceList(ctx, marketId).Execute()

Retrieve the price list associated to the market



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
    marketId := "marketId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriceListsApi.GETMarketIdPriceList(context.Background(), marketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceListsApi.GETMarketIdPriceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETMarketIdPriceListRequest struct via the builder pattern


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


## GETPriceIdPriceList

> GETPriceIdPriceList(ctx, priceId).Execute()

Retrieve the price list associated to the price



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
    resp, r, err := apiClient.PriceListsApi.GETPriceIdPriceList(context.Background(), priceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceListsApi.GETPriceIdPriceList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETPriceIdPriceListRequest struct via the builder pattern


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


## GETPriceLists

> GETPriceLists(ctx).Execute()

List all price lists



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
    resp, r, err := apiClient.PriceListsApi.GETPriceLists(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceListsApi.GETPriceLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETPriceListsRequest struct via the builder pattern


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


## GETPriceListsPriceListId

> PriceList GETPriceListsPriceListId(ctx, priceListId).Execute()

Retrieve a price list



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
    priceListId := "priceListId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriceListsApi.GETPriceListsPriceListId(context.Background(), priceListId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceListsApi.GETPriceListsPriceListId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPriceListsPriceListId`: PriceList
    fmt.Fprintf(os.Stdout, "Response from `PriceListsApi.GETPriceListsPriceListId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceListId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPriceListsPriceListIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PriceList**](PriceList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHPriceListsPriceListId

> PATCHPriceListsPriceListId(ctx, priceListId).PriceListUpdate(priceListUpdate).Execute()

Update a price list



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
    priceListUpdate := *openapiclient.NewPriceListUpdate(*openapiclient.NewPriceListUpdateData("price_lists", "XGZwpOSrWL", *openapiclient.NewPriceListUpdateDataAttributes())) // PriceListUpdate | 
    priceListId := "priceListId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriceListsApi.PATCHPriceListsPriceListId(context.Background(), priceListId).PriceListUpdate(priceListUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceListsApi.PATCHPriceListsPriceListId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceListId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHPriceListsPriceListIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **priceListUpdate** | [**PriceListUpdate**](PriceListUpdate.md) |  | 


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


## POSTPriceLists

> POSTPriceLists(ctx).PriceListCreate(priceListCreate).Execute()

Create a price list



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
    priceListCreate := *openapiclient.NewPriceListCreate(*openapiclient.NewPriceListCreateData("price_lists", *openapiclient.NewPriceListCreateDataAttributes("EU Price list", "EUR"))) // PriceListCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriceListsApi.POSTPriceLists(context.Background()).PriceListCreate(priceListCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceListsApi.POSTPriceLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTPriceListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **priceListCreate** | [**PriceListCreate**](PriceListCreate.md) |  | 

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

