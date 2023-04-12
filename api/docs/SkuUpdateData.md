# SkuUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHSkusSkuId200ResponseDataAttributes**](PATCHSkusSkuId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**SkuUpdateDataRelationships**](SkuUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewSkuUpdateData

`func NewSkuUpdateData(type_ interface{}, id interface{}, attributes PATCHSkusSkuId200ResponseDataAttributes, ) *SkuUpdateData`

NewSkuUpdateData instantiates a new SkuUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuUpdateDataWithDefaults

`func NewSkuUpdateDataWithDefaults() *SkuUpdateData`

NewSkuUpdateDataWithDefaults instantiates a new SkuUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SkuUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SkuUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SkuUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *SkuUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SkuUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *SkuUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SkuUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SkuUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *SkuUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SkuUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *SkuUpdateData) GetAttributes() PATCHSkusSkuId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SkuUpdateData) GetAttributesOk() (*PATCHSkusSkuId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SkuUpdateData) SetAttributes(v PATCHSkusSkuId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SkuUpdateData) GetRelationships() SkuUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SkuUpdateData) GetRelationshipsOk() (*SkuUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SkuUpdateData) SetRelationships(v SkuUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SkuUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


