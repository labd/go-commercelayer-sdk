# POSTCouponRecipientsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTCouponRecipientsRequestDataAttributes**](POSTCouponRecipientsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTCouponRecipientsRequestDataRelationships**](POSTCouponRecipientsRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTCouponRecipientsRequestData

`func NewPOSTCouponRecipientsRequestData(type_ interface{}, attributes POSTCouponRecipientsRequestDataAttributes, ) *POSTCouponRecipientsRequestData`

NewPOSTCouponRecipientsRequestData instantiates a new POSTCouponRecipientsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTCouponRecipientsRequestDataWithDefaults

`func NewPOSTCouponRecipientsRequestDataWithDefaults() *POSTCouponRecipientsRequestData`

NewPOSTCouponRecipientsRequestDataWithDefaults instantiates a new POSTCouponRecipientsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTCouponRecipientsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTCouponRecipientsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTCouponRecipientsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTCouponRecipientsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTCouponRecipientsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTCouponRecipientsRequestData) GetAttributes() POSTCouponRecipientsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTCouponRecipientsRequestData) GetAttributesOk() (*POSTCouponRecipientsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTCouponRecipientsRequestData) SetAttributes(v POSTCouponRecipientsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTCouponRecipientsRequestData) GetRelationships() POSTCouponRecipientsRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTCouponRecipientsRequestData) GetRelationshipsOk() (*POSTCouponRecipientsRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTCouponRecipientsRequestData) SetRelationships(v POSTCouponRecipientsRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTCouponRecipientsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


