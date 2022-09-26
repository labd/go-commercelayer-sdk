# POSTExports201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | The type of resource being exported. | 
**Format** | Pointer to **string** | The format of the export one of &#39;json&#39; (default) or &#39;csv&#39;. | [optional] 
**Includes** | Pointer to **[]string** | List of related resources that should be included in the export. | [optional] 
**Filters** | Pointer to **map[string]interface{}** | The filters used to select the records to be exported. | [optional] 
**SkipIds** | Pointer to **bool** | Send this attribute if you want to skip exporting of resources IDs. | [optional] 
**SkipRelIds** | Pointer to **bool** | Send this attribute if you want to skip exporting of relationships IDs. | [optional] 
**SkipTimestamps** | Pointer to **bool** | Send this attribute if you want to skip exporting of resources created_at and updated_at. | [optional] 
**SkipBlanks** | Pointer to **bool** | Send this attribute if you want to skip exporting of blank values (not suitable for csv format). | [optional] 
**SkipAmountFormats** | Pointer to **bool** | Send this attribute if you want to skip exporting of float and fromatted amounts. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTExports201ResponseDataAttributes

`func NewPOSTExports201ResponseDataAttributes(resourceType string, ) *POSTExports201ResponseDataAttributes`

NewPOSTExports201ResponseDataAttributes instantiates a new POSTExports201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTExports201ResponseDataAttributesWithDefaults

`func NewPOSTExports201ResponseDataAttributesWithDefaults() *POSTExports201ResponseDataAttributes`

NewPOSTExports201ResponseDataAttributesWithDefaults instantiates a new POSTExports201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *POSTExports201ResponseDataAttributes) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *POSTExports201ResponseDataAttributes) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *POSTExports201ResponseDataAttributes) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetFormat

`func (o *POSTExports201ResponseDataAttributes) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *POSTExports201ResponseDataAttributes) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *POSTExports201ResponseDataAttributes) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *POSTExports201ResponseDataAttributes) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetIncludes

`func (o *POSTExports201ResponseDataAttributes) GetIncludes() []string`

GetIncludes returns the Includes field if non-nil, zero value otherwise.

### GetIncludesOk

`func (o *POSTExports201ResponseDataAttributes) GetIncludesOk() (*[]string, bool)`

GetIncludesOk returns a tuple with the Includes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludes

`func (o *POSTExports201ResponseDataAttributes) SetIncludes(v []string)`

SetIncludes sets Includes field to given value.

### HasIncludes

`func (o *POSTExports201ResponseDataAttributes) HasIncludes() bool`

HasIncludes returns a boolean if a field has been set.

### GetFilters

`func (o *POSTExports201ResponseDataAttributes) GetFilters() map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *POSTExports201ResponseDataAttributes) GetFiltersOk() (*map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *POSTExports201ResponseDataAttributes) SetFilters(v map[string]interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *POSTExports201ResponseDataAttributes) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSkipIds

`func (o *POSTExports201ResponseDataAttributes) GetSkipIds() bool`

GetSkipIds returns the SkipIds field if non-nil, zero value otherwise.

### GetSkipIdsOk

`func (o *POSTExports201ResponseDataAttributes) GetSkipIdsOk() (*bool, bool)`

GetSkipIdsOk returns a tuple with the SkipIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipIds

`func (o *POSTExports201ResponseDataAttributes) SetSkipIds(v bool)`

SetSkipIds sets SkipIds field to given value.

### HasSkipIds

`func (o *POSTExports201ResponseDataAttributes) HasSkipIds() bool`

HasSkipIds returns a boolean if a field has been set.

### GetSkipRelIds

`func (o *POSTExports201ResponseDataAttributes) GetSkipRelIds() bool`

GetSkipRelIds returns the SkipRelIds field if non-nil, zero value otherwise.

### GetSkipRelIdsOk

`func (o *POSTExports201ResponseDataAttributes) GetSkipRelIdsOk() (*bool, bool)`

GetSkipRelIdsOk returns a tuple with the SkipRelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipRelIds

`func (o *POSTExports201ResponseDataAttributes) SetSkipRelIds(v bool)`

SetSkipRelIds sets SkipRelIds field to given value.

### HasSkipRelIds

`func (o *POSTExports201ResponseDataAttributes) HasSkipRelIds() bool`

HasSkipRelIds returns a boolean if a field has been set.

### GetSkipTimestamps

`func (o *POSTExports201ResponseDataAttributes) GetSkipTimestamps() bool`

GetSkipTimestamps returns the SkipTimestamps field if non-nil, zero value otherwise.

### GetSkipTimestampsOk

`func (o *POSTExports201ResponseDataAttributes) GetSkipTimestampsOk() (*bool, bool)`

GetSkipTimestampsOk returns a tuple with the SkipTimestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipTimestamps

`func (o *POSTExports201ResponseDataAttributes) SetSkipTimestamps(v bool)`

SetSkipTimestamps sets SkipTimestamps field to given value.

### HasSkipTimestamps

`func (o *POSTExports201ResponseDataAttributes) HasSkipTimestamps() bool`

HasSkipTimestamps returns a boolean if a field has been set.

### GetSkipBlanks

`func (o *POSTExports201ResponseDataAttributes) GetSkipBlanks() bool`

GetSkipBlanks returns the SkipBlanks field if non-nil, zero value otherwise.

### GetSkipBlanksOk

`func (o *POSTExports201ResponseDataAttributes) GetSkipBlanksOk() (*bool, bool)`

GetSkipBlanksOk returns a tuple with the SkipBlanks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipBlanks

`func (o *POSTExports201ResponseDataAttributes) SetSkipBlanks(v bool)`

SetSkipBlanks sets SkipBlanks field to given value.

### HasSkipBlanks

`func (o *POSTExports201ResponseDataAttributes) HasSkipBlanks() bool`

HasSkipBlanks returns a boolean if a field has been set.

### GetSkipAmountFormats

`func (o *POSTExports201ResponseDataAttributes) GetSkipAmountFormats() bool`

GetSkipAmountFormats returns the SkipAmountFormats field if non-nil, zero value otherwise.

### GetSkipAmountFormatsOk

`func (o *POSTExports201ResponseDataAttributes) GetSkipAmountFormatsOk() (*bool, bool)`

GetSkipAmountFormatsOk returns a tuple with the SkipAmountFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipAmountFormats

`func (o *POSTExports201ResponseDataAttributes) SetSkipAmountFormats(v bool)`

SetSkipAmountFormats sets SkipAmountFormats field to given value.

### HasSkipAmountFormats

`func (o *POSTExports201ResponseDataAttributes) HasSkipAmountFormats() bool`

HasSkipAmountFormats returns a boolean if a field has been set.

### GetReference

`func (o *POSTExports201ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTExports201ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTExports201ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTExports201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *POSTExports201ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTExports201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTExports201ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTExports201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *POSTExports201ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTExports201ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTExports201ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTExports201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


