# SkuResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShippingCategory** | Pointer to [**ShipmentResponseDataRelationshipsShippingCategory**](ShipmentResponseDataRelationshipsShippingCategory.md) |  | [optional] 
**Prices** | Pointer to [**PriceListResponseDataRelationshipsPrices**](PriceListResponseDataRelationshipsPrices.md) |  | [optional] 
**StockItems** | Pointer to [**SkuResponseDataRelationshipsStockItems**](SkuResponseDataRelationshipsStockItems.md) |  | [optional] 
**DeliveryLeadTimes** | Pointer to [**SkuResponseDataRelationshipsDeliveryLeadTimes**](SkuResponseDataRelationshipsDeliveryLeadTimes.md) |  | [optional] 
**SkuOptions** | Pointer to [**SkuResponseDataRelationshipsSkuOptions**](SkuResponseDataRelationshipsSkuOptions.md) |  | [optional] 
**Attachments** | Pointer to [**AvalaraAccountResponseDataRelationshipsAttachments**](AvalaraAccountResponseDataRelationshipsAttachments.md) |  | [optional] 

## Methods

### NewSkuResponseDataRelationships

`func NewSkuResponseDataRelationships() *SkuResponseDataRelationships`

NewSkuResponseDataRelationships instantiates a new SkuResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuResponseDataRelationshipsWithDefaults

`func NewSkuResponseDataRelationshipsWithDefaults() *SkuResponseDataRelationships`

NewSkuResponseDataRelationshipsWithDefaults instantiates a new SkuResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShippingCategory

`func (o *SkuResponseDataRelationships) GetShippingCategory() ShipmentResponseDataRelationshipsShippingCategory`

GetShippingCategory returns the ShippingCategory field if non-nil, zero value otherwise.

### GetShippingCategoryOk

`func (o *SkuResponseDataRelationships) GetShippingCategoryOk() (*ShipmentResponseDataRelationshipsShippingCategory, bool)`

GetShippingCategoryOk returns a tuple with the ShippingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCategory

`func (o *SkuResponseDataRelationships) SetShippingCategory(v ShipmentResponseDataRelationshipsShippingCategory)`

SetShippingCategory sets ShippingCategory field to given value.

### HasShippingCategory

`func (o *SkuResponseDataRelationships) HasShippingCategory() bool`

HasShippingCategory returns a boolean if a field has been set.

### GetPrices

`func (o *SkuResponseDataRelationships) GetPrices() PriceListResponseDataRelationshipsPrices`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *SkuResponseDataRelationships) GetPricesOk() (*PriceListResponseDataRelationshipsPrices, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *SkuResponseDataRelationships) SetPrices(v PriceListResponseDataRelationshipsPrices)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *SkuResponseDataRelationships) HasPrices() bool`

HasPrices returns a boolean if a field has been set.

### GetStockItems

`func (o *SkuResponseDataRelationships) GetStockItems() SkuResponseDataRelationshipsStockItems`

GetStockItems returns the StockItems field if non-nil, zero value otherwise.

### GetStockItemsOk

`func (o *SkuResponseDataRelationships) GetStockItemsOk() (*SkuResponseDataRelationshipsStockItems, bool)`

GetStockItemsOk returns a tuple with the StockItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockItems

`func (o *SkuResponseDataRelationships) SetStockItems(v SkuResponseDataRelationshipsStockItems)`

SetStockItems sets StockItems field to given value.

### HasStockItems

`func (o *SkuResponseDataRelationships) HasStockItems() bool`

HasStockItems returns a boolean if a field has been set.

### GetDeliveryLeadTimes

`func (o *SkuResponseDataRelationships) GetDeliveryLeadTimes() SkuResponseDataRelationshipsDeliveryLeadTimes`

GetDeliveryLeadTimes returns the DeliveryLeadTimes field if non-nil, zero value otherwise.

### GetDeliveryLeadTimesOk

`func (o *SkuResponseDataRelationships) GetDeliveryLeadTimesOk() (*SkuResponseDataRelationshipsDeliveryLeadTimes, bool)`

GetDeliveryLeadTimesOk returns a tuple with the DeliveryLeadTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryLeadTimes

`func (o *SkuResponseDataRelationships) SetDeliveryLeadTimes(v SkuResponseDataRelationshipsDeliveryLeadTimes)`

SetDeliveryLeadTimes sets DeliveryLeadTimes field to given value.

### HasDeliveryLeadTimes

`func (o *SkuResponseDataRelationships) HasDeliveryLeadTimes() bool`

HasDeliveryLeadTimes returns a boolean if a field has been set.

### GetSkuOptions

`func (o *SkuResponseDataRelationships) GetSkuOptions() SkuResponseDataRelationshipsSkuOptions`

GetSkuOptions returns the SkuOptions field if non-nil, zero value otherwise.

### GetSkuOptionsOk

`func (o *SkuResponseDataRelationships) GetSkuOptionsOk() (*SkuResponseDataRelationshipsSkuOptions, bool)`

GetSkuOptionsOk returns a tuple with the SkuOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuOptions

`func (o *SkuResponseDataRelationships) SetSkuOptions(v SkuResponseDataRelationshipsSkuOptions)`

SetSkuOptions sets SkuOptions field to given value.

### HasSkuOptions

`func (o *SkuResponseDataRelationships) HasSkuOptions() bool`

HasSkuOptions returns a boolean if a field has been set.

### GetAttachments

`func (o *SkuResponseDataRelationships) GetAttachments() AvalaraAccountResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *SkuResponseDataRelationships) GetAttachmentsOk() (*AvalaraAccountResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *SkuResponseDataRelationships) SetAttachments(v AvalaraAccountResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *SkuResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


