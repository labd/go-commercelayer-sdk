# GETSkus200ResponseDataInnerRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShippingCategory** | Pointer to [**GETShipments200ResponseDataInnerRelationshipsShippingCategory**](GETShipments200ResponseDataInnerRelationshipsShippingCategory.md) |  | [optional] 
**Prices** | Pointer to [**GETPriceLists200ResponseDataInnerRelationshipsPrices**](GETPriceLists200ResponseDataInnerRelationshipsPrices.md) |  | [optional] 
**StockItems** | Pointer to [**GETSkus200ResponseDataInnerRelationshipsStockItems**](GETSkus200ResponseDataInnerRelationshipsStockItems.md) |  | [optional] 
**DeliveryLeadTimes** | Pointer to [**GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime**](GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime.md) |  | [optional] 
**SkuOptions** | Pointer to [**GETLineItemOptions200ResponseDataInnerRelationshipsSkuOption**](GETLineItemOptions200ResponseDataInnerRelationshipsSkuOption.md) |  | [optional] 
**Attachments** | Pointer to [**GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments**](GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments.md) |  | [optional] 

## Methods

### NewGETSkus200ResponseDataInnerRelationships

`func NewGETSkus200ResponseDataInnerRelationships() *GETSkus200ResponseDataInnerRelationships`

NewGETSkus200ResponseDataInnerRelationships instantiates a new GETSkus200ResponseDataInnerRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETSkus200ResponseDataInnerRelationshipsWithDefaults

`func NewGETSkus200ResponseDataInnerRelationshipsWithDefaults() *GETSkus200ResponseDataInnerRelationships`

NewGETSkus200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETSkus200ResponseDataInnerRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShippingCategory

`func (o *GETSkus200ResponseDataInnerRelationships) GetShippingCategory() GETShipments200ResponseDataInnerRelationshipsShippingCategory`

GetShippingCategory returns the ShippingCategory field if non-nil, zero value otherwise.

### GetShippingCategoryOk

`func (o *GETSkus200ResponseDataInnerRelationships) GetShippingCategoryOk() (*GETShipments200ResponseDataInnerRelationshipsShippingCategory, bool)`

GetShippingCategoryOk returns a tuple with the ShippingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCategory

`func (o *GETSkus200ResponseDataInnerRelationships) SetShippingCategory(v GETShipments200ResponseDataInnerRelationshipsShippingCategory)`

SetShippingCategory sets ShippingCategory field to given value.

### HasShippingCategory

`func (o *GETSkus200ResponseDataInnerRelationships) HasShippingCategory() bool`

HasShippingCategory returns a boolean if a field has been set.

### GetPrices

`func (o *GETSkus200ResponseDataInnerRelationships) GetPrices() GETPriceLists200ResponseDataInnerRelationshipsPrices`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *GETSkus200ResponseDataInnerRelationships) GetPricesOk() (*GETPriceLists200ResponseDataInnerRelationshipsPrices, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *GETSkus200ResponseDataInnerRelationships) SetPrices(v GETPriceLists200ResponseDataInnerRelationshipsPrices)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *GETSkus200ResponseDataInnerRelationships) HasPrices() bool`

HasPrices returns a boolean if a field has been set.

### GetStockItems

`func (o *GETSkus200ResponseDataInnerRelationships) GetStockItems() GETSkus200ResponseDataInnerRelationshipsStockItems`

GetStockItems returns the StockItems field if non-nil, zero value otherwise.

### GetStockItemsOk

`func (o *GETSkus200ResponseDataInnerRelationships) GetStockItemsOk() (*GETSkus200ResponseDataInnerRelationshipsStockItems, bool)`

GetStockItemsOk returns a tuple with the StockItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockItems

`func (o *GETSkus200ResponseDataInnerRelationships) SetStockItems(v GETSkus200ResponseDataInnerRelationshipsStockItems)`

SetStockItems sets StockItems field to given value.

### HasStockItems

`func (o *GETSkus200ResponseDataInnerRelationships) HasStockItems() bool`

HasStockItems returns a boolean if a field has been set.

### GetDeliveryLeadTimes

`func (o *GETSkus200ResponseDataInnerRelationships) GetDeliveryLeadTimes() GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime`

GetDeliveryLeadTimes returns the DeliveryLeadTimes field if non-nil, zero value otherwise.

### GetDeliveryLeadTimesOk

`func (o *GETSkus200ResponseDataInnerRelationships) GetDeliveryLeadTimesOk() (*GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime, bool)`

GetDeliveryLeadTimesOk returns a tuple with the DeliveryLeadTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryLeadTimes

`func (o *GETSkus200ResponseDataInnerRelationships) SetDeliveryLeadTimes(v GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime)`

SetDeliveryLeadTimes sets DeliveryLeadTimes field to given value.

### HasDeliveryLeadTimes

`func (o *GETSkus200ResponseDataInnerRelationships) HasDeliveryLeadTimes() bool`

HasDeliveryLeadTimes returns a boolean if a field has been set.

### GetSkuOptions

`func (o *GETSkus200ResponseDataInnerRelationships) GetSkuOptions() GETLineItemOptions200ResponseDataInnerRelationshipsSkuOption`

GetSkuOptions returns the SkuOptions field if non-nil, zero value otherwise.

### GetSkuOptionsOk

`func (o *GETSkus200ResponseDataInnerRelationships) GetSkuOptionsOk() (*GETLineItemOptions200ResponseDataInnerRelationshipsSkuOption, bool)`

GetSkuOptionsOk returns a tuple with the SkuOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuOptions

`func (o *GETSkus200ResponseDataInnerRelationships) SetSkuOptions(v GETLineItemOptions200ResponseDataInnerRelationshipsSkuOption)`

SetSkuOptions sets SkuOptions field to given value.

### HasSkuOptions

`func (o *GETSkus200ResponseDataInnerRelationships) HasSkuOptions() bool`

HasSkuOptions returns a boolean if a field has been set.

### GetAttachments

`func (o *GETSkus200ResponseDataInnerRelationships) GetAttachments() GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GETSkus200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GETSkus200ResponseDataInnerRelationships) SetAttachments(v GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GETSkus200ResponseDataInnerRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


