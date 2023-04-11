# \InStockSubscriptionsApi

All URIs are relative to *https://.commercelayer.io/api*

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
    openapiclient "github.com/incentro-dc/go-commercelayer-sdk/api"
)

func main() {
    inStockSubscriptionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.InStockSubscriptionsApi.DELETEInStockSubscriptionsInStockSubscriptionId(context.Background(), inStockSubscriptionId).Execute()
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
**inStockSubscriptionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEInStockSubscriptionsInStockSubscriptionIdRequest struct via the builder pattern


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


## GETInStockSubscriptions

> GETInStockSubscriptions200Response GETInStockSubscriptions(ctx).Execute()

List all in stock subscriptions



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
    resp, r, err := apiClient.InStockSubscriptionsApi.GETInStockSubscriptions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InStockSubscriptionsApi.GETInStockSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETInStockSubscriptions`: GETInStockSubscriptions200Response
    fmt.Fprintf(os.Stdout, "Response from `InStockSubscriptionsApi.GETInStockSubscriptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETInStockSubscriptionsRequest struct via the builder pattern


### Return type

[**GETInStockSubscriptions200Response**](GETInStockSubscriptions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETInStockSubscriptionsInStockSubscriptionId

> GETInStockSubscriptionsInStockSubscriptionId200Response GETInStockSubscriptionsInStockSubscriptionId(ctx, inStockSubscriptionId).Execute()

Retrieve an in stock subscription



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
    inStockSubscriptionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InStockSubscriptionsApi.GETInStockSubscriptionsInStockSubscriptionId(context.Background(), inStockSubscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InStockSubscriptionsApi.GETInStockSubscriptionsInStockSubscriptionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETInStockSubscriptionsInStockSubscriptionId`: GETInStockSubscriptionsInStockSubscriptionId200Response
    fmt.Fprintf(os.Stdout, "Response from `InStockSubscriptionsApi.GETInStockSubscriptionsInStockSubscriptionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inStockSubscriptionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETInStockSubscriptionsInStockSubscriptionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETInStockSubscriptionsInStockSubscriptionId200Response**](GETInStockSubscriptionsInStockSubscriptionId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHInStockSubscriptionsInStockSubscriptionId

> PATCHInStockSubscriptionsInStockSubscriptionId200Response PATCHInStockSubscriptionsInStockSubscriptionId(ctx, inStockSubscriptionId).PATCHInStockSubscriptionsInStockSubscriptionIdRequest(pATCHInStockSubscriptionsInStockSubscriptionIdRequest).Execute()

Update an in stock subscription



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
    pATCHInStockSubscriptionsInStockSubscriptionIdRequest := *openapiclient.NewPATCHInStockSubscriptionsInStockSubscriptionIdRequest(*openapiclient.NewPATCHInStockSubscriptionsInStockSubscriptionIdRequestData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHInStockSubscriptionsInStockSubscriptionIdRequestDataAttributes())) // PATCHInStockSubscriptionsInStockSubscriptionIdRequest | 
    inStockSubscriptionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InStockSubscriptionsApi.PATCHInStockSubscriptionsInStockSubscriptionId(context.Background(), inStockSubscriptionId).PATCHInStockSubscriptionsInStockSubscriptionIdRequest(pATCHInStockSubscriptionsInStockSubscriptionIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InStockSubscriptionsApi.PATCHInStockSubscriptionsInStockSubscriptionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHInStockSubscriptionsInStockSubscriptionId`: PATCHInStockSubscriptionsInStockSubscriptionId200Response
    fmt.Fprintf(os.Stdout, "Response from `InStockSubscriptionsApi.PATCHInStockSubscriptionsInStockSubscriptionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inStockSubscriptionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHInStockSubscriptionsInStockSubscriptionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pATCHInStockSubscriptionsInStockSubscriptionIdRequest** | [**PATCHInStockSubscriptionsInStockSubscriptionIdRequest**](PATCHInStockSubscriptionsInStockSubscriptionIdRequest.md) |  | 


### Return type

[**PATCHInStockSubscriptionsInStockSubscriptionId200Response**](PATCHInStockSubscriptionsInStockSubscriptionId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTInStockSubscriptions

> POSTInStockSubscriptions201Response POSTInStockSubscriptions(ctx).POSTInStockSubscriptionsRequest(pOSTInStockSubscriptionsRequest).Execute()

Create an in stock subscription



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
    pOSTInStockSubscriptionsRequest := *openapiclient.NewPOSTInStockSubscriptionsRequest(*openapiclient.NewPOSTInStockSubscriptionsRequestData(interface{}(123), *openapiclient.NewPOSTInStockSubscriptionsRequestDataAttributes())) // POSTInStockSubscriptionsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InStockSubscriptionsApi.POSTInStockSubscriptions(context.Background()).POSTInStockSubscriptionsRequest(pOSTInStockSubscriptionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InStockSubscriptionsApi.POSTInStockSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTInStockSubscriptions`: POSTInStockSubscriptions201Response
    fmt.Fprintf(os.Stdout, "Response from `InStockSubscriptionsApi.POSTInStockSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTInStockSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pOSTInStockSubscriptionsRequest** | [**POSTInStockSubscriptionsRequest**](POSTInStockSubscriptionsRequest.md) |  | 

### Return type

[**POSTInStockSubscriptions201Response**](POSTInStockSubscriptions201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

