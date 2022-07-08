# \InventoryReturnLocationsApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEInventoryReturnLocationsInventoryReturnLocationId**](InventoryReturnLocationsApi.md#DELETEInventoryReturnLocationsInventoryReturnLocationId) | **Delete** /inventory_return_locations/{inventoryReturnLocationId} | Delete an inventory return location
[**GETInventoryModelIdInventoryReturnLocations**](InventoryReturnLocationsApi.md#GETInventoryModelIdInventoryReturnLocations) | **Get** /inventory_models/{inventoryModelId}/inventory_return_locations | Retrieve the inventory return locations associated to the inventory model
[**GETInventoryReturnLocations**](InventoryReturnLocationsApi.md#GETInventoryReturnLocations) | **Get** /inventory_return_locations | List all inventory return locations
[**GETInventoryReturnLocationsInventoryReturnLocationId**](InventoryReturnLocationsApi.md#GETInventoryReturnLocationsInventoryReturnLocationId) | **Get** /inventory_return_locations/{inventoryReturnLocationId} | Retrieve an inventory return location
[**GETStockLocationIdInventoryReturnLocations**](InventoryReturnLocationsApi.md#GETStockLocationIdInventoryReturnLocations) | **Get** /stock_locations/{stockLocationId}/inventory_return_locations | Retrieve the inventory return locations associated to the stock location
[**PATCHInventoryReturnLocationsInventoryReturnLocationId**](InventoryReturnLocationsApi.md#PATCHInventoryReturnLocationsInventoryReturnLocationId) | **Patch** /inventory_return_locations/{inventoryReturnLocationId} | Update an inventory return location
[**POSTInventoryReturnLocations**](InventoryReturnLocationsApi.md#POSTInventoryReturnLocations) | **Post** /inventory_return_locations | Create an inventory return location



## DELETEInventoryReturnLocationsInventoryReturnLocationId

> DELETEInventoryReturnLocationsInventoryReturnLocationId(ctx, inventoryReturnLocationId).Execute()

Delete an inventory return location



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
    inventoryReturnLocationId := "inventoryReturnLocationId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryReturnLocationsApi.DELETEInventoryReturnLocationsInventoryReturnLocationId(context.Background(), inventoryReturnLocationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryReturnLocationsApi.DELETEInventoryReturnLocationsInventoryReturnLocationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventoryReturnLocationId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEInventoryReturnLocationsInventoryReturnLocationIdRequest struct via the builder pattern


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


## GETInventoryModelIdInventoryReturnLocations

> GETInventoryModelIdInventoryReturnLocations(ctx, inventoryModelId).Execute()

Retrieve the inventory return locations associated to the inventory model



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
    resp, r, err := apiClient.InventoryReturnLocationsApi.GETInventoryModelIdInventoryReturnLocations(context.Background(), inventoryModelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryReturnLocationsApi.GETInventoryModelIdInventoryReturnLocations``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETInventoryModelIdInventoryReturnLocationsRequest struct via the builder pattern


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


## GETInventoryReturnLocations

> GETInventoryReturnLocations(ctx).Execute()

List all inventory return locations



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
    resp, r, err := apiClient.InventoryReturnLocationsApi.GETInventoryReturnLocations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryReturnLocationsApi.GETInventoryReturnLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETInventoryReturnLocationsRequest struct via the builder pattern


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


## GETInventoryReturnLocationsInventoryReturnLocationId

> InventoryReturnLocation GETInventoryReturnLocationsInventoryReturnLocationId(ctx, inventoryReturnLocationId).Execute()

Retrieve an inventory return location



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
    inventoryReturnLocationId := "inventoryReturnLocationId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryReturnLocationsApi.GETInventoryReturnLocationsInventoryReturnLocationId(context.Background(), inventoryReturnLocationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryReturnLocationsApi.GETInventoryReturnLocationsInventoryReturnLocationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETInventoryReturnLocationsInventoryReturnLocationId`: InventoryReturnLocation
    fmt.Fprintf(os.Stdout, "Response from `InventoryReturnLocationsApi.GETInventoryReturnLocationsInventoryReturnLocationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventoryReturnLocationId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETInventoryReturnLocationsInventoryReturnLocationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InventoryReturnLocation**](InventoryReturnLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETStockLocationIdInventoryReturnLocations

> GETStockLocationIdInventoryReturnLocations(ctx, stockLocationId).Execute()

Retrieve the inventory return locations associated to the stock location



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
    resp, r, err := apiClient.InventoryReturnLocationsApi.GETStockLocationIdInventoryReturnLocations(context.Background(), stockLocationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryReturnLocationsApi.GETStockLocationIdInventoryReturnLocations``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETStockLocationIdInventoryReturnLocationsRequest struct via the builder pattern


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


## PATCHInventoryReturnLocationsInventoryReturnLocationId

> PATCHInventoryReturnLocationsInventoryReturnLocationId(ctx, inventoryReturnLocationId).InventoryReturnLocationUpdate(inventoryReturnLocationUpdate).Execute()

Update an inventory return location



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
    inventoryReturnLocationUpdate := *openapiclient.NewInventoryReturnLocationUpdate(*openapiclient.NewInventoryReturnLocationUpdateData("inventory_return_locations", "XGZwpOSrWL", *openapiclient.NewInventoryReturnLocationUpdateDataAttributes())) // InventoryReturnLocationUpdate | 
    inventoryReturnLocationId := "inventoryReturnLocationId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryReturnLocationsApi.PATCHInventoryReturnLocationsInventoryReturnLocationId(context.Background(), inventoryReturnLocationId).InventoryReturnLocationUpdate(inventoryReturnLocationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryReturnLocationsApi.PATCHInventoryReturnLocationsInventoryReturnLocationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventoryReturnLocationId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHInventoryReturnLocationsInventoryReturnLocationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inventoryReturnLocationUpdate** | [**InventoryReturnLocationUpdate**](InventoryReturnLocationUpdate.md) |  | 


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


## POSTInventoryReturnLocations

> POSTInventoryReturnLocations(ctx).InventoryReturnLocationCreate(inventoryReturnLocationCreate).Execute()

Create an inventory return location



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
    inventoryReturnLocationCreate := *openapiclient.NewInventoryReturnLocationCreate(*openapiclient.NewInventoryReturnLocationCreateData("inventory_return_locations", *openapiclient.NewInventoryReturnLocationCreateDataAttributes(int32(1)))) // InventoryReturnLocationCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryReturnLocationsApi.POSTInventoryReturnLocations(context.Background()).InventoryReturnLocationCreate(inventoryReturnLocationCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryReturnLocationsApi.POSTInventoryReturnLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTInventoryReturnLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inventoryReturnLocationCreate** | [**InventoryReturnLocationCreate**](InventoryReturnLocationCreate.md) |  | 

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

