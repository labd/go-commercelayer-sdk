# GETKlarnaPayments200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | Pointer to **interface{}** | The identifier of the payment session, useful to updated it. | [optional] 
**ClientToken** | Pointer to **interface{}** | The public token linked to your API credential. Available upon session creation. | [optional] 
**PaymentMethods** | Pointer to **interface{}** | The merchant available payment methods for the assoiated order. Available upon session creation. | [optional] 
**AuthToken** | Pointer to **interface{}** | The token returned by a successful client authorization, mandatory to place the order. | [optional] 
**MismatchedAmounts** | Pointer to **interface{}** | Indicates if the order current amount differs form the one of the created payment intent. | [optional] 
**IntentAmountCents** | Pointer to **interface{}** | The amount of the associated payment intent, in cents. | [optional] 
**IntentAmountFloat** | Pointer to **interface{}** | The amount of the associated payment intent, float. | [optional] 
**FormattedIntentAmount** | Pointer to **interface{}** | The amount of the associated payment intent, formatted. | [optional] 
**PaymentInstrument** | Pointer to **interface{}** | Information about the payment instrument used in the transaction | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETKlarnaPayments200ResponseDataInnerAttributes

`func NewGETKlarnaPayments200ResponseDataInnerAttributes() *GETKlarnaPayments200ResponseDataInnerAttributes`

NewGETKlarnaPayments200ResponseDataInnerAttributes instantiates a new GETKlarnaPayments200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETKlarnaPayments200ResponseDataInnerAttributesWithDefaults

`func NewGETKlarnaPayments200ResponseDataInnerAttributesWithDefaults() *GETKlarnaPayments200ResponseDataInnerAttributes`

NewGETKlarnaPayments200ResponseDataInnerAttributesWithDefaults instantiates a new GETKlarnaPayments200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetSessionId() interface{}`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetSessionIdOk() (*interface{}, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetSessionId(v interface{})`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionIdNil

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
### GetClientToken

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetClientToken() interface{}`

GetClientToken returns the ClientToken field if non-nil, zero value otherwise.

### GetClientTokenOk

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetClientTokenOk() (*interface{}, bool)`

GetClientTokenOk returns a tuple with the ClientToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientToken

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetClientToken(v interface{})`

SetClientToken sets ClientToken field to given value.

### HasClientToken

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) HasClientToken() bool`

HasClientToken returns a boolean if a field has been set.

### SetClientTokenNil

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetClientTokenNil(b bool)`

 SetClientTokenNil sets the value for ClientToken to be an explicit nil

### UnsetClientToken
`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) UnsetClientToken()`

UnsetClientToken ensures that no value is present for ClientToken, not even an explicit nil
### GetPaymentMethods

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetPaymentMethods() interface{}`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetPaymentMethodsOk() (*interface{}, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetPaymentMethods(v interface{})`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### SetPaymentMethodsNil

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetPaymentMethodsNil(b bool)`

 SetPaymentMethodsNil sets the value for PaymentMethods to be an explicit nil

### UnsetPaymentMethods
`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) UnsetPaymentMethods()`

UnsetPaymentMethods ensures that no value is present for PaymentMethods, not even an explicit nil
### GetAuthToken

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetAuthToken() interface{}`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetAuthTokenOk() (*interface{}, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetAuthToken(v interface{})`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### SetAuthTokenNil

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetAuthTokenNil(b bool)`

 SetAuthTokenNil sets the value for AuthToken to be an explicit nil

### UnsetAuthToken
`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) UnsetAuthToken()`

UnsetAuthToken ensures that no value is present for AuthToken, not even an explicit nil
### GetMismatchedAmounts

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetMismatchedAmounts() interface{}`

GetMismatchedAmounts returns the MismatchedAmounts field if non-nil, zero value otherwise.

### GetMismatchedAmountsOk

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetMismatchedAmountsOk() (*interface{}, bool)`

GetMismatchedAmountsOk returns a tuple with the MismatchedAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMismatchedAmounts

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetMismatchedAmounts(v interface{})`

SetMismatchedAmounts sets MismatchedAmounts field to given value.

### HasMismatchedAmounts

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) HasMismatchedAmounts() bool`

HasMismatchedAmounts returns a boolean if a field has been set.

### SetMismatchedAmountsNil

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetMismatchedAmountsNil(b bool)`

 SetMismatchedAmountsNil sets the value for MismatchedAmounts to be an explicit nil

### UnsetMismatchedAmounts
`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) UnsetMismatchedAmounts()`

UnsetMismatchedAmounts ensures that no value is present for MismatchedAmounts, not even an explicit nil
### GetIntentAmountCents

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetIntentAmountCents() interface{}`

GetIntentAmountCents returns the IntentAmountCents field if non-nil, zero value otherwise.

### GetIntentAmountCentsOk

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetIntentAmountCentsOk() (*interface{}, bool)`

GetIntentAmountCentsOk returns a tuple with the IntentAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentAmountCents

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetIntentAmountCents(v interface{})`

SetIntentAmountCents sets IntentAmountCents field to given value.

### HasIntentAmountCents

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) HasIntentAmountCents() bool`

HasIntentAmountCents returns a boolean if a field has been set.

### SetIntentAmountCentsNil

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetIntentAmountCentsNil(b bool)`

 SetIntentAmountCentsNil sets the value for IntentAmountCents to be an explicit nil

### UnsetIntentAmountCents
`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) UnsetIntentAmountCents()`

UnsetIntentAmountCents ensures that no value is present for IntentAmountCents, not even an explicit nil
### GetIntentAmountFloat

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetIntentAmountFloat() interface{}`

GetIntentAmountFloat returns the IntentAmountFloat field if non-nil, zero value otherwise.

### GetIntentAmountFloatOk

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetIntentAmountFloatOk() (*interface{}, bool)`

GetIntentAmountFloatOk returns a tuple with the IntentAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentAmountFloat

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetIntentAmountFloat(v interface{})`

SetIntentAmountFloat sets IntentAmountFloat field to given value.

### HasIntentAmountFloat

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) HasIntentAmountFloat() bool`

HasIntentAmountFloat returns a boolean if a field has been set.

### SetIntentAmountFloatNil

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetIntentAmountFloatNil(b bool)`

 SetIntentAmountFloatNil sets the value for IntentAmountFloat to be an explicit nil

### UnsetIntentAmountFloat
`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) UnsetIntentAmountFloat()`

UnsetIntentAmountFloat ensures that no value is present for IntentAmountFloat, not even an explicit nil
### GetFormattedIntentAmount

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetFormattedIntentAmount() interface{}`

GetFormattedIntentAmount returns the FormattedIntentAmount field if non-nil, zero value otherwise.

### GetFormattedIntentAmountOk

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetFormattedIntentAmountOk() (*interface{}, bool)`

GetFormattedIntentAmountOk returns a tuple with the FormattedIntentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedIntentAmount

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetFormattedIntentAmount(v interface{})`

SetFormattedIntentAmount sets FormattedIntentAmount field to given value.

### HasFormattedIntentAmount

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) HasFormattedIntentAmount() bool`

HasFormattedIntentAmount returns a boolean if a field has been set.

### SetFormattedIntentAmountNil

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetFormattedIntentAmountNil(b bool)`

 SetFormattedIntentAmountNil sets the value for FormattedIntentAmount to be an explicit nil

### UnsetFormattedIntentAmount
`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) UnsetFormattedIntentAmount()`

UnsetFormattedIntentAmount ensures that no value is present for FormattedIntentAmount, not even an explicit nil
### GetPaymentInstrument

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetPaymentInstrument() interface{}`

GetPaymentInstrument returns the PaymentInstrument field if non-nil, zero value otherwise.

### GetPaymentInstrumentOk

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetPaymentInstrumentOk() (*interface{}, bool)`

GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrument

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetPaymentInstrument(v interface{})`

SetPaymentInstrument sets PaymentInstrument field to given value.

### HasPaymentInstrument

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) HasPaymentInstrument() bool`

HasPaymentInstrument returns a boolean if a field has been set.

### SetPaymentInstrumentNil

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetPaymentInstrumentNil(b bool)`

 SetPaymentInstrumentNil sets the value for PaymentInstrument to be an explicit nil

### UnsetPaymentInstrument
`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) UnsetPaymentInstrument()`

UnsetPaymentInstrument ensures that no value is present for PaymentInstrument, not even an explicit nil
### GetCreatedAt

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


