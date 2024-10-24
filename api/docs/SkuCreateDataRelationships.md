# SkuCreateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShippingCategory** | [**ShipmentCreateDataRelationshipsShippingCategory**](ShipmentCreateDataRelationshipsShippingCategory.md) |  | 
**Tags** | Pointer to [**AddressCreateDataRelationshipsTags**](AddressCreateDataRelationshipsTags.md) |  | [optional] 

## Methods

### NewSkuCreateDataRelationships

`func NewSkuCreateDataRelationships(shippingCategory ShipmentCreateDataRelationshipsShippingCategory, ) *SkuCreateDataRelationships`

NewSkuCreateDataRelationships instantiates a new SkuCreateDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuCreateDataRelationshipsWithDefaults

`func NewSkuCreateDataRelationshipsWithDefaults() *SkuCreateDataRelationships`

NewSkuCreateDataRelationshipsWithDefaults instantiates a new SkuCreateDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShippingCategory

`func (o *SkuCreateDataRelationships) GetShippingCategory() ShipmentCreateDataRelationshipsShippingCategory`

GetShippingCategory returns the ShippingCategory field if non-nil, zero value otherwise.

### GetShippingCategoryOk

`func (o *SkuCreateDataRelationships) GetShippingCategoryOk() (*ShipmentCreateDataRelationshipsShippingCategory, bool)`

GetShippingCategoryOk returns a tuple with the ShippingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCategory

`func (o *SkuCreateDataRelationships) SetShippingCategory(v ShipmentCreateDataRelationshipsShippingCategory)`

SetShippingCategory sets ShippingCategory field to given value.


### GetTags

`func (o *SkuCreateDataRelationships) GetTags() AddressCreateDataRelationshipsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SkuCreateDataRelationships) GetTagsOk() (*AddressCreateDataRelationshipsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SkuCreateDataRelationships) SetTags(v AddressCreateDataRelationshipsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SkuCreateDataRelationships) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


