# PackageUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**PackageUpdateDataAttributes**](PackageUpdateDataAttributes.md) |  | 
**Relationships** | Pointer to [**PackageUpdateDataRelationships**](PackageUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewPackageUpdateData

`func NewPackageUpdateData(type_ string, id string, attributes PackageUpdateDataAttributes, ) *PackageUpdateData`

NewPackageUpdateData instantiates a new PackageUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageUpdateDataWithDefaults

`func NewPackageUpdateDataWithDefaults() *PackageUpdateData`

NewPackageUpdateDataWithDefaults instantiates a new PackageUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PackageUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PackageUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PackageUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *PackageUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PackageUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PackageUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *PackageUpdateData) GetAttributes() PackageUpdateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PackageUpdateData) GetAttributesOk() (*PackageUpdateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PackageUpdateData) SetAttributes(v PackageUpdateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PackageUpdateData) GetRelationships() PackageUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PackageUpdateData) GetRelationshipsOk() (*PackageUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PackageUpdateData) SetRelationships(v PackageUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PackageUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


