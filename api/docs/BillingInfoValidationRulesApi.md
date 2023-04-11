# \BillingInfoValidationRulesApi

All URIs are relative to *https://.commercelayer.io/api*

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
    openapiclient "github.com/incentro-dc/go-commercelayer-sdk/api"
)

func main() {
    billingInfoValidationRuleId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BillingInfoValidationRulesApi.DELETEBillingInfoValidationRulesBillingInfoValidationRuleId(context.Background(), billingInfoValidationRuleId).Execute()
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
**billingInfoValidationRuleId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEBillingInfoValidationRulesBillingInfoValidationRuleIdRequest struct via the builder pattern


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


## GETBillingInfoValidationRules

> GETBillingInfoValidationRules200Response GETBillingInfoValidationRules(ctx).Execute()

List all billing info validation rules



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/incentro-dc/go-commercelayer-sdk/api"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingInfoValidationRulesApi.GETBillingInfoValidationRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingInfoValidationRulesApi.GETBillingInfoValidationRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETBillingInfoValidationRules`: GETBillingInfoValidationRules200Response
    fmt.Fprintf(os.Stdout, "Response from `BillingInfoValidationRulesApi.GETBillingInfoValidationRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETBillingInfoValidationRulesRequest struct via the builder pattern


### Return type

[**GETBillingInfoValidationRules200Response**](GETBillingInfoValidationRules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETBillingInfoValidationRulesBillingInfoValidationRuleId

> GETBillingInfoValidationRulesBillingInfoValidationRuleId200Response GETBillingInfoValidationRulesBillingInfoValidationRuleId(ctx, billingInfoValidationRuleId).Execute()

Retrieve a billing info validation rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/incentro-dc/go-commercelayer-sdk/api"
)

func main() {
    billingInfoValidationRuleId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingInfoValidationRulesApi.GETBillingInfoValidationRulesBillingInfoValidationRuleId(context.Background(), billingInfoValidationRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingInfoValidationRulesApi.GETBillingInfoValidationRulesBillingInfoValidationRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETBillingInfoValidationRulesBillingInfoValidationRuleId`: GETBillingInfoValidationRulesBillingInfoValidationRuleId200Response
    fmt.Fprintf(os.Stdout, "Response from `BillingInfoValidationRulesApi.GETBillingInfoValidationRulesBillingInfoValidationRuleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billingInfoValidationRuleId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETBillingInfoValidationRulesBillingInfoValidationRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETBillingInfoValidationRulesBillingInfoValidationRuleId200Response**](GETBillingInfoValidationRulesBillingInfoValidationRuleId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHBillingInfoValidationRulesBillingInfoValidationRuleId

> PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response PATCHBillingInfoValidationRulesBillingInfoValidationRuleId(ctx, billingInfoValidationRuleId).PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequest(pATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequest).Execute()

Update a billing info validation rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/incentro-dc/go-commercelayer-sdk/api"
)

func main() {
    pATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequest := *openapiclient.NewPATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequest(*openapiclient.NewPATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataAttributes())) // PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequest | 
    billingInfoValidationRuleId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingInfoValidationRulesApi.PATCHBillingInfoValidationRulesBillingInfoValidationRuleId(context.Background(), billingInfoValidationRuleId).PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequest(pATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingInfoValidationRulesApi.PATCHBillingInfoValidationRulesBillingInfoValidationRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHBillingInfoValidationRulesBillingInfoValidationRuleId`: PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response
    fmt.Fprintf(os.Stdout, "Response from `BillingInfoValidationRulesApi.PATCHBillingInfoValidationRulesBillingInfoValidationRuleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billingInfoValidationRuleId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequest** | [**PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequest**](PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequest.md) |  | 


### Return type

[**PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response**](PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTBillingInfoValidationRules

> POSTBillingInfoValidationRules201Response POSTBillingInfoValidationRules(ctx).POSTBillingInfoValidationRulesRequest(pOSTBillingInfoValidationRulesRequest).Execute()

Create a billing info validation rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/incentro-dc/go-commercelayer-sdk/api"
)

func main() {
    pOSTBillingInfoValidationRulesRequest := *openapiclient.NewPOSTBillingInfoValidationRulesRequest(*openapiclient.NewPOSTBillingInfoValidationRulesRequestData(interface{}(123), *openapiclient.NewPOSTAdyenPaymentsRequestDataAttributes())) // POSTBillingInfoValidationRulesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingInfoValidationRulesApi.POSTBillingInfoValidationRules(context.Background()).POSTBillingInfoValidationRulesRequest(pOSTBillingInfoValidationRulesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingInfoValidationRulesApi.POSTBillingInfoValidationRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTBillingInfoValidationRules`: POSTBillingInfoValidationRules201Response
    fmt.Fprintf(os.Stdout, "Response from `BillingInfoValidationRulesApi.POSTBillingInfoValidationRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTBillingInfoValidationRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pOSTBillingInfoValidationRulesRequest** | [**POSTBillingInfoValidationRulesRequest**](POSTBillingInfoValidationRulesRequest.md) |  | 

### Return type

[**POSTBillingInfoValidationRules201Response**](POSTBillingInfoValidationRules201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

