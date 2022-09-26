# BraintreePaymentDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientToken** | Pointer to **string** | The Braintree payment client token. Required by the Braintree JS SDK. | [optional] 
**PaymentMethodNonce** | Pointer to **string** | The Braintree payment method nonce. Sent by the Braintree JS SDK. | [optional] 
**PaymentId** | Pointer to **string** | The Braintree payment ID used by local payment and sent by the Braintree JS SDK. | [optional] 
**Local** | Pointer to **bool** | Indicates if the payment is local, in such case Braintree will trigger a webhook call passing the \&quot;payment_id\&quot; and \&quot;payment_method_nonce\&quot; in order to complete the transaction. | [optional] 
**Options** | Pointer to **map[string]interface{}** | Braintree payment options: &#39;customer_id&#39; and &#39;payment_method_token&#39; | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewBraintreePaymentDataAttributes

`func NewBraintreePaymentDataAttributes() *BraintreePaymentDataAttributes`

NewBraintreePaymentDataAttributes instantiates a new BraintreePaymentDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBraintreePaymentDataAttributesWithDefaults

`func NewBraintreePaymentDataAttributesWithDefaults() *BraintreePaymentDataAttributes`

NewBraintreePaymentDataAttributesWithDefaults instantiates a new BraintreePaymentDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientToken

`func (o *BraintreePaymentDataAttributes) GetClientToken() string`

GetClientToken returns the ClientToken field if non-nil, zero value otherwise.

### GetClientTokenOk

`func (o *BraintreePaymentDataAttributes) GetClientTokenOk() (*string, bool)`

GetClientTokenOk returns a tuple with the ClientToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientToken

`func (o *BraintreePaymentDataAttributes) SetClientToken(v string)`

SetClientToken sets ClientToken field to given value.

### HasClientToken

`func (o *BraintreePaymentDataAttributes) HasClientToken() bool`

HasClientToken returns a boolean if a field has been set.

### GetPaymentMethodNonce

`func (o *BraintreePaymentDataAttributes) GetPaymentMethodNonce() string`

GetPaymentMethodNonce returns the PaymentMethodNonce field if non-nil, zero value otherwise.

### GetPaymentMethodNonceOk

`func (o *BraintreePaymentDataAttributes) GetPaymentMethodNonceOk() (*string, bool)`

GetPaymentMethodNonceOk returns a tuple with the PaymentMethodNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodNonce

`func (o *BraintreePaymentDataAttributes) SetPaymentMethodNonce(v string)`

SetPaymentMethodNonce sets PaymentMethodNonce field to given value.

### HasPaymentMethodNonce

`func (o *BraintreePaymentDataAttributes) HasPaymentMethodNonce() bool`

HasPaymentMethodNonce returns a boolean if a field has been set.

### GetPaymentId

`func (o *BraintreePaymentDataAttributes) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *BraintreePaymentDataAttributes) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *BraintreePaymentDataAttributes) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *BraintreePaymentDataAttributes) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetLocal

`func (o *BraintreePaymentDataAttributes) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *BraintreePaymentDataAttributes) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *BraintreePaymentDataAttributes) SetLocal(v bool)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *BraintreePaymentDataAttributes) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetOptions

`func (o *BraintreePaymentDataAttributes) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *BraintreePaymentDataAttributes) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *BraintreePaymentDataAttributes) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *BraintreePaymentDataAttributes) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetId

`func (o *BraintreePaymentDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BraintreePaymentDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BraintreePaymentDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BraintreePaymentDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BraintreePaymentDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BraintreePaymentDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BraintreePaymentDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BraintreePaymentDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BraintreePaymentDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BraintreePaymentDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BraintreePaymentDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BraintreePaymentDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *BraintreePaymentDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *BraintreePaymentDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *BraintreePaymentDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *BraintreePaymentDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *BraintreePaymentDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *BraintreePaymentDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *BraintreePaymentDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *BraintreePaymentDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *BraintreePaymentDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BraintreePaymentDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BraintreePaymentDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BraintreePaymentDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


