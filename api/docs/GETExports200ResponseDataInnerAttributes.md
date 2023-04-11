# GETExports200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | Pointer to **interface{}** | The type of resource being exported. | [optional] 
**Format** | Pointer to **interface{}** | The format of the export one of &#39;json&#39; (default) or &#39;csv&#39;. | [optional] 
**Status** | Pointer to **interface{}** | The export job status. One of &#39;pending&#39; (default), &#39;in_progress&#39;, or &#39;completed&#39;. | [optional] 
**Includes** | Pointer to **[]interface{}** | List of related resources that should be included in the export. | [optional] 
**Filters** | Pointer to **interface{}** | The filters used to select the records to be exported. | [optional] 
**DryData** | Pointer to **interface{}** | Send this attribute if you want to skip exporting redundant attributes (IDs, timestamps, blanks, etc.), useful when combining export and import to duplicate your dataset. | [optional] 
**StartedAt** | Pointer to **interface{}** | Time at which the export was started. | [optional] 
**CompletedAt** | Pointer to **interface{}** | Time at which the export was completed. | [optional] 
**InterruptedAt** | Pointer to **interface{}** | Time at which the export was interrupted. | [optional] 
**RecordsCount** | Pointer to **interface{}** | Indicates the number of records to be exported. | [optional] 
**AttachmentUrl** | Pointer to **interface{}** | The URL to the output file, which will be generated upon export completion. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETExports200ResponseDataInnerAttributes

`func NewGETExports200ResponseDataInnerAttributes() *GETExports200ResponseDataInnerAttributes`

NewGETExports200ResponseDataInnerAttributes instantiates a new GETExports200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETExports200ResponseDataInnerAttributesWithDefaults

`func NewGETExports200ResponseDataInnerAttributesWithDefaults() *GETExports200ResponseDataInnerAttributes`

NewGETExports200ResponseDataInnerAttributesWithDefaults instantiates a new GETExports200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *GETExports200ResponseDataInnerAttributes) GetResourceType() interface{}`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *GETExports200ResponseDataInnerAttributes) GetResourceTypeOk() (*interface{}, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *GETExports200ResponseDataInnerAttributes) SetResourceType(v interface{})`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *GETExports200ResponseDataInnerAttributes) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *GETExports200ResponseDataInnerAttributes) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *GETExports200ResponseDataInnerAttributes) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetFormat

`func (o *GETExports200ResponseDataInnerAttributes) GetFormat() interface{}`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *GETExports200ResponseDataInnerAttributes) GetFormatOk() (*interface{}, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *GETExports200ResponseDataInnerAttributes) SetFormat(v interface{})`

SetFormat sets Format field to given value.

### HasFormat

`func (o *GETExports200ResponseDataInnerAttributes) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *GETExports200ResponseDataInnerAttributes) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *GETExports200ResponseDataInnerAttributes) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetStatus

`func (o *GETExports200ResponseDataInnerAttributes) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETExports200ResponseDataInnerAttributes) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETExports200ResponseDataInnerAttributes) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETExports200ResponseDataInnerAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GETExports200ResponseDataInnerAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GETExports200ResponseDataInnerAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetIncludes

`func (o *GETExports200ResponseDataInnerAttributes) GetIncludes() []interface{}`

GetIncludes returns the Includes field if non-nil, zero value otherwise.

### GetIncludesOk

`func (o *GETExports200ResponseDataInnerAttributes) GetIncludesOk() (*[]interface{}, bool)`

GetIncludesOk returns a tuple with the Includes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludes

`func (o *GETExports200ResponseDataInnerAttributes) SetIncludes(v []interface{})`

SetIncludes sets Includes field to given value.

### HasIncludes

`func (o *GETExports200ResponseDataInnerAttributes) HasIncludes() bool`

HasIncludes returns a boolean if a field has been set.

### GetFilters

`func (o *GETExports200ResponseDataInnerAttributes) GetFilters() interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GETExports200ResponseDataInnerAttributes) GetFiltersOk() (*interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GETExports200ResponseDataInnerAttributes) SetFilters(v interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *GETExports200ResponseDataInnerAttributes) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *GETExports200ResponseDataInnerAttributes) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *GETExports200ResponseDataInnerAttributes) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetDryData

`func (o *GETExports200ResponseDataInnerAttributes) GetDryData() interface{}`

GetDryData returns the DryData field if non-nil, zero value otherwise.

### GetDryDataOk

`func (o *GETExports200ResponseDataInnerAttributes) GetDryDataOk() (*interface{}, bool)`

GetDryDataOk returns a tuple with the DryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryData

`func (o *GETExports200ResponseDataInnerAttributes) SetDryData(v interface{})`

SetDryData sets DryData field to given value.

### HasDryData

`func (o *GETExports200ResponseDataInnerAttributes) HasDryData() bool`

HasDryData returns a boolean if a field has been set.

### SetDryDataNil

`func (o *GETExports200ResponseDataInnerAttributes) SetDryDataNil(b bool)`

 SetDryDataNil sets the value for DryData to be an explicit nil

### UnsetDryData
`func (o *GETExports200ResponseDataInnerAttributes) UnsetDryData()`

UnsetDryData ensures that no value is present for DryData, not even an explicit nil
### GetStartedAt

`func (o *GETExports200ResponseDataInnerAttributes) GetStartedAt() interface{}`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GETExports200ResponseDataInnerAttributes) GetStartedAtOk() (*interface{}, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GETExports200ResponseDataInnerAttributes) SetStartedAt(v interface{})`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *GETExports200ResponseDataInnerAttributes) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *GETExports200ResponseDataInnerAttributes) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *GETExports200ResponseDataInnerAttributes) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *GETExports200ResponseDataInnerAttributes) GetCompletedAt() interface{}`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GETExports200ResponseDataInnerAttributes) GetCompletedAtOk() (*interface{}, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GETExports200ResponseDataInnerAttributes) SetCompletedAt(v interface{})`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GETExports200ResponseDataInnerAttributes) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *GETExports200ResponseDataInnerAttributes) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *GETExports200ResponseDataInnerAttributes) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetInterruptedAt

`func (o *GETExports200ResponseDataInnerAttributes) GetInterruptedAt() interface{}`

GetInterruptedAt returns the InterruptedAt field if non-nil, zero value otherwise.

### GetInterruptedAtOk

`func (o *GETExports200ResponseDataInnerAttributes) GetInterruptedAtOk() (*interface{}, bool)`

GetInterruptedAtOk returns a tuple with the InterruptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptedAt

`func (o *GETExports200ResponseDataInnerAttributes) SetInterruptedAt(v interface{})`

SetInterruptedAt sets InterruptedAt field to given value.

### HasInterruptedAt

`func (o *GETExports200ResponseDataInnerAttributes) HasInterruptedAt() bool`

HasInterruptedAt returns a boolean if a field has been set.

### SetInterruptedAtNil

`func (o *GETExports200ResponseDataInnerAttributes) SetInterruptedAtNil(b bool)`

 SetInterruptedAtNil sets the value for InterruptedAt to be an explicit nil

### UnsetInterruptedAt
`func (o *GETExports200ResponseDataInnerAttributes) UnsetInterruptedAt()`

UnsetInterruptedAt ensures that no value is present for InterruptedAt, not even an explicit nil
### GetRecordsCount

`func (o *GETExports200ResponseDataInnerAttributes) GetRecordsCount() interface{}`

GetRecordsCount returns the RecordsCount field if non-nil, zero value otherwise.

### GetRecordsCountOk

`func (o *GETExports200ResponseDataInnerAttributes) GetRecordsCountOk() (*interface{}, bool)`

GetRecordsCountOk returns a tuple with the RecordsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsCount

`func (o *GETExports200ResponseDataInnerAttributes) SetRecordsCount(v interface{})`

SetRecordsCount sets RecordsCount field to given value.

### HasRecordsCount

`func (o *GETExports200ResponseDataInnerAttributes) HasRecordsCount() bool`

HasRecordsCount returns a boolean if a field has been set.

### SetRecordsCountNil

`func (o *GETExports200ResponseDataInnerAttributes) SetRecordsCountNil(b bool)`

 SetRecordsCountNil sets the value for RecordsCount to be an explicit nil

### UnsetRecordsCount
`func (o *GETExports200ResponseDataInnerAttributes) UnsetRecordsCount()`

UnsetRecordsCount ensures that no value is present for RecordsCount, not even an explicit nil
### GetAttachmentUrl

`func (o *GETExports200ResponseDataInnerAttributes) GetAttachmentUrl() interface{}`

GetAttachmentUrl returns the AttachmentUrl field if non-nil, zero value otherwise.

### GetAttachmentUrlOk

`func (o *GETExports200ResponseDataInnerAttributes) GetAttachmentUrlOk() (*interface{}, bool)`

GetAttachmentUrlOk returns a tuple with the AttachmentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentUrl

`func (o *GETExports200ResponseDataInnerAttributes) SetAttachmentUrl(v interface{})`

SetAttachmentUrl sets AttachmentUrl field to given value.

### HasAttachmentUrl

`func (o *GETExports200ResponseDataInnerAttributes) HasAttachmentUrl() bool`

HasAttachmentUrl returns a boolean if a field has been set.

### SetAttachmentUrlNil

`func (o *GETExports200ResponseDataInnerAttributes) SetAttachmentUrlNil(b bool)`

 SetAttachmentUrlNil sets the value for AttachmentUrl to be an explicit nil

### UnsetAttachmentUrl
`func (o *GETExports200ResponseDataInnerAttributes) UnsetAttachmentUrl()`

UnsetAttachmentUrl ensures that no value is present for AttachmentUrl, not even an explicit nil
### GetCreatedAt

`func (o *GETExports200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETExports200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETExports200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETExports200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETExports200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETExports200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETExports200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETExports200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETExports200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETExports200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETExports200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETExports200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETExports200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETExports200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETExports200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETExports200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETExports200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETExports200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETExports200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETExports200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETExports200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETExports200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETExports200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETExports200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETExports200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETExports200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETExports200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETExports200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETExports200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETExports200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


