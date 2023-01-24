# GETStockTransfers200ResponseDataInnerRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | Pointer to [**GETInStockSubscriptions200ResponseDataInnerRelationshipsSku**](GETInStockSubscriptions200ResponseDataInnerRelationshipsSku.md) |  | [optional] 
**OriginStockLocation** | Pointer to [**GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation**](GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation.md) |  | [optional] 
**DestinationStockLocation** | Pointer to [**GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocation**](GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocation.md) |  | [optional] 
**Shipment** | Pointer to [**GETParcels200ResponseDataInnerRelationshipsShipment**](GETParcels200ResponseDataInnerRelationshipsShipment.md) |  | [optional] 
**LineItem** | Pointer to [**GETLineItemOptions200ResponseDataInnerRelationshipsLineItem**](GETLineItemOptions200ResponseDataInnerRelationshipsLineItem.md) |  | [optional] 
**Events** | Pointer to [**GETCleanups200ResponseDataInnerRelationshipsEvents**](GETCleanups200ResponseDataInnerRelationshipsEvents.md) |  | [optional] 

## Methods

### NewGETStockTransfers200ResponseDataInnerRelationships

`func NewGETStockTransfers200ResponseDataInnerRelationships() *GETStockTransfers200ResponseDataInnerRelationships`

NewGETStockTransfers200ResponseDataInnerRelationships instantiates a new GETStockTransfers200ResponseDataInnerRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETStockTransfers200ResponseDataInnerRelationshipsWithDefaults

`func NewGETStockTransfers200ResponseDataInnerRelationshipsWithDefaults() *GETStockTransfers200ResponseDataInnerRelationships`

NewGETStockTransfers200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETStockTransfers200ResponseDataInnerRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *GETStockTransfers200ResponseDataInnerRelationships) GetSku() GETInStockSubscriptions200ResponseDataInnerRelationshipsSku`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *GETStockTransfers200ResponseDataInnerRelationships) GetSkuOk() (*GETInStockSubscriptions200ResponseDataInnerRelationshipsSku, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *GETStockTransfers200ResponseDataInnerRelationships) SetSku(v GETInStockSubscriptions200ResponseDataInnerRelationshipsSku)`

SetSku sets Sku field to given value.

### HasSku

`func (o *GETStockTransfers200ResponseDataInnerRelationships) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetOriginStockLocation

`func (o *GETStockTransfers200ResponseDataInnerRelationships) GetOriginStockLocation() GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation`

GetOriginStockLocation returns the OriginStockLocation field if non-nil, zero value otherwise.

### GetOriginStockLocationOk

`func (o *GETStockTransfers200ResponseDataInnerRelationships) GetOriginStockLocationOk() (*GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation, bool)`

GetOriginStockLocationOk returns a tuple with the OriginStockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStockLocation

`func (o *GETStockTransfers200ResponseDataInnerRelationships) SetOriginStockLocation(v GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation)`

SetOriginStockLocation sets OriginStockLocation field to given value.

### HasOriginStockLocation

`func (o *GETStockTransfers200ResponseDataInnerRelationships) HasOriginStockLocation() bool`

HasOriginStockLocation returns a boolean if a field has been set.

### GetDestinationStockLocation

`func (o *GETStockTransfers200ResponseDataInnerRelationships) GetDestinationStockLocation() GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocation`

GetDestinationStockLocation returns the DestinationStockLocation field if non-nil, zero value otherwise.

### GetDestinationStockLocationOk

`func (o *GETStockTransfers200ResponseDataInnerRelationships) GetDestinationStockLocationOk() (*GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocation, bool)`

GetDestinationStockLocationOk returns a tuple with the DestinationStockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationStockLocation

`func (o *GETStockTransfers200ResponseDataInnerRelationships) SetDestinationStockLocation(v GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocation)`

SetDestinationStockLocation sets DestinationStockLocation field to given value.

### HasDestinationStockLocation

`func (o *GETStockTransfers200ResponseDataInnerRelationships) HasDestinationStockLocation() bool`

HasDestinationStockLocation returns a boolean if a field has been set.

### GetShipment

`func (o *GETStockTransfers200ResponseDataInnerRelationships) GetShipment() GETParcels200ResponseDataInnerRelationshipsShipment`

GetShipment returns the Shipment field if non-nil, zero value otherwise.

### GetShipmentOk

`func (o *GETStockTransfers200ResponseDataInnerRelationships) GetShipmentOk() (*GETParcels200ResponseDataInnerRelationshipsShipment, bool)`

GetShipmentOk returns a tuple with the Shipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipment

`func (o *GETStockTransfers200ResponseDataInnerRelationships) SetShipment(v GETParcels200ResponseDataInnerRelationshipsShipment)`

SetShipment sets Shipment field to given value.

### HasShipment

`func (o *GETStockTransfers200ResponseDataInnerRelationships) HasShipment() bool`

HasShipment returns a boolean if a field has been set.

### GetLineItem

`func (o *GETStockTransfers200ResponseDataInnerRelationships) GetLineItem() GETLineItemOptions200ResponseDataInnerRelationshipsLineItem`

GetLineItem returns the LineItem field if non-nil, zero value otherwise.

### GetLineItemOk

`func (o *GETStockTransfers200ResponseDataInnerRelationships) GetLineItemOk() (*GETLineItemOptions200ResponseDataInnerRelationshipsLineItem, bool)`

GetLineItemOk returns a tuple with the LineItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItem

`func (o *GETStockTransfers200ResponseDataInnerRelationships) SetLineItem(v GETLineItemOptions200ResponseDataInnerRelationshipsLineItem)`

SetLineItem sets LineItem field to given value.

### HasLineItem

`func (o *GETStockTransfers200ResponseDataInnerRelationships) HasLineItem() bool`

HasLineItem returns a boolean if a field has been set.

### GetEvents

`func (o *GETStockTransfers200ResponseDataInnerRelationships) GetEvents() GETCleanups200ResponseDataInnerRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GETStockTransfers200ResponseDataInnerRelationships) GetEventsOk() (*GETCleanups200ResponseDataInnerRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GETStockTransfers200ResponseDataInnerRelationships) SetEvents(v GETCleanups200ResponseDataInnerRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *GETStockTransfers200ResponseDataInnerRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


