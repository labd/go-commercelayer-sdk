# KlarnaPaymentDataAttributes

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
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewKlarnaPaymentDataAttributes

`func NewKlarnaPaymentDataAttributes() *KlarnaPaymentDataAttributes`

NewKlarnaPaymentDataAttributes instantiates a new KlarnaPaymentDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKlarnaPaymentDataAttributesWithDefaults

`func NewKlarnaPaymentDataAttributesWithDefaults() *KlarnaPaymentDataAttributes`

NewKlarnaPaymentDataAttributesWithDefaults instantiates a new KlarnaPaymentDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *KlarnaPaymentDataAttributes) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *KlarnaPaymentDataAttributes) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *KlarnaPaymentDataAttributes) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *KlarnaPaymentDataAttributes) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetClientToken

`func (o *KlarnaPaymentDataAttributes) GetClientToken() string`

GetClientToken returns the ClientToken field if non-nil, zero value otherwise.

### GetClientTokenOk

`func (o *KlarnaPaymentDataAttributes) GetClientTokenOk() (*string, bool)`

GetClientTokenOk returns a tuple with the ClientToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientToken

`func (o *KlarnaPaymentDataAttributes) SetClientToken(v string)`

SetClientToken sets ClientToken field to given value.

### HasClientToken

`func (o *KlarnaPaymentDataAttributes) HasClientToken() bool`

HasClientToken returns a boolean if a field has been set.

### GetPaymentMethods

`func (o *KlarnaPaymentDataAttributes) GetPaymentMethods() []map[string]interface{}`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *KlarnaPaymentDataAttributes) GetPaymentMethodsOk() (*[]map[string]interface{}, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *KlarnaPaymentDataAttributes) SetPaymentMethods(v []map[string]interface{})`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *KlarnaPaymentDataAttributes) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### GetAuthToken

`func (o *KlarnaPaymentDataAttributes) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *KlarnaPaymentDataAttributes) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *KlarnaPaymentDataAttributes) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *KlarnaPaymentDataAttributes) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### GetMismatchedAmounts

`func (o *KlarnaPaymentDataAttributes) GetMismatchedAmounts() bool`

GetMismatchedAmounts returns the MismatchedAmounts field if non-nil, zero value otherwise.

### GetMismatchedAmountsOk

`func (o *KlarnaPaymentDataAttributes) GetMismatchedAmountsOk() (*bool, bool)`

GetMismatchedAmountsOk returns a tuple with the MismatchedAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMismatchedAmounts

`func (o *KlarnaPaymentDataAttributes) SetMismatchedAmounts(v bool)`

SetMismatchedAmounts sets MismatchedAmounts field to given value.

### HasMismatchedAmounts

`func (o *KlarnaPaymentDataAttributes) HasMismatchedAmounts() bool`

HasMismatchedAmounts returns a boolean if a field has been set.

### GetIntentAmountCents

`func (o *KlarnaPaymentDataAttributes) GetIntentAmountCents() int32`

GetIntentAmountCents returns the IntentAmountCents field if non-nil, zero value otherwise.

### GetIntentAmountCentsOk

`func (o *KlarnaPaymentDataAttributes) GetIntentAmountCentsOk() (*int32, bool)`

GetIntentAmountCentsOk returns a tuple with the IntentAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentAmountCents

`func (o *KlarnaPaymentDataAttributes) SetIntentAmountCents(v int32)`

SetIntentAmountCents sets IntentAmountCents field to given value.

### HasIntentAmountCents

`func (o *KlarnaPaymentDataAttributes) HasIntentAmountCents() bool`

HasIntentAmountCents returns a boolean if a field has been set.

### GetIntentAmountFloat

`func (o *KlarnaPaymentDataAttributes) GetIntentAmountFloat() float32`

GetIntentAmountFloat returns the IntentAmountFloat field if non-nil, zero value otherwise.

### GetIntentAmountFloatOk

`func (o *KlarnaPaymentDataAttributes) GetIntentAmountFloatOk() (*float32, bool)`

GetIntentAmountFloatOk returns a tuple with the IntentAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentAmountFloat

`func (o *KlarnaPaymentDataAttributes) SetIntentAmountFloat(v float32)`

SetIntentAmountFloat sets IntentAmountFloat field to given value.

### HasIntentAmountFloat

`func (o *KlarnaPaymentDataAttributes) HasIntentAmountFloat() bool`

HasIntentAmountFloat returns a boolean if a field has been set.

### GetFormattedIntentAmount

`func (o *KlarnaPaymentDataAttributes) GetFormattedIntentAmount() string`

GetFormattedIntentAmount returns the FormattedIntentAmount field if non-nil, zero value otherwise.

### GetFormattedIntentAmountOk

`func (o *KlarnaPaymentDataAttributes) GetFormattedIntentAmountOk() (*string, bool)`

GetFormattedIntentAmountOk returns a tuple with the FormattedIntentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedIntentAmount

`func (o *KlarnaPaymentDataAttributes) SetFormattedIntentAmount(v string)`

SetFormattedIntentAmount sets FormattedIntentAmount field to given value.

### HasFormattedIntentAmount

`func (o *KlarnaPaymentDataAttributes) HasFormattedIntentAmount() bool`

HasFormattedIntentAmount returns a boolean if a field has been set.

### GetId

`func (o *KlarnaPaymentDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KlarnaPaymentDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KlarnaPaymentDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KlarnaPaymentDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KlarnaPaymentDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KlarnaPaymentDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KlarnaPaymentDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KlarnaPaymentDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KlarnaPaymentDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KlarnaPaymentDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KlarnaPaymentDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KlarnaPaymentDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *KlarnaPaymentDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *KlarnaPaymentDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *KlarnaPaymentDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *KlarnaPaymentDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *KlarnaPaymentDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *KlarnaPaymentDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *KlarnaPaymentDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *KlarnaPaymentDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *KlarnaPaymentDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KlarnaPaymentDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KlarnaPaymentDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KlarnaPaymentDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


