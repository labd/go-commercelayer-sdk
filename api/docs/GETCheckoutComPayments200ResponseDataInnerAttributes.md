# GETCheckoutComPayments200ResponseDataInnerAttributes

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
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

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

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetPaymentType

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.

### GetToken

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetSessionId

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetSuccessUrl

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.

### HasSuccessUrl

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasSuccessUrl() bool`

HasSuccessUrl returns a boolean if a field has been set.

### GetFailureUrl

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetFailureUrl() string`

GetFailureUrl returns the FailureUrl field if non-nil, zero value otherwise.

### GetFailureUrlOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetFailureUrlOk() (*string, bool)`

GetFailureUrlOk returns a tuple with the FailureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureUrl

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetFailureUrl(v string)`

SetFailureUrl sets FailureUrl field to given value.

### HasFailureUrl

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasFailureUrl() bool`

HasFailureUrl returns a boolean if a field has been set.

### GetSourceId

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetCustomerToken

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetCustomerToken() string`

GetCustomerToken returns the CustomerToken field if non-nil, zero value otherwise.

### GetCustomerTokenOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetCustomerTokenOk() (*string, bool)`

GetCustomerTokenOk returns a tuple with the CustomerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerToken

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetCustomerToken(v string)`

SetCustomerToken sets CustomerToken field to given value.

### HasCustomerToken

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasCustomerToken() bool`

HasCustomerToken returns a boolean if a field has been set.

### GetRedirectUri

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### GetPaymentResponse

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetPaymentResponse() map[string]interface{}`

GetPaymentResponse returns the PaymentResponse field if non-nil, zero value otherwise.

### GetPaymentResponseOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetPaymentResponseOk() (*map[string]interface{}, bool)`

GetPaymentResponseOk returns a tuple with the PaymentResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentResponse

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetPaymentResponse(v map[string]interface{})`

SetPaymentResponse sets PaymentResponse field to given value.

### HasPaymentResponse

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasPaymentResponse() bool`

HasPaymentResponse returns a boolean if a field has been set.

### GetMismatchedAmounts

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetMismatchedAmounts() bool`

GetMismatchedAmounts returns the MismatchedAmounts field if non-nil, zero value otherwise.

### GetMismatchedAmountsOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetMismatchedAmountsOk() (*bool, bool)`

GetMismatchedAmountsOk returns a tuple with the MismatchedAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMismatchedAmounts

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetMismatchedAmounts(v bool)`

SetMismatchedAmounts sets MismatchedAmounts field to given value.

### HasMismatchedAmounts

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasMismatchedAmounts() bool`

HasMismatchedAmounts returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


