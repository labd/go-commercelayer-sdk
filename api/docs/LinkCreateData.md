# LinkCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTLinks201ResponseDataAttributes**](POSTLinks201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**LinkCreateDataRelationships**](LinkCreateDataRelationships.md) |  | [optional] 

## Methods

### NewLinkCreateData

`func NewLinkCreateData(type_ interface{}, attributes POSTLinks201ResponseDataAttributes, ) *LinkCreateData`

NewLinkCreateData instantiates a new LinkCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkCreateDataWithDefaults

`func NewLinkCreateDataWithDefaults() *LinkCreateData`

NewLinkCreateDataWithDefaults instantiates a new LinkCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LinkCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *LinkCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *LinkCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *LinkCreateData) GetAttributes() POSTLinks201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *LinkCreateData) GetAttributesOk() (*POSTLinks201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *LinkCreateData) SetAttributes(v POSTLinks201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *LinkCreateData) GetRelationships() LinkCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *LinkCreateData) GetRelationshipsOk() (*LinkCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *LinkCreateData) SetRelationships(v LinkCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *LinkCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


