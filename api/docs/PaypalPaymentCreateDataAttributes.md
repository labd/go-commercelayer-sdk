# PaypalPaymentCreateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnUrl** | **string** | The URL where the payer is redirected after they approve the payment. | 
**CancelUrl** | **string** | The URL where the payer is redirected after they cancel the payment. | 
**NoteToPayer** | Pointer to **string** | A free-form field that you can use to send a note to the payer on PayPal. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPaypalPaymentCreateDataAttributes

`func NewPaypalPaymentCreateDataAttributes(returnUrl string, cancelUrl string, ) *PaypalPaymentCreateDataAttributes`

NewPaypalPaymentCreateDataAttributes instantiates a new PaypalPaymentCreateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaypalPaymentCreateDataAttributesWithDefaults

`func NewPaypalPaymentCreateDataAttributesWithDefaults() *PaypalPaymentCreateDataAttributes`

NewPaypalPaymentCreateDataAttributesWithDefaults instantiates a new PaypalPaymentCreateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnUrl

`func (o *PaypalPaymentCreateDataAttributes) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *PaypalPaymentCreateDataAttributes) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *PaypalPaymentCreateDataAttributes) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.


### GetCancelUrl

`func (o *PaypalPaymentCreateDataAttributes) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *PaypalPaymentCreateDataAttributes) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *PaypalPaymentCreateDataAttributes) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.


### GetNoteToPayer

`func (o *PaypalPaymentCreateDataAttributes) GetNoteToPayer() string`

GetNoteToPayer returns the NoteToPayer field if non-nil, zero value otherwise.

### GetNoteToPayerOk

`func (o *PaypalPaymentCreateDataAttributes) GetNoteToPayerOk() (*string, bool)`

GetNoteToPayerOk returns a tuple with the NoteToPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteToPayer

`func (o *PaypalPaymentCreateDataAttributes) SetNoteToPayer(v string)`

SetNoteToPayer sets NoteToPayer field to given value.

### HasNoteToPayer

`func (o *PaypalPaymentCreateDataAttributes) HasNoteToPayer() bool`

HasNoteToPayer returns a boolean if a field has been set.

### GetReference

`func (o *PaypalPaymentCreateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaypalPaymentCreateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaypalPaymentCreateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PaypalPaymentCreateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PaypalPaymentCreateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PaypalPaymentCreateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PaypalPaymentCreateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PaypalPaymentCreateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PaypalPaymentCreateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PaypalPaymentCreateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PaypalPaymentCreateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PaypalPaymentCreateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


