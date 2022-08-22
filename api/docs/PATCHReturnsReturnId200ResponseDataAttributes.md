# PATCHReturnsReturnId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | Pointer to **bool** | Send this attribute if you want to activate this return. | [optional] 
**Approve** | Pointer to **bool** | Send this attribute if you want to mark this return as approved. | [optional] 
**Cancel** | Pointer to **bool** | Send this attribute if you want to mark this return as cancelled. | [optional] 
**Ship** | Pointer to **bool** | Send this attribute if you want to mark this return as shipped. | [optional] 
**Reject** | Pointer to **bool** | Send this attribute if you want to mark this return as rejected. | [optional] 
**Receive** | Pointer to **bool** | Send this attribute if you want to mark this return as received. | [optional] 
**Restock** | Pointer to **bool** | Send this attribute if you want to restock all of the return line items. | [optional] 
**Archive** | Pointer to **bool** | Send this attribute if you want to archive the return. | [optional] 
**Unarchive** | Pointer to **bool** | Send this attribute if you want to unarchive the return. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHReturnsReturnId200ResponseDataAttributes

`func NewPATCHReturnsReturnId200ResponseDataAttributes() *PATCHReturnsReturnId200ResponseDataAttributes`

NewPATCHReturnsReturnId200ResponseDataAttributes instantiates a new PATCHReturnsReturnId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHReturnsReturnId200ResponseDataAttributesWithDefaults

`func NewPATCHReturnsReturnId200ResponseDataAttributesWithDefaults() *PATCHReturnsReturnId200ResponseDataAttributes`

NewPATCHReturnsReturnId200ResponseDataAttributesWithDefaults instantiates a new PATCHReturnsReturnId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetRequest() bool`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetRequestOk() (*bool, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetRequest(v bool)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetApprove

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetApprove() bool`

GetApprove returns the Approve field if non-nil, zero value otherwise.

### GetApproveOk

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetApproveOk() (*bool, bool)`

GetApproveOk returns a tuple with the Approve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprove

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetApprove(v bool)`

SetApprove sets Approve field to given value.

### HasApprove

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) HasApprove() bool`

HasApprove returns a boolean if a field has been set.

### GetCancel

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetCancel() bool`

GetCancel returns the Cancel field if non-nil, zero value otherwise.

### GetCancelOk

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetCancelOk() (*bool, bool)`

GetCancelOk returns a tuple with the Cancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancel

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetCancel(v bool)`

SetCancel sets Cancel field to given value.

### HasCancel

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) HasCancel() bool`

HasCancel returns a boolean if a field has been set.

### GetShip

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetShip() bool`

GetShip returns the Ship field if non-nil, zero value otherwise.

### GetShipOk

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetShipOk() (*bool, bool)`

GetShipOk returns a tuple with the Ship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShip

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetShip(v bool)`

SetShip sets Ship field to given value.

### HasShip

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) HasShip() bool`

HasShip returns a boolean if a field has been set.

### GetReject

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetReject() bool`

GetReject returns the Reject field if non-nil, zero value otherwise.

### GetRejectOk

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetRejectOk() (*bool, bool)`

GetRejectOk returns a tuple with the Reject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReject

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetReject(v bool)`

SetReject sets Reject field to given value.

### HasReject

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) HasReject() bool`

HasReject returns a boolean if a field has been set.

### GetReceive

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetReceive() bool`

GetReceive returns the Receive field if non-nil, zero value otherwise.

### GetReceiveOk

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetReceiveOk() (*bool, bool)`

GetReceiveOk returns a tuple with the Receive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceive

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetReceive(v bool)`

SetReceive sets Receive field to given value.

### HasReceive

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) HasReceive() bool`

HasReceive returns a boolean if a field has been set.

### GetRestock

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetRestock() bool`

GetRestock returns the Restock field if non-nil, zero value otherwise.

### GetRestockOk

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetRestockOk() (*bool, bool)`

GetRestockOk returns a tuple with the Restock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestock

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetRestock(v bool)`

SetRestock sets Restock field to given value.

### HasRestock

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) HasRestock() bool`

HasRestock returns a boolean if a field has been set.

### GetArchive

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetArchive() bool`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetArchiveOk() (*bool, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetArchive(v bool)`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) HasArchive() bool`

HasArchive returns a boolean if a field has been set.

### GetUnarchive

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetUnarchive() bool`

GetUnarchive returns the Unarchive field if non-nil, zero value otherwise.

### GetUnarchiveOk

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetUnarchiveOk() (*bool, bool)`

GetUnarchiveOk returns a tuple with the Unarchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnarchive

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetUnarchive(v bool)`

SetUnarchive sets Unarchive field to given value.

### HasUnarchive

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) HasUnarchive() bool`

HasUnarchive returns a boolean if a field has been set.

### GetReference

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


