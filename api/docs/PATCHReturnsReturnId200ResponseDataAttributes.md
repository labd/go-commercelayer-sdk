# PATCHReturnsReturnId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | Pointer to **interface{}** | Send this attribute if you want to activate this return. | [optional] 
**Approve** | Pointer to **interface{}** | Send this attribute if you want to mark this return as approved. | [optional] 
**Cancel** | Pointer to **interface{}** | Send this attribute if you want to mark this return as cancelled. | [optional] 
**Ship** | Pointer to **interface{}** | Send this attribute if you want to mark this return as shipped. | [optional] 
**Reject** | Pointer to **interface{}** | Send this attribute if you want to mark this return as rejected. | [optional] 
**Receive** | Pointer to **interface{}** | Send this attribute if you want to mark this return as received. | [optional] 
**Restock** | Pointer to **interface{}** | Send this attribute if you want to restock all of the return line items. | [optional] 
**Archive** | Pointer to **interface{}** | Send this attribute if you want to archive the return. | [optional] 
**Unarchive** | Pointer to **interface{}** | Send this attribute if you want to unarchive the return. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

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

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetRequest() interface{}`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetRequestOk() (*interface{}, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetRequest(v interface{})`

SetRequest sets Request field to given value.

### HasRequest

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### SetRequestNil

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetRequestNil(b bool)`

 SetRequestNil sets the value for Request to be an explicit nil

### UnsetRequest
`func (o *PATCHReturnsReturnId200ResponseDataAttributes) UnsetRequest()`

UnsetRequest ensures that no value is present for Request, not even an explicit nil
### GetApprove

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetApprove() interface{}`

GetApprove returns the Approve field if non-nil, zero value otherwise.

### GetApproveOk

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetApproveOk() (*interface{}, bool)`

GetApproveOk returns a tuple with the Approve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprove

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetApprove(v interface{})`

SetApprove sets Approve field to given value.

### HasApprove

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) HasApprove() bool`

HasApprove returns a boolean if a field has been set.

### SetApproveNil

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetApproveNil(b bool)`

 SetApproveNil sets the value for Approve to be an explicit nil

### UnsetApprove
`func (o *PATCHReturnsReturnId200ResponseDataAttributes) UnsetApprove()`

UnsetApprove ensures that no value is present for Approve, not even an explicit nil
### GetCancel

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetCancel() interface{}`

GetCancel returns the Cancel field if non-nil, zero value otherwise.

### GetCancelOk

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetCancelOk() (*interface{}, bool)`

GetCancelOk returns a tuple with the Cancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancel

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetCancel(v interface{})`

SetCancel sets Cancel field to given value.

### HasCancel

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) HasCancel() bool`

HasCancel returns a boolean if a field has been set.

### SetCancelNil

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetCancelNil(b bool)`

 SetCancelNil sets the value for Cancel to be an explicit nil

### UnsetCancel
`func (o *PATCHReturnsReturnId200ResponseDataAttributes) UnsetCancel()`

UnsetCancel ensures that no value is present for Cancel, not even an explicit nil
### GetShip

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetShip() interface{}`

GetShip returns the Ship field if non-nil, zero value otherwise.

### GetShipOk

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetShipOk() (*interface{}, bool)`

GetShipOk returns a tuple with the Ship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShip

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetShip(v interface{})`

SetShip sets Ship field to given value.

### HasShip

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) HasShip() bool`

HasShip returns a boolean if a field has been set.

### SetShipNil

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetShipNil(b bool)`

 SetShipNil sets the value for Ship to be an explicit nil

### UnsetShip
`func (o *PATCHReturnsReturnId200ResponseDataAttributes) UnsetShip()`

UnsetShip ensures that no value is present for Ship, not even an explicit nil
### GetReject

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetReject() interface{}`

GetReject returns the Reject field if non-nil, zero value otherwise.

### GetRejectOk

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetRejectOk() (*interface{}, bool)`

GetRejectOk returns a tuple with the Reject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReject

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetReject(v interface{})`

SetReject sets Reject field to given value.

### HasReject

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) HasReject() bool`

HasReject returns a boolean if a field has been set.

### SetRejectNil

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetRejectNil(b bool)`

 SetRejectNil sets the value for Reject to be an explicit nil

### UnsetReject
`func (o *PATCHReturnsReturnId200ResponseDataAttributes) UnsetReject()`

UnsetReject ensures that no value is present for Reject, not even an explicit nil
### GetReceive

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetReceive() interface{}`

GetReceive returns the Receive field if non-nil, zero value otherwise.

### GetReceiveOk

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetReceiveOk() (*interface{}, bool)`

GetReceiveOk returns a tuple with the Receive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceive

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetReceive(v interface{})`

SetReceive sets Receive field to given value.

### HasReceive

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) HasReceive() bool`

HasReceive returns a boolean if a field has been set.

### SetReceiveNil

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetReceiveNil(b bool)`

 SetReceiveNil sets the value for Receive to be an explicit nil

### UnsetReceive
`func (o *PATCHReturnsReturnId200ResponseDataAttributes) UnsetReceive()`

UnsetReceive ensures that no value is present for Receive, not even an explicit nil
### GetRestock

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetRestock() interface{}`

GetRestock returns the Restock field if non-nil, zero value otherwise.

### GetRestockOk

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetRestockOk() (*interface{}, bool)`

GetRestockOk returns a tuple with the Restock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestock

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetRestock(v interface{})`

SetRestock sets Restock field to given value.

### HasRestock

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) HasRestock() bool`

HasRestock returns a boolean if a field has been set.

### SetRestockNil

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetRestockNil(b bool)`

 SetRestockNil sets the value for Restock to be an explicit nil

### UnsetRestock
`func (o *PATCHReturnsReturnId200ResponseDataAttributes) UnsetRestock()`

UnsetRestock ensures that no value is present for Restock, not even an explicit nil
### GetArchive

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetArchive() interface{}`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetArchiveOk() (*interface{}, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetArchive(v interface{})`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) HasArchive() bool`

HasArchive returns a boolean if a field has been set.

### SetArchiveNil

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetArchiveNil(b bool)`

 SetArchiveNil sets the value for Archive to be an explicit nil

### UnsetArchive
`func (o *PATCHReturnsReturnId200ResponseDataAttributes) UnsetArchive()`

UnsetArchive ensures that no value is present for Archive, not even an explicit nil
### GetUnarchive

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetUnarchive() interface{}`

GetUnarchive returns the Unarchive field if non-nil, zero value otherwise.

### GetUnarchiveOk

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetUnarchiveOk() (*interface{}, bool)`

GetUnarchiveOk returns a tuple with the Unarchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnarchive

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetUnarchive(v interface{})`

SetUnarchive sets Unarchive field to given value.

### HasUnarchive

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) HasUnarchive() bool`

HasUnarchive returns a boolean if a field has been set.

### SetUnarchiveNil

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetUnarchiveNil(b bool)`

 SetUnarchiveNil sets the value for Unarchive to be an explicit nil

### UnsetUnarchive
`func (o *PATCHReturnsReturnId200ResponseDataAttributes) UnsetUnarchive()`

UnsetUnarchive ensures that no value is present for Unarchive, not even an explicit nil
### GetReference

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHReturnsReturnId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHReturnsReturnId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHReturnsReturnId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHReturnsReturnId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


