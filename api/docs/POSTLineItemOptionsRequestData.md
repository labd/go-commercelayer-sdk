# POSTLineItemOptionsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTLineItemOptionsRequestDataAttributes**](POSTLineItemOptionsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTLineItemOptionsRequestDataRelationships**](POSTLineItemOptionsRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTLineItemOptionsRequestData

`func NewPOSTLineItemOptionsRequestData(type_ interface{}, attributes POSTLineItemOptionsRequestDataAttributes, ) *POSTLineItemOptionsRequestData`

NewPOSTLineItemOptionsRequestData instantiates a new POSTLineItemOptionsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTLineItemOptionsRequestDataWithDefaults

`func NewPOSTLineItemOptionsRequestDataWithDefaults() *POSTLineItemOptionsRequestData`

NewPOSTLineItemOptionsRequestDataWithDefaults instantiates a new POSTLineItemOptionsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTLineItemOptionsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTLineItemOptionsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTLineItemOptionsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTLineItemOptionsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTLineItemOptionsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTLineItemOptionsRequestData) GetAttributes() POSTLineItemOptionsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTLineItemOptionsRequestData) GetAttributesOk() (*POSTLineItemOptionsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTLineItemOptionsRequestData) SetAttributes(v POSTLineItemOptionsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTLineItemOptionsRequestData) GetRelationships() POSTLineItemOptionsRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTLineItemOptionsRequestData) GetRelationshipsOk() (*POSTLineItemOptionsRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTLineItemOptionsRequestData) SetRelationships(v POSTLineItemOptionsRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTLineItemOptionsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


