# \CustomerGroupsApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETECustomerGroupsCustomerGroupId**](CustomerGroupsApi.md#DELETECustomerGroupsCustomerGroupId) | **Delete** /customer_groups/{customerGroupId} | Delete a customer group
[**GETCustomerGroups**](CustomerGroupsApi.md#GETCustomerGroups) | **Get** /customer_groups | List all customer groups
[**GETCustomerGroupsCustomerGroupId**](CustomerGroupsApi.md#GETCustomerGroupsCustomerGroupId) | **Get** /customer_groups/{customerGroupId} | Retrieve a customer group
[**GETCustomerIdCustomerGroup**](CustomerGroupsApi.md#GETCustomerIdCustomerGroup) | **Get** /customers/{customerId}/customer_group | Retrieve the customer group associated to the customer
[**GETMarketIdCustomerGroup**](CustomerGroupsApi.md#GETMarketIdCustomerGroup) | **Get** /markets/{marketId}/customer_group | Retrieve the customer group associated to the market
[**PATCHCustomerGroupsCustomerGroupId**](CustomerGroupsApi.md#PATCHCustomerGroupsCustomerGroupId) | **Patch** /customer_groups/{customerGroupId} | Update a customer group
[**POSTCustomerGroups**](CustomerGroupsApi.md#POSTCustomerGroups) | **Post** /customer_groups | Create a customer group



## DELETECustomerGroupsCustomerGroupId

> DELETECustomerGroupsCustomerGroupId(ctx, customerGroupId).Execute()

Delete a customer group



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
    customerGroupId := "customerGroupId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerGroupsApi.DELETECustomerGroupsCustomerGroupId(context.Background(), customerGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerGroupsApi.DELETECustomerGroupsCustomerGroupId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerGroupId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETECustomerGroupsCustomerGroupIdRequest struct via the builder pattern


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


## GETCustomerGroups

> GETCustomerGroups(ctx).Execute()

List all customer groups



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
    resp, r, err := apiClient.CustomerGroupsApi.GETCustomerGroups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerGroupsApi.GETCustomerGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETCustomerGroupsRequest struct via the builder pattern


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


## GETCustomerGroupsCustomerGroupId

> CustomerGroup GETCustomerGroupsCustomerGroupId(ctx, customerGroupId).Execute()

Retrieve a customer group



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
    customerGroupId := "customerGroupId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerGroupsApi.GETCustomerGroupsCustomerGroupId(context.Background(), customerGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerGroupsApi.GETCustomerGroupsCustomerGroupId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCustomerGroupsCustomerGroupId`: CustomerGroup
    fmt.Fprintf(os.Stdout, "Response from `CustomerGroupsApi.GETCustomerGroupsCustomerGroupId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerGroupId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCustomerGroupsCustomerGroupIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerGroup**](CustomerGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCustomerIdCustomerGroup

> GETCustomerIdCustomerGroup(ctx, customerId).Execute()

Retrieve the customer group associated to the customer



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
    resp, r, err := apiClient.CustomerGroupsApi.GETCustomerIdCustomerGroup(context.Background(), customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerGroupsApi.GETCustomerIdCustomerGroup``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETCustomerIdCustomerGroupRequest struct via the builder pattern


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


## GETMarketIdCustomerGroup

> GETMarketIdCustomerGroup(ctx, marketId).Execute()

Retrieve the customer group associated to the market



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
    marketId := "marketId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerGroupsApi.GETMarketIdCustomerGroup(context.Background(), marketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerGroupsApi.GETMarketIdCustomerGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETMarketIdCustomerGroupRequest struct via the builder pattern


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


## PATCHCustomerGroupsCustomerGroupId

> PATCHCustomerGroupsCustomerGroupId(ctx, customerGroupId).CustomerGroupUpdate(customerGroupUpdate).Execute()

Update a customer group



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
    customerGroupId := "customerGroupId_example" // string | The resource's id
    customerGroupUpdate := *openapiclient.NewCustomerGroupUpdate(*openapiclient.NewCustomerGroupUpdateData("customer_groups", "XGZwpOSrWL", *openapiclient.NewCustomerGroupUpdateDataAttributes())) // CustomerGroupUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerGroupsApi.PATCHCustomerGroupsCustomerGroupId(context.Background(), customerGroupId).CustomerGroupUpdate(customerGroupUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerGroupsApi.PATCHCustomerGroupsCustomerGroupId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerGroupId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHCustomerGroupsCustomerGroupIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerGroupUpdate** | [**CustomerGroupUpdate**](CustomerGroupUpdate.md) |  | 

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


## POSTCustomerGroups

> POSTCustomerGroups(ctx).CustomerGroupCreate(customerGroupCreate).Execute()

Create a customer group



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
    customerGroupCreate := *openapiclient.NewCustomerGroupCreate(*openapiclient.NewCustomerGroupCreateData("customer_groups", *openapiclient.NewCustomerGroupCreateDataAttributes("VIP"))) // CustomerGroupCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerGroupsApi.POSTCustomerGroups(context.Background()).CustomerGroupCreate(customerGroupCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerGroupsApi.POSTCustomerGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTCustomerGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerGroupCreate** | [**CustomerGroupCreate**](CustomerGroupCreate.md) |  | 

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

