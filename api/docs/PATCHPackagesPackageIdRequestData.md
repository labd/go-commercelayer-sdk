# PATCHPackagesPackageIdRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHPackagesPackageIdRequestDataAttributes**](PATCHPackagesPackageIdRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**PATCHPackagesPackageIdRequestDataRelationships**](PATCHPackagesPackageIdRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPATCHPackagesPackageIdRequestData

`func NewPATCHPackagesPackageIdRequestData(type_ interface{}, id interface{}, attributes PATCHPackagesPackageIdRequestDataAttributes, ) *PATCHPackagesPackageIdRequestData`

NewPATCHPackagesPackageIdRequestData instantiates a new PATCHPackagesPackageIdRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHPackagesPackageIdRequestDataWithDefaults

`func NewPATCHPackagesPackageIdRequestDataWithDefaults() *PATCHPackagesPackageIdRequestData`

NewPATCHPackagesPackageIdRequestDataWithDefaults instantiates a new PATCHPackagesPackageIdRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PATCHPackagesPackageIdRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PATCHPackagesPackageIdRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PATCHPackagesPackageIdRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *PATCHPackagesPackageIdRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PATCHPackagesPackageIdRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *PATCHPackagesPackageIdRequestData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PATCHPackagesPackageIdRequestData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PATCHPackagesPackageIdRequestData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *PATCHPackagesPackageIdRequestData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PATCHPackagesPackageIdRequestData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *PATCHPackagesPackageIdRequestData) GetAttributes() PATCHPackagesPackageIdRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PATCHPackagesPackageIdRequestData) GetAttributesOk() (*PATCHPackagesPackageIdRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PATCHPackagesPackageIdRequestData) SetAttributes(v PATCHPackagesPackageIdRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PATCHPackagesPackageIdRequestData) GetRelationships() PATCHPackagesPackageIdRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PATCHPackagesPackageIdRequestData) GetRelationshipsOk() (*PATCHPackagesPackageIdRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PATCHPackagesPackageIdRequestData) SetRelationships(v PATCHPackagesPackageIdRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PATCHPackagesPackageIdRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


