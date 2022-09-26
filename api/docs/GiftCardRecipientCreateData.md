# GiftCardRecipientCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**CouponRecipientCreateDataAttributes**](CouponRecipientCreateDataAttributes.md) |  | 
**Relationships** | Pointer to [**CouponRecipientCreateDataRelationships**](CouponRecipientCreateDataRelationships.md) |  | [optional] 

## Methods

### NewGiftCardRecipientCreateData

`func NewGiftCardRecipientCreateData(type_ string, attributes CouponRecipientCreateDataAttributes, ) *GiftCardRecipientCreateData`

NewGiftCardRecipientCreateData instantiates a new GiftCardRecipientCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardRecipientCreateDataWithDefaults

`func NewGiftCardRecipientCreateDataWithDefaults() *GiftCardRecipientCreateData`

NewGiftCardRecipientCreateDataWithDefaults instantiates a new GiftCardRecipientCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GiftCardRecipientCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GiftCardRecipientCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GiftCardRecipientCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *GiftCardRecipientCreateData) GetAttributes() CouponRecipientCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GiftCardRecipientCreateData) GetAttributesOk() (*CouponRecipientCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GiftCardRecipientCreateData) SetAttributes(v CouponRecipientCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *GiftCardRecipientCreateData) GetRelationships() CouponRecipientCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GiftCardRecipientCreateData) GetRelationshipsOk() (*CouponRecipientCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GiftCardRecipientCreateData) SetRelationships(v CouponRecipientCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GiftCardRecipientCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


