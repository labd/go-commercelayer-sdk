# \TaxCategoriesApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETETaxCategoriesTaxCategoryId**](TaxCategoriesApi.md#DELETETaxCategoriesTaxCategoryId) | **Delete** /tax_categories/{taxCategoryId} | Delete a tax category
[**GETAvalaraAccountIdTaxCategories**](TaxCategoriesApi.md#GETAvalaraAccountIdTaxCategories) | **Get** /avalara_accounts/{avalaraAccountId}/tax_categories | Retrieve the tax categories associated to the avalara account
[**GETExternalTaxCalculatorIdTaxCategories**](TaxCategoriesApi.md#GETExternalTaxCalculatorIdTaxCategories) | **Get** /external_tax_calculators/{externalTaxCalculatorId}/tax_categories | Retrieve the tax categories associated to the external tax calculator
[**GETManualTaxCalculatorIdTaxCategories**](TaxCategoriesApi.md#GETManualTaxCalculatorIdTaxCategories) | **Get** /manual_tax_calculators/{manualTaxCalculatorId}/tax_categories | Retrieve the tax categories associated to the manual tax calculator
[**GETTaxCalculatorIdTaxCategories**](TaxCategoriesApi.md#GETTaxCalculatorIdTaxCategories) | **Get** /tax_calculators/{taxCalculatorId}/tax_categories | Retrieve the tax categories associated to the tax calculator
[**GETTaxCategories**](TaxCategoriesApi.md#GETTaxCategories) | **Get** /tax_categories | List all tax categories
[**GETTaxCategoriesTaxCategoryId**](TaxCategoriesApi.md#GETTaxCategoriesTaxCategoryId) | **Get** /tax_categories/{taxCategoryId} | Retrieve a tax category
[**GETTaxjarAccountIdTaxCategories**](TaxCategoriesApi.md#GETTaxjarAccountIdTaxCategories) | **Get** /taxjar_accounts/{taxjarAccountId}/tax_categories | Retrieve the tax categories associated to the taxjar account
[**PATCHTaxCategoriesTaxCategoryId**](TaxCategoriesApi.md#PATCHTaxCategoriesTaxCategoryId) | **Patch** /tax_categories/{taxCategoryId} | Update a tax category
[**POSTTaxCategories**](TaxCategoriesApi.md#POSTTaxCategories) | **Post** /tax_categories | Create a tax category



## DELETETaxCategoriesTaxCategoryId

> DELETETaxCategoriesTaxCategoryId(ctx, taxCategoryId).Execute()

Delete a tax category



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
    taxCategoryId := "taxCategoryId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxCategoriesApi.DELETETaxCategoriesTaxCategoryId(context.Background(), taxCategoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxCategoriesApi.DELETETaxCategoriesTaxCategoryId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxCategoryId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETETaxCategoriesTaxCategoryIdRequest struct via the builder pattern


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


## GETAvalaraAccountIdTaxCategories

> GETAvalaraAccountIdTaxCategories(ctx, avalaraAccountId).Execute()

Retrieve the tax categories associated to the avalara account



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
    avalaraAccountId := "avalaraAccountId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxCategoriesApi.GETAvalaraAccountIdTaxCategories(context.Background(), avalaraAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxCategoriesApi.GETAvalaraAccountIdTaxCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**avalaraAccountId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETAvalaraAccountIdTaxCategoriesRequest struct via the builder pattern


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


## GETExternalTaxCalculatorIdTaxCategories

> GETExternalTaxCalculatorIdTaxCategories(ctx, externalTaxCalculatorId).Execute()

Retrieve the tax categories associated to the external tax calculator



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
    externalTaxCalculatorId := "externalTaxCalculatorId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxCategoriesApi.GETExternalTaxCalculatorIdTaxCategories(context.Background(), externalTaxCalculatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxCategoriesApi.GETExternalTaxCalculatorIdTaxCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalTaxCalculatorId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETExternalTaxCalculatorIdTaxCategoriesRequest struct via the builder pattern


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


## GETManualTaxCalculatorIdTaxCategories

> GETManualTaxCalculatorIdTaxCategories(ctx, manualTaxCalculatorId).Execute()

Retrieve the tax categories associated to the manual tax calculator



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
    manualTaxCalculatorId := "manualTaxCalculatorId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxCategoriesApi.GETManualTaxCalculatorIdTaxCategories(context.Background(), manualTaxCalculatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxCategoriesApi.GETManualTaxCalculatorIdTaxCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**manualTaxCalculatorId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETManualTaxCalculatorIdTaxCategoriesRequest struct via the builder pattern


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


## GETTaxCalculatorIdTaxCategories

> GETTaxCalculatorIdTaxCategories(ctx, taxCalculatorId).Execute()

Retrieve the tax categories associated to the tax calculator



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
    taxCalculatorId := "taxCalculatorId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxCategoriesApi.GETTaxCalculatorIdTaxCategories(context.Background(), taxCalculatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxCategoriesApi.GETTaxCalculatorIdTaxCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxCalculatorId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETTaxCalculatorIdTaxCategoriesRequest struct via the builder pattern


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


## GETTaxCategories

> GETTaxCategories(ctx).Execute()

List all tax categories



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
    resp, r, err := apiClient.TaxCategoriesApi.GETTaxCategories(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxCategoriesApi.GETTaxCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETTaxCategoriesRequest struct via the builder pattern


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


## GETTaxCategoriesTaxCategoryId

> TaxCategory GETTaxCategoriesTaxCategoryId(ctx, taxCategoryId).Execute()

Retrieve a tax category



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
    taxCategoryId := "taxCategoryId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxCategoriesApi.GETTaxCategoriesTaxCategoryId(context.Background(), taxCategoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxCategoriesApi.GETTaxCategoriesTaxCategoryId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTaxCategoriesTaxCategoryId`: TaxCategory
    fmt.Fprintf(os.Stdout, "Response from `TaxCategoriesApi.GETTaxCategoriesTaxCategoryId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxCategoryId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETTaxCategoriesTaxCategoryIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaxCategory**](TaxCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETTaxjarAccountIdTaxCategories

> GETTaxjarAccountIdTaxCategories(ctx, taxjarAccountId).Execute()

Retrieve the tax categories associated to the taxjar account



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
    taxjarAccountId := "taxjarAccountId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxCategoriesApi.GETTaxjarAccountIdTaxCategories(context.Background(), taxjarAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxCategoriesApi.GETTaxjarAccountIdTaxCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxjarAccountId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETTaxjarAccountIdTaxCategoriesRequest struct via the builder pattern


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


## PATCHTaxCategoriesTaxCategoryId

> PATCHTaxCategoriesTaxCategoryId(ctx, taxCategoryId).TaxCategoryUpdate(taxCategoryUpdate).Execute()

Update a tax category



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
    taxCategoryId := "taxCategoryId_example" // string | The resource's id
    taxCategoryUpdate := *openapiclient.NewTaxCategoryUpdate(*openapiclient.NewTaxCategoryUpdateData("tax_categories", "XGZwpOSrWL", *openapiclient.NewTaxCategoryUpdateDataAttributes())) // TaxCategoryUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxCategoriesApi.PATCHTaxCategoriesTaxCategoryId(context.Background(), taxCategoryId).TaxCategoryUpdate(taxCategoryUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxCategoriesApi.PATCHTaxCategoriesTaxCategoryId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxCategoryId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHTaxCategoriesTaxCategoryIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taxCategoryUpdate** | [**TaxCategoryUpdate**](TaxCategoryUpdate.md) |  | 

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


## POSTTaxCategories

> POSTTaxCategories(ctx).TaxCategoryCreate(taxCategoryCreate).Execute()

Create a tax category



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
    taxCategoryCreate := *openapiclient.NewTaxCategoryCreate(*openapiclient.NewTaxCategoryCreateData("tax_categories", *openapiclient.NewTaxCategoryCreateDataAttributes("31000"))) // TaxCategoryCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxCategoriesApi.POSTTaxCategories(context.Background()).TaxCategoryCreate(taxCategoryCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxCategoriesApi.POSTTaxCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTTaxCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxCategoryCreate** | [**TaxCategoryCreate**](TaxCategoryCreate.md) |  | 

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

