# \BillingInfoValidationRulesApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEBillingInfoValidationRulesBillingInfoValidationRuleId**](BillingInfoValidationRulesApi.md#DELETEBillingInfoValidationRulesBillingInfoValidationRuleId) | **Delete** /billing_info_validation_rules/{billingInfoValidationRuleId} | Delete a billing info validation rule
[**GETBillingInfoValidationRules**](BillingInfoValidationRulesApi.md#GETBillingInfoValidationRules) | **Get** /billing_info_validation_rules | List all billing info validation rules
[**GETBillingInfoValidationRulesBillingInfoValidationRuleId**](BillingInfoValidationRulesApi.md#GETBillingInfoValidationRulesBillingInfoValidationRuleId) | **Get** /billing_info_validation_rules/{billingInfoValidationRuleId} | Retrieve a billing info validation rule
[**PATCHBillingInfoValidationRulesBillingInfoValidationRuleId**](BillingInfoValidationRulesApi.md#PATCHBillingInfoValidationRulesBillingInfoValidationRuleId) | **Patch** /billing_info_validation_rules/{billingInfoValidationRuleId} | Update a billing info validation rule
[**POSTBillingInfoValidationRules**](BillingInfoValidationRulesApi.md#POSTBillingInfoValidationRules) | **Post** /billing_info_validation_rules | Create a billing info validation rule



## DELETEBillingInfoValidationRulesBillingInfoValidationRuleId

> DELETEBillingInfoValidationRulesBillingInfoValidationRuleId(ctx, billingInfoValidationRuleId).Execute()

Delete a billing info validation rule



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
    billingInfoValidationRuleId := "billingInfoValidationRuleId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingInfoValidationRulesApi.DELETEBillingInfoValidationRulesBillingInfoValidationRuleId(context.Background(), billingInfoValidationRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingInfoValidationRulesApi.DELETEBillingInfoValidationRulesBillingInfoValidationRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billingInfoValidationRuleId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEBillingInfoValidationRulesBillingInfoValidationRuleIdRequest struct via the builder pattern


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


## GETBillingInfoValidationRules

> GETBillingInfoValidationRules(ctx).Execute()

List all billing info validation rules



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
    resp, r, err := apiClient.BillingInfoValidationRulesApi.GETBillingInfoValidationRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingInfoValidationRulesApi.GETBillingInfoValidationRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETBillingInfoValidationRulesRequest struct via the builder pattern


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


## GETBillingInfoValidationRulesBillingInfoValidationRuleId

> BillingInfoValidationRule GETBillingInfoValidationRulesBillingInfoValidationRuleId(ctx, billingInfoValidationRuleId).Execute()

Retrieve a billing info validation rule



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
    billingInfoValidationRuleId := "billingInfoValidationRuleId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingInfoValidationRulesApi.GETBillingInfoValidationRulesBillingInfoValidationRuleId(context.Background(), billingInfoValidationRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingInfoValidationRulesApi.GETBillingInfoValidationRulesBillingInfoValidationRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETBillingInfoValidationRulesBillingInfoValidationRuleId`: BillingInfoValidationRule
    fmt.Fprintf(os.Stdout, "Response from `BillingInfoValidationRulesApi.GETBillingInfoValidationRulesBillingInfoValidationRuleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billingInfoValidationRuleId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETBillingInfoValidationRulesBillingInfoValidationRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BillingInfoValidationRule**](BillingInfoValidationRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHBillingInfoValidationRulesBillingInfoValidationRuleId

> PATCHBillingInfoValidationRulesBillingInfoValidationRuleId(ctx, billingInfoValidationRuleId).BillingInfoValidationRuleUpdate(billingInfoValidationRuleUpdate).Execute()

Update a billing info validation rule



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
    billingInfoValidationRuleId := "billingInfoValidationRuleId_example" // string | The resource's id
    billingInfoValidationRuleUpdate := *openapiclient.NewBillingInfoValidationRuleUpdate(*openapiclient.NewBillingInfoValidationRuleUpdateData("billing_info_validation_rules", "XGZwpOSrWL", *openapiclient.NewAdyenPaymentCreateDataAttributes())) // BillingInfoValidationRuleUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingInfoValidationRulesApi.PATCHBillingInfoValidationRulesBillingInfoValidationRuleId(context.Background(), billingInfoValidationRuleId).BillingInfoValidationRuleUpdate(billingInfoValidationRuleUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingInfoValidationRulesApi.PATCHBillingInfoValidationRulesBillingInfoValidationRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billingInfoValidationRuleId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **billingInfoValidationRuleUpdate** | [**BillingInfoValidationRuleUpdate**](BillingInfoValidationRuleUpdate.md) |  | 

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


## POSTBillingInfoValidationRules

> POSTBillingInfoValidationRules(ctx).BillingInfoValidationRuleCreate(billingInfoValidationRuleCreate).Execute()

Create a billing info validation rule



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
    billingInfoValidationRuleCreate := *openapiclient.NewBillingInfoValidationRuleCreate(*openapiclient.NewBillingInfoValidationRuleCreateData("billing_info_validation_rules", *openapiclient.NewAdyenPaymentCreateDataAttributes())) // BillingInfoValidationRuleCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingInfoValidationRulesApi.POSTBillingInfoValidationRules(context.Background()).BillingInfoValidationRuleCreate(billingInfoValidationRuleCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingInfoValidationRulesApi.POSTBillingInfoValidationRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTBillingInfoValidationRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **billingInfoValidationRuleCreate** | [**BillingInfoValidationRuleCreate**](BillingInfoValidationRuleCreate.md) |  | 

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

