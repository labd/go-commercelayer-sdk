# POSTAxerveGatewaysRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTAxerveGatewaysRequestDataAttributes**](POSTAxerveGatewaysRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTAxerveGatewaysRequestDataRelationships**](POSTAxerveGatewaysRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTAxerveGatewaysRequestData

`func NewPOSTAxerveGatewaysRequestData(type_ interface{}, attributes POSTAxerveGatewaysRequestDataAttributes, ) *POSTAxerveGatewaysRequestData`

NewPOSTAxerveGatewaysRequestData instantiates a new POSTAxerveGatewaysRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTAxerveGatewaysRequestDataWithDefaults

`func NewPOSTAxerveGatewaysRequestDataWithDefaults() *POSTAxerveGatewaysRequestData`

NewPOSTAxerveGatewaysRequestDataWithDefaults instantiates a new POSTAxerveGatewaysRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTAxerveGatewaysRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTAxerveGatewaysRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTAxerveGatewaysRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTAxerveGatewaysRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTAxerveGatewaysRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTAxerveGatewaysRequestData) GetAttributes() POSTAxerveGatewaysRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTAxerveGatewaysRequestData) GetAttributesOk() (*POSTAxerveGatewaysRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTAxerveGatewaysRequestData) SetAttributes(v POSTAxerveGatewaysRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTAxerveGatewaysRequestData) GetRelationships() POSTAxerveGatewaysRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTAxerveGatewaysRequestData) GetRelationshipsOk() (*POSTAxerveGatewaysRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTAxerveGatewaysRequestData) SetRelationships(v POSTAxerveGatewaysRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTAxerveGatewaysRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


