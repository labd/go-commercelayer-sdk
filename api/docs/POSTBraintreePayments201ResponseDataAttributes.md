# POSTBraintreePayments201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | Pointer to **string** | The Braintree payment ID used by local payment and sent by the Braintree JS SDK. | [optional] 
**Local** | Pointer to **bool** | Indicates if the payment is local, in such case Braintree will trigger a webhook call passing the \&quot;payment_id\&quot; and \&quot;payment_method_nonce\&quot; in order to complete the transaction. | [optional] 
**Options** | Pointer to **map[string]interface{}** | Braintree payment options: &#39;customer_id&#39; and &#39;payment_method_token&#39; | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTBraintreePayments201ResponseDataAttributes

`func NewPOSTBraintreePayments201ResponseDataAttributes() *POSTBraintreePayments201ResponseDataAttributes`

NewPOSTBraintreePayments201ResponseDataAttributes instantiates a new POSTBraintreePayments201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTBraintreePayments201ResponseDataAttributesWithDefaults

`func NewPOSTBraintreePayments201ResponseDataAttributesWithDefaults() *POSTBraintreePayments201ResponseDataAttributes`

NewPOSTBraintreePayments201ResponseDataAttributesWithDefaults instantiates a new POSTBraintreePayments201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *POSTBraintreePayments201ResponseDataAttributes) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *POSTBraintreePayments201ResponseDataAttributes) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *POSTBraintreePayments201ResponseDataAttributes) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *POSTBraintreePayments201ResponseDataAttributes) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetLocal

`func (o *POSTBraintreePayments201ResponseDataAttributes) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *POSTBraintreePayments201ResponseDataAttributes) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *POSTBraintreePayments201ResponseDataAttributes) SetLocal(v bool)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *POSTBraintreePayments201ResponseDataAttributes) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetOptions

`func (o *POSTBraintreePayments201ResponseDataAttributes) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *POSTBraintreePayments201ResponseDataAttributes) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *POSTBraintreePayments201ResponseDataAttributes) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *POSTBraintreePayments201ResponseDataAttributes) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetReference

`func (o *POSTBraintreePayments201ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTBraintreePayments201ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTBraintreePayments201ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTBraintreePayments201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *POSTBraintreePayments201ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTBraintreePayments201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTBraintreePayments201ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTBraintreePayments201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *POSTBraintreePayments201ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTBraintreePayments201ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTBraintreePayments201ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTBraintreePayments201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


