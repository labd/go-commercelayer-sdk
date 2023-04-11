# POSTGiftCardRecipientsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTCouponRecipientsRequestDataAttributes**](POSTCouponRecipientsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTCouponRecipientsRequestDataRelationships**](POSTCouponRecipientsRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTGiftCardRecipientsRequestData

`func NewPOSTGiftCardRecipientsRequestData(type_ interface{}, attributes POSTCouponRecipientsRequestDataAttributes, ) *POSTGiftCardRecipientsRequestData`

NewPOSTGiftCardRecipientsRequestData instantiates a new POSTGiftCardRecipientsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTGiftCardRecipientsRequestDataWithDefaults

`func NewPOSTGiftCardRecipientsRequestDataWithDefaults() *POSTGiftCardRecipientsRequestData`

NewPOSTGiftCardRecipientsRequestDataWithDefaults instantiates a new POSTGiftCardRecipientsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTGiftCardRecipientsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTGiftCardRecipientsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTGiftCardRecipientsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTGiftCardRecipientsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTGiftCardRecipientsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTGiftCardRecipientsRequestData) GetAttributes() POSTCouponRecipientsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTGiftCardRecipientsRequestData) GetAttributesOk() (*POSTCouponRecipientsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTGiftCardRecipientsRequestData) SetAttributes(v POSTCouponRecipientsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTGiftCardRecipientsRequestData) GetRelationships() POSTCouponRecipientsRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTGiftCardRecipientsRequestData) GetRelationshipsOk() (*POSTCouponRecipientsRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTGiftCardRecipientsRequestData) SetRelationships(v POSTCouponRecipientsRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTGiftCardRecipientsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


