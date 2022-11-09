# POSTExports201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | The type of resource being exported. | 
**Format** | Pointer to **string** | The format of the export one of &#39;json&#39; (default) or &#39;csv&#39;. | [optional] 
**Includes** | Pointer to **[]string** | List of related resources that should be included in the export. | [optional] 
**Filters** | Pointer to **map[string]interface{}** | The filters used to select the records to be exported. | [optional] 
**DryData** | Pointer to **bool** | Send this attribute if you want to skip exporting redundant attributes (IDs, timespamps, blanks, etc), useful when combining export and import to duplicate your dataset. | [optional] 
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

### GetDryData

`func (o *POSTExports201ResponseDataAttributes) GetDryData() bool`

GetDryData returns the DryData field if non-nil, zero value otherwise.

### GetDryDataOk

`func (o *POSTExports201ResponseDataAttributes) GetDryDataOk() (*bool, bool)`

GetDryDataOk returns a tuple with the DryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryData

`func (o *POSTExports201ResponseDataAttributes) SetDryData(v bool)`

SetDryData sets DryData field to given value.

### HasDryData

`func (o *POSTExports201ResponseDataAttributes) HasDryData() bool`

HasDryData returns a boolean if a field has been set.

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


