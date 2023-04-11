# POSTPackagesRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTPackagesRequestDataAttributes**](POSTPackagesRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTPackagesRequestDataRelationships**](POSTPackagesRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTPackagesRequestData

`func NewPOSTPackagesRequestData(type_ interface{}, attributes POSTPackagesRequestDataAttributes, ) *POSTPackagesRequestData`

NewPOSTPackagesRequestData instantiates a new POSTPackagesRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTPackagesRequestDataWithDefaults

`func NewPOSTPackagesRequestDataWithDefaults() *POSTPackagesRequestData`

NewPOSTPackagesRequestDataWithDefaults instantiates a new POSTPackagesRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTPackagesRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTPackagesRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTPackagesRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTPackagesRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTPackagesRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTPackagesRequestData) GetAttributes() POSTPackagesRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTPackagesRequestData) GetAttributesOk() (*POSTPackagesRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTPackagesRequestData) SetAttributes(v POSTPackagesRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTPackagesRequestData) GetRelationships() POSTPackagesRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTPackagesRequestData) GetRelationshipsOk() (*POSTPackagesRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTPackagesRequestData) SetRelationships(v POSTPackagesRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTPackagesRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


