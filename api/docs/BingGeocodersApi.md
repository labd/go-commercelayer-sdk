# \BingGeocodersApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEBingGeocodersBingGeocoderId**](BingGeocodersApi.md#DELETEBingGeocodersBingGeocoderId) | **Delete** /bing_geocoders/{bingGeocoderId} | Delete a bing geocoder
[**GETBingGeocoders**](BingGeocodersApi.md#GETBingGeocoders) | **Get** /bing_geocoders | List all bing geocoders
[**GETBingGeocodersBingGeocoderId**](BingGeocodersApi.md#GETBingGeocodersBingGeocoderId) | **Get** /bing_geocoders/{bingGeocoderId} | Retrieve a bing geocoder
[**PATCHBingGeocodersBingGeocoderId**](BingGeocodersApi.md#PATCHBingGeocodersBingGeocoderId) | **Patch** /bing_geocoders/{bingGeocoderId} | Update a bing geocoder
[**POSTBingGeocoders**](BingGeocodersApi.md#POSTBingGeocoders) | **Post** /bing_geocoders | Create a bing geocoder



## DELETEBingGeocodersBingGeocoderId

> DELETEBingGeocodersBingGeocoderId(ctx, bingGeocoderId).Execute()

Delete a bing geocoder



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
    bingGeocoderId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BingGeocodersApi.DELETEBingGeocodersBingGeocoderId(context.Background(), bingGeocoderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BingGeocodersApi.DELETEBingGeocodersBingGeocoderId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bingGeocoderId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEBingGeocodersBingGeocoderIdRequest struct via the builder pattern


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


## GETBingGeocoders

> GETBingGeocoders200Response GETBingGeocoders(ctx).Execute()

List all bing geocoders



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
    resp, r, err := apiClient.BingGeocodersApi.GETBingGeocoders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BingGeocodersApi.GETBingGeocoders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETBingGeocoders`: GETBingGeocoders200Response
    fmt.Fprintf(os.Stdout, "Response from `BingGeocodersApi.GETBingGeocoders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETBingGeocodersRequest struct via the builder pattern


### Return type

[**GETBingGeocoders200Response**](GETBingGeocoders200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETBingGeocodersBingGeocoderId

> GETBingGeocodersBingGeocoderId200Response GETBingGeocodersBingGeocoderId(ctx, bingGeocoderId).Execute()

Retrieve a bing geocoder



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
    bingGeocoderId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BingGeocodersApi.GETBingGeocodersBingGeocoderId(context.Background(), bingGeocoderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BingGeocodersApi.GETBingGeocodersBingGeocoderId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETBingGeocodersBingGeocoderId`: GETBingGeocodersBingGeocoderId200Response
    fmt.Fprintf(os.Stdout, "Response from `BingGeocodersApi.GETBingGeocodersBingGeocoderId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bingGeocoderId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETBingGeocodersBingGeocoderIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETBingGeocodersBingGeocoderId200Response**](GETBingGeocodersBingGeocoderId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHBingGeocodersBingGeocoderId

> PATCHBingGeocodersBingGeocoderId200Response PATCHBingGeocodersBingGeocoderId(ctx, bingGeocoderId).PATCHBingGeocodersBingGeocoderIdRequest(pATCHBingGeocodersBingGeocoderIdRequest).Execute()

Update a bing geocoder



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
    pATCHBingGeocodersBingGeocoderIdRequest := *openapiclient.NewPATCHBingGeocodersBingGeocoderIdRequest(*openapiclient.NewPATCHBingGeocodersBingGeocoderIdRequestData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHBingGeocodersBingGeocoderIdRequestDataAttributes())) // PATCHBingGeocodersBingGeocoderIdRequest | 
    bingGeocoderId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BingGeocodersApi.PATCHBingGeocodersBingGeocoderId(context.Background(), bingGeocoderId).PATCHBingGeocodersBingGeocoderIdRequest(pATCHBingGeocodersBingGeocoderIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BingGeocodersApi.PATCHBingGeocodersBingGeocoderId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHBingGeocodersBingGeocoderId`: PATCHBingGeocodersBingGeocoderId200Response
    fmt.Fprintf(os.Stdout, "Response from `BingGeocodersApi.PATCHBingGeocodersBingGeocoderId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bingGeocoderId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHBingGeocodersBingGeocoderIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pATCHBingGeocodersBingGeocoderIdRequest** | [**PATCHBingGeocodersBingGeocoderIdRequest**](PATCHBingGeocodersBingGeocoderIdRequest.md) |  | 


### Return type

[**PATCHBingGeocodersBingGeocoderId200Response**](PATCHBingGeocodersBingGeocoderId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTBingGeocoders

> POSTBingGeocoders201Response POSTBingGeocoders(ctx).POSTBingGeocodersRequest(pOSTBingGeocodersRequest).Execute()

Create a bing geocoder



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
    pOSTBingGeocodersRequest := *openapiclient.NewPOSTBingGeocodersRequest(*openapiclient.NewPOSTBingGeocodersRequestData(interface{}(123), *openapiclient.NewPOSTBingGeocodersRequestDataAttributes(interface{}(Default geocoder), interface{}(xxxx-yyyy-zzzz)))) // POSTBingGeocodersRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BingGeocodersApi.POSTBingGeocoders(context.Background()).POSTBingGeocodersRequest(pOSTBingGeocodersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BingGeocodersApi.POSTBingGeocoders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTBingGeocoders`: POSTBingGeocoders201Response
    fmt.Fprintf(os.Stdout, "Response from `BingGeocodersApi.POSTBingGeocoders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTBingGeocodersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pOSTBingGeocodersRequest** | [**POSTBingGeocodersRequest**](POSTBingGeocodersRequest.md) |  | 

### Return type

[**POSTBingGeocoders201Response**](POSTBingGeocoders201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

