# GETExports200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | Pointer to **string** | The type of resource being exported. | [optional] 
**Format** | Pointer to **string** | The format of the export one of &#39;json&#39; (default) or &#39;csv&#39;. | [optional] 
**Status** | Pointer to **string** | The export job status. One of &#39;pending&#39; (default), &#39;in_progress&#39;, or &#39;completed&#39;. | [optional] 
**Includes** | Pointer to **[]string** | List of related resources that should be included in the export. | [optional] 
**Filters** | Pointer to **map[string]interface{}** | The filters used to select the records to be exported. | [optional] 
**SkipIds** | Pointer to **bool** | Send this attribute if you want to skip exporting of resources IDs. | [optional] 
**SkipRelIds** | Pointer to **bool** | Send this attribute if you want to skip exporting of relationships IDs. | [optional] 
**SkipTimestamps** | Pointer to **bool** | Send this attribute if you want to skip exporting of resources created_at and updated_at. | [optional] 
**SkipBlanks** | Pointer to **bool** | Send this attribute if you want to skip exporting of blank values (not suitable for csv format). | [optional] 
**SkipAmountFormats** | Pointer to **bool** | Send this attribute if you want to skip exporting of float and fromatted amounts. | [optional] 
**StartedAt** | Pointer to **string** | Time at which the export was started. | [optional] 
**CompletedAt** | Pointer to **string** | Time at which the export was completed. | [optional] 
**InterruptedAt** | Pointer to **string** | Time at which the export was interrupted. | [optional] 
**RecordsCount** | Pointer to **int32** | Indicates the number of records to be exported. | [optional] 
**AttachmentUrl** | Pointer to **string** | The URL to the output file, which will be generated upon export completion. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

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

`func (o *GETExports200ResponseDataInnerAttributes) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *GETExports200ResponseDataInnerAttributes) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *GETExports200ResponseDataInnerAttributes) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *GETExports200ResponseDataInnerAttributes) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetFormat

`func (o *GETExports200ResponseDataInnerAttributes) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *GETExports200ResponseDataInnerAttributes) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *GETExports200ResponseDataInnerAttributes) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *GETExports200ResponseDataInnerAttributes) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetStatus

`func (o *GETExports200ResponseDataInnerAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETExports200ResponseDataInnerAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETExports200ResponseDataInnerAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETExports200ResponseDataInnerAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIncludes

`func (o *GETExports200ResponseDataInnerAttributes) GetIncludes() []string`

GetIncludes returns the Includes field if non-nil, zero value otherwise.

### GetIncludesOk

`func (o *GETExports200ResponseDataInnerAttributes) GetIncludesOk() (*[]string, bool)`

GetIncludesOk returns a tuple with the Includes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludes

`func (o *GETExports200ResponseDataInnerAttributes) SetIncludes(v []string)`

SetIncludes sets Includes field to given value.

### HasIncludes

`func (o *GETExports200ResponseDataInnerAttributes) HasIncludes() bool`

HasIncludes returns a boolean if a field has been set.

### GetFilters

`func (o *GETExports200ResponseDataInnerAttributes) GetFilters() map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GETExports200ResponseDataInnerAttributes) GetFiltersOk() (*map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GETExports200ResponseDataInnerAttributes) SetFilters(v map[string]interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *GETExports200ResponseDataInnerAttributes) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSkipIds

`func (o *GETExports200ResponseDataInnerAttributes) GetSkipIds() bool`

GetSkipIds returns the SkipIds field if non-nil, zero value otherwise.

### GetSkipIdsOk

`func (o *GETExports200ResponseDataInnerAttributes) GetSkipIdsOk() (*bool, bool)`

GetSkipIdsOk returns a tuple with the SkipIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipIds

`func (o *GETExports200ResponseDataInnerAttributes) SetSkipIds(v bool)`

SetSkipIds sets SkipIds field to given value.

### HasSkipIds

`func (o *GETExports200ResponseDataInnerAttributes) HasSkipIds() bool`

HasSkipIds returns a boolean if a field has been set.

### GetSkipRelIds

`func (o *GETExports200ResponseDataInnerAttributes) GetSkipRelIds() bool`

GetSkipRelIds returns the SkipRelIds field if non-nil, zero value otherwise.

### GetSkipRelIdsOk

`func (o *GETExports200ResponseDataInnerAttributes) GetSkipRelIdsOk() (*bool, bool)`

GetSkipRelIdsOk returns a tuple with the SkipRelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipRelIds

`func (o *GETExports200ResponseDataInnerAttributes) SetSkipRelIds(v bool)`

SetSkipRelIds sets SkipRelIds field to given value.

### HasSkipRelIds

`func (o *GETExports200ResponseDataInnerAttributes) HasSkipRelIds() bool`

HasSkipRelIds returns a boolean if a field has been set.

### GetSkipTimestamps

`func (o *GETExports200ResponseDataInnerAttributes) GetSkipTimestamps() bool`

GetSkipTimestamps returns the SkipTimestamps field if non-nil, zero value otherwise.

### GetSkipTimestampsOk

`func (o *GETExports200ResponseDataInnerAttributes) GetSkipTimestampsOk() (*bool, bool)`

GetSkipTimestampsOk returns a tuple with the SkipTimestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipTimestamps

`func (o *GETExports200ResponseDataInnerAttributes) SetSkipTimestamps(v bool)`

SetSkipTimestamps sets SkipTimestamps field to given value.

### HasSkipTimestamps

`func (o *GETExports200ResponseDataInnerAttributes) HasSkipTimestamps() bool`

HasSkipTimestamps returns a boolean if a field has been set.

### GetSkipBlanks

`func (o *GETExports200ResponseDataInnerAttributes) GetSkipBlanks() bool`

GetSkipBlanks returns the SkipBlanks field if non-nil, zero value otherwise.

### GetSkipBlanksOk

`func (o *GETExports200ResponseDataInnerAttributes) GetSkipBlanksOk() (*bool, bool)`

GetSkipBlanksOk returns a tuple with the SkipBlanks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipBlanks

`func (o *GETExports200ResponseDataInnerAttributes) SetSkipBlanks(v bool)`

SetSkipBlanks sets SkipBlanks field to given value.

### HasSkipBlanks

`func (o *GETExports200ResponseDataInnerAttributes) HasSkipBlanks() bool`

HasSkipBlanks returns a boolean if a field has been set.

### GetSkipAmountFormats

`func (o *GETExports200ResponseDataInnerAttributes) GetSkipAmountFormats() bool`

GetSkipAmountFormats returns the SkipAmountFormats field if non-nil, zero value otherwise.

### GetSkipAmountFormatsOk

`func (o *GETExports200ResponseDataInnerAttributes) GetSkipAmountFormatsOk() (*bool, bool)`

GetSkipAmountFormatsOk returns a tuple with the SkipAmountFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipAmountFormats

`func (o *GETExports200ResponseDataInnerAttributes) SetSkipAmountFormats(v bool)`

SetSkipAmountFormats sets SkipAmountFormats field to given value.

### HasSkipAmountFormats

`func (o *GETExports200ResponseDataInnerAttributes) HasSkipAmountFormats() bool`

HasSkipAmountFormats returns a boolean if a field has been set.

### GetStartedAt

`func (o *GETExports200ResponseDataInnerAttributes) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GETExports200ResponseDataInnerAttributes) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GETExports200ResponseDataInnerAttributes) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *GETExports200ResponseDataInnerAttributes) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *GETExports200ResponseDataInnerAttributes) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GETExports200ResponseDataInnerAttributes) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GETExports200ResponseDataInnerAttributes) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GETExports200ResponseDataInnerAttributes) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetInterruptedAt

`func (o *GETExports200ResponseDataInnerAttributes) GetInterruptedAt() string`

GetInterruptedAt returns the InterruptedAt field if non-nil, zero value otherwise.

### GetInterruptedAtOk

`func (o *GETExports200ResponseDataInnerAttributes) GetInterruptedAtOk() (*string, bool)`

GetInterruptedAtOk returns a tuple with the InterruptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptedAt

`func (o *GETExports200ResponseDataInnerAttributes) SetInterruptedAt(v string)`

SetInterruptedAt sets InterruptedAt field to given value.

### HasInterruptedAt

`func (o *GETExports200ResponseDataInnerAttributes) HasInterruptedAt() bool`

HasInterruptedAt returns a boolean if a field has been set.

### GetRecordsCount

`func (o *GETExports200ResponseDataInnerAttributes) GetRecordsCount() int32`

GetRecordsCount returns the RecordsCount field if non-nil, zero value otherwise.

### GetRecordsCountOk

`func (o *GETExports200ResponseDataInnerAttributes) GetRecordsCountOk() (*int32, bool)`

GetRecordsCountOk returns a tuple with the RecordsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsCount

`func (o *GETExports200ResponseDataInnerAttributes) SetRecordsCount(v int32)`

SetRecordsCount sets RecordsCount field to given value.

### HasRecordsCount

`func (o *GETExports200ResponseDataInnerAttributes) HasRecordsCount() bool`

HasRecordsCount returns a boolean if a field has been set.

### GetAttachmentUrl

`func (o *GETExports200ResponseDataInnerAttributes) GetAttachmentUrl() string`

GetAttachmentUrl returns the AttachmentUrl field if non-nil, zero value otherwise.

### GetAttachmentUrlOk

`func (o *GETExports200ResponseDataInnerAttributes) GetAttachmentUrlOk() (*string, bool)`

GetAttachmentUrlOk returns a tuple with the AttachmentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentUrl

`func (o *GETExports200ResponseDataInnerAttributes) SetAttachmentUrl(v string)`

SetAttachmentUrl sets AttachmentUrl field to given value.

### HasAttachmentUrl

`func (o *GETExports200ResponseDataInnerAttributes) HasAttachmentUrl() bool`

HasAttachmentUrl returns a boolean if a field has been set.

### GetId

`func (o *GETExports200ResponseDataInnerAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETExports200ResponseDataInnerAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETExports200ResponseDataInnerAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETExports200ResponseDataInnerAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETExports200ResponseDataInnerAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETExports200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETExports200ResponseDataInnerAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETExports200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETExports200ResponseDataInnerAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETExports200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETExports200ResponseDataInnerAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETExports200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETExports200ResponseDataInnerAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETExports200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETExports200ResponseDataInnerAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETExports200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETExports200ResponseDataInnerAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETExports200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETExports200ResponseDataInnerAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETExports200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETExports200ResponseDataInnerAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETExports200ResponseDataInnerAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETExports200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETExports200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


