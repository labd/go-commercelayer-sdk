# \StockItemsApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEStockItemsStockItemId**](StockItemsApi.md#DELETEStockItemsStockItemId) | **Delete** /stock_items/{stockItemId} | Delete a stock item
[**GETSkuIdStockItems**](StockItemsApi.md#GETSkuIdStockItems) | **Get** /skus/{skuId}/stock_items | Retrieve the stock items associated to the SKU
[**GETStockItems**](StockItemsApi.md#GETStockItems) | **Get** /stock_items | List all stock items
[**GETStockItemsStockItemId**](StockItemsApi.md#GETStockItemsStockItemId) | **Get** /stock_items/{stockItemId} | Retrieve a stock item
[**GETStockLineItemIdStockItem**](StockItemsApi.md#GETStockLineItemIdStockItem) | **Get** /stock_line_items/{stockLineItemId}/stock_item | Retrieve the stock item associated to the stock line item
[**GETStockLocationIdStockItems**](StockItemsApi.md#GETStockLocationIdStockItems) | **Get** /stock_locations/{stockLocationId}/stock_items | Retrieve the stock items associated to the stock location
[**PATCHStockItemsStockItemId**](StockItemsApi.md#PATCHStockItemsStockItemId) | **Patch** /stock_items/{stockItemId} | Update a stock item
[**POSTStockItems**](StockItemsApi.md#POSTStockItems) | **Post** /stock_items | Create a stock item



## DELETEStockItemsStockItemId

> DELETEStockItemsStockItemId(ctx, stockItemId).Execute()

Delete a stock item



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
    stockItemId := "stockItemId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StockItemsApi.DELETEStockItemsStockItemId(context.Background(), stockItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockItemsApi.DELETEStockItemsStockItemId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stockItemId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEStockItemsStockItemIdRequest struct via the builder pattern


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


## GETSkuIdStockItems

> GETSkuIdStockItems(ctx, skuId).Execute()

Retrieve the stock items associated to the SKU



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
    skuId := "skuId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StockItemsApi.GETSkuIdStockItems(context.Background(), skuId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockItemsApi.GETSkuIdStockItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skuId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETSkuIdStockItemsRequest struct via the builder pattern


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


## GETStockItems

> GETStockItems(ctx).Execute()

List all stock items



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
    resp, r, err := apiClient.StockItemsApi.GETStockItems(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockItemsApi.GETStockItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETStockItemsRequest struct via the builder pattern


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


## GETStockItemsStockItemId

> StockItem GETStockItemsStockItemId(ctx, stockItemId).Execute()

Retrieve a stock item



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
    stockItemId := "stockItemId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StockItemsApi.GETStockItemsStockItemId(context.Background(), stockItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockItemsApi.GETStockItemsStockItemId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETStockItemsStockItemId`: StockItem
    fmt.Fprintf(os.Stdout, "Response from `StockItemsApi.GETStockItemsStockItemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stockItemId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETStockItemsStockItemIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StockItem**](StockItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETStockLineItemIdStockItem

> GETStockLineItemIdStockItem(ctx, stockLineItemId).Execute()

Retrieve the stock item associated to the stock line item



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
    stockLineItemId := "stockLineItemId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StockItemsApi.GETStockLineItemIdStockItem(context.Background(), stockLineItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockItemsApi.GETStockLineItemIdStockItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stockLineItemId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETStockLineItemIdStockItemRequest struct via the builder pattern


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


## GETStockLocationIdStockItems

> GETStockLocationIdStockItems(ctx, stockLocationId).Execute()

Retrieve the stock items associated to the stock location



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
    stockLocationId := "stockLocationId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StockItemsApi.GETStockLocationIdStockItems(context.Background(), stockLocationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockItemsApi.GETStockLocationIdStockItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stockLocationId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETStockLocationIdStockItemsRequest struct via the builder pattern


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


## PATCHStockItemsStockItemId

> PATCHStockItemsStockItemId(ctx, stockItemId).StockItemUpdate(stockItemUpdate).Execute()

Update a stock item



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
    stockItemUpdate := *openapiclient.NewStockItemUpdate(*openapiclient.NewStockItemUpdateData("stock_items", "XGZwpOSrWL", *openapiclient.NewStockItemUpdateDataAttributes())) // StockItemUpdate | 
    stockItemId := "stockItemId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StockItemsApi.PATCHStockItemsStockItemId(context.Background(), stockItemId).StockItemUpdate(stockItemUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockItemsApi.PATCHStockItemsStockItemId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stockItemId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHStockItemsStockItemIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stockItemUpdate** | [**StockItemUpdate**](StockItemUpdate.md) |  | 


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


## POSTStockItems

> POSTStockItems(ctx).StockItemCreate(stockItemCreate).Execute()

Create a stock item



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
    stockItemCreate := *openapiclient.NewStockItemCreate(*openapiclient.NewStockItemCreateData("stock_items", *openapiclient.NewStockItemCreateDataAttributes(int32(100)))) // StockItemCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StockItemsApi.POSTStockItems(context.Background()).StockItemCreate(stockItemCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockItemsApi.POSTStockItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTStockItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stockItemCreate** | [**StockItemCreate**](StockItemCreate.md) |  | 

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

