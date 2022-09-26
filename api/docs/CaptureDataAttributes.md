# CaptureDataAttributes

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

### NewCaptureDataAttributes

`func NewCaptureDataAttributes() *CaptureDataAttributes`

NewCaptureDataAttributes instantiates a new CaptureDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptureDataAttributesWithDefaults

`func NewCaptureDataAttributesWithDefaults() *CaptureDataAttributes`

NewCaptureDataAttributesWithDefaults instantiates a new CaptureDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *CaptureDataAttributes) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CaptureDataAttributes) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CaptureDataAttributes) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CaptureDataAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *CaptureDataAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *CaptureDataAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *CaptureDataAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *CaptureDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetAmountCents

`func (o *CaptureDataAttributes) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *CaptureDataAttributes) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *CaptureDataAttributes) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *CaptureDataAttributes) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetAmountFloat

`func (o *CaptureDataAttributes) GetAmountFloat() float32`

GetAmountFloat returns the AmountFloat field if non-nil, zero value otherwise.

### GetAmountFloatOk

`func (o *CaptureDataAttributes) GetAmountFloatOk() (*float32, bool)`

GetAmountFloatOk returns a tuple with the AmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountFloat

`func (o *CaptureDataAttributes) SetAmountFloat(v float32)`

SetAmountFloat sets AmountFloat field to given value.

### HasAmountFloat

`func (o *CaptureDataAttributes) HasAmountFloat() bool`

HasAmountFloat returns a boolean if a field has been set.

### GetFormattedAmount

`func (o *CaptureDataAttributes) GetFormattedAmount() string`

GetFormattedAmount returns the FormattedAmount field if non-nil, zero value otherwise.

### GetFormattedAmountOk

`func (o *CaptureDataAttributes) GetFormattedAmountOk() (*string, bool)`

GetFormattedAmountOk returns a tuple with the FormattedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAmount

`func (o *CaptureDataAttributes) SetFormattedAmount(v string)`

SetFormattedAmount sets FormattedAmount field to given value.

### HasFormattedAmount

`func (o *CaptureDataAttributes) HasFormattedAmount() bool`

HasFormattedAmount returns a boolean if a field has been set.

### GetSucceeded

`func (o *CaptureDataAttributes) GetSucceeded() bool`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *CaptureDataAttributes) GetSucceededOk() (*bool, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *CaptureDataAttributes) SetSucceeded(v bool)`

SetSucceeded sets Succeeded field to given value.

### HasSucceeded

`func (o *CaptureDataAttributes) HasSucceeded() bool`

HasSucceeded returns a boolean if a field has been set.

### GetMessage

`func (o *CaptureDataAttributes) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CaptureDataAttributes) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CaptureDataAttributes) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CaptureDataAttributes) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrorCode

`func (o *CaptureDataAttributes) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *CaptureDataAttributes) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *CaptureDataAttributes) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *CaptureDataAttributes) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDetail

`func (o *CaptureDataAttributes) GetErrorDetail() string`

GetErrorDetail returns the ErrorDetail field if non-nil, zero value otherwise.

### GetErrorDetailOk

`func (o *CaptureDataAttributes) GetErrorDetailOk() (*string, bool)`

GetErrorDetailOk returns a tuple with the ErrorDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetail

`func (o *CaptureDataAttributes) SetErrorDetail(v string)`

SetErrorDetail sets ErrorDetail field to given value.

### HasErrorDetail

`func (o *CaptureDataAttributes) HasErrorDetail() bool`

HasErrorDetail returns a boolean if a field has been set.

### GetToken

`func (o *CaptureDataAttributes) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CaptureDataAttributes) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CaptureDataAttributes) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CaptureDataAttributes) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetGatewayTransactionId

`func (o *CaptureDataAttributes) GetGatewayTransactionId() string`

GetGatewayTransactionId returns the GatewayTransactionId field if non-nil, zero value otherwise.

### GetGatewayTransactionIdOk

`func (o *CaptureDataAttributes) GetGatewayTransactionIdOk() (*string, bool)`

GetGatewayTransactionIdOk returns a tuple with the GatewayTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayTransactionId

`func (o *CaptureDataAttributes) SetGatewayTransactionId(v string)`

SetGatewayTransactionId sets GatewayTransactionId field to given value.

### HasGatewayTransactionId

`func (o *CaptureDataAttributes) HasGatewayTransactionId() bool`

HasGatewayTransactionId returns a boolean if a field has been set.

### GetId

`func (o *CaptureDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CaptureDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CaptureDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CaptureDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CaptureDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CaptureDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CaptureDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CaptureDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CaptureDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CaptureDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CaptureDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CaptureDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *CaptureDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CaptureDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CaptureDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *CaptureDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *CaptureDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *CaptureDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *CaptureDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *CaptureDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *CaptureDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CaptureDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CaptureDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CaptureDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRefundAmountCents

`func (o *CaptureDataAttributes) GetRefundAmountCents() int32`

GetRefundAmountCents returns the RefundAmountCents field if non-nil, zero value otherwise.

### GetRefundAmountCentsOk

`func (o *CaptureDataAttributes) GetRefundAmountCentsOk() (*int32, bool)`

GetRefundAmountCentsOk returns a tuple with the RefundAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmountCents

`func (o *CaptureDataAttributes) SetRefundAmountCents(v int32)`

SetRefundAmountCents sets RefundAmountCents field to given value.

### HasRefundAmountCents

`func (o *CaptureDataAttributes) HasRefundAmountCents() bool`

HasRefundAmountCents returns a boolean if a field has been set.

### GetRefundAmountFloat

`func (o *CaptureDataAttributes) GetRefundAmountFloat() float32`

GetRefundAmountFloat returns the RefundAmountFloat field if non-nil, zero value otherwise.

### GetRefundAmountFloatOk

`func (o *CaptureDataAttributes) GetRefundAmountFloatOk() (*float32, bool)`

GetRefundAmountFloatOk returns a tuple with the RefundAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmountFloat

`func (o *CaptureDataAttributes) SetRefundAmountFloat(v float32)`

SetRefundAmountFloat sets RefundAmountFloat field to given value.

### HasRefundAmountFloat

`func (o *CaptureDataAttributes) HasRefundAmountFloat() bool`

HasRefundAmountFloat returns a boolean if a field has been set.

### GetFormattedRefundAmount

`func (o *CaptureDataAttributes) GetFormattedRefundAmount() string`

GetFormattedRefundAmount returns the FormattedRefundAmount field if non-nil, zero value otherwise.

### GetFormattedRefundAmountOk

`func (o *CaptureDataAttributes) GetFormattedRefundAmountOk() (*string, bool)`

GetFormattedRefundAmountOk returns a tuple with the FormattedRefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedRefundAmount

`func (o *CaptureDataAttributes) SetFormattedRefundAmount(v string)`

SetFormattedRefundAmount sets FormattedRefundAmount field to given value.

### HasFormattedRefundAmount

`func (o *CaptureDataAttributes) HasFormattedRefundAmount() bool`

HasFormattedRefundAmount returns a boolean if a field has been set.

### GetRefundBalanceCents

`func (o *CaptureDataAttributes) GetRefundBalanceCents() int32`

GetRefundBalanceCents returns the RefundBalanceCents field if non-nil, zero value otherwise.

### GetRefundBalanceCentsOk

`func (o *CaptureDataAttributes) GetRefundBalanceCentsOk() (*int32, bool)`

GetRefundBalanceCentsOk returns a tuple with the RefundBalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundBalanceCents

`func (o *CaptureDataAttributes) SetRefundBalanceCents(v int32)`

SetRefundBalanceCents sets RefundBalanceCents field to given value.

### HasRefundBalanceCents

`func (o *CaptureDataAttributes) HasRefundBalanceCents() bool`

HasRefundBalanceCents returns a boolean if a field has been set.

### GetRefundBalanceFloat

`func (o *CaptureDataAttributes) GetRefundBalanceFloat() float32`

GetRefundBalanceFloat returns the RefundBalanceFloat field if non-nil, zero value otherwise.

### GetRefundBalanceFloatOk

`func (o *CaptureDataAttributes) GetRefundBalanceFloatOk() (*float32, bool)`

GetRefundBalanceFloatOk returns a tuple with the RefundBalanceFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundBalanceFloat

`func (o *CaptureDataAttributes) SetRefundBalanceFloat(v float32)`

SetRefundBalanceFloat sets RefundBalanceFloat field to given value.

### HasRefundBalanceFloat

`func (o *CaptureDataAttributes) HasRefundBalanceFloat() bool`

HasRefundBalanceFloat returns a boolean if a field has been set.

### GetFormattedRefundBalance

`func (o *CaptureDataAttributes) GetFormattedRefundBalance() string`

GetFormattedRefundBalance returns the FormattedRefundBalance field if non-nil, zero value otherwise.

### GetFormattedRefundBalanceOk

`func (o *CaptureDataAttributes) GetFormattedRefundBalanceOk() (*string, bool)`

GetFormattedRefundBalanceOk returns a tuple with the FormattedRefundBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedRefundBalance

`func (o *CaptureDataAttributes) SetFormattedRefundBalance(v string)`

SetFormattedRefundBalance sets FormattedRefundBalance field to given value.

### HasFormattedRefundBalance

`func (o *CaptureDataAttributes) HasFormattedRefundBalance() bool`

HasFormattedRefundBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


