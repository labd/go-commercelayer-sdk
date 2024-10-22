# PATCHCapturesCaptureId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Succeeded** | Pointer to **interface{}** | Indicates if the transaction is successful. | [optional] 
**Forward** | Pointer to **interface{}** | Send this attribute if you want to forward a stuck transaction to succeeded and update associated order states accordingly. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**Refund** | Pointer to **interface{}** | Send this attribute if you want to create a refund for this capture. | [optional] 
**RefundAmountCents** | Pointer to **interface{}** | Send this attribute as a value in cents if you want to overwrite the amount to be refunded. | [optional] 

## Methods

### NewPATCHCapturesCaptureId200ResponseDataAttributes

`func NewPATCHCapturesCaptureId200ResponseDataAttributes() *PATCHCapturesCaptureId200ResponseDataAttributes`

NewPATCHCapturesCaptureId200ResponseDataAttributes instantiates a new PATCHCapturesCaptureId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHCapturesCaptureId200ResponseDataAttributesWithDefaults

`func NewPATCHCapturesCaptureId200ResponseDataAttributesWithDefaults() *PATCHCapturesCaptureId200ResponseDataAttributes`

NewPATCHCapturesCaptureId200ResponseDataAttributesWithDefaults instantiates a new PATCHCapturesCaptureId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSucceeded

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) GetSucceeded() interface{}`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) GetSucceededOk() (*interface{}, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) SetSucceeded(v interface{})`

SetSucceeded sets Succeeded field to given value.

### HasSucceeded

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) HasSucceeded() bool`

HasSucceeded returns a boolean if a field has been set.

### SetSucceededNil

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) SetSucceededNil(b bool)`

 SetSucceededNil sets the value for Succeeded to be an explicit nil

### UnsetSucceeded
`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) UnsetSucceeded()`

UnsetSucceeded ensures that no value is present for Succeeded, not even an explicit nil
### GetForward

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) GetForward() interface{}`

GetForward returns the Forward field if non-nil, zero value otherwise.

### GetForwardOk

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) GetForwardOk() (*interface{}, bool)`

GetForwardOk returns a tuple with the Forward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForward

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) SetForward(v interface{})`

SetForward sets Forward field to given value.

### HasForward

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) HasForward() bool`

HasForward returns a boolean if a field has been set.

### SetForwardNil

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) SetForwardNil(b bool)`

 SetForwardNil sets the value for Forward to be an explicit nil

### UnsetForward
`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) UnsetForward()`

UnsetForward ensures that no value is present for Forward, not even an explicit nil
### GetReference

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetRefund

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) GetRefund() interface{}`

GetRefund returns the Refund field if non-nil, zero value otherwise.

### GetRefundOk

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) GetRefundOk() (*interface{}, bool)`

GetRefundOk returns a tuple with the Refund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefund

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) SetRefund(v interface{})`

SetRefund sets Refund field to given value.

### HasRefund

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) HasRefund() bool`

HasRefund returns a boolean if a field has been set.

### SetRefundNil

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) SetRefundNil(b bool)`

 SetRefundNil sets the value for Refund to be an explicit nil

### UnsetRefund
`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) UnsetRefund()`

UnsetRefund ensures that no value is present for Refund, not even an explicit nil
### GetRefundAmountCents

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) GetRefundAmountCents() interface{}`

GetRefundAmountCents returns the RefundAmountCents field if non-nil, zero value otherwise.

### GetRefundAmountCentsOk

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) GetRefundAmountCentsOk() (*interface{}, bool)`

GetRefundAmountCentsOk returns a tuple with the RefundAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmountCents

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) SetRefundAmountCents(v interface{})`

SetRefundAmountCents sets RefundAmountCents field to given value.

### HasRefundAmountCents

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) HasRefundAmountCents() bool`

HasRefundAmountCents returns a boolean if a field has been set.

### SetRefundAmountCentsNil

`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) SetRefundAmountCentsNil(b bool)`

 SetRefundAmountCentsNil sets the value for RefundAmountCents to be an explicit nil

### UnsetRefundAmountCents
`func (o *PATCHCapturesCaptureId200ResponseDataAttributes) UnsetRefundAmountCents()`

UnsetRefundAmountCents ensures that no value is present for RefundAmountCents, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


