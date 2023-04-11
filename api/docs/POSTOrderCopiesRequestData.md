# POSTOrderCopiesRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTOrderCopiesRequestDataAttributes**](POSTOrderCopiesRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTOrderCopiesRequestDataRelationships**](POSTOrderCopiesRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTOrderCopiesRequestData

`func NewPOSTOrderCopiesRequestData(type_ interface{}, attributes POSTOrderCopiesRequestDataAttributes, ) *POSTOrderCopiesRequestData`

NewPOSTOrderCopiesRequestData instantiates a new POSTOrderCopiesRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTOrderCopiesRequestDataWithDefaults

`func NewPOSTOrderCopiesRequestDataWithDefaults() *POSTOrderCopiesRequestData`

NewPOSTOrderCopiesRequestDataWithDefaults instantiates a new POSTOrderCopiesRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTOrderCopiesRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTOrderCopiesRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTOrderCopiesRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTOrderCopiesRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTOrderCopiesRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTOrderCopiesRequestData) GetAttributes() POSTOrderCopiesRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTOrderCopiesRequestData) GetAttributesOk() (*POSTOrderCopiesRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTOrderCopiesRequestData) SetAttributes(v POSTOrderCopiesRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTOrderCopiesRequestData) GetRelationships() POSTOrderCopiesRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTOrderCopiesRequestData) GetRelationshipsOk() (*POSTOrderCopiesRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTOrderCopiesRequestData) SetRelationships(v POSTOrderCopiesRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTOrderCopiesRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


