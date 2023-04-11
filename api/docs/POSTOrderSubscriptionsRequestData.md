# POSTOrderSubscriptionsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTOrderSubscriptionsRequestDataAttributes**](POSTOrderSubscriptionsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTOrderSubscriptionsRequestDataRelationships**](POSTOrderSubscriptionsRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTOrderSubscriptionsRequestData

`func NewPOSTOrderSubscriptionsRequestData(type_ interface{}, attributes POSTOrderSubscriptionsRequestDataAttributes, ) *POSTOrderSubscriptionsRequestData`

NewPOSTOrderSubscriptionsRequestData instantiates a new POSTOrderSubscriptionsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTOrderSubscriptionsRequestDataWithDefaults

`func NewPOSTOrderSubscriptionsRequestDataWithDefaults() *POSTOrderSubscriptionsRequestData`

NewPOSTOrderSubscriptionsRequestDataWithDefaults instantiates a new POSTOrderSubscriptionsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTOrderSubscriptionsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTOrderSubscriptionsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTOrderSubscriptionsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTOrderSubscriptionsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTOrderSubscriptionsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTOrderSubscriptionsRequestData) GetAttributes() POSTOrderSubscriptionsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTOrderSubscriptionsRequestData) GetAttributesOk() (*POSTOrderSubscriptionsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTOrderSubscriptionsRequestData) SetAttributes(v POSTOrderSubscriptionsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTOrderSubscriptionsRequestData) GetRelationships() POSTOrderSubscriptionsRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTOrderSubscriptionsRequestData) GetRelationshipsOk() (*POSTOrderSubscriptionsRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTOrderSubscriptionsRequestData) SetRelationships(v POSTOrderSubscriptionsRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTOrderSubscriptionsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


