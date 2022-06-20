# PaymentMethodUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "payment_methods"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**PaymentMethodUpdateDataAttributes**](PaymentMethodUpdateDataAttributes.md) |  | 
**Relationships** | Pointer to [**PaymentMethodUpdateDataRelationships**](PaymentMethodUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewPaymentMethodUpdateData

`func NewPaymentMethodUpdateData(type_ string, id string, attributes PaymentMethodUpdateDataAttributes, ) *PaymentMethodUpdateData`

NewPaymentMethodUpdateData instantiates a new PaymentMethodUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodUpdateDataWithDefaults

`func NewPaymentMethodUpdateDataWithDefaults() *PaymentMethodUpdateData`

NewPaymentMethodUpdateDataWithDefaults instantiates a new PaymentMethodUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentMethodUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *PaymentMethodUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethodUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethodUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *PaymentMethodUpdateData) GetAttributes() PaymentMethodUpdateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PaymentMethodUpdateData) GetAttributesOk() (*PaymentMethodUpdateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PaymentMethodUpdateData) SetAttributes(v PaymentMethodUpdateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PaymentMethodUpdateData) GetRelationships() PaymentMethodUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PaymentMethodUpdateData) GetRelationshipsOk() (*PaymentMethodUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PaymentMethodUpdateData) SetRelationships(v PaymentMethodUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PaymentMethodUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


