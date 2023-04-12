# LineItemOptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETLineItemOptionsLineItemOptionId200ResponseDataAttributes**](GETLineItemOptionsLineItemOptionId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**LineItemOptionDataRelationships**](LineItemOptionDataRelationships.md) |  | [optional] 

## Methods

### NewLineItemOptionData

`func NewLineItemOptionData(type_ interface{}, attributes GETLineItemOptionsLineItemOptionId200ResponseDataAttributes, ) *LineItemOptionData`

NewLineItemOptionData instantiates a new LineItemOptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemOptionDataWithDefaults

`func NewLineItemOptionDataWithDefaults() *LineItemOptionData`

NewLineItemOptionDataWithDefaults instantiates a new LineItemOptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LineItemOptionData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LineItemOptionData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LineItemOptionData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *LineItemOptionData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *LineItemOptionData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *LineItemOptionData) GetAttributes() GETLineItemOptionsLineItemOptionId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *LineItemOptionData) GetAttributesOk() (*GETLineItemOptionsLineItemOptionId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *LineItemOptionData) SetAttributes(v GETLineItemOptionsLineItemOptionId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *LineItemOptionData) GetRelationships() LineItemOptionDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *LineItemOptionData) GetRelationshipsOk() (*LineItemOptionDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *LineItemOptionData) SetRelationships(v LineItemOptionDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *LineItemOptionData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


