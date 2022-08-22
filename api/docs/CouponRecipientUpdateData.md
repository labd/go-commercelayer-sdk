# CouponRecipientUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "coupon_recipients"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes**](PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTCouponRecipients201ResponseDataRelationships**](POSTCouponRecipients201ResponseDataRelationships.md) |  | [optional] 

## Methods

### NewCouponRecipientUpdateData

`func NewCouponRecipientUpdateData(type_ string, id string, attributes PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes, ) *CouponRecipientUpdateData`

NewCouponRecipientUpdateData instantiates a new CouponRecipientUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponRecipientUpdateDataWithDefaults

`func NewCouponRecipientUpdateDataWithDefaults() *CouponRecipientUpdateData`

NewCouponRecipientUpdateDataWithDefaults instantiates a new CouponRecipientUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CouponRecipientUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CouponRecipientUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CouponRecipientUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *CouponRecipientUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CouponRecipientUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CouponRecipientUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CouponRecipientUpdateData) GetAttributes() PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CouponRecipientUpdateData) GetAttributesOk() (*PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CouponRecipientUpdateData) SetAttributes(v PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CouponRecipientUpdateData) GetRelationships() POSTCouponRecipients201ResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CouponRecipientUpdateData) GetRelationshipsOk() (*POSTCouponRecipients201ResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CouponRecipientUpdateData) SetRelationships(v POSTCouponRecipients201ResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CouponRecipientUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


