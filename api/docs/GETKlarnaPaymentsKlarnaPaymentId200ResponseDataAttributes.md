# GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | Pointer to **interface{}** | The identifier of the payment session. | [optional] 
**ClientToken** | Pointer to **interface{}** | The public token linked to your API credential. Available upon session creation. | [optional] 
**PaymentMethods** | Pointer to **interface{}** | The merchant available payment methods for the assoiated order. Available upon session creation. | [optional] 
**AuthToken** | Pointer to **interface{}** | The token returned by a successful client authorization, mandatory to place the order. | [optional] 
**MismatchedAmounts** | Pointer to **interface{}** | Indicates if the order current amount differs form the one of the created payment intent. | [optional] 
**PaymentInstrument** | Pointer to **interface{}** | Information about the payment instrument used in the transaction. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes

`func NewGETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes() *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes`

NewGETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes instantiates a new GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributesWithDefaults

`func NewGETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributesWithDefaults() *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes`

NewGETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributesWithDefaults instantiates a new GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetSessionId() interface{}`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetSessionIdOk() (*interface{}, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetSessionId(v interface{})`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionIdNil

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
### GetClientToken

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetClientToken() interface{}`

GetClientToken returns the ClientToken field if non-nil, zero value otherwise.

### GetClientTokenOk

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetClientTokenOk() (*interface{}, bool)`

GetClientTokenOk returns a tuple with the ClientToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientToken

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetClientToken(v interface{})`

SetClientToken sets ClientToken field to given value.

### HasClientToken

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) HasClientToken() bool`

HasClientToken returns a boolean if a field has been set.

### SetClientTokenNil

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetClientTokenNil(b bool)`

 SetClientTokenNil sets the value for ClientToken to be an explicit nil

### UnsetClientToken
`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) UnsetClientToken()`

UnsetClientToken ensures that no value is present for ClientToken, not even an explicit nil
### GetPaymentMethods

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetPaymentMethods() interface{}`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetPaymentMethodsOk() (*interface{}, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetPaymentMethods(v interface{})`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### SetPaymentMethodsNil

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetPaymentMethodsNil(b bool)`

 SetPaymentMethodsNil sets the value for PaymentMethods to be an explicit nil

### UnsetPaymentMethods
`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) UnsetPaymentMethods()`

UnsetPaymentMethods ensures that no value is present for PaymentMethods, not even an explicit nil
### GetAuthToken

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetAuthToken() interface{}`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetAuthTokenOk() (*interface{}, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetAuthToken(v interface{})`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### SetAuthTokenNil

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetAuthTokenNil(b bool)`

 SetAuthTokenNil sets the value for AuthToken to be an explicit nil

### UnsetAuthToken
`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) UnsetAuthToken()`

UnsetAuthToken ensures that no value is present for AuthToken, not even an explicit nil
### GetMismatchedAmounts

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetMismatchedAmounts() interface{}`

GetMismatchedAmounts returns the MismatchedAmounts field if non-nil, zero value otherwise.

### GetMismatchedAmountsOk

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetMismatchedAmountsOk() (*interface{}, bool)`

GetMismatchedAmountsOk returns a tuple with the MismatchedAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMismatchedAmounts

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetMismatchedAmounts(v interface{})`

SetMismatchedAmounts sets MismatchedAmounts field to given value.

### HasMismatchedAmounts

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) HasMismatchedAmounts() bool`

HasMismatchedAmounts returns a boolean if a field has been set.

### SetMismatchedAmountsNil

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetMismatchedAmountsNil(b bool)`

 SetMismatchedAmountsNil sets the value for MismatchedAmounts to be an explicit nil

### UnsetMismatchedAmounts
`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) UnsetMismatchedAmounts()`

UnsetMismatchedAmounts ensures that no value is present for MismatchedAmounts, not even an explicit nil
### GetPaymentInstrument

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetPaymentInstrument() interface{}`

GetPaymentInstrument returns the PaymentInstrument field if non-nil, zero value otherwise.

### GetPaymentInstrumentOk

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetPaymentInstrumentOk() (*interface{}, bool)`

GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrument

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetPaymentInstrument(v interface{})`

SetPaymentInstrument sets PaymentInstrument field to given value.

### HasPaymentInstrument

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) HasPaymentInstrument() bool`

HasPaymentInstrument returns a boolean if a field has been set.

### SetPaymentInstrumentNil

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetPaymentInstrumentNil(b bool)`

 SetPaymentInstrumentNil sets the value for PaymentInstrument to be an explicit nil

### UnsetPaymentInstrument
`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) UnsetPaymentInstrument()`

UnsetPaymentInstrument ensures that no value is present for PaymentInstrument, not even an explicit nil
### GetCreatedAt

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


