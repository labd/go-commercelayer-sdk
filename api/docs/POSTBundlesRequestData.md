# POSTBundlesRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTBundlesRequestDataAttributes**](POSTBundlesRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTBundlesRequestDataRelationships**](POSTBundlesRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTBundlesRequestData

`func NewPOSTBundlesRequestData(type_ interface{}, attributes POSTBundlesRequestDataAttributes, ) *POSTBundlesRequestData`

NewPOSTBundlesRequestData instantiates a new POSTBundlesRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTBundlesRequestDataWithDefaults

`func NewPOSTBundlesRequestDataWithDefaults() *POSTBundlesRequestData`

NewPOSTBundlesRequestDataWithDefaults instantiates a new POSTBundlesRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTBundlesRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTBundlesRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTBundlesRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTBundlesRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTBundlesRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTBundlesRequestData) GetAttributes() POSTBundlesRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTBundlesRequestData) GetAttributesOk() (*POSTBundlesRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTBundlesRequestData) SetAttributes(v POSTBundlesRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTBundlesRequestData) GetRelationships() POSTBundlesRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTBundlesRequestData) GetRelationshipsOk() (*POSTBundlesRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTBundlesRequestData) SetRelationships(v POSTBundlesRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTBundlesRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


