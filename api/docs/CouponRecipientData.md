# CouponRecipientData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETCouponRecipientsCouponRecipientId200ResponseDataAttributes**](GETCouponRecipientsCouponRecipientId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**CouponRecipientDataRelationships**](CouponRecipientDataRelationships.md) |  | [optional] 

## Methods

### NewCouponRecipientData

`func NewCouponRecipientData(type_ interface{}, attributes GETCouponRecipientsCouponRecipientId200ResponseDataAttributes, ) *CouponRecipientData`

NewCouponRecipientData instantiates a new CouponRecipientData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponRecipientDataWithDefaults

`func NewCouponRecipientDataWithDefaults() *CouponRecipientData`

NewCouponRecipientDataWithDefaults instantiates a new CouponRecipientData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CouponRecipientData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CouponRecipientData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CouponRecipientData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CouponRecipientData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CouponRecipientData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *CouponRecipientData) GetAttributes() GETCouponRecipientsCouponRecipientId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CouponRecipientData) GetAttributesOk() (*GETCouponRecipientsCouponRecipientId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CouponRecipientData) SetAttributes(v GETCouponRecipientsCouponRecipientId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CouponRecipientData) GetRelationships() CouponRecipientDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CouponRecipientData) GetRelationshipsOk() (*CouponRecipientDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CouponRecipientData) SetRelationships(v CouponRecipientDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CouponRecipientData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


