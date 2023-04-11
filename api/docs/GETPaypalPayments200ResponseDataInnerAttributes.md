# GETPaypalPayments200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnUrl** | Pointer to **interface{}** | The URL where the payer is redirected after they approve the payment. | [optional] 
**CancelUrl** | Pointer to **interface{}** | The URL where the payer is redirected after they cancel the payment. | [optional] 
**NoteToPayer** | Pointer to **interface{}** | A free-form field that you can use to send a note to the payer on PayPal. | [optional] 
**PaypalPayerId** | Pointer to **interface{}** | The id of the payer that PayPal passes in the return_url. | [optional] 
**Name** | Pointer to **interface{}** | The PayPal payer id (if present) | [optional] 
**PaypalId** | Pointer to **interface{}** | The id of the PayPal payment object. | [optional] 
**Status** | Pointer to **interface{}** | The PayPal payment status. One of &#39;created&#39; (default) or &#39;approved&#39;. | [optional] 
**ApprovalUrl** | Pointer to **interface{}** | The URL the customer should be redirected to approve the payment. | [optional] 
**MismatchedAmounts** | Pointer to **interface{}** | Indicates if the order current amount differs form the one of the created payment intent. | [optional] 
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

### NewGETPaypalPayments200ResponseDataInnerAttributes

`func NewGETPaypalPayments200ResponseDataInnerAttributes() *GETPaypalPayments200ResponseDataInnerAttributes`

NewGETPaypalPayments200ResponseDataInnerAttributes instantiates a new GETPaypalPayments200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETPaypalPayments200ResponseDataInnerAttributesWithDefaults

`func NewGETPaypalPayments200ResponseDataInnerAttributesWithDefaults() *GETPaypalPayments200ResponseDataInnerAttributes`

NewGETPaypalPayments200ResponseDataInnerAttributesWithDefaults instantiates a new GETPaypalPayments200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnUrl

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetReturnUrl() interface{}`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetReturnUrlOk() (*interface{}, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetReturnUrl(v interface{})`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### SetReturnUrlNil

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetReturnUrlNil(b bool)`

 SetReturnUrlNil sets the value for ReturnUrl to be an explicit nil

### UnsetReturnUrl
`func (o *GETPaypalPayments200ResponseDataInnerAttributes) UnsetReturnUrl()`

UnsetReturnUrl ensures that no value is present for ReturnUrl, not even an explicit nil
### GetCancelUrl

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetCancelUrl() interface{}`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetCancelUrlOk() (*interface{}, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetCancelUrl(v interface{})`

SetCancelUrl sets CancelUrl field to given value.

### HasCancelUrl

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasCancelUrl() bool`

HasCancelUrl returns a boolean if a field has been set.

### SetCancelUrlNil

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetCancelUrlNil(b bool)`

 SetCancelUrlNil sets the value for CancelUrl to be an explicit nil

### UnsetCancelUrl
`func (o *GETPaypalPayments200ResponseDataInnerAttributes) UnsetCancelUrl()`

UnsetCancelUrl ensures that no value is present for CancelUrl, not even an explicit nil
### GetNoteToPayer

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetNoteToPayer() interface{}`

GetNoteToPayer returns the NoteToPayer field if non-nil, zero value otherwise.

### GetNoteToPayerOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetNoteToPayerOk() (*interface{}, bool)`

GetNoteToPayerOk returns a tuple with the NoteToPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteToPayer

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetNoteToPayer(v interface{})`

SetNoteToPayer sets NoteToPayer field to given value.

### HasNoteToPayer

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasNoteToPayer() bool`

HasNoteToPayer returns a boolean if a field has been set.

### SetNoteToPayerNil

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetNoteToPayerNil(b bool)`

 SetNoteToPayerNil sets the value for NoteToPayer to be an explicit nil

### UnsetNoteToPayer
`func (o *GETPaypalPayments200ResponseDataInnerAttributes) UnsetNoteToPayer()`

UnsetNoteToPayer ensures that no value is present for NoteToPayer, not even an explicit nil
### GetPaypalPayerId

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetPaypalPayerId() interface{}`

GetPaypalPayerId returns the PaypalPayerId field if non-nil, zero value otherwise.

### GetPaypalPayerIdOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetPaypalPayerIdOk() (*interface{}, bool)`

GetPaypalPayerIdOk returns a tuple with the PaypalPayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalPayerId

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetPaypalPayerId(v interface{})`

SetPaypalPayerId sets PaypalPayerId field to given value.

### HasPaypalPayerId

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasPaypalPayerId() bool`

HasPaypalPayerId returns a boolean if a field has been set.

### SetPaypalPayerIdNil

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetPaypalPayerIdNil(b bool)`

 SetPaypalPayerIdNil sets the value for PaypalPayerId to be an explicit nil

### UnsetPaypalPayerId
`func (o *GETPaypalPayments200ResponseDataInnerAttributes) UnsetPaypalPayerId()`

UnsetPaypalPayerId ensures that no value is present for PaypalPayerId, not even an explicit nil
### GetName

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETPaypalPayments200ResponseDataInnerAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPaypalId

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetPaypalId() interface{}`

GetPaypalId returns the PaypalId field if non-nil, zero value otherwise.

### GetPaypalIdOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetPaypalIdOk() (*interface{}, bool)`

GetPaypalIdOk returns a tuple with the PaypalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalId

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetPaypalId(v interface{})`

SetPaypalId sets PaypalId field to given value.

### HasPaypalId

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasPaypalId() bool`

HasPaypalId returns a boolean if a field has been set.

### SetPaypalIdNil

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetPaypalIdNil(b bool)`

 SetPaypalIdNil sets the value for PaypalId to be an explicit nil

### UnsetPaypalId
`func (o *GETPaypalPayments200ResponseDataInnerAttributes) UnsetPaypalId()`

UnsetPaypalId ensures that no value is present for PaypalId, not even an explicit nil
### GetStatus

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GETPaypalPayments200ResponseDataInnerAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetApprovalUrl

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetApprovalUrl() interface{}`

GetApprovalUrl returns the ApprovalUrl field if non-nil, zero value otherwise.

### GetApprovalUrlOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetApprovalUrlOk() (*interface{}, bool)`

GetApprovalUrlOk returns a tuple with the ApprovalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalUrl

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetApprovalUrl(v interface{})`

SetApprovalUrl sets ApprovalUrl field to given value.

### HasApprovalUrl

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasApprovalUrl() bool`

HasApprovalUrl returns a boolean if a field has been set.

### SetApprovalUrlNil

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetApprovalUrlNil(b bool)`

 SetApprovalUrlNil sets the value for ApprovalUrl to be an explicit nil

### UnsetApprovalUrl
`func (o *GETPaypalPayments200ResponseDataInnerAttributes) UnsetApprovalUrl()`

UnsetApprovalUrl ensures that no value is present for ApprovalUrl, not even an explicit nil
### GetMismatchedAmounts

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetMismatchedAmounts() interface{}`

GetMismatchedAmounts returns the MismatchedAmounts field if non-nil, zero value otherwise.

### GetMismatchedAmountsOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetMismatchedAmountsOk() (*interface{}, bool)`

GetMismatchedAmountsOk returns a tuple with the MismatchedAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMismatchedAmounts

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetMismatchedAmounts(v interface{})`

SetMismatchedAmounts sets MismatchedAmounts field to given value.

### HasMismatchedAmounts

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasMismatchedAmounts() bool`

HasMismatchedAmounts returns a boolean if a field has been set.

### SetMismatchedAmountsNil

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetMismatchedAmountsNil(b bool)`

 SetMismatchedAmountsNil sets the value for MismatchedAmounts to be an explicit nil

### UnsetMismatchedAmounts
`func (o *GETPaypalPayments200ResponseDataInnerAttributes) UnsetMismatchedAmounts()`

UnsetMismatchedAmounts ensures that no value is present for MismatchedAmounts, not even an explicit nil
### GetIntentAmountCents

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetIntentAmountCents() interface{}`

GetIntentAmountCents returns the IntentAmountCents field if non-nil, zero value otherwise.

### GetIntentAmountCentsOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetIntentAmountCentsOk() (*interface{}, bool)`

GetIntentAmountCentsOk returns a tuple with the IntentAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentAmountCents

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetIntentAmountCents(v interface{})`

SetIntentAmountCents sets IntentAmountCents field to given value.

### HasIntentAmountCents

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasIntentAmountCents() bool`

HasIntentAmountCents returns a boolean if a field has been set.

### SetIntentAmountCentsNil

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetIntentAmountCentsNil(b bool)`

 SetIntentAmountCentsNil sets the value for IntentAmountCents to be an explicit nil

### UnsetIntentAmountCents
`func (o *GETPaypalPayments200ResponseDataInnerAttributes) UnsetIntentAmountCents()`

UnsetIntentAmountCents ensures that no value is present for IntentAmountCents, not even an explicit nil
### GetIntentAmountFloat

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetIntentAmountFloat() interface{}`

GetIntentAmountFloat returns the IntentAmountFloat field if non-nil, zero value otherwise.

### GetIntentAmountFloatOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetIntentAmountFloatOk() (*interface{}, bool)`

GetIntentAmountFloatOk returns a tuple with the IntentAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentAmountFloat

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetIntentAmountFloat(v interface{})`

SetIntentAmountFloat sets IntentAmountFloat field to given value.

### HasIntentAmountFloat

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasIntentAmountFloat() bool`

HasIntentAmountFloat returns a boolean if a field has been set.

### SetIntentAmountFloatNil

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetIntentAmountFloatNil(b bool)`

 SetIntentAmountFloatNil sets the value for IntentAmountFloat to be an explicit nil

### UnsetIntentAmountFloat
`func (o *GETPaypalPayments200ResponseDataInnerAttributes) UnsetIntentAmountFloat()`

UnsetIntentAmountFloat ensures that no value is present for IntentAmountFloat, not even an explicit nil
### GetFormattedIntentAmount

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetFormattedIntentAmount() interface{}`

GetFormattedIntentAmount returns the FormattedIntentAmount field if non-nil, zero value otherwise.

### GetFormattedIntentAmountOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetFormattedIntentAmountOk() (*interface{}, bool)`

GetFormattedIntentAmountOk returns a tuple with the FormattedIntentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedIntentAmount

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetFormattedIntentAmount(v interface{})`

SetFormattedIntentAmount sets FormattedIntentAmount field to given value.

### HasFormattedIntentAmount

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasFormattedIntentAmount() bool`

HasFormattedIntentAmount returns a boolean if a field has been set.

### SetFormattedIntentAmountNil

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetFormattedIntentAmountNil(b bool)`

 SetFormattedIntentAmountNil sets the value for FormattedIntentAmount to be an explicit nil

### UnsetFormattedIntentAmount
`func (o *GETPaypalPayments200ResponseDataInnerAttributes) UnsetFormattedIntentAmount()`

UnsetFormattedIntentAmount ensures that no value is present for FormattedIntentAmount, not even an explicit nil
### GetPaymentInstrument

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetPaymentInstrument() interface{}`

GetPaymentInstrument returns the PaymentInstrument field if non-nil, zero value otherwise.

### GetPaymentInstrumentOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetPaymentInstrumentOk() (*interface{}, bool)`

GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrument

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetPaymentInstrument(v interface{})`

SetPaymentInstrument sets PaymentInstrument field to given value.

### HasPaymentInstrument

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasPaymentInstrument() bool`

HasPaymentInstrument returns a boolean if a field has been set.

### SetPaymentInstrumentNil

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetPaymentInstrumentNil(b bool)`

 SetPaymentInstrumentNil sets the value for PaymentInstrument to be an explicit nil

### UnsetPaymentInstrument
`func (o *GETPaypalPayments200ResponseDataInnerAttributes) UnsetPaymentInstrument()`

UnsetPaymentInstrument ensures that no value is present for PaymentInstrument, not even an explicit nil
### GetCreatedAt

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETPaypalPayments200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETPaypalPayments200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETPaypalPayments200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETPaypalPayments200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETPaypalPayments200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


