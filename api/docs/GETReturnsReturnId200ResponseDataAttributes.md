# GETReturnsReturnId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **interface{}** | Unique identifier for the return. | [optional] 
**Status** | Pointer to **interface{}** | The return status. One of &#39;draft&#39; (default), &#39;requested&#39;, &#39;approved&#39;, &#39;cancelled&#39;, &#39;shipped&#39;, &#39;rejected&#39;, &#39;received&#39;, or &#39;refunded&#39;. | [optional] 
**CustomerEmail** | Pointer to **interface{}** | The email address of the associated customer. | [optional] 
**SkusCount** | Pointer to **interface{}** | The total number of SKUs in the return&#39;s line items. This can be useful to display a preview of the return content. | [optional] 
**ApprovedAt** | Pointer to **interface{}** | Time at which the return was approved. | [optional] 
**CancelledAt** | Pointer to **interface{}** | Time at which the return was cancelled. | [optional] 
**ShippedAt** | Pointer to **interface{}** | Time at which the return was shipped. | [optional] 
**RejectedAt** | Pointer to **interface{}** | Time at which the return was rejected. | [optional] 
**ReceivedAt** | Pointer to **interface{}** | Time at which the return was received. | [optional] 
**RefundedAt** | Pointer to **interface{}** | Time at which the return was refunded. | [optional] 
**ArchivedAt** | Pointer to **interface{}** | Time at which the resource has been archived. | [optional] 
**EstimatedRefundAmountCents** | Pointer to **interface{}** | The amount to be refunded, estimated by associated return line items, in cents. | [optional] 
**EstimatedRefundAmountFloat** | Pointer to **interface{}** | The amount to be refunded, estimated by associated return line items, float. | [optional] 
**FormattedEstimatedRefundAmount** | Pointer to **interface{}** | The amount to be refunded, estimated by associated return line items, formatted. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETReturnsReturnId200ResponseDataAttributes

`func NewGETReturnsReturnId200ResponseDataAttributes() *GETReturnsReturnId200ResponseDataAttributes`

NewGETReturnsReturnId200ResponseDataAttributes instantiates a new GETReturnsReturnId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETReturnsReturnId200ResponseDataAttributesWithDefaults

`func NewGETReturnsReturnId200ResponseDataAttributesWithDefaults() *GETReturnsReturnId200ResponseDataAttributes`

NewGETReturnsReturnId200ResponseDataAttributesWithDefaults instantiates a new GETReturnsReturnId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetNumber() interface{}`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetNumberOk() (*interface{}, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetNumber(v interface{})`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GETReturnsReturnId200ResponseDataAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *GETReturnsReturnId200ResponseDataAttributes) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetStatus

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETReturnsReturnId200ResponseDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GETReturnsReturnId200ResponseDataAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetCustomerEmail

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetCustomerEmail() interface{}`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetCustomerEmailOk() (*interface{}, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetCustomerEmail(v interface{})`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *GETReturnsReturnId200ResponseDataAttributes) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### SetCustomerEmailNil

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetCustomerEmailNil(b bool)`

 SetCustomerEmailNil sets the value for CustomerEmail to be an explicit nil

### UnsetCustomerEmail
`func (o *GETReturnsReturnId200ResponseDataAttributes) UnsetCustomerEmail()`

UnsetCustomerEmail ensures that no value is present for CustomerEmail, not even an explicit nil
### GetSkusCount

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetSkusCount() interface{}`

GetSkusCount returns the SkusCount field if non-nil, zero value otherwise.

### GetSkusCountOk

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetSkusCountOk() (*interface{}, bool)`

GetSkusCountOk returns a tuple with the SkusCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkusCount

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetSkusCount(v interface{})`

SetSkusCount sets SkusCount field to given value.

### HasSkusCount

`func (o *GETReturnsReturnId200ResponseDataAttributes) HasSkusCount() bool`

HasSkusCount returns a boolean if a field has been set.

### SetSkusCountNil

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetSkusCountNil(b bool)`

 SetSkusCountNil sets the value for SkusCount to be an explicit nil

### UnsetSkusCount
`func (o *GETReturnsReturnId200ResponseDataAttributes) UnsetSkusCount()`

UnsetSkusCount ensures that no value is present for SkusCount, not even an explicit nil
### GetApprovedAt

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetApprovedAt() interface{}`

GetApprovedAt returns the ApprovedAt field if non-nil, zero value otherwise.

### GetApprovedAtOk

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetApprovedAtOk() (*interface{}, bool)`

GetApprovedAtOk returns a tuple with the ApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAt

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetApprovedAt(v interface{})`

SetApprovedAt sets ApprovedAt field to given value.

### HasApprovedAt

`func (o *GETReturnsReturnId200ResponseDataAttributes) HasApprovedAt() bool`

HasApprovedAt returns a boolean if a field has been set.

### SetApprovedAtNil

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetApprovedAtNil(b bool)`

 SetApprovedAtNil sets the value for ApprovedAt to be an explicit nil

### UnsetApprovedAt
`func (o *GETReturnsReturnId200ResponseDataAttributes) UnsetApprovedAt()`

UnsetApprovedAt ensures that no value is present for ApprovedAt, not even an explicit nil
### GetCancelledAt

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetCancelledAt() interface{}`

GetCancelledAt returns the CancelledAt field if non-nil, zero value otherwise.

### GetCancelledAtOk

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetCancelledAtOk() (*interface{}, bool)`

GetCancelledAtOk returns a tuple with the CancelledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAt

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetCancelledAt(v interface{})`

SetCancelledAt sets CancelledAt field to given value.

### HasCancelledAt

`func (o *GETReturnsReturnId200ResponseDataAttributes) HasCancelledAt() bool`

HasCancelledAt returns a boolean if a field has been set.

### SetCancelledAtNil

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetCancelledAtNil(b bool)`

 SetCancelledAtNil sets the value for CancelledAt to be an explicit nil

### UnsetCancelledAt
`func (o *GETReturnsReturnId200ResponseDataAttributes) UnsetCancelledAt()`

UnsetCancelledAt ensures that no value is present for CancelledAt, not even an explicit nil
### GetShippedAt

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetShippedAt() interface{}`

GetShippedAt returns the ShippedAt field if non-nil, zero value otherwise.

### GetShippedAtOk

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetShippedAtOk() (*interface{}, bool)`

GetShippedAtOk returns a tuple with the ShippedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedAt

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetShippedAt(v interface{})`

SetShippedAt sets ShippedAt field to given value.

### HasShippedAt

`func (o *GETReturnsReturnId200ResponseDataAttributes) HasShippedAt() bool`

HasShippedAt returns a boolean if a field has been set.

### SetShippedAtNil

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetShippedAtNil(b bool)`

 SetShippedAtNil sets the value for ShippedAt to be an explicit nil

### UnsetShippedAt
`func (o *GETReturnsReturnId200ResponseDataAttributes) UnsetShippedAt()`

UnsetShippedAt ensures that no value is present for ShippedAt, not even an explicit nil
### GetRejectedAt

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetRejectedAt() interface{}`

GetRejectedAt returns the RejectedAt field if non-nil, zero value otherwise.

### GetRejectedAtOk

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetRejectedAtOk() (*interface{}, bool)`

GetRejectedAtOk returns a tuple with the RejectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedAt

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetRejectedAt(v interface{})`

SetRejectedAt sets RejectedAt field to given value.

### HasRejectedAt

`func (o *GETReturnsReturnId200ResponseDataAttributes) HasRejectedAt() bool`

HasRejectedAt returns a boolean if a field has been set.

### SetRejectedAtNil

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetRejectedAtNil(b bool)`

 SetRejectedAtNil sets the value for RejectedAt to be an explicit nil

### UnsetRejectedAt
`func (o *GETReturnsReturnId200ResponseDataAttributes) UnsetRejectedAt()`

UnsetRejectedAt ensures that no value is present for RejectedAt, not even an explicit nil
### GetReceivedAt

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetReceivedAt() interface{}`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetReceivedAtOk() (*interface{}, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetReceivedAt(v interface{})`

SetReceivedAt sets ReceivedAt field to given value.

### HasReceivedAt

`func (o *GETReturnsReturnId200ResponseDataAttributes) HasReceivedAt() bool`

HasReceivedAt returns a boolean if a field has been set.

### SetReceivedAtNil

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetReceivedAtNil(b bool)`

 SetReceivedAtNil sets the value for ReceivedAt to be an explicit nil

### UnsetReceivedAt
`func (o *GETReturnsReturnId200ResponseDataAttributes) UnsetReceivedAt()`

UnsetReceivedAt ensures that no value is present for ReceivedAt, not even an explicit nil
### GetRefundedAt

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetRefundedAt() interface{}`

GetRefundedAt returns the RefundedAt field if non-nil, zero value otherwise.

### GetRefundedAtOk

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetRefundedAtOk() (*interface{}, bool)`

GetRefundedAtOk returns a tuple with the RefundedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedAt

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetRefundedAt(v interface{})`

SetRefundedAt sets RefundedAt field to given value.

### HasRefundedAt

`func (o *GETReturnsReturnId200ResponseDataAttributes) HasRefundedAt() bool`

HasRefundedAt returns a boolean if a field has been set.

### SetRefundedAtNil

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetRefundedAtNil(b bool)`

 SetRefundedAtNil sets the value for RefundedAt to be an explicit nil

### UnsetRefundedAt
`func (o *GETReturnsReturnId200ResponseDataAttributes) UnsetRefundedAt()`

UnsetRefundedAt ensures that no value is present for RefundedAt, not even an explicit nil
### GetArchivedAt

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetArchivedAt() interface{}`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetArchivedAtOk() (*interface{}, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetArchivedAt(v interface{})`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *GETReturnsReturnId200ResponseDataAttributes) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### SetArchivedAtNil

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetArchivedAtNil(b bool)`

 SetArchivedAtNil sets the value for ArchivedAt to be an explicit nil

### UnsetArchivedAt
`func (o *GETReturnsReturnId200ResponseDataAttributes) UnsetArchivedAt()`

UnsetArchivedAt ensures that no value is present for ArchivedAt, not even an explicit nil
### GetEstimatedRefundAmountCents

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetEstimatedRefundAmountCents() interface{}`

GetEstimatedRefundAmountCents returns the EstimatedRefundAmountCents field if non-nil, zero value otherwise.

### GetEstimatedRefundAmountCentsOk

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetEstimatedRefundAmountCentsOk() (*interface{}, bool)`

GetEstimatedRefundAmountCentsOk returns a tuple with the EstimatedRefundAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedRefundAmountCents

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetEstimatedRefundAmountCents(v interface{})`

SetEstimatedRefundAmountCents sets EstimatedRefundAmountCents field to given value.

### HasEstimatedRefundAmountCents

`func (o *GETReturnsReturnId200ResponseDataAttributes) HasEstimatedRefundAmountCents() bool`

HasEstimatedRefundAmountCents returns a boolean if a field has been set.

### SetEstimatedRefundAmountCentsNil

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetEstimatedRefundAmountCentsNil(b bool)`

 SetEstimatedRefundAmountCentsNil sets the value for EstimatedRefundAmountCents to be an explicit nil

### UnsetEstimatedRefundAmountCents
`func (o *GETReturnsReturnId200ResponseDataAttributes) UnsetEstimatedRefundAmountCents()`

UnsetEstimatedRefundAmountCents ensures that no value is present for EstimatedRefundAmountCents, not even an explicit nil
### GetEstimatedRefundAmountFloat

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetEstimatedRefundAmountFloat() interface{}`

GetEstimatedRefundAmountFloat returns the EstimatedRefundAmountFloat field if non-nil, zero value otherwise.

### GetEstimatedRefundAmountFloatOk

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetEstimatedRefundAmountFloatOk() (*interface{}, bool)`

GetEstimatedRefundAmountFloatOk returns a tuple with the EstimatedRefundAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedRefundAmountFloat

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetEstimatedRefundAmountFloat(v interface{})`

SetEstimatedRefundAmountFloat sets EstimatedRefundAmountFloat field to given value.

### HasEstimatedRefundAmountFloat

`func (o *GETReturnsReturnId200ResponseDataAttributes) HasEstimatedRefundAmountFloat() bool`

HasEstimatedRefundAmountFloat returns a boolean if a field has been set.

### SetEstimatedRefundAmountFloatNil

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetEstimatedRefundAmountFloatNil(b bool)`

 SetEstimatedRefundAmountFloatNil sets the value for EstimatedRefundAmountFloat to be an explicit nil

### UnsetEstimatedRefundAmountFloat
`func (o *GETReturnsReturnId200ResponseDataAttributes) UnsetEstimatedRefundAmountFloat()`

UnsetEstimatedRefundAmountFloat ensures that no value is present for EstimatedRefundAmountFloat, not even an explicit nil
### GetFormattedEstimatedRefundAmount

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetFormattedEstimatedRefundAmount() interface{}`

GetFormattedEstimatedRefundAmount returns the FormattedEstimatedRefundAmount field if non-nil, zero value otherwise.

### GetFormattedEstimatedRefundAmountOk

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetFormattedEstimatedRefundAmountOk() (*interface{}, bool)`

GetFormattedEstimatedRefundAmountOk returns a tuple with the FormattedEstimatedRefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedEstimatedRefundAmount

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetFormattedEstimatedRefundAmount(v interface{})`

SetFormattedEstimatedRefundAmount sets FormattedEstimatedRefundAmount field to given value.

### HasFormattedEstimatedRefundAmount

`func (o *GETReturnsReturnId200ResponseDataAttributes) HasFormattedEstimatedRefundAmount() bool`

HasFormattedEstimatedRefundAmount returns a boolean if a field has been set.

### SetFormattedEstimatedRefundAmountNil

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetFormattedEstimatedRefundAmountNil(b bool)`

 SetFormattedEstimatedRefundAmountNil sets the value for FormattedEstimatedRefundAmount to be an explicit nil

### UnsetFormattedEstimatedRefundAmount
`func (o *GETReturnsReturnId200ResponseDataAttributes) UnsetFormattedEstimatedRefundAmount()`

UnsetFormattedEstimatedRefundAmount ensures that no value is present for FormattedEstimatedRefundAmount, not even an explicit nil
### GetCreatedAt

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETReturnsReturnId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETReturnsReturnId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETReturnsReturnId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETReturnsReturnId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETReturnsReturnId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETReturnsReturnId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETReturnsReturnId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETReturnsReturnId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETReturnsReturnId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETReturnsReturnId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETReturnsReturnId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETReturnsReturnId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


