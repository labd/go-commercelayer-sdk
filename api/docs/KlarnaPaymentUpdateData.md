# KlarnaPaymentUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes**](PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**AdyenPaymentUpdateDataRelationships**](AdyenPaymentUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewKlarnaPaymentUpdateData

`func NewKlarnaPaymentUpdateData(type_ string, id string, attributes PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes, ) *KlarnaPaymentUpdateData`

NewKlarnaPaymentUpdateData instantiates a new KlarnaPaymentUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKlarnaPaymentUpdateDataWithDefaults

`func NewKlarnaPaymentUpdateDataWithDefaults() *KlarnaPaymentUpdateData`

NewKlarnaPaymentUpdateDataWithDefaults instantiates a new KlarnaPaymentUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *KlarnaPaymentUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KlarnaPaymentUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KlarnaPaymentUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *KlarnaPaymentUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KlarnaPaymentUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KlarnaPaymentUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *KlarnaPaymentUpdateData) GetAttributes() PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *KlarnaPaymentUpdateData) GetAttributesOk() (*PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *KlarnaPaymentUpdateData) SetAttributes(v PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *KlarnaPaymentUpdateData) GetRelationships() AdyenPaymentUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *KlarnaPaymentUpdateData) GetRelationshipsOk() (*AdyenPaymentUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *KlarnaPaymentUpdateData) SetRelationships(v AdyenPaymentUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *KlarnaPaymentUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


