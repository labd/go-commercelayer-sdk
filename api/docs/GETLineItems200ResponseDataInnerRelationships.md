# GETLineItems200ResponseDataInnerRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**GETAddresses200ResponseDataInnerRelationshipsGeocoder**](GETAddresses200ResponseDataInnerRelationshipsGeocoder.md) |  | [optional] 
**Item** | Pointer to [**GETAddresses200ResponseDataInnerRelationshipsGeocoder**](GETAddresses200ResponseDataInnerRelationshipsGeocoder.md) |  | [optional] 
**LineItemOptions** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**ShipmentLineItems** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**StockLineItems** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**StockTransfers** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 

## Methods

### NewGETLineItems200ResponseDataInnerRelationships

`func NewGETLineItems200ResponseDataInnerRelationships() *GETLineItems200ResponseDataInnerRelationships`

NewGETLineItems200ResponseDataInnerRelationships instantiates a new GETLineItems200ResponseDataInnerRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETLineItems200ResponseDataInnerRelationshipsWithDefaults

`func NewGETLineItems200ResponseDataInnerRelationshipsWithDefaults() *GETLineItems200ResponseDataInnerRelationships`

NewGETLineItems200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETLineItems200ResponseDataInnerRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *GETLineItems200ResponseDataInnerRelationships) GetOrder() GETAddresses200ResponseDataInnerRelationshipsGeocoder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GETLineItems200ResponseDataInnerRelationships) GetOrderOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GETLineItems200ResponseDataInnerRelationships) SetOrder(v GETAddresses200ResponseDataInnerRelationshipsGeocoder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GETLineItems200ResponseDataInnerRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetItem

`func (o *GETLineItems200ResponseDataInnerRelationships) GetItem() GETAddresses200ResponseDataInnerRelationshipsGeocoder`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *GETLineItems200ResponseDataInnerRelationships) GetItemOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *GETLineItems200ResponseDataInnerRelationships) SetItem(v GETAddresses200ResponseDataInnerRelationshipsGeocoder)`

SetItem sets Item field to given value.

### HasItem

`func (o *GETLineItems200ResponseDataInnerRelationships) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetLineItemOptions

`func (o *GETLineItems200ResponseDataInnerRelationships) GetLineItemOptions() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetLineItemOptions returns the LineItemOptions field if non-nil, zero value otherwise.

### GetLineItemOptionsOk

`func (o *GETLineItems200ResponseDataInnerRelationships) GetLineItemOptionsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetLineItemOptionsOk returns a tuple with the LineItemOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItemOptions

`func (o *GETLineItems200ResponseDataInnerRelationships) SetLineItemOptions(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetLineItemOptions sets LineItemOptions field to given value.

### HasLineItemOptions

`func (o *GETLineItems200ResponseDataInnerRelationships) HasLineItemOptions() bool`

HasLineItemOptions returns a boolean if a field has been set.

### GetShipmentLineItems

`func (o *GETLineItems200ResponseDataInnerRelationships) GetShipmentLineItems() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetShipmentLineItems returns the ShipmentLineItems field if non-nil, zero value otherwise.

### GetShipmentLineItemsOk

`func (o *GETLineItems200ResponseDataInnerRelationships) GetShipmentLineItemsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetShipmentLineItemsOk returns a tuple with the ShipmentLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentLineItems

`func (o *GETLineItems200ResponseDataInnerRelationships) SetShipmentLineItems(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetShipmentLineItems sets ShipmentLineItems field to given value.

### HasShipmentLineItems

`func (o *GETLineItems200ResponseDataInnerRelationships) HasShipmentLineItems() bool`

HasShipmentLineItems returns a boolean if a field has been set.

### GetStockLineItems

`func (o *GETLineItems200ResponseDataInnerRelationships) GetStockLineItems() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetStockLineItems returns the StockLineItems field if non-nil, zero value otherwise.

### GetStockLineItemsOk

`func (o *GETLineItems200ResponseDataInnerRelationships) GetStockLineItemsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetStockLineItemsOk returns a tuple with the StockLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLineItems

`func (o *GETLineItems200ResponseDataInnerRelationships) SetStockLineItems(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetStockLineItems sets StockLineItems field to given value.

### HasStockLineItems

`func (o *GETLineItems200ResponseDataInnerRelationships) HasStockLineItems() bool`

HasStockLineItems returns a boolean if a field has been set.

### GetStockTransfers

`func (o *GETLineItems200ResponseDataInnerRelationships) GetStockTransfers() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetStockTransfers returns the StockTransfers field if non-nil, zero value otherwise.

### GetStockTransfersOk

`func (o *GETLineItems200ResponseDataInnerRelationships) GetStockTransfersOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetStockTransfersOk returns a tuple with the StockTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockTransfers

`func (o *GETLineItems200ResponseDataInnerRelationships) SetStockTransfers(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetStockTransfers sets StockTransfers field to given value.

### HasStockTransfers

`func (o *GETLineItems200ResponseDataInnerRelationships) HasStockTransfers() bool`

HasStockTransfers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


