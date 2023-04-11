# POSTCheckoutComPayments201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentType** | **interface{}** | The payment source type. | 
**Token** | **interface{}** | The Checkout.com card or digital wallet token. | 
**SessionId** | Pointer to **interface{}** | A payment session ID used to obtain the details. | [optional] 
**SuccessUrl** | Pointer to **interface{}** | The URL to redirect your customer upon 3DS succeeded authentication. | [optional] 
**FailureUrl** | Pointer to **interface{}** | The URL to redirect your customer upon 3DS failed authentication. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTCheckoutComPayments201ResponseDataAttributes

`func NewPOSTCheckoutComPayments201ResponseDataAttributes(paymentType interface{}, token interface{}, ) *POSTCheckoutComPayments201ResponseDataAttributes`

NewPOSTCheckoutComPayments201ResponseDataAttributes instantiates a new POSTCheckoutComPayments201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTCheckoutComPayments201ResponseDataAttributesWithDefaults

`func NewPOSTCheckoutComPayments201ResponseDataAttributesWithDefaults() *POSTCheckoutComPayments201ResponseDataAttributes`

NewPOSTCheckoutComPayments201ResponseDataAttributesWithDefaults instantiates a new POSTCheckoutComPayments201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentType

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetPaymentType() interface{}`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetPaymentTypeOk() (*interface{}, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) SetPaymentType(v interface{})`

SetPaymentType sets PaymentType field to given value.


### SetPaymentTypeNil

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) SetPaymentTypeNil(b bool)`

 SetPaymentTypeNil sets the value for PaymentType to be an explicit nil

### UnsetPaymentType
`func (o *POSTCheckoutComPayments201ResponseDataAttributes) UnsetPaymentType()`

UnsetPaymentType ensures that no value is present for PaymentType, not even an explicit nil
### GetToken

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetToken() interface{}`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetTokenOk() (*interface{}, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) SetToken(v interface{})`

SetToken sets Token field to given value.


### SetTokenNil

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *POSTCheckoutComPayments201ResponseDataAttributes) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetSessionId

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetSessionId() interface{}`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetSessionIdOk() (*interface{}, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) SetSessionId(v interface{})`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionIdNil

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *POSTCheckoutComPayments201ResponseDataAttributes) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
### GetSuccessUrl

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetSuccessUrl() interface{}`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetSuccessUrlOk() (*interface{}, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) SetSuccessUrl(v interface{})`

SetSuccessUrl sets SuccessUrl field to given value.

### HasSuccessUrl

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) HasSuccessUrl() bool`

HasSuccessUrl returns a boolean if a field has been set.

### SetSuccessUrlNil

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) SetSuccessUrlNil(b bool)`

 SetSuccessUrlNil sets the value for SuccessUrl to be an explicit nil

### UnsetSuccessUrl
`func (o *POSTCheckoutComPayments201ResponseDataAttributes) UnsetSuccessUrl()`

UnsetSuccessUrl ensures that no value is present for SuccessUrl, not even an explicit nil
### GetFailureUrl

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetFailureUrl() interface{}`

GetFailureUrl returns the FailureUrl field if non-nil, zero value otherwise.

### GetFailureUrlOk

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetFailureUrlOk() (*interface{}, bool)`

GetFailureUrlOk returns a tuple with the FailureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureUrl

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) SetFailureUrl(v interface{})`

SetFailureUrl sets FailureUrl field to given value.

### HasFailureUrl

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) HasFailureUrl() bool`

HasFailureUrl returns a boolean if a field has been set.

### SetFailureUrlNil

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) SetFailureUrlNil(b bool)`

 SetFailureUrlNil sets the value for FailureUrl to be an explicit nil

### UnsetFailureUrl
`func (o *POSTCheckoutComPayments201ResponseDataAttributes) UnsetFailureUrl()`

UnsetFailureUrl ensures that no value is present for FailureUrl, not even an explicit nil
### GetReference

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTCheckoutComPayments201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTCheckoutComPayments201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTCheckoutComPayments201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTCheckoutComPayments201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


