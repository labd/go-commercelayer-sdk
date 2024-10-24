# POSTStripePayments201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StripeId** | Pointer to **interface{}** | The Stripe payment intent ID. Required to identify a payment session on stripe. | [optional] 
**ClientSecret** | Pointer to **interface{}** | The Stripe payment intent client secret. Required to create a charge through Stripe.js. | [optional] 
**Options** | Pointer to **interface{}** | Stripe payment options: &#39;customer&#39;, &#39;payment_method&#39;, &#39;return_url&#39;, etc. Check Stripe payment intent API for more details. | [optional] 
**ReturnUrl** | Pointer to **interface{}** | The URL to return to when a redirect payment is completed. | [optional] 
**ReceiptEmail** | Pointer to **interface{}** | The email address to send the receipt to. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTStripePayments201ResponseDataAttributes

`func NewPOSTStripePayments201ResponseDataAttributes() *POSTStripePayments201ResponseDataAttributes`

NewPOSTStripePayments201ResponseDataAttributes instantiates a new POSTStripePayments201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTStripePayments201ResponseDataAttributesWithDefaults

`func NewPOSTStripePayments201ResponseDataAttributesWithDefaults() *POSTStripePayments201ResponseDataAttributes`

NewPOSTStripePayments201ResponseDataAttributesWithDefaults instantiates a new POSTStripePayments201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStripeId

`func (o *POSTStripePayments201ResponseDataAttributes) GetStripeId() interface{}`

GetStripeId returns the StripeId field if non-nil, zero value otherwise.

### GetStripeIdOk

`func (o *POSTStripePayments201ResponseDataAttributes) GetStripeIdOk() (*interface{}, bool)`

GetStripeIdOk returns a tuple with the StripeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeId

`func (o *POSTStripePayments201ResponseDataAttributes) SetStripeId(v interface{})`

SetStripeId sets StripeId field to given value.

### HasStripeId

`func (o *POSTStripePayments201ResponseDataAttributes) HasStripeId() bool`

HasStripeId returns a boolean if a field has been set.

### SetStripeIdNil

`func (o *POSTStripePayments201ResponseDataAttributes) SetStripeIdNil(b bool)`

 SetStripeIdNil sets the value for StripeId to be an explicit nil

### UnsetStripeId
`func (o *POSTStripePayments201ResponseDataAttributes) UnsetStripeId()`

UnsetStripeId ensures that no value is present for StripeId, not even an explicit nil
### GetClientSecret

`func (o *POSTStripePayments201ResponseDataAttributes) GetClientSecret() interface{}`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *POSTStripePayments201ResponseDataAttributes) GetClientSecretOk() (*interface{}, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *POSTStripePayments201ResponseDataAttributes) SetClientSecret(v interface{})`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *POSTStripePayments201ResponseDataAttributes) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### SetClientSecretNil

`func (o *POSTStripePayments201ResponseDataAttributes) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *POSTStripePayments201ResponseDataAttributes) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
### GetOptions

`func (o *POSTStripePayments201ResponseDataAttributes) GetOptions() interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *POSTStripePayments201ResponseDataAttributes) GetOptionsOk() (*interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *POSTStripePayments201ResponseDataAttributes) SetOptions(v interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *POSTStripePayments201ResponseDataAttributes) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *POSTStripePayments201ResponseDataAttributes) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *POSTStripePayments201ResponseDataAttributes) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetReturnUrl

`func (o *POSTStripePayments201ResponseDataAttributes) GetReturnUrl() interface{}`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *POSTStripePayments201ResponseDataAttributes) GetReturnUrlOk() (*interface{}, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *POSTStripePayments201ResponseDataAttributes) SetReturnUrl(v interface{})`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *POSTStripePayments201ResponseDataAttributes) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### SetReturnUrlNil

`func (o *POSTStripePayments201ResponseDataAttributes) SetReturnUrlNil(b bool)`

 SetReturnUrlNil sets the value for ReturnUrl to be an explicit nil

### UnsetReturnUrl
`func (o *POSTStripePayments201ResponseDataAttributes) UnsetReturnUrl()`

UnsetReturnUrl ensures that no value is present for ReturnUrl, not even an explicit nil
### GetReceiptEmail

`func (o *POSTStripePayments201ResponseDataAttributes) GetReceiptEmail() interface{}`

GetReceiptEmail returns the ReceiptEmail field if non-nil, zero value otherwise.

### GetReceiptEmailOk

`func (o *POSTStripePayments201ResponseDataAttributes) GetReceiptEmailOk() (*interface{}, bool)`

GetReceiptEmailOk returns a tuple with the ReceiptEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptEmail

`func (o *POSTStripePayments201ResponseDataAttributes) SetReceiptEmail(v interface{})`

SetReceiptEmail sets ReceiptEmail field to given value.

### HasReceiptEmail

`func (o *POSTStripePayments201ResponseDataAttributes) HasReceiptEmail() bool`

HasReceiptEmail returns a boolean if a field has been set.

### SetReceiptEmailNil

`func (o *POSTStripePayments201ResponseDataAttributes) SetReceiptEmailNil(b bool)`

 SetReceiptEmailNil sets the value for ReceiptEmail to be an explicit nil

### UnsetReceiptEmail
`func (o *POSTStripePayments201ResponseDataAttributes) UnsetReceiptEmail()`

UnsetReceiptEmail ensures that no value is present for ReceiptEmail, not even an explicit nil
### GetReference

`func (o *POSTStripePayments201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTStripePayments201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTStripePayments201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTStripePayments201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTStripePayments201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTStripePayments201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTStripePayments201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTStripePayments201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTStripePayments201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTStripePayments201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTStripePayments201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTStripePayments201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTStripePayments201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTStripePayments201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTStripePayments201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTStripePayments201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTStripePayments201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTStripePayments201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


