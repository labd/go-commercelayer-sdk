# POSTStockTransfers201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | [**GETBundles200ResponseDataInnerRelationshipsSkus**](GETBundles200ResponseDataInnerRelationshipsSkus.md) |  | 
**OriginStockLocation** | [**GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation**](GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation.md) |  | 
**DestinationStockLocation** | [**GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation**](GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation.md) |  | 
**Shipment** | Pointer to [**GETOrders200ResponseDataInnerRelationshipsShipments**](GETOrders200ResponseDataInnerRelationshipsShipments.md) |  | [optional] 
**LineItem** | Pointer to [**GETLineItemOptions200ResponseDataInnerRelationshipsLineItem**](GETLineItemOptions200ResponseDataInnerRelationshipsLineItem.md) |  | [optional] 

## Methods

### NewPOSTStockTransfers201ResponseDataRelationships

`func NewPOSTStockTransfers201ResponseDataRelationships(sku GETBundles200ResponseDataInnerRelationshipsSkus, originStockLocation GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation, destinationStockLocation GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation, ) *POSTStockTransfers201ResponseDataRelationships`

NewPOSTStockTransfers201ResponseDataRelationships instantiates a new POSTStockTransfers201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTStockTransfers201ResponseDataRelationshipsWithDefaults

`func NewPOSTStockTransfers201ResponseDataRelationshipsWithDefaults() *POSTStockTransfers201ResponseDataRelationships`

NewPOSTStockTransfers201ResponseDataRelationshipsWithDefaults instantiates a new POSTStockTransfers201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *POSTStockTransfers201ResponseDataRelationships) GetSku() GETBundles200ResponseDataInnerRelationshipsSkus`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *POSTStockTransfers201ResponseDataRelationships) GetSkuOk() (*GETBundles200ResponseDataInnerRelationshipsSkus, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *POSTStockTransfers201ResponseDataRelationships) SetSku(v GETBundles200ResponseDataInnerRelationshipsSkus)`

SetSku sets Sku field to given value.


### GetOriginStockLocation

`func (o *POSTStockTransfers201ResponseDataRelationships) GetOriginStockLocation() GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation`

GetOriginStockLocation returns the OriginStockLocation field if non-nil, zero value otherwise.

### GetOriginStockLocationOk

`func (o *POSTStockTransfers201ResponseDataRelationships) GetOriginStockLocationOk() (*GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation, bool)`

GetOriginStockLocationOk returns a tuple with the OriginStockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStockLocation

`func (o *POSTStockTransfers201ResponseDataRelationships) SetOriginStockLocation(v GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation)`

SetOriginStockLocation sets OriginStockLocation field to given value.


### GetDestinationStockLocation

`func (o *POSTStockTransfers201ResponseDataRelationships) GetDestinationStockLocation() GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation`

GetDestinationStockLocation returns the DestinationStockLocation field if non-nil, zero value otherwise.

### GetDestinationStockLocationOk

`func (o *POSTStockTransfers201ResponseDataRelationships) GetDestinationStockLocationOk() (*GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation, bool)`

GetDestinationStockLocationOk returns a tuple with the DestinationStockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationStockLocation

`func (o *POSTStockTransfers201ResponseDataRelationships) SetDestinationStockLocation(v GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation)`

SetDestinationStockLocation sets DestinationStockLocation field to given value.


### GetShipment

`func (o *POSTStockTransfers201ResponseDataRelationships) GetShipment() GETOrders200ResponseDataInnerRelationshipsShipments`

GetShipment returns the Shipment field if non-nil, zero value otherwise.

### GetShipmentOk

`func (o *POSTStockTransfers201ResponseDataRelationships) GetShipmentOk() (*GETOrders200ResponseDataInnerRelationshipsShipments, bool)`

GetShipmentOk returns a tuple with the Shipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipment

`func (o *POSTStockTransfers201ResponseDataRelationships) SetShipment(v GETOrders200ResponseDataInnerRelationshipsShipments)`

SetShipment sets Shipment field to given value.

### HasShipment

`func (o *POSTStockTransfers201ResponseDataRelationships) HasShipment() bool`

HasShipment returns a boolean if a field has been set.

### GetLineItem

`func (o *POSTStockTransfers201ResponseDataRelationships) GetLineItem() GETLineItemOptions200ResponseDataInnerRelationshipsLineItem`

GetLineItem returns the LineItem field if non-nil, zero value otherwise.

### GetLineItemOk

`func (o *POSTStockTransfers201ResponseDataRelationships) GetLineItemOk() (*GETLineItemOptions200ResponseDataInnerRelationshipsLineItem, bool)`

GetLineItemOk returns a tuple with the LineItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItem

`func (o *POSTStockTransfers201ResponseDataRelationships) SetLineItem(v GETLineItemOptions200ResponseDataInnerRelationshipsLineItem)`

SetLineItem sets LineItem field to given value.

### HasLineItem

`func (o *POSTStockTransfers201ResponseDataRelationships) HasLineItem() bool`

HasLineItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


