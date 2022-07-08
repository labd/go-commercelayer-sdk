# \InventoryStockLocationsApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEInventoryStockLocationsInventoryStockLocationId**](InventoryStockLocationsApi.md#DELETEInventoryStockLocationsInventoryStockLocationId) | **Delete** /inventory_stock_locations/{inventoryStockLocationId} | Delete an inventory stock location
[**GETInventoryModelIdInventoryStockLocations**](InventoryStockLocationsApi.md#GETInventoryModelIdInventoryStockLocations) | **Get** /inventory_models/{inventoryModelId}/inventory_stock_locations | Retrieve the inventory stock locations associated to the inventory model
[**GETInventoryStockLocations**](InventoryStockLocationsApi.md#GETInventoryStockLocations) | **Get** /inventory_stock_locations | List all inventory stock locations
[**GETInventoryStockLocationsInventoryStockLocationId**](InventoryStockLocationsApi.md#GETInventoryStockLocationsInventoryStockLocationId) | **Get** /inventory_stock_locations/{inventoryStockLocationId} | Retrieve an inventory stock location
[**GETStockLocationIdInventoryStockLocations**](InventoryStockLocationsApi.md#GETStockLocationIdInventoryStockLocations) | **Get** /stock_locations/{stockLocationId}/inventory_stock_locations | Retrieve the inventory stock locations associated to the stock location
[**PATCHInventoryStockLocationsInventoryStockLocationId**](InventoryStockLocationsApi.md#PATCHInventoryStockLocationsInventoryStockLocationId) | **Patch** /inventory_stock_locations/{inventoryStockLocationId} | Update an inventory stock location
[**POSTInventoryStockLocations**](InventoryStockLocationsApi.md#POSTInventoryStockLocations) | **Post** /inventory_stock_locations | Create an inventory stock location



## DELETEInventoryStockLocationsInventoryStockLocationId

> DELETEInventoryStockLocationsInventoryStockLocationId(ctx, inventoryStockLocationId).Execute()

Delete an inventory stock location



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
    inventoryStockLocationId := "inventoryStockLocationId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryStockLocationsApi.DELETEInventoryStockLocationsInventoryStockLocationId(context.Background(), inventoryStockLocationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryStockLocationsApi.DELETEInventoryStockLocationsInventoryStockLocationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventoryStockLocationId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEInventoryStockLocationsInventoryStockLocationIdRequest struct via the builder pattern


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


## GETInventoryModelIdInventoryStockLocations

> GETInventoryModelIdInventoryStockLocations(ctx, inventoryModelId).Execute()

Retrieve the inventory stock locations associated to the inventory model



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
    inventoryModelId := "inventoryModelId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryStockLocationsApi.GETInventoryModelIdInventoryStockLocations(context.Background(), inventoryModelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryStockLocationsApi.GETInventoryModelIdInventoryStockLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventoryModelId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETInventoryModelIdInventoryStockLocationsRequest struct via the builder pattern


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


## GETInventoryStockLocations

> GETInventoryStockLocations(ctx).Execute()

List all inventory stock locations



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
    resp, r, err := apiClient.InventoryStockLocationsApi.GETInventoryStockLocations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryStockLocationsApi.GETInventoryStockLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETInventoryStockLocationsRequest struct via the builder pattern


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


## GETInventoryStockLocationsInventoryStockLocationId

> InventoryStockLocation GETInventoryStockLocationsInventoryStockLocationId(ctx, inventoryStockLocationId).Execute()

Retrieve an inventory stock location



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
    inventoryStockLocationId := "inventoryStockLocationId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryStockLocationsApi.GETInventoryStockLocationsInventoryStockLocationId(context.Background(), inventoryStockLocationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryStockLocationsApi.GETInventoryStockLocationsInventoryStockLocationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETInventoryStockLocationsInventoryStockLocationId`: InventoryStockLocation
    fmt.Fprintf(os.Stdout, "Response from `InventoryStockLocationsApi.GETInventoryStockLocationsInventoryStockLocationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventoryStockLocationId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETInventoryStockLocationsInventoryStockLocationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InventoryStockLocation**](InventoryStockLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETStockLocationIdInventoryStockLocations

> GETStockLocationIdInventoryStockLocations(ctx, stockLocationId).Execute()

Retrieve the inventory stock locations associated to the stock location



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
    resp, r, err := apiClient.InventoryStockLocationsApi.GETStockLocationIdInventoryStockLocations(context.Background(), stockLocationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryStockLocationsApi.GETStockLocationIdInventoryStockLocations``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETStockLocationIdInventoryStockLocationsRequest struct via the builder pattern


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


## PATCHInventoryStockLocationsInventoryStockLocationId

> PATCHInventoryStockLocationsInventoryStockLocationId(ctx, inventoryStockLocationId).InventoryStockLocationUpdate(inventoryStockLocationUpdate).Execute()

Update an inventory stock location



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
    inventoryStockLocationUpdate := *openapiclient.NewInventoryStockLocationUpdate(*openapiclient.NewInventoryStockLocationUpdateData("inventory_stock_locations", "XGZwpOSrWL", *openapiclient.NewInventoryStockLocationUpdateDataAttributes())) // InventoryStockLocationUpdate | 
    inventoryStockLocationId := "inventoryStockLocationId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryStockLocationsApi.PATCHInventoryStockLocationsInventoryStockLocationId(context.Background(), inventoryStockLocationId).InventoryStockLocationUpdate(inventoryStockLocationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryStockLocationsApi.PATCHInventoryStockLocationsInventoryStockLocationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventoryStockLocationId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHInventoryStockLocationsInventoryStockLocationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inventoryStockLocationUpdate** | [**InventoryStockLocationUpdate**](InventoryStockLocationUpdate.md) |  | 


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


## POSTInventoryStockLocations

> POSTInventoryStockLocations(ctx).InventoryStockLocationCreate(inventoryStockLocationCreate).Execute()

Create an inventory stock location



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
    inventoryStockLocationCreate := *openapiclient.NewInventoryStockLocationCreate(*openapiclient.NewInventoryStockLocationCreateData("inventory_stock_locations", *openapiclient.NewInventoryStockLocationCreateDataAttributes(int32(1)))) // InventoryStockLocationCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryStockLocationsApi.POSTInventoryStockLocations(context.Background()).InventoryStockLocationCreate(inventoryStockLocationCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryStockLocationsApi.POSTInventoryStockLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTInventoryStockLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inventoryStockLocationCreate** | [**InventoryStockLocationCreate**](InventoryStockLocationCreate.md) |  | 

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

