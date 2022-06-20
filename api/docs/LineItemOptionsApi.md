# \LineItemOptionsApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETELineItemOptionsLineItemOptionId**](LineItemOptionsApi.md#DELETELineItemOptionsLineItemOptionId) | **Delete** /line_item_options/{lineItemOptionId} | Delete a line item option
[**GETLineItemIdLineItemOptions**](LineItemOptionsApi.md#GETLineItemIdLineItemOptions) | **Get** /line_items/{lineItemId}/line_item_options | Retrieve the line item options associated to the line item
[**GETLineItemOptions**](LineItemOptionsApi.md#GETLineItemOptions) | **Get** /line_item_options | List all line item options
[**GETLineItemOptionsLineItemOptionId**](LineItemOptionsApi.md#GETLineItemOptionsLineItemOptionId) | **Get** /line_item_options/{lineItemOptionId} | Retrieve a line item option
[**PATCHLineItemOptionsLineItemOptionId**](LineItemOptionsApi.md#PATCHLineItemOptionsLineItemOptionId) | **Patch** /line_item_options/{lineItemOptionId} | Update a line item option
[**POSTLineItemOptions**](LineItemOptionsApi.md#POSTLineItemOptions) | **Post** /line_item_options | Create a line item option



## DELETELineItemOptionsLineItemOptionId

> DELETELineItemOptionsLineItemOptionId(ctx, lineItemOptionId).Execute()

Delete a line item option



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
    lineItemOptionId := "lineItemOptionId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LineItemOptionsApi.DELETELineItemOptionsLineItemOptionId(context.Background(), lineItemOptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineItemOptionsApi.DELETELineItemOptionsLineItemOptionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineItemOptionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETELineItemOptionsLineItemOptionIdRequest struct via the builder pattern


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


## GETLineItemIdLineItemOptions

> GETLineItemIdLineItemOptions(ctx, lineItemId).Execute()

Retrieve the line item options associated to the line item



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
    lineItemId := "lineItemId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LineItemOptionsApi.GETLineItemIdLineItemOptions(context.Background(), lineItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineItemOptionsApi.GETLineItemIdLineItemOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineItemId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETLineItemIdLineItemOptionsRequest struct via the builder pattern


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


## GETLineItemOptions

> GETLineItemOptions(ctx).Execute()

List all line item options



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
    resp, r, err := apiClient.LineItemOptionsApi.GETLineItemOptions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineItemOptionsApi.GETLineItemOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETLineItemOptionsRequest struct via the builder pattern


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


## GETLineItemOptionsLineItemOptionId

> LineItemOption GETLineItemOptionsLineItemOptionId(ctx, lineItemOptionId).Execute()

Retrieve a line item option



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
    lineItemOptionId := "lineItemOptionId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LineItemOptionsApi.GETLineItemOptionsLineItemOptionId(context.Background(), lineItemOptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineItemOptionsApi.GETLineItemOptionsLineItemOptionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETLineItemOptionsLineItemOptionId`: LineItemOption
    fmt.Fprintf(os.Stdout, "Response from `LineItemOptionsApi.GETLineItemOptionsLineItemOptionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineItemOptionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETLineItemOptionsLineItemOptionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LineItemOption**](LineItemOption.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHLineItemOptionsLineItemOptionId

> PATCHLineItemOptionsLineItemOptionId(ctx, lineItemOptionId).LineItemOptionUpdate(lineItemOptionUpdate).Execute()

Update a line item option



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
    lineItemOptionId := "lineItemOptionId_example" // string | The resource's id
    lineItemOptionUpdate := *openapiclient.NewLineItemOptionUpdate(*openapiclient.NewLineItemOptionUpdateData("line_item_options", "XGZwpOSrWL", *openapiclient.NewLineItemOptionUpdateDataAttributes())) // LineItemOptionUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LineItemOptionsApi.PATCHLineItemOptionsLineItemOptionId(context.Background(), lineItemOptionId).LineItemOptionUpdate(lineItemOptionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineItemOptionsApi.PATCHLineItemOptionsLineItemOptionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineItemOptionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHLineItemOptionsLineItemOptionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lineItemOptionUpdate** | [**LineItemOptionUpdate**](LineItemOptionUpdate.md) |  | 

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


## POSTLineItemOptions

> POSTLineItemOptions(ctx).LineItemOptionCreate(lineItemOptionCreate).Execute()

Create a line item option



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
    lineItemOptionCreate := *openapiclient.NewLineItemOptionCreate(*openapiclient.NewLineItemOptionCreateData("line_item_options", *openapiclient.NewLineItemOptionCreateDataAttributes(int32(2), map[string]interface{}({"embossing_text":"Happy Birthday!"})))) // LineItemOptionCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LineItemOptionsApi.POSTLineItemOptions(context.Background()).LineItemOptionCreate(lineItemOptionCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineItemOptionsApi.POSTLineItemOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTLineItemOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lineItemOptionCreate** | [**LineItemOptionCreate**](LineItemOptionCreate.md) |  | 

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

