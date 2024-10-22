# ShipmentUpdateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShippingCategory** | Pointer to [**ShipmentCreateDataRelationshipsShippingCategory**](ShipmentCreateDataRelationshipsShippingCategory.md) |  | [optional] 
**InventoryStockLocation** | Pointer to [**ShipmentCreateDataRelationshipsInventoryStockLocation**](ShipmentCreateDataRelationshipsInventoryStockLocation.md) |  | [optional] 
**ShippingAddress** | Pointer to [**CustomerAddressCreateDataRelationshipsAddress**](CustomerAddressCreateDataRelationshipsAddress.md) |  | [optional] 
**ShippingMethod** | Pointer to [**DeliveryLeadTimeCreateDataRelationshipsShippingMethod**](DeliveryLeadTimeCreateDataRelationshipsShippingMethod.md) |  | [optional] 
**Tags** | Pointer to [**AddressCreateDataRelationshipsTags**](AddressCreateDataRelationshipsTags.md) |  | [optional] 

## Methods

### NewShipmentUpdateDataRelationships

`func NewShipmentUpdateDataRelationships() *ShipmentUpdateDataRelationships`

NewShipmentUpdateDataRelationships instantiates a new ShipmentUpdateDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentUpdateDataRelationshipsWithDefaults

`func NewShipmentUpdateDataRelationshipsWithDefaults() *ShipmentUpdateDataRelationships`

NewShipmentUpdateDataRelationshipsWithDefaults instantiates a new ShipmentUpdateDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShippingCategory

`func (o *ShipmentUpdateDataRelationships) GetShippingCategory() ShipmentCreateDataRelationshipsShippingCategory`

GetShippingCategory returns the ShippingCategory field if non-nil, zero value otherwise.

### GetShippingCategoryOk

`func (o *ShipmentUpdateDataRelationships) GetShippingCategoryOk() (*ShipmentCreateDataRelationshipsShippingCategory, bool)`

GetShippingCategoryOk returns a tuple with the ShippingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCategory

`func (o *ShipmentUpdateDataRelationships) SetShippingCategory(v ShipmentCreateDataRelationshipsShippingCategory)`

SetShippingCategory sets ShippingCategory field to given value.

### HasShippingCategory

`func (o *ShipmentUpdateDataRelationships) HasShippingCategory() bool`

HasShippingCategory returns a boolean if a field has been set.

### GetInventoryStockLocation

`func (o *ShipmentUpdateDataRelationships) GetInventoryStockLocation() ShipmentCreateDataRelationshipsInventoryStockLocation`

GetInventoryStockLocation returns the InventoryStockLocation field if non-nil, zero value otherwise.

### GetInventoryStockLocationOk

`func (o *ShipmentUpdateDataRelationships) GetInventoryStockLocationOk() (*ShipmentCreateDataRelationshipsInventoryStockLocation, bool)`

GetInventoryStockLocationOk returns a tuple with the InventoryStockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryStockLocation

`func (o *ShipmentUpdateDataRelationships) SetInventoryStockLocation(v ShipmentCreateDataRelationshipsInventoryStockLocation)`

SetInventoryStockLocation sets InventoryStockLocation field to given value.

### HasInventoryStockLocation

`func (o *ShipmentUpdateDataRelationships) HasInventoryStockLocation() bool`

HasInventoryStockLocation returns a boolean if a field has been set.

### GetShippingAddress

`func (o *ShipmentUpdateDataRelationships) GetShippingAddress() CustomerAddressCreateDataRelationshipsAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *ShipmentUpdateDataRelationships) GetShippingAddressOk() (*CustomerAddressCreateDataRelationshipsAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *ShipmentUpdateDataRelationships) SetShippingAddress(v CustomerAddressCreateDataRelationshipsAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *ShipmentUpdateDataRelationships) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetShippingMethod

`func (o *ShipmentUpdateDataRelationships) GetShippingMethod() DeliveryLeadTimeCreateDataRelationshipsShippingMethod`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *ShipmentUpdateDataRelationships) GetShippingMethodOk() (*DeliveryLeadTimeCreateDataRelationshipsShippingMethod, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *ShipmentUpdateDataRelationships) SetShippingMethod(v DeliveryLeadTimeCreateDataRelationshipsShippingMethod)`

SetShippingMethod sets ShippingMethod field to given value.

### HasShippingMethod

`func (o *ShipmentUpdateDataRelationships) HasShippingMethod() bool`

HasShippingMethod returns a boolean if a field has been set.

### GetTags

`func (o *ShipmentUpdateDataRelationships) GetTags() AddressCreateDataRelationshipsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ShipmentUpdateDataRelationships) GetTagsOk() (*AddressCreateDataRelationshipsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ShipmentUpdateDataRelationships) SetTags(v AddressCreateDataRelationshipsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ShipmentUpdateDataRelationships) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


