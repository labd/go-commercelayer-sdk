# GETCheckoutComPayments200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicKey** | Pointer to **interface{}** | The Checkout.com publishable API key. | [optional] 
**PaymentType** | Pointer to **interface{}** | The payment source type. | [optional] 
**Token** | Pointer to **interface{}** | The Checkout.com card or digital wallet token. | [optional] 
**SessionId** | Pointer to **interface{}** | A payment session ID used to obtain the details. | [optional] 
**SuccessUrl** | Pointer to **interface{}** | The URL to redirect your customer upon 3DS succeeded authentication. | [optional] 
**FailureUrl** | Pointer to **interface{}** | The URL to redirect your customer upon 3DS failed authentication. | [optional] 
**SourceId** | Pointer to **interface{}** | The payment source identifier that can be used for subsequent payments. | [optional] 
**CustomerToken** | Pointer to **interface{}** | The customer&#39;s unique identifier. This can be passed as a source when making a payment. | [optional] 
**RedirectUri** | Pointer to **interface{}** | The URI that the customer should be redirected to in order to complete the payment. | [optional] 
**PaymentResponse** | Pointer to **interface{}** | The Checkout.com payment response, used to fetch internal data. | [optional] 
**MismatchedAmounts** | Pointer to **interface{}** | Indicates if the order current amount differs form the one of the associated authorization. | [optional] 
**PaymentInstrument** | Pointer to **interface{}** | Information about the payment instrument used in the transaction | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETCheckoutComPayments200ResponseDataInnerAttributes

`func NewGETCheckoutComPayments200ResponseDataInnerAttributes() *GETCheckoutComPayments200ResponseDataInnerAttributes`

NewGETCheckoutComPayments200ResponseDataInnerAttributes instantiates a new GETCheckoutComPayments200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETCheckoutComPayments200ResponseDataInnerAttributesWithDefaults

`func NewGETCheckoutComPayments200ResponseDataInnerAttributesWithDefaults() *GETCheckoutComPayments200ResponseDataInnerAttributes`

NewGETCheckoutComPayments200ResponseDataInnerAttributesWithDefaults instantiates a new GETCheckoutComPayments200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicKey

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetPublicKey() interface{}`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetPublicKeyOk() (*interface{}, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetPublicKey(v interface{})`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetPaymentType

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetPaymentType() interface{}`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetPaymentTypeOk() (*interface{}, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetPaymentType(v interface{})`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.

### SetPaymentTypeNil

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetPaymentTypeNil(b bool)`

 SetPaymentTypeNil sets the value for PaymentType to be an explicit nil

### UnsetPaymentType
`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) UnsetPaymentType()`

UnsetPaymentType ensures that no value is present for PaymentType, not even an explicit nil
### GetToken

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetToken() interface{}`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetTokenOk() (*interface{}, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetToken(v interface{})`

SetToken sets Token field to given value.

### HasToken

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetSessionId

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetSessionId() interface{}`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetSessionIdOk() (*interface{}, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetSessionId(v interface{})`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionIdNil

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
### GetSuccessUrl

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetSuccessUrl() interface{}`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetSuccessUrlOk() (*interface{}, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetSuccessUrl(v interface{})`

SetSuccessUrl sets SuccessUrl field to given value.

### HasSuccessUrl

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasSuccessUrl() bool`

HasSuccessUrl returns a boolean if a field has been set.

### SetSuccessUrlNil

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetSuccessUrlNil(b bool)`

 SetSuccessUrlNil sets the value for SuccessUrl to be an explicit nil

### UnsetSuccessUrl
`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) UnsetSuccessUrl()`

UnsetSuccessUrl ensures that no value is present for SuccessUrl, not even an explicit nil
### GetFailureUrl

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetFailureUrl() interface{}`

GetFailureUrl returns the FailureUrl field if non-nil, zero value otherwise.

### GetFailureUrlOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetFailureUrlOk() (*interface{}, bool)`

GetFailureUrlOk returns a tuple with the FailureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureUrl

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetFailureUrl(v interface{})`

SetFailureUrl sets FailureUrl field to given value.

### HasFailureUrl

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasFailureUrl() bool`

HasFailureUrl returns a boolean if a field has been set.

### SetFailureUrlNil

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetFailureUrlNil(b bool)`

 SetFailureUrlNil sets the value for FailureUrl to be an explicit nil

### UnsetFailureUrl
`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) UnsetFailureUrl()`

UnsetFailureUrl ensures that no value is present for FailureUrl, not even an explicit nil
### GetSourceId

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetSourceId() interface{}`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetSourceIdOk() (*interface{}, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetSourceId(v interface{})`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetCustomerToken

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetCustomerToken() interface{}`

GetCustomerToken returns the CustomerToken field if non-nil, zero value otherwise.

### GetCustomerTokenOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetCustomerTokenOk() (*interface{}, bool)`

GetCustomerTokenOk returns a tuple with the CustomerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerToken

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetCustomerToken(v interface{})`

SetCustomerToken sets CustomerToken field to given value.

### HasCustomerToken

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasCustomerToken() bool`

HasCustomerToken returns a boolean if a field has been set.

### SetCustomerTokenNil

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetCustomerTokenNil(b bool)`

 SetCustomerTokenNil sets the value for CustomerToken to be an explicit nil

### UnsetCustomerToken
`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) UnsetCustomerToken()`

UnsetCustomerToken ensures that no value is present for CustomerToken, not even an explicit nil
### GetRedirectUri

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetRedirectUri() interface{}`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetRedirectUriOk() (*interface{}, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetRedirectUri(v interface{})`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### SetRedirectUriNil

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetRedirectUriNil(b bool)`

 SetRedirectUriNil sets the value for RedirectUri to be an explicit nil

### UnsetRedirectUri
`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) UnsetRedirectUri()`

UnsetRedirectUri ensures that no value is present for RedirectUri, not even an explicit nil
### GetPaymentResponse

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetPaymentResponse() interface{}`

GetPaymentResponse returns the PaymentResponse field if non-nil, zero value otherwise.

### GetPaymentResponseOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetPaymentResponseOk() (*interface{}, bool)`

GetPaymentResponseOk returns a tuple with the PaymentResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentResponse

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetPaymentResponse(v interface{})`

SetPaymentResponse sets PaymentResponse field to given value.

### HasPaymentResponse

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasPaymentResponse() bool`

HasPaymentResponse returns a boolean if a field has been set.

### SetPaymentResponseNil

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetPaymentResponseNil(b bool)`

 SetPaymentResponseNil sets the value for PaymentResponse to be an explicit nil

### UnsetPaymentResponse
`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) UnsetPaymentResponse()`

UnsetPaymentResponse ensures that no value is present for PaymentResponse, not even an explicit nil
### GetMismatchedAmounts

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetMismatchedAmounts() interface{}`

GetMismatchedAmounts returns the MismatchedAmounts field if non-nil, zero value otherwise.

### GetMismatchedAmountsOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetMismatchedAmountsOk() (*interface{}, bool)`

GetMismatchedAmountsOk returns a tuple with the MismatchedAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMismatchedAmounts

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetMismatchedAmounts(v interface{})`

SetMismatchedAmounts sets MismatchedAmounts field to given value.

### HasMismatchedAmounts

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasMismatchedAmounts() bool`

HasMismatchedAmounts returns a boolean if a field has been set.

### SetMismatchedAmountsNil

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetMismatchedAmountsNil(b bool)`

 SetMismatchedAmountsNil sets the value for MismatchedAmounts to be an explicit nil

### UnsetMismatchedAmounts
`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) UnsetMismatchedAmounts()`

UnsetMismatchedAmounts ensures that no value is present for MismatchedAmounts, not even an explicit nil
### GetPaymentInstrument

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetPaymentInstrument() interface{}`

GetPaymentInstrument returns the PaymentInstrument field if non-nil, zero value otherwise.

### GetPaymentInstrumentOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetPaymentInstrumentOk() (*interface{}, bool)`

GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrument

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetPaymentInstrument(v interface{})`

SetPaymentInstrument sets PaymentInstrument field to given value.

### HasPaymentInstrument

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasPaymentInstrument() bool`

HasPaymentInstrument returns a boolean if a field has been set.

### SetPaymentInstrumentNil

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetPaymentInstrumentNil(b bool)`

 SetPaymentInstrumentNil sets the value for PaymentInstrument to be an explicit nil

### UnsetPaymentInstrument
`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) UnsetPaymentInstrument()`

UnsetPaymentInstrument ensures that no value is present for PaymentInstrument, not even an explicit nil
### GetCreatedAt

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


