# SkuListData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETSkuListsSkuListId200ResponseDataAttributes**](GETSkuListsSkuListId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**SkuListDataRelationships**](SkuListDataRelationships.md) |  | [optional] 

## Methods

### NewSkuListData

`func NewSkuListData(type_ interface{}, attributes GETSkuListsSkuListId200ResponseDataAttributes, ) *SkuListData`

NewSkuListData instantiates a new SkuListData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuListDataWithDefaults

`func NewSkuListDataWithDefaults() *SkuListData`

NewSkuListDataWithDefaults instantiates a new SkuListData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SkuListData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SkuListData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SkuListData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *SkuListData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SkuListData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *SkuListData) GetAttributes() GETSkuListsSkuListId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SkuListData) GetAttributesOk() (*GETSkuListsSkuListId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SkuListData) SetAttributes(v GETSkuListsSkuListId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SkuListData) GetRelationships() SkuListDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SkuListData) GetRelationshipsOk() (*SkuListDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SkuListData) SetRelationships(v SkuListDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SkuListData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


