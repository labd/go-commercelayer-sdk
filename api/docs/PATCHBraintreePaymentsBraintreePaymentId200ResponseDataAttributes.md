# PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethodNonce** | Pointer to **interface{}** | The Braintree payment method nonce. Sent by the Braintree JS SDK. | [optional] 
**PaymentId** | Pointer to **interface{}** | The Braintree payment ID used by local payment and sent by the Braintree JS SDK. | [optional] 
**Local** | Pointer to **interface{}** | Indicates if the payment is local, in such case Braintree will trigger a webhook call passing the \&quot;payment_id\&quot; and \&quot;payment_method_nonce\&quot; in order to complete the transaction. | [optional] 
**Options** | Pointer to **interface{}** | Braintree payment options, &#39;customer_id&#39; and &#39;payment_method_token&#39;. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes

`func NewPATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes() *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes`

NewPATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes instantiates a new PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributesWithDefaults

`func NewPATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributesWithDefaults() *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes`

NewPATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributesWithDefaults instantiates a new PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethodNonce

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetPaymentMethodNonce() interface{}`

GetPaymentMethodNonce returns the PaymentMethodNonce field if non-nil, zero value otherwise.

### GetPaymentMethodNonceOk

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetPaymentMethodNonceOk() (*interface{}, bool)`

GetPaymentMethodNonceOk returns a tuple with the PaymentMethodNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodNonce

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetPaymentMethodNonce(v interface{})`

SetPaymentMethodNonce sets PaymentMethodNonce field to given value.

### HasPaymentMethodNonce

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasPaymentMethodNonce() bool`

HasPaymentMethodNonce returns a boolean if a field has been set.

### SetPaymentMethodNonceNil

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetPaymentMethodNonceNil(b bool)`

 SetPaymentMethodNonceNil sets the value for PaymentMethodNonce to be an explicit nil

### UnsetPaymentMethodNonce
`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) UnsetPaymentMethodNonce()`

UnsetPaymentMethodNonce ensures that no value is present for PaymentMethodNonce, not even an explicit nil
### GetPaymentId

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetPaymentId() interface{}`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetPaymentIdOk() (*interface{}, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetPaymentId(v interface{})`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### SetPaymentIdNil

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetPaymentIdNil(b bool)`

 SetPaymentIdNil sets the value for PaymentId to be an explicit nil

### UnsetPaymentId
`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) UnsetPaymentId()`

UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil
### GetLocal

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetLocal() interface{}`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetLocalOk() (*interface{}, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetLocal(v interface{})`

SetLocal sets Local field to given value.

### HasLocal

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### SetLocalNil

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetLocalNil(b bool)`

 SetLocalNil sets the value for Local to be an explicit nil

### UnsetLocal
`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) UnsetLocal()`

UnsetLocal ensures that no value is present for Local, not even an explicit nil
### GetOptions

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetOptions() interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetOptionsOk() (*interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetOptions(v interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetReference

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


