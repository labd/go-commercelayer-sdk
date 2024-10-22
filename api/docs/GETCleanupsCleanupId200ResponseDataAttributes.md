# GETCleanupsCleanupId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | Pointer to **interface{}** | The type of resource being cleaned. | [optional] 
**Status** | Pointer to **interface{}** | The cleanup job status. One of &#39;pending&#39; (default), &#39;in_progress&#39;, &#39;interrupted&#39;, or &#39;completed&#39;. | [optional] 
**StartedAt** | Pointer to **interface{}** | Time at which the cleanup was started. | [optional] 
**CompletedAt** | Pointer to **interface{}** | Time at which the cleanup was completed. | [optional] 
**InterruptedAt** | Pointer to **interface{}** | Time at which the cleanup was interrupted. | [optional] 
**Filters** | Pointer to **interface{}** | The filters used to select the records to be cleaned. | [optional] 
**RecordsCount** | Pointer to **interface{}** | Indicates the number of records to be cleaned. | [optional] 
**ErrorsCount** | Pointer to **interface{}** | Indicates the number of cleanup errors, if any. | [optional] 
**ProcessedCount** | Pointer to **interface{}** | Indicates the number of records that have been cleaned. | [optional] 
**ErrorsLog** | Pointer to **interface{}** | Contains the cleanup errors, if any. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETCleanupsCleanupId200ResponseDataAttributes

`func NewGETCleanupsCleanupId200ResponseDataAttributes() *GETCleanupsCleanupId200ResponseDataAttributes`

NewGETCleanupsCleanupId200ResponseDataAttributes instantiates a new GETCleanupsCleanupId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETCleanupsCleanupId200ResponseDataAttributesWithDefaults

`func NewGETCleanupsCleanupId200ResponseDataAttributesWithDefaults() *GETCleanupsCleanupId200ResponseDataAttributes`

NewGETCleanupsCleanupId200ResponseDataAttributesWithDefaults instantiates a new GETCleanupsCleanupId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetResourceType() interface{}`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetResourceTypeOk() (*interface{}, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetResourceType(v interface{})`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *GETCleanupsCleanupId200ResponseDataAttributes) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetStatus

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GETCleanupsCleanupId200ResponseDataAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStartedAt

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetStartedAt() interface{}`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetStartedAtOk() (*interface{}, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetStartedAt(v interface{})`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *GETCleanupsCleanupId200ResponseDataAttributes) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetCompletedAt() interface{}`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetCompletedAtOk() (*interface{}, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetCompletedAt(v interface{})`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *GETCleanupsCleanupId200ResponseDataAttributes) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetInterruptedAt

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetInterruptedAt() interface{}`

GetInterruptedAt returns the InterruptedAt field if non-nil, zero value otherwise.

### GetInterruptedAtOk

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetInterruptedAtOk() (*interface{}, bool)`

GetInterruptedAtOk returns a tuple with the InterruptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptedAt

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetInterruptedAt(v interface{})`

SetInterruptedAt sets InterruptedAt field to given value.

### HasInterruptedAt

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) HasInterruptedAt() bool`

HasInterruptedAt returns a boolean if a field has been set.

### SetInterruptedAtNil

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetInterruptedAtNil(b bool)`

 SetInterruptedAtNil sets the value for InterruptedAt to be an explicit nil

### UnsetInterruptedAt
`func (o *GETCleanupsCleanupId200ResponseDataAttributes) UnsetInterruptedAt()`

UnsetInterruptedAt ensures that no value is present for InterruptedAt, not even an explicit nil
### GetFilters

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetFilters() interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetFiltersOk() (*interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetFilters(v interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *GETCleanupsCleanupId200ResponseDataAttributes) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetRecordsCount

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetRecordsCount() interface{}`

GetRecordsCount returns the RecordsCount field if non-nil, zero value otherwise.

### GetRecordsCountOk

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetRecordsCountOk() (*interface{}, bool)`

GetRecordsCountOk returns a tuple with the RecordsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsCount

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetRecordsCount(v interface{})`

SetRecordsCount sets RecordsCount field to given value.

### HasRecordsCount

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) HasRecordsCount() bool`

HasRecordsCount returns a boolean if a field has been set.

### SetRecordsCountNil

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetRecordsCountNil(b bool)`

 SetRecordsCountNil sets the value for RecordsCount to be an explicit nil

### UnsetRecordsCount
`func (o *GETCleanupsCleanupId200ResponseDataAttributes) UnsetRecordsCount()`

UnsetRecordsCount ensures that no value is present for RecordsCount, not even an explicit nil
### GetErrorsCount

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetErrorsCount() interface{}`

GetErrorsCount returns the ErrorsCount field if non-nil, zero value otherwise.

### GetErrorsCountOk

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetErrorsCountOk() (*interface{}, bool)`

GetErrorsCountOk returns a tuple with the ErrorsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsCount

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetErrorsCount(v interface{})`

SetErrorsCount sets ErrorsCount field to given value.

### HasErrorsCount

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) HasErrorsCount() bool`

HasErrorsCount returns a boolean if a field has been set.

### SetErrorsCountNil

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetErrorsCountNil(b bool)`

 SetErrorsCountNil sets the value for ErrorsCount to be an explicit nil

### UnsetErrorsCount
`func (o *GETCleanupsCleanupId200ResponseDataAttributes) UnsetErrorsCount()`

UnsetErrorsCount ensures that no value is present for ErrorsCount, not even an explicit nil
### GetProcessedCount

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetProcessedCount() interface{}`

GetProcessedCount returns the ProcessedCount field if non-nil, zero value otherwise.

### GetProcessedCountOk

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetProcessedCountOk() (*interface{}, bool)`

GetProcessedCountOk returns a tuple with the ProcessedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedCount

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetProcessedCount(v interface{})`

SetProcessedCount sets ProcessedCount field to given value.

### HasProcessedCount

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) HasProcessedCount() bool`

HasProcessedCount returns a boolean if a field has been set.

### SetProcessedCountNil

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetProcessedCountNil(b bool)`

 SetProcessedCountNil sets the value for ProcessedCount to be an explicit nil

### UnsetProcessedCount
`func (o *GETCleanupsCleanupId200ResponseDataAttributes) UnsetProcessedCount()`

UnsetProcessedCount ensures that no value is present for ProcessedCount, not even an explicit nil
### GetErrorsLog

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetErrorsLog() interface{}`

GetErrorsLog returns the ErrorsLog field if non-nil, zero value otherwise.

### GetErrorsLogOk

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetErrorsLogOk() (*interface{}, bool)`

GetErrorsLogOk returns a tuple with the ErrorsLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsLog

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetErrorsLog(v interface{})`

SetErrorsLog sets ErrorsLog field to given value.

### HasErrorsLog

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) HasErrorsLog() bool`

HasErrorsLog returns a boolean if a field has been set.

### SetErrorsLogNil

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetErrorsLogNil(b bool)`

 SetErrorsLogNil sets the value for ErrorsLog to be an explicit nil

### UnsetErrorsLog
`func (o *GETCleanupsCleanupId200ResponseDataAttributes) UnsetErrorsLog()`

UnsetErrorsLog ensures that no value is present for ErrorsLog, not even an explicit nil
### GetCreatedAt

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETCleanupsCleanupId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETCleanupsCleanupId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETCleanupsCleanupId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETCleanupsCleanupId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETCleanupsCleanupId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETCleanupsCleanupId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


