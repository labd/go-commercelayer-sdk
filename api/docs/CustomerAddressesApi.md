# \CustomerAddressesApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETECustomerAddressesCustomerAddressId**](CustomerAddressesApi.md#DELETECustomerAddressesCustomerAddressId) | **Delete** /customer_addresses/{customerAddressId} | Delete a customer address
[**GETCustomerAddresses**](CustomerAddressesApi.md#GETCustomerAddresses) | **Get** /customer_addresses | List all customer addresses
[**GETCustomerAddressesCustomerAddressId**](CustomerAddressesApi.md#GETCustomerAddressesCustomerAddressId) | **Get** /customer_addresses/{customerAddressId} | Retrieve a customer address
[**GETCustomerIdCustomerAddresses**](CustomerAddressesApi.md#GETCustomerIdCustomerAddresses) | **Get** /customers/{customerId}/customer_addresses | Retrieve the customer addresses associated to the customer
[**PATCHCustomerAddressesCustomerAddressId**](CustomerAddressesApi.md#PATCHCustomerAddressesCustomerAddressId) | **Patch** /customer_addresses/{customerAddressId} | Update a customer address
[**POSTCustomerAddresses**](CustomerAddressesApi.md#POSTCustomerAddresses) | **Post** /customer_addresses | Create a customer address



## DELETECustomerAddressesCustomerAddressId

> DELETECustomerAddressesCustomerAddressId(ctx, customerAddressId).Execute()

Delete a customer address



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
    customerAddressId := "customerAddressId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerAddressesApi.DELETECustomerAddressesCustomerAddressId(context.Background(), customerAddressId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerAddressesApi.DELETECustomerAddressesCustomerAddressId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerAddressId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETECustomerAddressesCustomerAddressIdRequest struct via the builder pattern


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


## GETCustomerAddresses

> GETCustomerAddresses(ctx).Execute()

List all customer addresses



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
    resp, r, err := apiClient.CustomerAddressesApi.GETCustomerAddresses(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerAddressesApi.GETCustomerAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETCustomerAddressesRequest struct via the builder pattern


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


## GETCustomerAddressesCustomerAddressId

> CustomerAddress GETCustomerAddressesCustomerAddressId(ctx, customerAddressId).Execute()

Retrieve a customer address



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
    customerAddressId := "customerAddressId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerAddressesApi.GETCustomerAddressesCustomerAddressId(context.Background(), customerAddressId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerAddressesApi.GETCustomerAddressesCustomerAddressId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCustomerAddressesCustomerAddressId`: CustomerAddress
    fmt.Fprintf(os.Stdout, "Response from `CustomerAddressesApi.GETCustomerAddressesCustomerAddressId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerAddressId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCustomerAddressesCustomerAddressIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerAddress**](CustomerAddress.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCustomerIdCustomerAddresses

> GETCustomerIdCustomerAddresses(ctx, customerId).Execute()

Retrieve the customer addresses associated to the customer



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
    customerId := "customerId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerAddressesApi.GETCustomerIdCustomerAddresses(context.Background(), customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerAddressesApi.GETCustomerIdCustomerAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCustomerIdCustomerAddressesRequest struct via the builder pattern


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


## PATCHCustomerAddressesCustomerAddressId

> PATCHCustomerAddressesCustomerAddressId(ctx, customerAddressId).CustomerAddressUpdate(customerAddressUpdate).Execute()

Update a customer address



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
    customerAddressUpdate := *openapiclient.NewCustomerAddressUpdate(*openapiclient.NewCustomerAddressUpdateData("customer_addresses", "XGZwpOSrWL", *openapiclient.NewAdyenPaymentCreateDataAttributes())) // CustomerAddressUpdate | 
    customerAddressId := "customerAddressId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerAddressesApi.PATCHCustomerAddressesCustomerAddressId(context.Background(), customerAddressId).CustomerAddressUpdate(customerAddressUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerAddressesApi.PATCHCustomerAddressesCustomerAddressId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerAddressId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHCustomerAddressesCustomerAddressIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerAddressUpdate** | [**CustomerAddressUpdate**](CustomerAddressUpdate.md) |  | 


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


## POSTCustomerAddresses

> POSTCustomerAddresses(ctx).CustomerAddressCreate(customerAddressCreate).Execute()

Create a customer address



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
    customerAddressCreate := *openapiclient.NewCustomerAddressCreate(*openapiclient.NewCustomerAddressCreateData("customer_addresses", *openapiclient.NewAdyenPaymentCreateDataAttributes())) // CustomerAddressCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerAddressesApi.POSTCustomerAddresses(context.Background()).CustomerAddressCreate(customerAddressCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerAddressesApi.POSTCustomerAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTCustomerAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerAddressCreate** | [**CustomerAddressCreate**](CustomerAddressCreate.md) |  | 

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

