# LineItemDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**AdyenPaymentDataRelationshipsOrder**](AdyenPaymentDataRelationshipsOrder.md) |  | [optional] 
**Item** | Pointer to [**LineItemDataRelationshipsItem**](LineItemDataRelationshipsItem.md) |  | [optional] 
**LineItemOptions** | Pointer to [**LineItemDataRelationshipsLineItemOptions**](LineItemDataRelationshipsLineItemOptions.md) |  | [optional] 
**ShipmentLineItems** | Pointer to [**LineItemDataRelationshipsShipmentLineItems**](LineItemDataRelationshipsShipmentLineItems.md) |  | [optional] 
**StockLineItems** | Pointer to [**LineItemDataRelationshipsStockLineItems**](LineItemDataRelationshipsStockLineItems.md) |  | [optional] 
**StockTransfers** | Pointer to [**LineItemDataRelationshipsStockTransfers**](LineItemDataRelationshipsStockTransfers.md) |  | [optional] 

## Methods

### NewLineItemDataRelationships

`func NewLineItemDataRelationships() *LineItemDataRelationships`

NewLineItemDataRelationships instantiates a new LineItemDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemDataRelationshipsWithDefaults

`func NewLineItemDataRelationshipsWithDefaults() *LineItemDataRelationships`

NewLineItemDataRelationshipsWithDefaults instantiates a new LineItemDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *LineItemDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *LineItemDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *LineItemDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *LineItemDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetItem

`func (o *LineItemDataRelationships) GetItem() LineItemDataRelationshipsItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *LineItemDataRelationships) GetItemOk() (*LineItemDataRelationshipsItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *LineItemDataRelationships) SetItem(v LineItemDataRelationshipsItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *LineItemDataRelationships) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetLineItemOptions

`func (o *LineItemDataRelationships) GetLineItemOptions() LineItemDataRelationshipsLineItemOptions`

GetLineItemOptions returns the LineItemOptions field if non-nil, zero value otherwise.

### GetLineItemOptionsOk

`func (o *LineItemDataRelationships) GetLineItemOptionsOk() (*LineItemDataRelationshipsLineItemOptions, bool)`

GetLineItemOptionsOk returns a tuple with the LineItemOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItemOptions

`func (o *LineItemDataRelationships) SetLineItemOptions(v LineItemDataRelationshipsLineItemOptions)`

SetLineItemOptions sets LineItemOptions field to given value.

### HasLineItemOptions

`func (o *LineItemDataRelationships) HasLineItemOptions() bool`

HasLineItemOptions returns a boolean if a field has been set.

### GetShipmentLineItems

`func (o *LineItemDataRelationships) GetShipmentLineItems() LineItemDataRelationshipsShipmentLineItems`

GetShipmentLineItems returns the ShipmentLineItems field if non-nil, zero value otherwise.

### GetShipmentLineItemsOk

`func (o *LineItemDataRelationships) GetShipmentLineItemsOk() (*LineItemDataRelationshipsShipmentLineItems, bool)`

GetShipmentLineItemsOk returns a tuple with the ShipmentLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentLineItems

`func (o *LineItemDataRelationships) SetShipmentLineItems(v LineItemDataRelationshipsShipmentLineItems)`

SetShipmentLineItems sets ShipmentLineItems field to given value.

### HasShipmentLineItems

`func (o *LineItemDataRelationships) HasShipmentLineItems() bool`

HasShipmentLineItems returns a boolean if a field has been set.

### GetStockLineItems

`func (o *LineItemDataRelationships) GetStockLineItems() LineItemDataRelationshipsStockLineItems`

GetStockLineItems returns the StockLineItems field if non-nil, zero value otherwise.

### GetStockLineItemsOk

`func (o *LineItemDataRelationships) GetStockLineItemsOk() (*LineItemDataRelationshipsStockLineItems, bool)`

GetStockLineItemsOk returns a tuple with the StockLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLineItems

`func (o *LineItemDataRelationships) SetStockLineItems(v LineItemDataRelationshipsStockLineItems)`

SetStockLineItems sets StockLineItems field to given value.

### HasStockLineItems

`func (o *LineItemDataRelationships) HasStockLineItems() bool`

HasStockLineItems returns a boolean if a field has been set.

### GetStockTransfers

`func (o *LineItemDataRelationships) GetStockTransfers() LineItemDataRelationshipsStockTransfers`

GetStockTransfers returns the StockTransfers field if non-nil, zero value otherwise.

### GetStockTransfersOk

`func (o *LineItemDataRelationships) GetStockTransfersOk() (*LineItemDataRelationshipsStockTransfers, bool)`

GetStockTransfersOk returns a tuple with the StockTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockTransfers

`func (o *LineItemDataRelationships) SetStockTransfers(v LineItemDataRelationshipsStockTransfers)`

SetStockTransfers sets StockTransfers field to given value.

### HasStockTransfers

`func (o *LineItemDataRelationships) HasStockTransfers() bool`

HasStockTransfers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


