# StockLineItemCreateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shipment** | Pointer to [**ParcelCreateDataRelationshipsShipment**](ParcelCreateDataRelationshipsShipment.md) |  | [optional] 
**LineItem** | Pointer to [**LineItemOptionCreateDataRelationshipsLineItem**](LineItemOptionCreateDataRelationshipsLineItem.md) |  | [optional] 
**StockItem** | Pointer to [**StockLineItemCreateDataRelationshipsStockItem**](StockLineItemCreateDataRelationshipsStockItem.md) |  | [optional] 
**Sku** | Pointer to [**InStockSubscriptionCreateDataRelationshipsSku**](InStockSubscriptionCreateDataRelationshipsSku.md) |  | [optional] 

## Methods

### NewStockLineItemCreateDataRelationships

`func NewStockLineItemCreateDataRelationships() *StockLineItemCreateDataRelationships`

NewStockLineItemCreateDataRelationships instantiates a new StockLineItemCreateDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockLineItemCreateDataRelationshipsWithDefaults

`func NewStockLineItemCreateDataRelationshipsWithDefaults() *StockLineItemCreateDataRelationships`

NewStockLineItemCreateDataRelationshipsWithDefaults instantiates a new StockLineItemCreateDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipment

`func (o *StockLineItemCreateDataRelationships) GetShipment() ParcelCreateDataRelationshipsShipment`

GetShipment returns the Shipment field if non-nil, zero value otherwise.

### GetShipmentOk

`func (o *StockLineItemCreateDataRelationships) GetShipmentOk() (*ParcelCreateDataRelationshipsShipment, bool)`

GetShipmentOk returns a tuple with the Shipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipment

`func (o *StockLineItemCreateDataRelationships) SetShipment(v ParcelCreateDataRelationshipsShipment)`

SetShipment sets Shipment field to given value.

### HasShipment

`func (o *StockLineItemCreateDataRelationships) HasShipment() bool`

HasShipment returns a boolean if a field has been set.

### GetLineItem

`func (o *StockLineItemCreateDataRelationships) GetLineItem() LineItemOptionCreateDataRelationshipsLineItem`

GetLineItem returns the LineItem field if non-nil, zero value otherwise.

### GetLineItemOk

`func (o *StockLineItemCreateDataRelationships) GetLineItemOk() (*LineItemOptionCreateDataRelationshipsLineItem, bool)`

GetLineItemOk returns a tuple with the LineItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItem

`func (o *StockLineItemCreateDataRelationships) SetLineItem(v LineItemOptionCreateDataRelationshipsLineItem)`

SetLineItem sets LineItem field to given value.

### HasLineItem

`func (o *StockLineItemCreateDataRelationships) HasLineItem() bool`

HasLineItem returns a boolean if a field has been set.

### GetStockItem

`func (o *StockLineItemCreateDataRelationships) GetStockItem() StockLineItemCreateDataRelationshipsStockItem`

GetStockItem returns the StockItem field if non-nil, zero value otherwise.

### GetStockItemOk

`func (o *StockLineItemCreateDataRelationships) GetStockItemOk() (*StockLineItemCreateDataRelationshipsStockItem, bool)`

GetStockItemOk returns a tuple with the StockItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockItem

`func (o *StockLineItemCreateDataRelationships) SetStockItem(v StockLineItemCreateDataRelationshipsStockItem)`

SetStockItem sets StockItem field to given value.

### HasStockItem

`func (o *StockLineItemCreateDataRelationships) HasStockItem() bool`

HasStockItem returns a boolean if a field has been set.

### GetSku

`func (o *StockLineItemCreateDataRelationships) GetSku() InStockSubscriptionCreateDataRelationshipsSku`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *StockLineItemCreateDataRelationships) GetSkuOk() (*InStockSubscriptionCreateDataRelationshipsSku, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *StockLineItemCreateDataRelationships) SetSku(v InStockSubscriptionCreateDataRelationshipsSku)`

SetSku sets Sku field to given value.

### HasSku

`func (o *StockLineItemCreateDataRelationships) HasSku() bool`

HasSku returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


