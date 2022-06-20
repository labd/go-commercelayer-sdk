# CustomerPaymentSourceDataRelationshipsPaymentSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "adyen_payments"]
**Id** | **string** | The resource&#39;s id | 
**Data** | [**WireTransferData**](WireTransferData.md) |  | 

## Methods

### NewCustomerPaymentSourceDataRelationshipsPaymentSource

`func NewCustomerPaymentSourceDataRelationshipsPaymentSource(type_ string, id string, data WireTransferData, ) *CustomerPaymentSourceDataRelationshipsPaymentSource`

NewCustomerPaymentSourceDataRelationshipsPaymentSource instantiates a new CustomerPaymentSourceDataRelationshipsPaymentSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerPaymentSourceDataRelationshipsPaymentSourceWithDefaults

`func NewCustomerPaymentSourceDataRelationshipsPaymentSourceWithDefaults() *CustomerPaymentSourceDataRelationshipsPaymentSource`

NewCustomerPaymentSourceDataRelationshipsPaymentSourceWithDefaults instantiates a new CustomerPaymentSourceDataRelationshipsPaymentSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomerPaymentSourceDataRelationshipsPaymentSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerPaymentSourceDataRelationshipsPaymentSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerPaymentSourceDataRelationshipsPaymentSource) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *CustomerPaymentSourceDataRelationshipsPaymentSource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerPaymentSourceDataRelationshipsPaymentSource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerPaymentSourceDataRelationshipsPaymentSource) SetId(v string)`

SetId sets Id field to given value.


### GetData

`func (o *CustomerPaymentSourceDataRelationshipsPaymentSource) GetData() WireTransferData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CustomerPaymentSourceDataRelationshipsPaymentSource) GetDataOk() (*WireTransferData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CustomerPaymentSourceDataRelationshipsPaymentSource) SetData(v WireTransferData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


