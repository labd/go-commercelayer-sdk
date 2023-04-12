# \CarrierAccountsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETCarrierAccounts**](CarrierAccountsApi.md#GETCarrierAccounts) | **Get** /carrier_accounts | List all carrier accounts
[**GETCarrierAccountsCarrierAccountId**](CarrierAccountsApi.md#GETCarrierAccountsCarrierAccountId) | **Get** /carrier_accounts/{carrierAccountId} | Retrieve a carrier account
[**GETShipmentIdCarrierAccounts**](CarrierAccountsApi.md#GETShipmentIdCarrierAccounts) | **Get** /shipments/{shipmentId}/carrier_accounts | Retrieve the carrier accounts associated to the shipment



## GETCarrierAccounts

> GETCarrierAccounts200Response GETCarrierAccounts(ctx).Execute()

List all carrier accounts



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
    resp, r, err := apiClient.CarrierAccountsApi.GETCarrierAccounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CarrierAccountsApi.GETCarrierAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCarrierAccounts`: GETCarrierAccounts200Response
    fmt.Fprintf(os.Stdout, "Response from `CarrierAccountsApi.GETCarrierAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETCarrierAccountsRequest struct via the builder pattern


### Return type

[**GETCarrierAccounts200Response**](GETCarrierAccounts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCarrierAccountsCarrierAccountId

> GETCarrierAccountsCarrierAccountId200Response GETCarrierAccountsCarrierAccountId(ctx, carrierAccountId).Execute()

Retrieve a carrier account



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
    carrierAccountId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CarrierAccountsApi.GETCarrierAccountsCarrierAccountId(context.Background(), carrierAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CarrierAccountsApi.GETCarrierAccountsCarrierAccountId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCarrierAccountsCarrierAccountId`: GETCarrierAccountsCarrierAccountId200Response
    fmt.Fprintf(os.Stdout, "Response from `CarrierAccountsApi.GETCarrierAccountsCarrierAccountId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**carrierAccountId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCarrierAccountsCarrierAccountIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETCarrierAccountsCarrierAccountId200Response**](GETCarrierAccountsCarrierAccountId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETShipmentIdCarrierAccounts

> GETShipmentIdCarrierAccounts(ctx, shipmentId).Execute()

Retrieve the carrier accounts associated to the shipment



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
    shipmentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CarrierAccountsApi.GETShipmentIdCarrierAccounts(context.Background(), shipmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CarrierAccountsApi.GETShipmentIdCarrierAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETShipmentIdCarrierAccountsRequest struct via the builder pattern


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

