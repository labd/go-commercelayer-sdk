# GETAuthorizations200ResponseDataInnerAttributes

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
**CvvCode** | Pointer to **string** | The CVV code returned by the payment gateway | [optional] 
**CvvMessage** | Pointer to **string** | The CVV message returned by the payment gateway | [optional] 
**AvsCode** | Pointer to **string** | The AVS code returned by the payment gateway | [optional] 
**AvsMessage** | Pointer to **string** | The AVS message returned by the payment gateway | [optional] 
**FraudReview** | Pointer to **string** | The fraud review message, if any, returned by the payment gateway | [optional] 
**CaptureAmountCents** | Pointer to **int32** | The amount to be captured, in cents. | [optional] 
**CaptureAmountFloat** | Pointer to **float32** | The amount to be captured, float. | [optional] 
**FormattedCaptureAmount** | Pointer to **string** | The amount to be captured, formatted. | [optional] 
**CaptureBalanceCents** | Pointer to **int32** | The balance to be captured, in cents. | [optional] 
**CaptureBalanceFloat** | Pointer to **float32** | The balance to be captured, float. | [optional] 
**FormattedCaptureBalance** | Pointer to **string** | The balance to be captured, formatted. | [optional] 
**VoidBalanceCents** | Pointer to **int32** | The balance to be voided, in cents. | [optional] 
**VoidBalanceFloat** | Pointer to **float32** | The balance to be voided, float. | [optional] 
**FormattedVoidBalance** | Pointer to **string** | The balance to be voided, formatted. | [optional] 

## Methods

### NewGETAuthorizations200ResponseDataInnerAttributes

`func NewGETAuthorizations200ResponseDataInnerAttributes() *GETAuthorizations200ResponseDataInnerAttributes`

NewGETAuthorizations200ResponseDataInnerAttributes instantiates a new GETAuthorizations200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETAuthorizations200ResponseDataInnerAttributesWithDefaults

`func NewGETAuthorizations200ResponseDataInnerAttributesWithDefaults() *GETAuthorizations200ResponseDataInnerAttributes`

NewGETAuthorizations200ResponseDataInnerAttributesWithDefaults instantiates a new GETAuthorizations200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetAmountCents

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetAmountFloat

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetAmountFloat() float32`

GetAmountFloat returns the AmountFloat field if non-nil, zero value otherwise.

### GetAmountFloatOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetAmountFloatOk() (*float32, bool)`

GetAmountFloatOk returns a tuple with the AmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountFloat

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetAmountFloat(v float32)`

SetAmountFloat sets AmountFloat field to given value.

### HasAmountFloat

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasAmountFloat() bool`

HasAmountFloat returns a boolean if a field has been set.

### GetFormattedAmount

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetFormattedAmount() string`

GetFormattedAmount returns the FormattedAmount field if non-nil, zero value otherwise.

### GetFormattedAmountOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetFormattedAmountOk() (*string, bool)`

GetFormattedAmountOk returns a tuple with the FormattedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAmount

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetFormattedAmount(v string)`

SetFormattedAmount sets FormattedAmount field to given value.

### HasFormattedAmount

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasFormattedAmount() bool`

HasFormattedAmount returns a boolean if a field has been set.

### GetSucceeded

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetSucceeded() bool`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetSucceededOk() (*bool, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetSucceeded(v bool)`

SetSucceeded sets Succeeded field to given value.

### HasSucceeded

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasSucceeded() bool`

HasSucceeded returns a boolean if a field has been set.

### GetMessage

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrorCode

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDetail

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetErrorDetail() string`

GetErrorDetail returns the ErrorDetail field if non-nil, zero value otherwise.

### GetErrorDetailOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetErrorDetailOk() (*string, bool)`

GetErrorDetailOk returns a tuple with the ErrorDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetail

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetErrorDetail(v string)`

SetErrorDetail sets ErrorDetail field to given value.

### HasErrorDetail

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasErrorDetail() bool`

HasErrorDetail returns a boolean if a field has been set.

### GetToken

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetGatewayTransactionId

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetGatewayTransactionId() string`

GetGatewayTransactionId returns the GatewayTransactionId field if non-nil, zero value otherwise.

### GetGatewayTransactionIdOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetGatewayTransactionIdOk() (*string, bool)`

GetGatewayTransactionIdOk returns a tuple with the GatewayTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayTransactionId

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetGatewayTransactionId(v string)`

SetGatewayTransactionId sets GatewayTransactionId field to given value.

### HasGatewayTransactionId

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasGatewayTransactionId() bool`

HasGatewayTransactionId returns a boolean if a field has been set.

### GetId

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCvvCode

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCvvCode() string`

GetCvvCode returns the CvvCode field if non-nil, zero value otherwise.

### GetCvvCodeOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCvvCodeOk() (*string, bool)`

GetCvvCodeOk returns a tuple with the CvvCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvvCode

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetCvvCode(v string)`

SetCvvCode sets CvvCode field to given value.

### HasCvvCode

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasCvvCode() bool`

HasCvvCode returns a boolean if a field has been set.

### GetCvvMessage

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCvvMessage() string`

GetCvvMessage returns the CvvMessage field if non-nil, zero value otherwise.

### GetCvvMessageOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCvvMessageOk() (*string, bool)`

GetCvvMessageOk returns a tuple with the CvvMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvvMessage

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetCvvMessage(v string)`

SetCvvMessage sets CvvMessage field to given value.

### HasCvvMessage

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasCvvMessage() bool`

HasCvvMessage returns a boolean if a field has been set.

### GetAvsCode

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetAvsCode() string`

GetAvsCode returns the AvsCode field if non-nil, zero value otherwise.

### GetAvsCodeOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetAvsCodeOk() (*string, bool)`

GetAvsCodeOk returns a tuple with the AvsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsCode

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetAvsCode(v string)`

SetAvsCode sets AvsCode field to given value.

### HasAvsCode

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasAvsCode() bool`

HasAvsCode returns a boolean if a field has been set.

### GetAvsMessage

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetAvsMessage() string`

GetAvsMessage returns the AvsMessage field if non-nil, zero value otherwise.

### GetAvsMessageOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetAvsMessageOk() (*string, bool)`

GetAvsMessageOk returns a tuple with the AvsMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsMessage

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetAvsMessage(v string)`

SetAvsMessage sets AvsMessage field to given value.

### HasAvsMessage

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasAvsMessage() bool`

HasAvsMessage returns a boolean if a field has been set.

### GetFraudReview

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetFraudReview() string`

GetFraudReview returns the FraudReview field if non-nil, zero value otherwise.

### GetFraudReviewOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetFraudReviewOk() (*string, bool)`

GetFraudReviewOk returns a tuple with the FraudReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudReview

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetFraudReview(v string)`

SetFraudReview sets FraudReview field to given value.

### HasFraudReview

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasFraudReview() bool`

HasFraudReview returns a boolean if a field has been set.

### GetCaptureAmountCents

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCaptureAmountCents() int32`

GetCaptureAmountCents returns the CaptureAmountCents field if non-nil, zero value otherwise.

### GetCaptureAmountCentsOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCaptureAmountCentsOk() (*int32, bool)`

GetCaptureAmountCentsOk returns a tuple with the CaptureAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureAmountCents

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetCaptureAmountCents(v int32)`

SetCaptureAmountCents sets CaptureAmountCents field to given value.

### HasCaptureAmountCents

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasCaptureAmountCents() bool`

HasCaptureAmountCents returns a boolean if a field has been set.

### GetCaptureAmountFloat

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCaptureAmountFloat() float32`

GetCaptureAmountFloat returns the CaptureAmountFloat field if non-nil, zero value otherwise.

### GetCaptureAmountFloatOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCaptureAmountFloatOk() (*float32, bool)`

GetCaptureAmountFloatOk returns a tuple with the CaptureAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureAmountFloat

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetCaptureAmountFloat(v float32)`

SetCaptureAmountFloat sets CaptureAmountFloat field to given value.

### HasCaptureAmountFloat

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasCaptureAmountFloat() bool`

HasCaptureAmountFloat returns a boolean if a field has been set.

### GetFormattedCaptureAmount

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetFormattedCaptureAmount() string`

GetFormattedCaptureAmount returns the FormattedCaptureAmount field if non-nil, zero value otherwise.

### GetFormattedCaptureAmountOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetFormattedCaptureAmountOk() (*string, bool)`

GetFormattedCaptureAmountOk returns a tuple with the FormattedCaptureAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedCaptureAmount

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetFormattedCaptureAmount(v string)`

SetFormattedCaptureAmount sets FormattedCaptureAmount field to given value.

### HasFormattedCaptureAmount

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasFormattedCaptureAmount() bool`

HasFormattedCaptureAmount returns a boolean if a field has been set.

### GetCaptureBalanceCents

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCaptureBalanceCents() int32`

GetCaptureBalanceCents returns the CaptureBalanceCents field if non-nil, zero value otherwise.

### GetCaptureBalanceCentsOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCaptureBalanceCentsOk() (*int32, bool)`

GetCaptureBalanceCentsOk returns a tuple with the CaptureBalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureBalanceCents

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetCaptureBalanceCents(v int32)`

SetCaptureBalanceCents sets CaptureBalanceCents field to given value.

### HasCaptureBalanceCents

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasCaptureBalanceCents() bool`

HasCaptureBalanceCents returns a boolean if a field has been set.

### GetCaptureBalanceFloat

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCaptureBalanceFloat() float32`

GetCaptureBalanceFloat returns the CaptureBalanceFloat field if non-nil, zero value otherwise.

### GetCaptureBalanceFloatOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCaptureBalanceFloatOk() (*float32, bool)`

GetCaptureBalanceFloatOk returns a tuple with the CaptureBalanceFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureBalanceFloat

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetCaptureBalanceFloat(v float32)`

SetCaptureBalanceFloat sets CaptureBalanceFloat field to given value.

### HasCaptureBalanceFloat

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasCaptureBalanceFloat() bool`

HasCaptureBalanceFloat returns a boolean if a field has been set.

### GetFormattedCaptureBalance

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetFormattedCaptureBalance() string`

GetFormattedCaptureBalance returns the FormattedCaptureBalance field if non-nil, zero value otherwise.

### GetFormattedCaptureBalanceOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetFormattedCaptureBalanceOk() (*string, bool)`

GetFormattedCaptureBalanceOk returns a tuple with the FormattedCaptureBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedCaptureBalance

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetFormattedCaptureBalance(v string)`

SetFormattedCaptureBalance sets FormattedCaptureBalance field to given value.

### HasFormattedCaptureBalance

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasFormattedCaptureBalance() bool`

HasFormattedCaptureBalance returns a boolean if a field has been set.

### GetVoidBalanceCents

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetVoidBalanceCents() int32`

GetVoidBalanceCents returns the VoidBalanceCents field if non-nil, zero value otherwise.

### GetVoidBalanceCentsOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetVoidBalanceCentsOk() (*int32, bool)`

GetVoidBalanceCentsOk returns a tuple with the VoidBalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidBalanceCents

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetVoidBalanceCents(v int32)`

SetVoidBalanceCents sets VoidBalanceCents field to given value.

### HasVoidBalanceCents

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasVoidBalanceCents() bool`

HasVoidBalanceCents returns a boolean if a field has been set.

### GetVoidBalanceFloat

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetVoidBalanceFloat() float32`

GetVoidBalanceFloat returns the VoidBalanceFloat field if non-nil, zero value otherwise.

### GetVoidBalanceFloatOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetVoidBalanceFloatOk() (*float32, bool)`

GetVoidBalanceFloatOk returns a tuple with the VoidBalanceFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidBalanceFloat

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetVoidBalanceFloat(v float32)`

SetVoidBalanceFloat sets VoidBalanceFloat field to given value.

### HasVoidBalanceFloat

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasVoidBalanceFloat() bool`

HasVoidBalanceFloat returns a boolean if a field has been set.

### GetFormattedVoidBalance

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetFormattedVoidBalance() string`

GetFormattedVoidBalance returns the FormattedVoidBalance field if non-nil, zero value otherwise.

### GetFormattedVoidBalanceOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetFormattedVoidBalanceOk() (*string, bool)`

GetFormattedVoidBalanceOk returns a tuple with the FormattedVoidBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedVoidBalance

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetFormattedVoidBalance(v string)`

SetFormattedVoidBalance sets FormattedVoidBalance field to given value.

### HasFormattedVoidBalance

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasFormattedVoidBalance() bool`

HasFormattedVoidBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


