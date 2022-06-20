# \ManualTaxCalculatorsApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEManualTaxCalculatorsManualTaxCalculatorId**](ManualTaxCalculatorsApi.md#DELETEManualTaxCalculatorsManualTaxCalculatorId) | **Delete** /manual_tax_calculators/{manualTaxCalculatorId} | Delete a manual tax calculator
[**GETManualTaxCalculators**](ManualTaxCalculatorsApi.md#GETManualTaxCalculators) | **Get** /manual_tax_calculators | List all manual tax calculators
[**GETManualTaxCalculatorsManualTaxCalculatorId**](ManualTaxCalculatorsApi.md#GETManualTaxCalculatorsManualTaxCalculatorId) | **Get** /manual_tax_calculators/{manualTaxCalculatorId} | Retrieve a manual tax calculator
[**GETTaxRuleIdManualTaxCalculator**](ManualTaxCalculatorsApi.md#GETTaxRuleIdManualTaxCalculator) | **Get** /tax_rules/{taxRuleId}/manual_tax_calculator | Retrieve the manual tax calculator associated to the tax rule
[**PATCHManualTaxCalculatorsManualTaxCalculatorId**](ManualTaxCalculatorsApi.md#PATCHManualTaxCalculatorsManualTaxCalculatorId) | **Patch** /manual_tax_calculators/{manualTaxCalculatorId} | Update a manual tax calculator
[**POSTManualTaxCalculators**](ManualTaxCalculatorsApi.md#POSTManualTaxCalculators) | **Post** /manual_tax_calculators | Create a manual tax calculator



## DELETEManualTaxCalculatorsManualTaxCalculatorId

> DELETEManualTaxCalculatorsManualTaxCalculatorId(ctx, manualTaxCalculatorId).Execute()

Delete a manual tax calculator



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
    resp, r, err := apiClient.ManualTaxCalculatorsApi.DELETEManualTaxCalculatorsManualTaxCalculatorId(context.Background(), manualTaxCalculatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManualTaxCalculatorsApi.DELETEManualTaxCalculatorsManualTaxCalculatorId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDELETEManualTaxCalculatorsManualTaxCalculatorIdRequest struct via the builder pattern


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


## GETManualTaxCalculators

> GETManualTaxCalculators(ctx).Execute()

List all manual tax calculators



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
    resp, r, err := apiClient.ManualTaxCalculatorsApi.GETManualTaxCalculators(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManualTaxCalculatorsApi.GETManualTaxCalculators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETManualTaxCalculatorsRequest struct via the builder pattern


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


## GETManualTaxCalculatorsManualTaxCalculatorId

> ManualTaxCalculator GETManualTaxCalculatorsManualTaxCalculatorId(ctx, manualTaxCalculatorId).Execute()

Retrieve a manual tax calculator



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
    resp, r, err := apiClient.ManualTaxCalculatorsApi.GETManualTaxCalculatorsManualTaxCalculatorId(context.Background(), manualTaxCalculatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManualTaxCalculatorsApi.GETManualTaxCalculatorsManualTaxCalculatorId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETManualTaxCalculatorsManualTaxCalculatorId`: ManualTaxCalculator
    fmt.Fprintf(os.Stdout, "Response from `ManualTaxCalculatorsApi.GETManualTaxCalculatorsManualTaxCalculatorId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**manualTaxCalculatorId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETManualTaxCalculatorsManualTaxCalculatorIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManualTaxCalculator**](ManualTaxCalculator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETTaxRuleIdManualTaxCalculator

> GETTaxRuleIdManualTaxCalculator(ctx, taxRuleId).Execute()

Retrieve the manual tax calculator associated to the tax rule



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
    taxRuleId := "taxRuleId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManualTaxCalculatorsApi.GETTaxRuleIdManualTaxCalculator(context.Background(), taxRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManualTaxCalculatorsApi.GETTaxRuleIdManualTaxCalculator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxRuleId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETTaxRuleIdManualTaxCalculatorRequest struct via the builder pattern


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


## PATCHManualTaxCalculatorsManualTaxCalculatorId

> PATCHManualTaxCalculatorsManualTaxCalculatorId(ctx, manualTaxCalculatorId).ManualTaxCalculatorUpdate(manualTaxCalculatorUpdate).Execute()

Update a manual tax calculator



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
    manualTaxCalculatorUpdate := *openapiclient.NewManualTaxCalculatorUpdate(*openapiclient.NewManualTaxCalculatorUpdateData("manual_tax_calculators", "XGZwpOSrWL", *openapiclient.NewManualTaxCalculatorUpdateDataAttributes())) // ManualTaxCalculatorUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManualTaxCalculatorsApi.PATCHManualTaxCalculatorsManualTaxCalculatorId(context.Background(), manualTaxCalculatorId).ManualTaxCalculatorUpdate(manualTaxCalculatorUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManualTaxCalculatorsApi.PATCHManualTaxCalculatorsManualTaxCalculatorId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPATCHManualTaxCalculatorsManualTaxCalculatorIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **manualTaxCalculatorUpdate** | [**ManualTaxCalculatorUpdate**](ManualTaxCalculatorUpdate.md) |  | 

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


## POSTManualTaxCalculators

> POSTManualTaxCalculators(ctx).ManualTaxCalculatorCreate(manualTaxCalculatorCreate).Execute()

Create a manual tax calculator



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
    manualTaxCalculatorCreate := *openapiclient.NewManualTaxCalculatorCreate(*openapiclient.NewManualTaxCalculatorCreateData("manual_tax_calculators", *openapiclient.NewManualTaxCalculatorCreateDataAttributes("Personal tax calculator"))) // ManualTaxCalculatorCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManualTaxCalculatorsApi.POSTManualTaxCalculators(context.Background()).ManualTaxCalculatorCreate(manualTaxCalculatorCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManualTaxCalculatorsApi.POSTManualTaxCalculators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTManualTaxCalculatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **manualTaxCalculatorCreate** | [**ManualTaxCalculatorCreate**](ManualTaxCalculatorCreate.md) |  | 

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

