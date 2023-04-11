# POSTStockTransfersRequestDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | [**POSTInStockSubscriptionsRequestDataRelationshipsSku**](POSTInStockSubscriptionsRequestDataRelationshipsSku.md) |  | 
**OriginStockLocation** | [**POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation**](POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation.md) |  | 
**DestinationStockLocation** | [**POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation**](POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation.md) |  | 
**Shipment** | Pointer to [**POSTParcelsRequestDataRelationshipsShipment**](POSTParcelsRequestDataRelationshipsShipment.md) |  | [optional] 
**LineItem** | Pointer to [**POSTLineItemOptionsRequestDataRelationshipsLineItem**](POSTLineItemOptionsRequestDataRelationshipsLineItem.md) |  | [optional] 

## Methods

### NewPOSTStockTransfersRequestDataRelationships

`func NewPOSTStockTransfersRequestDataRelationships(sku POSTInStockSubscriptionsRequestDataRelationshipsSku, originStockLocation POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation, destinationStockLocation POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation, ) *POSTStockTransfersRequestDataRelationships`

NewPOSTStockTransfersRequestDataRelationships instantiates a new POSTStockTransfersRequestDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTStockTransfersRequestDataRelationshipsWithDefaults

`func NewPOSTStockTransfersRequestDataRelationshipsWithDefaults() *POSTStockTransfersRequestDataRelationships`

NewPOSTStockTransfersRequestDataRelationshipsWithDefaults instantiates a new POSTStockTransfersRequestDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *POSTStockTransfersRequestDataRelationships) GetSku() POSTInStockSubscriptionsRequestDataRelationshipsSku`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *POSTStockTransfersRequestDataRelationships) GetSkuOk() (*POSTInStockSubscriptionsRequestDataRelationshipsSku, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *POSTStockTransfersRequestDataRelationships) SetSku(v POSTInStockSubscriptionsRequestDataRelationshipsSku)`

SetSku sets Sku field to given value.


### GetOriginStockLocation

`func (o *POSTStockTransfersRequestDataRelationships) GetOriginStockLocation() POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation`

GetOriginStockLocation returns the OriginStockLocation field if non-nil, zero value otherwise.

### GetOriginStockLocationOk

`func (o *POSTStockTransfersRequestDataRelationships) GetOriginStockLocationOk() (*POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation, bool)`

GetOriginStockLocationOk returns a tuple with the OriginStockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStockLocation

`func (o *POSTStockTransfersRequestDataRelationships) SetOriginStockLocation(v POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation)`

SetOriginStockLocation sets OriginStockLocation field to given value.


### GetDestinationStockLocation

`func (o *POSTStockTransfersRequestDataRelationships) GetDestinationStockLocation() POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation`

GetDestinationStockLocation returns the DestinationStockLocation field if non-nil, zero value otherwise.

### GetDestinationStockLocationOk

`func (o *POSTStockTransfersRequestDataRelationships) GetDestinationStockLocationOk() (*POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation, bool)`

GetDestinationStockLocationOk returns a tuple with the DestinationStockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationStockLocation

`func (o *POSTStockTransfersRequestDataRelationships) SetDestinationStockLocation(v POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation)`

SetDestinationStockLocation sets DestinationStockLocation field to given value.


### GetShipment

`func (o *POSTStockTransfersRequestDataRelationships) GetShipment() POSTParcelsRequestDataRelationshipsShipment`

GetShipment returns the Shipment field if non-nil, zero value otherwise.

### GetShipmentOk

`func (o *POSTStockTransfersRequestDataRelationships) GetShipmentOk() (*POSTParcelsRequestDataRelationshipsShipment, bool)`

GetShipmentOk returns a tuple with the Shipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipment

`func (o *POSTStockTransfersRequestDataRelationships) SetShipment(v POSTParcelsRequestDataRelationshipsShipment)`

SetShipment sets Shipment field to given value.

### HasShipment

`func (o *POSTStockTransfersRequestDataRelationships) HasShipment() bool`

HasShipment returns a boolean if a field has been set.

### GetLineItem

`func (o *POSTStockTransfersRequestDataRelationships) GetLineItem() POSTLineItemOptionsRequestDataRelationshipsLineItem`

GetLineItem returns the LineItem field if non-nil, zero value otherwise.

### GetLineItemOk

`func (o *POSTStockTransfersRequestDataRelationships) GetLineItemOk() (*POSTLineItemOptionsRequestDataRelationshipsLineItem, bool)`

GetLineItemOk returns a tuple with the LineItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItem

`func (o *POSTStockTransfersRequestDataRelationships) SetLineItem(v POSTLineItemOptionsRequestDataRelationshipsLineItem)`

SetLineItem sets LineItem field to given value.

### HasLineItem

`func (o *POSTStockTransfersRequestDataRelationships) HasLineItem() bool`

HasLineItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


