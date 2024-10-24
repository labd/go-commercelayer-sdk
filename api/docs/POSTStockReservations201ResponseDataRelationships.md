# POSTStockReservations201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineItem** | Pointer to [**POSTLineItemOptions201ResponseDataRelationshipsLineItem**](POSTLineItemOptions201ResponseDataRelationshipsLineItem.md) |  | [optional] 
**Order** | Pointer to [**POSTAdyenPayments201ResponseDataRelationshipsOrder**](POSTAdyenPayments201ResponseDataRelationshipsOrder.md) |  | [optional] 
**StockLineItem** | Pointer to [**POSTParcelLineItems201ResponseDataRelationshipsStockLineItem**](POSTParcelLineItems201ResponseDataRelationshipsStockLineItem.md) |  | [optional] 
**StockTransfer** | Pointer to [**POSTStockReservations201ResponseDataRelationshipsStockTransfer**](POSTStockReservations201ResponseDataRelationshipsStockTransfer.md) |  | [optional] 
**StockItem** | Pointer to [**GETReservedStocksReservedStockId200ResponseDataRelationshipsStockItem**](GETReservedStocksReservedStockId200ResponseDataRelationshipsStockItem.md) |  | [optional] 
**ReservedStock** | Pointer to [**POSTStockItems201ResponseDataRelationshipsReservedStock**](POSTStockItems201ResponseDataRelationshipsReservedStock.md) |  | [optional] 
**Sku** | Pointer to [**POSTInStockSubscriptions201ResponseDataRelationshipsSku**](POSTInStockSubscriptions201ResponseDataRelationshipsSku.md) |  | [optional] 

## Methods

### NewPOSTStockReservations201ResponseDataRelationships

`func NewPOSTStockReservations201ResponseDataRelationships() *POSTStockReservations201ResponseDataRelationships`

NewPOSTStockReservations201ResponseDataRelationships instantiates a new POSTStockReservations201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTStockReservations201ResponseDataRelationshipsWithDefaults

`func NewPOSTStockReservations201ResponseDataRelationshipsWithDefaults() *POSTStockReservations201ResponseDataRelationships`

NewPOSTStockReservations201ResponseDataRelationshipsWithDefaults instantiates a new POSTStockReservations201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLineItem

`func (o *POSTStockReservations201ResponseDataRelationships) GetLineItem() POSTLineItemOptions201ResponseDataRelationshipsLineItem`

GetLineItem returns the LineItem field if non-nil, zero value otherwise.

### GetLineItemOk

`func (o *POSTStockReservations201ResponseDataRelationships) GetLineItemOk() (*POSTLineItemOptions201ResponseDataRelationshipsLineItem, bool)`

GetLineItemOk returns a tuple with the LineItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItem

`func (o *POSTStockReservations201ResponseDataRelationships) SetLineItem(v POSTLineItemOptions201ResponseDataRelationshipsLineItem)`

SetLineItem sets LineItem field to given value.

### HasLineItem

`func (o *POSTStockReservations201ResponseDataRelationships) HasLineItem() bool`

HasLineItem returns a boolean if a field has been set.

### GetOrder

`func (o *POSTStockReservations201ResponseDataRelationships) GetOrder() POSTAdyenPayments201ResponseDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *POSTStockReservations201ResponseDataRelationships) GetOrderOk() (*POSTAdyenPayments201ResponseDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *POSTStockReservations201ResponseDataRelationships) SetOrder(v POSTAdyenPayments201ResponseDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *POSTStockReservations201ResponseDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetStockLineItem

`func (o *POSTStockReservations201ResponseDataRelationships) GetStockLineItem() POSTParcelLineItems201ResponseDataRelationshipsStockLineItem`

GetStockLineItem returns the StockLineItem field if non-nil, zero value otherwise.

### GetStockLineItemOk

`func (o *POSTStockReservations201ResponseDataRelationships) GetStockLineItemOk() (*POSTParcelLineItems201ResponseDataRelationshipsStockLineItem, bool)`

GetStockLineItemOk returns a tuple with the StockLineItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLineItem

`func (o *POSTStockReservations201ResponseDataRelationships) SetStockLineItem(v POSTParcelLineItems201ResponseDataRelationshipsStockLineItem)`

SetStockLineItem sets StockLineItem field to given value.

### HasStockLineItem

`func (o *POSTStockReservations201ResponseDataRelationships) HasStockLineItem() bool`

HasStockLineItem returns a boolean if a field has been set.

### GetStockTransfer

`func (o *POSTStockReservations201ResponseDataRelationships) GetStockTransfer() POSTStockReservations201ResponseDataRelationshipsStockTransfer`

GetStockTransfer returns the StockTransfer field if non-nil, zero value otherwise.

### GetStockTransferOk

`func (o *POSTStockReservations201ResponseDataRelationships) GetStockTransferOk() (*POSTStockReservations201ResponseDataRelationshipsStockTransfer, bool)`

GetStockTransferOk returns a tuple with the StockTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockTransfer

`func (o *POSTStockReservations201ResponseDataRelationships) SetStockTransfer(v POSTStockReservations201ResponseDataRelationshipsStockTransfer)`

SetStockTransfer sets StockTransfer field to given value.

### HasStockTransfer

`func (o *POSTStockReservations201ResponseDataRelationships) HasStockTransfer() bool`

HasStockTransfer returns a boolean if a field has been set.

### GetStockItem

`func (o *POSTStockReservations201ResponseDataRelationships) GetStockItem() GETReservedStocksReservedStockId200ResponseDataRelationshipsStockItem`

GetStockItem returns the StockItem field if non-nil, zero value otherwise.

### GetStockItemOk

`func (o *POSTStockReservations201ResponseDataRelationships) GetStockItemOk() (*GETReservedStocksReservedStockId200ResponseDataRelationshipsStockItem, bool)`

GetStockItemOk returns a tuple with the StockItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockItem

`func (o *POSTStockReservations201ResponseDataRelationships) SetStockItem(v GETReservedStocksReservedStockId200ResponseDataRelationshipsStockItem)`

SetStockItem sets StockItem field to given value.

### HasStockItem

`func (o *POSTStockReservations201ResponseDataRelationships) HasStockItem() bool`

HasStockItem returns a boolean if a field has been set.

### GetReservedStock

`func (o *POSTStockReservations201ResponseDataRelationships) GetReservedStock() POSTStockItems201ResponseDataRelationshipsReservedStock`

GetReservedStock returns the ReservedStock field if non-nil, zero value otherwise.

### GetReservedStockOk

`func (o *POSTStockReservations201ResponseDataRelationships) GetReservedStockOk() (*POSTStockItems201ResponseDataRelationshipsReservedStock, bool)`

GetReservedStockOk returns a tuple with the ReservedStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedStock

`func (o *POSTStockReservations201ResponseDataRelationships) SetReservedStock(v POSTStockItems201ResponseDataRelationshipsReservedStock)`

SetReservedStock sets ReservedStock field to given value.

### HasReservedStock

`func (o *POSTStockReservations201ResponseDataRelationships) HasReservedStock() bool`

HasReservedStock returns a boolean if a field has been set.

### GetSku

`func (o *POSTStockReservations201ResponseDataRelationships) GetSku() POSTInStockSubscriptions201ResponseDataRelationshipsSku`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *POSTStockReservations201ResponseDataRelationships) GetSkuOk() (*POSTInStockSubscriptions201ResponseDataRelationshipsSku, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *POSTStockReservations201ResponseDataRelationships) SetSku(v POSTInStockSubscriptions201ResponseDataRelationshipsSku)`

SetSku sets Sku field to given value.

### HasSku

`func (o *POSTStockReservations201ResponseDataRelationships) HasSku() bool`

HasSku returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


