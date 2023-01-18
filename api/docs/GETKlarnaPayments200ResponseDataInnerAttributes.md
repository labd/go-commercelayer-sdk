# GETKlarnaPayments200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | Pointer to **string** | The identifier of the payment session, useful to updated it. | [optional] 
**ClientToken** | Pointer to **string** | The public token linked to your API credential. Available upon session creation. | [optional] 
**PaymentMethods** | Pointer to **[]map[string]interface{}** | The merchant available payment methods for the assoiated order. Available upon session creation. | [optional] 
**AuthToken** | Pointer to **string** | The token returned by a successful client authorization, mandatory to place the order. | [optional] 
**MismatchedAmounts** | Pointer to **bool** | Indicates if the order current amount differs form the one of the created payment intent. | [optional] 
**IntentAmountCents** | Pointer to **int32** | The amount of the associated payment intent, in cents. | [optional] 
**IntentAmountFloat** | Pointer to **float32** | The amount of the associated payment intent, float. | [optional] 
**FormattedIntentAmount** | Pointer to **string** | The amount of the associated payment intent, formatted. | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

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

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetClientToken

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetClientToken() string`

GetClientToken returns the ClientToken field if non-nil, zero value otherwise.

### GetClientTokenOk

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetClientTokenOk() (*string, bool)`

GetClientTokenOk returns a tuple with the ClientToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientToken

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetClientToken(v string)`

SetClientToken sets ClientToken field to given value.

### HasClientToken

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) HasClientToken() bool`

HasClientToken returns a boolean if a field has been set.

### GetPaymentMethods

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetPaymentMethods() []map[string]interface{}`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetPaymentMethodsOk() (*[]map[string]interface{}, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetPaymentMethods(v []map[string]interface{})`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### GetAuthToken

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### GetMismatchedAmounts

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetMismatchedAmounts() bool`

GetMismatchedAmounts returns the MismatchedAmounts field if non-nil, zero value otherwise.

### GetMismatchedAmountsOk

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetMismatchedAmountsOk() (*bool, bool)`

GetMismatchedAmountsOk returns a tuple with the MismatchedAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMismatchedAmounts

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetMismatchedAmounts(v bool)`

SetMismatchedAmounts sets MismatchedAmounts field to given value.

### HasMismatchedAmounts

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) HasMismatchedAmounts() bool`

HasMismatchedAmounts returns a boolean if a field has been set.

### GetIntentAmountCents

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetIntentAmountCents() int32`

GetIntentAmountCents returns the IntentAmountCents field if non-nil, zero value otherwise.

### GetIntentAmountCentsOk

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetIntentAmountCentsOk() (*int32, bool)`

GetIntentAmountCentsOk returns a tuple with the IntentAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentAmountCents

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetIntentAmountCents(v int32)`

SetIntentAmountCents sets IntentAmountCents field to given value.

### HasIntentAmountCents

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) HasIntentAmountCents() bool`

HasIntentAmountCents returns a boolean if a field has been set.

### GetIntentAmountFloat

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetIntentAmountFloat() float32`

GetIntentAmountFloat returns the IntentAmountFloat field if non-nil, zero value otherwise.

### GetIntentAmountFloatOk

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetIntentAmountFloatOk() (*float32, bool)`

GetIntentAmountFloatOk returns a tuple with the IntentAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentAmountFloat

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetIntentAmountFloat(v float32)`

SetIntentAmountFloat sets IntentAmountFloat field to given value.

### HasIntentAmountFloat

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) HasIntentAmountFloat() bool`

HasIntentAmountFloat returns a boolean if a field has been set.

### GetFormattedIntentAmount

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetFormattedIntentAmount() string`

GetFormattedIntentAmount returns the FormattedIntentAmount field if non-nil, zero value otherwise.

### GetFormattedIntentAmountOk

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetFormattedIntentAmountOk() (*string, bool)`

GetFormattedIntentAmountOk returns a tuple with the FormattedIntentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedIntentAmount

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetFormattedIntentAmount(v string)`

SetFormattedIntentAmount sets FormattedIntentAmount field to given value.

### HasFormattedIntentAmount

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) HasFormattedIntentAmount() bool`

HasFormattedIntentAmount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETKlarnaPayments200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


