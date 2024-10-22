# GETRefundsRefundId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **interface{}** | The transaction number, auto generated. | [optional] 
**Type** | Pointer to **interface{}** | The transaction&#39;s type. | [optional] 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard, inherited from the associated order. | [optional] 
**AmountCents** | Pointer to **interface{}** | The transaction amount, in cents. | [optional] 
**AmountFloat** | Pointer to **interface{}** | The transaction amount, float. | [optional] 
**FormattedAmount** | Pointer to **interface{}** | The transaction amount, formatted. | [optional] 
**Succeeded** | Pointer to **interface{}** | Indicates if the transaction is successful. | [optional] 
**Message** | Pointer to **interface{}** | The message returned by the payment gateway. | [optional] 
**ErrorCode** | Pointer to **interface{}** | The error code, if any, returned by the payment gateway. | [optional] 
**ErrorDetail** | Pointer to **interface{}** | The error detail, if any, returned by the payment gateway. | [optional] 
**Token** | Pointer to **interface{}** | The token identifying the transaction, returned by the payment gateway. | [optional] 
**GatewayTransactionId** | Pointer to **interface{}** | The ID identifying the transaction, returned by the payment gateway. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETRefundsRefundId200ResponseDataAttributes

`func NewGETRefundsRefundId200ResponseDataAttributes() *GETRefundsRefundId200ResponseDataAttributes`

NewGETRefundsRefundId200ResponseDataAttributes instantiates a new GETRefundsRefundId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETRefundsRefundId200ResponseDataAttributesWithDefaults

`func NewGETRefundsRefundId200ResponseDataAttributesWithDefaults() *GETRefundsRefundId200ResponseDataAttributes`

NewGETRefundsRefundId200ResponseDataAttributesWithDefaults instantiates a new GETRefundsRefundId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetNumber() interface{}`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetNumberOk() (*interface{}, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetNumber(v interface{})`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GETRefundsRefundId200ResponseDataAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *GETRefundsRefundId200ResponseDataAttributes) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetType

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetType(v interface{})`

SetType sets Type field to given value.

### HasType

`func (o *GETRefundsRefundId200ResponseDataAttributes) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *GETRefundsRefundId200ResponseDataAttributes) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetCurrencyCode

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETRefundsRefundId200ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *GETRefundsRefundId200ResponseDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetAmountCents

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetAmountCents() interface{}`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetAmountCentsOk() (*interface{}, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetAmountCents(v interface{})`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *GETRefundsRefundId200ResponseDataAttributes) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### SetAmountCentsNil

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetAmountCentsNil(b bool)`

 SetAmountCentsNil sets the value for AmountCents to be an explicit nil

### UnsetAmountCents
`func (o *GETRefundsRefundId200ResponseDataAttributes) UnsetAmountCents()`

UnsetAmountCents ensures that no value is present for AmountCents, not even an explicit nil
### GetAmountFloat

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetAmountFloat() interface{}`

GetAmountFloat returns the AmountFloat field if non-nil, zero value otherwise.

### GetAmountFloatOk

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetAmountFloatOk() (*interface{}, bool)`

GetAmountFloatOk returns a tuple with the AmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountFloat

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetAmountFloat(v interface{})`

SetAmountFloat sets AmountFloat field to given value.

### HasAmountFloat

`func (o *GETRefundsRefundId200ResponseDataAttributes) HasAmountFloat() bool`

HasAmountFloat returns a boolean if a field has been set.

### SetAmountFloatNil

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetAmountFloatNil(b bool)`

 SetAmountFloatNil sets the value for AmountFloat to be an explicit nil

### UnsetAmountFloat
`func (o *GETRefundsRefundId200ResponseDataAttributes) UnsetAmountFloat()`

UnsetAmountFloat ensures that no value is present for AmountFloat, not even an explicit nil
### GetFormattedAmount

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetFormattedAmount() interface{}`

GetFormattedAmount returns the FormattedAmount field if non-nil, zero value otherwise.

### GetFormattedAmountOk

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetFormattedAmountOk() (*interface{}, bool)`

GetFormattedAmountOk returns a tuple with the FormattedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAmount

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetFormattedAmount(v interface{})`

SetFormattedAmount sets FormattedAmount field to given value.

### HasFormattedAmount

`func (o *GETRefundsRefundId200ResponseDataAttributes) HasFormattedAmount() bool`

HasFormattedAmount returns a boolean if a field has been set.

### SetFormattedAmountNil

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetFormattedAmountNil(b bool)`

 SetFormattedAmountNil sets the value for FormattedAmount to be an explicit nil

### UnsetFormattedAmount
`func (o *GETRefundsRefundId200ResponseDataAttributes) UnsetFormattedAmount()`

UnsetFormattedAmount ensures that no value is present for FormattedAmount, not even an explicit nil
### GetSucceeded

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetSucceeded() interface{}`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetSucceededOk() (*interface{}, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetSucceeded(v interface{})`

SetSucceeded sets Succeeded field to given value.

### HasSucceeded

`func (o *GETRefundsRefundId200ResponseDataAttributes) HasSucceeded() bool`

HasSucceeded returns a boolean if a field has been set.

### SetSucceededNil

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetSucceededNil(b bool)`

 SetSucceededNil sets the value for Succeeded to be an explicit nil

### UnsetSucceeded
`func (o *GETRefundsRefundId200ResponseDataAttributes) UnsetSucceeded()`

UnsetSucceeded ensures that no value is present for Succeeded, not even an explicit nil
### GetMessage

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetMessage() interface{}`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetMessageOk() (*interface{}, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetMessage(v interface{})`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GETRefundsRefundId200ResponseDataAttributes) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *GETRefundsRefundId200ResponseDataAttributes) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetErrorCode

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetErrorCode() interface{}`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetErrorCodeOk() (*interface{}, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetErrorCode(v interface{})`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *GETRefundsRefundId200ResponseDataAttributes) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *GETRefundsRefundId200ResponseDataAttributes) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorDetail

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetErrorDetail() interface{}`

GetErrorDetail returns the ErrorDetail field if non-nil, zero value otherwise.

### GetErrorDetailOk

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetErrorDetailOk() (*interface{}, bool)`

GetErrorDetailOk returns a tuple with the ErrorDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetail

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetErrorDetail(v interface{})`

SetErrorDetail sets ErrorDetail field to given value.

### HasErrorDetail

`func (o *GETRefundsRefundId200ResponseDataAttributes) HasErrorDetail() bool`

HasErrorDetail returns a boolean if a field has been set.

### SetErrorDetailNil

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetErrorDetailNil(b bool)`

 SetErrorDetailNil sets the value for ErrorDetail to be an explicit nil

### UnsetErrorDetail
`func (o *GETRefundsRefundId200ResponseDataAttributes) UnsetErrorDetail()`

UnsetErrorDetail ensures that no value is present for ErrorDetail, not even an explicit nil
### GetToken

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetToken() interface{}`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetTokenOk() (*interface{}, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetToken(v interface{})`

SetToken sets Token field to given value.

### HasToken

`func (o *GETRefundsRefundId200ResponseDataAttributes) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *GETRefundsRefundId200ResponseDataAttributes) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetGatewayTransactionId

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetGatewayTransactionId() interface{}`

GetGatewayTransactionId returns the GatewayTransactionId field if non-nil, zero value otherwise.

### GetGatewayTransactionIdOk

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetGatewayTransactionIdOk() (*interface{}, bool)`

GetGatewayTransactionIdOk returns a tuple with the GatewayTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayTransactionId

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetGatewayTransactionId(v interface{})`

SetGatewayTransactionId sets GatewayTransactionId field to given value.

### HasGatewayTransactionId

`func (o *GETRefundsRefundId200ResponseDataAttributes) HasGatewayTransactionId() bool`

HasGatewayTransactionId returns a boolean if a field has been set.

### SetGatewayTransactionIdNil

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetGatewayTransactionIdNil(b bool)`

 SetGatewayTransactionIdNil sets the value for GatewayTransactionId to be an explicit nil

### UnsetGatewayTransactionId
`func (o *GETRefundsRefundId200ResponseDataAttributes) UnsetGatewayTransactionId()`

UnsetGatewayTransactionId ensures that no value is present for GatewayTransactionId, not even an explicit nil
### GetCreatedAt

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETRefundsRefundId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETRefundsRefundId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETRefundsRefundId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETRefundsRefundId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETRefundsRefundId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETRefundsRefundId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETRefundsRefundId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETRefundsRefundId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETRefundsRefundId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETRefundsRefundId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETRefundsRefundId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETRefundsRefundId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


