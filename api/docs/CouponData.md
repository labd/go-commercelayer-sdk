# CouponData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "coupons"]
**Attributes** | [**GETCoupons200ResponseDataInnerAttributes**](GETCoupons200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**GETCoupons200ResponseDataInnerRelationships**](GETCoupons200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewCouponData

`func NewCouponData(type_ string, attributes GETCoupons200ResponseDataInnerAttributes, ) *CouponData`

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

`func (o *CouponData) GetAttributes() GETCoupons200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CouponData) GetAttributesOk() (*GETCoupons200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CouponData) SetAttributes(v GETCoupons200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CouponData) GetRelationships() GETCoupons200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CouponData) GetRelationshipsOk() (*GETCoupons200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CouponData) SetRelationships(v GETCoupons200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CouponData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


