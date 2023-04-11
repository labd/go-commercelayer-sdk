# POSTOrdersRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTOrdersRequestDataAttributes**](POSTOrdersRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTOrdersRequestDataRelationships**](POSTOrdersRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTOrdersRequestData

`func NewPOSTOrdersRequestData(type_ interface{}, attributes POSTOrdersRequestDataAttributes, ) *POSTOrdersRequestData`

NewPOSTOrdersRequestData instantiates a new POSTOrdersRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTOrdersRequestDataWithDefaults

`func NewPOSTOrdersRequestDataWithDefaults() *POSTOrdersRequestData`

NewPOSTOrdersRequestDataWithDefaults instantiates a new POSTOrdersRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTOrdersRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTOrdersRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTOrdersRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTOrdersRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTOrdersRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTOrdersRequestData) GetAttributes() POSTOrdersRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTOrdersRequestData) GetAttributesOk() (*POSTOrdersRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTOrdersRequestData) SetAttributes(v POSTOrdersRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTOrdersRequestData) GetRelationships() POSTOrdersRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTOrdersRequestData) GetRelationshipsOk() (*POSTOrdersRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTOrdersRequestData) SetRelationships(v POSTOrdersRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTOrdersRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


