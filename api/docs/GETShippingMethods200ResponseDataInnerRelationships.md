# GETShippingMethods200ResponseDataInnerRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket**](GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket.md) |  | [optional] 
**ShippingZone** | Pointer to [**GETShippingMethods200ResponseDataInnerRelationshipsShippingZone**](GETShippingMethods200ResponseDataInnerRelationshipsShippingZone.md) |  | [optional] 
**ShippingCategory** | Pointer to [**GETShipments200ResponseDataInnerRelationshipsShippingCategory**](GETShipments200ResponseDataInnerRelationshipsShippingCategory.md) |  | [optional] 
**StockLocation** | Pointer to [**GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation**](GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation.md) |  | [optional] 
**DeliveryLeadTimeForShipment** | Pointer to [**GETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipment**](GETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipment.md) |  | [optional] 
**ShippingMethodTiers** | Pointer to [**GETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiers**](GETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiers.md) |  | [optional] 
**ShippingWeightTiers** | Pointer to [**GETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers**](GETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers.md) |  | [optional] 
**Attachments** | Pointer to [**GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments**](GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments.md) |  | [optional] 

## Methods

### NewGETShippingMethods200ResponseDataInnerRelationships

`func NewGETShippingMethods200ResponseDataInnerRelationships() *GETShippingMethods200ResponseDataInnerRelationships`

NewGETShippingMethods200ResponseDataInnerRelationships instantiates a new GETShippingMethods200ResponseDataInnerRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETShippingMethods200ResponseDataInnerRelationshipsWithDefaults

`func NewGETShippingMethods200ResponseDataInnerRelationshipsWithDefaults() *GETShippingMethods200ResponseDataInnerRelationships`

NewGETShippingMethods200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETShippingMethods200ResponseDataInnerRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *GETShippingMethods200ResponseDataInnerRelationships) GetMarket() GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *GETShippingMethods200ResponseDataInnerRelationships) GetMarketOk() (*GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *GETShippingMethods200ResponseDataInnerRelationships) SetMarket(v GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *GETShippingMethods200ResponseDataInnerRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetShippingZone

`func (o *GETShippingMethods200ResponseDataInnerRelationships) GetShippingZone() GETShippingMethods200ResponseDataInnerRelationshipsShippingZone`

GetShippingZone returns the ShippingZone field if non-nil, zero value otherwise.

### GetShippingZoneOk

`func (o *GETShippingMethods200ResponseDataInnerRelationships) GetShippingZoneOk() (*GETShippingMethods200ResponseDataInnerRelationshipsShippingZone, bool)`

GetShippingZoneOk returns a tuple with the ShippingZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingZone

`func (o *GETShippingMethods200ResponseDataInnerRelationships) SetShippingZone(v GETShippingMethods200ResponseDataInnerRelationshipsShippingZone)`

SetShippingZone sets ShippingZone field to given value.

### HasShippingZone

`func (o *GETShippingMethods200ResponseDataInnerRelationships) HasShippingZone() bool`

HasShippingZone returns a boolean if a field has been set.

### GetShippingCategory

`func (o *GETShippingMethods200ResponseDataInnerRelationships) GetShippingCategory() GETShipments200ResponseDataInnerRelationshipsShippingCategory`

GetShippingCategory returns the ShippingCategory field if non-nil, zero value otherwise.

### GetShippingCategoryOk

`func (o *GETShippingMethods200ResponseDataInnerRelationships) GetShippingCategoryOk() (*GETShipments200ResponseDataInnerRelationshipsShippingCategory, bool)`

GetShippingCategoryOk returns a tuple with the ShippingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCategory

`func (o *GETShippingMethods200ResponseDataInnerRelationships) SetShippingCategory(v GETShipments200ResponseDataInnerRelationshipsShippingCategory)`

SetShippingCategory sets ShippingCategory field to given value.

### HasShippingCategory

`func (o *GETShippingMethods200ResponseDataInnerRelationships) HasShippingCategory() bool`

HasShippingCategory returns a boolean if a field has been set.

### GetStockLocation

`func (o *GETShippingMethods200ResponseDataInnerRelationships) GetStockLocation() GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *GETShippingMethods200ResponseDataInnerRelationships) GetStockLocationOk() (*GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *GETShippingMethods200ResponseDataInnerRelationships) SetStockLocation(v GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation)`

SetStockLocation sets StockLocation field to given value.

### HasStockLocation

`func (o *GETShippingMethods200ResponseDataInnerRelationships) HasStockLocation() bool`

HasStockLocation returns a boolean if a field has been set.

### GetDeliveryLeadTimeForShipment

`func (o *GETShippingMethods200ResponseDataInnerRelationships) GetDeliveryLeadTimeForShipment() GETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipment`

GetDeliveryLeadTimeForShipment returns the DeliveryLeadTimeForShipment field if non-nil, zero value otherwise.

### GetDeliveryLeadTimeForShipmentOk

`func (o *GETShippingMethods200ResponseDataInnerRelationships) GetDeliveryLeadTimeForShipmentOk() (*GETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipment, bool)`

GetDeliveryLeadTimeForShipmentOk returns a tuple with the DeliveryLeadTimeForShipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryLeadTimeForShipment

`func (o *GETShippingMethods200ResponseDataInnerRelationships) SetDeliveryLeadTimeForShipment(v GETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipment)`

SetDeliveryLeadTimeForShipment sets DeliveryLeadTimeForShipment field to given value.

### HasDeliveryLeadTimeForShipment

`func (o *GETShippingMethods200ResponseDataInnerRelationships) HasDeliveryLeadTimeForShipment() bool`

HasDeliveryLeadTimeForShipment returns a boolean if a field has been set.

### GetShippingMethodTiers

`func (o *GETShippingMethods200ResponseDataInnerRelationships) GetShippingMethodTiers() GETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiers`

GetShippingMethodTiers returns the ShippingMethodTiers field if non-nil, zero value otherwise.

### GetShippingMethodTiersOk

`func (o *GETShippingMethods200ResponseDataInnerRelationships) GetShippingMethodTiersOk() (*GETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiers, bool)`

GetShippingMethodTiersOk returns a tuple with the ShippingMethodTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethodTiers

`func (o *GETShippingMethods200ResponseDataInnerRelationships) SetShippingMethodTiers(v GETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiers)`

SetShippingMethodTiers sets ShippingMethodTiers field to given value.

### HasShippingMethodTiers

`func (o *GETShippingMethods200ResponseDataInnerRelationships) HasShippingMethodTiers() bool`

HasShippingMethodTiers returns a boolean if a field has been set.

### GetShippingWeightTiers

`func (o *GETShippingMethods200ResponseDataInnerRelationships) GetShippingWeightTiers() GETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers`

GetShippingWeightTiers returns the ShippingWeightTiers field if non-nil, zero value otherwise.

### GetShippingWeightTiersOk

`func (o *GETShippingMethods200ResponseDataInnerRelationships) GetShippingWeightTiersOk() (*GETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers, bool)`

GetShippingWeightTiersOk returns a tuple with the ShippingWeightTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingWeightTiers

`func (o *GETShippingMethods200ResponseDataInnerRelationships) SetShippingWeightTiers(v GETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers)`

SetShippingWeightTiers sets ShippingWeightTiers field to given value.

### HasShippingWeightTiers

`func (o *GETShippingMethods200ResponseDataInnerRelationships) HasShippingWeightTiers() bool`

HasShippingWeightTiers returns a boolean if a field has been set.

### GetAttachments

`func (o *GETShippingMethods200ResponseDataInnerRelationships) GetAttachments() GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GETShippingMethods200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GETShippingMethods200ResponseDataInnerRelationships) SetAttachments(v GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GETShippingMethods200ResponseDataInnerRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


