# GETAuthorizationsAuthorizationId200ResponseDataAttributes

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

### NewGETAuthorizationsAuthorizationId200ResponseDataAttributes

`func NewGETAuthorizationsAuthorizationId200ResponseDataAttributes() *GETAuthorizationsAuthorizationId200ResponseDataAttributes`

NewGETAuthorizationsAuthorizationId200ResponseDataAttributes instantiates a new GETAuthorizationsAuthorizationId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETAuthorizationsAuthorizationId200ResponseDataAttributesWithDefaults

`func NewGETAuthorizationsAuthorizationId200ResponseDataAttributesWithDefaults() *GETAuthorizationsAuthorizationId200ResponseDataAttributes`

NewGETAuthorizationsAuthorizationId200ResponseDataAttributesWithDefaults instantiates a new GETAuthorizationsAuthorizationId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetNumber() interface{}`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetNumberOk() (*interface{}, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetNumber(v interface{})`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetCurrencyCode

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetAmountCents

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetAmountCents() interface{}`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetAmountCentsOk() (*interface{}, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetAmountCents(v interface{})`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### SetAmountCentsNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetAmountCentsNil(b bool)`

 SetAmountCentsNil sets the value for AmountCents to be an explicit nil

### UnsetAmountCents
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetAmountCents()`

UnsetAmountCents ensures that no value is present for AmountCents, not even an explicit nil
### GetAmountFloat

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetAmountFloat() interface{}`

GetAmountFloat returns the AmountFloat field if non-nil, zero value otherwise.

### GetAmountFloatOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetAmountFloatOk() (*interface{}, bool)`

GetAmountFloatOk returns a tuple with the AmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountFloat

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetAmountFloat(v interface{})`

SetAmountFloat sets AmountFloat field to given value.

### HasAmountFloat

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasAmountFloat() bool`

HasAmountFloat returns a boolean if a field has been set.

### SetAmountFloatNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetAmountFloatNil(b bool)`

 SetAmountFloatNil sets the value for AmountFloat to be an explicit nil

### UnsetAmountFloat
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetAmountFloat()`

UnsetAmountFloat ensures that no value is present for AmountFloat, not even an explicit nil
### GetFormattedAmount

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetFormattedAmount() interface{}`

GetFormattedAmount returns the FormattedAmount field if non-nil, zero value otherwise.

### GetFormattedAmountOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetFormattedAmountOk() (*interface{}, bool)`

GetFormattedAmountOk returns a tuple with the FormattedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAmount

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetFormattedAmount(v interface{})`

SetFormattedAmount sets FormattedAmount field to given value.

### HasFormattedAmount

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasFormattedAmount() bool`

HasFormattedAmount returns a boolean if a field has been set.

### SetFormattedAmountNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetFormattedAmountNil(b bool)`

 SetFormattedAmountNil sets the value for FormattedAmount to be an explicit nil

### UnsetFormattedAmount
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetFormattedAmount()`

UnsetFormattedAmount ensures that no value is present for FormattedAmount, not even an explicit nil
### GetSucceeded

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetSucceeded() interface{}`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetSucceededOk() (*interface{}, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetSucceeded(v interface{})`

SetSucceeded sets Succeeded field to given value.

### HasSucceeded

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasSucceeded() bool`

HasSucceeded returns a boolean if a field has been set.

### SetSucceededNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetSucceededNil(b bool)`

 SetSucceededNil sets the value for Succeeded to be an explicit nil

### UnsetSucceeded
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetSucceeded()`

UnsetSucceeded ensures that no value is present for Succeeded, not even an explicit nil
### GetMessage

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetMessage() interface{}`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetMessageOk() (*interface{}, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetMessage(v interface{})`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetErrorCode

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetErrorCode() interface{}`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetErrorCodeOk() (*interface{}, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetErrorCode(v interface{})`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorDetail

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetErrorDetail() interface{}`

GetErrorDetail returns the ErrorDetail field if non-nil, zero value otherwise.

### GetErrorDetailOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetErrorDetailOk() (*interface{}, bool)`

GetErrorDetailOk returns a tuple with the ErrorDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetail

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetErrorDetail(v interface{})`

SetErrorDetail sets ErrorDetail field to given value.

### HasErrorDetail

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasErrorDetail() bool`

HasErrorDetail returns a boolean if a field has been set.

### SetErrorDetailNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetErrorDetailNil(b bool)`

 SetErrorDetailNil sets the value for ErrorDetail to be an explicit nil

### UnsetErrorDetail
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetErrorDetail()`

UnsetErrorDetail ensures that no value is present for ErrorDetail, not even an explicit nil
### GetToken

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetToken() interface{}`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetTokenOk() (*interface{}, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetToken(v interface{})`

SetToken sets Token field to given value.

### HasToken

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetGatewayTransactionId

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetGatewayTransactionId() interface{}`

GetGatewayTransactionId returns the GatewayTransactionId field if non-nil, zero value otherwise.

### GetGatewayTransactionIdOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetGatewayTransactionIdOk() (*interface{}, bool)`

GetGatewayTransactionIdOk returns a tuple with the GatewayTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayTransactionId

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetGatewayTransactionId(v interface{})`

SetGatewayTransactionId sets GatewayTransactionId field to given value.

### HasGatewayTransactionId

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasGatewayTransactionId() bool`

HasGatewayTransactionId returns a boolean if a field has been set.

### SetGatewayTransactionIdNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetGatewayTransactionIdNil(b bool)`

 SetGatewayTransactionIdNil sets the value for GatewayTransactionId to be an explicit nil

### UnsetGatewayTransactionId
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetGatewayTransactionId()`

UnsetGatewayTransactionId ensures that no value is present for GatewayTransactionId, not even an explicit nil
### GetCreatedAt

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCvvCode

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetCvvCode() interface{}`

GetCvvCode returns the CvvCode field if non-nil, zero value otherwise.

### GetCvvCodeOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetCvvCodeOk() (*interface{}, bool)`

GetCvvCodeOk returns a tuple with the CvvCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvvCode

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetCvvCode(v interface{})`

SetCvvCode sets CvvCode field to given value.

### HasCvvCode

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasCvvCode() bool`

HasCvvCode returns a boolean if a field has been set.

### SetCvvCodeNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetCvvCodeNil(b bool)`

 SetCvvCodeNil sets the value for CvvCode to be an explicit nil

### UnsetCvvCode
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetCvvCode()`

UnsetCvvCode ensures that no value is present for CvvCode, not even an explicit nil
### GetCvvMessage

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetCvvMessage() interface{}`

GetCvvMessage returns the CvvMessage field if non-nil, zero value otherwise.

### GetCvvMessageOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetCvvMessageOk() (*interface{}, bool)`

GetCvvMessageOk returns a tuple with the CvvMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvvMessage

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetCvvMessage(v interface{})`

SetCvvMessage sets CvvMessage field to given value.

### HasCvvMessage

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasCvvMessage() bool`

HasCvvMessage returns a boolean if a field has been set.

### SetCvvMessageNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetCvvMessageNil(b bool)`

 SetCvvMessageNil sets the value for CvvMessage to be an explicit nil

### UnsetCvvMessage
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetCvvMessage()`

UnsetCvvMessage ensures that no value is present for CvvMessage, not even an explicit nil
### GetAvsCode

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetAvsCode() interface{}`

GetAvsCode returns the AvsCode field if non-nil, zero value otherwise.

### GetAvsCodeOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetAvsCodeOk() (*interface{}, bool)`

GetAvsCodeOk returns a tuple with the AvsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsCode

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetAvsCode(v interface{})`

SetAvsCode sets AvsCode field to given value.

### HasAvsCode

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasAvsCode() bool`

HasAvsCode returns a boolean if a field has been set.

### SetAvsCodeNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetAvsCodeNil(b bool)`

 SetAvsCodeNil sets the value for AvsCode to be an explicit nil

### UnsetAvsCode
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetAvsCode()`

UnsetAvsCode ensures that no value is present for AvsCode, not even an explicit nil
### GetAvsMessage

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetAvsMessage() interface{}`

GetAvsMessage returns the AvsMessage field if non-nil, zero value otherwise.

### GetAvsMessageOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetAvsMessageOk() (*interface{}, bool)`

GetAvsMessageOk returns a tuple with the AvsMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsMessage

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetAvsMessage(v interface{})`

SetAvsMessage sets AvsMessage field to given value.

### HasAvsMessage

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasAvsMessage() bool`

HasAvsMessage returns a boolean if a field has been set.

### SetAvsMessageNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetAvsMessageNil(b bool)`

 SetAvsMessageNil sets the value for AvsMessage to be an explicit nil

### UnsetAvsMessage
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetAvsMessage()`

UnsetAvsMessage ensures that no value is present for AvsMessage, not even an explicit nil
### GetFraudReview

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetFraudReview() interface{}`

GetFraudReview returns the FraudReview field if non-nil, zero value otherwise.

### GetFraudReviewOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetFraudReviewOk() (*interface{}, bool)`

GetFraudReviewOk returns a tuple with the FraudReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudReview

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetFraudReview(v interface{})`

SetFraudReview sets FraudReview field to given value.

### HasFraudReview

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasFraudReview() bool`

HasFraudReview returns a boolean if a field has been set.

### SetFraudReviewNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetFraudReviewNil(b bool)`

 SetFraudReviewNil sets the value for FraudReview to be an explicit nil

### UnsetFraudReview
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetFraudReview()`

UnsetFraudReview ensures that no value is present for FraudReview, not even an explicit nil
### GetCaptureAmountCents

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetCaptureAmountCents() interface{}`

GetCaptureAmountCents returns the CaptureAmountCents field if non-nil, zero value otherwise.

### GetCaptureAmountCentsOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetCaptureAmountCentsOk() (*interface{}, bool)`

GetCaptureAmountCentsOk returns a tuple with the CaptureAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureAmountCents

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetCaptureAmountCents(v interface{})`

SetCaptureAmountCents sets CaptureAmountCents field to given value.

### HasCaptureAmountCents

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasCaptureAmountCents() bool`

HasCaptureAmountCents returns a boolean if a field has been set.

### SetCaptureAmountCentsNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetCaptureAmountCentsNil(b bool)`

 SetCaptureAmountCentsNil sets the value for CaptureAmountCents to be an explicit nil

### UnsetCaptureAmountCents
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetCaptureAmountCents()`

UnsetCaptureAmountCents ensures that no value is present for CaptureAmountCents, not even an explicit nil
### GetCaptureAmountFloat

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetCaptureAmountFloat() interface{}`

GetCaptureAmountFloat returns the CaptureAmountFloat field if non-nil, zero value otherwise.

### GetCaptureAmountFloatOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetCaptureAmountFloatOk() (*interface{}, bool)`

GetCaptureAmountFloatOk returns a tuple with the CaptureAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureAmountFloat

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetCaptureAmountFloat(v interface{})`

SetCaptureAmountFloat sets CaptureAmountFloat field to given value.

### HasCaptureAmountFloat

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasCaptureAmountFloat() bool`

HasCaptureAmountFloat returns a boolean if a field has been set.

### SetCaptureAmountFloatNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetCaptureAmountFloatNil(b bool)`

 SetCaptureAmountFloatNil sets the value for CaptureAmountFloat to be an explicit nil

### UnsetCaptureAmountFloat
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetCaptureAmountFloat()`

UnsetCaptureAmountFloat ensures that no value is present for CaptureAmountFloat, not even an explicit nil
### GetFormattedCaptureAmount

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetFormattedCaptureAmount() interface{}`

GetFormattedCaptureAmount returns the FormattedCaptureAmount field if non-nil, zero value otherwise.

### GetFormattedCaptureAmountOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetFormattedCaptureAmountOk() (*interface{}, bool)`

GetFormattedCaptureAmountOk returns a tuple with the FormattedCaptureAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedCaptureAmount

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetFormattedCaptureAmount(v interface{})`

SetFormattedCaptureAmount sets FormattedCaptureAmount field to given value.

### HasFormattedCaptureAmount

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasFormattedCaptureAmount() bool`

HasFormattedCaptureAmount returns a boolean if a field has been set.

### SetFormattedCaptureAmountNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetFormattedCaptureAmountNil(b bool)`

 SetFormattedCaptureAmountNil sets the value for FormattedCaptureAmount to be an explicit nil

### UnsetFormattedCaptureAmount
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetFormattedCaptureAmount()`

UnsetFormattedCaptureAmount ensures that no value is present for FormattedCaptureAmount, not even an explicit nil
### GetCaptureBalanceCents

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetCaptureBalanceCents() interface{}`

GetCaptureBalanceCents returns the CaptureBalanceCents field if non-nil, zero value otherwise.

### GetCaptureBalanceCentsOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetCaptureBalanceCentsOk() (*interface{}, bool)`

GetCaptureBalanceCentsOk returns a tuple with the CaptureBalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureBalanceCents

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetCaptureBalanceCents(v interface{})`

SetCaptureBalanceCents sets CaptureBalanceCents field to given value.

### HasCaptureBalanceCents

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasCaptureBalanceCents() bool`

HasCaptureBalanceCents returns a boolean if a field has been set.

### SetCaptureBalanceCentsNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetCaptureBalanceCentsNil(b bool)`

 SetCaptureBalanceCentsNil sets the value for CaptureBalanceCents to be an explicit nil

### UnsetCaptureBalanceCents
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetCaptureBalanceCents()`

UnsetCaptureBalanceCents ensures that no value is present for CaptureBalanceCents, not even an explicit nil
### GetCaptureBalanceFloat

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetCaptureBalanceFloat() interface{}`

GetCaptureBalanceFloat returns the CaptureBalanceFloat field if non-nil, zero value otherwise.

### GetCaptureBalanceFloatOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetCaptureBalanceFloatOk() (*interface{}, bool)`

GetCaptureBalanceFloatOk returns a tuple with the CaptureBalanceFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureBalanceFloat

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetCaptureBalanceFloat(v interface{})`

SetCaptureBalanceFloat sets CaptureBalanceFloat field to given value.

### HasCaptureBalanceFloat

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasCaptureBalanceFloat() bool`

HasCaptureBalanceFloat returns a boolean if a field has been set.

### SetCaptureBalanceFloatNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetCaptureBalanceFloatNil(b bool)`

 SetCaptureBalanceFloatNil sets the value for CaptureBalanceFloat to be an explicit nil

### UnsetCaptureBalanceFloat
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetCaptureBalanceFloat()`

UnsetCaptureBalanceFloat ensures that no value is present for CaptureBalanceFloat, not even an explicit nil
### GetFormattedCaptureBalance

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetFormattedCaptureBalance() interface{}`

GetFormattedCaptureBalance returns the FormattedCaptureBalance field if non-nil, zero value otherwise.

### GetFormattedCaptureBalanceOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetFormattedCaptureBalanceOk() (*interface{}, bool)`

GetFormattedCaptureBalanceOk returns a tuple with the FormattedCaptureBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedCaptureBalance

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetFormattedCaptureBalance(v interface{})`

SetFormattedCaptureBalance sets FormattedCaptureBalance field to given value.

### HasFormattedCaptureBalance

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasFormattedCaptureBalance() bool`

HasFormattedCaptureBalance returns a boolean if a field has been set.

### SetFormattedCaptureBalanceNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetFormattedCaptureBalanceNil(b bool)`

 SetFormattedCaptureBalanceNil sets the value for FormattedCaptureBalance to be an explicit nil

### UnsetFormattedCaptureBalance
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetFormattedCaptureBalance()`

UnsetFormattedCaptureBalance ensures that no value is present for FormattedCaptureBalance, not even an explicit nil
### GetVoidBalanceCents

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetVoidBalanceCents() interface{}`

GetVoidBalanceCents returns the VoidBalanceCents field if non-nil, zero value otherwise.

### GetVoidBalanceCentsOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetVoidBalanceCentsOk() (*interface{}, bool)`

GetVoidBalanceCentsOk returns a tuple with the VoidBalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidBalanceCents

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetVoidBalanceCents(v interface{})`

SetVoidBalanceCents sets VoidBalanceCents field to given value.

### HasVoidBalanceCents

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasVoidBalanceCents() bool`

HasVoidBalanceCents returns a boolean if a field has been set.

### SetVoidBalanceCentsNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetVoidBalanceCentsNil(b bool)`

 SetVoidBalanceCentsNil sets the value for VoidBalanceCents to be an explicit nil

### UnsetVoidBalanceCents
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetVoidBalanceCents()`

UnsetVoidBalanceCents ensures that no value is present for VoidBalanceCents, not even an explicit nil
### GetVoidBalanceFloat

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetVoidBalanceFloat() interface{}`

GetVoidBalanceFloat returns the VoidBalanceFloat field if non-nil, zero value otherwise.

### GetVoidBalanceFloatOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetVoidBalanceFloatOk() (*interface{}, bool)`

GetVoidBalanceFloatOk returns a tuple with the VoidBalanceFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidBalanceFloat

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetVoidBalanceFloat(v interface{})`

SetVoidBalanceFloat sets VoidBalanceFloat field to given value.

### HasVoidBalanceFloat

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasVoidBalanceFloat() bool`

HasVoidBalanceFloat returns a boolean if a field has been set.

### SetVoidBalanceFloatNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetVoidBalanceFloatNil(b bool)`

 SetVoidBalanceFloatNil sets the value for VoidBalanceFloat to be an explicit nil

### UnsetVoidBalanceFloat
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetVoidBalanceFloat()`

UnsetVoidBalanceFloat ensures that no value is present for VoidBalanceFloat, not even an explicit nil
### GetFormattedVoidBalance

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetFormattedVoidBalance() interface{}`

GetFormattedVoidBalance returns the FormattedVoidBalance field if non-nil, zero value otherwise.

### GetFormattedVoidBalanceOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) GetFormattedVoidBalanceOk() (*interface{}, bool)`

GetFormattedVoidBalanceOk returns a tuple with the FormattedVoidBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedVoidBalance

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetFormattedVoidBalance(v interface{})`

SetFormattedVoidBalance sets FormattedVoidBalance field to given value.

### HasFormattedVoidBalance

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) HasFormattedVoidBalance() bool`

HasFormattedVoidBalance returns a boolean if a field has been set.

### SetFormattedVoidBalanceNil

`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) SetFormattedVoidBalanceNil(b bool)`

 SetFormattedVoidBalanceNil sets the value for FormattedVoidBalance to be an explicit nil

### UnsetFormattedVoidBalance
`func (o *GETAuthorizationsAuthorizationId200ResponseDataAttributes) UnsetFormattedVoidBalance()`

UnsetFormattedVoidBalance ensures that no value is present for FormattedVoidBalance, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


