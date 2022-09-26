# PATCHAuthorizationsAuthorizationId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**Capture** | Pointer to **bool** | Send this attribute if you want to create a capture for this authorization. | [optional] 
**CaptureAmountCents** | Pointer to **int32** | The associated capture amount, in cents. | [optional] 
**Void** | Pointer to **bool** | Send this attribute if you want to create a void for this authorization. | [optional] 

## Methods

### NewPATCHAuthorizationsAuthorizationId200ResponseDataAttributes

`func NewPATCHAuthorizationsAuthorizationId200ResponseDataAttributes() *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes`

NewPATCHAuthorizationsAuthorizationId200ResponseDataAttributes instantiates a new PATCHAuthorizationsAuthorizationId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHAuthorizationsAuthorizationId200ResponseDataAttributesWithDefaults

`func NewPATCHAuthorizationsAuthorizationId200ResponseDataAttributesWithDefaults() *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes`

NewPATCHAuthorizationsAuthorizationId200ResponseDataAttributesWithDefaults instantiates a new PATCHAuthorizationsAuthorizationId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReference

`func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCapture

`func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetCapture() bool`

GetCapture returns the Capture field if non-nil, zero value otherwise.

### GetCaptureOk

`func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetCaptureOk() (*bool, bool)`

GetCaptureOk returns a tuple with the Capture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapture

`func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) SetCapture(v bool)`

SetCapture sets Capture field to given value.

### HasCapture

`func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) HasCapture() bool`

HasCapture returns a boolean if a field has been set.

### GetCaptureAmountCents

`func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetCaptureAmountCents() int32`

GetCaptureAmountCents returns the CaptureAmountCents field if non-nil, zero value otherwise.

### GetCaptureAmountCentsOk

`func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetCaptureAmountCentsOk() (*int32, bool)`

GetCaptureAmountCentsOk returns a tuple with the CaptureAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureAmountCents

`func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) SetCaptureAmountCents(v int32)`

SetCaptureAmountCents sets CaptureAmountCents field to given value.

### HasCaptureAmountCents

`func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) HasCaptureAmountCents() bool`

HasCaptureAmountCents returns a boolean if a field has been set.

### GetVoid

`func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetVoid() bool`

GetVoid returns the Void field if non-nil, zero value otherwise.

### GetVoidOk

`func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetVoidOk() (*bool, bool)`

GetVoidOk returns a tuple with the Void field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoid

`func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) SetVoid(v bool)`

SetVoid sets Void field to given value.

### HasVoid

`func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) HasVoid() bool`

HasVoid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


