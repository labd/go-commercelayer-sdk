# StockLineItemDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shipment** | Pointer to [**OrderDataRelationshipsShipments**](OrderDataRelationshipsShipments.md) |  | [optional] 
**LineItem** | Pointer to [**LineItemOptionDataRelationshipsLineItem**](LineItemOptionDataRelationshipsLineItem.md) |  | [optional] 
**StockItem** | Pointer to [**SkuDataRelationshipsStockItems**](SkuDataRelationshipsStockItems.md) |  | [optional] 

## Methods

### NewStockLineItemDataRelationships

`func NewStockLineItemDataRelationships() *StockLineItemDataRelationships`

NewStockLineItemDataRelationships instantiates a new StockLineItemDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockLineItemDataRelationshipsWithDefaults

`func NewStockLineItemDataRelationshipsWithDefaults() *StockLineItemDataRelationships`

NewStockLineItemDataRelationshipsWithDefaults instantiates a new StockLineItemDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipment

`func (o *StockLineItemDataRelationships) GetShipment() OrderDataRelationshipsShipments`

GetShipment returns the Shipment field if non-nil, zero value otherwise.

### GetShipmentOk

`func (o *StockLineItemDataRelationships) GetShipmentOk() (*OrderDataRelationshipsShipments, bool)`

GetShipmentOk returns a tuple with the Shipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipment

`func (o *StockLineItemDataRelationships) SetShipment(v OrderDataRelationshipsShipments)`

SetShipment sets Shipment field to given value.

### HasShipment

`func (o *StockLineItemDataRelationships) HasShipment() bool`

HasShipment returns a boolean if a field has been set.

### GetLineItem

`func (o *StockLineItemDataRelationships) GetLineItem() LineItemOptionDataRelationshipsLineItem`

GetLineItem returns the LineItem field if non-nil, zero value otherwise.

### GetLineItemOk

`func (o *StockLineItemDataRelationships) GetLineItemOk() (*LineItemOptionDataRelationshipsLineItem, bool)`

GetLineItemOk returns a tuple with the LineItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItem

`func (o *StockLineItemDataRelationships) SetLineItem(v LineItemOptionDataRelationshipsLineItem)`

SetLineItem sets LineItem field to given value.

### HasLineItem

`func (o *StockLineItemDataRelationships) HasLineItem() bool`

HasLineItem returns a boolean if a field has been set.

### GetStockItem

`func (o *StockLineItemDataRelationships) GetStockItem() SkuDataRelationshipsStockItems`

GetStockItem returns the StockItem field if non-nil, zero value otherwise.

### GetStockItemOk

`func (o *StockLineItemDataRelationships) GetStockItemOk() (*SkuDataRelationshipsStockItems, bool)`

GetStockItemOk returns a tuple with the StockItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockItem

`func (o *StockLineItemDataRelationships) SetStockItem(v SkuDataRelationshipsStockItems)`

SetStockItem sets StockItem field to given value.

### HasStockItem

`func (o *StockLineItemDataRelationships) HasStockItem() bool`

HasStockItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


