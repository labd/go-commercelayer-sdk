# GETCaptures200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **string** | The transaction number, auto generated | [optional] 
**CurrencyCode** | Pointer to **string** | The international 3-letter currency code as defined by the ISO 4217 standard, inherited from the associated order. | [optional] 
**AmountCents** | Pointer to **int32** | The transaction amount, in cents. | [optional] 
**AmountFloat** | Pointer to **float32** | The transaction amount, float. | [optional] 
**FormattedAmount** | Pointer to **string** | The transaction amount, formatted. | [optional] 
**Succeeded** | Pointer to **bool** | Indicates if the transaction is successful | [optional] 
**Message** | Pointer to **string** | The message returned by the payment gateway | [optional] 
**ErrorCode** | Pointer to **string** | The error code, if any, returned by the payment gateway | [optional] 
**ErrorDetail** | Pointer to **string** | The error detail, if any, returned by the payment gateway | [optional] 
**Token** | Pointer to **string** | The token identifying the transaction, returned by the payment gateway | [optional] 
**GatewayTransactionId** | Pointer to **string** | The ID identifying the transaction, returned by the payment gateway | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**RefundAmountCents** | Pointer to **int32** | The amount to be refunded, in cents. | [optional] 
**RefundAmountFloat** | Pointer to **float32** | The amount to be refunded, float. | [optional] 
**FormattedRefundAmount** | Pointer to **string** | The amount to be refunded, formatted. | [optional] 
**RefundBalanceCents** | Pointer to **int32** | The balance to be refunded, in cents. | [optional] 
**RefundBalanceFloat** | Pointer to **float32** | The balance to be refunded, float. | [optional] 
**FormattedRefundBalance** | Pointer to **string** | The balance to be refunded, formatted. | [optional] 

## Methods

### NewGETCaptures200ResponseDataInnerAttributes

`func NewGETCaptures200ResponseDataInnerAttributes() *GETCaptures200ResponseDataInnerAttributes`

NewGETCaptures200ResponseDataInnerAttributes instantiates a new GETCaptures200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETCaptures200ResponseDataInnerAttributesWithDefaults

`func NewGETCaptures200ResponseDataInnerAttributesWithDefaults() *GETCaptures200ResponseDataInnerAttributes`

NewGETCaptures200ResponseDataInnerAttributesWithDefaults instantiates a new GETCaptures200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *GETCaptures200ResponseDataInnerAttributes) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GETCaptures200ResponseDataInnerAttributes) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GETCaptures200ResponseDataInnerAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *GETCaptures200ResponseDataInnerAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETCaptures200ResponseDataInnerAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETCaptures200ResponseDataInnerAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetAmountCents

`func (o *GETCaptures200ResponseDataInnerAttributes) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *GETCaptures200ResponseDataInnerAttributes) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *GETCaptures200ResponseDataInnerAttributes) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetAmountFloat

`func (o *GETCaptures200ResponseDataInnerAttributes) GetAmountFloat() float32`

GetAmountFloat returns the AmountFloat field if non-nil, zero value otherwise.

### GetAmountFloatOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetAmountFloatOk() (*float32, bool)`

GetAmountFloatOk returns a tuple with the AmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountFloat

`func (o *GETCaptures200ResponseDataInnerAttributes) SetAmountFloat(v float32)`

SetAmountFloat sets AmountFloat field to given value.

### HasAmountFloat

`func (o *GETCaptures200ResponseDataInnerAttributes) HasAmountFloat() bool`

HasAmountFloat returns a boolean if a field has been set.

### GetFormattedAmount

`func (o *GETCaptures200ResponseDataInnerAttributes) GetFormattedAmount() string`

GetFormattedAmount returns the FormattedAmount field if non-nil, zero value otherwise.

### GetFormattedAmountOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetFormattedAmountOk() (*string, bool)`

GetFormattedAmountOk returns a tuple with the FormattedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAmount

`func (o *GETCaptures200ResponseDataInnerAttributes) SetFormattedAmount(v string)`

SetFormattedAmount sets FormattedAmount field to given value.

### HasFormattedAmount

`func (o *GETCaptures200ResponseDataInnerAttributes) HasFormattedAmount() bool`

HasFormattedAmount returns a boolean if a field has been set.

### GetSucceeded

`func (o *GETCaptures200ResponseDataInnerAttributes) GetSucceeded() bool`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetSucceededOk() (*bool, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *GETCaptures200ResponseDataInnerAttributes) SetSucceeded(v bool)`

SetSucceeded sets Succeeded field to given value.

### HasSucceeded

`func (o *GETCaptures200ResponseDataInnerAttributes) HasSucceeded() bool`

HasSucceeded returns a boolean if a field has been set.

### GetMessage

`func (o *GETCaptures200ResponseDataInnerAttributes) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GETCaptures200ResponseDataInnerAttributes) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GETCaptures200ResponseDataInnerAttributes) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrorCode

`func (o *GETCaptures200ResponseDataInnerAttributes) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GETCaptures200ResponseDataInnerAttributes) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *GETCaptures200ResponseDataInnerAttributes) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDetail

`func (o *GETCaptures200ResponseDataInnerAttributes) GetErrorDetail() string`

GetErrorDetail returns the ErrorDetail field if non-nil, zero value otherwise.

### GetErrorDetailOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetErrorDetailOk() (*string, bool)`

GetErrorDetailOk returns a tuple with the ErrorDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetail

`func (o *GETCaptures200ResponseDataInnerAttributes) SetErrorDetail(v string)`

SetErrorDetail sets ErrorDetail field to given value.

### HasErrorDetail

`func (o *GETCaptures200ResponseDataInnerAttributes) HasErrorDetail() bool`

HasErrorDetail returns a boolean if a field has been set.

### GetToken

`func (o *GETCaptures200ResponseDataInnerAttributes) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GETCaptures200ResponseDataInnerAttributes) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GETCaptures200ResponseDataInnerAttributes) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetGatewayTransactionId

`func (o *GETCaptures200ResponseDataInnerAttributes) GetGatewayTransactionId() string`

GetGatewayTransactionId returns the GatewayTransactionId field if non-nil, zero value otherwise.

### GetGatewayTransactionIdOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetGatewayTransactionIdOk() (*string, bool)`

GetGatewayTransactionIdOk returns a tuple with the GatewayTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayTransactionId

`func (o *GETCaptures200ResponseDataInnerAttributes) SetGatewayTransactionId(v string)`

SetGatewayTransactionId sets GatewayTransactionId field to given value.

### HasGatewayTransactionId

`func (o *GETCaptures200ResponseDataInnerAttributes) HasGatewayTransactionId() bool`

HasGatewayTransactionId returns a boolean if a field has been set.

### GetId

`func (o *GETCaptures200ResponseDataInnerAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETCaptures200ResponseDataInnerAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETCaptures200ResponseDataInnerAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETCaptures200ResponseDataInnerAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETCaptures200ResponseDataInnerAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETCaptures200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETCaptures200ResponseDataInnerAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETCaptures200ResponseDataInnerAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETCaptures200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETCaptures200ResponseDataInnerAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETCaptures200ResponseDataInnerAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETCaptures200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETCaptures200ResponseDataInnerAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETCaptures200ResponseDataInnerAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETCaptures200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETCaptures200ResponseDataInnerAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETCaptures200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETCaptures200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRefundAmountCents

`func (o *GETCaptures200ResponseDataInnerAttributes) GetRefundAmountCents() int32`

GetRefundAmountCents returns the RefundAmountCents field if non-nil, zero value otherwise.

### GetRefundAmountCentsOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetRefundAmountCentsOk() (*int32, bool)`

GetRefundAmountCentsOk returns a tuple with the RefundAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmountCents

`func (o *GETCaptures200ResponseDataInnerAttributes) SetRefundAmountCents(v int32)`

SetRefundAmountCents sets RefundAmountCents field to given value.

### HasRefundAmountCents

`func (o *GETCaptures200ResponseDataInnerAttributes) HasRefundAmountCents() bool`

HasRefundAmountCents returns a boolean if a field has been set.

### GetRefundAmountFloat

`func (o *GETCaptures200ResponseDataInnerAttributes) GetRefundAmountFloat() float32`

GetRefundAmountFloat returns the RefundAmountFloat field if non-nil, zero value otherwise.

### GetRefundAmountFloatOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetRefundAmountFloatOk() (*float32, bool)`

GetRefundAmountFloatOk returns a tuple with the RefundAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmountFloat

`func (o *GETCaptures200ResponseDataInnerAttributes) SetRefundAmountFloat(v float32)`

SetRefundAmountFloat sets RefundAmountFloat field to given value.

### HasRefundAmountFloat

`func (o *GETCaptures200ResponseDataInnerAttributes) HasRefundAmountFloat() bool`

HasRefundAmountFloat returns a boolean if a field has been set.

### GetFormattedRefundAmount

`func (o *GETCaptures200ResponseDataInnerAttributes) GetFormattedRefundAmount() string`

GetFormattedRefundAmount returns the FormattedRefundAmount field if non-nil, zero value otherwise.

### GetFormattedRefundAmountOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetFormattedRefundAmountOk() (*string, bool)`

GetFormattedRefundAmountOk returns a tuple with the FormattedRefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedRefundAmount

`func (o *GETCaptures200ResponseDataInnerAttributes) SetFormattedRefundAmount(v string)`

SetFormattedRefundAmount sets FormattedRefundAmount field to given value.

### HasFormattedRefundAmount

`func (o *GETCaptures200ResponseDataInnerAttributes) HasFormattedRefundAmount() bool`

HasFormattedRefundAmount returns a boolean if a field has been set.

### GetRefundBalanceCents

`func (o *GETCaptures200ResponseDataInnerAttributes) GetRefundBalanceCents() int32`

GetRefundBalanceCents returns the RefundBalanceCents field if non-nil, zero value otherwise.

### GetRefundBalanceCentsOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetRefundBalanceCentsOk() (*int32, bool)`

GetRefundBalanceCentsOk returns a tuple with the RefundBalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundBalanceCents

`func (o *GETCaptures200ResponseDataInnerAttributes) SetRefundBalanceCents(v int32)`

SetRefundBalanceCents sets RefundBalanceCents field to given value.

### HasRefundBalanceCents

`func (o *GETCaptures200ResponseDataInnerAttributes) HasRefundBalanceCents() bool`

HasRefundBalanceCents returns a boolean if a field has been set.

### GetRefundBalanceFloat

`func (o *GETCaptures200ResponseDataInnerAttributes) GetRefundBalanceFloat() float32`

GetRefundBalanceFloat returns the RefundBalanceFloat field if non-nil, zero value otherwise.

### GetRefundBalanceFloatOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetRefundBalanceFloatOk() (*float32, bool)`

GetRefundBalanceFloatOk returns a tuple with the RefundBalanceFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundBalanceFloat

`func (o *GETCaptures200ResponseDataInnerAttributes) SetRefundBalanceFloat(v float32)`

SetRefundBalanceFloat sets RefundBalanceFloat field to given value.

### HasRefundBalanceFloat

`func (o *GETCaptures200ResponseDataInnerAttributes) HasRefundBalanceFloat() bool`

HasRefundBalanceFloat returns a boolean if a field has been set.

### GetFormattedRefundBalance

`func (o *GETCaptures200ResponseDataInnerAttributes) GetFormattedRefundBalance() string`

GetFormattedRefundBalance returns the FormattedRefundBalance field if non-nil, zero value otherwise.

### GetFormattedRefundBalanceOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetFormattedRefundBalanceOk() (*string, bool)`

GetFormattedRefundBalanceOk returns a tuple with the FormattedRefundBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedRefundBalance

`func (o *GETCaptures200ResponseDataInnerAttributes) SetFormattedRefundBalance(v string)`

SetFormattedRefundBalance sets FormattedRefundBalance field to given value.

### HasFormattedRefundBalance

`func (o *GETCaptures200ResponseDataInnerAttributes) HasFormattedRefundBalance() bool`

HasFormattedRefundBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


