# GETImports200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | Pointer to **string** | The type of resource being imported. | [optional] 
**ParentResourceId** | Pointer to **string** | The ID of the parent resource to be associated with imported data. | [optional] 
**Status** | Pointer to **string** | The import job status. One of &#39;pending&#39; (default), &#39;in_progress&#39;, &#39;interrupted&#39;, or &#39;completed&#39;. | [optional] 
**StartedAt** | Pointer to **string** | Time at which the import was started. | [optional] 
**CompletedAt** | Pointer to **string** | Time at which the import was completed. | [optional] 
**InterruptedAt** | Pointer to **string** | Time at which the import was interrupted. | [optional] 
**Inputs** | Pointer to **[]map[string]interface{}** | Array of objects representing the resources that are being imported. | [optional] 
**InputsSize** | Pointer to **int32** | Indicates the size of the objects to be imported. | [optional] 
**ErrorsCount** | Pointer to **int32** | Indicates the number of import errors, if any. | [optional] 
**WarningsCount** | Pointer to **int32** | Indicates the number of import warnings, if any. | [optional] 
**DestroyedCount** | Pointer to **int32** | Indicates the number of records that have been destroyed, if any. | [optional] 
**ProcessedCount** | Pointer to **int32** | Indicates the number of records that have been processed (created or updated). | [optional] 
**ErrorsLog** | Pointer to **map[string]interface{}** | Contains the import errors, if any. | [optional] 
**WarningsLog** | Pointer to **map[string]interface{}** | Contains the import warnings, if any. | [optional] 
**CleanupRecords** | Pointer to **bool** | Indicates if the import should cleanup records that are not included in the inputs array. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

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

`func (o *GETImports200ResponseDataInnerAttributes) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *GETImports200ResponseDataInnerAttributes) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *GETImports200ResponseDataInnerAttributes) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *GETImports200ResponseDataInnerAttributes) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetParentResourceId

`func (o *GETImports200ResponseDataInnerAttributes) GetParentResourceId() string`

GetParentResourceId returns the ParentResourceId field if non-nil, zero value otherwise.

### GetParentResourceIdOk

`func (o *GETImports200ResponseDataInnerAttributes) GetParentResourceIdOk() (*string, bool)`

GetParentResourceIdOk returns a tuple with the ParentResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentResourceId

`func (o *GETImports200ResponseDataInnerAttributes) SetParentResourceId(v string)`

SetParentResourceId sets ParentResourceId field to given value.

### HasParentResourceId

`func (o *GETImports200ResponseDataInnerAttributes) HasParentResourceId() bool`

HasParentResourceId returns a boolean if a field has been set.

### GetStatus

`func (o *GETImports200ResponseDataInnerAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETImports200ResponseDataInnerAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETImports200ResponseDataInnerAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETImports200ResponseDataInnerAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartedAt

`func (o *GETImports200ResponseDataInnerAttributes) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GETImports200ResponseDataInnerAttributes) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GETImports200ResponseDataInnerAttributes) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *GETImports200ResponseDataInnerAttributes) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *GETImports200ResponseDataInnerAttributes) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GETImports200ResponseDataInnerAttributes) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GETImports200ResponseDataInnerAttributes) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GETImports200ResponseDataInnerAttributes) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetInterruptedAt

`func (o *GETImports200ResponseDataInnerAttributes) GetInterruptedAt() string`

GetInterruptedAt returns the InterruptedAt field if non-nil, zero value otherwise.

### GetInterruptedAtOk

`func (o *GETImports200ResponseDataInnerAttributes) GetInterruptedAtOk() (*string, bool)`

GetInterruptedAtOk returns a tuple with the InterruptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptedAt

`func (o *GETImports200ResponseDataInnerAttributes) SetInterruptedAt(v string)`

SetInterruptedAt sets InterruptedAt field to given value.

### HasInterruptedAt

`func (o *GETImports200ResponseDataInnerAttributes) HasInterruptedAt() bool`

HasInterruptedAt returns a boolean if a field has been set.

### GetInputs

`func (o *GETImports200ResponseDataInnerAttributes) GetInputs() []map[string]interface{}`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *GETImports200ResponseDataInnerAttributes) GetInputsOk() (*[]map[string]interface{}, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *GETImports200ResponseDataInnerAttributes) SetInputs(v []map[string]interface{})`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *GETImports200ResponseDataInnerAttributes) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetInputsSize

`func (o *GETImports200ResponseDataInnerAttributes) GetInputsSize() int32`

GetInputsSize returns the InputsSize field if non-nil, zero value otherwise.

### GetInputsSizeOk

`func (o *GETImports200ResponseDataInnerAttributes) GetInputsSizeOk() (*int32, bool)`

GetInputsSizeOk returns a tuple with the InputsSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputsSize

`func (o *GETImports200ResponseDataInnerAttributes) SetInputsSize(v int32)`

SetInputsSize sets InputsSize field to given value.

### HasInputsSize

`func (o *GETImports200ResponseDataInnerAttributes) HasInputsSize() bool`

HasInputsSize returns a boolean if a field has been set.

### GetErrorsCount

`func (o *GETImports200ResponseDataInnerAttributes) GetErrorsCount() int32`

GetErrorsCount returns the ErrorsCount field if non-nil, zero value otherwise.

### GetErrorsCountOk

`func (o *GETImports200ResponseDataInnerAttributes) GetErrorsCountOk() (*int32, bool)`

GetErrorsCountOk returns a tuple with the ErrorsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsCount

`func (o *GETImports200ResponseDataInnerAttributes) SetErrorsCount(v int32)`

SetErrorsCount sets ErrorsCount field to given value.

### HasErrorsCount

`func (o *GETImports200ResponseDataInnerAttributes) HasErrorsCount() bool`

HasErrorsCount returns a boolean if a field has been set.

### GetWarningsCount

`func (o *GETImports200ResponseDataInnerAttributes) GetWarningsCount() int32`

GetWarningsCount returns the WarningsCount field if non-nil, zero value otherwise.

### GetWarningsCountOk

`func (o *GETImports200ResponseDataInnerAttributes) GetWarningsCountOk() (*int32, bool)`

GetWarningsCountOk returns a tuple with the WarningsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningsCount

`func (o *GETImports200ResponseDataInnerAttributes) SetWarningsCount(v int32)`

SetWarningsCount sets WarningsCount field to given value.

### HasWarningsCount

`func (o *GETImports200ResponseDataInnerAttributes) HasWarningsCount() bool`

HasWarningsCount returns a boolean if a field has been set.

### GetDestroyedCount

`func (o *GETImports200ResponseDataInnerAttributes) GetDestroyedCount() int32`

GetDestroyedCount returns the DestroyedCount field if non-nil, zero value otherwise.

### GetDestroyedCountOk

`func (o *GETImports200ResponseDataInnerAttributes) GetDestroyedCountOk() (*int32, bool)`

GetDestroyedCountOk returns a tuple with the DestroyedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyedCount

`func (o *GETImports200ResponseDataInnerAttributes) SetDestroyedCount(v int32)`

SetDestroyedCount sets DestroyedCount field to given value.

### HasDestroyedCount

`func (o *GETImports200ResponseDataInnerAttributes) HasDestroyedCount() bool`

HasDestroyedCount returns a boolean if a field has been set.

### GetProcessedCount

`func (o *GETImports200ResponseDataInnerAttributes) GetProcessedCount() int32`

GetProcessedCount returns the ProcessedCount field if non-nil, zero value otherwise.

### GetProcessedCountOk

`func (o *GETImports200ResponseDataInnerAttributes) GetProcessedCountOk() (*int32, bool)`

GetProcessedCountOk returns a tuple with the ProcessedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedCount

`func (o *GETImports200ResponseDataInnerAttributes) SetProcessedCount(v int32)`

SetProcessedCount sets ProcessedCount field to given value.

### HasProcessedCount

`func (o *GETImports200ResponseDataInnerAttributes) HasProcessedCount() bool`

HasProcessedCount returns a boolean if a field has been set.

### GetErrorsLog

`func (o *GETImports200ResponseDataInnerAttributes) GetErrorsLog() map[string]interface{}`

GetErrorsLog returns the ErrorsLog field if non-nil, zero value otherwise.

### GetErrorsLogOk

`func (o *GETImports200ResponseDataInnerAttributes) GetErrorsLogOk() (*map[string]interface{}, bool)`

GetErrorsLogOk returns a tuple with the ErrorsLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsLog

`func (o *GETImports200ResponseDataInnerAttributes) SetErrorsLog(v map[string]interface{})`

SetErrorsLog sets ErrorsLog field to given value.

### HasErrorsLog

`func (o *GETImports200ResponseDataInnerAttributes) HasErrorsLog() bool`

HasErrorsLog returns a boolean if a field has been set.

### GetWarningsLog

`func (o *GETImports200ResponseDataInnerAttributes) GetWarningsLog() map[string]interface{}`

GetWarningsLog returns the WarningsLog field if non-nil, zero value otherwise.

### GetWarningsLogOk

`func (o *GETImports200ResponseDataInnerAttributes) GetWarningsLogOk() (*map[string]interface{}, bool)`

GetWarningsLogOk returns a tuple with the WarningsLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningsLog

`func (o *GETImports200ResponseDataInnerAttributes) SetWarningsLog(v map[string]interface{})`

SetWarningsLog sets WarningsLog field to given value.

### HasWarningsLog

`func (o *GETImports200ResponseDataInnerAttributes) HasWarningsLog() bool`

HasWarningsLog returns a boolean if a field has been set.

### GetCleanupRecords

`func (o *GETImports200ResponseDataInnerAttributes) GetCleanupRecords() bool`

GetCleanupRecords returns the CleanupRecords field if non-nil, zero value otherwise.

### GetCleanupRecordsOk

`func (o *GETImports200ResponseDataInnerAttributes) GetCleanupRecordsOk() (*bool, bool)`

GetCleanupRecordsOk returns a tuple with the CleanupRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupRecords

`func (o *GETImports200ResponseDataInnerAttributes) SetCleanupRecords(v bool)`

SetCleanupRecords sets CleanupRecords field to given value.

### HasCleanupRecords

`func (o *GETImports200ResponseDataInnerAttributes) HasCleanupRecords() bool`

HasCleanupRecords returns a boolean if a field has been set.

### GetId

`func (o *GETImports200ResponseDataInnerAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETImports200ResponseDataInnerAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETImports200ResponseDataInnerAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETImports200ResponseDataInnerAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETImports200ResponseDataInnerAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETImports200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETImports200ResponseDataInnerAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETImports200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETImports200ResponseDataInnerAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETImports200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETImports200ResponseDataInnerAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETImports200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETImports200ResponseDataInnerAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETImports200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETImports200ResponseDataInnerAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETImports200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETImports200ResponseDataInnerAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETImports200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETImports200ResponseDataInnerAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETImports200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETImports200ResponseDataInnerAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETImports200ResponseDataInnerAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETImports200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETImports200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


