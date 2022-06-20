# \CustomerPasswordResetsApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETECustomerPasswordResetsCustomerPasswordResetId**](CustomerPasswordResetsApi.md#DELETECustomerPasswordResetsCustomerPasswordResetId) | **Delete** /customer_password_resets/{customerPasswordResetId} | Delete a customer password reset
[**GETCustomerPasswordResets**](CustomerPasswordResetsApi.md#GETCustomerPasswordResets) | **Get** /customer_password_resets | List all customer password resets
[**GETCustomerPasswordResetsCustomerPasswordResetId**](CustomerPasswordResetsApi.md#GETCustomerPasswordResetsCustomerPasswordResetId) | **Get** /customer_password_resets/{customerPasswordResetId} | Retrieve a customer password reset
[**PATCHCustomerPasswordResetsCustomerPasswordResetId**](CustomerPasswordResetsApi.md#PATCHCustomerPasswordResetsCustomerPasswordResetId) | **Patch** /customer_password_resets/{customerPasswordResetId} | Update a customer password reset
[**POSTCustomerPasswordResets**](CustomerPasswordResetsApi.md#POSTCustomerPasswordResets) | **Post** /customer_password_resets | Create a customer password reset



## DELETECustomerPasswordResetsCustomerPasswordResetId

> DELETECustomerPasswordResetsCustomerPasswordResetId(ctx, customerPasswordResetId).Execute()

Delete a customer password reset



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
    customerPasswordResetId := "customerPasswordResetId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerPasswordResetsApi.DELETECustomerPasswordResetsCustomerPasswordResetId(context.Background(), customerPasswordResetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerPasswordResetsApi.DELETECustomerPasswordResetsCustomerPasswordResetId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerPasswordResetId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETECustomerPasswordResetsCustomerPasswordResetIdRequest struct via the builder pattern


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


## GETCustomerPasswordResets

> GETCustomerPasswordResets(ctx).Execute()

List all customer password resets



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
    resp, r, err := apiClient.CustomerPasswordResetsApi.GETCustomerPasswordResets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerPasswordResetsApi.GETCustomerPasswordResets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETCustomerPasswordResetsRequest struct via the builder pattern


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


## GETCustomerPasswordResetsCustomerPasswordResetId

> CustomerPasswordReset GETCustomerPasswordResetsCustomerPasswordResetId(ctx, customerPasswordResetId).Execute()

Retrieve a customer password reset



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
    customerPasswordResetId := "customerPasswordResetId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerPasswordResetsApi.GETCustomerPasswordResetsCustomerPasswordResetId(context.Background(), customerPasswordResetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerPasswordResetsApi.GETCustomerPasswordResetsCustomerPasswordResetId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCustomerPasswordResetsCustomerPasswordResetId`: CustomerPasswordReset
    fmt.Fprintf(os.Stdout, "Response from `CustomerPasswordResetsApi.GETCustomerPasswordResetsCustomerPasswordResetId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerPasswordResetId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCustomerPasswordResetsCustomerPasswordResetIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerPasswordReset**](CustomerPasswordReset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHCustomerPasswordResetsCustomerPasswordResetId

> PATCHCustomerPasswordResetsCustomerPasswordResetId(ctx, customerPasswordResetId).CustomerPasswordResetUpdate(customerPasswordResetUpdate).Execute()

Update a customer password reset



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
    customerPasswordResetId := "customerPasswordResetId_example" // string | The resource's id
    customerPasswordResetUpdate := *openapiclient.NewCustomerPasswordResetUpdate(*openapiclient.NewCustomerPasswordResetUpdateData("customer_password_resets", "XGZwpOSrWL", *openapiclient.NewCustomerPasswordResetUpdateDataAttributes())) // CustomerPasswordResetUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerPasswordResetsApi.PATCHCustomerPasswordResetsCustomerPasswordResetId(context.Background(), customerPasswordResetId).CustomerPasswordResetUpdate(customerPasswordResetUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerPasswordResetsApi.PATCHCustomerPasswordResetsCustomerPasswordResetId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerPasswordResetId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHCustomerPasswordResetsCustomerPasswordResetIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerPasswordResetUpdate** | [**CustomerPasswordResetUpdate**](CustomerPasswordResetUpdate.md) |  | 

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


## POSTCustomerPasswordResets

> POSTCustomerPasswordResets(ctx).CustomerPasswordResetCreate(customerPasswordResetCreate).Execute()

Create a customer password reset



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
    customerPasswordResetCreate := *openapiclient.NewCustomerPasswordResetCreate(*openapiclient.NewCustomerPasswordResetCreateData("customer_password_resets", *openapiclient.NewCustomerPasswordResetCreateDataAttributes("john@example.com"))) // CustomerPasswordResetCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerPasswordResetsApi.POSTCustomerPasswordResets(context.Background()).CustomerPasswordResetCreate(customerPasswordResetCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerPasswordResetsApi.POSTCustomerPasswordResets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTCustomerPasswordResetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerPasswordResetCreate** | [**CustomerPasswordResetCreate**](CustomerPasswordResetCreate.md) |  | 

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

