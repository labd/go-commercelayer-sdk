# LinkUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHLinksLinkId200ResponseDataAttributes**](PATCHLinksLinkId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**LinkUpdateDataRelationships**](LinkUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewLinkUpdateData

`func NewLinkUpdateData(type_ interface{}, id interface{}, attributes PATCHLinksLinkId200ResponseDataAttributes, ) *LinkUpdateData`

NewLinkUpdateData instantiates a new LinkUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkUpdateDataWithDefaults

`func NewLinkUpdateDataWithDefaults() *LinkUpdateData`

NewLinkUpdateDataWithDefaults instantiates a new LinkUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LinkUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *LinkUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *LinkUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *LinkUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinkUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinkUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *LinkUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *LinkUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *LinkUpdateData) GetAttributes() PATCHLinksLinkId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *LinkUpdateData) GetAttributesOk() (*PATCHLinksLinkId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *LinkUpdateData) SetAttributes(v PATCHLinksLinkId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *LinkUpdateData) GetRelationships() LinkUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *LinkUpdateData) GetRelationshipsOk() (*LinkUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *LinkUpdateData) SetRelationships(v LinkUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *LinkUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


