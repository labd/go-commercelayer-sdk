# LineItemOptionUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes**](PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**LineItemOptionUpdateDataRelationships**](LineItemOptionUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewLineItemOptionUpdateData

`func NewLineItemOptionUpdateData(type_ interface{}, id interface{}, attributes PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes, ) *LineItemOptionUpdateData`

NewLineItemOptionUpdateData instantiates a new LineItemOptionUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemOptionUpdateDataWithDefaults

`func NewLineItemOptionUpdateDataWithDefaults() *LineItemOptionUpdateData`

NewLineItemOptionUpdateDataWithDefaults instantiates a new LineItemOptionUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LineItemOptionUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LineItemOptionUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LineItemOptionUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *LineItemOptionUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *LineItemOptionUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *LineItemOptionUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LineItemOptionUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LineItemOptionUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *LineItemOptionUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *LineItemOptionUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *LineItemOptionUpdateData) GetAttributes() PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *LineItemOptionUpdateData) GetAttributesOk() (*PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *LineItemOptionUpdateData) SetAttributes(v PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *LineItemOptionUpdateData) GetRelationships() LineItemOptionUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *LineItemOptionUpdateData) GetRelationshipsOk() (*LineItemOptionUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *LineItemOptionUpdateData) SetRelationships(v LineItemOptionUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *LineItemOptionUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


