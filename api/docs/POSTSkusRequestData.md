# POSTSkusRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTSkusRequestDataAttributes**](POSTSkusRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTSkusRequestDataRelationships**](POSTSkusRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTSkusRequestData

`func NewPOSTSkusRequestData(type_ interface{}, attributes POSTSkusRequestDataAttributes, ) *POSTSkusRequestData`

NewPOSTSkusRequestData instantiates a new POSTSkusRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTSkusRequestDataWithDefaults

`func NewPOSTSkusRequestDataWithDefaults() *POSTSkusRequestData`

NewPOSTSkusRequestDataWithDefaults instantiates a new POSTSkusRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTSkusRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTSkusRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTSkusRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTSkusRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTSkusRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTSkusRequestData) GetAttributes() POSTSkusRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTSkusRequestData) GetAttributesOk() (*POSTSkusRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTSkusRequestData) SetAttributes(v POSTSkusRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTSkusRequestData) GetRelationships() POSTSkusRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTSkusRequestData) GetRelationshipsOk() (*POSTSkusRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTSkusRequestData) SetRelationships(v POSTSkusRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTSkusRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


