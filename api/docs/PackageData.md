# PackageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETPackagesPackageId200ResponseDataAttributes**](GETPackagesPackageId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**PackageDataRelationships**](PackageDataRelationships.md) |  | [optional] 

## Methods

### NewPackageData

`func NewPackageData(type_ interface{}, attributes GETPackagesPackageId200ResponseDataAttributes, ) *PackageData`

NewPackageData instantiates a new PackageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageDataWithDefaults

`func NewPackageDataWithDefaults() *PackageData`

NewPackageDataWithDefaults instantiates a new PackageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PackageData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PackageData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PackageData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *PackageData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PackageData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *PackageData) GetAttributes() GETPackagesPackageId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PackageData) GetAttributesOk() (*GETPackagesPackageId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PackageData) SetAttributes(v GETPackagesPackageId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PackageData) GetRelationships() PackageDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PackageData) GetRelationshipsOk() (*PackageDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PackageData) SetRelationships(v PackageDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PackageData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


