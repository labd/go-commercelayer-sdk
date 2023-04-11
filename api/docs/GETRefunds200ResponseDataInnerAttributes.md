# GETRefunds200ResponseDataInnerAttributes

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

## Methods

### NewGETRefunds200ResponseDataInnerAttributes

`func NewGETRefunds200ResponseDataInnerAttributes() *GETRefunds200ResponseDataInnerAttributes`

NewGETRefunds200ResponseDataInnerAttributes instantiates a new GETRefunds200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETRefunds200ResponseDataInnerAttributesWithDefaults

`func NewGETRefunds200ResponseDataInnerAttributesWithDefaults() *GETRefunds200ResponseDataInnerAttributes`

NewGETRefunds200ResponseDataInnerAttributesWithDefaults instantiates a new GETRefunds200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *GETRefunds200ResponseDataInnerAttributes) GetNumber() interface{}`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GETRefunds200ResponseDataInnerAttributes) GetNumberOk() (*interface{}, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GETRefunds200ResponseDataInnerAttributes) SetNumber(v interface{})`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GETRefunds200ResponseDataInnerAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *GETRefunds200ResponseDataInnerAttributes) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *GETRefunds200ResponseDataInnerAttributes) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetCurrencyCode

`func (o *GETRefunds200ResponseDataInnerAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETRefunds200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETRefunds200ResponseDataInnerAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETRefunds200ResponseDataInnerAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *GETRefunds200ResponseDataInnerAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *GETRefunds200ResponseDataInnerAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetAmountCents

`func (o *GETRefunds200ResponseDataInnerAttributes) GetAmountCents() interface{}`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *GETRefunds200ResponseDataInnerAttributes) GetAmountCentsOk() (*interface{}, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *GETRefunds200ResponseDataInnerAttributes) SetAmountCents(v interface{})`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *GETRefunds200ResponseDataInnerAttributes) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### SetAmountCentsNil

`func (o *GETRefunds200ResponseDataInnerAttributes) SetAmountCentsNil(b bool)`

 SetAmountCentsNil sets the value for AmountCents to be an explicit nil

### UnsetAmountCents
`func (o *GETRefunds200ResponseDataInnerAttributes) UnsetAmountCents()`

UnsetAmountCents ensures that no value is present for AmountCents, not even an explicit nil
### GetAmountFloat

`func (o *GETRefunds200ResponseDataInnerAttributes) GetAmountFloat() interface{}`

GetAmountFloat returns the AmountFloat field if non-nil, zero value otherwise.

### GetAmountFloatOk

`func (o *GETRefunds200ResponseDataInnerAttributes) GetAmountFloatOk() (*interface{}, bool)`

GetAmountFloatOk returns a tuple with the AmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountFloat

`func (o *GETRefunds200ResponseDataInnerAttributes) SetAmountFloat(v interface{})`

SetAmountFloat sets AmountFloat field to given value.

### HasAmountFloat

`func (o *GETRefunds200ResponseDataInnerAttributes) HasAmountFloat() bool`

HasAmountFloat returns a boolean if a field has been set.

### SetAmountFloatNil

`func (o *GETRefunds200ResponseDataInnerAttributes) SetAmountFloatNil(b bool)`

 SetAmountFloatNil sets the value for AmountFloat to be an explicit nil

### UnsetAmountFloat
`func (o *GETRefunds200ResponseDataInnerAttributes) UnsetAmountFloat()`

UnsetAmountFloat ensures that no value is present for AmountFloat, not even an explicit nil
### GetFormattedAmount

`func (o *GETRefunds200ResponseDataInnerAttributes) GetFormattedAmount() interface{}`

GetFormattedAmount returns the FormattedAmount field if non-nil, zero value otherwise.

### GetFormattedAmountOk

`func (o *GETRefunds200ResponseDataInnerAttributes) GetFormattedAmountOk() (*interface{}, bool)`

GetFormattedAmountOk returns a tuple with the FormattedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAmount

`func (o *GETRefunds200ResponseDataInnerAttributes) SetFormattedAmount(v interface{})`

SetFormattedAmount sets FormattedAmount field to given value.

### HasFormattedAmount

`func (o *GETRefunds200ResponseDataInnerAttributes) HasFormattedAmount() bool`

HasFormattedAmount returns a boolean if a field has been set.

### SetFormattedAmountNil

`func (o *GETRefunds200ResponseDataInnerAttributes) SetFormattedAmountNil(b bool)`

 SetFormattedAmountNil sets the value for FormattedAmount to be an explicit nil

### UnsetFormattedAmount
`func (o *GETRefunds200ResponseDataInnerAttributes) UnsetFormattedAmount()`

UnsetFormattedAmount ensures that no value is present for FormattedAmount, not even an explicit nil
### GetSucceeded

`func (o *GETRefunds200ResponseDataInnerAttributes) GetSucceeded() interface{}`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *GETRefunds200ResponseDataInnerAttributes) GetSucceededOk() (*interface{}, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *GETRefunds200ResponseDataInnerAttributes) SetSucceeded(v interface{})`

SetSucceeded sets Succeeded field to given value.

### HasSucceeded

`func (o *GETRefunds200ResponseDataInnerAttributes) HasSucceeded() bool`

HasSucceeded returns a boolean if a field has been set.

### SetSucceededNil

`func (o *GETRefunds200ResponseDataInnerAttributes) SetSucceededNil(b bool)`

 SetSucceededNil sets the value for Succeeded to be an explicit nil

### UnsetSucceeded
`func (o *GETRefunds200ResponseDataInnerAttributes) UnsetSucceeded()`

UnsetSucceeded ensures that no value is present for Succeeded, not even an explicit nil
### GetMessage

`func (o *GETRefunds200ResponseDataInnerAttributes) GetMessage() interface{}`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GETRefunds200ResponseDataInnerAttributes) GetMessageOk() (*interface{}, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GETRefunds200ResponseDataInnerAttributes) SetMessage(v interface{})`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GETRefunds200ResponseDataInnerAttributes) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *GETRefunds200ResponseDataInnerAttributes) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *GETRefunds200ResponseDataInnerAttributes) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetErrorCode

`func (o *GETRefunds200ResponseDataInnerAttributes) GetErrorCode() interface{}`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GETRefunds200ResponseDataInnerAttributes) GetErrorCodeOk() (*interface{}, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GETRefunds200ResponseDataInnerAttributes) SetErrorCode(v interface{})`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *GETRefunds200ResponseDataInnerAttributes) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *GETRefunds200ResponseDataInnerAttributes) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *GETRefunds200ResponseDataInnerAttributes) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorDetail

`func (o *GETRefunds200ResponseDataInnerAttributes) GetErrorDetail() interface{}`

GetErrorDetail returns the ErrorDetail field if non-nil, zero value otherwise.

### GetErrorDetailOk

`func (o *GETRefunds200ResponseDataInnerAttributes) GetErrorDetailOk() (*interface{}, bool)`

GetErrorDetailOk returns a tuple with the ErrorDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetail

`func (o *GETRefunds200ResponseDataInnerAttributes) SetErrorDetail(v interface{})`

SetErrorDetail sets ErrorDetail field to given value.

### HasErrorDetail

`func (o *GETRefunds200ResponseDataInnerAttributes) HasErrorDetail() bool`

HasErrorDetail returns a boolean if a field has been set.

### SetErrorDetailNil

`func (o *GETRefunds200ResponseDataInnerAttributes) SetErrorDetailNil(b bool)`

 SetErrorDetailNil sets the value for ErrorDetail to be an explicit nil

### UnsetErrorDetail
`func (o *GETRefunds200ResponseDataInnerAttributes) UnsetErrorDetail()`

UnsetErrorDetail ensures that no value is present for ErrorDetail, not even an explicit nil
### GetToken

`func (o *GETRefunds200ResponseDataInnerAttributes) GetToken() interface{}`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GETRefunds200ResponseDataInnerAttributes) GetTokenOk() (*interface{}, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GETRefunds200ResponseDataInnerAttributes) SetToken(v interface{})`

SetToken sets Token field to given value.

### HasToken

`func (o *GETRefunds200ResponseDataInnerAttributes) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *GETRefunds200ResponseDataInnerAttributes) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *GETRefunds200ResponseDataInnerAttributes) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetGatewayTransactionId

`func (o *GETRefunds200ResponseDataInnerAttributes) GetGatewayTransactionId() interface{}`

GetGatewayTransactionId returns the GatewayTransactionId field if non-nil, zero value otherwise.

### GetGatewayTransactionIdOk

`func (o *GETRefunds200ResponseDataInnerAttributes) GetGatewayTransactionIdOk() (*interface{}, bool)`

GetGatewayTransactionIdOk returns a tuple with the GatewayTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayTransactionId

`func (o *GETRefunds200ResponseDataInnerAttributes) SetGatewayTransactionId(v interface{})`

SetGatewayTransactionId sets GatewayTransactionId field to given value.

### HasGatewayTransactionId

`func (o *GETRefunds200ResponseDataInnerAttributes) HasGatewayTransactionId() bool`

HasGatewayTransactionId returns a boolean if a field has been set.

### SetGatewayTransactionIdNil

`func (o *GETRefunds200ResponseDataInnerAttributes) SetGatewayTransactionIdNil(b bool)`

 SetGatewayTransactionIdNil sets the value for GatewayTransactionId to be an explicit nil

### UnsetGatewayTransactionId
`func (o *GETRefunds200ResponseDataInnerAttributes) UnsetGatewayTransactionId()`

UnsetGatewayTransactionId ensures that no value is present for GatewayTransactionId, not even an explicit nil
### GetCreatedAt

`func (o *GETRefunds200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETRefunds200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETRefunds200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETRefunds200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETRefunds200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETRefunds200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETRefunds200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETRefunds200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETRefunds200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETRefunds200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETRefunds200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETRefunds200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETRefunds200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETRefunds200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETRefunds200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETRefunds200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETRefunds200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETRefunds200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETRefunds200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETRefunds200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETRefunds200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETRefunds200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETRefunds200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETRefunds200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETRefunds200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETRefunds200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETRefunds200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETRefunds200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETRefunds200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETRefunds200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


