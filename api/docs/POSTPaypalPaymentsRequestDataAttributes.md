# POSTPaypalPaymentsRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnUrl** | **interface{}** | The URL where the payer is redirected after they approve the payment. | 
**CancelUrl** | **interface{}** | The URL where the payer is redirected after they cancel the payment. | 
**NoteToPayer** | Pointer to **interface{}** | A free-form field that you can use to send a note to the payer on PayPal. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTPaypalPaymentsRequestDataAttributes

`func NewPOSTPaypalPaymentsRequestDataAttributes(returnUrl interface{}, cancelUrl interface{}, ) *POSTPaypalPaymentsRequestDataAttributes`

NewPOSTPaypalPaymentsRequestDataAttributes instantiates a new POSTPaypalPaymentsRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTPaypalPaymentsRequestDataAttributesWithDefaults

`func NewPOSTPaypalPaymentsRequestDataAttributesWithDefaults() *POSTPaypalPaymentsRequestDataAttributes`

NewPOSTPaypalPaymentsRequestDataAttributesWithDefaults instantiates a new POSTPaypalPaymentsRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnUrl

`func (o *POSTPaypalPaymentsRequestDataAttributes) GetReturnUrl() interface{}`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *POSTPaypalPaymentsRequestDataAttributes) GetReturnUrlOk() (*interface{}, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *POSTPaypalPaymentsRequestDataAttributes) SetReturnUrl(v interface{})`

SetReturnUrl sets ReturnUrl field to given value.


### SetReturnUrlNil

`func (o *POSTPaypalPaymentsRequestDataAttributes) SetReturnUrlNil(b bool)`

 SetReturnUrlNil sets the value for ReturnUrl to be an explicit nil

### UnsetReturnUrl
`func (o *POSTPaypalPaymentsRequestDataAttributes) UnsetReturnUrl()`

UnsetReturnUrl ensures that no value is present for ReturnUrl, not even an explicit nil
### GetCancelUrl

`func (o *POSTPaypalPaymentsRequestDataAttributes) GetCancelUrl() interface{}`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *POSTPaypalPaymentsRequestDataAttributes) GetCancelUrlOk() (*interface{}, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *POSTPaypalPaymentsRequestDataAttributes) SetCancelUrl(v interface{})`

SetCancelUrl sets CancelUrl field to given value.


### SetCancelUrlNil

`func (o *POSTPaypalPaymentsRequestDataAttributes) SetCancelUrlNil(b bool)`

 SetCancelUrlNil sets the value for CancelUrl to be an explicit nil

### UnsetCancelUrl
`func (o *POSTPaypalPaymentsRequestDataAttributes) UnsetCancelUrl()`

UnsetCancelUrl ensures that no value is present for CancelUrl, not even an explicit nil
### GetNoteToPayer

`func (o *POSTPaypalPaymentsRequestDataAttributes) GetNoteToPayer() interface{}`

GetNoteToPayer returns the NoteToPayer field if non-nil, zero value otherwise.

### GetNoteToPayerOk

`func (o *POSTPaypalPaymentsRequestDataAttributes) GetNoteToPayerOk() (*interface{}, bool)`

GetNoteToPayerOk returns a tuple with the NoteToPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteToPayer

`func (o *POSTPaypalPaymentsRequestDataAttributes) SetNoteToPayer(v interface{})`

SetNoteToPayer sets NoteToPayer field to given value.

### HasNoteToPayer

`func (o *POSTPaypalPaymentsRequestDataAttributes) HasNoteToPayer() bool`

HasNoteToPayer returns a boolean if a field has been set.

### SetNoteToPayerNil

`func (o *POSTPaypalPaymentsRequestDataAttributes) SetNoteToPayerNil(b bool)`

 SetNoteToPayerNil sets the value for NoteToPayer to be an explicit nil

### UnsetNoteToPayer
`func (o *POSTPaypalPaymentsRequestDataAttributes) UnsetNoteToPayer()`

UnsetNoteToPayer ensures that no value is present for NoteToPayer, not even an explicit nil
### GetReference

`func (o *POSTPaypalPaymentsRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTPaypalPaymentsRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTPaypalPaymentsRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTPaypalPaymentsRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTPaypalPaymentsRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTPaypalPaymentsRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTPaypalPaymentsRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTPaypalPaymentsRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTPaypalPaymentsRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTPaypalPaymentsRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTPaypalPaymentsRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTPaypalPaymentsRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTPaypalPaymentsRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTPaypalPaymentsRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTPaypalPaymentsRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTPaypalPaymentsRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTPaypalPaymentsRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTPaypalPaymentsRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


