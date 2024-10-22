# PaymentOptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETPaymentOptionsPaymentOptionId200ResponseDataAttributes**](GETPaymentOptionsPaymentOptionId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**PaymentOptionDataRelationships**](PaymentOptionDataRelationships.md) |  | [optional] 

## Methods

### NewPaymentOptionData

`func NewPaymentOptionData(type_ interface{}, attributes GETPaymentOptionsPaymentOptionId200ResponseDataAttributes, ) *PaymentOptionData`

NewPaymentOptionData instantiates a new PaymentOptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentOptionDataWithDefaults

`func NewPaymentOptionDataWithDefaults() *PaymentOptionData`

NewPaymentOptionDataWithDefaults instantiates a new PaymentOptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentOptionData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentOptionData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentOptionData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *PaymentOptionData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PaymentOptionData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *PaymentOptionData) GetAttributes() GETPaymentOptionsPaymentOptionId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PaymentOptionData) GetAttributesOk() (*GETPaymentOptionsPaymentOptionId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PaymentOptionData) SetAttributes(v GETPaymentOptionsPaymentOptionId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PaymentOptionData) GetRelationships() PaymentOptionDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PaymentOptionData) GetRelationshipsOk() (*PaymentOptionDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PaymentOptionData) SetRelationships(v PaymentOptionDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PaymentOptionData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


