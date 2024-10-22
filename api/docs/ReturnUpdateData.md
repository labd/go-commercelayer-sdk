# ReturnUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHReturnsReturnId200ResponseDataAttributes**](PATCHReturnsReturnId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**ReturnUpdateDataRelationships**](ReturnUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewReturnUpdateData

`func NewReturnUpdateData(type_ interface{}, id interface{}, attributes PATCHReturnsReturnId200ResponseDataAttributes, ) *ReturnUpdateData`

NewReturnUpdateData instantiates a new ReturnUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnUpdateDataWithDefaults

`func NewReturnUpdateDataWithDefaults() *ReturnUpdateData`

NewReturnUpdateDataWithDefaults instantiates a new ReturnUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ReturnUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReturnUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReturnUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ReturnUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ReturnUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *ReturnUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReturnUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReturnUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *ReturnUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ReturnUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *ReturnUpdateData) GetAttributes() PATCHReturnsReturnId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ReturnUpdateData) GetAttributesOk() (*PATCHReturnsReturnId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ReturnUpdateData) SetAttributes(v PATCHReturnsReturnId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ReturnUpdateData) GetRelationships() ReturnUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ReturnUpdateData) GetRelationshipsOk() (*ReturnUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ReturnUpdateData) SetRelationships(v ReturnUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ReturnUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


