# ResourceErrorData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETResourceErrorsResourceErrorId200ResponseDataAttributes**](GETResourceErrorsResourceErrorId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**ResourceErrorDataRelationships**](ResourceErrorDataRelationships.md) |  | [optional] 

## Methods

### NewResourceErrorData

`func NewResourceErrorData(type_ interface{}, attributes GETResourceErrorsResourceErrorId200ResponseDataAttributes, ) *ResourceErrorData`

NewResourceErrorData instantiates a new ResourceErrorData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceErrorDataWithDefaults

`func NewResourceErrorDataWithDefaults() *ResourceErrorData`

NewResourceErrorDataWithDefaults instantiates a new ResourceErrorData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResourceErrorData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceErrorData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceErrorData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ResourceErrorData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ResourceErrorData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *ResourceErrorData) GetAttributes() GETResourceErrorsResourceErrorId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ResourceErrorData) GetAttributesOk() (*GETResourceErrorsResourceErrorId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ResourceErrorData) SetAttributes(v GETResourceErrorsResourceErrorId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ResourceErrorData) GetRelationships() ResourceErrorDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ResourceErrorData) GetRelationshipsOk() (*ResourceErrorDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ResourceErrorData) SetRelationships(v ResourceErrorDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ResourceErrorData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


