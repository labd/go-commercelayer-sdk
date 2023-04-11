# GETReturns200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **interface{}** | Unique identifier for the return | [optional] 
**Status** | Pointer to **interface{}** | The return status, one of &#39;draft&#39;, &#39;requested&#39;, &#39;approved&#39;, &#39;cancelled&#39;, &#39;shipped&#39;, &#39;rejected&#39; or &#39;received&#39; | [optional] 
**CustomerEmail** | Pointer to **interface{}** | The email address of the associated customer. | [optional] 
**SkusCount** | Pointer to **interface{}** | The total number of SKUs in the return&#39;s line items. This can be useful to display a preview of the return content. | [optional] 
**ApprovedAt** | Pointer to **interface{}** | Time at which the return was approved. | [optional] 
**CancelledAt** | Pointer to **interface{}** | Time at which the return was cancelled. | [optional] 
**ShippedAt** | Pointer to **interface{}** | Time at which the return was shipped. | [optional] 
**RejectedAt** | Pointer to **interface{}** | Time at which the return was rejected. | [optional] 
**ReceivedAt** | Pointer to **interface{}** | Time at which the return was received. | [optional] 
**ArchivedAt** | Pointer to **interface{}** | Time at which the resource has been archived. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETReturns200ResponseDataInnerAttributes

`func NewGETReturns200ResponseDataInnerAttributes() *GETReturns200ResponseDataInnerAttributes`

NewGETReturns200ResponseDataInnerAttributes instantiates a new GETReturns200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETReturns200ResponseDataInnerAttributesWithDefaults

`func NewGETReturns200ResponseDataInnerAttributesWithDefaults() *GETReturns200ResponseDataInnerAttributes`

NewGETReturns200ResponseDataInnerAttributesWithDefaults instantiates a new GETReturns200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *GETReturns200ResponseDataInnerAttributes) GetNumber() interface{}`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GETReturns200ResponseDataInnerAttributes) GetNumberOk() (*interface{}, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GETReturns200ResponseDataInnerAttributes) SetNumber(v interface{})`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GETReturns200ResponseDataInnerAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *GETReturns200ResponseDataInnerAttributes) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *GETReturns200ResponseDataInnerAttributes) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetStatus

`func (o *GETReturns200ResponseDataInnerAttributes) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETReturns200ResponseDataInnerAttributes) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETReturns200ResponseDataInnerAttributes) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETReturns200ResponseDataInnerAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GETReturns200ResponseDataInnerAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GETReturns200ResponseDataInnerAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetCustomerEmail

`func (o *GETReturns200ResponseDataInnerAttributes) GetCustomerEmail() interface{}`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *GETReturns200ResponseDataInnerAttributes) GetCustomerEmailOk() (*interface{}, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *GETReturns200ResponseDataInnerAttributes) SetCustomerEmail(v interface{})`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *GETReturns200ResponseDataInnerAttributes) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### SetCustomerEmailNil

`func (o *GETReturns200ResponseDataInnerAttributes) SetCustomerEmailNil(b bool)`

 SetCustomerEmailNil sets the value for CustomerEmail to be an explicit nil

### UnsetCustomerEmail
`func (o *GETReturns200ResponseDataInnerAttributes) UnsetCustomerEmail()`

UnsetCustomerEmail ensures that no value is present for CustomerEmail, not even an explicit nil
### GetSkusCount

`func (o *GETReturns200ResponseDataInnerAttributes) GetSkusCount() interface{}`

GetSkusCount returns the SkusCount field if non-nil, zero value otherwise.

### GetSkusCountOk

`func (o *GETReturns200ResponseDataInnerAttributes) GetSkusCountOk() (*interface{}, bool)`

GetSkusCountOk returns a tuple with the SkusCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkusCount

`func (o *GETReturns200ResponseDataInnerAttributes) SetSkusCount(v interface{})`

SetSkusCount sets SkusCount field to given value.

### HasSkusCount

`func (o *GETReturns200ResponseDataInnerAttributes) HasSkusCount() bool`

HasSkusCount returns a boolean if a field has been set.

### SetSkusCountNil

`func (o *GETReturns200ResponseDataInnerAttributes) SetSkusCountNil(b bool)`

 SetSkusCountNil sets the value for SkusCount to be an explicit nil

### UnsetSkusCount
`func (o *GETReturns200ResponseDataInnerAttributes) UnsetSkusCount()`

UnsetSkusCount ensures that no value is present for SkusCount, not even an explicit nil
### GetApprovedAt

`func (o *GETReturns200ResponseDataInnerAttributes) GetApprovedAt() interface{}`

GetApprovedAt returns the ApprovedAt field if non-nil, zero value otherwise.

### GetApprovedAtOk

`func (o *GETReturns200ResponseDataInnerAttributes) GetApprovedAtOk() (*interface{}, bool)`

GetApprovedAtOk returns a tuple with the ApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAt

`func (o *GETReturns200ResponseDataInnerAttributes) SetApprovedAt(v interface{})`

SetApprovedAt sets ApprovedAt field to given value.

### HasApprovedAt

`func (o *GETReturns200ResponseDataInnerAttributes) HasApprovedAt() bool`

HasApprovedAt returns a boolean if a field has been set.

### SetApprovedAtNil

`func (o *GETReturns200ResponseDataInnerAttributes) SetApprovedAtNil(b bool)`

 SetApprovedAtNil sets the value for ApprovedAt to be an explicit nil

### UnsetApprovedAt
`func (o *GETReturns200ResponseDataInnerAttributes) UnsetApprovedAt()`

UnsetApprovedAt ensures that no value is present for ApprovedAt, not even an explicit nil
### GetCancelledAt

`func (o *GETReturns200ResponseDataInnerAttributes) GetCancelledAt() interface{}`

GetCancelledAt returns the CancelledAt field if non-nil, zero value otherwise.

### GetCancelledAtOk

`func (o *GETReturns200ResponseDataInnerAttributes) GetCancelledAtOk() (*interface{}, bool)`

GetCancelledAtOk returns a tuple with the CancelledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAt

`func (o *GETReturns200ResponseDataInnerAttributes) SetCancelledAt(v interface{})`

SetCancelledAt sets CancelledAt field to given value.

### HasCancelledAt

`func (o *GETReturns200ResponseDataInnerAttributes) HasCancelledAt() bool`

HasCancelledAt returns a boolean if a field has been set.

### SetCancelledAtNil

`func (o *GETReturns200ResponseDataInnerAttributes) SetCancelledAtNil(b bool)`

 SetCancelledAtNil sets the value for CancelledAt to be an explicit nil

### UnsetCancelledAt
`func (o *GETReturns200ResponseDataInnerAttributes) UnsetCancelledAt()`

UnsetCancelledAt ensures that no value is present for CancelledAt, not even an explicit nil
### GetShippedAt

`func (o *GETReturns200ResponseDataInnerAttributes) GetShippedAt() interface{}`

GetShippedAt returns the ShippedAt field if non-nil, zero value otherwise.

### GetShippedAtOk

`func (o *GETReturns200ResponseDataInnerAttributes) GetShippedAtOk() (*interface{}, bool)`

GetShippedAtOk returns a tuple with the ShippedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedAt

`func (o *GETReturns200ResponseDataInnerAttributes) SetShippedAt(v interface{})`

SetShippedAt sets ShippedAt field to given value.

### HasShippedAt

`func (o *GETReturns200ResponseDataInnerAttributes) HasShippedAt() bool`

HasShippedAt returns a boolean if a field has been set.

### SetShippedAtNil

`func (o *GETReturns200ResponseDataInnerAttributes) SetShippedAtNil(b bool)`

 SetShippedAtNil sets the value for ShippedAt to be an explicit nil

### UnsetShippedAt
`func (o *GETReturns200ResponseDataInnerAttributes) UnsetShippedAt()`

UnsetShippedAt ensures that no value is present for ShippedAt, not even an explicit nil
### GetRejectedAt

`func (o *GETReturns200ResponseDataInnerAttributes) GetRejectedAt() interface{}`

GetRejectedAt returns the RejectedAt field if non-nil, zero value otherwise.

### GetRejectedAtOk

`func (o *GETReturns200ResponseDataInnerAttributes) GetRejectedAtOk() (*interface{}, bool)`

GetRejectedAtOk returns a tuple with the RejectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedAt

`func (o *GETReturns200ResponseDataInnerAttributes) SetRejectedAt(v interface{})`

SetRejectedAt sets RejectedAt field to given value.

### HasRejectedAt

`func (o *GETReturns200ResponseDataInnerAttributes) HasRejectedAt() bool`

HasRejectedAt returns a boolean if a field has been set.

### SetRejectedAtNil

`func (o *GETReturns200ResponseDataInnerAttributes) SetRejectedAtNil(b bool)`

 SetRejectedAtNil sets the value for RejectedAt to be an explicit nil

### UnsetRejectedAt
`func (o *GETReturns200ResponseDataInnerAttributes) UnsetRejectedAt()`

UnsetRejectedAt ensures that no value is present for RejectedAt, not even an explicit nil
### GetReceivedAt

`func (o *GETReturns200ResponseDataInnerAttributes) GetReceivedAt() interface{}`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *GETReturns200ResponseDataInnerAttributes) GetReceivedAtOk() (*interface{}, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *GETReturns200ResponseDataInnerAttributes) SetReceivedAt(v interface{})`

SetReceivedAt sets ReceivedAt field to given value.

### HasReceivedAt

`func (o *GETReturns200ResponseDataInnerAttributes) HasReceivedAt() bool`

HasReceivedAt returns a boolean if a field has been set.

### SetReceivedAtNil

`func (o *GETReturns200ResponseDataInnerAttributes) SetReceivedAtNil(b bool)`

 SetReceivedAtNil sets the value for ReceivedAt to be an explicit nil

### UnsetReceivedAt
`func (o *GETReturns200ResponseDataInnerAttributes) UnsetReceivedAt()`

UnsetReceivedAt ensures that no value is present for ReceivedAt, not even an explicit nil
### GetArchivedAt

`func (o *GETReturns200ResponseDataInnerAttributes) GetArchivedAt() interface{}`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *GETReturns200ResponseDataInnerAttributes) GetArchivedAtOk() (*interface{}, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *GETReturns200ResponseDataInnerAttributes) SetArchivedAt(v interface{})`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *GETReturns200ResponseDataInnerAttributes) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### SetArchivedAtNil

`func (o *GETReturns200ResponseDataInnerAttributes) SetArchivedAtNil(b bool)`

 SetArchivedAtNil sets the value for ArchivedAt to be an explicit nil

### UnsetArchivedAt
`func (o *GETReturns200ResponseDataInnerAttributes) UnsetArchivedAt()`

UnsetArchivedAt ensures that no value is present for ArchivedAt, not even an explicit nil
### GetCreatedAt

`func (o *GETReturns200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETReturns200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETReturns200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETReturns200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETReturns200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETReturns200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETReturns200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETReturns200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETReturns200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETReturns200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETReturns200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETReturns200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETReturns200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETReturns200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETReturns200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETReturns200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETReturns200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETReturns200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETReturns200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETReturns200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETReturns200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETReturns200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETReturns200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETReturns200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETReturns200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETReturns200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETReturns200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETReturns200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETReturns200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETReturns200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


