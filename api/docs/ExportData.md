# ExportData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**ExportDataAttributes**](ExportDataAttributes.md) |  | 
**Relationships** | Pointer to [**ExportDataRelationships**](ExportDataRelationships.md) |  | [optional] 

## Methods

### NewExportData

`func NewExportData(type_ string, attributes ExportDataAttributes, ) *ExportData`

NewExportData instantiates a new ExportData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportDataWithDefaults

`func NewExportDataWithDefaults() *ExportData`

NewExportDataWithDefaults instantiates a new ExportData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExportData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExportData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExportData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ExportData) GetAttributes() ExportDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ExportData) GetAttributesOk() (*ExportDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ExportData) SetAttributes(v ExportDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ExportData) GetRelationships() ExportDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ExportData) GetRelationshipsOk() (*ExportDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ExportData) SetRelationships(v ExportDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ExportData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


