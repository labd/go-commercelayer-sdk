# SkuDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShippingCategory** | Pointer to [**ShipmentDataRelationshipsShippingCategory**](ShipmentDataRelationshipsShippingCategory.md) |  | [optional] 
**Prices** | Pointer to [**PriceFrequencyTierDataRelationshipsPrice**](PriceFrequencyTierDataRelationshipsPrice.md) |  | [optional] 
**StockItems** | Pointer to [**SkuDataRelationshipsStockItems**](SkuDataRelationshipsStockItems.md) |  | [optional] 
**DeliveryLeadTimes** | Pointer to [**ShipmentDataRelationshipsDeliveryLeadTime**](ShipmentDataRelationshipsDeliveryLeadTime.md) |  | [optional] 
**SkuOptions** | Pointer to [**LineItemOptionDataRelationshipsSkuOption**](LineItemOptionDataRelationshipsSkuOption.md) |  | [optional] 
**Attachments** | Pointer to [**AvalaraAccountDataRelationshipsAttachments**](AvalaraAccountDataRelationshipsAttachments.md) |  | [optional] 

## Methods

### NewSkuDataRelationships

`func NewSkuDataRelationships() *SkuDataRelationships`

NewSkuDataRelationships instantiates a new SkuDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuDataRelationshipsWithDefaults

`func NewSkuDataRelationshipsWithDefaults() *SkuDataRelationships`

NewSkuDataRelationshipsWithDefaults instantiates a new SkuDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShippingCategory

`func (o *SkuDataRelationships) GetShippingCategory() ShipmentDataRelationshipsShippingCategory`

GetShippingCategory returns the ShippingCategory field if non-nil, zero value otherwise.

### GetShippingCategoryOk

`func (o *SkuDataRelationships) GetShippingCategoryOk() (*ShipmentDataRelationshipsShippingCategory, bool)`

GetShippingCategoryOk returns a tuple with the ShippingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCategory

`func (o *SkuDataRelationships) SetShippingCategory(v ShipmentDataRelationshipsShippingCategory)`

SetShippingCategory sets ShippingCategory field to given value.

### HasShippingCategory

`func (o *SkuDataRelationships) HasShippingCategory() bool`

HasShippingCategory returns a boolean if a field has been set.

### GetPrices

`func (o *SkuDataRelationships) GetPrices() PriceFrequencyTierDataRelationshipsPrice`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *SkuDataRelationships) GetPricesOk() (*PriceFrequencyTierDataRelationshipsPrice, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *SkuDataRelationships) SetPrices(v PriceFrequencyTierDataRelationshipsPrice)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *SkuDataRelationships) HasPrices() bool`

HasPrices returns a boolean if a field has been set.

### GetStockItems

`func (o *SkuDataRelationships) GetStockItems() SkuDataRelationshipsStockItems`

GetStockItems returns the StockItems field if non-nil, zero value otherwise.

### GetStockItemsOk

`func (o *SkuDataRelationships) GetStockItemsOk() (*SkuDataRelationshipsStockItems, bool)`

GetStockItemsOk returns a tuple with the StockItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockItems

`func (o *SkuDataRelationships) SetStockItems(v SkuDataRelationshipsStockItems)`

SetStockItems sets StockItems field to given value.

### HasStockItems

`func (o *SkuDataRelationships) HasStockItems() bool`

HasStockItems returns a boolean if a field has been set.

### GetDeliveryLeadTimes

`func (o *SkuDataRelationships) GetDeliveryLeadTimes() ShipmentDataRelationshipsDeliveryLeadTime`

GetDeliveryLeadTimes returns the DeliveryLeadTimes field if non-nil, zero value otherwise.

### GetDeliveryLeadTimesOk

`func (o *SkuDataRelationships) GetDeliveryLeadTimesOk() (*ShipmentDataRelationshipsDeliveryLeadTime, bool)`

GetDeliveryLeadTimesOk returns a tuple with the DeliveryLeadTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryLeadTimes

`func (o *SkuDataRelationships) SetDeliveryLeadTimes(v ShipmentDataRelationshipsDeliveryLeadTime)`

SetDeliveryLeadTimes sets DeliveryLeadTimes field to given value.

### HasDeliveryLeadTimes

`func (o *SkuDataRelationships) HasDeliveryLeadTimes() bool`

HasDeliveryLeadTimes returns a boolean if a field has been set.

### GetSkuOptions

`func (o *SkuDataRelationships) GetSkuOptions() LineItemOptionDataRelationshipsSkuOption`

GetSkuOptions returns the SkuOptions field if non-nil, zero value otherwise.

### GetSkuOptionsOk

`func (o *SkuDataRelationships) GetSkuOptionsOk() (*LineItemOptionDataRelationshipsSkuOption, bool)`

GetSkuOptionsOk returns a tuple with the SkuOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuOptions

`func (o *SkuDataRelationships) SetSkuOptions(v LineItemOptionDataRelationshipsSkuOption)`

SetSkuOptions sets SkuOptions field to given value.

### HasSkuOptions

`func (o *SkuDataRelationships) HasSkuOptions() bool`

HasSkuOptions returns a boolean if a field has been set.

### GetAttachments

`func (o *SkuDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *SkuDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *SkuDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *SkuDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


