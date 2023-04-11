# GETCaptures200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **interface{}** | The transaction number, auto generated | [optional] 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard, inherited from the associated order. | [optional] 
**AmountCents** | Pointer to **interface{}** | The transaction amount, in cents. | [optional] 
**AmountFloat** | Pointer to **interface{}** | The transaction amount, float. | [optional] 
**FormattedAmount** | Pointer to **interface{}** | The transaction amount, formatted. | [optional] 
**Succeeded** | Pointer to **interface{}** | Indicates if the transaction is successful | [optional] 
**Message** | Pointer to **interface{}** | The message returned by the payment gateway | [optional] 
**ErrorCode** | Pointer to **interface{}** | The error code, if any, returned by the payment gateway | [optional] 
**ErrorDetail** | Pointer to **interface{}** | The error detail, if any, returned by the payment gateway | [optional] 
**Token** | Pointer to **interface{}** | The token identifying the transaction, returned by the payment gateway | [optional] 
**GatewayTransactionId** | Pointer to **interface{}** | The ID identifying the transaction, returned by the payment gateway | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**RefundAmountCents** | Pointer to **interface{}** | The amount to be refunded, in cents. | [optional] 
**RefundAmountFloat** | Pointer to **interface{}** | The amount to be refunded, float. | [optional] 
**FormattedRefundAmount** | Pointer to **interface{}** | The amount to be refunded, formatted. | [optional] 
**RefundBalanceCents** | Pointer to **interface{}** | The balance to be refunded, in cents. | [optional] 
**RefundBalanceFloat** | Pointer to **interface{}** | The balance to be refunded, float. | [optional] 
**FormattedRefundBalance** | Pointer to **interface{}** | The balance to be refunded, formatted. | [optional] 

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

`func (o *GETCaptures200ResponseDataInnerAttributes) GetNumber() interface{}`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetNumberOk() (*interface{}, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GETCaptures200ResponseDataInnerAttributes) SetNumber(v interface{})`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GETCaptures200ResponseDataInnerAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *GETCaptures200ResponseDataInnerAttributes) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *GETCaptures200ResponseDataInnerAttributes) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetCurrencyCode

`func (o *GETCaptures200ResponseDataInnerAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETCaptures200ResponseDataInnerAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETCaptures200ResponseDataInnerAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *GETCaptures200ResponseDataInnerAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *GETCaptures200ResponseDataInnerAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetAmountCents

`func (o *GETCaptures200ResponseDataInnerAttributes) GetAmountCents() interface{}`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetAmountCentsOk() (*interface{}, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *GETCaptures200ResponseDataInnerAttributes) SetAmountCents(v interface{})`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *GETCaptures200ResponseDataInnerAttributes) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### SetAmountCentsNil

`func (o *GETCaptures200ResponseDataInnerAttributes) SetAmountCentsNil(b bool)`

 SetAmountCentsNil sets the value for AmountCents to be an explicit nil

### UnsetAmountCents
`func (o *GETCaptures200ResponseDataInnerAttributes) UnsetAmountCents()`

UnsetAmountCents ensures that no value is present for AmountCents, not even an explicit nil
### GetAmountFloat

`func (o *GETCaptures200ResponseDataInnerAttributes) GetAmountFloat() interface{}`

GetAmountFloat returns the AmountFloat field if non-nil, zero value otherwise.

### GetAmountFloatOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetAmountFloatOk() (*interface{}, bool)`

GetAmountFloatOk returns a tuple with the AmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountFloat

`func (o *GETCaptures200ResponseDataInnerAttributes) SetAmountFloat(v interface{})`

SetAmountFloat sets AmountFloat field to given value.

### HasAmountFloat

`func (o *GETCaptures200ResponseDataInnerAttributes) HasAmountFloat() bool`

HasAmountFloat returns a boolean if a field has been set.

### SetAmountFloatNil

`func (o *GETCaptures200ResponseDataInnerAttributes) SetAmountFloatNil(b bool)`

 SetAmountFloatNil sets the value for AmountFloat to be an explicit nil

### UnsetAmountFloat
`func (o *GETCaptures200ResponseDataInnerAttributes) UnsetAmountFloat()`

UnsetAmountFloat ensures that no value is present for AmountFloat, not even an explicit nil
### GetFormattedAmount

`func (o *GETCaptures200ResponseDataInnerAttributes) GetFormattedAmount() interface{}`

GetFormattedAmount returns the FormattedAmount field if non-nil, zero value otherwise.

### GetFormattedAmountOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetFormattedAmountOk() (*interface{}, bool)`

GetFormattedAmountOk returns a tuple with the FormattedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAmount

`func (o *GETCaptures200ResponseDataInnerAttributes) SetFormattedAmount(v interface{})`

SetFormattedAmount sets FormattedAmount field to given value.

### HasFormattedAmount

`func (o *GETCaptures200ResponseDataInnerAttributes) HasFormattedAmount() bool`

HasFormattedAmount returns a boolean if a field has been set.

### SetFormattedAmountNil

`func (o *GETCaptures200ResponseDataInnerAttributes) SetFormattedAmountNil(b bool)`

 SetFormattedAmountNil sets the value for FormattedAmount to be an explicit nil

### UnsetFormattedAmount
`func (o *GETCaptures200ResponseDataInnerAttributes) UnsetFormattedAmount()`

UnsetFormattedAmount ensures that no value is present for FormattedAmount, not even an explicit nil
### GetSucceeded

`func (o *GETCaptures200ResponseDataInnerAttributes) GetSucceeded() interface{}`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetSucceededOk() (*interface{}, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *GETCaptures200ResponseDataInnerAttributes) SetSucceeded(v interface{})`

SetSucceeded sets Succeeded field to given value.

### HasSucceeded

`func (o *GETCaptures200ResponseDataInnerAttributes) HasSucceeded() bool`

HasSucceeded returns a boolean if a field has been set.

### SetSucceededNil

`func (o *GETCaptures200ResponseDataInnerAttributes) SetSucceededNil(b bool)`

 SetSucceededNil sets the value for Succeeded to be an explicit nil

### UnsetSucceeded
`func (o *GETCaptures200ResponseDataInnerAttributes) UnsetSucceeded()`

UnsetSucceeded ensures that no value is present for Succeeded, not even an explicit nil
### GetMessage

`func (o *GETCaptures200ResponseDataInnerAttributes) GetMessage() interface{}`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetMessageOk() (*interface{}, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GETCaptures200ResponseDataInnerAttributes) SetMessage(v interface{})`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GETCaptures200ResponseDataInnerAttributes) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *GETCaptures200ResponseDataInnerAttributes) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *GETCaptures200ResponseDataInnerAttributes) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetErrorCode

`func (o *GETCaptures200ResponseDataInnerAttributes) GetErrorCode() interface{}`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetErrorCodeOk() (*interface{}, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GETCaptures200ResponseDataInnerAttributes) SetErrorCode(v interface{})`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *GETCaptures200ResponseDataInnerAttributes) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *GETCaptures200ResponseDataInnerAttributes) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *GETCaptures200ResponseDataInnerAttributes) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorDetail

`func (o *GETCaptures200ResponseDataInnerAttributes) GetErrorDetail() interface{}`

GetErrorDetail returns the ErrorDetail field if non-nil, zero value otherwise.

### GetErrorDetailOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetErrorDetailOk() (*interface{}, bool)`

GetErrorDetailOk returns a tuple with the ErrorDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetail

`func (o *GETCaptures200ResponseDataInnerAttributes) SetErrorDetail(v interface{})`

SetErrorDetail sets ErrorDetail field to given value.

### HasErrorDetail

`func (o *GETCaptures200ResponseDataInnerAttributes) HasErrorDetail() bool`

HasErrorDetail returns a boolean if a field has been set.

### SetErrorDetailNil

`func (o *GETCaptures200ResponseDataInnerAttributes) SetErrorDetailNil(b bool)`

 SetErrorDetailNil sets the value for ErrorDetail to be an explicit nil

### UnsetErrorDetail
`func (o *GETCaptures200ResponseDataInnerAttributes) UnsetErrorDetail()`

UnsetErrorDetail ensures that no value is present for ErrorDetail, not even an explicit nil
### GetToken

`func (o *GETCaptures200ResponseDataInnerAttributes) GetToken() interface{}`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetTokenOk() (*interface{}, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GETCaptures200ResponseDataInnerAttributes) SetToken(v interface{})`

SetToken sets Token field to given value.

### HasToken

`func (o *GETCaptures200ResponseDataInnerAttributes) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *GETCaptures200ResponseDataInnerAttributes) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *GETCaptures200ResponseDataInnerAttributes) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetGatewayTransactionId

`func (o *GETCaptures200ResponseDataInnerAttributes) GetGatewayTransactionId() interface{}`

GetGatewayTransactionId returns the GatewayTransactionId field if non-nil, zero value otherwise.

### GetGatewayTransactionIdOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetGatewayTransactionIdOk() (*interface{}, bool)`

GetGatewayTransactionIdOk returns a tuple with the GatewayTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayTransactionId

`func (o *GETCaptures200ResponseDataInnerAttributes) SetGatewayTransactionId(v interface{})`

SetGatewayTransactionId sets GatewayTransactionId field to given value.

### HasGatewayTransactionId

`func (o *GETCaptures200ResponseDataInnerAttributes) HasGatewayTransactionId() bool`

HasGatewayTransactionId returns a boolean if a field has been set.

### SetGatewayTransactionIdNil

`func (o *GETCaptures200ResponseDataInnerAttributes) SetGatewayTransactionIdNil(b bool)`

 SetGatewayTransactionIdNil sets the value for GatewayTransactionId to be an explicit nil

### UnsetGatewayTransactionId
`func (o *GETCaptures200ResponseDataInnerAttributes) UnsetGatewayTransactionId()`

UnsetGatewayTransactionId ensures that no value is present for GatewayTransactionId, not even an explicit nil
### GetCreatedAt

`func (o *GETCaptures200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETCaptures200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETCaptures200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETCaptures200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETCaptures200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETCaptures200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETCaptures200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETCaptures200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETCaptures200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETCaptures200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETCaptures200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETCaptures200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETCaptures200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETCaptures200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETCaptures200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETCaptures200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETCaptures200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETCaptures200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETCaptures200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETCaptures200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETCaptures200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETCaptures200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETCaptures200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETCaptures200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETCaptures200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetRefundAmountCents

`func (o *GETCaptures200ResponseDataInnerAttributes) GetRefundAmountCents() interface{}`

GetRefundAmountCents returns the RefundAmountCents field if non-nil, zero value otherwise.

### GetRefundAmountCentsOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetRefundAmountCentsOk() (*interface{}, bool)`

GetRefundAmountCentsOk returns a tuple with the RefundAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmountCents

`func (o *GETCaptures200ResponseDataInnerAttributes) SetRefundAmountCents(v interface{})`

SetRefundAmountCents sets RefundAmountCents field to given value.

### HasRefundAmountCents

`func (o *GETCaptures200ResponseDataInnerAttributes) HasRefundAmountCents() bool`

HasRefundAmountCents returns a boolean if a field has been set.

### SetRefundAmountCentsNil

`func (o *GETCaptures200ResponseDataInnerAttributes) SetRefundAmountCentsNil(b bool)`

 SetRefundAmountCentsNil sets the value for RefundAmountCents to be an explicit nil

### UnsetRefundAmountCents
`func (o *GETCaptures200ResponseDataInnerAttributes) UnsetRefundAmountCents()`

UnsetRefundAmountCents ensures that no value is present for RefundAmountCents, not even an explicit nil
### GetRefundAmountFloat

`func (o *GETCaptures200ResponseDataInnerAttributes) GetRefundAmountFloat() interface{}`

GetRefundAmountFloat returns the RefundAmountFloat field if non-nil, zero value otherwise.

### GetRefundAmountFloatOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetRefundAmountFloatOk() (*interface{}, bool)`

GetRefundAmountFloatOk returns a tuple with the RefundAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmountFloat

`func (o *GETCaptures200ResponseDataInnerAttributes) SetRefundAmountFloat(v interface{})`

SetRefundAmountFloat sets RefundAmountFloat field to given value.

### HasRefundAmountFloat

`func (o *GETCaptures200ResponseDataInnerAttributes) HasRefundAmountFloat() bool`

HasRefundAmountFloat returns a boolean if a field has been set.

### SetRefundAmountFloatNil

`func (o *GETCaptures200ResponseDataInnerAttributes) SetRefundAmountFloatNil(b bool)`

 SetRefundAmountFloatNil sets the value for RefundAmountFloat to be an explicit nil

### UnsetRefundAmountFloat
`func (o *GETCaptures200ResponseDataInnerAttributes) UnsetRefundAmountFloat()`

UnsetRefundAmountFloat ensures that no value is present for RefundAmountFloat, not even an explicit nil
### GetFormattedRefundAmount

`func (o *GETCaptures200ResponseDataInnerAttributes) GetFormattedRefundAmount() interface{}`

GetFormattedRefundAmount returns the FormattedRefundAmount field if non-nil, zero value otherwise.

### GetFormattedRefundAmountOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetFormattedRefundAmountOk() (*interface{}, bool)`

GetFormattedRefundAmountOk returns a tuple with the FormattedRefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedRefundAmount

`func (o *GETCaptures200ResponseDataInnerAttributes) SetFormattedRefundAmount(v interface{})`

SetFormattedRefundAmount sets FormattedRefundAmount field to given value.

### HasFormattedRefundAmount

`func (o *GETCaptures200ResponseDataInnerAttributes) HasFormattedRefundAmount() bool`

HasFormattedRefundAmount returns a boolean if a field has been set.

### SetFormattedRefundAmountNil

`func (o *GETCaptures200ResponseDataInnerAttributes) SetFormattedRefundAmountNil(b bool)`

 SetFormattedRefundAmountNil sets the value for FormattedRefundAmount to be an explicit nil

### UnsetFormattedRefundAmount
`func (o *GETCaptures200ResponseDataInnerAttributes) UnsetFormattedRefundAmount()`

UnsetFormattedRefundAmount ensures that no value is present for FormattedRefundAmount, not even an explicit nil
### GetRefundBalanceCents

`func (o *GETCaptures200ResponseDataInnerAttributes) GetRefundBalanceCents() interface{}`

GetRefundBalanceCents returns the RefundBalanceCents field if non-nil, zero value otherwise.

### GetRefundBalanceCentsOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetRefundBalanceCentsOk() (*interface{}, bool)`

GetRefundBalanceCentsOk returns a tuple with the RefundBalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundBalanceCents

`func (o *GETCaptures200ResponseDataInnerAttributes) SetRefundBalanceCents(v interface{})`

SetRefundBalanceCents sets RefundBalanceCents field to given value.

### HasRefundBalanceCents

`func (o *GETCaptures200ResponseDataInnerAttributes) HasRefundBalanceCents() bool`

HasRefundBalanceCents returns a boolean if a field has been set.

### SetRefundBalanceCentsNil

`func (o *GETCaptures200ResponseDataInnerAttributes) SetRefundBalanceCentsNil(b bool)`

 SetRefundBalanceCentsNil sets the value for RefundBalanceCents to be an explicit nil

### UnsetRefundBalanceCents
`func (o *GETCaptures200ResponseDataInnerAttributes) UnsetRefundBalanceCents()`

UnsetRefundBalanceCents ensures that no value is present for RefundBalanceCents, not even an explicit nil
### GetRefundBalanceFloat

`func (o *GETCaptures200ResponseDataInnerAttributes) GetRefundBalanceFloat() interface{}`

GetRefundBalanceFloat returns the RefundBalanceFloat field if non-nil, zero value otherwise.

### GetRefundBalanceFloatOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetRefundBalanceFloatOk() (*interface{}, bool)`

GetRefundBalanceFloatOk returns a tuple with the RefundBalanceFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundBalanceFloat

`func (o *GETCaptures200ResponseDataInnerAttributes) SetRefundBalanceFloat(v interface{})`

SetRefundBalanceFloat sets RefundBalanceFloat field to given value.

### HasRefundBalanceFloat

`func (o *GETCaptures200ResponseDataInnerAttributes) HasRefundBalanceFloat() bool`

HasRefundBalanceFloat returns a boolean if a field has been set.

### SetRefundBalanceFloatNil

`func (o *GETCaptures200ResponseDataInnerAttributes) SetRefundBalanceFloatNil(b bool)`

 SetRefundBalanceFloatNil sets the value for RefundBalanceFloat to be an explicit nil

### UnsetRefundBalanceFloat
`func (o *GETCaptures200ResponseDataInnerAttributes) UnsetRefundBalanceFloat()`

UnsetRefundBalanceFloat ensures that no value is present for RefundBalanceFloat, not even an explicit nil
### GetFormattedRefundBalance

`func (o *GETCaptures200ResponseDataInnerAttributes) GetFormattedRefundBalance() interface{}`

GetFormattedRefundBalance returns the FormattedRefundBalance field if non-nil, zero value otherwise.

### GetFormattedRefundBalanceOk

`func (o *GETCaptures200ResponseDataInnerAttributes) GetFormattedRefundBalanceOk() (*interface{}, bool)`

GetFormattedRefundBalanceOk returns a tuple with the FormattedRefundBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedRefundBalance

`func (o *GETCaptures200ResponseDataInnerAttributes) SetFormattedRefundBalance(v interface{})`

SetFormattedRefundBalance sets FormattedRefundBalance field to given value.

### HasFormattedRefundBalance

`func (o *GETCaptures200ResponseDataInnerAttributes) HasFormattedRefundBalance() bool`

HasFormattedRefundBalance returns a boolean if a field has been set.

### SetFormattedRefundBalanceNil

`func (o *GETCaptures200ResponseDataInnerAttributes) SetFormattedRefundBalanceNil(b bool)`

 SetFormattedRefundBalanceNil sets the value for FormattedRefundBalance to be an explicit nil

### UnsetFormattedRefundBalance
`func (o *GETCaptures200ResponseDataInnerAttributes) UnsetFormattedRefundBalance()`

UnsetFormattedRefundBalance ensures that no value is present for FormattedRefundBalance, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


