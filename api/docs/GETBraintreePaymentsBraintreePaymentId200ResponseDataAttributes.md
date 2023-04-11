# GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientToken** | Pointer to **interface{}** | The Braintree payment client token. Required by the Braintree JS SDK. | [optional] 
**PaymentMethodNonce** | Pointer to **interface{}** | The Braintree payment method nonce. Sent by the Braintree JS SDK. | [optional] 
**PaymentId** | Pointer to **interface{}** | The Braintree payment ID used by local payment and sent by the Braintree JS SDK. | [optional] 
**Local** | Pointer to **interface{}** | Indicates if the payment is local, in such case Braintree will trigger a webhook call passing the \&quot;payment_id\&quot; and \&quot;payment_method_nonce\&quot; in order to complete the transaction. | [optional] 
**Options** | Pointer to **interface{}** | Braintree payment options: &#39;customer_id&#39; and &#39;payment_method_token&#39; | [optional] 
**PaymentInstrument** | Pointer to **interface{}** | Information about the payment instrument used in the transaction | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes

`func NewGETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes() *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes`

NewGETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes instantiates a new GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETBraintreePaymentsBraintreePaymentId200ResponseDataAttributesWithDefaults

`func NewGETBraintreePaymentsBraintreePaymentId200ResponseDataAttributesWithDefaults() *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes`

NewGETBraintreePaymentsBraintreePaymentId200ResponseDataAttributesWithDefaults instantiates a new GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientToken

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetClientToken() interface{}`

GetClientToken returns the ClientToken field if non-nil, zero value otherwise.

### GetClientTokenOk

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetClientTokenOk() (*interface{}, bool)`

GetClientTokenOk returns a tuple with the ClientToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientToken

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetClientToken(v interface{})`

SetClientToken sets ClientToken field to given value.

### HasClientToken

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasClientToken() bool`

HasClientToken returns a boolean if a field has been set.

### SetClientTokenNil

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetClientTokenNil(b bool)`

 SetClientTokenNil sets the value for ClientToken to be an explicit nil

### UnsetClientToken
`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) UnsetClientToken()`

UnsetClientToken ensures that no value is present for ClientToken, not even an explicit nil
### GetPaymentMethodNonce

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetPaymentMethodNonce() interface{}`

GetPaymentMethodNonce returns the PaymentMethodNonce field if non-nil, zero value otherwise.

### GetPaymentMethodNonceOk

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetPaymentMethodNonceOk() (*interface{}, bool)`

GetPaymentMethodNonceOk returns a tuple with the PaymentMethodNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodNonce

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetPaymentMethodNonce(v interface{})`

SetPaymentMethodNonce sets PaymentMethodNonce field to given value.

### HasPaymentMethodNonce

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasPaymentMethodNonce() bool`

HasPaymentMethodNonce returns a boolean if a field has been set.

### SetPaymentMethodNonceNil

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetPaymentMethodNonceNil(b bool)`

 SetPaymentMethodNonceNil sets the value for PaymentMethodNonce to be an explicit nil

### UnsetPaymentMethodNonce
`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) UnsetPaymentMethodNonce()`

UnsetPaymentMethodNonce ensures that no value is present for PaymentMethodNonce, not even an explicit nil
### GetPaymentId

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetPaymentId() interface{}`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetPaymentIdOk() (*interface{}, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetPaymentId(v interface{})`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### SetPaymentIdNil

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetPaymentIdNil(b bool)`

 SetPaymentIdNil sets the value for PaymentId to be an explicit nil

### UnsetPaymentId
`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) UnsetPaymentId()`

UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil
### GetLocal

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetLocal() interface{}`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetLocalOk() (*interface{}, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetLocal(v interface{})`

SetLocal sets Local field to given value.

### HasLocal

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### SetLocalNil

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetLocalNil(b bool)`

 SetLocalNil sets the value for Local to be an explicit nil

### UnsetLocal
`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) UnsetLocal()`

UnsetLocal ensures that no value is present for Local, not even an explicit nil
### GetOptions

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetOptions() interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetOptionsOk() (*interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetOptions(v interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetPaymentInstrument

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetPaymentInstrument() interface{}`

GetPaymentInstrument returns the PaymentInstrument field if non-nil, zero value otherwise.

### GetPaymentInstrumentOk

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetPaymentInstrumentOk() (*interface{}, bool)`

GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrument

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetPaymentInstrument(v interface{})`

SetPaymentInstrument sets PaymentInstrument field to given value.

### HasPaymentInstrument

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasPaymentInstrument() bool`

HasPaymentInstrument returns a boolean if a field has been set.

### SetPaymentInstrumentNil

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetPaymentInstrumentNil(b bool)`

 SetPaymentInstrumentNil sets the value for PaymentInstrument to be an explicit nil

### UnsetPaymentInstrument
`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) UnsetPaymentInstrument()`

UnsetPaymentInstrument ensures that no value is present for PaymentInstrument, not even an explicit nil
### GetCreatedAt

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


