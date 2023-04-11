# \CouponsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETECouponsCouponId**](CouponsApi.md#DELETECouponsCouponId) | **Delete** /coupons/{couponId} | Delete a coupon
[**GETCouponCodesPromotionRuleIdCoupons**](CouponsApi.md#GETCouponCodesPromotionRuleIdCoupons) | **Get** /coupon_codes_promotion_rules/{couponCodesPromotionRuleId}/coupons | Retrieve the coupons associated to the coupon codes promotion rule
[**GETCoupons**](CouponsApi.md#GETCoupons) | **Get** /coupons | List all coupons
[**GETCouponsCouponId**](CouponsApi.md#GETCouponsCouponId) | **Get** /coupons/{couponId} | Retrieve a coupon
[**PATCHCouponsCouponId**](CouponsApi.md#PATCHCouponsCouponId) | **Patch** /coupons/{couponId} | Update a coupon
[**POSTCoupons**](CouponsApi.md#POSTCoupons) | **Post** /coupons | Create a coupon



## DELETECouponsCouponId

> DELETECouponsCouponId(ctx, couponId).Execute()

Delete a coupon



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
    couponId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponsApi.DELETECouponsCouponId(context.Background(), couponId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponsApi.DELETECouponsCouponId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETECouponsCouponIdRequest struct via the builder pattern


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


## GETCouponCodesPromotionRuleIdCoupons

> GETCouponCodesPromotionRuleIdCoupons(ctx, couponCodesPromotionRuleId).Execute()

Retrieve the coupons associated to the coupon codes promotion rule



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
    couponCodesPromotionRuleId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponsApi.GETCouponCodesPromotionRuleIdCoupons(context.Background(), couponCodesPromotionRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponsApi.GETCouponCodesPromotionRuleIdCoupons``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponCodesPromotionRuleId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCouponCodesPromotionRuleIdCouponsRequest struct via the builder pattern


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


## GETCoupons

> GETCoupons200Response GETCoupons(ctx).Execute()

List all coupons



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
    resp, r, err := apiClient.CouponsApi.GETCoupons(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponsApi.GETCoupons``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCoupons`: GETCoupons200Response
    fmt.Fprintf(os.Stdout, "Response from `CouponsApi.GETCoupons`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETCouponsRequest struct via the builder pattern


### Return type

[**GETCoupons200Response**](GETCoupons200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCouponsCouponId

> GETCouponsCouponId200Response GETCouponsCouponId(ctx, couponId).Execute()

Retrieve a coupon



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
    couponId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponsApi.GETCouponsCouponId(context.Background(), couponId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponsApi.GETCouponsCouponId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCouponsCouponId`: GETCouponsCouponId200Response
    fmt.Fprintf(os.Stdout, "Response from `CouponsApi.GETCouponsCouponId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCouponsCouponIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETCouponsCouponId200Response**](GETCouponsCouponId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHCouponsCouponId

> PATCHCouponsCouponId200Response PATCHCouponsCouponId(ctx, couponId).CouponUpdate(couponUpdate).Execute()

Update a coupon



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
    couponUpdate := *openapiclient.NewCouponUpdate(*openapiclient.NewCouponUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHCouponsCouponId200ResponseDataAttributes())) // CouponUpdate | 
    couponId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponsApi.PATCHCouponsCouponId(context.Background(), couponId).CouponUpdate(couponUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponsApi.PATCHCouponsCouponId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHCouponsCouponId`: PATCHCouponsCouponId200Response
    fmt.Fprintf(os.Stdout, "Response from `CouponsApi.PATCHCouponsCouponId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHCouponsCouponIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **couponUpdate** | [**CouponUpdate**](CouponUpdate.md) |  | 


### Return type

[**PATCHCouponsCouponId200Response**](PATCHCouponsCouponId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTCoupons

> POSTCoupons201Response POSTCoupons(ctx).CouponCreate(couponCreate).Execute()

Create a coupon



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
    couponCreate := *openapiclient.NewCouponCreate(*openapiclient.NewCouponCreateData(interface{}(123), *openapiclient.NewPOSTCoupons201ResponseDataAttributes(interface{}(04371af2-70b3-48d7-8f4e-316b374224c3), interface{}(50)))) // CouponCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponsApi.POSTCoupons(context.Background()).CouponCreate(couponCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponsApi.POSTCoupons``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTCoupons`: POSTCoupons201Response
    fmt.Fprintf(os.Stdout, "Response from `CouponsApi.POSTCoupons`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTCouponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **couponCreate** | [**CouponCreate**](CouponCreate.md) |  | 

### Return type

[**POSTCoupons201Response**](POSTCoupons201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

