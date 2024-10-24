# StockLineItemDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shipment** | Pointer to [**LineItemDataRelationshipsShipment**](LineItemDataRelationshipsShipment.md) |  | [optional] 
**LineItem** | Pointer to [**LineItemOptionDataRelationshipsLineItem**](LineItemOptionDataRelationshipsLineItem.md) |  | [optional] 
**StockItem** | Pointer to [**ReservedStockDataRelationshipsStockItem**](ReservedStockDataRelationshipsStockItem.md) |  | [optional] 
**Sku** | Pointer to [**BundleDataRelationshipsSkus**](BundleDataRelationshipsSkus.md) |  | [optional] 
**StockReservation** | Pointer to [**LineItemDataRelationshipsStockReservations**](LineItemDataRelationshipsStockReservations.md) |  | [optional] 
**Versions** | Pointer to [**AddressDataRelationshipsVersions**](AddressDataRelationshipsVersions.md) |  | [optional] 

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

`func (o *StockLineItemDataRelationships) GetShipment() LineItemDataRelationshipsShipment`

GetShipment returns the Shipment field if non-nil, zero value otherwise.

### GetShipmentOk

`func (o *StockLineItemDataRelationships) GetShipmentOk() (*LineItemDataRelationshipsShipment, bool)`

GetShipmentOk returns a tuple with the Shipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipment

`func (o *StockLineItemDataRelationships) SetShipment(v LineItemDataRelationshipsShipment)`

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

`func (o *StockLineItemDataRelationships) GetStockItem() ReservedStockDataRelationshipsStockItem`

GetStockItem returns the StockItem field if non-nil, zero value otherwise.

### GetStockItemOk

`func (o *StockLineItemDataRelationships) GetStockItemOk() (*ReservedStockDataRelationshipsStockItem, bool)`

GetStockItemOk returns a tuple with the StockItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockItem

`func (o *StockLineItemDataRelationships) SetStockItem(v ReservedStockDataRelationshipsStockItem)`

SetStockItem sets StockItem field to given value.

### HasStockItem

`func (o *StockLineItemDataRelationships) HasStockItem() bool`

HasStockItem returns a boolean if a field has been set.

### GetSku

`func (o *StockLineItemDataRelationships) GetSku() BundleDataRelationshipsSkus`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *StockLineItemDataRelationships) GetSkuOk() (*BundleDataRelationshipsSkus, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *StockLineItemDataRelationships) SetSku(v BundleDataRelationshipsSkus)`

SetSku sets Sku field to given value.

### HasSku

`func (o *StockLineItemDataRelationships) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetStockReservation

`func (o *StockLineItemDataRelationships) GetStockReservation() LineItemDataRelationshipsStockReservations`

GetStockReservation returns the StockReservation field if non-nil, zero value otherwise.

### GetStockReservationOk

`func (o *StockLineItemDataRelationships) GetStockReservationOk() (*LineItemDataRelationshipsStockReservations, bool)`

GetStockReservationOk returns a tuple with the StockReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockReservation

`func (o *StockLineItemDataRelationships) SetStockReservation(v LineItemDataRelationshipsStockReservations)`

SetStockReservation sets StockReservation field to given value.

### HasStockReservation

`func (o *StockLineItemDataRelationships) HasStockReservation() bool`

HasStockReservation returns a boolean if a field has been set.

### GetVersions

`func (o *StockLineItemDataRelationships) GetVersions() AddressDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *StockLineItemDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *StockLineItemDataRelationships) SetVersions(v AddressDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *StockLineItemDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


