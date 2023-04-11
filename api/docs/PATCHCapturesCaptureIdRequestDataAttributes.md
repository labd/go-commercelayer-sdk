# PATCHCapturesCaptureIdRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**Refund** | Pointer to **interface{}** | Send this attribute if you want to create a refund for this capture. | [optional] 
**RefundAmountCents** | Pointer to **interface{}** | The associated refund amount, in cents. | [optional] 

## Methods

### NewPATCHCapturesCaptureIdRequestDataAttributes

`func NewPATCHCapturesCaptureIdRequestDataAttributes() *PATCHCapturesCaptureIdRequestDataAttributes`

NewPATCHCapturesCaptureIdRequestDataAttributes instantiates a new PATCHCapturesCaptureIdRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHCapturesCaptureIdRequestDataAttributesWithDefaults

`func NewPATCHCapturesCaptureIdRequestDataAttributesWithDefaults() *PATCHCapturesCaptureIdRequestDataAttributes`

NewPATCHCapturesCaptureIdRequestDataAttributesWithDefaults instantiates a new PATCHCapturesCaptureIdRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReference

`func (o *PATCHCapturesCaptureIdRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHCapturesCaptureIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHCapturesCaptureIdRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHCapturesCaptureIdRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHCapturesCaptureIdRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHCapturesCaptureIdRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHCapturesCaptureIdRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHCapturesCaptureIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHCapturesCaptureIdRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHCapturesCaptureIdRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHCapturesCaptureIdRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHCapturesCaptureIdRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHCapturesCaptureIdRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHCapturesCaptureIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHCapturesCaptureIdRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHCapturesCaptureIdRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHCapturesCaptureIdRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHCapturesCaptureIdRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetRefund

`func (o *PATCHCapturesCaptureIdRequestDataAttributes) GetRefund() interface{}`

GetRefund returns the Refund field if non-nil, zero value otherwise.

### GetRefundOk

`func (o *PATCHCapturesCaptureIdRequestDataAttributes) GetRefundOk() (*interface{}, bool)`

GetRefundOk returns a tuple with the Refund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefund

`func (o *PATCHCapturesCaptureIdRequestDataAttributes) SetRefund(v interface{})`

SetRefund sets Refund field to given value.

### HasRefund

`func (o *PATCHCapturesCaptureIdRequestDataAttributes) HasRefund() bool`

HasRefund returns a boolean if a field has been set.

### SetRefundNil

`func (o *PATCHCapturesCaptureIdRequestDataAttributes) SetRefundNil(b bool)`

 SetRefundNil sets the value for Refund to be an explicit nil

### UnsetRefund
`func (o *PATCHCapturesCaptureIdRequestDataAttributes) UnsetRefund()`

UnsetRefund ensures that no value is present for Refund, not even an explicit nil
### GetRefundAmountCents

`func (o *PATCHCapturesCaptureIdRequestDataAttributes) GetRefundAmountCents() interface{}`

GetRefundAmountCents returns the RefundAmountCents field if non-nil, zero value otherwise.

### GetRefundAmountCentsOk

`func (o *PATCHCapturesCaptureIdRequestDataAttributes) GetRefundAmountCentsOk() (*interface{}, bool)`

GetRefundAmountCentsOk returns a tuple with the RefundAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmountCents

`func (o *PATCHCapturesCaptureIdRequestDataAttributes) SetRefundAmountCents(v interface{})`

SetRefundAmountCents sets RefundAmountCents field to given value.

### HasRefundAmountCents

`func (o *PATCHCapturesCaptureIdRequestDataAttributes) HasRefundAmountCents() bool`

HasRefundAmountCents returns a boolean if a field has been set.

### SetRefundAmountCentsNil

`func (o *PATCHCapturesCaptureIdRequestDataAttributes) SetRefundAmountCentsNil(b bool)`

 SetRefundAmountCentsNil sets the value for RefundAmountCents to be an explicit nil

### UnsetRefundAmountCents
`func (o *PATCHCapturesCaptureIdRequestDataAttributes) UnsetRefundAmountCents()`

UnsetRefundAmountCents ensures that no value is present for RefundAmountCents, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


