# \EventCallbacksApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETEventCallbacks**](EventCallbacksApi.md#GETEventCallbacks) | **Get** /event_callbacks | List all event callbacks
[**GETEventCallbacksEventCallbackId**](EventCallbacksApi.md#GETEventCallbacksEventCallbackId) | **Get** /event_callbacks/{eventCallbackId} | Retrieve an event callback
[**GETEventIdLastEventCallbacks**](EventCallbacksApi.md#GETEventIdLastEventCallbacks) | **Get** /events/{eventId}/last_event_callbacks | Retrieve the last event callbacks associated to the event
[**GETWebhookIdLastEventCallbacks**](EventCallbacksApi.md#GETWebhookIdLastEventCallbacks) | **Get** /webhooks/{webhookId}/last_event_callbacks | Retrieve the last event callbacks associated to the webhook



## GETEventCallbacks

> GETEventCallbacks(ctx).Execute()

List all event callbacks



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
    resp, r, err := apiClient.EventCallbacksApi.GETEventCallbacks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventCallbacksApi.GETEventCallbacks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETEventCallbacksRequest struct via the builder pattern


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


## GETEventCallbacksEventCallbackId

> EventCallback GETEventCallbacksEventCallbackId(ctx, eventCallbackId).Execute()

Retrieve an event callback



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
    eventCallbackId := "eventCallbackId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventCallbacksApi.GETEventCallbacksEventCallbackId(context.Background(), eventCallbackId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventCallbacksApi.GETEventCallbacksEventCallbackId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETEventCallbacksEventCallbackId`: EventCallback
    fmt.Fprintf(os.Stdout, "Response from `EventCallbacksApi.GETEventCallbacksEventCallbackId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventCallbackId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETEventCallbacksEventCallbackIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventCallback**](EventCallback.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETEventIdLastEventCallbacks

> GETEventIdLastEventCallbacks(ctx, eventId).Execute()

Retrieve the last event callbacks associated to the event



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
    eventId := "eventId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventCallbacksApi.GETEventIdLastEventCallbacks(context.Background(), eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventCallbacksApi.GETEventIdLastEventCallbacks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETEventIdLastEventCallbacksRequest struct via the builder pattern


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


## GETWebhookIdLastEventCallbacks

> GETWebhookIdLastEventCallbacks(ctx, webhookId).Execute()

Retrieve the last event callbacks associated to the webhook



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
    webhookId := "webhookId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventCallbacksApi.GETWebhookIdLastEventCallbacks(context.Background(), webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventCallbacksApi.GETWebhookIdLastEventCallbacks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETWebhookIdLastEventCallbacksRequest struct via the builder pattern


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

