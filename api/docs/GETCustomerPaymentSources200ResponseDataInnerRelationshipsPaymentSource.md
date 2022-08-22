# GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "adyen_payments"]
**Id** | **string** | The resource&#39;s id | 
**Data** | [**WireTransferData**](WireTransferData.md) |  | 

## Methods

### NewGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource

`func NewGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource(type_ string, id string, data WireTransferData, ) *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource`

NewGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource instantiates a new GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceWithDefaults

`func NewGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceWithDefaults() *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource`

NewGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceWithDefaults instantiates a new GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) SetId(v string)`

SetId sets Id field to given value.


### GetData

`func (o *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) GetData() WireTransferData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) GetDataOk() (*WireTransferData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) SetData(v WireTransferData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


