# ParcelDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shipment** | Pointer to [**LineItemDataRelationshipsShipment**](LineItemDataRelationshipsShipment.md) |  | [optional] 
**Package** | Pointer to [**ParcelDataRelationshipsPackage**](ParcelDataRelationshipsPackage.md) |  | [optional] 
**ParcelLineItems** | Pointer to [**ParcelDataRelationshipsParcelLineItems**](ParcelDataRelationshipsParcelLineItems.md) |  | [optional] 
**Attachments** | Pointer to [**AuthorizationDataRelationshipsAttachments**](AuthorizationDataRelationshipsAttachments.md) |  | [optional] 
**Events** | Pointer to [**AddressDataRelationshipsEvents**](AddressDataRelationshipsEvents.md) |  | [optional] 
**Versions** | Pointer to [**AddressDataRelationshipsVersions**](AddressDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewParcelDataRelationships

`func NewParcelDataRelationships() *ParcelDataRelationships`

NewParcelDataRelationships instantiates a new ParcelDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParcelDataRelationshipsWithDefaults

`func NewParcelDataRelationshipsWithDefaults() *ParcelDataRelationships`

NewParcelDataRelationshipsWithDefaults instantiates a new ParcelDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipment

`func (o *ParcelDataRelationships) GetShipment() LineItemDataRelationshipsShipment`

GetShipment returns the Shipment field if non-nil, zero value otherwise.

### GetShipmentOk

`func (o *ParcelDataRelationships) GetShipmentOk() (*LineItemDataRelationshipsShipment, bool)`

GetShipmentOk returns a tuple with the Shipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipment

`func (o *ParcelDataRelationships) SetShipment(v LineItemDataRelationshipsShipment)`

SetShipment sets Shipment field to given value.

### HasShipment

`func (o *ParcelDataRelationships) HasShipment() bool`

HasShipment returns a boolean if a field has been set.

### GetPackage

`func (o *ParcelDataRelationships) GetPackage() ParcelDataRelationshipsPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *ParcelDataRelationships) GetPackageOk() (*ParcelDataRelationshipsPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *ParcelDataRelationships) SetPackage(v ParcelDataRelationshipsPackage)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *ParcelDataRelationships) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetParcelLineItems

`func (o *ParcelDataRelationships) GetParcelLineItems() ParcelDataRelationshipsParcelLineItems`

GetParcelLineItems returns the ParcelLineItems field if non-nil, zero value otherwise.

### GetParcelLineItemsOk

`func (o *ParcelDataRelationships) GetParcelLineItemsOk() (*ParcelDataRelationshipsParcelLineItems, bool)`

GetParcelLineItemsOk returns a tuple with the ParcelLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelLineItems

`func (o *ParcelDataRelationships) SetParcelLineItems(v ParcelDataRelationshipsParcelLineItems)`

SetParcelLineItems sets ParcelLineItems field to given value.

### HasParcelLineItems

`func (o *ParcelDataRelationships) HasParcelLineItems() bool`

HasParcelLineItems returns a boolean if a field has been set.

### GetAttachments

`func (o *ParcelDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ParcelDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ParcelDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ParcelDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *ParcelDataRelationships) GetEvents() AddressDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ParcelDataRelationships) GetEventsOk() (*AddressDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ParcelDataRelationships) SetEvents(v AddressDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ParcelDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetVersions

`func (o *ParcelDataRelationships) GetVersions() AddressDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ParcelDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ParcelDataRelationships) SetVersions(v AddressDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *ParcelDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


