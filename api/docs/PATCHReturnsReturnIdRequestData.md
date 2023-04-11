# PATCHReturnsReturnIdRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHReturnsReturnIdRequestDataAttributes**](PATCHReturnsReturnIdRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**PATCHPackagesPackageIdRequestDataRelationships**](PATCHPackagesPackageIdRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPATCHReturnsReturnIdRequestData

`func NewPATCHReturnsReturnIdRequestData(type_ interface{}, id interface{}, attributes PATCHReturnsReturnIdRequestDataAttributes, ) *PATCHReturnsReturnIdRequestData`

NewPATCHReturnsReturnIdRequestData instantiates a new PATCHReturnsReturnIdRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHReturnsReturnIdRequestDataWithDefaults

`func NewPATCHReturnsReturnIdRequestDataWithDefaults() *PATCHReturnsReturnIdRequestData`

NewPATCHReturnsReturnIdRequestDataWithDefaults instantiates a new PATCHReturnsReturnIdRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PATCHReturnsReturnIdRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PATCHReturnsReturnIdRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PATCHReturnsReturnIdRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *PATCHReturnsReturnIdRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PATCHReturnsReturnIdRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *PATCHReturnsReturnIdRequestData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PATCHReturnsReturnIdRequestData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PATCHReturnsReturnIdRequestData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *PATCHReturnsReturnIdRequestData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PATCHReturnsReturnIdRequestData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *PATCHReturnsReturnIdRequestData) GetAttributes() PATCHReturnsReturnIdRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PATCHReturnsReturnIdRequestData) GetAttributesOk() (*PATCHReturnsReturnIdRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PATCHReturnsReturnIdRequestData) SetAttributes(v PATCHReturnsReturnIdRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PATCHReturnsReturnIdRequestData) GetRelationships() PATCHPackagesPackageIdRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PATCHReturnsReturnIdRequestData) GetRelationshipsOk() (*PATCHPackagesPackageIdRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PATCHReturnsReturnIdRequestData) SetRelationships(v PATCHPackagesPackageIdRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PATCHReturnsReturnIdRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


