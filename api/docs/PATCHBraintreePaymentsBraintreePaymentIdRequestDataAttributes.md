# PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethodNonce** | Pointer to **interface{}** | The Braintree payment method nonce. Sent by the Braintree JS SDK. | [optional] 
**PaymentId** | Pointer to **interface{}** | The Braintree payment ID used by local payment and sent by the Braintree JS SDK. | [optional] 
**Local** | Pointer to **interface{}** | Indicates if the payment is local, in such case Braintree will trigger a webhook call passing the \&quot;payment_id\&quot; and \&quot;payment_method_nonce\&quot; in order to complete the transaction. | [optional] 
**Options** | Pointer to **interface{}** | Braintree payment options: &#39;customer_id&#39; and &#39;payment_method_token&#39; | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes

`func NewPATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes() *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes`

NewPATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes instantiates a new PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributesWithDefaults

`func NewPATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributesWithDefaults() *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes`

NewPATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributesWithDefaults instantiates a new PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethodNonce

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) GetPaymentMethodNonce() interface{}`

GetPaymentMethodNonce returns the PaymentMethodNonce field if non-nil, zero value otherwise.

### GetPaymentMethodNonceOk

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) GetPaymentMethodNonceOk() (*interface{}, bool)`

GetPaymentMethodNonceOk returns a tuple with the PaymentMethodNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodNonce

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) SetPaymentMethodNonce(v interface{})`

SetPaymentMethodNonce sets PaymentMethodNonce field to given value.

### HasPaymentMethodNonce

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) HasPaymentMethodNonce() bool`

HasPaymentMethodNonce returns a boolean if a field has been set.

### SetPaymentMethodNonceNil

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) SetPaymentMethodNonceNil(b bool)`

 SetPaymentMethodNonceNil sets the value for PaymentMethodNonce to be an explicit nil

### UnsetPaymentMethodNonce
`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) UnsetPaymentMethodNonce()`

UnsetPaymentMethodNonce ensures that no value is present for PaymentMethodNonce, not even an explicit nil
### GetPaymentId

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) GetPaymentId() interface{}`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) GetPaymentIdOk() (*interface{}, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) SetPaymentId(v interface{})`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### SetPaymentIdNil

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) SetPaymentIdNil(b bool)`

 SetPaymentIdNil sets the value for PaymentId to be an explicit nil

### UnsetPaymentId
`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) UnsetPaymentId()`

UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil
### GetLocal

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) GetLocal() interface{}`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) GetLocalOk() (*interface{}, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) SetLocal(v interface{})`

SetLocal sets Local field to given value.

### HasLocal

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### SetLocalNil

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) SetLocalNil(b bool)`

 SetLocalNil sets the value for Local to be an explicit nil

### UnsetLocal
`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) UnsetLocal()`

UnsetLocal ensures that no value is present for Local, not even an explicit nil
### GetOptions

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) GetOptions() interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) GetOptionsOk() (*interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) SetOptions(v interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetReference

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHBraintreePaymentsBraintreePaymentIdRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


