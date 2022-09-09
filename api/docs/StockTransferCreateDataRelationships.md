# StockTransferCreateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | [**BundleDataRelationshipsSkus**](BundleDataRelationshipsSkus.md) |  | 
**OriginStockLocation** | [**DeliveryLeadTimeDataRelationshipsStockLocation**](DeliveryLeadTimeDataRelationshipsStockLocation.md) |  | 
**DestinationStockLocation** | [**DeliveryLeadTimeDataRelationshipsStockLocation**](DeliveryLeadTimeDataRelationshipsStockLocation.md) |  | 
**Shipment** | Pointer to [**OrderDataRelationshipsShipments**](OrderDataRelationshipsShipments.md) |  | [optional] 
**LineItem** | Pointer to [**LineItemOptionDataRelationshipsLineItem**](LineItemOptionDataRelationshipsLineItem.md) |  | [optional] 

## Methods

### NewStockTransferCreateDataRelationships

`func NewStockTransferCreateDataRelationships(sku BundleDataRelationshipsSkus, originStockLocation DeliveryLeadTimeDataRelationshipsStockLocation, destinationStockLocation DeliveryLeadTimeDataRelationshipsStockLocation, ) *StockTransferCreateDataRelationships`

NewStockTransferCreateDataRelationships instantiates a new StockTransferCreateDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockTransferCreateDataRelationshipsWithDefaults

`func NewStockTransferCreateDataRelationshipsWithDefaults() *StockTransferCreateDataRelationships`

NewStockTransferCreateDataRelationshipsWithDefaults instantiates a new StockTransferCreateDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *StockTransferCreateDataRelationships) GetSku() BundleDataRelationshipsSkus`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *StockTransferCreateDataRelationships) GetSkuOk() (*BundleDataRelationshipsSkus, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *StockTransferCreateDataRelationships) SetSku(v BundleDataRelationshipsSkus)`

SetSku sets Sku field to given value.


### GetOriginStockLocation

`func (o *StockTransferCreateDataRelationships) GetOriginStockLocation() DeliveryLeadTimeDataRelationshipsStockLocation`

GetOriginStockLocation returns the OriginStockLocation field if non-nil, zero value otherwise.

### GetOriginStockLocationOk

`func (o *StockTransferCreateDataRelationships) GetOriginStockLocationOk() (*DeliveryLeadTimeDataRelationshipsStockLocation, bool)`

GetOriginStockLocationOk returns a tuple with the OriginStockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStockLocation

`func (o *StockTransferCreateDataRelationships) SetOriginStockLocation(v DeliveryLeadTimeDataRelationshipsStockLocation)`

SetOriginStockLocation sets OriginStockLocation field to given value.


### GetDestinationStockLocation

`func (o *StockTransferCreateDataRelationships) GetDestinationStockLocation() DeliveryLeadTimeDataRelationshipsStockLocation`

GetDestinationStockLocation returns the DestinationStockLocation field if non-nil, zero value otherwise.

### GetDestinationStockLocationOk

`func (o *StockTransferCreateDataRelationships) GetDestinationStockLocationOk() (*DeliveryLeadTimeDataRelationshipsStockLocation, bool)`

GetDestinationStockLocationOk returns a tuple with the DestinationStockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationStockLocation

`func (o *StockTransferCreateDataRelationships) SetDestinationStockLocation(v DeliveryLeadTimeDataRelationshipsStockLocation)`

SetDestinationStockLocation sets DestinationStockLocation field to given value.


### GetShipment

`func (o *StockTransferCreateDataRelationships) GetShipment() OrderDataRelationshipsShipments`

GetShipment returns the Shipment field if non-nil, zero value otherwise.

### GetShipmentOk

`func (o *StockTransferCreateDataRelationships) GetShipmentOk() (*OrderDataRelationshipsShipments, bool)`

GetShipmentOk returns a tuple with the Shipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipment

`func (o *StockTransferCreateDataRelationships) SetShipment(v OrderDataRelationshipsShipments)`

SetShipment sets Shipment field to given value.

### HasShipment

`func (o *StockTransferCreateDataRelationships) HasShipment() bool`

HasShipment returns a boolean if a field has been set.

### GetLineItem

`func (o *StockTransferCreateDataRelationships) GetLineItem() LineItemOptionDataRelationshipsLineItem`

GetLineItem returns the LineItem field if non-nil, zero value otherwise.

### GetLineItemOk

`func (o *StockTransferCreateDataRelationships) GetLineItemOk() (*LineItemOptionDataRelationshipsLineItem, bool)`

GetLineItemOk returns a tuple with the LineItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItem

`func (o *StockTransferCreateDataRelationships) SetLineItem(v LineItemOptionDataRelationshipsLineItem)`

SetLineItem sets LineItem field to given value.

### HasLineItem

`func (o *StockTransferCreateDataRelationships) HasLineItem() bool`

HasLineItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


