# CouponData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**CouponDataAttributes**](CouponDataAttributes.md) |  | 
**Relationships** | Pointer to [**CouponDataRelationships**](CouponDataRelationships.md) |  | [optional] 

## Methods

### NewCouponData

`func NewCouponData(type_ string, attributes CouponDataAttributes, ) *CouponData`

NewCouponData instantiates a new CouponData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponDataWithDefaults

`func NewCouponDataWithDefaults() *CouponData`

NewCouponDataWithDefaults instantiates a new CouponData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CouponData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CouponData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CouponData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *CouponData) GetAttributes() CouponDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CouponData) GetAttributesOk() (*CouponDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CouponData) SetAttributes(v CouponDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CouponData) GetRelationships() CouponDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CouponData) GetRelationshipsOk() (*CouponDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CouponData) SetRelationships(v CouponDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CouponData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


