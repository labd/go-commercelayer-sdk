# SkuListItemUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**SkuListItemCreateDataAttributes**](SkuListItemCreateDataAttributes.md) |  | 
**Relationships** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSkuListItemUpdateData

`func NewSkuListItemUpdateData(type_ string, id string, attributes SkuListItemCreateDataAttributes, ) *SkuListItemUpdateData`

NewSkuListItemUpdateData instantiates a new SkuListItemUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuListItemUpdateDataWithDefaults

`func NewSkuListItemUpdateDataWithDefaults() *SkuListItemUpdateData`

NewSkuListItemUpdateDataWithDefaults instantiates a new SkuListItemUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SkuListItemUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SkuListItemUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SkuListItemUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *SkuListItemUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SkuListItemUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SkuListItemUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *SkuListItemUpdateData) GetAttributes() SkuListItemCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SkuListItemUpdateData) GetAttributesOk() (*SkuListItemCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SkuListItemUpdateData) SetAttributes(v SkuListItemCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SkuListItemUpdateData) GetRelationships() map[string]interface{}`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SkuListItemUpdateData) GetRelationshipsOk() (*map[string]interface{}, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SkuListItemUpdateData) SetRelationships(v map[string]interface{})`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SkuListItemUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


