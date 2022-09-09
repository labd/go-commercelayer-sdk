# ShippingMethodDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**AvalaraAccountDataRelationshipsMarkets**](AvalaraAccountDataRelationshipsMarkets.md) |  | [optional] 
**ShippingZone** | Pointer to [**ShippingMethodDataRelationshipsShippingZone**](ShippingMethodDataRelationshipsShippingZone.md) |  | [optional] 
**ShippingCategory** | Pointer to [**ShipmentDataRelationshipsShippingCategory**](ShipmentDataRelationshipsShippingCategory.md) |  | [optional] 
**StockLocation** | Pointer to [**DeliveryLeadTimeDataRelationshipsStockLocation**](DeliveryLeadTimeDataRelationshipsStockLocation.md) |  | [optional] 
**DeliveryLeadTimeForShipment** | Pointer to [**ShipmentDataRelationshipsDeliveryLeadTime**](ShipmentDataRelationshipsDeliveryLeadTime.md) |  | [optional] 
**ShippingMethodTiers** | Pointer to [**ShippingMethodDataRelationshipsShippingMethodTiers**](ShippingMethodDataRelationshipsShippingMethodTiers.md) |  | [optional] 
**ShippingWeightTiers** | Pointer to [**ShippingMethodDataRelationshipsShippingWeightTiers**](ShippingMethodDataRelationshipsShippingWeightTiers.md) |  | [optional] 
**Attachments** | Pointer to [**AvalaraAccountDataRelationshipsAttachments**](AvalaraAccountDataRelationshipsAttachments.md) |  | [optional] 

## Methods

### NewShippingMethodDataRelationships

`func NewShippingMethodDataRelationships() *ShippingMethodDataRelationships`

NewShippingMethodDataRelationships instantiates a new ShippingMethodDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingMethodDataRelationshipsWithDefaults

`func NewShippingMethodDataRelationshipsWithDefaults() *ShippingMethodDataRelationships`

NewShippingMethodDataRelationshipsWithDefaults instantiates a new ShippingMethodDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *ShippingMethodDataRelationships) GetMarket() AvalaraAccountDataRelationshipsMarkets`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *ShippingMethodDataRelationships) GetMarketOk() (*AvalaraAccountDataRelationshipsMarkets, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *ShippingMethodDataRelationships) SetMarket(v AvalaraAccountDataRelationshipsMarkets)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *ShippingMethodDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetShippingZone

`func (o *ShippingMethodDataRelationships) GetShippingZone() ShippingMethodDataRelationshipsShippingZone`

GetShippingZone returns the ShippingZone field if non-nil, zero value otherwise.

### GetShippingZoneOk

`func (o *ShippingMethodDataRelationships) GetShippingZoneOk() (*ShippingMethodDataRelationshipsShippingZone, bool)`

GetShippingZoneOk returns a tuple with the ShippingZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingZone

`func (o *ShippingMethodDataRelationships) SetShippingZone(v ShippingMethodDataRelationshipsShippingZone)`

SetShippingZone sets ShippingZone field to given value.

### HasShippingZone

`func (o *ShippingMethodDataRelationships) HasShippingZone() bool`

HasShippingZone returns a boolean if a field has been set.

### GetShippingCategory

`func (o *ShippingMethodDataRelationships) GetShippingCategory() ShipmentDataRelationshipsShippingCategory`

GetShippingCategory returns the ShippingCategory field if non-nil, zero value otherwise.

### GetShippingCategoryOk

`func (o *ShippingMethodDataRelationships) GetShippingCategoryOk() (*ShipmentDataRelationshipsShippingCategory, bool)`

GetShippingCategoryOk returns a tuple with the ShippingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCategory

`func (o *ShippingMethodDataRelationships) SetShippingCategory(v ShipmentDataRelationshipsShippingCategory)`

SetShippingCategory sets ShippingCategory field to given value.

### HasShippingCategory

`func (o *ShippingMethodDataRelationships) HasShippingCategory() bool`

HasShippingCategory returns a boolean if a field has been set.

### GetStockLocation

`func (o *ShippingMethodDataRelationships) GetStockLocation() DeliveryLeadTimeDataRelationshipsStockLocation`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *ShippingMethodDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeDataRelationshipsStockLocation, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *ShippingMethodDataRelationships) SetStockLocation(v DeliveryLeadTimeDataRelationshipsStockLocation)`

SetStockLocation sets StockLocation field to given value.

### HasStockLocation

`func (o *ShippingMethodDataRelationships) HasStockLocation() bool`

HasStockLocation returns a boolean if a field has been set.

### GetDeliveryLeadTimeForShipment

`func (o *ShippingMethodDataRelationships) GetDeliveryLeadTimeForShipment() ShipmentDataRelationshipsDeliveryLeadTime`

GetDeliveryLeadTimeForShipment returns the DeliveryLeadTimeForShipment field if non-nil, zero value otherwise.

### GetDeliveryLeadTimeForShipmentOk

`func (o *ShippingMethodDataRelationships) GetDeliveryLeadTimeForShipmentOk() (*ShipmentDataRelationshipsDeliveryLeadTime, bool)`

GetDeliveryLeadTimeForShipmentOk returns a tuple with the DeliveryLeadTimeForShipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryLeadTimeForShipment

`func (o *ShippingMethodDataRelationships) SetDeliveryLeadTimeForShipment(v ShipmentDataRelationshipsDeliveryLeadTime)`

SetDeliveryLeadTimeForShipment sets DeliveryLeadTimeForShipment field to given value.

### HasDeliveryLeadTimeForShipment

`func (o *ShippingMethodDataRelationships) HasDeliveryLeadTimeForShipment() bool`

HasDeliveryLeadTimeForShipment returns a boolean if a field has been set.

### GetShippingMethodTiers

`func (o *ShippingMethodDataRelationships) GetShippingMethodTiers() ShippingMethodDataRelationshipsShippingMethodTiers`

GetShippingMethodTiers returns the ShippingMethodTiers field if non-nil, zero value otherwise.

### GetShippingMethodTiersOk

`func (o *ShippingMethodDataRelationships) GetShippingMethodTiersOk() (*ShippingMethodDataRelationshipsShippingMethodTiers, bool)`

GetShippingMethodTiersOk returns a tuple with the ShippingMethodTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethodTiers

`func (o *ShippingMethodDataRelationships) SetShippingMethodTiers(v ShippingMethodDataRelationshipsShippingMethodTiers)`

SetShippingMethodTiers sets ShippingMethodTiers field to given value.

### HasShippingMethodTiers

`func (o *ShippingMethodDataRelationships) HasShippingMethodTiers() bool`

HasShippingMethodTiers returns a boolean if a field has been set.

### GetShippingWeightTiers

`func (o *ShippingMethodDataRelationships) GetShippingWeightTiers() ShippingMethodDataRelationshipsShippingWeightTiers`

GetShippingWeightTiers returns the ShippingWeightTiers field if non-nil, zero value otherwise.

### GetShippingWeightTiersOk

`func (o *ShippingMethodDataRelationships) GetShippingWeightTiersOk() (*ShippingMethodDataRelationshipsShippingWeightTiers, bool)`

GetShippingWeightTiersOk returns a tuple with the ShippingWeightTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingWeightTiers

`func (o *ShippingMethodDataRelationships) SetShippingWeightTiers(v ShippingMethodDataRelationshipsShippingWeightTiers)`

SetShippingWeightTiers sets ShippingWeightTiers field to given value.

### HasShippingWeightTiers

`func (o *ShippingMethodDataRelationships) HasShippingWeightTiers() bool`

HasShippingWeightTiers returns a boolean if a field has been set.

### GetAttachments

`func (o *ShippingMethodDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ShippingMethodDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ShippingMethodDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ShippingMethodDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


