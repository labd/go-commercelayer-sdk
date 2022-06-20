# \ShippingCategoriesApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEShippingCategoriesShippingCategoryId**](ShippingCategoriesApi.md#DELETEShippingCategoriesShippingCategoryId) | **Delete** /shipping_categories/{shippingCategoryId} | Delete a shipping category
[**GETShipmentIdShippingCategory**](ShippingCategoriesApi.md#GETShipmentIdShippingCategory) | **Get** /shipments/{shipmentId}/shipping_category | Retrieve the shipping category associated to the shipment
[**GETShippingCategories**](ShippingCategoriesApi.md#GETShippingCategories) | **Get** /shipping_categories | List all shipping categories
[**GETShippingCategoriesShippingCategoryId**](ShippingCategoriesApi.md#GETShippingCategoriesShippingCategoryId) | **Get** /shipping_categories/{shippingCategoryId} | Retrieve a shipping category
[**GETShippingMethodIdShippingCategory**](ShippingCategoriesApi.md#GETShippingMethodIdShippingCategory) | **Get** /shipping_methods/{shippingMethodId}/shipping_category | Retrieve the shipping category associated to the shipping method
[**GETSkuIdShippingCategory**](ShippingCategoriesApi.md#GETSkuIdShippingCategory) | **Get** /skus/{skuId}/shipping_category | Retrieve the shipping category associated to the SKU
[**PATCHShippingCategoriesShippingCategoryId**](ShippingCategoriesApi.md#PATCHShippingCategoriesShippingCategoryId) | **Patch** /shipping_categories/{shippingCategoryId} | Update a shipping category
[**POSTShippingCategories**](ShippingCategoriesApi.md#POSTShippingCategories) | **Post** /shipping_categories | Create a shipping category



## DELETEShippingCategoriesShippingCategoryId

> DELETEShippingCategoriesShippingCategoryId(ctx, shippingCategoryId).Execute()

Delete a shipping category



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
    shippingCategoryId := "shippingCategoryId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingCategoriesApi.DELETEShippingCategoriesShippingCategoryId(context.Background(), shippingCategoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingCategoriesApi.DELETEShippingCategoriesShippingCategoryId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shippingCategoryId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEShippingCategoriesShippingCategoryIdRequest struct via the builder pattern


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


## GETShipmentIdShippingCategory

> GETShipmentIdShippingCategory(ctx, shipmentId).Execute()

Retrieve the shipping category associated to the shipment



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
    resp, r, err := apiClient.ShippingCategoriesApi.GETShipmentIdShippingCategory(context.Background(), shipmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingCategoriesApi.GETShipmentIdShippingCategory``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETShipmentIdShippingCategoryRequest struct via the builder pattern


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


## GETShippingCategories

> GETShippingCategories(ctx).Execute()

List all shipping categories



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
    resp, r, err := apiClient.ShippingCategoriesApi.GETShippingCategories(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingCategoriesApi.GETShippingCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETShippingCategoriesRequest struct via the builder pattern


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


## GETShippingCategoriesShippingCategoryId

> ShippingCategory GETShippingCategoriesShippingCategoryId(ctx, shippingCategoryId).Execute()

Retrieve a shipping category



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
    shippingCategoryId := "shippingCategoryId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingCategoriesApi.GETShippingCategoriesShippingCategoryId(context.Background(), shippingCategoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingCategoriesApi.GETShippingCategoriesShippingCategoryId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETShippingCategoriesShippingCategoryId`: ShippingCategory
    fmt.Fprintf(os.Stdout, "Response from `ShippingCategoriesApi.GETShippingCategoriesShippingCategoryId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shippingCategoryId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETShippingCategoriesShippingCategoryIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ShippingCategory**](ShippingCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETShippingMethodIdShippingCategory

> GETShippingMethodIdShippingCategory(ctx, shippingMethodId).Execute()

Retrieve the shipping category associated to the shipping method



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
    resp, r, err := apiClient.ShippingCategoriesApi.GETShippingMethodIdShippingCategory(context.Background(), shippingMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingCategoriesApi.GETShippingMethodIdShippingCategory``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETShippingMethodIdShippingCategoryRequest struct via the builder pattern


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


## GETSkuIdShippingCategory

> GETSkuIdShippingCategory(ctx, skuId).Execute()

Retrieve the shipping category associated to the SKU



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
    resp, r, err := apiClient.ShippingCategoriesApi.GETSkuIdShippingCategory(context.Background(), skuId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingCategoriesApi.GETSkuIdShippingCategory``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETSkuIdShippingCategoryRequest struct via the builder pattern


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


## PATCHShippingCategoriesShippingCategoryId

> PATCHShippingCategoriesShippingCategoryId(ctx, shippingCategoryId).ShippingCategoryUpdate(shippingCategoryUpdate).Execute()

Update a shipping category



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
    shippingCategoryId := "shippingCategoryId_example" // string | The resource's id
    shippingCategoryUpdate := *openapiclient.NewShippingCategoryUpdate(*openapiclient.NewShippingCategoryUpdateData("shipping_categories", "XGZwpOSrWL", *openapiclient.NewShippingCategoryUpdateDataAttributes())) // ShippingCategoryUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingCategoriesApi.PATCHShippingCategoriesShippingCategoryId(context.Background(), shippingCategoryId).ShippingCategoryUpdate(shippingCategoryUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingCategoriesApi.PATCHShippingCategoriesShippingCategoryId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shippingCategoryId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHShippingCategoriesShippingCategoryIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shippingCategoryUpdate** | [**ShippingCategoryUpdate**](ShippingCategoryUpdate.md) |  | 

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


## POSTShippingCategories

> POSTShippingCategories(ctx).ShippingCategoryCreate(shippingCategoryCreate).Execute()

Create a shipping category



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
    shippingCategoryCreate := *openapiclient.NewShippingCategoryCreate(*openapiclient.NewShippingCategoryCreateData("shipping_categories", *openapiclient.NewShippingCategoryCreateDataAttributes("Merchandise"))) // ShippingCategoryCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingCategoriesApi.POSTShippingCategories(context.Background()).ShippingCategoryCreate(shippingCategoryCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingCategoriesApi.POSTShippingCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTShippingCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shippingCategoryCreate** | [**ShippingCategoryCreate**](ShippingCategoryCreate.md) |  | 

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

