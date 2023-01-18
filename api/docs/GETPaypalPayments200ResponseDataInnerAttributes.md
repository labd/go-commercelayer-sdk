# GETPaypalPayments200ResponseDataInnerAttributes

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
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

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

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetCancelUrl

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.

### HasCancelUrl

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasCancelUrl() bool`

HasCancelUrl returns a boolean if a field has been set.

### GetNoteToPayer

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetNoteToPayer() string`

GetNoteToPayer returns the NoteToPayer field if non-nil, zero value otherwise.

### GetNoteToPayerOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetNoteToPayerOk() (*string, bool)`

GetNoteToPayerOk returns a tuple with the NoteToPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteToPayer

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetNoteToPayer(v string)`

SetNoteToPayer sets NoteToPayer field to given value.

### HasNoteToPayer

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasNoteToPayer() bool`

HasNoteToPayer returns a boolean if a field has been set.

### GetPaypalPayerId

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetPaypalPayerId() string`

GetPaypalPayerId returns the PaypalPayerId field if non-nil, zero value otherwise.

### GetPaypalPayerIdOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetPaypalPayerIdOk() (*string, bool)`

GetPaypalPayerIdOk returns a tuple with the PaypalPayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalPayerId

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetPaypalPayerId(v string)`

SetPaypalPayerId sets PaypalPayerId field to given value.

### HasPaypalPayerId

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasPaypalPayerId() bool`

HasPaypalPayerId returns a boolean if a field has been set.

### GetName

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPaypalId

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetPaypalId() string`

GetPaypalId returns the PaypalId field if non-nil, zero value otherwise.

### GetPaypalIdOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetPaypalIdOk() (*string, bool)`

GetPaypalIdOk returns a tuple with the PaypalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalId

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetPaypalId(v string)`

SetPaypalId sets PaypalId field to given value.

### HasPaypalId

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasPaypalId() bool`

HasPaypalId returns a boolean if a field has been set.

### GetStatus

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApprovalUrl

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetApprovalUrl() string`

GetApprovalUrl returns the ApprovalUrl field if non-nil, zero value otherwise.

### GetApprovalUrlOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetApprovalUrlOk() (*string, bool)`

GetApprovalUrlOk returns a tuple with the ApprovalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalUrl

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetApprovalUrl(v string)`

SetApprovalUrl sets ApprovalUrl field to given value.

### HasApprovalUrl

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasApprovalUrl() bool`

HasApprovalUrl returns a boolean if a field has been set.

### GetMismatchedAmounts

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetMismatchedAmounts() bool`

GetMismatchedAmounts returns the MismatchedAmounts field if non-nil, zero value otherwise.

### GetMismatchedAmountsOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetMismatchedAmountsOk() (*bool, bool)`

GetMismatchedAmountsOk returns a tuple with the MismatchedAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMismatchedAmounts

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetMismatchedAmounts(v bool)`

SetMismatchedAmounts sets MismatchedAmounts field to given value.

### HasMismatchedAmounts

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasMismatchedAmounts() bool`

HasMismatchedAmounts returns a boolean if a field has been set.

### GetIntentAmountCents

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetIntentAmountCents() int32`

GetIntentAmountCents returns the IntentAmountCents field if non-nil, zero value otherwise.

### GetIntentAmountCentsOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetIntentAmountCentsOk() (*int32, bool)`

GetIntentAmountCentsOk returns a tuple with the IntentAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentAmountCents

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetIntentAmountCents(v int32)`

SetIntentAmountCents sets IntentAmountCents field to given value.

### HasIntentAmountCents

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasIntentAmountCents() bool`

HasIntentAmountCents returns a boolean if a field has been set.

### GetIntentAmountFloat

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetIntentAmountFloat() float32`

GetIntentAmountFloat returns the IntentAmountFloat field if non-nil, zero value otherwise.

### GetIntentAmountFloatOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetIntentAmountFloatOk() (*float32, bool)`

GetIntentAmountFloatOk returns a tuple with the IntentAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentAmountFloat

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetIntentAmountFloat(v float32)`

SetIntentAmountFloat sets IntentAmountFloat field to given value.

### HasIntentAmountFloat

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasIntentAmountFloat() bool`

HasIntentAmountFloat returns a boolean if a field has been set.

### GetFormattedIntentAmount

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetFormattedIntentAmount() string`

GetFormattedIntentAmount returns the FormattedIntentAmount field if non-nil, zero value otherwise.

### GetFormattedIntentAmountOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetFormattedIntentAmountOk() (*string, bool)`

GetFormattedIntentAmountOk returns a tuple with the FormattedIntentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedIntentAmount

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetFormattedIntentAmount(v string)`

SetFormattedIntentAmount sets FormattedIntentAmount field to given value.

### HasFormattedIntentAmount

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasFormattedIntentAmount() bool`

HasFormattedIntentAmount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETPaypalPayments200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


