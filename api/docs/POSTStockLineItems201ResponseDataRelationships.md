# POSTStockLineItems201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shipment** | Pointer to [**POSTLineItems201ResponseDataRelationshipsShipment**](POSTLineItems201ResponseDataRelationshipsShipment.md) |  | [optional] 
**LineItem** | Pointer to [**POSTLineItemOptions201ResponseDataRelationshipsLineItem**](POSTLineItemOptions201ResponseDataRelationshipsLineItem.md) |  | [optional] 
**StockItem** | Pointer to [**GETReservedStocksReservedStockId200ResponseDataRelationshipsStockItem**](GETReservedStocksReservedStockId200ResponseDataRelationshipsStockItem.md) |  | [optional] 
**Sku** | Pointer to [**POSTInStockSubscriptions201ResponseDataRelationshipsSku**](POSTInStockSubscriptions201ResponseDataRelationshipsSku.md) |  | [optional] 
**StockReservation** | Pointer to [**POSTStockLineItems201ResponseDataRelationshipsStockReservation**](POSTStockLineItems201ResponseDataRelationshipsStockReservation.md) |  | [optional] 
**Versions** | Pointer to [**POSTAddresses201ResponseDataRelationshipsVersions**](POSTAddresses201ResponseDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewPOSTStockLineItems201ResponseDataRelationships

`func NewPOSTStockLineItems201ResponseDataRelationships() *POSTStockLineItems201ResponseDataRelationships`

NewPOSTStockLineItems201ResponseDataRelationships instantiates a new POSTStockLineItems201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTStockLineItems201ResponseDataRelationshipsWithDefaults

`func NewPOSTStockLineItems201ResponseDataRelationshipsWithDefaults() *POSTStockLineItems201ResponseDataRelationships`

NewPOSTStockLineItems201ResponseDataRelationshipsWithDefaults instantiates a new POSTStockLineItems201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipment

`func (o *POSTStockLineItems201ResponseDataRelationships) GetShipment() POSTLineItems201ResponseDataRelationshipsShipment`

GetShipment returns the Shipment field if non-nil, zero value otherwise.

### GetShipmentOk

`func (o *POSTStockLineItems201ResponseDataRelationships) GetShipmentOk() (*POSTLineItems201ResponseDataRelationshipsShipment, bool)`

GetShipmentOk returns a tuple with the Shipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipment

`func (o *POSTStockLineItems201ResponseDataRelationships) SetShipment(v POSTLineItems201ResponseDataRelationshipsShipment)`

SetShipment sets Shipment field to given value.

### HasShipment

`func (o *POSTStockLineItems201ResponseDataRelationships) HasShipment() bool`

HasShipment returns a boolean if a field has been set.

### GetLineItem

`func (o *POSTStockLineItems201ResponseDataRelationships) GetLineItem() POSTLineItemOptions201ResponseDataRelationshipsLineItem`

GetLineItem returns the LineItem field if non-nil, zero value otherwise.

### GetLineItemOk

`func (o *POSTStockLineItems201ResponseDataRelationships) GetLineItemOk() (*POSTLineItemOptions201ResponseDataRelationshipsLineItem, bool)`

GetLineItemOk returns a tuple with the LineItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItem

`func (o *POSTStockLineItems201ResponseDataRelationships) SetLineItem(v POSTLineItemOptions201ResponseDataRelationshipsLineItem)`

SetLineItem sets LineItem field to given value.

### HasLineItem

`func (o *POSTStockLineItems201ResponseDataRelationships) HasLineItem() bool`

HasLineItem returns a boolean if a field has been set.

### GetStockItem

`func (o *POSTStockLineItems201ResponseDataRelationships) GetStockItem() GETReservedStocksReservedStockId200ResponseDataRelationshipsStockItem`

GetStockItem returns the StockItem field if non-nil, zero value otherwise.

### GetStockItemOk

`func (o *POSTStockLineItems201ResponseDataRelationships) GetStockItemOk() (*GETReservedStocksReservedStockId200ResponseDataRelationshipsStockItem, bool)`

GetStockItemOk returns a tuple with the StockItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockItem

`func (o *POSTStockLineItems201ResponseDataRelationships) SetStockItem(v GETReservedStocksReservedStockId200ResponseDataRelationshipsStockItem)`

SetStockItem sets StockItem field to given value.

### HasStockItem

`func (o *POSTStockLineItems201ResponseDataRelationships) HasStockItem() bool`

HasStockItem returns a boolean if a field has been set.

### GetSku

`func (o *POSTStockLineItems201ResponseDataRelationships) GetSku() POSTInStockSubscriptions201ResponseDataRelationshipsSku`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *POSTStockLineItems201ResponseDataRelationships) GetSkuOk() (*POSTInStockSubscriptions201ResponseDataRelationshipsSku, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *POSTStockLineItems201ResponseDataRelationships) SetSku(v POSTInStockSubscriptions201ResponseDataRelationshipsSku)`

SetSku sets Sku field to given value.

### HasSku

`func (o *POSTStockLineItems201ResponseDataRelationships) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetStockReservation

`func (o *POSTStockLineItems201ResponseDataRelationships) GetStockReservation() POSTStockLineItems201ResponseDataRelationshipsStockReservation`

GetStockReservation returns the StockReservation field if non-nil, zero value otherwise.

### GetStockReservationOk

`func (o *POSTStockLineItems201ResponseDataRelationships) GetStockReservationOk() (*POSTStockLineItems201ResponseDataRelationshipsStockReservation, bool)`

GetStockReservationOk returns a tuple with the StockReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockReservation

`func (o *POSTStockLineItems201ResponseDataRelationships) SetStockReservation(v POSTStockLineItems201ResponseDataRelationshipsStockReservation)`

SetStockReservation sets StockReservation field to given value.

### HasStockReservation

`func (o *POSTStockLineItems201ResponseDataRelationships) HasStockReservation() bool`

HasStockReservation returns a boolean if a field has been set.

### GetVersions

`func (o *POSTStockLineItems201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *POSTStockLineItems201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *POSTStockLineItems201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *POSTStockLineItems201ResponseDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


