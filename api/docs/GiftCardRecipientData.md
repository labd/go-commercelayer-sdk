# GiftCardRecipientData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETCouponRecipientsCouponRecipientId200ResponseDataAttributes**](GETCouponRecipientsCouponRecipientId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**CouponRecipientDataRelationships**](CouponRecipientDataRelationships.md) |  | [optional] 

## Methods

### NewGiftCardRecipientData

`func NewGiftCardRecipientData(type_ interface{}, attributes GETCouponRecipientsCouponRecipientId200ResponseDataAttributes, ) *GiftCardRecipientData`

NewGiftCardRecipientData instantiates a new GiftCardRecipientData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardRecipientDataWithDefaults

`func NewGiftCardRecipientDataWithDefaults() *GiftCardRecipientData`

NewGiftCardRecipientDataWithDefaults instantiates a new GiftCardRecipientData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GiftCardRecipientData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GiftCardRecipientData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GiftCardRecipientData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *GiftCardRecipientData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *GiftCardRecipientData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *GiftCardRecipientData) GetAttributes() GETCouponRecipientsCouponRecipientId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GiftCardRecipientData) GetAttributesOk() (*GETCouponRecipientsCouponRecipientId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GiftCardRecipientData) SetAttributes(v GETCouponRecipientsCouponRecipientId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *GiftCardRecipientData) GetRelationships() CouponRecipientDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GiftCardRecipientData) GetRelationshipsOk() (*CouponRecipientDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GiftCardRecipientData) SetRelationships(v CouponRecipientDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GiftCardRecipientData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


