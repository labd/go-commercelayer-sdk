# \GoogleGeocodersApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEGoogleGeocodersGoogleGeocoderId**](GoogleGeocodersApi.md#DELETEGoogleGeocodersGoogleGeocoderId) | **Delete** /google_geocoders/{googleGeocoderId} | Delete a google geocoder
[**GETGoogleGeocoders**](GoogleGeocodersApi.md#GETGoogleGeocoders) | **Get** /google_geocoders | List all google geocoders
[**GETGoogleGeocodersGoogleGeocoderId**](GoogleGeocodersApi.md#GETGoogleGeocodersGoogleGeocoderId) | **Get** /google_geocoders/{googleGeocoderId} | Retrieve a google geocoder
[**PATCHGoogleGeocodersGoogleGeocoderId**](GoogleGeocodersApi.md#PATCHGoogleGeocodersGoogleGeocoderId) | **Patch** /google_geocoders/{googleGeocoderId} | Update a google geocoder
[**POSTGoogleGeocoders**](GoogleGeocodersApi.md#POSTGoogleGeocoders) | **Post** /google_geocoders | Create a google geocoder



## DELETEGoogleGeocodersGoogleGeocoderId

> DELETEGoogleGeocodersGoogleGeocoderId(ctx, googleGeocoderId).Execute()

Delete a google geocoder



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
    googleGeocoderId := "googleGeocoderId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleGeocodersApi.DELETEGoogleGeocodersGoogleGeocoderId(context.Background(), googleGeocoderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleGeocodersApi.DELETEGoogleGeocodersGoogleGeocoderId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**googleGeocoderId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEGoogleGeocodersGoogleGeocoderIdRequest struct via the builder pattern


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


## GETGoogleGeocoders

> GETGoogleGeocoders(ctx).Execute()

List all google geocoders



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
    resp, r, err := apiClient.GoogleGeocodersApi.GETGoogleGeocoders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleGeocodersApi.GETGoogleGeocoders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETGoogleGeocodersRequest struct via the builder pattern


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


## GETGoogleGeocodersGoogleGeocoderId

> GoogleGeocoder GETGoogleGeocodersGoogleGeocoderId(ctx, googleGeocoderId).Execute()

Retrieve a google geocoder



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
    googleGeocoderId := "googleGeocoderId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleGeocodersApi.GETGoogleGeocodersGoogleGeocoderId(context.Background(), googleGeocoderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleGeocodersApi.GETGoogleGeocodersGoogleGeocoderId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETGoogleGeocodersGoogleGeocoderId`: GoogleGeocoder
    fmt.Fprintf(os.Stdout, "Response from `GoogleGeocodersApi.GETGoogleGeocodersGoogleGeocoderId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**googleGeocoderId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETGoogleGeocodersGoogleGeocoderIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GoogleGeocoder**](GoogleGeocoder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHGoogleGeocodersGoogleGeocoderId

> PATCHGoogleGeocodersGoogleGeocoderId(ctx, googleGeocoderId).GoogleGeocoderUpdate(googleGeocoderUpdate).Execute()

Update a google geocoder



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
    googleGeocoderId := "googleGeocoderId_example" // string | The resource's id
    googleGeocoderUpdate := *openapiclient.NewGoogleGeocoderUpdate(*openapiclient.NewGoogleGeocoderUpdateData("google_geocoders", "XGZwpOSrWL", *openapiclient.NewGoogleGeocoderUpdateDataAttributes())) // GoogleGeocoderUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleGeocodersApi.PATCHGoogleGeocodersGoogleGeocoderId(context.Background(), googleGeocoderId).GoogleGeocoderUpdate(googleGeocoderUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleGeocodersApi.PATCHGoogleGeocodersGoogleGeocoderId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**googleGeocoderId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHGoogleGeocodersGoogleGeocoderIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **googleGeocoderUpdate** | [**GoogleGeocoderUpdate**](GoogleGeocoderUpdate.md) |  | 

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


## POSTGoogleGeocoders

> POSTGoogleGeocoders(ctx).GoogleGeocoderCreate(googleGeocoderCreate).Execute()

Create a google geocoder



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
    googleGeocoderCreate := *openapiclient.NewGoogleGeocoderCreate(*openapiclient.NewGoogleGeocoderCreateData("google_geocoders", *openapiclient.NewGoogleGeocoderCreateDataAttributes("Default geocoder", "xxxx-yyyy-zzzz"))) // GoogleGeocoderCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleGeocodersApi.POSTGoogleGeocoders(context.Background()).GoogleGeocoderCreate(googleGeocoderCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleGeocodersApi.POSTGoogleGeocoders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTGoogleGeocodersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **googleGeocoderCreate** | [**GoogleGeocoderCreate**](GoogleGeocoderCreate.md) |  | 

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

