# ShippingMethodResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**BillingInfoValidationRuleResponseDataRelationshipsMarket**](BillingInfoValidationRuleResponseDataRelationshipsMarket.md) |  | [optional] 
**ShippingZone** | Pointer to [**ShippingMethodResponseDataRelationshipsShippingZone**](ShippingMethodResponseDataRelationshipsShippingZone.md) |  | [optional] 
**ShippingCategory** | Pointer to [**ShipmentResponseDataRelationshipsShippingCategory**](ShipmentResponseDataRelationshipsShippingCategory.md) |  | [optional] 
**StockLocation** | Pointer to [**DeliveryLeadTimeResponseDataRelationshipsStockLocation**](DeliveryLeadTimeResponseDataRelationshipsStockLocation.md) |  | [optional] 
**DeliveryLeadTimeForShipment** | Pointer to [**ShippingMethodResponseDataRelationshipsDeliveryLeadTimeForShipment**](ShippingMethodResponseDataRelationshipsDeliveryLeadTimeForShipment.md) |  | [optional] 
**ShippingMethodTiers** | Pointer to [**ShippingMethodResponseDataRelationshipsShippingMethodTiers**](ShippingMethodResponseDataRelationshipsShippingMethodTiers.md) |  | [optional] 
**ShippingWeightTiers** | Pointer to [**ShippingMethodResponseDataRelationshipsShippingWeightTiers**](ShippingMethodResponseDataRelationshipsShippingWeightTiers.md) |  | [optional] 
**Attachments** | Pointer to [**AvalaraAccountResponseDataRelationshipsAttachments**](AvalaraAccountResponseDataRelationshipsAttachments.md) |  | [optional] 

## Methods

### NewShippingMethodResponseDataRelationships

`func NewShippingMethodResponseDataRelationships() *ShippingMethodResponseDataRelationships`

NewShippingMethodResponseDataRelationships instantiates a new ShippingMethodResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingMethodResponseDataRelationshipsWithDefaults

`func NewShippingMethodResponseDataRelationshipsWithDefaults() *ShippingMethodResponseDataRelationships`

NewShippingMethodResponseDataRelationshipsWithDefaults instantiates a new ShippingMethodResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *ShippingMethodResponseDataRelationships) GetMarket() BillingInfoValidationRuleResponseDataRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *ShippingMethodResponseDataRelationships) GetMarketOk() (*BillingInfoValidationRuleResponseDataRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *ShippingMethodResponseDataRelationships) SetMarket(v BillingInfoValidationRuleResponseDataRelationshipsMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *ShippingMethodResponseDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetShippingZone

`func (o *ShippingMethodResponseDataRelationships) GetShippingZone() ShippingMethodResponseDataRelationshipsShippingZone`

GetShippingZone returns the ShippingZone field if non-nil, zero value otherwise.

### GetShippingZoneOk

`func (o *ShippingMethodResponseDataRelationships) GetShippingZoneOk() (*ShippingMethodResponseDataRelationshipsShippingZone, bool)`

GetShippingZoneOk returns a tuple with the ShippingZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingZone

`func (o *ShippingMethodResponseDataRelationships) SetShippingZone(v ShippingMethodResponseDataRelationshipsShippingZone)`

SetShippingZone sets ShippingZone field to given value.

### HasShippingZone

`func (o *ShippingMethodResponseDataRelationships) HasShippingZone() bool`

HasShippingZone returns a boolean if a field has been set.

### GetShippingCategory

`func (o *ShippingMethodResponseDataRelationships) GetShippingCategory() ShipmentResponseDataRelationshipsShippingCategory`

GetShippingCategory returns the ShippingCategory field if non-nil, zero value otherwise.

### GetShippingCategoryOk

`func (o *ShippingMethodResponseDataRelationships) GetShippingCategoryOk() (*ShipmentResponseDataRelationshipsShippingCategory, bool)`

GetShippingCategoryOk returns a tuple with the ShippingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCategory

`func (o *ShippingMethodResponseDataRelationships) SetShippingCategory(v ShipmentResponseDataRelationshipsShippingCategory)`

SetShippingCategory sets ShippingCategory field to given value.

### HasShippingCategory

`func (o *ShippingMethodResponseDataRelationships) HasShippingCategory() bool`

HasShippingCategory returns a boolean if a field has been set.

### GetStockLocation

`func (o *ShippingMethodResponseDataRelationships) GetStockLocation() DeliveryLeadTimeResponseDataRelationshipsStockLocation`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *ShippingMethodResponseDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeResponseDataRelationshipsStockLocation, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *ShippingMethodResponseDataRelationships) SetStockLocation(v DeliveryLeadTimeResponseDataRelationshipsStockLocation)`

SetStockLocation sets StockLocation field to given value.

### HasStockLocation

`func (o *ShippingMethodResponseDataRelationships) HasStockLocation() bool`

HasStockLocation returns a boolean if a field has been set.

### GetDeliveryLeadTimeForShipment

`func (o *ShippingMethodResponseDataRelationships) GetDeliveryLeadTimeForShipment() ShippingMethodResponseDataRelationshipsDeliveryLeadTimeForShipment`

GetDeliveryLeadTimeForShipment returns the DeliveryLeadTimeForShipment field if non-nil, zero value otherwise.

### GetDeliveryLeadTimeForShipmentOk

`func (o *ShippingMethodResponseDataRelationships) GetDeliveryLeadTimeForShipmentOk() (*ShippingMethodResponseDataRelationshipsDeliveryLeadTimeForShipment, bool)`

GetDeliveryLeadTimeForShipmentOk returns a tuple with the DeliveryLeadTimeForShipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryLeadTimeForShipment

`func (o *ShippingMethodResponseDataRelationships) SetDeliveryLeadTimeForShipment(v ShippingMethodResponseDataRelationshipsDeliveryLeadTimeForShipment)`

SetDeliveryLeadTimeForShipment sets DeliveryLeadTimeForShipment field to given value.

### HasDeliveryLeadTimeForShipment

`func (o *ShippingMethodResponseDataRelationships) HasDeliveryLeadTimeForShipment() bool`

HasDeliveryLeadTimeForShipment returns a boolean if a field has been set.

### GetShippingMethodTiers

`func (o *ShippingMethodResponseDataRelationships) GetShippingMethodTiers() ShippingMethodResponseDataRelationshipsShippingMethodTiers`

GetShippingMethodTiers returns the ShippingMethodTiers field if non-nil, zero value otherwise.

### GetShippingMethodTiersOk

`func (o *ShippingMethodResponseDataRelationships) GetShippingMethodTiersOk() (*ShippingMethodResponseDataRelationshipsShippingMethodTiers, bool)`

GetShippingMethodTiersOk returns a tuple with the ShippingMethodTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethodTiers

`func (o *ShippingMethodResponseDataRelationships) SetShippingMethodTiers(v ShippingMethodResponseDataRelationshipsShippingMethodTiers)`

SetShippingMethodTiers sets ShippingMethodTiers field to given value.

### HasShippingMethodTiers

`func (o *ShippingMethodResponseDataRelationships) HasShippingMethodTiers() bool`

HasShippingMethodTiers returns a boolean if a field has been set.

### GetShippingWeightTiers

`func (o *ShippingMethodResponseDataRelationships) GetShippingWeightTiers() ShippingMethodResponseDataRelationshipsShippingWeightTiers`

GetShippingWeightTiers returns the ShippingWeightTiers field if non-nil, zero value otherwise.

### GetShippingWeightTiersOk

`func (o *ShippingMethodResponseDataRelationships) GetShippingWeightTiersOk() (*ShippingMethodResponseDataRelationshipsShippingWeightTiers, bool)`

GetShippingWeightTiersOk returns a tuple with the ShippingWeightTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingWeightTiers

`func (o *ShippingMethodResponseDataRelationships) SetShippingWeightTiers(v ShippingMethodResponseDataRelationshipsShippingWeightTiers)`

SetShippingWeightTiers sets ShippingWeightTiers field to given value.

### HasShippingWeightTiers

`func (o *ShippingMethodResponseDataRelationships) HasShippingWeightTiers() bool`

HasShippingWeightTiers returns a boolean if a field has been set.

### GetAttachments

`func (o *ShippingMethodResponseDataRelationships) GetAttachments() AvalaraAccountResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ShippingMethodResponseDataRelationships) GetAttachmentsOk() (*AvalaraAccountResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ShippingMethodResponseDataRelationships) SetAttachments(v AvalaraAccountResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ShippingMethodResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


