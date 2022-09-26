# RefundDataAttributes

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

## Methods

### NewRefundDataAttributes

`func NewRefundDataAttributes() *RefundDataAttributes`

NewRefundDataAttributes instantiates a new RefundDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundDataAttributesWithDefaults

`func NewRefundDataAttributesWithDefaults() *RefundDataAttributes`

NewRefundDataAttributesWithDefaults instantiates a new RefundDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *RefundDataAttributes) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *RefundDataAttributes) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *RefundDataAttributes) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *RefundDataAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *RefundDataAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *RefundDataAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *RefundDataAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *RefundDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetAmountCents

`func (o *RefundDataAttributes) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *RefundDataAttributes) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *RefundDataAttributes) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *RefundDataAttributes) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetAmountFloat

`func (o *RefundDataAttributes) GetAmountFloat() float32`

GetAmountFloat returns the AmountFloat field if non-nil, zero value otherwise.

### GetAmountFloatOk

`func (o *RefundDataAttributes) GetAmountFloatOk() (*float32, bool)`

GetAmountFloatOk returns a tuple with the AmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountFloat

`func (o *RefundDataAttributes) SetAmountFloat(v float32)`

SetAmountFloat sets AmountFloat field to given value.

### HasAmountFloat

`func (o *RefundDataAttributes) HasAmountFloat() bool`

HasAmountFloat returns a boolean if a field has been set.

### GetFormattedAmount

`func (o *RefundDataAttributes) GetFormattedAmount() string`

GetFormattedAmount returns the FormattedAmount field if non-nil, zero value otherwise.

### GetFormattedAmountOk

`func (o *RefundDataAttributes) GetFormattedAmountOk() (*string, bool)`

GetFormattedAmountOk returns a tuple with the FormattedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAmount

`func (o *RefundDataAttributes) SetFormattedAmount(v string)`

SetFormattedAmount sets FormattedAmount field to given value.

### HasFormattedAmount

`func (o *RefundDataAttributes) HasFormattedAmount() bool`

HasFormattedAmount returns a boolean if a field has been set.

### GetSucceeded

`func (o *RefundDataAttributes) GetSucceeded() bool`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *RefundDataAttributes) GetSucceededOk() (*bool, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *RefundDataAttributes) SetSucceeded(v bool)`

SetSucceeded sets Succeeded field to given value.

### HasSucceeded

`func (o *RefundDataAttributes) HasSucceeded() bool`

HasSucceeded returns a boolean if a field has been set.

### GetMessage

`func (o *RefundDataAttributes) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RefundDataAttributes) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RefundDataAttributes) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RefundDataAttributes) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrorCode

`func (o *RefundDataAttributes) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *RefundDataAttributes) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *RefundDataAttributes) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *RefundDataAttributes) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDetail

`func (o *RefundDataAttributes) GetErrorDetail() string`

GetErrorDetail returns the ErrorDetail field if non-nil, zero value otherwise.

### GetErrorDetailOk

`func (o *RefundDataAttributes) GetErrorDetailOk() (*string, bool)`

GetErrorDetailOk returns a tuple with the ErrorDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetail

`func (o *RefundDataAttributes) SetErrorDetail(v string)`

SetErrorDetail sets ErrorDetail field to given value.

### HasErrorDetail

`func (o *RefundDataAttributes) HasErrorDetail() bool`

HasErrorDetail returns a boolean if a field has been set.

### GetToken

`func (o *RefundDataAttributes) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RefundDataAttributes) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RefundDataAttributes) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RefundDataAttributes) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetGatewayTransactionId

`func (o *RefundDataAttributes) GetGatewayTransactionId() string`

GetGatewayTransactionId returns the GatewayTransactionId field if non-nil, zero value otherwise.

### GetGatewayTransactionIdOk

`func (o *RefundDataAttributes) GetGatewayTransactionIdOk() (*string, bool)`

GetGatewayTransactionIdOk returns a tuple with the GatewayTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayTransactionId

`func (o *RefundDataAttributes) SetGatewayTransactionId(v string)`

SetGatewayTransactionId sets GatewayTransactionId field to given value.

### HasGatewayTransactionId

`func (o *RefundDataAttributes) HasGatewayTransactionId() bool`

HasGatewayTransactionId returns a boolean if a field has been set.

### GetId

`func (o *RefundDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RefundDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RefundDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RefundDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RefundDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RefundDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RefundDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RefundDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RefundDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RefundDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RefundDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RefundDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *RefundDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *RefundDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *RefundDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *RefundDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *RefundDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *RefundDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *RefundDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *RefundDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *RefundDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RefundDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RefundDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RefundDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


