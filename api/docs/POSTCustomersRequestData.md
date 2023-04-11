# POSTCustomersRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTCustomersRequestDataAttributes**](POSTCustomersRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTCustomersRequestDataRelationships**](POSTCustomersRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTCustomersRequestData

`func NewPOSTCustomersRequestData(type_ interface{}, attributes POSTCustomersRequestDataAttributes, ) *POSTCustomersRequestData`

NewPOSTCustomersRequestData instantiates a new POSTCustomersRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTCustomersRequestDataWithDefaults

`func NewPOSTCustomersRequestDataWithDefaults() *POSTCustomersRequestData`

NewPOSTCustomersRequestDataWithDefaults instantiates a new POSTCustomersRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTCustomersRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTCustomersRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTCustomersRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTCustomersRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTCustomersRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTCustomersRequestData) GetAttributes() POSTCustomersRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTCustomersRequestData) GetAttributesOk() (*POSTCustomersRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTCustomersRequestData) SetAttributes(v POSTCustomersRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTCustomersRequestData) GetRelationships() POSTCustomersRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTCustomersRequestData) GetRelationshipsOk() (*POSTCustomersRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTCustomersRequestData) SetRelationships(v POSTCustomersRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTCustomersRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


