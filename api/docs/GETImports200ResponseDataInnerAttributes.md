# GETImports200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | Pointer to **interface{}** | The type of resource being imported. | [optional] 
**ParentResourceId** | Pointer to **interface{}** | The ID of the parent resource to be associated with imported data. | [optional] 
**Status** | Pointer to **interface{}** | The import job status. One of &#39;pending&#39; (default), &#39;in_progress&#39;, &#39;interrupted&#39;, or &#39;completed&#39;. | [optional] 
**StartedAt** | Pointer to **interface{}** | Time at which the import was started. | [optional] 
**CompletedAt** | Pointer to **interface{}** | Time at which the import was completed. | [optional] 
**InterruptedAt** | Pointer to **interface{}** | Time at which the import was interrupted. | [optional] 
**Inputs** | Pointer to **[]interface{}** | Array of objects representing the resources that are being imported. | [optional] 
**InputsSize** | Pointer to **interface{}** | Indicates the size of the objects to be imported. | [optional] 
**ErrorsCount** | Pointer to **interface{}** | Indicates the number of import errors, if any. | [optional] 
**WarningsCount** | Pointer to **interface{}** | Indicates the number of import warnings, if any. | [optional] 
**DestroyedCount** | Pointer to **interface{}** | Indicates the number of records that have been destroyed, if any. | [optional] 
**ProcessedCount** | Pointer to **interface{}** | Indicates the number of records that have been processed (created or updated). | [optional] 
**ErrorsLog** | Pointer to **interface{}** | Contains the import errors, if any. | [optional] 
**WarningsLog** | Pointer to **interface{}** | Contains the import warnings, if any. | [optional] 
**CleanupRecords** | Pointer to **interface{}** | Indicates if the import should cleanup records that are not included in the inputs array. | [optional] 
**AttachmentUrl** | Pointer to **interface{}** | The URL the the raw inputs file, which will be generated at import start. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETImports200ResponseDataInnerAttributes

`func NewGETImports200ResponseDataInnerAttributes() *GETImports200ResponseDataInnerAttributes`

NewGETImports200ResponseDataInnerAttributes instantiates a new GETImports200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETImports200ResponseDataInnerAttributesWithDefaults

`func NewGETImports200ResponseDataInnerAttributesWithDefaults() *GETImports200ResponseDataInnerAttributes`

NewGETImports200ResponseDataInnerAttributesWithDefaults instantiates a new GETImports200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *GETImports200ResponseDataInnerAttributes) GetResourceType() interface{}`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *GETImports200ResponseDataInnerAttributes) GetResourceTypeOk() (*interface{}, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *GETImports200ResponseDataInnerAttributes) SetResourceType(v interface{})`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *GETImports200ResponseDataInnerAttributes) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *GETImports200ResponseDataInnerAttributes) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *GETImports200ResponseDataInnerAttributes) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetParentResourceId

`func (o *GETImports200ResponseDataInnerAttributes) GetParentResourceId() interface{}`

GetParentResourceId returns the ParentResourceId field if non-nil, zero value otherwise.

### GetParentResourceIdOk

`func (o *GETImports200ResponseDataInnerAttributes) GetParentResourceIdOk() (*interface{}, bool)`

GetParentResourceIdOk returns a tuple with the ParentResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentResourceId

`func (o *GETImports200ResponseDataInnerAttributes) SetParentResourceId(v interface{})`

SetParentResourceId sets ParentResourceId field to given value.

### HasParentResourceId

`func (o *GETImports200ResponseDataInnerAttributes) HasParentResourceId() bool`

HasParentResourceId returns a boolean if a field has been set.

### SetParentResourceIdNil

`func (o *GETImports200ResponseDataInnerAttributes) SetParentResourceIdNil(b bool)`

 SetParentResourceIdNil sets the value for ParentResourceId to be an explicit nil

### UnsetParentResourceId
`func (o *GETImports200ResponseDataInnerAttributes) UnsetParentResourceId()`

UnsetParentResourceId ensures that no value is present for ParentResourceId, not even an explicit nil
### GetStatus

`func (o *GETImports200ResponseDataInnerAttributes) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETImports200ResponseDataInnerAttributes) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETImports200ResponseDataInnerAttributes) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETImports200ResponseDataInnerAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GETImports200ResponseDataInnerAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GETImports200ResponseDataInnerAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStartedAt

`func (o *GETImports200ResponseDataInnerAttributes) GetStartedAt() interface{}`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GETImports200ResponseDataInnerAttributes) GetStartedAtOk() (*interface{}, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GETImports200ResponseDataInnerAttributes) SetStartedAt(v interface{})`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *GETImports200ResponseDataInnerAttributes) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *GETImports200ResponseDataInnerAttributes) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *GETImports200ResponseDataInnerAttributes) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *GETImports200ResponseDataInnerAttributes) GetCompletedAt() interface{}`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GETImports200ResponseDataInnerAttributes) GetCompletedAtOk() (*interface{}, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GETImports200ResponseDataInnerAttributes) SetCompletedAt(v interface{})`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GETImports200ResponseDataInnerAttributes) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *GETImports200ResponseDataInnerAttributes) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *GETImports200ResponseDataInnerAttributes) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetInterruptedAt

`func (o *GETImports200ResponseDataInnerAttributes) GetInterruptedAt() interface{}`

GetInterruptedAt returns the InterruptedAt field if non-nil, zero value otherwise.

### GetInterruptedAtOk

`func (o *GETImports200ResponseDataInnerAttributes) GetInterruptedAtOk() (*interface{}, bool)`

GetInterruptedAtOk returns a tuple with the InterruptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptedAt

`func (o *GETImports200ResponseDataInnerAttributes) SetInterruptedAt(v interface{})`

SetInterruptedAt sets InterruptedAt field to given value.

### HasInterruptedAt

`func (o *GETImports200ResponseDataInnerAttributes) HasInterruptedAt() bool`

HasInterruptedAt returns a boolean if a field has been set.

### SetInterruptedAtNil

`func (o *GETImports200ResponseDataInnerAttributes) SetInterruptedAtNil(b bool)`

 SetInterruptedAtNil sets the value for InterruptedAt to be an explicit nil

### UnsetInterruptedAt
`func (o *GETImports200ResponseDataInnerAttributes) UnsetInterruptedAt()`

UnsetInterruptedAt ensures that no value is present for InterruptedAt, not even an explicit nil
### GetInputs

`func (o *GETImports200ResponseDataInnerAttributes) GetInputs() []interface{}`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *GETImports200ResponseDataInnerAttributes) GetInputsOk() (*[]interface{}, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *GETImports200ResponseDataInnerAttributes) SetInputs(v []interface{})`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *GETImports200ResponseDataInnerAttributes) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetInputsSize

`func (o *GETImports200ResponseDataInnerAttributes) GetInputsSize() interface{}`

GetInputsSize returns the InputsSize field if non-nil, zero value otherwise.

### GetInputsSizeOk

`func (o *GETImports200ResponseDataInnerAttributes) GetInputsSizeOk() (*interface{}, bool)`

GetInputsSizeOk returns a tuple with the InputsSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputsSize

`func (o *GETImports200ResponseDataInnerAttributes) SetInputsSize(v interface{})`

SetInputsSize sets InputsSize field to given value.

### HasInputsSize

`func (o *GETImports200ResponseDataInnerAttributes) HasInputsSize() bool`

HasInputsSize returns a boolean if a field has been set.

### SetInputsSizeNil

`func (o *GETImports200ResponseDataInnerAttributes) SetInputsSizeNil(b bool)`

 SetInputsSizeNil sets the value for InputsSize to be an explicit nil

### UnsetInputsSize
`func (o *GETImports200ResponseDataInnerAttributes) UnsetInputsSize()`

UnsetInputsSize ensures that no value is present for InputsSize, not even an explicit nil
### GetErrorsCount

`func (o *GETImports200ResponseDataInnerAttributes) GetErrorsCount() interface{}`

GetErrorsCount returns the ErrorsCount field if non-nil, zero value otherwise.

### GetErrorsCountOk

`func (o *GETImports200ResponseDataInnerAttributes) GetErrorsCountOk() (*interface{}, bool)`

GetErrorsCountOk returns a tuple with the ErrorsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsCount

`func (o *GETImports200ResponseDataInnerAttributes) SetErrorsCount(v interface{})`

SetErrorsCount sets ErrorsCount field to given value.

### HasErrorsCount

`func (o *GETImports200ResponseDataInnerAttributes) HasErrorsCount() bool`

HasErrorsCount returns a boolean if a field has been set.

### SetErrorsCountNil

`func (o *GETImports200ResponseDataInnerAttributes) SetErrorsCountNil(b bool)`

 SetErrorsCountNil sets the value for ErrorsCount to be an explicit nil

### UnsetErrorsCount
`func (o *GETImports200ResponseDataInnerAttributes) UnsetErrorsCount()`

UnsetErrorsCount ensures that no value is present for ErrorsCount, not even an explicit nil
### GetWarningsCount

`func (o *GETImports200ResponseDataInnerAttributes) GetWarningsCount() interface{}`

GetWarningsCount returns the WarningsCount field if non-nil, zero value otherwise.

### GetWarningsCountOk

`func (o *GETImports200ResponseDataInnerAttributes) GetWarningsCountOk() (*interface{}, bool)`

GetWarningsCountOk returns a tuple with the WarningsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningsCount

`func (o *GETImports200ResponseDataInnerAttributes) SetWarningsCount(v interface{})`

SetWarningsCount sets WarningsCount field to given value.

### HasWarningsCount

`func (o *GETImports200ResponseDataInnerAttributes) HasWarningsCount() bool`

HasWarningsCount returns a boolean if a field has been set.

### SetWarningsCountNil

`func (o *GETImports200ResponseDataInnerAttributes) SetWarningsCountNil(b bool)`

 SetWarningsCountNil sets the value for WarningsCount to be an explicit nil

### UnsetWarningsCount
`func (o *GETImports200ResponseDataInnerAttributes) UnsetWarningsCount()`

UnsetWarningsCount ensures that no value is present for WarningsCount, not even an explicit nil
### GetDestroyedCount

`func (o *GETImports200ResponseDataInnerAttributes) GetDestroyedCount() interface{}`

GetDestroyedCount returns the DestroyedCount field if non-nil, zero value otherwise.

### GetDestroyedCountOk

`func (o *GETImports200ResponseDataInnerAttributes) GetDestroyedCountOk() (*interface{}, bool)`

GetDestroyedCountOk returns a tuple with the DestroyedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyedCount

`func (o *GETImports200ResponseDataInnerAttributes) SetDestroyedCount(v interface{})`

SetDestroyedCount sets DestroyedCount field to given value.

### HasDestroyedCount

`func (o *GETImports200ResponseDataInnerAttributes) HasDestroyedCount() bool`

HasDestroyedCount returns a boolean if a field has been set.

### SetDestroyedCountNil

`func (o *GETImports200ResponseDataInnerAttributes) SetDestroyedCountNil(b bool)`

 SetDestroyedCountNil sets the value for DestroyedCount to be an explicit nil

### UnsetDestroyedCount
`func (o *GETImports200ResponseDataInnerAttributes) UnsetDestroyedCount()`

UnsetDestroyedCount ensures that no value is present for DestroyedCount, not even an explicit nil
### GetProcessedCount

`func (o *GETImports200ResponseDataInnerAttributes) GetProcessedCount() interface{}`

GetProcessedCount returns the ProcessedCount field if non-nil, zero value otherwise.

### GetProcessedCountOk

`func (o *GETImports200ResponseDataInnerAttributes) GetProcessedCountOk() (*interface{}, bool)`

GetProcessedCountOk returns a tuple with the ProcessedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedCount

`func (o *GETImports200ResponseDataInnerAttributes) SetProcessedCount(v interface{})`

SetProcessedCount sets ProcessedCount field to given value.

### HasProcessedCount

`func (o *GETImports200ResponseDataInnerAttributes) HasProcessedCount() bool`

HasProcessedCount returns a boolean if a field has been set.

### SetProcessedCountNil

`func (o *GETImports200ResponseDataInnerAttributes) SetProcessedCountNil(b bool)`

 SetProcessedCountNil sets the value for ProcessedCount to be an explicit nil

### UnsetProcessedCount
`func (o *GETImports200ResponseDataInnerAttributes) UnsetProcessedCount()`

UnsetProcessedCount ensures that no value is present for ProcessedCount, not even an explicit nil
### GetErrorsLog

`func (o *GETImports200ResponseDataInnerAttributes) GetErrorsLog() interface{}`

GetErrorsLog returns the ErrorsLog field if non-nil, zero value otherwise.

### GetErrorsLogOk

`func (o *GETImports200ResponseDataInnerAttributes) GetErrorsLogOk() (*interface{}, bool)`

GetErrorsLogOk returns a tuple with the ErrorsLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsLog

`func (o *GETImports200ResponseDataInnerAttributes) SetErrorsLog(v interface{})`

SetErrorsLog sets ErrorsLog field to given value.

### HasErrorsLog

`func (o *GETImports200ResponseDataInnerAttributes) HasErrorsLog() bool`

HasErrorsLog returns a boolean if a field has been set.

### SetErrorsLogNil

`func (o *GETImports200ResponseDataInnerAttributes) SetErrorsLogNil(b bool)`

 SetErrorsLogNil sets the value for ErrorsLog to be an explicit nil

### UnsetErrorsLog
`func (o *GETImports200ResponseDataInnerAttributes) UnsetErrorsLog()`

UnsetErrorsLog ensures that no value is present for ErrorsLog, not even an explicit nil
### GetWarningsLog

`func (o *GETImports200ResponseDataInnerAttributes) GetWarningsLog() interface{}`

GetWarningsLog returns the WarningsLog field if non-nil, zero value otherwise.

### GetWarningsLogOk

`func (o *GETImports200ResponseDataInnerAttributes) GetWarningsLogOk() (*interface{}, bool)`

GetWarningsLogOk returns a tuple with the WarningsLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningsLog

`func (o *GETImports200ResponseDataInnerAttributes) SetWarningsLog(v interface{})`

SetWarningsLog sets WarningsLog field to given value.

### HasWarningsLog

`func (o *GETImports200ResponseDataInnerAttributes) HasWarningsLog() bool`

HasWarningsLog returns a boolean if a field has been set.

### SetWarningsLogNil

`func (o *GETImports200ResponseDataInnerAttributes) SetWarningsLogNil(b bool)`

 SetWarningsLogNil sets the value for WarningsLog to be an explicit nil

### UnsetWarningsLog
`func (o *GETImports200ResponseDataInnerAttributes) UnsetWarningsLog()`

UnsetWarningsLog ensures that no value is present for WarningsLog, not even an explicit nil
### GetCleanupRecords

`func (o *GETImports200ResponseDataInnerAttributes) GetCleanupRecords() interface{}`

GetCleanupRecords returns the CleanupRecords field if non-nil, zero value otherwise.

### GetCleanupRecordsOk

`func (o *GETImports200ResponseDataInnerAttributes) GetCleanupRecordsOk() (*interface{}, bool)`

GetCleanupRecordsOk returns a tuple with the CleanupRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupRecords

`func (o *GETImports200ResponseDataInnerAttributes) SetCleanupRecords(v interface{})`

SetCleanupRecords sets CleanupRecords field to given value.

### HasCleanupRecords

`func (o *GETImports200ResponseDataInnerAttributes) HasCleanupRecords() bool`

HasCleanupRecords returns a boolean if a field has been set.

### SetCleanupRecordsNil

`func (o *GETImports200ResponseDataInnerAttributes) SetCleanupRecordsNil(b bool)`

 SetCleanupRecordsNil sets the value for CleanupRecords to be an explicit nil

### UnsetCleanupRecords
`func (o *GETImports200ResponseDataInnerAttributes) UnsetCleanupRecords()`

UnsetCleanupRecords ensures that no value is present for CleanupRecords, not even an explicit nil
### GetAttachmentUrl

`func (o *GETImports200ResponseDataInnerAttributes) GetAttachmentUrl() interface{}`

GetAttachmentUrl returns the AttachmentUrl field if non-nil, zero value otherwise.

### GetAttachmentUrlOk

`func (o *GETImports200ResponseDataInnerAttributes) GetAttachmentUrlOk() (*interface{}, bool)`

GetAttachmentUrlOk returns a tuple with the AttachmentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentUrl

`func (o *GETImports200ResponseDataInnerAttributes) SetAttachmentUrl(v interface{})`

SetAttachmentUrl sets AttachmentUrl field to given value.

### HasAttachmentUrl

`func (o *GETImports200ResponseDataInnerAttributes) HasAttachmentUrl() bool`

HasAttachmentUrl returns a boolean if a field has been set.

### SetAttachmentUrlNil

`func (o *GETImports200ResponseDataInnerAttributes) SetAttachmentUrlNil(b bool)`

 SetAttachmentUrlNil sets the value for AttachmentUrl to be an explicit nil

### UnsetAttachmentUrl
`func (o *GETImports200ResponseDataInnerAttributes) UnsetAttachmentUrl()`

UnsetAttachmentUrl ensures that no value is present for AttachmentUrl, not even an explicit nil
### GetCreatedAt

`func (o *GETImports200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETImports200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETImports200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETImports200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETImports200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETImports200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETImports200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETImports200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETImports200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETImports200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETImports200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETImports200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETImports200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETImports200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETImports200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETImports200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETImports200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETImports200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETImports200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETImports200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETImports200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETImports200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETImports200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETImports200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETImports200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETImports200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETImports200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETImports200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETImports200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETImports200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


