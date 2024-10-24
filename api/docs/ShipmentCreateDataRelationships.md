# ShipmentCreateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | [**AdyenPaymentCreateDataRelationshipsOrder**](AdyenPaymentCreateDataRelationshipsOrder.md) |  | 
**ShippingCategory** | Pointer to [**ShipmentCreateDataRelationshipsShippingCategory**](ShipmentCreateDataRelationshipsShippingCategory.md) |  | [optional] 
**InventoryStockLocation** | [**ShipmentCreateDataRelationshipsInventoryStockLocation**](ShipmentCreateDataRelationshipsInventoryStockLocation.md) |  | 
**ShippingAddress** | Pointer to [**CustomerAddressCreateDataRelationshipsAddress**](CustomerAddressCreateDataRelationshipsAddress.md) |  | [optional] 
**ShippingMethod** | Pointer to [**DeliveryLeadTimeCreateDataRelationshipsShippingMethod**](DeliveryLeadTimeCreateDataRelationshipsShippingMethod.md) |  | [optional] 
**Tags** | Pointer to [**AddressCreateDataRelationshipsTags**](AddressCreateDataRelationshipsTags.md) |  | [optional] 

## Methods

### NewShipmentCreateDataRelationships

`func NewShipmentCreateDataRelationships(order AdyenPaymentCreateDataRelationshipsOrder, inventoryStockLocation ShipmentCreateDataRelationshipsInventoryStockLocation, ) *ShipmentCreateDataRelationships`

NewShipmentCreateDataRelationships instantiates a new ShipmentCreateDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentCreateDataRelationshipsWithDefaults

`func NewShipmentCreateDataRelationshipsWithDefaults() *ShipmentCreateDataRelationships`

NewShipmentCreateDataRelationshipsWithDefaults instantiates a new ShipmentCreateDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *ShipmentCreateDataRelationships) GetOrder() AdyenPaymentCreateDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ShipmentCreateDataRelationships) GetOrderOk() (*AdyenPaymentCreateDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ShipmentCreateDataRelationships) SetOrder(v AdyenPaymentCreateDataRelationshipsOrder)`

SetOrder sets Order field to given value.


### GetShippingCategory

`func (o *ShipmentCreateDataRelationships) GetShippingCategory() ShipmentCreateDataRelationshipsShippingCategory`

GetShippingCategory returns the ShippingCategory field if non-nil, zero value otherwise.

### GetShippingCategoryOk

`func (o *ShipmentCreateDataRelationships) GetShippingCategoryOk() (*ShipmentCreateDataRelationshipsShippingCategory, bool)`

GetShippingCategoryOk returns a tuple with the ShippingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCategory

`func (o *ShipmentCreateDataRelationships) SetShippingCategory(v ShipmentCreateDataRelationshipsShippingCategory)`

SetShippingCategory sets ShippingCategory field to given value.

### HasShippingCategory

`func (o *ShipmentCreateDataRelationships) HasShippingCategory() bool`

HasShippingCategory returns a boolean if a field has been set.

### GetInventoryStockLocation

`func (o *ShipmentCreateDataRelationships) GetInventoryStockLocation() ShipmentCreateDataRelationshipsInventoryStockLocation`

GetInventoryStockLocation returns the InventoryStockLocation field if non-nil, zero value otherwise.

### GetInventoryStockLocationOk

`func (o *ShipmentCreateDataRelationships) GetInventoryStockLocationOk() (*ShipmentCreateDataRelationshipsInventoryStockLocation, bool)`

GetInventoryStockLocationOk returns a tuple with the InventoryStockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryStockLocation

`func (o *ShipmentCreateDataRelationships) SetInventoryStockLocation(v ShipmentCreateDataRelationshipsInventoryStockLocation)`

SetInventoryStockLocation sets InventoryStockLocation field to given value.


### GetShippingAddress

`func (o *ShipmentCreateDataRelationships) GetShippingAddress() CustomerAddressCreateDataRelationshipsAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *ShipmentCreateDataRelationships) GetShippingAddressOk() (*CustomerAddressCreateDataRelationshipsAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *ShipmentCreateDataRelationships) SetShippingAddress(v CustomerAddressCreateDataRelationshipsAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *ShipmentCreateDataRelationships) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetShippingMethod

`func (o *ShipmentCreateDataRelationships) GetShippingMethod() DeliveryLeadTimeCreateDataRelationshipsShippingMethod`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *ShipmentCreateDataRelationships) GetShippingMethodOk() (*DeliveryLeadTimeCreateDataRelationshipsShippingMethod, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *ShipmentCreateDataRelationships) SetShippingMethod(v DeliveryLeadTimeCreateDataRelationshipsShippingMethod)`

SetShippingMethod sets ShippingMethod field to given value.

### HasShippingMethod

`func (o *ShipmentCreateDataRelationships) HasShippingMethod() bool`

HasShippingMethod returns a boolean if a field has been set.

### GetTags

`func (o *ShipmentCreateDataRelationships) GetTags() AddressCreateDataRelationshipsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ShipmentCreateDataRelationships) GetTagsOk() (*AddressCreateDataRelationshipsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ShipmentCreateDataRelationships) SetTags(v AddressCreateDataRelationshipsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ShipmentCreateDataRelationships) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


