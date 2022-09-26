# ImportData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**ImportDataAttributes**](ImportDataAttributes.md) |  | 
**Relationships** | Pointer to [**ExportDataRelationships**](ExportDataRelationships.md) |  | [optional] 

## Methods

### NewImportData

`func NewImportData(type_ string, attributes ImportDataAttributes, ) *ImportData`

NewImportData instantiates a new ImportData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportDataWithDefaults

`func NewImportDataWithDefaults() *ImportData`

NewImportDataWithDefaults instantiates a new ImportData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ImportData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImportData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImportData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ImportData) GetAttributes() ImportDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ImportData) GetAttributesOk() (*ImportDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ImportData) SetAttributes(v ImportDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ImportData) GetRelationships() ExportDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ImportData) GetRelationshipsOk() (*ExportDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ImportData) SetRelationships(v ExportDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ImportData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


