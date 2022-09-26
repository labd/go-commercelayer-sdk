# \CouponRecipientsApi

All URIs are relative to *https://}.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETECouponRecipientsCouponRecipientId**](CouponRecipientsApi.md#DELETECouponRecipientsCouponRecipientId) | **Delete** /coupon_recipients/{couponRecipientId} | Delete a coupon recipient
[**GETCouponRecipients**](CouponRecipientsApi.md#GETCouponRecipients) | **Get** /coupon_recipients | List all coupon recipients
[**GETCouponRecipientsCouponRecipientId**](CouponRecipientsApi.md#GETCouponRecipientsCouponRecipientId) | **Get** /coupon_recipients/{couponRecipientId} | Retrieve a coupon recipient
[**PATCHCouponRecipientsCouponRecipientId**](CouponRecipientsApi.md#PATCHCouponRecipientsCouponRecipientId) | **Patch** /coupon_recipients/{couponRecipientId} | Update a coupon recipient
[**POSTCouponRecipients**](CouponRecipientsApi.md#POSTCouponRecipients) | **Post** /coupon_recipients | Create a coupon recipient



## DELETECouponRecipientsCouponRecipientId

> DELETECouponRecipientsCouponRecipientId(ctx, couponRecipientId).Execute()

Delete a coupon recipient



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
    couponRecipientId := "couponRecipientId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponRecipientsApi.DELETECouponRecipientsCouponRecipientId(context.Background(), couponRecipientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponRecipientsApi.DELETECouponRecipientsCouponRecipientId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponRecipientId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETECouponRecipientsCouponRecipientIdRequest struct via the builder pattern


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


## GETCouponRecipients

> CouponRecipientResponseList GETCouponRecipients(ctx).Execute()

List all coupon recipients



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
    resp, r, err := apiClient.CouponRecipientsApi.GETCouponRecipients(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponRecipientsApi.GETCouponRecipients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCouponRecipients`: CouponRecipientResponseList
    fmt.Fprintf(os.Stdout, "Response from `CouponRecipientsApi.GETCouponRecipients`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETCouponRecipientsRequest struct via the builder pattern


### Return type

[**CouponRecipientResponseList**](CouponRecipientResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCouponRecipientsCouponRecipientId

> CouponRecipientResponse GETCouponRecipientsCouponRecipientId(ctx, couponRecipientId).Execute()

Retrieve a coupon recipient



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
    couponRecipientId := "couponRecipientId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponRecipientsApi.GETCouponRecipientsCouponRecipientId(context.Background(), couponRecipientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponRecipientsApi.GETCouponRecipientsCouponRecipientId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCouponRecipientsCouponRecipientId`: CouponRecipientResponse
    fmt.Fprintf(os.Stdout, "Response from `CouponRecipientsApi.GETCouponRecipientsCouponRecipientId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponRecipientId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCouponRecipientsCouponRecipientIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CouponRecipientResponse**](CouponRecipientResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHCouponRecipientsCouponRecipientId

> CouponRecipientResponse PATCHCouponRecipientsCouponRecipientId(ctx, couponRecipientId).CouponRecipientUpdate(couponRecipientUpdate).Execute()

Update a coupon recipient



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
    couponRecipientUpdate := *openapiclient.NewCouponRecipientUpdate(*openapiclient.NewCouponRecipientUpdateData("Type_example", "XGZwpOSrWL", *openapiclient.NewCouponRecipientUpdateDataAttributes())) // CouponRecipientUpdate | 
    couponRecipientId := "couponRecipientId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponRecipientsApi.PATCHCouponRecipientsCouponRecipientId(context.Background(), couponRecipientId).CouponRecipientUpdate(couponRecipientUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponRecipientsApi.PATCHCouponRecipientsCouponRecipientId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHCouponRecipientsCouponRecipientId`: CouponRecipientResponse
    fmt.Fprintf(os.Stdout, "Response from `CouponRecipientsApi.PATCHCouponRecipientsCouponRecipientId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponRecipientId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHCouponRecipientsCouponRecipientIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **couponRecipientUpdate** | [**CouponRecipientUpdate**](CouponRecipientUpdate.md) |  | 


### Return type

[**CouponRecipientResponse**](CouponRecipientResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTCouponRecipients

> CouponRecipientResponse POSTCouponRecipients(ctx).CouponRecipientCreate(couponRecipientCreate).Execute()

Create a coupon recipient



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
    couponRecipientCreate := *openapiclient.NewCouponRecipientCreate(*openapiclient.NewCouponRecipientCreateData("Type_example", *openapiclient.NewCouponRecipientCreateDataAttributes("john@example.com"))) // CouponRecipientCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponRecipientsApi.POSTCouponRecipients(context.Background()).CouponRecipientCreate(couponRecipientCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponRecipientsApi.POSTCouponRecipients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTCouponRecipients`: CouponRecipientResponse
    fmt.Fprintf(os.Stdout, "Response from `CouponRecipientsApi.POSTCouponRecipients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTCouponRecipientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **couponRecipientCreate** | [**CouponRecipientCreate**](CouponRecipientCreate.md) |  | 

### Return type

[**CouponRecipientResponse**](CouponRecipientResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

