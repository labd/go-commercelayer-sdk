# GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnUrl** | Pointer to **interface{}** | The URL where the payer is redirected after they approve the payment. | [optional] 
**CancelUrl** | Pointer to **interface{}** | The URL where the payer is redirected after they cancel the payment. | [optional] 
**NoteToPayer** | Pointer to **interface{}** | A free-form field that you can use to send a note to the payer on PayPal. | [optional] 
**PaypalPayerId** | Pointer to **interface{}** | The id of the payer that PayPal passes in the return_url. | [optional] 
**Name** | Pointer to **interface{}** | The PayPal payer id (if present). | [optional] 
**PaypalId** | Pointer to **interface{}** | The id of the PayPal payment object. | [optional] 
**Status** | Pointer to **interface{}** | The PayPal payment status. One of &#39;created&#39;, or &#39;approved&#39;. | [optional] 
**ApprovalUrl** | Pointer to **interface{}** | The URL the customer should be redirected to approve the payment. | [optional] 
**MismatchedAmounts** | Pointer to **interface{}** | Indicates if the order current amount differs form the one of the created payment intent. | [optional] 
**PaymentInstrument** | Pointer to **interface{}** | Information about the payment instrument used in the transaction. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes

`func NewGETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes() *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes`

NewGETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes instantiates a new GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETPaypalPaymentsPaypalPaymentId200ResponseDataAttributesWithDefaults

`func NewGETPaypalPaymentsPaypalPaymentId200ResponseDataAttributesWithDefaults() *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes`

NewGETPaypalPaymentsPaypalPaymentId200ResponseDataAttributesWithDefaults instantiates a new GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnUrl

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetReturnUrl() interface{}`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetReturnUrlOk() (*interface{}, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetReturnUrl(v interface{})`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### SetReturnUrlNil

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetReturnUrlNil(b bool)`

 SetReturnUrlNil sets the value for ReturnUrl to be an explicit nil

### UnsetReturnUrl
`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) UnsetReturnUrl()`

UnsetReturnUrl ensures that no value is present for ReturnUrl, not even an explicit nil
### GetCancelUrl

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetCancelUrl() interface{}`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetCancelUrlOk() (*interface{}, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetCancelUrl(v interface{})`

SetCancelUrl sets CancelUrl field to given value.

### HasCancelUrl

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) HasCancelUrl() bool`

HasCancelUrl returns a boolean if a field has been set.

### SetCancelUrlNil

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetCancelUrlNil(b bool)`

 SetCancelUrlNil sets the value for CancelUrl to be an explicit nil

### UnsetCancelUrl
`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) UnsetCancelUrl()`

UnsetCancelUrl ensures that no value is present for CancelUrl, not even an explicit nil
### GetNoteToPayer

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetNoteToPayer() interface{}`

GetNoteToPayer returns the NoteToPayer field if non-nil, zero value otherwise.

### GetNoteToPayerOk

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetNoteToPayerOk() (*interface{}, bool)`

GetNoteToPayerOk returns a tuple with the NoteToPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteToPayer

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetNoteToPayer(v interface{})`

SetNoteToPayer sets NoteToPayer field to given value.

### HasNoteToPayer

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) HasNoteToPayer() bool`

HasNoteToPayer returns a boolean if a field has been set.

### SetNoteToPayerNil

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetNoteToPayerNil(b bool)`

 SetNoteToPayerNil sets the value for NoteToPayer to be an explicit nil

### UnsetNoteToPayer
`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) UnsetNoteToPayer()`

UnsetNoteToPayer ensures that no value is present for NoteToPayer, not even an explicit nil
### GetPaypalPayerId

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetPaypalPayerId() interface{}`

GetPaypalPayerId returns the PaypalPayerId field if non-nil, zero value otherwise.

### GetPaypalPayerIdOk

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetPaypalPayerIdOk() (*interface{}, bool)`

GetPaypalPayerIdOk returns a tuple with the PaypalPayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalPayerId

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetPaypalPayerId(v interface{})`

SetPaypalPayerId sets PaypalPayerId field to given value.

### HasPaypalPayerId

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) HasPaypalPayerId() bool`

HasPaypalPayerId returns a boolean if a field has been set.

### SetPaypalPayerIdNil

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetPaypalPayerIdNil(b bool)`

 SetPaypalPayerIdNil sets the value for PaypalPayerId to be an explicit nil

### UnsetPaypalPayerId
`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) UnsetPaypalPayerId()`

UnsetPaypalPayerId ensures that no value is present for PaypalPayerId, not even an explicit nil
### GetName

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPaypalId

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetPaypalId() interface{}`

GetPaypalId returns the PaypalId field if non-nil, zero value otherwise.

### GetPaypalIdOk

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetPaypalIdOk() (*interface{}, bool)`

GetPaypalIdOk returns a tuple with the PaypalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalId

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetPaypalId(v interface{})`

SetPaypalId sets PaypalId field to given value.

### HasPaypalId

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) HasPaypalId() bool`

HasPaypalId returns a boolean if a field has been set.

### SetPaypalIdNil

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetPaypalIdNil(b bool)`

 SetPaypalIdNil sets the value for PaypalId to be an explicit nil

### UnsetPaypalId
`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) UnsetPaypalId()`

UnsetPaypalId ensures that no value is present for PaypalId, not even an explicit nil
### GetStatus

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetApprovalUrl

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetApprovalUrl() interface{}`

GetApprovalUrl returns the ApprovalUrl field if non-nil, zero value otherwise.

### GetApprovalUrlOk

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetApprovalUrlOk() (*interface{}, bool)`

GetApprovalUrlOk returns a tuple with the ApprovalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalUrl

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetApprovalUrl(v interface{})`

SetApprovalUrl sets ApprovalUrl field to given value.

### HasApprovalUrl

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) HasApprovalUrl() bool`

HasApprovalUrl returns a boolean if a field has been set.

### SetApprovalUrlNil

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetApprovalUrlNil(b bool)`

 SetApprovalUrlNil sets the value for ApprovalUrl to be an explicit nil

### UnsetApprovalUrl
`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) UnsetApprovalUrl()`

UnsetApprovalUrl ensures that no value is present for ApprovalUrl, not even an explicit nil
### GetMismatchedAmounts

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetMismatchedAmounts() interface{}`

GetMismatchedAmounts returns the MismatchedAmounts field if non-nil, zero value otherwise.

### GetMismatchedAmountsOk

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetMismatchedAmountsOk() (*interface{}, bool)`

GetMismatchedAmountsOk returns a tuple with the MismatchedAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMismatchedAmounts

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetMismatchedAmounts(v interface{})`

SetMismatchedAmounts sets MismatchedAmounts field to given value.

### HasMismatchedAmounts

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) HasMismatchedAmounts() bool`

HasMismatchedAmounts returns a boolean if a field has been set.

### SetMismatchedAmountsNil

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetMismatchedAmountsNil(b bool)`

 SetMismatchedAmountsNil sets the value for MismatchedAmounts to be an explicit nil

### UnsetMismatchedAmounts
`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) UnsetMismatchedAmounts()`

UnsetMismatchedAmounts ensures that no value is present for MismatchedAmounts, not even an explicit nil
### GetPaymentInstrument

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetPaymentInstrument() interface{}`

GetPaymentInstrument returns the PaymentInstrument field if non-nil, zero value otherwise.

### GetPaymentInstrumentOk

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetPaymentInstrumentOk() (*interface{}, bool)`

GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrument

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetPaymentInstrument(v interface{})`

SetPaymentInstrument sets PaymentInstrument field to given value.

### HasPaymentInstrument

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) HasPaymentInstrument() bool`

HasPaymentInstrument returns a boolean if a field has been set.

### SetPaymentInstrumentNil

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetPaymentInstrumentNil(b bool)`

 SetPaymentInstrumentNil sets the value for PaymentInstrument to be an explicit nil

### UnsetPaymentInstrument
`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) UnsetPaymentInstrument()`

UnsetPaymentInstrument ensures that no value is present for PaymentInstrument, not even an explicit nil
### GetCreatedAt

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


