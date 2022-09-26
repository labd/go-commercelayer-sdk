# LineItemResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**AdyenPaymentResponseDataRelationshipsOrder**](AdyenPaymentResponseDataRelationshipsOrder.md) |  | [optional] 
**Item** | Pointer to [**LineItemResponseDataRelationshipsItem**](LineItemResponseDataRelationshipsItem.md) |  | [optional] 
**LineItemOptions** | Pointer to [**LineItemResponseDataRelationshipsLineItemOptions**](LineItemResponseDataRelationshipsLineItemOptions.md) |  | [optional] 
**ShipmentLineItems** | Pointer to [**LineItemResponseDataRelationshipsShipmentLineItems**](LineItemResponseDataRelationshipsShipmentLineItems.md) |  | [optional] 
**StockLineItems** | Pointer to [**LineItemResponseDataRelationshipsStockLineItems**](LineItemResponseDataRelationshipsStockLineItems.md) |  | [optional] 
**StockTransfers** | Pointer to [**LineItemResponseDataRelationshipsStockTransfers**](LineItemResponseDataRelationshipsStockTransfers.md) |  | [optional] 

## Methods

### NewLineItemResponseDataRelationships

`func NewLineItemResponseDataRelationships() *LineItemResponseDataRelationships`

NewLineItemResponseDataRelationships instantiates a new LineItemResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemResponseDataRelationshipsWithDefaults

`func NewLineItemResponseDataRelationshipsWithDefaults() *LineItemResponseDataRelationships`

NewLineItemResponseDataRelationshipsWithDefaults instantiates a new LineItemResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *LineItemResponseDataRelationships) GetOrder() AdyenPaymentResponseDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *LineItemResponseDataRelationships) GetOrderOk() (*AdyenPaymentResponseDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *LineItemResponseDataRelationships) SetOrder(v AdyenPaymentResponseDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *LineItemResponseDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetItem

`func (o *LineItemResponseDataRelationships) GetItem() LineItemResponseDataRelationshipsItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *LineItemResponseDataRelationships) GetItemOk() (*LineItemResponseDataRelationshipsItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *LineItemResponseDataRelationships) SetItem(v LineItemResponseDataRelationshipsItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *LineItemResponseDataRelationships) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetLineItemOptions

`func (o *LineItemResponseDataRelationships) GetLineItemOptions() LineItemResponseDataRelationshipsLineItemOptions`

GetLineItemOptions returns the LineItemOptions field if non-nil, zero value otherwise.

### GetLineItemOptionsOk

`func (o *LineItemResponseDataRelationships) GetLineItemOptionsOk() (*LineItemResponseDataRelationshipsLineItemOptions, bool)`

GetLineItemOptionsOk returns a tuple with the LineItemOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItemOptions

`func (o *LineItemResponseDataRelationships) SetLineItemOptions(v LineItemResponseDataRelationshipsLineItemOptions)`

SetLineItemOptions sets LineItemOptions field to given value.

### HasLineItemOptions

`func (o *LineItemResponseDataRelationships) HasLineItemOptions() bool`

HasLineItemOptions returns a boolean if a field has been set.

### GetShipmentLineItems

`func (o *LineItemResponseDataRelationships) GetShipmentLineItems() LineItemResponseDataRelationshipsShipmentLineItems`

GetShipmentLineItems returns the ShipmentLineItems field if non-nil, zero value otherwise.

### GetShipmentLineItemsOk

`func (o *LineItemResponseDataRelationships) GetShipmentLineItemsOk() (*LineItemResponseDataRelationshipsShipmentLineItems, bool)`

GetShipmentLineItemsOk returns a tuple with the ShipmentLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentLineItems

`func (o *LineItemResponseDataRelationships) SetShipmentLineItems(v LineItemResponseDataRelationshipsShipmentLineItems)`

SetShipmentLineItems sets ShipmentLineItems field to given value.

### HasShipmentLineItems

`func (o *LineItemResponseDataRelationships) HasShipmentLineItems() bool`

HasShipmentLineItems returns a boolean if a field has been set.

### GetStockLineItems

`func (o *LineItemResponseDataRelationships) GetStockLineItems() LineItemResponseDataRelationshipsStockLineItems`

GetStockLineItems returns the StockLineItems field if non-nil, zero value otherwise.

### GetStockLineItemsOk

`func (o *LineItemResponseDataRelationships) GetStockLineItemsOk() (*LineItemResponseDataRelationshipsStockLineItems, bool)`

GetStockLineItemsOk returns a tuple with the StockLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLineItems

`func (o *LineItemResponseDataRelationships) SetStockLineItems(v LineItemResponseDataRelationshipsStockLineItems)`

SetStockLineItems sets StockLineItems field to given value.

### HasStockLineItems

`func (o *LineItemResponseDataRelationships) HasStockLineItems() bool`

HasStockLineItems returns a boolean if a field has been set.

### GetStockTransfers

`func (o *LineItemResponseDataRelationships) GetStockTransfers() LineItemResponseDataRelationshipsStockTransfers`

GetStockTransfers returns the StockTransfers field if non-nil, zero value otherwise.

### GetStockTransfersOk

`func (o *LineItemResponseDataRelationships) GetStockTransfersOk() (*LineItemResponseDataRelationshipsStockTransfers, bool)`

GetStockTransfersOk returns a tuple with the StockTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockTransfers

`func (o *LineItemResponseDataRelationships) SetStockTransfers(v LineItemResponseDataRelationshipsStockTransfers)`

SetStockTransfers sets StockTransfers field to given value.

### HasStockTransfers

`func (o *LineItemResponseDataRelationships) HasStockTransfers() bool`

HasStockTransfers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


