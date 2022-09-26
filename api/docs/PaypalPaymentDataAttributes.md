# PaypalPaymentDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnUrl** | Pointer to **string** | The URL where the payer is redirected after they approve the payment. | [optional] 
**CancelUrl** | Pointer to **string** | The URL where the payer is redirected after they cancel the payment. | [optional] 
**NoteToPayer** | Pointer to **string** | A free-form field that you can use to send a note to the payer on PayPal. | [optional] 
**PaypalPayerId** | Pointer to **string** | The id of the payer that PayPal passes in the return_url. | [optional] 
**Name** | Pointer to **string** | The PayPal payer id (if present) | [optional] 
**PaypalId** | Pointer to **string** | The id of the PayPal payment object. | [optional] 
**Status** | Pointer to **string** | The PayPal payment status. One of &#39;created&#39; (default) or &#39;approved&#39;. | [optional] 
**ApprovalUrl** | Pointer to **string** | The URL the customer should be redirected to approve the payment. | [optional] 
**MismatchedAmounts** | Pointer to **bool** | Indicates if the order current amount differs form the one of the created payment intent. | [optional] 
**IntentAmountCents** | Pointer to **int32** | The amount of the associated payment intent, in cents. | [optional] 
**IntentAmountFloat** | Pointer to **float32** | The amount of the associated payment intent, float. | [optional] 
**FormattedIntentAmount** | Pointer to **string** | The amount of the associated payment intent, formatted. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPaypalPaymentDataAttributes

`func NewPaypalPaymentDataAttributes() *PaypalPaymentDataAttributes`

NewPaypalPaymentDataAttributes instantiates a new PaypalPaymentDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaypalPaymentDataAttributesWithDefaults

`func NewPaypalPaymentDataAttributesWithDefaults() *PaypalPaymentDataAttributes`

NewPaypalPaymentDataAttributesWithDefaults instantiates a new PaypalPaymentDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnUrl

`func (o *PaypalPaymentDataAttributes) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *PaypalPaymentDataAttributes) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *PaypalPaymentDataAttributes) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *PaypalPaymentDataAttributes) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetCancelUrl

`func (o *PaypalPaymentDataAttributes) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *PaypalPaymentDataAttributes) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *PaypalPaymentDataAttributes) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.

### HasCancelUrl

`func (o *PaypalPaymentDataAttributes) HasCancelUrl() bool`

HasCancelUrl returns a boolean if a field has been set.

### GetNoteToPayer

`func (o *PaypalPaymentDataAttributes) GetNoteToPayer() string`

GetNoteToPayer returns the NoteToPayer field if non-nil, zero value otherwise.

### GetNoteToPayerOk

`func (o *PaypalPaymentDataAttributes) GetNoteToPayerOk() (*string, bool)`

GetNoteToPayerOk returns a tuple with the NoteToPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteToPayer

`func (o *PaypalPaymentDataAttributes) SetNoteToPayer(v string)`

SetNoteToPayer sets NoteToPayer field to given value.

### HasNoteToPayer

`func (o *PaypalPaymentDataAttributes) HasNoteToPayer() bool`

HasNoteToPayer returns a boolean if a field has been set.

### GetPaypalPayerId

`func (o *PaypalPaymentDataAttributes) GetPaypalPayerId() string`

GetPaypalPayerId returns the PaypalPayerId field if non-nil, zero value otherwise.

### GetPaypalPayerIdOk

`func (o *PaypalPaymentDataAttributes) GetPaypalPayerIdOk() (*string, bool)`

GetPaypalPayerIdOk returns a tuple with the PaypalPayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalPayerId

`func (o *PaypalPaymentDataAttributes) SetPaypalPayerId(v string)`

SetPaypalPayerId sets PaypalPayerId field to given value.

### HasPaypalPayerId

`func (o *PaypalPaymentDataAttributes) HasPaypalPayerId() bool`

HasPaypalPayerId returns a boolean if a field has been set.

### GetName

`func (o *PaypalPaymentDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaypalPaymentDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaypalPaymentDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaypalPaymentDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPaypalId

`func (o *PaypalPaymentDataAttributes) GetPaypalId() string`

GetPaypalId returns the PaypalId field if non-nil, zero value otherwise.

### GetPaypalIdOk

`func (o *PaypalPaymentDataAttributes) GetPaypalIdOk() (*string, bool)`

GetPaypalIdOk returns a tuple with the PaypalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalId

`func (o *PaypalPaymentDataAttributes) SetPaypalId(v string)`

SetPaypalId sets PaypalId field to given value.

### HasPaypalId

`func (o *PaypalPaymentDataAttributes) HasPaypalId() bool`

HasPaypalId returns a boolean if a field has been set.

### GetStatus

`func (o *PaypalPaymentDataAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaypalPaymentDataAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaypalPaymentDataAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaypalPaymentDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApprovalUrl

`func (o *PaypalPaymentDataAttributes) GetApprovalUrl() string`

GetApprovalUrl returns the ApprovalUrl field if non-nil, zero value otherwise.

### GetApprovalUrlOk

`func (o *PaypalPaymentDataAttributes) GetApprovalUrlOk() (*string, bool)`

GetApprovalUrlOk returns a tuple with the ApprovalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalUrl

`func (o *PaypalPaymentDataAttributes) SetApprovalUrl(v string)`

SetApprovalUrl sets ApprovalUrl field to given value.

### HasApprovalUrl

`func (o *PaypalPaymentDataAttributes) HasApprovalUrl() bool`

HasApprovalUrl returns a boolean if a field has been set.

### GetMismatchedAmounts

`func (o *PaypalPaymentDataAttributes) GetMismatchedAmounts() bool`

GetMismatchedAmounts returns the MismatchedAmounts field if non-nil, zero value otherwise.

### GetMismatchedAmountsOk

`func (o *PaypalPaymentDataAttributes) GetMismatchedAmountsOk() (*bool, bool)`

GetMismatchedAmountsOk returns a tuple with the MismatchedAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMismatchedAmounts

`func (o *PaypalPaymentDataAttributes) SetMismatchedAmounts(v bool)`

SetMismatchedAmounts sets MismatchedAmounts field to given value.

### HasMismatchedAmounts

`func (o *PaypalPaymentDataAttributes) HasMismatchedAmounts() bool`

HasMismatchedAmounts returns a boolean if a field has been set.

### GetIntentAmountCents

`func (o *PaypalPaymentDataAttributes) GetIntentAmountCents() int32`

GetIntentAmountCents returns the IntentAmountCents field if non-nil, zero value otherwise.

### GetIntentAmountCentsOk

`func (o *PaypalPaymentDataAttributes) GetIntentAmountCentsOk() (*int32, bool)`

GetIntentAmountCentsOk returns a tuple with the IntentAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentAmountCents

`func (o *PaypalPaymentDataAttributes) SetIntentAmountCents(v int32)`

SetIntentAmountCents sets IntentAmountCents field to given value.

### HasIntentAmountCents

`func (o *PaypalPaymentDataAttributes) HasIntentAmountCents() bool`

HasIntentAmountCents returns a boolean if a field has been set.

### GetIntentAmountFloat

`func (o *PaypalPaymentDataAttributes) GetIntentAmountFloat() float32`

GetIntentAmountFloat returns the IntentAmountFloat field if non-nil, zero value otherwise.

### GetIntentAmountFloatOk

`func (o *PaypalPaymentDataAttributes) GetIntentAmountFloatOk() (*float32, bool)`

GetIntentAmountFloatOk returns a tuple with the IntentAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentAmountFloat

`func (o *PaypalPaymentDataAttributes) SetIntentAmountFloat(v float32)`

SetIntentAmountFloat sets IntentAmountFloat field to given value.

### HasIntentAmountFloat

`func (o *PaypalPaymentDataAttributes) HasIntentAmountFloat() bool`

HasIntentAmountFloat returns a boolean if a field has been set.

### GetFormattedIntentAmount

`func (o *PaypalPaymentDataAttributes) GetFormattedIntentAmount() string`

GetFormattedIntentAmount returns the FormattedIntentAmount field if non-nil, zero value otherwise.

### GetFormattedIntentAmountOk

`func (o *PaypalPaymentDataAttributes) GetFormattedIntentAmountOk() (*string, bool)`

GetFormattedIntentAmountOk returns a tuple with the FormattedIntentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedIntentAmount

`func (o *PaypalPaymentDataAttributes) SetFormattedIntentAmount(v string)`

SetFormattedIntentAmount sets FormattedIntentAmount field to given value.

### HasFormattedIntentAmount

`func (o *PaypalPaymentDataAttributes) HasFormattedIntentAmount() bool`

HasFormattedIntentAmount returns a boolean if a field has been set.

### GetId

`func (o *PaypalPaymentDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaypalPaymentDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaypalPaymentDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaypalPaymentDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PaypalPaymentDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaypalPaymentDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaypalPaymentDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PaypalPaymentDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PaypalPaymentDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PaypalPaymentDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PaypalPaymentDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PaypalPaymentDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *PaypalPaymentDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaypalPaymentDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaypalPaymentDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PaypalPaymentDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PaypalPaymentDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PaypalPaymentDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PaypalPaymentDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PaypalPaymentDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PaypalPaymentDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PaypalPaymentDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PaypalPaymentDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PaypalPaymentDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


