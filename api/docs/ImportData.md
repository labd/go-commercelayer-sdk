# ImportData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETImportsImportId200ResponseDataAttributes**](GETImportsImportId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**ExportDataRelationships**](ExportDataRelationships.md) |  | [optional] 

## Methods

### NewImportData

`func NewImportData(type_ interface{}, attributes GETImportsImportId200ResponseDataAttributes, ) *ImportData`

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

`func (o *ImportData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImportData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImportData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ImportData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ImportData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *ImportData) GetAttributes() GETImportsImportId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ImportData) GetAttributesOk() (*GETImportsImportId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ImportData) SetAttributes(v GETImportsImportId200ResponseDataAttributes)`

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


