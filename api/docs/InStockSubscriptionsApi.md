# \InStockSubscriptionsApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEInStockSubscriptionsInStockSubscriptionId**](InStockSubscriptionsApi.md#DELETEInStockSubscriptionsInStockSubscriptionId) | **Delete** /in_stock_subscriptions/{inStockSubscriptionId} | Delete an in stock subscription
[**GETInStockSubscriptions**](InStockSubscriptionsApi.md#GETInStockSubscriptions) | **Get** /in_stock_subscriptions | List all in stock subscriptions
[**GETInStockSubscriptionsInStockSubscriptionId**](InStockSubscriptionsApi.md#GETInStockSubscriptionsInStockSubscriptionId) | **Get** /in_stock_subscriptions/{inStockSubscriptionId} | Retrieve an in stock subscription
[**PATCHInStockSubscriptionsInStockSubscriptionId**](InStockSubscriptionsApi.md#PATCHInStockSubscriptionsInStockSubscriptionId) | **Patch** /in_stock_subscriptions/{inStockSubscriptionId} | Update an in stock subscription
[**POSTInStockSubscriptions**](InStockSubscriptionsApi.md#POSTInStockSubscriptions) | **Post** /in_stock_subscriptions | Create an in stock subscription



## DELETEInStockSubscriptionsInStockSubscriptionId

> DELETEInStockSubscriptionsInStockSubscriptionId(ctx, inStockSubscriptionId).Execute()

Delete an in stock subscription



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
    inStockSubscriptionId := "inStockSubscriptionId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InStockSubscriptionsApi.DELETEInStockSubscriptionsInStockSubscriptionId(context.Background(), inStockSubscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InStockSubscriptionsApi.DELETEInStockSubscriptionsInStockSubscriptionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inStockSubscriptionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEInStockSubscriptionsInStockSubscriptionIdRequest struct via the builder pattern


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


## GETInStockSubscriptions

> GETInStockSubscriptions(ctx).Execute()

List all in stock subscriptions



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
    resp, r, err := apiClient.InStockSubscriptionsApi.GETInStockSubscriptions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InStockSubscriptionsApi.GETInStockSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETInStockSubscriptionsRequest struct via the builder pattern


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


## GETInStockSubscriptionsInStockSubscriptionId

> InStockSubscription GETInStockSubscriptionsInStockSubscriptionId(ctx, inStockSubscriptionId).Execute()

Retrieve an in stock subscription



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
    inStockSubscriptionId := "inStockSubscriptionId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InStockSubscriptionsApi.GETInStockSubscriptionsInStockSubscriptionId(context.Background(), inStockSubscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InStockSubscriptionsApi.GETInStockSubscriptionsInStockSubscriptionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETInStockSubscriptionsInStockSubscriptionId`: InStockSubscription
    fmt.Fprintf(os.Stdout, "Response from `InStockSubscriptionsApi.GETInStockSubscriptionsInStockSubscriptionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inStockSubscriptionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETInStockSubscriptionsInStockSubscriptionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InStockSubscription**](InStockSubscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHInStockSubscriptionsInStockSubscriptionId

> PATCHInStockSubscriptionsInStockSubscriptionId(ctx, inStockSubscriptionId).InStockSubscriptionUpdate(inStockSubscriptionUpdate).Execute()

Update an in stock subscription



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
    inStockSubscriptionUpdate := *openapiclient.NewInStockSubscriptionUpdate(*openapiclient.NewInStockSubscriptionUpdateData("in_stock_subscriptions", "XGZwpOSrWL", *openapiclient.NewInStockSubscriptionUpdateDataAttributes())) // InStockSubscriptionUpdate | 
    inStockSubscriptionId := "inStockSubscriptionId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InStockSubscriptionsApi.PATCHInStockSubscriptionsInStockSubscriptionId(context.Background(), inStockSubscriptionId).InStockSubscriptionUpdate(inStockSubscriptionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InStockSubscriptionsApi.PATCHInStockSubscriptionsInStockSubscriptionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inStockSubscriptionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHInStockSubscriptionsInStockSubscriptionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inStockSubscriptionUpdate** | [**InStockSubscriptionUpdate**](InStockSubscriptionUpdate.md) |  | 


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


## POSTInStockSubscriptions

> POSTInStockSubscriptions(ctx).InStockSubscriptionCreate(inStockSubscriptionCreate).Execute()

Create an in stock subscription



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
    inStockSubscriptionCreate := *openapiclient.NewInStockSubscriptionCreate(*openapiclient.NewInStockSubscriptionCreateData("in_stock_subscriptions", *openapiclient.NewInStockSubscriptionCreateDataAttributes())) // InStockSubscriptionCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InStockSubscriptionsApi.POSTInStockSubscriptions(context.Background()).InStockSubscriptionCreate(inStockSubscriptionCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InStockSubscriptionsApi.POSTInStockSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTInStockSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inStockSubscriptionCreate** | [**InStockSubscriptionCreate**](InStockSubscriptionCreate.md) |  | 

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

