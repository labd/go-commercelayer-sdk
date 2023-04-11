# GETAuthorizations200ResponseDataInnerAttributes

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
**CvvCode** | Pointer to **interface{}** | The CVV code returned by the payment gateway | [optional] 
**CvvMessage** | Pointer to **interface{}** | The CVV message returned by the payment gateway | [optional] 
**AvsCode** | Pointer to **interface{}** | The AVS code returned by the payment gateway | [optional] 
**AvsMessage** | Pointer to **interface{}** | The AVS message returned by the payment gateway | [optional] 
**FraudReview** | Pointer to **interface{}** | The fraud review message, if any, returned by the payment gateway | [optional] 
**CaptureAmountCents** | Pointer to **interface{}** | The amount to be captured, in cents. | [optional] 
**CaptureAmountFloat** | Pointer to **interface{}** | The amount to be captured, float. | [optional] 
**FormattedCaptureAmount** | Pointer to **interface{}** | The amount to be captured, formatted. | [optional] 
**CaptureBalanceCents** | Pointer to **interface{}** | The balance to be captured, in cents. | [optional] 
**CaptureBalanceFloat** | Pointer to **interface{}** | The balance to be captured, float. | [optional] 
**FormattedCaptureBalance** | Pointer to **interface{}** | The balance to be captured, formatted. | [optional] 
**VoidBalanceCents** | Pointer to **interface{}** | The balance to be voided, in cents. | [optional] 
**VoidBalanceFloat** | Pointer to **interface{}** | The balance to be voided, float. | [optional] 
**FormattedVoidBalance** | Pointer to **interface{}** | The balance to be voided, formatted. | [optional] 

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

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetNumber() interface{}`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetNumberOk() (*interface{}, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetNumber(v interface{})`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetCurrencyCode

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetAmountCents

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetAmountCents() interface{}`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetAmountCentsOk() (*interface{}, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetAmountCents(v interface{})`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### SetAmountCentsNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetAmountCentsNil(b bool)`

 SetAmountCentsNil sets the value for AmountCents to be an explicit nil

### UnsetAmountCents
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetAmountCents()`

UnsetAmountCents ensures that no value is present for AmountCents, not even an explicit nil
### GetAmountFloat

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetAmountFloat() interface{}`

GetAmountFloat returns the AmountFloat field if non-nil, zero value otherwise.

### GetAmountFloatOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetAmountFloatOk() (*interface{}, bool)`

GetAmountFloatOk returns a tuple with the AmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountFloat

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetAmountFloat(v interface{})`

SetAmountFloat sets AmountFloat field to given value.

### HasAmountFloat

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasAmountFloat() bool`

HasAmountFloat returns a boolean if a field has been set.

### SetAmountFloatNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetAmountFloatNil(b bool)`

 SetAmountFloatNil sets the value for AmountFloat to be an explicit nil

### UnsetAmountFloat
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetAmountFloat()`

UnsetAmountFloat ensures that no value is present for AmountFloat, not even an explicit nil
### GetFormattedAmount

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetFormattedAmount() interface{}`

GetFormattedAmount returns the FormattedAmount field if non-nil, zero value otherwise.

### GetFormattedAmountOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetFormattedAmountOk() (*interface{}, bool)`

GetFormattedAmountOk returns a tuple with the FormattedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAmount

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetFormattedAmount(v interface{})`

SetFormattedAmount sets FormattedAmount field to given value.

### HasFormattedAmount

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasFormattedAmount() bool`

HasFormattedAmount returns a boolean if a field has been set.

### SetFormattedAmountNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetFormattedAmountNil(b bool)`

 SetFormattedAmountNil sets the value for FormattedAmount to be an explicit nil

### UnsetFormattedAmount
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetFormattedAmount()`

UnsetFormattedAmount ensures that no value is present for FormattedAmount, not even an explicit nil
### GetSucceeded

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetSucceeded() interface{}`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetSucceededOk() (*interface{}, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetSucceeded(v interface{})`

SetSucceeded sets Succeeded field to given value.

### HasSucceeded

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasSucceeded() bool`

HasSucceeded returns a boolean if a field has been set.

### SetSucceededNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetSucceededNil(b bool)`

 SetSucceededNil sets the value for Succeeded to be an explicit nil

### UnsetSucceeded
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetSucceeded()`

UnsetSucceeded ensures that no value is present for Succeeded, not even an explicit nil
### GetMessage

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetMessage() interface{}`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetMessageOk() (*interface{}, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetMessage(v interface{})`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetErrorCode

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetErrorCode() interface{}`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetErrorCodeOk() (*interface{}, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetErrorCode(v interface{})`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorDetail

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetErrorDetail() interface{}`

GetErrorDetail returns the ErrorDetail field if non-nil, zero value otherwise.

### GetErrorDetailOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetErrorDetailOk() (*interface{}, bool)`

GetErrorDetailOk returns a tuple with the ErrorDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetail

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetErrorDetail(v interface{})`

SetErrorDetail sets ErrorDetail field to given value.

### HasErrorDetail

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasErrorDetail() bool`

HasErrorDetail returns a boolean if a field has been set.

### SetErrorDetailNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetErrorDetailNil(b bool)`

 SetErrorDetailNil sets the value for ErrorDetail to be an explicit nil

### UnsetErrorDetail
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetErrorDetail()`

UnsetErrorDetail ensures that no value is present for ErrorDetail, not even an explicit nil
### GetToken

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetToken() interface{}`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetTokenOk() (*interface{}, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetToken(v interface{})`

SetToken sets Token field to given value.

### HasToken

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetGatewayTransactionId

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetGatewayTransactionId() interface{}`

GetGatewayTransactionId returns the GatewayTransactionId field if non-nil, zero value otherwise.

### GetGatewayTransactionIdOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetGatewayTransactionIdOk() (*interface{}, bool)`

GetGatewayTransactionIdOk returns a tuple with the GatewayTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayTransactionId

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetGatewayTransactionId(v interface{})`

SetGatewayTransactionId sets GatewayTransactionId field to given value.

### HasGatewayTransactionId

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasGatewayTransactionId() bool`

HasGatewayTransactionId returns a boolean if a field has been set.

### SetGatewayTransactionIdNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetGatewayTransactionIdNil(b bool)`

 SetGatewayTransactionIdNil sets the value for GatewayTransactionId to be an explicit nil

### UnsetGatewayTransactionId
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetGatewayTransactionId()`

UnsetGatewayTransactionId ensures that no value is present for GatewayTransactionId, not even an explicit nil
### GetCreatedAt

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCvvCode

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCvvCode() interface{}`

GetCvvCode returns the CvvCode field if non-nil, zero value otherwise.

### GetCvvCodeOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCvvCodeOk() (*interface{}, bool)`

GetCvvCodeOk returns a tuple with the CvvCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvvCode

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetCvvCode(v interface{})`

SetCvvCode sets CvvCode field to given value.

### HasCvvCode

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasCvvCode() bool`

HasCvvCode returns a boolean if a field has been set.

### SetCvvCodeNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetCvvCodeNil(b bool)`

 SetCvvCodeNil sets the value for CvvCode to be an explicit nil

### UnsetCvvCode
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetCvvCode()`

UnsetCvvCode ensures that no value is present for CvvCode, not even an explicit nil
### GetCvvMessage

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCvvMessage() interface{}`

GetCvvMessage returns the CvvMessage field if non-nil, zero value otherwise.

### GetCvvMessageOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCvvMessageOk() (*interface{}, bool)`

GetCvvMessageOk returns a tuple with the CvvMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvvMessage

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetCvvMessage(v interface{})`

SetCvvMessage sets CvvMessage field to given value.

### HasCvvMessage

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasCvvMessage() bool`

HasCvvMessage returns a boolean if a field has been set.

### SetCvvMessageNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetCvvMessageNil(b bool)`

 SetCvvMessageNil sets the value for CvvMessage to be an explicit nil

### UnsetCvvMessage
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetCvvMessage()`

UnsetCvvMessage ensures that no value is present for CvvMessage, not even an explicit nil
### GetAvsCode

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetAvsCode() interface{}`

GetAvsCode returns the AvsCode field if non-nil, zero value otherwise.

### GetAvsCodeOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetAvsCodeOk() (*interface{}, bool)`

GetAvsCodeOk returns a tuple with the AvsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsCode

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetAvsCode(v interface{})`

SetAvsCode sets AvsCode field to given value.

### HasAvsCode

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasAvsCode() bool`

HasAvsCode returns a boolean if a field has been set.

### SetAvsCodeNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetAvsCodeNil(b bool)`

 SetAvsCodeNil sets the value for AvsCode to be an explicit nil

### UnsetAvsCode
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetAvsCode()`

UnsetAvsCode ensures that no value is present for AvsCode, not even an explicit nil
### GetAvsMessage

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetAvsMessage() interface{}`

GetAvsMessage returns the AvsMessage field if non-nil, zero value otherwise.

### GetAvsMessageOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetAvsMessageOk() (*interface{}, bool)`

GetAvsMessageOk returns a tuple with the AvsMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsMessage

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetAvsMessage(v interface{})`

SetAvsMessage sets AvsMessage field to given value.

### HasAvsMessage

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasAvsMessage() bool`

HasAvsMessage returns a boolean if a field has been set.

### SetAvsMessageNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetAvsMessageNil(b bool)`

 SetAvsMessageNil sets the value for AvsMessage to be an explicit nil

### UnsetAvsMessage
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetAvsMessage()`

UnsetAvsMessage ensures that no value is present for AvsMessage, not even an explicit nil
### GetFraudReview

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetFraudReview() interface{}`

GetFraudReview returns the FraudReview field if non-nil, zero value otherwise.

### GetFraudReviewOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetFraudReviewOk() (*interface{}, bool)`

GetFraudReviewOk returns a tuple with the FraudReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudReview

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetFraudReview(v interface{})`

SetFraudReview sets FraudReview field to given value.

### HasFraudReview

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasFraudReview() bool`

HasFraudReview returns a boolean if a field has been set.

### SetFraudReviewNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetFraudReviewNil(b bool)`

 SetFraudReviewNil sets the value for FraudReview to be an explicit nil

### UnsetFraudReview
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetFraudReview()`

UnsetFraudReview ensures that no value is present for FraudReview, not even an explicit nil
### GetCaptureAmountCents

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCaptureAmountCents() interface{}`

GetCaptureAmountCents returns the CaptureAmountCents field if non-nil, zero value otherwise.

### GetCaptureAmountCentsOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCaptureAmountCentsOk() (*interface{}, bool)`

GetCaptureAmountCentsOk returns a tuple with the CaptureAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureAmountCents

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetCaptureAmountCents(v interface{})`

SetCaptureAmountCents sets CaptureAmountCents field to given value.

### HasCaptureAmountCents

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasCaptureAmountCents() bool`

HasCaptureAmountCents returns a boolean if a field has been set.

### SetCaptureAmountCentsNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetCaptureAmountCentsNil(b bool)`

 SetCaptureAmountCentsNil sets the value for CaptureAmountCents to be an explicit nil

### UnsetCaptureAmountCents
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetCaptureAmountCents()`

UnsetCaptureAmountCents ensures that no value is present for CaptureAmountCents, not even an explicit nil
### GetCaptureAmountFloat

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCaptureAmountFloat() interface{}`

GetCaptureAmountFloat returns the CaptureAmountFloat field if non-nil, zero value otherwise.

### GetCaptureAmountFloatOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCaptureAmountFloatOk() (*interface{}, bool)`

GetCaptureAmountFloatOk returns a tuple with the CaptureAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureAmountFloat

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetCaptureAmountFloat(v interface{})`

SetCaptureAmountFloat sets CaptureAmountFloat field to given value.

### HasCaptureAmountFloat

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasCaptureAmountFloat() bool`

HasCaptureAmountFloat returns a boolean if a field has been set.

### SetCaptureAmountFloatNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetCaptureAmountFloatNil(b bool)`

 SetCaptureAmountFloatNil sets the value for CaptureAmountFloat to be an explicit nil

### UnsetCaptureAmountFloat
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetCaptureAmountFloat()`

UnsetCaptureAmountFloat ensures that no value is present for CaptureAmountFloat, not even an explicit nil
### GetFormattedCaptureAmount

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetFormattedCaptureAmount() interface{}`

GetFormattedCaptureAmount returns the FormattedCaptureAmount field if non-nil, zero value otherwise.

### GetFormattedCaptureAmountOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetFormattedCaptureAmountOk() (*interface{}, bool)`

GetFormattedCaptureAmountOk returns a tuple with the FormattedCaptureAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedCaptureAmount

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetFormattedCaptureAmount(v interface{})`

SetFormattedCaptureAmount sets FormattedCaptureAmount field to given value.

### HasFormattedCaptureAmount

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasFormattedCaptureAmount() bool`

HasFormattedCaptureAmount returns a boolean if a field has been set.

### SetFormattedCaptureAmountNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetFormattedCaptureAmountNil(b bool)`

 SetFormattedCaptureAmountNil sets the value for FormattedCaptureAmount to be an explicit nil

### UnsetFormattedCaptureAmount
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetFormattedCaptureAmount()`

UnsetFormattedCaptureAmount ensures that no value is present for FormattedCaptureAmount, not even an explicit nil
### GetCaptureBalanceCents

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCaptureBalanceCents() interface{}`

GetCaptureBalanceCents returns the CaptureBalanceCents field if non-nil, zero value otherwise.

### GetCaptureBalanceCentsOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCaptureBalanceCentsOk() (*interface{}, bool)`

GetCaptureBalanceCentsOk returns a tuple with the CaptureBalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureBalanceCents

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetCaptureBalanceCents(v interface{})`

SetCaptureBalanceCents sets CaptureBalanceCents field to given value.

### HasCaptureBalanceCents

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasCaptureBalanceCents() bool`

HasCaptureBalanceCents returns a boolean if a field has been set.

### SetCaptureBalanceCentsNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetCaptureBalanceCentsNil(b bool)`

 SetCaptureBalanceCentsNil sets the value for CaptureBalanceCents to be an explicit nil

### UnsetCaptureBalanceCents
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetCaptureBalanceCents()`

UnsetCaptureBalanceCents ensures that no value is present for CaptureBalanceCents, not even an explicit nil
### GetCaptureBalanceFloat

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCaptureBalanceFloat() interface{}`

GetCaptureBalanceFloat returns the CaptureBalanceFloat field if non-nil, zero value otherwise.

### GetCaptureBalanceFloatOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetCaptureBalanceFloatOk() (*interface{}, bool)`

GetCaptureBalanceFloatOk returns a tuple with the CaptureBalanceFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureBalanceFloat

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetCaptureBalanceFloat(v interface{})`

SetCaptureBalanceFloat sets CaptureBalanceFloat field to given value.

### HasCaptureBalanceFloat

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasCaptureBalanceFloat() bool`

HasCaptureBalanceFloat returns a boolean if a field has been set.

### SetCaptureBalanceFloatNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetCaptureBalanceFloatNil(b bool)`

 SetCaptureBalanceFloatNil sets the value for CaptureBalanceFloat to be an explicit nil

### UnsetCaptureBalanceFloat
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetCaptureBalanceFloat()`

UnsetCaptureBalanceFloat ensures that no value is present for CaptureBalanceFloat, not even an explicit nil
### GetFormattedCaptureBalance

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetFormattedCaptureBalance() interface{}`

GetFormattedCaptureBalance returns the FormattedCaptureBalance field if non-nil, zero value otherwise.

### GetFormattedCaptureBalanceOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetFormattedCaptureBalanceOk() (*interface{}, bool)`

GetFormattedCaptureBalanceOk returns a tuple with the FormattedCaptureBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedCaptureBalance

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetFormattedCaptureBalance(v interface{})`

SetFormattedCaptureBalance sets FormattedCaptureBalance field to given value.

### HasFormattedCaptureBalance

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasFormattedCaptureBalance() bool`

HasFormattedCaptureBalance returns a boolean if a field has been set.

### SetFormattedCaptureBalanceNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetFormattedCaptureBalanceNil(b bool)`

 SetFormattedCaptureBalanceNil sets the value for FormattedCaptureBalance to be an explicit nil

### UnsetFormattedCaptureBalance
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetFormattedCaptureBalance()`

UnsetFormattedCaptureBalance ensures that no value is present for FormattedCaptureBalance, not even an explicit nil
### GetVoidBalanceCents

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetVoidBalanceCents() interface{}`

GetVoidBalanceCents returns the VoidBalanceCents field if non-nil, zero value otherwise.

### GetVoidBalanceCentsOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetVoidBalanceCentsOk() (*interface{}, bool)`

GetVoidBalanceCentsOk returns a tuple with the VoidBalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidBalanceCents

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetVoidBalanceCents(v interface{})`

SetVoidBalanceCents sets VoidBalanceCents field to given value.

### HasVoidBalanceCents

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasVoidBalanceCents() bool`

HasVoidBalanceCents returns a boolean if a field has been set.

### SetVoidBalanceCentsNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetVoidBalanceCentsNil(b bool)`

 SetVoidBalanceCentsNil sets the value for VoidBalanceCents to be an explicit nil

### UnsetVoidBalanceCents
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetVoidBalanceCents()`

UnsetVoidBalanceCents ensures that no value is present for VoidBalanceCents, not even an explicit nil
### GetVoidBalanceFloat

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetVoidBalanceFloat() interface{}`

GetVoidBalanceFloat returns the VoidBalanceFloat field if non-nil, zero value otherwise.

### GetVoidBalanceFloatOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetVoidBalanceFloatOk() (*interface{}, bool)`

GetVoidBalanceFloatOk returns a tuple with the VoidBalanceFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidBalanceFloat

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetVoidBalanceFloat(v interface{})`

SetVoidBalanceFloat sets VoidBalanceFloat field to given value.

### HasVoidBalanceFloat

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasVoidBalanceFloat() bool`

HasVoidBalanceFloat returns a boolean if a field has been set.

### SetVoidBalanceFloatNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetVoidBalanceFloatNil(b bool)`

 SetVoidBalanceFloatNil sets the value for VoidBalanceFloat to be an explicit nil

### UnsetVoidBalanceFloat
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetVoidBalanceFloat()`

UnsetVoidBalanceFloat ensures that no value is present for VoidBalanceFloat, not even an explicit nil
### GetFormattedVoidBalance

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetFormattedVoidBalance() interface{}`

GetFormattedVoidBalance returns the FormattedVoidBalance field if non-nil, zero value otherwise.

### GetFormattedVoidBalanceOk

`func (o *GETAuthorizations200ResponseDataInnerAttributes) GetFormattedVoidBalanceOk() (*interface{}, bool)`

GetFormattedVoidBalanceOk returns a tuple with the FormattedVoidBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedVoidBalance

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetFormattedVoidBalance(v interface{})`

SetFormattedVoidBalance sets FormattedVoidBalance field to given value.

### HasFormattedVoidBalance

`func (o *GETAuthorizations200ResponseDataInnerAttributes) HasFormattedVoidBalance() bool`

HasFormattedVoidBalance returns a boolean if a field has been set.

### SetFormattedVoidBalanceNil

`func (o *GETAuthorizations200ResponseDataInnerAttributes) SetFormattedVoidBalanceNil(b bool)`

 SetFormattedVoidBalanceNil sets the value for FormattedVoidBalance to be an explicit nil

### UnsetFormattedVoidBalance
`func (o *GETAuthorizations200ResponseDataInnerAttributes) UnsetFormattedVoidBalance()`

UnsetFormattedVoidBalance ensures that no value is present for FormattedVoidBalance, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


