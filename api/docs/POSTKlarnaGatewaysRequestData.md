# POSTKlarnaGatewaysRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTKlarnaGatewaysRequestDataAttributes**](POSTKlarnaGatewaysRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTKlarnaGatewaysRequestDataRelationships**](POSTKlarnaGatewaysRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTKlarnaGatewaysRequestData

`func NewPOSTKlarnaGatewaysRequestData(type_ interface{}, attributes POSTKlarnaGatewaysRequestDataAttributes, ) *POSTKlarnaGatewaysRequestData`

NewPOSTKlarnaGatewaysRequestData instantiates a new POSTKlarnaGatewaysRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTKlarnaGatewaysRequestDataWithDefaults

`func NewPOSTKlarnaGatewaysRequestDataWithDefaults() *POSTKlarnaGatewaysRequestData`

NewPOSTKlarnaGatewaysRequestDataWithDefaults instantiates a new POSTKlarnaGatewaysRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTKlarnaGatewaysRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTKlarnaGatewaysRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTKlarnaGatewaysRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTKlarnaGatewaysRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTKlarnaGatewaysRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTKlarnaGatewaysRequestData) GetAttributes() POSTKlarnaGatewaysRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTKlarnaGatewaysRequestData) GetAttributesOk() (*POSTKlarnaGatewaysRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTKlarnaGatewaysRequestData) SetAttributes(v POSTKlarnaGatewaysRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTKlarnaGatewaysRequestData) GetRelationships() POSTKlarnaGatewaysRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTKlarnaGatewaysRequestData) GetRelationshipsOk() (*POSTKlarnaGatewaysRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTKlarnaGatewaysRequestData) SetRelationships(v POSTKlarnaGatewaysRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTKlarnaGatewaysRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


