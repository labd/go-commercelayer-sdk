# ParcelResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shipment** | Pointer to [**ParcelResponseDataRelationshipsShipment**](ParcelResponseDataRelationshipsShipment.md) |  | [optional] 
**Package** | Pointer to [**ParcelResponseDataRelationshipsPackage**](ParcelResponseDataRelationshipsPackage.md) |  | [optional] 
**ParcelLineItems** | Pointer to [**ParcelResponseDataRelationshipsParcelLineItems**](ParcelResponseDataRelationshipsParcelLineItems.md) |  | [optional] 
**Attachments** | Pointer to [**AvalaraAccountResponseDataRelationshipsAttachments**](AvalaraAccountResponseDataRelationshipsAttachments.md) |  | [optional] 
**Events** | Pointer to [**CustomerAddressResponseDataRelationshipsEvents**](CustomerAddressResponseDataRelationshipsEvents.md) |  | [optional] 

## Methods

### NewParcelResponseDataRelationships

`func NewParcelResponseDataRelationships() *ParcelResponseDataRelationships`

NewParcelResponseDataRelationships instantiates a new ParcelResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParcelResponseDataRelationshipsWithDefaults

`func NewParcelResponseDataRelationshipsWithDefaults() *ParcelResponseDataRelationships`

NewParcelResponseDataRelationshipsWithDefaults instantiates a new ParcelResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipment

`func (o *ParcelResponseDataRelationships) GetShipment() ParcelResponseDataRelationshipsShipment`

GetShipment returns the Shipment field if non-nil, zero value otherwise.

### GetShipmentOk

`func (o *ParcelResponseDataRelationships) GetShipmentOk() (*ParcelResponseDataRelationshipsShipment, bool)`

GetShipmentOk returns a tuple with the Shipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipment

`func (o *ParcelResponseDataRelationships) SetShipment(v ParcelResponseDataRelationshipsShipment)`

SetShipment sets Shipment field to given value.

### HasShipment

`func (o *ParcelResponseDataRelationships) HasShipment() bool`

HasShipment returns a boolean if a field has been set.

### GetPackage

`func (o *ParcelResponseDataRelationships) GetPackage() ParcelResponseDataRelationshipsPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *ParcelResponseDataRelationships) GetPackageOk() (*ParcelResponseDataRelationshipsPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *ParcelResponseDataRelationships) SetPackage(v ParcelResponseDataRelationshipsPackage)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *ParcelResponseDataRelationships) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetParcelLineItems

`func (o *ParcelResponseDataRelationships) GetParcelLineItems() ParcelResponseDataRelationshipsParcelLineItems`

GetParcelLineItems returns the ParcelLineItems field if non-nil, zero value otherwise.

### GetParcelLineItemsOk

`func (o *ParcelResponseDataRelationships) GetParcelLineItemsOk() (*ParcelResponseDataRelationshipsParcelLineItems, bool)`

GetParcelLineItemsOk returns a tuple with the ParcelLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelLineItems

`func (o *ParcelResponseDataRelationships) SetParcelLineItems(v ParcelResponseDataRelationshipsParcelLineItems)`

SetParcelLineItems sets ParcelLineItems field to given value.

### HasParcelLineItems

`func (o *ParcelResponseDataRelationships) HasParcelLineItems() bool`

HasParcelLineItems returns a boolean if a field has been set.

### GetAttachments

`func (o *ParcelResponseDataRelationships) GetAttachments() AvalaraAccountResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ParcelResponseDataRelationships) GetAttachmentsOk() (*AvalaraAccountResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ParcelResponseDataRelationships) SetAttachments(v AvalaraAccountResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ParcelResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *ParcelResponseDataRelationships) GetEvents() CustomerAddressResponseDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ParcelResponseDataRelationships) GetEventsOk() (*CustomerAddressResponseDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ParcelResponseDataRelationships) SetEvents(v CustomerAddressResponseDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ParcelResponseDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


