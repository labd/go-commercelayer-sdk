# POSTStockTransfers201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | Pointer to [**POSTInStockSubscriptions201ResponseDataRelationshipsSku**](POSTInStockSubscriptions201ResponseDataRelationshipsSku.md) |  | [optional] 
**OriginStockLocation** | Pointer to [**POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation**](POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation.md) |  | [optional] 
**DestinationStockLocation** | Pointer to [**POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation**](POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation.md) |  | [optional] 
**Shipment** | Pointer to [**POSTLineItems201ResponseDataRelationshipsShipment**](POSTLineItems201ResponseDataRelationshipsShipment.md) |  | [optional] 
**LineItem** | Pointer to [**POSTLineItemOptions201ResponseDataRelationshipsLineItem**](POSTLineItemOptions201ResponseDataRelationshipsLineItem.md) |  | [optional] 
**StockReservation** | Pointer to [**POSTStockLineItems201ResponseDataRelationshipsStockReservation**](POSTStockLineItems201ResponseDataRelationshipsStockReservation.md) |  | [optional] 
**Attachments** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments.md) |  | [optional] 
**Events** | Pointer to [**POSTAddresses201ResponseDataRelationshipsEvents**](POSTAddresses201ResponseDataRelationshipsEvents.md) |  | [optional] 
**Versions** | Pointer to [**POSTAddresses201ResponseDataRelationshipsVersions**](POSTAddresses201ResponseDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewPOSTStockTransfers201ResponseDataRelationships

`func NewPOSTStockTransfers201ResponseDataRelationships() *POSTStockTransfers201ResponseDataRelationships`

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

`func (o *POSTStockTransfers201ResponseDataRelationships) GetSku() POSTInStockSubscriptions201ResponseDataRelationshipsSku`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *POSTStockTransfers201ResponseDataRelationships) GetSkuOk() (*POSTInStockSubscriptions201ResponseDataRelationshipsSku, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *POSTStockTransfers201ResponseDataRelationships) SetSku(v POSTInStockSubscriptions201ResponseDataRelationshipsSku)`

SetSku sets Sku field to given value.

### HasSku

`func (o *POSTStockTransfers201ResponseDataRelationships) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetOriginStockLocation

`func (o *POSTStockTransfers201ResponseDataRelationships) GetOriginStockLocation() POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation`

GetOriginStockLocation returns the OriginStockLocation field if non-nil, zero value otherwise.

### GetOriginStockLocationOk

`func (o *POSTStockTransfers201ResponseDataRelationships) GetOriginStockLocationOk() (*POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation, bool)`

GetOriginStockLocationOk returns a tuple with the OriginStockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStockLocation

`func (o *POSTStockTransfers201ResponseDataRelationships) SetOriginStockLocation(v POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation)`

SetOriginStockLocation sets OriginStockLocation field to given value.

### HasOriginStockLocation

`func (o *POSTStockTransfers201ResponseDataRelationships) HasOriginStockLocation() bool`

HasOriginStockLocation returns a boolean if a field has been set.

### GetDestinationStockLocation

`func (o *POSTStockTransfers201ResponseDataRelationships) GetDestinationStockLocation() POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation`

GetDestinationStockLocation returns the DestinationStockLocation field if non-nil, zero value otherwise.

### GetDestinationStockLocationOk

`func (o *POSTStockTransfers201ResponseDataRelationships) GetDestinationStockLocationOk() (*POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation, bool)`

GetDestinationStockLocationOk returns a tuple with the DestinationStockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationStockLocation

`func (o *POSTStockTransfers201ResponseDataRelationships) SetDestinationStockLocation(v POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation)`

SetDestinationStockLocation sets DestinationStockLocation field to given value.

### HasDestinationStockLocation

`func (o *POSTStockTransfers201ResponseDataRelationships) HasDestinationStockLocation() bool`

HasDestinationStockLocation returns a boolean if a field has been set.

### GetShipment

`func (o *POSTStockTransfers201ResponseDataRelationships) GetShipment() POSTLineItems201ResponseDataRelationshipsShipment`

GetShipment returns the Shipment field if non-nil, zero value otherwise.

### GetShipmentOk

`func (o *POSTStockTransfers201ResponseDataRelationships) GetShipmentOk() (*POSTLineItems201ResponseDataRelationshipsShipment, bool)`

GetShipmentOk returns a tuple with the Shipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipment

`func (o *POSTStockTransfers201ResponseDataRelationships) SetShipment(v POSTLineItems201ResponseDataRelationshipsShipment)`

SetShipment sets Shipment field to given value.

### HasShipment

`func (o *POSTStockTransfers201ResponseDataRelationships) HasShipment() bool`

HasShipment returns a boolean if a field has been set.

### GetLineItem

`func (o *POSTStockTransfers201ResponseDataRelationships) GetLineItem() POSTLineItemOptions201ResponseDataRelationshipsLineItem`

GetLineItem returns the LineItem field if non-nil, zero value otherwise.

### GetLineItemOk

`func (o *POSTStockTransfers201ResponseDataRelationships) GetLineItemOk() (*POSTLineItemOptions201ResponseDataRelationshipsLineItem, bool)`

GetLineItemOk returns a tuple with the LineItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItem

`func (o *POSTStockTransfers201ResponseDataRelationships) SetLineItem(v POSTLineItemOptions201ResponseDataRelationshipsLineItem)`

SetLineItem sets LineItem field to given value.

### HasLineItem

`func (o *POSTStockTransfers201ResponseDataRelationships) HasLineItem() bool`

HasLineItem returns a boolean if a field has been set.

### GetStockReservation

`func (o *POSTStockTransfers201ResponseDataRelationships) GetStockReservation() POSTStockLineItems201ResponseDataRelationshipsStockReservation`

GetStockReservation returns the StockReservation field if non-nil, zero value otherwise.

### GetStockReservationOk

`func (o *POSTStockTransfers201ResponseDataRelationships) GetStockReservationOk() (*POSTStockLineItems201ResponseDataRelationshipsStockReservation, bool)`

GetStockReservationOk returns a tuple with the StockReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockReservation

`func (o *POSTStockTransfers201ResponseDataRelationships) SetStockReservation(v POSTStockLineItems201ResponseDataRelationshipsStockReservation)`

SetStockReservation sets StockReservation field to given value.

### HasStockReservation

`func (o *POSTStockTransfers201ResponseDataRelationships) HasStockReservation() bool`

HasStockReservation returns a boolean if a field has been set.

### GetAttachments

`func (o *POSTStockTransfers201ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *POSTStockTransfers201ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *POSTStockTransfers201ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *POSTStockTransfers201ResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *POSTStockTransfers201ResponseDataRelationships) GetEvents() POSTAddresses201ResponseDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *POSTStockTransfers201ResponseDataRelationships) GetEventsOk() (*POSTAddresses201ResponseDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *POSTStockTransfers201ResponseDataRelationships) SetEvents(v POSTAddresses201ResponseDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *POSTStockTransfers201ResponseDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetVersions

`func (o *POSTStockTransfers201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *POSTStockTransfers201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *POSTStockTransfers201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *POSTStockTransfers201ResponseDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


