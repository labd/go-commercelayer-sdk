# PATCHAuthorizationsAuthorizationIdRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**Capture** | Pointer to **interface{}** | Send this attribute if you want to create a capture for this authorization. | [optional] 
**CaptureAmountCents** | Pointer to **interface{}** | The associated capture amount, in cents. | [optional] 
**Void** | Pointer to **interface{}** | Send this attribute if you want to create a void for this authorization. | [optional] 

## Methods

### NewPATCHAuthorizationsAuthorizationIdRequestDataAttributes

`func NewPATCHAuthorizationsAuthorizationIdRequestDataAttributes() *PATCHAuthorizationsAuthorizationIdRequestDataAttributes`

NewPATCHAuthorizationsAuthorizationIdRequestDataAttributes instantiates a new PATCHAuthorizationsAuthorizationIdRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHAuthorizationsAuthorizationIdRequestDataAttributesWithDefaults

`func NewPATCHAuthorizationsAuthorizationIdRequestDataAttributesWithDefaults() *PATCHAuthorizationsAuthorizationIdRequestDataAttributes`

NewPATCHAuthorizationsAuthorizationIdRequestDataAttributesWithDefaults instantiates a new PATCHAuthorizationsAuthorizationIdRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReference

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCapture

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) GetCapture() interface{}`

GetCapture returns the Capture field if non-nil, zero value otherwise.

### GetCaptureOk

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) GetCaptureOk() (*interface{}, bool)`

GetCaptureOk returns a tuple with the Capture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapture

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) SetCapture(v interface{})`

SetCapture sets Capture field to given value.

### HasCapture

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) HasCapture() bool`

HasCapture returns a boolean if a field has been set.

### SetCaptureNil

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) SetCaptureNil(b bool)`

 SetCaptureNil sets the value for Capture to be an explicit nil

### UnsetCapture
`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) UnsetCapture()`

UnsetCapture ensures that no value is present for Capture, not even an explicit nil
### GetCaptureAmountCents

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) GetCaptureAmountCents() interface{}`

GetCaptureAmountCents returns the CaptureAmountCents field if non-nil, zero value otherwise.

### GetCaptureAmountCentsOk

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) GetCaptureAmountCentsOk() (*interface{}, bool)`

GetCaptureAmountCentsOk returns a tuple with the CaptureAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureAmountCents

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) SetCaptureAmountCents(v interface{})`

SetCaptureAmountCents sets CaptureAmountCents field to given value.

### HasCaptureAmountCents

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) HasCaptureAmountCents() bool`

HasCaptureAmountCents returns a boolean if a field has been set.

### SetCaptureAmountCentsNil

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) SetCaptureAmountCentsNil(b bool)`

 SetCaptureAmountCentsNil sets the value for CaptureAmountCents to be an explicit nil

### UnsetCaptureAmountCents
`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) UnsetCaptureAmountCents()`

UnsetCaptureAmountCents ensures that no value is present for CaptureAmountCents, not even an explicit nil
### GetVoid

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) GetVoid() interface{}`

GetVoid returns the Void field if non-nil, zero value otherwise.

### GetVoidOk

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) GetVoidOk() (*interface{}, bool)`

GetVoidOk returns a tuple with the Void field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoid

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) SetVoid(v interface{})`

SetVoid sets Void field to given value.

### HasVoid

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) HasVoid() bool`

HasVoid returns a boolean if a field has been set.

### SetVoidNil

`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) SetVoidNil(b bool)`

 SetVoidNil sets the value for Void to be an explicit nil

### UnsetVoid
`func (o *PATCHAuthorizationsAuthorizationIdRequestDataAttributes) UnsetVoid()`

UnsetVoid ensures that no value is present for Void, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


