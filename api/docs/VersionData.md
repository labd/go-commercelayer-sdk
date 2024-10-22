# VersionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETVersionsVersionId200ResponseDataAttributes**](GETVersionsVersionId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewVersionData

`func NewVersionData(type_ interface{}, attributes GETVersionsVersionId200ResponseDataAttributes, ) *VersionData`

NewVersionData instantiates a new VersionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionDataWithDefaults

`func NewVersionDataWithDefaults() *VersionData`

NewVersionDataWithDefaults instantiates a new VersionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VersionData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VersionData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VersionData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *VersionData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *VersionData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *VersionData) GetAttributes() GETVersionsVersionId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *VersionData) GetAttributesOk() (*GETVersionsVersionId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *VersionData) SetAttributes(v GETVersionsVersionId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *VersionData) GetRelationships() interface{}`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *VersionData) GetRelationshipsOk() (*interface{}, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *VersionData) SetRelationships(v interface{})`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *VersionData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### SetRelationshipsNil

`func (o *VersionData) SetRelationshipsNil(b bool)`

 SetRelationshipsNil sets the value for Relationships to be an explicit nil

### UnsetRelationships
`func (o *VersionData) UnsetRelationships()`

UnsetRelationships ensures that no value is present for Relationships, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


