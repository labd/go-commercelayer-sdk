# CouponUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**PATCHCouponsCouponId200ResponseDataAttributes**](PATCHCouponsCouponId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**CouponUpdateDataRelationships**](CouponUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewCouponUpdateData

`func NewCouponUpdateData(type_ string, id string, attributes PATCHCouponsCouponId200ResponseDataAttributes, ) *CouponUpdateData`

NewCouponUpdateData instantiates a new CouponUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponUpdateDataWithDefaults

`func NewCouponUpdateDataWithDefaults() *CouponUpdateData`

NewCouponUpdateDataWithDefaults instantiates a new CouponUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CouponUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CouponUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CouponUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *CouponUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CouponUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CouponUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CouponUpdateData) GetAttributes() PATCHCouponsCouponId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CouponUpdateData) GetAttributesOk() (*PATCHCouponsCouponId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CouponUpdateData) SetAttributes(v PATCHCouponsCouponId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CouponUpdateData) GetRelationships() CouponUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CouponUpdateData) GetRelationshipsOk() (*CouponUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CouponUpdateData) SetRelationships(v CouponUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CouponUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


