# PaymentOptionUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes**](PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**AdyenPaymentUpdateDataRelationships**](AdyenPaymentUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewPaymentOptionUpdateData

`func NewPaymentOptionUpdateData(type_ interface{}, id interface{}, attributes PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes, ) *PaymentOptionUpdateData`

NewPaymentOptionUpdateData instantiates a new PaymentOptionUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentOptionUpdateDataWithDefaults

`func NewPaymentOptionUpdateDataWithDefaults() *PaymentOptionUpdateData`

NewPaymentOptionUpdateDataWithDefaults instantiates a new PaymentOptionUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentOptionUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentOptionUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentOptionUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *PaymentOptionUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PaymentOptionUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *PaymentOptionUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentOptionUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentOptionUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *PaymentOptionUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PaymentOptionUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *PaymentOptionUpdateData) GetAttributes() PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PaymentOptionUpdateData) GetAttributesOk() (*PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PaymentOptionUpdateData) SetAttributes(v PATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PaymentOptionUpdateData) GetRelationships() AdyenPaymentUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PaymentOptionUpdateData) GetRelationshipsOk() (*AdyenPaymentUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PaymentOptionUpdateData) SetRelationships(v AdyenPaymentUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PaymentOptionUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


