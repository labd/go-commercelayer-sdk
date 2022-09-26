# AuthorizationDataAttributes

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

### NewAuthorizationDataAttributes

`func NewAuthorizationDataAttributes() *AuthorizationDataAttributes`

NewAuthorizationDataAttributes instantiates a new AuthorizationDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationDataAttributesWithDefaults

`func NewAuthorizationDataAttributesWithDefaults() *AuthorizationDataAttributes`

NewAuthorizationDataAttributesWithDefaults instantiates a new AuthorizationDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *AuthorizationDataAttributes) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *AuthorizationDataAttributes) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *AuthorizationDataAttributes) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *AuthorizationDataAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *AuthorizationDataAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AuthorizationDataAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AuthorizationDataAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *AuthorizationDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetAmountCents

`func (o *AuthorizationDataAttributes) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *AuthorizationDataAttributes) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *AuthorizationDataAttributes) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *AuthorizationDataAttributes) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetAmountFloat

`func (o *AuthorizationDataAttributes) GetAmountFloat() float32`

GetAmountFloat returns the AmountFloat field if non-nil, zero value otherwise.

### GetAmountFloatOk

`func (o *AuthorizationDataAttributes) GetAmountFloatOk() (*float32, bool)`

GetAmountFloatOk returns a tuple with the AmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountFloat

`func (o *AuthorizationDataAttributes) SetAmountFloat(v float32)`

SetAmountFloat sets AmountFloat field to given value.

### HasAmountFloat

`func (o *AuthorizationDataAttributes) HasAmountFloat() bool`

HasAmountFloat returns a boolean if a field has been set.

### GetFormattedAmount

`func (o *AuthorizationDataAttributes) GetFormattedAmount() string`

GetFormattedAmount returns the FormattedAmount field if non-nil, zero value otherwise.

### GetFormattedAmountOk

`func (o *AuthorizationDataAttributes) GetFormattedAmountOk() (*string, bool)`

GetFormattedAmountOk returns a tuple with the FormattedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAmount

`func (o *AuthorizationDataAttributes) SetFormattedAmount(v string)`

SetFormattedAmount sets FormattedAmount field to given value.

### HasFormattedAmount

`func (o *AuthorizationDataAttributes) HasFormattedAmount() bool`

HasFormattedAmount returns a boolean if a field has been set.

### GetSucceeded

`func (o *AuthorizationDataAttributes) GetSucceeded() bool`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *AuthorizationDataAttributes) GetSucceededOk() (*bool, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *AuthorizationDataAttributes) SetSucceeded(v bool)`

SetSucceeded sets Succeeded field to given value.

### HasSucceeded

`func (o *AuthorizationDataAttributes) HasSucceeded() bool`

HasSucceeded returns a boolean if a field has been set.

### GetMessage

`func (o *AuthorizationDataAttributes) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AuthorizationDataAttributes) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AuthorizationDataAttributes) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AuthorizationDataAttributes) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrorCode

`func (o *AuthorizationDataAttributes) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *AuthorizationDataAttributes) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *AuthorizationDataAttributes) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *AuthorizationDataAttributes) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDetail

`func (o *AuthorizationDataAttributes) GetErrorDetail() string`

GetErrorDetail returns the ErrorDetail field if non-nil, zero value otherwise.

### GetErrorDetailOk

`func (o *AuthorizationDataAttributes) GetErrorDetailOk() (*string, bool)`

GetErrorDetailOk returns a tuple with the ErrorDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetail

`func (o *AuthorizationDataAttributes) SetErrorDetail(v string)`

SetErrorDetail sets ErrorDetail field to given value.

### HasErrorDetail

`func (o *AuthorizationDataAttributes) HasErrorDetail() bool`

HasErrorDetail returns a boolean if a field has been set.

### GetToken

`func (o *AuthorizationDataAttributes) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthorizationDataAttributes) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthorizationDataAttributes) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthorizationDataAttributes) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetGatewayTransactionId

`func (o *AuthorizationDataAttributes) GetGatewayTransactionId() string`

GetGatewayTransactionId returns the GatewayTransactionId field if non-nil, zero value otherwise.

### GetGatewayTransactionIdOk

`func (o *AuthorizationDataAttributes) GetGatewayTransactionIdOk() (*string, bool)`

GetGatewayTransactionIdOk returns a tuple with the GatewayTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayTransactionId

`func (o *AuthorizationDataAttributes) SetGatewayTransactionId(v string)`

SetGatewayTransactionId sets GatewayTransactionId field to given value.

### HasGatewayTransactionId

`func (o *AuthorizationDataAttributes) HasGatewayTransactionId() bool`

HasGatewayTransactionId returns a boolean if a field has been set.

### GetId

`func (o *AuthorizationDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorizationDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorizationDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthorizationDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuthorizationDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuthorizationDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuthorizationDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuthorizationDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuthorizationDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuthorizationDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuthorizationDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuthorizationDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *AuthorizationDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *AuthorizationDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *AuthorizationDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *AuthorizationDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *AuthorizationDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *AuthorizationDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *AuthorizationDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *AuthorizationDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *AuthorizationDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AuthorizationDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AuthorizationDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AuthorizationDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCvvCode

`func (o *AuthorizationDataAttributes) GetCvvCode() string`

GetCvvCode returns the CvvCode field if non-nil, zero value otherwise.

### GetCvvCodeOk

`func (o *AuthorizationDataAttributes) GetCvvCodeOk() (*string, bool)`

GetCvvCodeOk returns a tuple with the CvvCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvvCode

`func (o *AuthorizationDataAttributes) SetCvvCode(v string)`

SetCvvCode sets CvvCode field to given value.

### HasCvvCode

`func (o *AuthorizationDataAttributes) HasCvvCode() bool`

HasCvvCode returns a boolean if a field has been set.

### GetCvvMessage

`func (o *AuthorizationDataAttributes) GetCvvMessage() string`

GetCvvMessage returns the CvvMessage field if non-nil, zero value otherwise.

### GetCvvMessageOk

`func (o *AuthorizationDataAttributes) GetCvvMessageOk() (*string, bool)`

GetCvvMessageOk returns a tuple with the CvvMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvvMessage

`func (o *AuthorizationDataAttributes) SetCvvMessage(v string)`

SetCvvMessage sets CvvMessage field to given value.

### HasCvvMessage

`func (o *AuthorizationDataAttributes) HasCvvMessage() bool`

HasCvvMessage returns a boolean if a field has been set.

### GetAvsCode

`func (o *AuthorizationDataAttributes) GetAvsCode() string`

GetAvsCode returns the AvsCode field if non-nil, zero value otherwise.

### GetAvsCodeOk

`func (o *AuthorizationDataAttributes) GetAvsCodeOk() (*string, bool)`

GetAvsCodeOk returns a tuple with the AvsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsCode

`func (o *AuthorizationDataAttributes) SetAvsCode(v string)`

SetAvsCode sets AvsCode field to given value.

### HasAvsCode

`func (o *AuthorizationDataAttributes) HasAvsCode() bool`

HasAvsCode returns a boolean if a field has been set.

### GetAvsMessage

`func (o *AuthorizationDataAttributes) GetAvsMessage() string`

GetAvsMessage returns the AvsMessage field if non-nil, zero value otherwise.

### GetAvsMessageOk

`func (o *AuthorizationDataAttributes) GetAvsMessageOk() (*string, bool)`

GetAvsMessageOk returns a tuple with the AvsMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsMessage

`func (o *AuthorizationDataAttributes) SetAvsMessage(v string)`

SetAvsMessage sets AvsMessage field to given value.

### HasAvsMessage

`func (o *AuthorizationDataAttributes) HasAvsMessage() bool`

HasAvsMessage returns a boolean if a field has been set.

### GetFraudReview

`func (o *AuthorizationDataAttributes) GetFraudReview() string`

GetFraudReview returns the FraudReview field if non-nil, zero value otherwise.

### GetFraudReviewOk

`func (o *AuthorizationDataAttributes) GetFraudReviewOk() (*string, bool)`

GetFraudReviewOk returns a tuple with the FraudReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudReview

`func (o *AuthorizationDataAttributes) SetFraudReview(v string)`

SetFraudReview sets FraudReview field to given value.

### HasFraudReview

`func (o *AuthorizationDataAttributes) HasFraudReview() bool`

HasFraudReview returns a boolean if a field has been set.

### GetCaptureAmountCents

`func (o *AuthorizationDataAttributes) GetCaptureAmountCents() int32`

GetCaptureAmountCents returns the CaptureAmountCents field if non-nil, zero value otherwise.

### GetCaptureAmountCentsOk

`func (o *AuthorizationDataAttributes) GetCaptureAmountCentsOk() (*int32, bool)`

GetCaptureAmountCentsOk returns a tuple with the CaptureAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureAmountCents

`func (o *AuthorizationDataAttributes) SetCaptureAmountCents(v int32)`

SetCaptureAmountCents sets CaptureAmountCents field to given value.

### HasCaptureAmountCents

`func (o *AuthorizationDataAttributes) HasCaptureAmountCents() bool`

HasCaptureAmountCents returns a boolean if a field has been set.

### GetCaptureAmountFloat

`func (o *AuthorizationDataAttributes) GetCaptureAmountFloat() float32`

GetCaptureAmountFloat returns the CaptureAmountFloat field if non-nil, zero value otherwise.

### GetCaptureAmountFloatOk

`func (o *AuthorizationDataAttributes) GetCaptureAmountFloatOk() (*float32, bool)`

GetCaptureAmountFloatOk returns a tuple with the CaptureAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureAmountFloat

`func (o *AuthorizationDataAttributes) SetCaptureAmountFloat(v float32)`

SetCaptureAmountFloat sets CaptureAmountFloat field to given value.

### HasCaptureAmountFloat

`func (o *AuthorizationDataAttributes) HasCaptureAmountFloat() bool`

HasCaptureAmountFloat returns a boolean if a field has been set.

### GetFormattedCaptureAmount

`func (o *AuthorizationDataAttributes) GetFormattedCaptureAmount() string`

GetFormattedCaptureAmount returns the FormattedCaptureAmount field if non-nil, zero value otherwise.

### GetFormattedCaptureAmountOk

`func (o *AuthorizationDataAttributes) GetFormattedCaptureAmountOk() (*string, bool)`

GetFormattedCaptureAmountOk returns a tuple with the FormattedCaptureAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedCaptureAmount

`func (o *AuthorizationDataAttributes) SetFormattedCaptureAmount(v string)`

SetFormattedCaptureAmount sets FormattedCaptureAmount field to given value.

### HasFormattedCaptureAmount

`func (o *AuthorizationDataAttributes) HasFormattedCaptureAmount() bool`

HasFormattedCaptureAmount returns a boolean if a field has been set.

### GetCaptureBalanceCents

`func (o *AuthorizationDataAttributes) GetCaptureBalanceCents() int32`

GetCaptureBalanceCents returns the CaptureBalanceCents field if non-nil, zero value otherwise.

### GetCaptureBalanceCentsOk

`func (o *AuthorizationDataAttributes) GetCaptureBalanceCentsOk() (*int32, bool)`

GetCaptureBalanceCentsOk returns a tuple with the CaptureBalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureBalanceCents

`func (o *AuthorizationDataAttributes) SetCaptureBalanceCents(v int32)`

SetCaptureBalanceCents sets CaptureBalanceCents field to given value.

### HasCaptureBalanceCents

`func (o *AuthorizationDataAttributes) HasCaptureBalanceCents() bool`

HasCaptureBalanceCents returns a boolean if a field has been set.

### GetCaptureBalanceFloat

`func (o *AuthorizationDataAttributes) GetCaptureBalanceFloat() float32`

GetCaptureBalanceFloat returns the CaptureBalanceFloat field if non-nil, zero value otherwise.

### GetCaptureBalanceFloatOk

`func (o *AuthorizationDataAttributes) GetCaptureBalanceFloatOk() (*float32, bool)`

GetCaptureBalanceFloatOk returns a tuple with the CaptureBalanceFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureBalanceFloat

`func (o *AuthorizationDataAttributes) SetCaptureBalanceFloat(v float32)`

SetCaptureBalanceFloat sets CaptureBalanceFloat field to given value.

### HasCaptureBalanceFloat

`func (o *AuthorizationDataAttributes) HasCaptureBalanceFloat() bool`

HasCaptureBalanceFloat returns a boolean if a field has been set.

### GetFormattedCaptureBalance

`func (o *AuthorizationDataAttributes) GetFormattedCaptureBalance() string`

GetFormattedCaptureBalance returns the FormattedCaptureBalance field if non-nil, zero value otherwise.

### GetFormattedCaptureBalanceOk

`func (o *AuthorizationDataAttributes) GetFormattedCaptureBalanceOk() (*string, bool)`

GetFormattedCaptureBalanceOk returns a tuple with the FormattedCaptureBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedCaptureBalance

`func (o *AuthorizationDataAttributes) SetFormattedCaptureBalance(v string)`

SetFormattedCaptureBalance sets FormattedCaptureBalance field to given value.

### HasFormattedCaptureBalance

`func (o *AuthorizationDataAttributes) HasFormattedCaptureBalance() bool`

HasFormattedCaptureBalance returns a boolean if a field has been set.

### GetVoidBalanceCents

`func (o *AuthorizationDataAttributes) GetVoidBalanceCents() int32`

GetVoidBalanceCents returns the VoidBalanceCents field if non-nil, zero value otherwise.

### GetVoidBalanceCentsOk

`func (o *AuthorizationDataAttributes) GetVoidBalanceCentsOk() (*int32, bool)`

GetVoidBalanceCentsOk returns a tuple with the VoidBalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidBalanceCents

`func (o *AuthorizationDataAttributes) SetVoidBalanceCents(v int32)`

SetVoidBalanceCents sets VoidBalanceCents field to given value.

### HasVoidBalanceCents

`func (o *AuthorizationDataAttributes) HasVoidBalanceCents() bool`

HasVoidBalanceCents returns a boolean if a field has been set.

### GetVoidBalanceFloat

`func (o *AuthorizationDataAttributes) GetVoidBalanceFloat() float32`

GetVoidBalanceFloat returns the VoidBalanceFloat field if non-nil, zero value otherwise.

### GetVoidBalanceFloatOk

`func (o *AuthorizationDataAttributes) GetVoidBalanceFloatOk() (*float32, bool)`

GetVoidBalanceFloatOk returns a tuple with the VoidBalanceFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidBalanceFloat

`func (o *AuthorizationDataAttributes) SetVoidBalanceFloat(v float32)`

SetVoidBalanceFloat sets VoidBalanceFloat field to given value.

### HasVoidBalanceFloat

`func (o *AuthorizationDataAttributes) HasVoidBalanceFloat() bool`

HasVoidBalanceFloat returns a boolean if a field has been set.

### GetFormattedVoidBalance

`func (o *AuthorizationDataAttributes) GetFormattedVoidBalance() string`

GetFormattedVoidBalance returns the FormattedVoidBalance field if non-nil, zero value otherwise.

### GetFormattedVoidBalanceOk

`func (o *AuthorizationDataAttributes) GetFormattedVoidBalanceOk() (*string, bool)`

GetFormattedVoidBalanceOk returns a tuple with the FormattedVoidBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedVoidBalance

`func (o *AuthorizationDataAttributes) SetFormattedVoidBalance(v string)`

SetFormattedVoidBalance sets FormattedVoidBalance field to given value.

### HasFormattedVoidBalance

`func (o *AuthorizationDataAttributes) HasFormattedVoidBalance() bool`

HasFormattedVoidBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


