# StockTransferCreateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | [**InStockSubscriptionCreateDataRelationshipsSku**](InStockSubscriptionCreateDataRelationshipsSku.md) |  | 
**OriginStockLocation** | [**DeliveryLeadTimeCreateDataRelationshipsStockLocation**](DeliveryLeadTimeCreateDataRelationshipsStockLocation.md) |  | 
**DestinationStockLocation** | [**DeliveryLeadTimeCreateDataRelationshipsStockLocation**](DeliveryLeadTimeCreateDataRelationshipsStockLocation.md) |  | 
**Shipment** | Pointer to [**ParcelCreateDataRelationshipsShipment**](ParcelCreateDataRelationshipsShipment.md) |  | [optional] 
**LineItem** | Pointer to [**LineItemOptionCreateDataRelationshipsLineItem**](LineItemOptionCreateDataRelationshipsLineItem.md) |  | [optional] 

## Methods

### NewStockTransferCreateDataRelationships

`func NewStockTransferCreateDataRelationships(sku InStockSubscriptionCreateDataRelationshipsSku, originStockLocation DeliveryLeadTimeCreateDataRelationshipsStockLocation, destinationStockLocation DeliveryLeadTimeCreateDataRelationshipsStockLocation, ) *StockTransferCreateDataRelationships`

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

`func (o *StockTransferCreateDataRelationships) GetSku() InStockSubscriptionCreateDataRelationshipsSku`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *StockTransferCreateDataRelationships) GetSkuOk() (*InStockSubscriptionCreateDataRelationshipsSku, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *StockTransferCreateDataRelationships) SetSku(v InStockSubscriptionCreateDataRelationshipsSku)`

SetSku sets Sku field to given value.


### GetOriginStockLocation

`func (o *StockTransferCreateDataRelationships) GetOriginStockLocation() DeliveryLeadTimeCreateDataRelationshipsStockLocation`

GetOriginStockLocation returns the OriginStockLocation field if non-nil, zero value otherwise.

### GetOriginStockLocationOk

`func (o *StockTransferCreateDataRelationships) GetOriginStockLocationOk() (*DeliveryLeadTimeCreateDataRelationshipsStockLocation, bool)`

GetOriginStockLocationOk returns a tuple with the OriginStockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStockLocation

`func (o *StockTransferCreateDataRelationships) SetOriginStockLocation(v DeliveryLeadTimeCreateDataRelationshipsStockLocation)`

SetOriginStockLocation sets OriginStockLocation field to given value.


### GetDestinationStockLocation

`func (o *StockTransferCreateDataRelationships) GetDestinationStockLocation() DeliveryLeadTimeCreateDataRelationshipsStockLocation`

GetDestinationStockLocation returns the DestinationStockLocation field if non-nil, zero value otherwise.

### GetDestinationStockLocationOk

`func (o *StockTransferCreateDataRelationships) GetDestinationStockLocationOk() (*DeliveryLeadTimeCreateDataRelationshipsStockLocation, bool)`

GetDestinationStockLocationOk returns a tuple with the DestinationStockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationStockLocation

`func (o *StockTransferCreateDataRelationships) SetDestinationStockLocation(v DeliveryLeadTimeCreateDataRelationshipsStockLocation)`

SetDestinationStockLocation sets DestinationStockLocation field to given value.


### GetShipment

`func (o *StockTransferCreateDataRelationships) GetShipment() ParcelCreateDataRelationshipsShipment`

GetShipment returns the Shipment field if non-nil, zero value otherwise.

### GetShipmentOk

`func (o *StockTransferCreateDataRelationships) GetShipmentOk() (*ParcelCreateDataRelationshipsShipment, bool)`

GetShipmentOk returns a tuple with the Shipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipment

`func (o *StockTransferCreateDataRelationships) SetShipment(v ParcelCreateDataRelationshipsShipment)`

SetShipment sets Shipment field to given value.

### HasShipment

`func (o *StockTransferCreateDataRelationships) HasShipment() bool`

HasShipment returns a boolean if a field has been set.

### GetLineItem

`func (o *StockTransferCreateDataRelationships) GetLineItem() LineItemOptionCreateDataRelationshipsLineItem`

GetLineItem returns the LineItem field if non-nil, zero value otherwise.

### GetLineItemOk

`func (o *StockTransferCreateDataRelationships) GetLineItemOk() (*LineItemOptionCreateDataRelationshipsLineItem, bool)`

GetLineItemOk returns a tuple with the LineItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItem

`func (o *StockTransferCreateDataRelationships) SetLineItem(v LineItemOptionCreateDataRelationshipsLineItem)`

SetLineItem sets LineItem field to given value.

### HasLineItem

`func (o *StockTransferCreateDataRelationships) HasLineItem() bool`

HasLineItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


