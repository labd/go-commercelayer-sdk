# CouponRecipientCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTCouponRecipients201ResponseDataAttributes**](POSTCouponRecipients201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**CouponRecipientCreateDataRelationships**](CouponRecipientCreateDataRelationships.md) |  | [optional] 

## Methods

### NewCouponRecipientCreateData

`func NewCouponRecipientCreateData(type_ interface{}, attributes POSTCouponRecipients201ResponseDataAttributes, ) *CouponRecipientCreateData`

NewCouponRecipientCreateData instantiates a new CouponRecipientCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponRecipientCreateDataWithDefaults

`func NewCouponRecipientCreateDataWithDefaults() *CouponRecipientCreateData`

NewCouponRecipientCreateDataWithDefaults instantiates a new CouponRecipientCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CouponRecipientCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CouponRecipientCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CouponRecipientCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CouponRecipientCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CouponRecipientCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *CouponRecipientCreateData) GetAttributes() POSTCouponRecipients201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CouponRecipientCreateData) GetAttributesOk() (*POSTCouponRecipients201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CouponRecipientCreateData) SetAttributes(v POSTCouponRecipients201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CouponRecipientCreateData) GetRelationships() CouponRecipientCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CouponRecipientCreateData) GetRelationshipsOk() (*CouponRecipientCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CouponRecipientCreateData) SetRelationships(v CouponRecipientCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CouponRecipientCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


