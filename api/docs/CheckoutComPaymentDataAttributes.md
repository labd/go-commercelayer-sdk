# CheckoutComPaymentDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicKey** | Pointer to **string** | The Checkout.com publishable API key. | [optional] 
**PaymentType** | Pointer to **string** | The payment source type. | [optional] 
**Token** | Pointer to **string** | The Checkout.com card or digital wallet token. | [optional] 
**SessionId** | Pointer to **string** | A payment session ID used to obtain the details. | [optional] 
**SuccessUrl** | Pointer to **string** | The URL to redirect your customer upon 3DS succeeded authentication. | [optional] 
**FailureUrl** | Pointer to **string** | The URL to redirect your customer upon 3DS failed authentication. | [optional] 
**SourceId** | Pointer to **string** | The payment source identifier that can be used for subsequent payments. | [optional] 
**CustomerToken** | Pointer to **string** | The customer&#39;s unique identifier. This can be passed as a source when making a payment. | [optional] 
**RedirectUri** | Pointer to **string** | The URI that the customer should be redirected to in order to complete the payment. | [optional] 
**PaymentResponse** | Pointer to **map[string]interface{}** | The Checkout.com payment response, used to fetch internal data. | [optional] 
**MismatchedAmounts** | Pointer to **bool** | Indicates if the order current amount differs form the one of the associated authorization. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewCheckoutComPaymentDataAttributes

`func NewCheckoutComPaymentDataAttributes() *CheckoutComPaymentDataAttributes`

NewCheckoutComPaymentDataAttributes instantiates a new CheckoutComPaymentDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutComPaymentDataAttributesWithDefaults

`func NewCheckoutComPaymentDataAttributesWithDefaults() *CheckoutComPaymentDataAttributes`

NewCheckoutComPaymentDataAttributesWithDefaults instantiates a new CheckoutComPaymentDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicKey

`func (o *CheckoutComPaymentDataAttributes) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *CheckoutComPaymentDataAttributes) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *CheckoutComPaymentDataAttributes) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *CheckoutComPaymentDataAttributes) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetPaymentType

`func (o *CheckoutComPaymentDataAttributes) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *CheckoutComPaymentDataAttributes) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *CheckoutComPaymentDataAttributes) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *CheckoutComPaymentDataAttributes) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.

### GetToken

`func (o *CheckoutComPaymentDataAttributes) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CheckoutComPaymentDataAttributes) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CheckoutComPaymentDataAttributes) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CheckoutComPaymentDataAttributes) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetSessionId

`func (o *CheckoutComPaymentDataAttributes) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *CheckoutComPaymentDataAttributes) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *CheckoutComPaymentDataAttributes) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *CheckoutComPaymentDataAttributes) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetSuccessUrl

`func (o *CheckoutComPaymentDataAttributes) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *CheckoutComPaymentDataAttributes) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *CheckoutComPaymentDataAttributes) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.

### HasSuccessUrl

`func (o *CheckoutComPaymentDataAttributes) HasSuccessUrl() bool`

HasSuccessUrl returns a boolean if a field has been set.

### GetFailureUrl

`func (o *CheckoutComPaymentDataAttributes) GetFailureUrl() string`

GetFailureUrl returns the FailureUrl field if non-nil, zero value otherwise.

### GetFailureUrlOk

`func (o *CheckoutComPaymentDataAttributes) GetFailureUrlOk() (*string, bool)`

GetFailureUrlOk returns a tuple with the FailureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureUrl

`func (o *CheckoutComPaymentDataAttributes) SetFailureUrl(v string)`

SetFailureUrl sets FailureUrl field to given value.

### HasFailureUrl

`func (o *CheckoutComPaymentDataAttributes) HasFailureUrl() bool`

HasFailureUrl returns a boolean if a field has been set.

### GetSourceId

`func (o *CheckoutComPaymentDataAttributes) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *CheckoutComPaymentDataAttributes) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *CheckoutComPaymentDataAttributes) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *CheckoutComPaymentDataAttributes) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetCustomerToken

`func (o *CheckoutComPaymentDataAttributes) GetCustomerToken() string`

GetCustomerToken returns the CustomerToken field if non-nil, zero value otherwise.

### GetCustomerTokenOk

`func (o *CheckoutComPaymentDataAttributes) GetCustomerTokenOk() (*string, bool)`

GetCustomerTokenOk returns a tuple with the CustomerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerToken

`func (o *CheckoutComPaymentDataAttributes) SetCustomerToken(v string)`

SetCustomerToken sets CustomerToken field to given value.

### HasCustomerToken

`func (o *CheckoutComPaymentDataAttributes) HasCustomerToken() bool`

HasCustomerToken returns a boolean if a field has been set.

### GetRedirectUri

`func (o *CheckoutComPaymentDataAttributes) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *CheckoutComPaymentDataAttributes) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *CheckoutComPaymentDataAttributes) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *CheckoutComPaymentDataAttributes) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### GetPaymentResponse

`func (o *CheckoutComPaymentDataAttributes) GetPaymentResponse() map[string]interface{}`

GetPaymentResponse returns the PaymentResponse field if non-nil, zero value otherwise.

### GetPaymentResponseOk

`func (o *CheckoutComPaymentDataAttributes) GetPaymentResponseOk() (*map[string]interface{}, bool)`

GetPaymentResponseOk returns a tuple with the PaymentResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentResponse

`func (o *CheckoutComPaymentDataAttributes) SetPaymentResponse(v map[string]interface{})`

SetPaymentResponse sets PaymentResponse field to given value.

### HasPaymentResponse

`func (o *CheckoutComPaymentDataAttributes) HasPaymentResponse() bool`

HasPaymentResponse returns a boolean if a field has been set.

### GetMismatchedAmounts

`func (o *CheckoutComPaymentDataAttributes) GetMismatchedAmounts() bool`

GetMismatchedAmounts returns the MismatchedAmounts field if non-nil, zero value otherwise.

### GetMismatchedAmountsOk

`func (o *CheckoutComPaymentDataAttributes) GetMismatchedAmountsOk() (*bool, bool)`

GetMismatchedAmountsOk returns a tuple with the MismatchedAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMismatchedAmounts

`func (o *CheckoutComPaymentDataAttributes) SetMismatchedAmounts(v bool)`

SetMismatchedAmounts sets MismatchedAmounts field to given value.

### HasMismatchedAmounts

`func (o *CheckoutComPaymentDataAttributes) HasMismatchedAmounts() bool`

HasMismatchedAmounts returns a boolean if a field has been set.

### GetId

`func (o *CheckoutComPaymentDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckoutComPaymentDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckoutComPaymentDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CheckoutComPaymentDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CheckoutComPaymentDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CheckoutComPaymentDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CheckoutComPaymentDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CheckoutComPaymentDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CheckoutComPaymentDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CheckoutComPaymentDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CheckoutComPaymentDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CheckoutComPaymentDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *CheckoutComPaymentDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CheckoutComPaymentDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CheckoutComPaymentDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *CheckoutComPaymentDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *CheckoutComPaymentDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *CheckoutComPaymentDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *CheckoutComPaymentDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *CheckoutComPaymentDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *CheckoutComPaymentDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CheckoutComPaymentDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CheckoutComPaymentDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CheckoutComPaymentDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


