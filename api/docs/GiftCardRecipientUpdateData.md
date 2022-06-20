# GiftCardRecipientUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "gift_card_recipients"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**CouponRecipientUpdateDataAttributes**](CouponRecipientUpdateDataAttributes.md) |  | 
**Relationships** | Pointer to [**CouponRecipientCreateDataRelationships**](CouponRecipientCreateDataRelationships.md) |  | [optional] 

## Methods

### NewGiftCardRecipientUpdateData

`func NewGiftCardRecipientUpdateData(type_ string, id string, attributes CouponRecipientUpdateDataAttributes, ) *GiftCardRecipientUpdateData`

NewGiftCardRecipientUpdateData instantiates a new GiftCardRecipientUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardRecipientUpdateDataWithDefaults

`func NewGiftCardRecipientUpdateDataWithDefaults() *GiftCardRecipientUpdateData`

NewGiftCardRecipientUpdateDataWithDefaults instantiates a new GiftCardRecipientUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GiftCardRecipientUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GiftCardRecipientUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GiftCardRecipientUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *GiftCardRecipientUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GiftCardRecipientUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GiftCardRecipientUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GiftCardRecipientUpdateData) GetAttributes() CouponRecipientUpdateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GiftCardRecipientUpdateData) GetAttributesOk() (*CouponRecipientUpdateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GiftCardRecipientUpdateData) SetAttributes(v CouponRecipientUpdateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *GiftCardRecipientUpdateData) GetRelationships() CouponRecipientCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GiftCardRecipientUpdateData) GetRelationshipsOk() (*CouponRecipientCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GiftCardRecipientUpdateData) SetRelationships(v CouponRecipientCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GiftCardRecipientUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


