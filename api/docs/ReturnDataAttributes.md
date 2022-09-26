# ReturnDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **string** | Unique identifier for the return | [optional] 
**Status** | Pointer to **string** | The return status, one of &#39;draft&#39;, &#39;requested&#39;, &#39;approved&#39;, &#39;cancelled&#39;, &#39;shipped&#39;, &#39;rejected&#39; or &#39;received&#39; | [optional] 
**CustomerEmail** | Pointer to **string** | The email address of the associated customer. | [optional] 
**SkusCount** | Pointer to **int32** | The total number of SKUs in the return&#39;s line items. This can be useful to display a preview of the return content. | [optional] 
**ApprovedAt** | Pointer to **string** | Time at which the return was approved. | [optional] 
**CancelledAt** | Pointer to **string** | Time at which the return was cancelled. | [optional] 
**ShippedAt** | Pointer to **string** | Time at which the return was shipped. | [optional] 
**RejectedAt** | Pointer to **string** | Time at which the return was rejected. | [optional] 
**ReceivedAt** | Pointer to **string** | Time at which the return was received. | [optional] 
**ArchivedAt** | Pointer to **string** | Time at which the resource has been archived. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewReturnDataAttributes

`func NewReturnDataAttributes() *ReturnDataAttributes`

NewReturnDataAttributes instantiates a new ReturnDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnDataAttributesWithDefaults

`func NewReturnDataAttributesWithDefaults() *ReturnDataAttributes`

NewReturnDataAttributesWithDefaults instantiates a new ReturnDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *ReturnDataAttributes) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ReturnDataAttributes) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ReturnDataAttributes) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *ReturnDataAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetStatus

`func (o *ReturnDataAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReturnDataAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReturnDataAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReturnDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *ReturnDataAttributes) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *ReturnDataAttributes) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *ReturnDataAttributes) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *ReturnDataAttributes) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetSkusCount

`func (o *ReturnDataAttributes) GetSkusCount() int32`

GetSkusCount returns the SkusCount field if non-nil, zero value otherwise.

### GetSkusCountOk

`func (o *ReturnDataAttributes) GetSkusCountOk() (*int32, bool)`

GetSkusCountOk returns a tuple with the SkusCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkusCount

`func (o *ReturnDataAttributes) SetSkusCount(v int32)`

SetSkusCount sets SkusCount field to given value.

### HasSkusCount

`func (o *ReturnDataAttributes) HasSkusCount() bool`

HasSkusCount returns a boolean if a field has been set.

### GetApprovedAt

`func (o *ReturnDataAttributes) GetApprovedAt() string`

GetApprovedAt returns the ApprovedAt field if non-nil, zero value otherwise.

### GetApprovedAtOk

`func (o *ReturnDataAttributes) GetApprovedAtOk() (*string, bool)`

GetApprovedAtOk returns a tuple with the ApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAt

`func (o *ReturnDataAttributes) SetApprovedAt(v string)`

SetApprovedAt sets ApprovedAt field to given value.

### HasApprovedAt

`func (o *ReturnDataAttributes) HasApprovedAt() bool`

HasApprovedAt returns a boolean if a field has been set.

### GetCancelledAt

`func (o *ReturnDataAttributes) GetCancelledAt() string`

GetCancelledAt returns the CancelledAt field if non-nil, zero value otherwise.

### GetCancelledAtOk

`func (o *ReturnDataAttributes) GetCancelledAtOk() (*string, bool)`

GetCancelledAtOk returns a tuple with the CancelledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAt

`func (o *ReturnDataAttributes) SetCancelledAt(v string)`

SetCancelledAt sets CancelledAt field to given value.

### HasCancelledAt

`func (o *ReturnDataAttributes) HasCancelledAt() bool`

HasCancelledAt returns a boolean if a field has been set.

### GetShippedAt

`func (o *ReturnDataAttributes) GetShippedAt() string`

GetShippedAt returns the ShippedAt field if non-nil, zero value otherwise.

### GetShippedAtOk

`func (o *ReturnDataAttributes) GetShippedAtOk() (*string, bool)`

GetShippedAtOk returns a tuple with the ShippedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedAt

`func (o *ReturnDataAttributes) SetShippedAt(v string)`

SetShippedAt sets ShippedAt field to given value.

### HasShippedAt

`func (o *ReturnDataAttributes) HasShippedAt() bool`

HasShippedAt returns a boolean if a field has been set.

### GetRejectedAt

`func (o *ReturnDataAttributes) GetRejectedAt() string`

GetRejectedAt returns the RejectedAt field if non-nil, zero value otherwise.

### GetRejectedAtOk

`func (o *ReturnDataAttributes) GetRejectedAtOk() (*string, bool)`

GetRejectedAtOk returns a tuple with the RejectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedAt

`func (o *ReturnDataAttributes) SetRejectedAt(v string)`

SetRejectedAt sets RejectedAt field to given value.

### HasRejectedAt

`func (o *ReturnDataAttributes) HasRejectedAt() bool`

HasRejectedAt returns a boolean if a field has been set.

### GetReceivedAt

`func (o *ReturnDataAttributes) GetReceivedAt() string`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *ReturnDataAttributes) GetReceivedAtOk() (*string, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *ReturnDataAttributes) SetReceivedAt(v string)`

SetReceivedAt sets ReceivedAt field to given value.

### HasReceivedAt

`func (o *ReturnDataAttributes) HasReceivedAt() bool`

HasReceivedAt returns a boolean if a field has been set.

### GetArchivedAt

`func (o *ReturnDataAttributes) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *ReturnDataAttributes) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *ReturnDataAttributes) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *ReturnDataAttributes) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetId

`func (o *ReturnDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReturnDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReturnDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReturnDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ReturnDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReturnDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReturnDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ReturnDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ReturnDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ReturnDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ReturnDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ReturnDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *ReturnDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ReturnDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ReturnDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ReturnDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *ReturnDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *ReturnDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *ReturnDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *ReturnDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *ReturnDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ReturnDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ReturnDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ReturnDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


