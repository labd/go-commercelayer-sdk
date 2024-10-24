# POSTParcels201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shipment** | Pointer to [**POSTLineItems201ResponseDataRelationshipsShipment**](POSTLineItems201ResponseDataRelationshipsShipment.md) |  | [optional] 
**Package** | Pointer to [**POSTParcels201ResponseDataRelationshipsPackage**](POSTParcels201ResponseDataRelationshipsPackage.md) |  | [optional] 
**ParcelLineItems** | Pointer to [**POSTParcels201ResponseDataRelationshipsParcelLineItems**](POSTParcels201ResponseDataRelationshipsParcelLineItems.md) |  | [optional] 
**Attachments** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments.md) |  | [optional] 
**Events** | Pointer to [**POSTAddresses201ResponseDataRelationshipsEvents**](POSTAddresses201ResponseDataRelationshipsEvents.md) |  | [optional] 
**Versions** | Pointer to [**POSTAddresses201ResponseDataRelationshipsVersions**](POSTAddresses201ResponseDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewPOSTParcels201ResponseDataRelationships

`func NewPOSTParcels201ResponseDataRelationships() *POSTParcels201ResponseDataRelationships`

NewPOSTParcels201ResponseDataRelationships instantiates a new POSTParcels201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTParcels201ResponseDataRelationshipsWithDefaults

`func NewPOSTParcels201ResponseDataRelationshipsWithDefaults() *POSTParcels201ResponseDataRelationships`

NewPOSTParcels201ResponseDataRelationshipsWithDefaults instantiates a new POSTParcels201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipment

`func (o *POSTParcels201ResponseDataRelationships) GetShipment() POSTLineItems201ResponseDataRelationshipsShipment`

GetShipment returns the Shipment field if non-nil, zero value otherwise.

### GetShipmentOk

`func (o *POSTParcels201ResponseDataRelationships) GetShipmentOk() (*POSTLineItems201ResponseDataRelationshipsShipment, bool)`

GetShipmentOk returns a tuple with the Shipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipment

`func (o *POSTParcels201ResponseDataRelationships) SetShipment(v POSTLineItems201ResponseDataRelationshipsShipment)`

SetShipment sets Shipment field to given value.

### HasShipment

`func (o *POSTParcels201ResponseDataRelationships) HasShipment() bool`

HasShipment returns a boolean if a field has been set.

### GetPackage

`func (o *POSTParcels201ResponseDataRelationships) GetPackage() POSTParcels201ResponseDataRelationshipsPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *POSTParcels201ResponseDataRelationships) GetPackageOk() (*POSTParcels201ResponseDataRelationshipsPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *POSTParcels201ResponseDataRelationships) SetPackage(v POSTParcels201ResponseDataRelationshipsPackage)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *POSTParcels201ResponseDataRelationships) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetParcelLineItems

`func (o *POSTParcels201ResponseDataRelationships) GetParcelLineItems() POSTParcels201ResponseDataRelationshipsParcelLineItems`

GetParcelLineItems returns the ParcelLineItems field if non-nil, zero value otherwise.

### GetParcelLineItemsOk

`func (o *POSTParcels201ResponseDataRelationships) GetParcelLineItemsOk() (*POSTParcels201ResponseDataRelationshipsParcelLineItems, bool)`

GetParcelLineItemsOk returns a tuple with the ParcelLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelLineItems

`func (o *POSTParcels201ResponseDataRelationships) SetParcelLineItems(v POSTParcels201ResponseDataRelationshipsParcelLineItems)`

SetParcelLineItems sets ParcelLineItems field to given value.

### HasParcelLineItems

`func (o *POSTParcels201ResponseDataRelationships) HasParcelLineItems() bool`

HasParcelLineItems returns a boolean if a field has been set.

### GetAttachments

`func (o *POSTParcels201ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *POSTParcels201ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *POSTParcels201ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *POSTParcels201ResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *POSTParcels201ResponseDataRelationships) GetEvents() POSTAddresses201ResponseDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *POSTParcels201ResponseDataRelationships) GetEventsOk() (*POSTAddresses201ResponseDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *POSTParcels201ResponseDataRelationships) SetEvents(v POSTAddresses201ResponseDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *POSTParcels201ResponseDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetVersions

`func (o *POSTParcels201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *POSTParcels201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *POSTParcels201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *POSTParcels201ResponseDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


