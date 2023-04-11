# POSTImports201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **interface{}** | The type of resource being imported. | 
**Format** | Pointer to **interface{}** | The format of the import inputs one of &#39;json&#39; (default) or &#39;csv&#39;. | [optional] 
**ParentResourceId** | Pointer to **interface{}** | The ID of the parent resource to be associated with imported data. | [optional] 
**Inputs** | **interface{}** | Array of objects representing the resources that are being imported. | 
**CleanupRecords** | Pointer to **interface{}** | Indicates if the import should cleanup records that are not included in the inputs array. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTImports201ResponseDataAttributes

`func NewPOSTImports201ResponseDataAttributes(resourceType interface{}, inputs interface{}, ) *POSTImports201ResponseDataAttributes`

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

`func (o *POSTImports201ResponseDataAttributes) GetResourceType() interface{}`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *POSTImports201ResponseDataAttributes) GetResourceTypeOk() (*interface{}, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *POSTImports201ResponseDataAttributes) SetResourceType(v interface{})`

SetResourceType sets ResourceType field to given value.


### SetResourceTypeNil

`func (o *POSTImports201ResponseDataAttributes) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *POSTImports201ResponseDataAttributes) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetFormat

`func (o *POSTImports201ResponseDataAttributes) GetFormat() interface{}`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *POSTImports201ResponseDataAttributes) GetFormatOk() (*interface{}, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *POSTImports201ResponseDataAttributes) SetFormat(v interface{})`

SetFormat sets Format field to given value.

### HasFormat

`func (o *POSTImports201ResponseDataAttributes) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *POSTImports201ResponseDataAttributes) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *POSTImports201ResponseDataAttributes) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetParentResourceId

`func (o *POSTImports201ResponseDataAttributes) GetParentResourceId() interface{}`

GetParentResourceId returns the ParentResourceId field if non-nil, zero value otherwise.

### GetParentResourceIdOk

`func (o *POSTImports201ResponseDataAttributes) GetParentResourceIdOk() (*interface{}, bool)`

GetParentResourceIdOk returns a tuple with the ParentResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentResourceId

`func (o *POSTImports201ResponseDataAttributes) SetParentResourceId(v interface{})`

SetParentResourceId sets ParentResourceId field to given value.

### HasParentResourceId

`func (o *POSTImports201ResponseDataAttributes) HasParentResourceId() bool`

HasParentResourceId returns a boolean if a field has been set.

### SetParentResourceIdNil

`func (o *POSTImports201ResponseDataAttributes) SetParentResourceIdNil(b bool)`

 SetParentResourceIdNil sets the value for ParentResourceId to be an explicit nil

### UnsetParentResourceId
`func (o *POSTImports201ResponseDataAttributes) UnsetParentResourceId()`

UnsetParentResourceId ensures that no value is present for ParentResourceId, not even an explicit nil
### GetInputs

`func (o *POSTImports201ResponseDataAttributes) GetInputs() interface{}`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *POSTImports201ResponseDataAttributes) GetInputsOk() (*interface{}, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *POSTImports201ResponseDataAttributes) SetInputs(v interface{})`

SetInputs sets Inputs field to given value.


### SetInputsNil

`func (o *POSTImports201ResponseDataAttributes) SetInputsNil(b bool)`

 SetInputsNil sets the value for Inputs to be an explicit nil

### UnsetInputs
`func (o *POSTImports201ResponseDataAttributes) UnsetInputs()`

UnsetInputs ensures that no value is present for Inputs, not even an explicit nil
### GetCleanupRecords

`func (o *POSTImports201ResponseDataAttributes) GetCleanupRecords() interface{}`

GetCleanupRecords returns the CleanupRecords field if non-nil, zero value otherwise.

### GetCleanupRecordsOk

`func (o *POSTImports201ResponseDataAttributes) GetCleanupRecordsOk() (*interface{}, bool)`

GetCleanupRecordsOk returns a tuple with the CleanupRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupRecords

`func (o *POSTImports201ResponseDataAttributes) SetCleanupRecords(v interface{})`

SetCleanupRecords sets CleanupRecords field to given value.

### HasCleanupRecords

`func (o *POSTImports201ResponseDataAttributes) HasCleanupRecords() bool`

HasCleanupRecords returns a boolean if a field has been set.

### SetCleanupRecordsNil

`func (o *POSTImports201ResponseDataAttributes) SetCleanupRecordsNil(b bool)`

 SetCleanupRecordsNil sets the value for CleanupRecords to be an explicit nil

### UnsetCleanupRecords
`func (o *POSTImports201ResponseDataAttributes) UnsetCleanupRecords()`

UnsetCleanupRecords ensures that no value is present for CleanupRecords, not even an explicit nil
### GetReference

`func (o *POSTImports201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTImports201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTImports201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTImports201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTImports201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTImports201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTImports201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTImports201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTImports201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTImports201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTImports201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTImports201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTImports201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTImports201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTImports201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTImports201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTImports201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTImports201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


