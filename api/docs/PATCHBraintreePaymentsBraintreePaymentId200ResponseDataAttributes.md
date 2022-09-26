# PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethodNonce** | Pointer to **string** | The Braintree payment method nonce. Sent by the Braintree JS SDK. | [optional] 
**PaymentId** | Pointer to **string** | The Braintree payment ID used by local payment and sent by the Braintree JS SDK. | [optional] 
**Local** | Pointer to **bool** | Indicates if the payment is local, in such case Braintree will trigger a webhook call passing the \&quot;payment_id\&quot; and \&quot;payment_method_nonce\&quot; in order to complete the transaction. | [optional] 
**Options** | Pointer to **map[string]interface{}** | Braintree payment options: &#39;customer_id&#39; and &#39;payment_method_token&#39; | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

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

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetPaymentMethodNonce() string`

GetPaymentMethodNonce returns the PaymentMethodNonce field if non-nil, zero value otherwise.

### GetPaymentMethodNonceOk

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetPaymentMethodNonceOk() (*string, bool)`

GetPaymentMethodNonceOk returns a tuple with the PaymentMethodNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodNonce

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetPaymentMethodNonce(v string)`

SetPaymentMethodNonce sets PaymentMethodNonce field to given value.

### HasPaymentMethodNonce

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasPaymentMethodNonce() bool`

HasPaymentMethodNonce returns a boolean if a field has been set.

### GetPaymentId

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetLocal

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetLocal(v bool)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetOptions

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetReference

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


