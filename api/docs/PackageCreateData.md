# PackageCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**PackageCreateDataAttributes**](PackageCreateDataAttributes.md) |  | 
**Relationships** | Pointer to [**PackageCreateDataRelationships**](PackageCreateDataRelationships.md) |  | [optional] 

## Methods

### NewPackageCreateData

`func NewPackageCreateData(type_ string, attributes PackageCreateDataAttributes, ) *PackageCreateData`

NewPackageCreateData instantiates a new PackageCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageCreateDataWithDefaults

`func NewPackageCreateDataWithDefaults() *PackageCreateData`

NewPackageCreateDataWithDefaults instantiates a new PackageCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PackageCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PackageCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PackageCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *PackageCreateData) GetAttributes() PackageCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PackageCreateData) GetAttributesOk() (*PackageCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PackageCreateData) SetAttributes(v PackageCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PackageCreateData) GetRelationships() PackageCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PackageCreateData) GetRelationshipsOk() (*PackageCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PackageCreateData) SetRelationships(v PackageCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PackageCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


