# StockReservationDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineItem** | Pointer to [**LineItemOptionDataRelationshipsLineItem**](LineItemOptionDataRelationshipsLineItem.md) |  | [optional] 
**Order** | Pointer to [**AdyenPaymentDataRelationshipsOrder**](AdyenPaymentDataRelationshipsOrder.md) |  | [optional] 
**StockLineItem** | Pointer to [**LineItemDataRelationshipsStockLineItems**](LineItemDataRelationshipsStockLineItems.md) |  | [optional] 
**StockTransfer** | Pointer to [**LineItemDataRelationshipsStockTransfers**](LineItemDataRelationshipsStockTransfers.md) |  | [optional] 
**StockItem** | Pointer to [**ReservedStockDataRelationshipsStockItem**](ReservedStockDataRelationshipsStockItem.md) |  | [optional] 
**ReservedStock** | Pointer to [**StockItemDataRelationshipsReservedStock**](StockItemDataRelationshipsReservedStock.md) |  | [optional] 
**Sku** | Pointer to [**BundleDataRelationshipsSkus**](BundleDataRelationshipsSkus.md) |  | [optional] 

## Methods

### NewStockReservationDataRelationships

`func NewStockReservationDataRelationships() *StockReservationDataRelationships`

NewStockReservationDataRelationships instantiates a new StockReservationDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockReservationDataRelationshipsWithDefaults

`func NewStockReservationDataRelationshipsWithDefaults() *StockReservationDataRelationships`

NewStockReservationDataRelationshipsWithDefaults instantiates a new StockReservationDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLineItem

`func (o *StockReservationDataRelationships) GetLineItem() LineItemOptionDataRelationshipsLineItem`

GetLineItem returns the LineItem field if non-nil, zero value otherwise.

### GetLineItemOk

`func (o *StockReservationDataRelationships) GetLineItemOk() (*LineItemOptionDataRelationshipsLineItem, bool)`

GetLineItemOk returns a tuple with the LineItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItem

`func (o *StockReservationDataRelationships) SetLineItem(v LineItemOptionDataRelationshipsLineItem)`

SetLineItem sets LineItem field to given value.

### HasLineItem

`func (o *StockReservationDataRelationships) HasLineItem() bool`

HasLineItem returns a boolean if a field has been set.

### GetOrder

`func (o *StockReservationDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *StockReservationDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *StockReservationDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *StockReservationDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetStockLineItem

`func (o *StockReservationDataRelationships) GetStockLineItem() LineItemDataRelationshipsStockLineItems`

GetStockLineItem returns the StockLineItem field if non-nil, zero value otherwise.

### GetStockLineItemOk

`func (o *StockReservationDataRelationships) GetStockLineItemOk() (*LineItemDataRelationshipsStockLineItems, bool)`

GetStockLineItemOk returns a tuple with the StockLineItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLineItem

`func (o *StockReservationDataRelationships) SetStockLineItem(v LineItemDataRelationshipsStockLineItems)`

SetStockLineItem sets StockLineItem field to given value.

### HasStockLineItem

`func (o *StockReservationDataRelationships) HasStockLineItem() bool`

HasStockLineItem returns a boolean if a field has been set.

### GetStockTransfer

`func (o *StockReservationDataRelationships) GetStockTransfer() LineItemDataRelationshipsStockTransfers`

GetStockTransfer returns the StockTransfer field if non-nil, zero value otherwise.

### GetStockTransferOk

`func (o *StockReservationDataRelationships) GetStockTransferOk() (*LineItemDataRelationshipsStockTransfers, bool)`

GetStockTransferOk returns a tuple with the StockTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockTransfer

`func (o *StockReservationDataRelationships) SetStockTransfer(v LineItemDataRelationshipsStockTransfers)`

SetStockTransfer sets StockTransfer field to given value.

### HasStockTransfer

`func (o *StockReservationDataRelationships) HasStockTransfer() bool`

HasStockTransfer returns a boolean if a field has been set.

### GetStockItem

`func (o *StockReservationDataRelationships) GetStockItem() ReservedStockDataRelationshipsStockItem`

GetStockItem returns the StockItem field if non-nil, zero value otherwise.

### GetStockItemOk

`func (o *StockReservationDataRelationships) GetStockItemOk() (*ReservedStockDataRelationshipsStockItem, bool)`

GetStockItemOk returns a tuple with the StockItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockItem

`func (o *StockReservationDataRelationships) SetStockItem(v ReservedStockDataRelationshipsStockItem)`

SetStockItem sets StockItem field to given value.

### HasStockItem

`func (o *StockReservationDataRelationships) HasStockItem() bool`

HasStockItem returns a boolean if a field has been set.

### GetReservedStock

`func (o *StockReservationDataRelationships) GetReservedStock() StockItemDataRelationshipsReservedStock`

GetReservedStock returns the ReservedStock field if non-nil, zero value otherwise.

### GetReservedStockOk

`func (o *StockReservationDataRelationships) GetReservedStockOk() (*StockItemDataRelationshipsReservedStock, bool)`

GetReservedStockOk returns a tuple with the ReservedStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedStock

`func (o *StockReservationDataRelationships) SetReservedStock(v StockItemDataRelationshipsReservedStock)`

SetReservedStock sets ReservedStock field to given value.

### HasReservedStock

`func (o *StockReservationDataRelationships) HasReservedStock() bool`

HasReservedStock returns a boolean if a field has been set.

### GetSku

`func (o *StockReservationDataRelationships) GetSku() BundleDataRelationshipsSkus`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *StockReservationDataRelationships) GetSkuOk() (*BundleDataRelationshipsSkus, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *StockReservationDataRelationships) SetSku(v BundleDataRelationshipsSkus)`

SetSku sets Sku field to given value.

### HasSku

`func (o *StockReservationDataRelationships) HasSku() bool`

HasSku returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


