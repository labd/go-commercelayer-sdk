# GETAxervePayments200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | Pointer to **interface{}** | The merchant login code. | [optional] 
**ReturnUrl** | Pointer to **interface{}** | The URL where the payer is redirected after they approve the payment. | [optional] 
**PaymentRequestData** | Pointer to **interface{}** | The Axerve payment request data, collected by client. | [optional] 
**MismatchedAmounts** | Pointer to **interface{}** | Indicates if the order current amount differs form the one of the associated authorization. | [optional] 
**IntentAmountCents** | Pointer to **interface{}** | The amount of the associated payment intent, in cents. | [optional] 
**IntentAmountFloat** | Pointer to **interface{}** | The amount of the associated payment intent, float. | [optional] 
**FormattedIntentAmount** | Pointer to **interface{}** | The amount of the associated payment intent, formatted. | [optional] 
**PaymentInstrument** | Pointer to **interface{}** | Information about the payment instrument used in the transaction | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETAxervePayments200ResponseDataInnerAttributes

`func NewGETAxervePayments200ResponseDataInnerAttributes() *GETAxervePayments200ResponseDataInnerAttributes`

NewGETAxervePayments200ResponseDataInnerAttributes instantiates a new GETAxervePayments200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETAxervePayments200ResponseDataInnerAttributesWithDefaults

`func NewGETAxervePayments200ResponseDataInnerAttributesWithDefaults() *GETAxervePayments200ResponseDataInnerAttributes`

NewGETAxervePayments200ResponseDataInnerAttributesWithDefaults instantiates a new GETAxervePayments200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *GETAxervePayments200ResponseDataInnerAttributes) GetLogin() interface{}`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *GETAxervePayments200ResponseDataInnerAttributes) GetLoginOk() (*interface{}, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *GETAxervePayments200ResponseDataInnerAttributes) SetLogin(v interface{})`

SetLogin sets Login field to given value.

### HasLogin

`func (o *GETAxervePayments200ResponseDataInnerAttributes) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### SetLoginNil

`func (o *GETAxervePayments200ResponseDataInnerAttributes) SetLoginNil(b bool)`

 SetLoginNil sets the value for Login to be an explicit nil

### UnsetLogin
`func (o *GETAxervePayments200ResponseDataInnerAttributes) UnsetLogin()`

UnsetLogin ensures that no value is present for Login, not even an explicit nil
### GetReturnUrl

`func (o *GETAxervePayments200ResponseDataInnerAttributes) GetReturnUrl() interface{}`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *GETAxervePayments200ResponseDataInnerAttributes) GetReturnUrlOk() (*interface{}, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *GETAxervePayments200ResponseDataInnerAttributes) SetReturnUrl(v interface{})`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *GETAxervePayments200ResponseDataInnerAttributes) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### SetReturnUrlNil

`func (o *GETAxervePayments200ResponseDataInnerAttributes) SetReturnUrlNil(b bool)`

 SetReturnUrlNil sets the value for ReturnUrl to be an explicit nil

### UnsetReturnUrl
`func (o *GETAxervePayments200ResponseDataInnerAttributes) UnsetReturnUrl()`

UnsetReturnUrl ensures that no value is present for ReturnUrl, not even an explicit nil
### GetPaymentRequestData

`func (o *GETAxervePayments200ResponseDataInnerAttributes) GetPaymentRequestData() interface{}`

GetPaymentRequestData returns the PaymentRequestData field if non-nil, zero value otherwise.

### GetPaymentRequestDataOk

`func (o *GETAxervePayments200ResponseDataInnerAttributes) GetPaymentRequestDataOk() (*interface{}, bool)`

GetPaymentRequestDataOk returns a tuple with the PaymentRequestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequestData

`func (o *GETAxervePayments200ResponseDataInnerAttributes) SetPaymentRequestData(v interface{})`

SetPaymentRequestData sets PaymentRequestData field to given value.

### HasPaymentRequestData

`func (o *GETAxervePayments200ResponseDataInnerAttributes) HasPaymentRequestData() bool`

HasPaymentRequestData returns a boolean if a field has been set.

### SetPaymentRequestDataNil

`func (o *GETAxervePayments200ResponseDataInnerAttributes) SetPaymentRequestDataNil(b bool)`

 SetPaymentRequestDataNil sets the value for PaymentRequestData to be an explicit nil

### UnsetPaymentRequestData
`func (o *GETAxervePayments200ResponseDataInnerAttributes) UnsetPaymentRequestData()`

UnsetPaymentRequestData ensures that no value is present for PaymentRequestData, not even an explicit nil
### GetMismatchedAmounts

`func (o *GETAxervePayments200ResponseDataInnerAttributes) GetMismatchedAmounts() interface{}`

GetMismatchedAmounts returns the MismatchedAmounts field if non-nil, zero value otherwise.

### GetMismatchedAmountsOk

`func (o *GETAxervePayments200ResponseDataInnerAttributes) GetMismatchedAmountsOk() (*interface{}, bool)`

GetMismatchedAmountsOk returns a tuple with the MismatchedAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMismatchedAmounts

`func (o *GETAxervePayments200ResponseDataInnerAttributes) SetMismatchedAmounts(v interface{})`

SetMismatchedAmounts sets MismatchedAmounts field to given value.

### HasMismatchedAmounts

`func (o *GETAxervePayments200ResponseDataInnerAttributes) HasMismatchedAmounts() bool`

HasMismatchedAmounts returns a boolean if a field has been set.

### SetMismatchedAmountsNil

`func (o *GETAxervePayments200ResponseDataInnerAttributes) SetMismatchedAmountsNil(b bool)`

 SetMismatchedAmountsNil sets the value for MismatchedAmounts to be an explicit nil

### UnsetMismatchedAmounts
`func (o *GETAxervePayments200ResponseDataInnerAttributes) UnsetMismatchedAmounts()`

UnsetMismatchedAmounts ensures that no value is present for MismatchedAmounts, not even an explicit nil
### GetIntentAmountCents

`func (o *GETAxervePayments200ResponseDataInnerAttributes) GetIntentAmountCents() interface{}`

GetIntentAmountCents returns the IntentAmountCents field if non-nil, zero value otherwise.

### GetIntentAmountCentsOk

`func (o *GETAxervePayments200ResponseDataInnerAttributes) GetIntentAmountCentsOk() (*interface{}, bool)`

GetIntentAmountCentsOk returns a tuple with the IntentAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentAmountCents

`func (o *GETAxervePayments200ResponseDataInnerAttributes) SetIntentAmountCents(v interface{})`

SetIntentAmountCents sets IntentAmountCents field to given value.

### HasIntentAmountCents

`func (o *GETAxervePayments200ResponseDataInnerAttributes) HasIntentAmountCents() bool`

HasIntentAmountCents returns a boolean if a field has been set.

### SetIntentAmountCentsNil

`func (o *GETAxervePayments200ResponseDataInnerAttributes) SetIntentAmountCentsNil(b bool)`

 SetIntentAmountCentsNil sets the value for IntentAmountCents to be an explicit nil

### UnsetIntentAmountCents
`func (o *GETAxervePayments200ResponseDataInnerAttributes) UnsetIntentAmountCents()`

UnsetIntentAmountCents ensures that no value is present for IntentAmountCents, not even an explicit nil
### GetIntentAmountFloat

`func (o *GETAxervePayments200ResponseDataInnerAttributes) GetIntentAmountFloat() interface{}`

GetIntentAmountFloat returns the IntentAmountFloat field if non-nil, zero value otherwise.

### GetIntentAmountFloatOk

`func (o *GETAxervePayments200ResponseDataInnerAttributes) GetIntentAmountFloatOk() (*interface{}, bool)`

GetIntentAmountFloatOk returns a tuple with the IntentAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentAmountFloat

`func (o *GETAxervePayments200ResponseDataInnerAttributes) SetIntentAmountFloat(v interface{})`

SetIntentAmountFloat sets IntentAmountFloat field to given value.

### HasIntentAmountFloat

`func (o *GETAxervePayments200ResponseDataInnerAttributes) HasIntentAmountFloat() bool`

HasIntentAmountFloat returns a boolean if a field has been set.

### SetIntentAmountFloatNil

`func (o *GETAxervePayments200ResponseDataInnerAttributes) SetIntentAmountFloatNil(b bool)`

 SetIntentAmountFloatNil sets the value for IntentAmountFloat to be an explicit nil

### UnsetIntentAmountFloat
`func (o *GETAxervePayments200ResponseDataInnerAttributes) UnsetIntentAmountFloat()`

UnsetIntentAmountFloat ensures that no value is present for IntentAmountFloat, not even an explicit nil
### GetFormattedIntentAmount

`func (o *GETAxervePayments200ResponseDataInnerAttributes) GetFormattedIntentAmount() interface{}`

GetFormattedIntentAmount returns the FormattedIntentAmount field if non-nil, zero value otherwise.

### GetFormattedIntentAmountOk

`func (o *GETAxervePayments200ResponseDataInnerAttributes) GetFormattedIntentAmountOk() (*interface{}, bool)`

GetFormattedIntentAmountOk returns a tuple with the FormattedIntentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedIntentAmount

`func (o *GETAxervePayments200ResponseDataInnerAttributes) SetFormattedIntentAmount(v interface{})`

SetFormattedIntentAmount sets FormattedIntentAmount field to given value.

### HasFormattedIntentAmount

`func (o *GETAxervePayments200ResponseDataInnerAttributes) HasFormattedIntentAmount() bool`

HasFormattedIntentAmount returns a boolean if a field has been set.

### SetFormattedIntentAmountNil

`func (o *GETAxervePayments200ResponseDataInnerAttributes) SetFormattedIntentAmountNil(b bool)`

 SetFormattedIntentAmountNil sets the value for FormattedIntentAmount to be an explicit nil

### UnsetFormattedIntentAmount
`func (o *GETAxervePayments200ResponseDataInnerAttributes) UnsetFormattedIntentAmount()`

UnsetFormattedIntentAmount ensures that no value is present for FormattedIntentAmount, not even an explicit nil
### GetPaymentInstrument

`func (o *GETAxervePayments200ResponseDataInnerAttributes) GetPaymentInstrument() interface{}`

GetPaymentInstrument returns the PaymentInstrument field if non-nil, zero value otherwise.

### GetPaymentInstrumentOk

`func (o *GETAxervePayments200ResponseDataInnerAttributes) GetPaymentInstrumentOk() (*interface{}, bool)`

GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrument

`func (o *GETAxervePayments200ResponseDataInnerAttributes) SetPaymentInstrument(v interface{})`

SetPaymentInstrument sets PaymentInstrument field to given value.

### HasPaymentInstrument

`func (o *GETAxervePayments200ResponseDataInnerAttributes) HasPaymentInstrument() bool`

HasPaymentInstrument returns a boolean if a field has been set.

### SetPaymentInstrumentNil

`func (o *GETAxervePayments200ResponseDataInnerAttributes) SetPaymentInstrumentNil(b bool)`

 SetPaymentInstrumentNil sets the value for PaymentInstrument to be an explicit nil

### UnsetPaymentInstrument
`func (o *GETAxervePayments200ResponseDataInnerAttributes) UnsetPaymentInstrument()`

UnsetPaymentInstrument ensures that no value is present for PaymentInstrument, not even an explicit nil
### GetCreatedAt

`func (o *GETAxervePayments200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETAxervePayments200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETAxervePayments200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETAxervePayments200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETAxervePayments200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETAxervePayments200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETAxervePayments200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETAxervePayments200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETAxervePayments200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETAxervePayments200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETAxervePayments200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETAxervePayments200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETAxervePayments200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETAxervePayments200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETAxervePayments200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETAxervePayments200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETAxervePayments200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETAxervePayments200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETAxervePayments200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETAxervePayments200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETAxervePayments200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETAxervePayments200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETAxervePayments200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETAxervePayments200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETAxervePayments200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETAxervePayments200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETAxervePayments200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETAxervePayments200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETAxervePayments200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETAxervePayments200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


