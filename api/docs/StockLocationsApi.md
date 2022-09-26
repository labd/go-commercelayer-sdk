# \StockLocationsApi

All URIs are relative to *https://}.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEStockLocationsStockLocationId**](StockLocationsApi.md#DELETEStockLocationsStockLocationId) | **Delete** /stock_locations/{stockLocationId} | Delete a stock location
[**GETDeliveryLeadTimeIdStockLocation**](StockLocationsApi.md#GETDeliveryLeadTimeIdStockLocation) | **Get** /delivery_lead_times/{deliveryLeadTimeId}/stock_location | Retrieve the stock location associated to the delivery lead time
[**GETInventoryReturnLocationIdStockLocation**](StockLocationsApi.md#GETInventoryReturnLocationIdStockLocation) | **Get** /inventory_return_locations/{inventoryReturnLocationId}/stock_location | Retrieve the stock location associated to the inventory return location
[**GETInventoryStockLocationIdStockLocation**](StockLocationsApi.md#GETInventoryStockLocationIdStockLocation) | **Get** /inventory_stock_locations/{inventoryStockLocationId}/stock_location | Retrieve the stock location associated to the inventory stock location
[**GETPackageIdStockLocation**](StockLocationsApi.md#GETPackageIdStockLocation) | **Get** /packages/{packageId}/stock_location | Retrieve the stock location associated to the package
[**GETReturnIdStockLocation**](StockLocationsApi.md#GETReturnIdStockLocation) | **Get** /returns/{returnId}/stock_location | Retrieve the stock location associated to the return
[**GETShipmentIdStockLocation**](StockLocationsApi.md#GETShipmentIdStockLocation) | **Get** /shipments/{shipmentId}/stock_location | Retrieve the stock location associated to the shipment
[**GETShippingMethodIdStockLocation**](StockLocationsApi.md#GETShippingMethodIdStockLocation) | **Get** /shipping_methods/{shippingMethodId}/stock_location | Retrieve the stock location associated to the shipping method
[**GETStockItemIdStockLocation**](StockLocationsApi.md#GETStockItemIdStockLocation) | **Get** /stock_items/{stockItemId}/stock_location | Retrieve the stock location associated to the stock item
[**GETStockLocations**](StockLocationsApi.md#GETStockLocations) | **Get** /stock_locations | List all stock locations
[**GETStockLocationsStockLocationId**](StockLocationsApi.md#GETStockLocationsStockLocationId) | **Get** /stock_locations/{stockLocationId} | Retrieve a stock location
[**GETStockTransferIdDestinationStockLocation**](StockLocationsApi.md#GETStockTransferIdDestinationStockLocation) | **Get** /stock_transfers/{stockTransferId}/destination_stock_location | Retrieve the destination stock location associated to the stock transfer
[**GETStockTransferIdOriginStockLocation**](StockLocationsApi.md#GETStockTransferIdOriginStockLocation) | **Get** /stock_transfers/{stockTransferId}/origin_stock_location | Retrieve the origin stock location associated to the stock transfer
[**PATCHStockLocationsStockLocationId**](StockLocationsApi.md#PATCHStockLocationsStockLocationId) | **Patch** /stock_locations/{stockLocationId} | Update a stock location
[**POSTStockLocations**](StockLocationsApi.md#POSTStockLocations) | **Post** /stock_locations | Create a stock location



## DELETEStockLocationsStockLocationId

> DELETEStockLocationsStockLocationId(ctx, stockLocationId).Execute()

Delete a stock location



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
    resp, r, err := apiClient.StockLocationsApi.DELETEStockLocationsStockLocationId(context.Background(), stockLocationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockLocationsApi.DELETEStockLocationsStockLocationId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDELETEStockLocationsStockLocationIdRequest struct via the builder pattern


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


## GETDeliveryLeadTimeIdStockLocation

> GETDeliveryLeadTimeIdStockLocation(ctx, deliveryLeadTimeId).Execute()

Retrieve the stock location associated to the delivery lead time



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
    deliveryLeadTimeId := "deliveryLeadTimeId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StockLocationsApi.GETDeliveryLeadTimeIdStockLocation(context.Background(), deliveryLeadTimeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockLocationsApi.GETDeliveryLeadTimeIdStockLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deliveryLeadTimeId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETDeliveryLeadTimeIdStockLocationRequest struct via the builder pattern


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


## GETInventoryReturnLocationIdStockLocation

> GETInventoryReturnLocationIdStockLocation(ctx, inventoryReturnLocationId).Execute()

Retrieve the stock location associated to the inventory return location



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
    resp, r, err := apiClient.StockLocationsApi.GETInventoryReturnLocationIdStockLocation(context.Background(), inventoryReturnLocationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockLocationsApi.GETInventoryReturnLocationIdStockLocation``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETInventoryReturnLocationIdStockLocationRequest struct via the builder pattern


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


## GETInventoryStockLocationIdStockLocation

> GETInventoryStockLocationIdStockLocation(ctx, inventoryStockLocationId).Execute()

Retrieve the stock location associated to the inventory stock location



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
    resp, r, err := apiClient.StockLocationsApi.GETInventoryStockLocationIdStockLocation(context.Background(), inventoryStockLocationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockLocationsApi.GETInventoryStockLocationIdStockLocation``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETInventoryStockLocationIdStockLocationRequest struct via the builder pattern


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


## GETPackageIdStockLocation

> GETPackageIdStockLocation(ctx, packageId).Execute()

Retrieve the stock location associated to the package



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
    packageId := "packageId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StockLocationsApi.GETPackageIdStockLocation(context.Background(), packageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockLocationsApi.GETPackageIdStockLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPackageIdStockLocationRequest struct via the builder pattern


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


## GETReturnIdStockLocation

> GETReturnIdStockLocation(ctx, returnId).Execute()

Retrieve the stock location associated to the return



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
    returnId := "returnId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StockLocationsApi.GETReturnIdStockLocation(context.Background(), returnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockLocationsApi.GETReturnIdStockLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**returnId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETReturnIdStockLocationRequest struct via the builder pattern


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


## GETShipmentIdStockLocation

> GETShipmentIdStockLocation(ctx, shipmentId).Execute()

Retrieve the stock location associated to the shipment



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
    shipmentId := "shipmentId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StockLocationsApi.GETShipmentIdStockLocation(context.Background(), shipmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockLocationsApi.GETShipmentIdStockLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETShipmentIdStockLocationRequest struct via the builder pattern


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


## GETShippingMethodIdStockLocation

> GETShippingMethodIdStockLocation(ctx, shippingMethodId).Execute()

Retrieve the stock location associated to the shipping method



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
    resp, r, err := apiClient.StockLocationsApi.GETShippingMethodIdStockLocation(context.Background(), shippingMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockLocationsApi.GETShippingMethodIdStockLocation``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETShippingMethodIdStockLocationRequest struct via the builder pattern


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


## GETStockItemIdStockLocation

> GETStockItemIdStockLocation(ctx, stockItemId).Execute()

Retrieve the stock location associated to the stock item



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
    resp, r, err := apiClient.StockLocationsApi.GETStockItemIdStockLocation(context.Background(), stockItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockLocationsApi.GETStockItemIdStockLocation``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETStockItemIdStockLocationRequest struct via the builder pattern


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


## GETStockLocations

> GETStockLocations200Response GETStockLocations(ctx).Execute()

List all stock locations



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
    resp, r, err := apiClient.StockLocationsApi.GETStockLocations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockLocationsApi.GETStockLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETStockLocations`: GETStockLocations200Response
    fmt.Fprintf(os.Stdout, "Response from `StockLocationsApi.GETStockLocations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETStockLocationsRequest struct via the builder pattern


### Return type

[**GETStockLocations200Response**](GETStockLocations200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETStockLocationsStockLocationId

> GETStockLocationsStockLocationId200Response GETStockLocationsStockLocationId(ctx, stockLocationId).Execute()

Retrieve a stock location



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
    resp, r, err := apiClient.StockLocationsApi.GETStockLocationsStockLocationId(context.Background(), stockLocationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockLocationsApi.GETStockLocationsStockLocationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETStockLocationsStockLocationId`: GETStockLocationsStockLocationId200Response
    fmt.Fprintf(os.Stdout, "Response from `StockLocationsApi.GETStockLocationsStockLocationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stockLocationId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETStockLocationsStockLocationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETStockLocationsStockLocationId200Response**](GETStockLocationsStockLocationId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETStockTransferIdDestinationStockLocation

> GETStockTransferIdDestinationStockLocation(ctx, stockTransferId).Execute()

Retrieve the destination stock location associated to the stock transfer



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
    stockTransferId := "stockTransferId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StockLocationsApi.GETStockTransferIdDestinationStockLocation(context.Background(), stockTransferId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockLocationsApi.GETStockTransferIdDestinationStockLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stockTransferId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETStockTransferIdDestinationStockLocationRequest struct via the builder pattern


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


## GETStockTransferIdOriginStockLocation

> GETStockTransferIdOriginStockLocation(ctx, stockTransferId).Execute()

Retrieve the origin stock location associated to the stock transfer



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
    stockTransferId := "stockTransferId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StockLocationsApi.GETStockTransferIdOriginStockLocation(context.Background(), stockTransferId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockLocationsApi.GETStockTransferIdOriginStockLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stockTransferId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETStockTransferIdOriginStockLocationRequest struct via the builder pattern


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


## PATCHStockLocationsStockLocationId

> PATCHStockLocationsStockLocationId200Response PATCHStockLocationsStockLocationId(ctx, stockLocationId).StockLocationUpdate(stockLocationUpdate).Execute()

Update a stock location



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
    stockLocationUpdate := *openapiclient.NewStockLocationUpdate(*openapiclient.NewStockLocationUpdateData("Type_example", "XGZwpOSrWL", *openapiclient.NewPATCHStockLocationsStockLocationId200ResponseDataAttributes())) // StockLocationUpdate | 
    stockLocationId := "stockLocationId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StockLocationsApi.PATCHStockLocationsStockLocationId(context.Background(), stockLocationId).StockLocationUpdate(stockLocationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockLocationsApi.PATCHStockLocationsStockLocationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHStockLocationsStockLocationId`: PATCHStockLocationsStockLocationId200Response
    fmt.Fprintf(os.Stdout, "Response from `StockLocationsApi.PATCHStockLocationsStockLocationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stockLocationId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHStockLocationsStockLocationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stockLocationUpdate** | [**StockLocationUpdate**](StockLocationUpdate.md) |  | 


### Return type

[**PATCHStockLocationsStockLocationId200Response**](PATCHStockLocationsStockLocationId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTStockLocations

> POSTStockLocations201Response POSTStockLocations(ctx).StockLocationCreate(stockLocationCreate).Execute()

Create a stock location



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
    stockLocationCreate := *openapiclient.NewStockLocationCreate(*openapiclient.NewStockLocationCreateData("Type_example", *openapiclient.NewPOSTStockLocations201ResponseDataAttributes("Primary warehouse"))) // StockLocationCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StockLocationsApi.POSTStockLocations(context.Background()).StockLocationCreate(stockLocationCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockLocationsApi.POSTStockLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTStockLocations`: POSTStockLocations201Response
    fmt.Fprintf(os.Stdout, "Response from `StockLocationsApi.POSTStockLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTStockLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stockLocationCreate** | [**StockLocationCreate**](StockLocationCreate.md) |  | 

### Return type

[**POSTStockLocations201Response**](POSTStockLocations201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

