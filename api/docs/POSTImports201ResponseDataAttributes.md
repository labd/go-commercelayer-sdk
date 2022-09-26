# POSTImports201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | The type of resource being imported. | 
**Format** | Pointer to **string** | The format of the import inputs one of &#39;json&#39; (default) or &#39;csv&#39;. | [optional] 
**ParentResourceId** | Pointer to **string** | The ID of the parent resource to be associated with imported data. | [optional] 
**Inputs** | **[]map[string]interface{}** | Array of objects representing the resources that are being imported. | 
**CleanupRecords** | Pointer to **bool** | Indicates if the import should cleanup records that are not included in the inputs array. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTImports201ResponseDataAttributes

`func NewPOSTImports201ResponseDataAttributes(resourceType string, inputs []map[string]interface{}, ) *POSTImports201ResponseDataAttributes`

NewPOSTImports201ResponseDataAttributes instantiates a new POSTImports201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTImports201ResponseDataAttributesWithDefaults

`func NewPOSTImports201ResponseDataAttributesWithDefaults() *POSTImports201ResponseDataAttributes`

NewPOSTImports201ResponseDataAttributesWithDefaults instantiates a new POSTImports201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *POSTImports201ResponseDataAttributes) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *POSTImports201ResponseDataAttributes) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *POSTImports201ResponseDataAttributes) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetFormat

`func (o *POSTImports201ResponseDataAttributes) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *POSTImports201ResponseDataAttributes) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *POSTImports201ResponseDataAttributes) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *POSTImports201ResponseDataAttributes) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetParentResourceId

`func (o *POSTImports201ResponseDataAttributes) GetParentResourceId() string`

GetParentResourceId returns the ParentResourceId field if non-nil, zero value otherwise.

### GetParentResourceIdOk

`func (o *POSTImports201ResponseDataAttributes) GetParentResourceIdOk() (*string, bool)`

GetParentResourceIdOk returns a tuple with the ParentResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentResourceId

`func (o *POSTImports201ResponseDataAttributes) SetParentResourceId(v string)`

SetParentResourceId sets ParentResourceId field to given value.

### HasParentResourceId

`func (o *POSTImports201ResponseDataAttributes) HasParentResourceId() bool`

HasParentResourceId returns a boolean if a field has been set.

### GetInputs

`func (o *POSTImports201ResponseDataAttributes) GetInputs() []map[string]interface{}`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *POSTImports201ResponseDataAttributes) GetInputsOk() (*[]map[string]interface{}, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *POSTImports201ResponseDataAttributes) SetInputs(v []map[string]interface{})`

SetInputs sets Inputs field to given value.


### GetCleanupRecords

`func (o *POSTImports201ResponseDataAttributes) GetCleanupRecords() bool`

GetCleanupRecords returns the CleanupRecords field if non-nil, zero value otherwise.

### GetCleanupRecordsOk

`func (o *POSTImports201ResponseDataAttributes) GetCleanupRecordsOk() (*bool, bool)`

GetCleanupRecordsOk returns a tuple with the CleanupRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupRecords

`func (o *POSTImports201ResponseDataAttributes) SetCleanupRecords(v bool)`

SetCleanupRecords sets CleanupRecords field to given value.

### HasCleanupRecords

`func (o *POSTImports201ResponseDataAttributes) HasCleanupRecords() bool`

HasCleanupRecords returns a boolean if a field has been set.

### GetReference

`func (o *POSTImports201ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTImports201ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTImports201ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTImports201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *POSTImports201ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTImports201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTImports201ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTImports201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *POSTImports201ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTImports201ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTImports201ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTImports201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


