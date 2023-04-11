# POSTSatispayGatewaysRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTManualGatewaysRequestDataAttributes**](POSTManualGatewaysRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTSatispayGatewaysRequestDataRelationships**](POSTSatispayGatewaysRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTSatispayGatewaysRequestData

`func NewPOSTSatispayGatewaysRequestData(type_ interface{}, attributes POSTManualGatewaysRequestDataAttributes, ) *POSTSatispayGatewaysRequestData`

NewPOSTSatispayGatewaysRequestData instantiates a new POSTSatispayGatewaysRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTSatispayGatewaysRequestDataWithDefaults

`func NewPOSTSatispayGatewaysRequestDataWithDefaults() *POSTSatispayGatewaysRequestData`

NewPOSTSatispayGatewaysRequestDataWithDefaults instantiates a new POSTSatispayGatewaysRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTSatispayGatewaysRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTSatispayGatewaysRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTSatispayGatewaysRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTSatispayGatewaysRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTSatispayGatewaysRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTSatispayGatewaysRequestData) GetAttributes() POSTManualGatewaysRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTSatispayGatewaysRequestData) GetAttributesOk() (*POSTManualGatewaysRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTSatispayGatewaysRequestData) SetAttributes(v POSTManualGatewaysRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTSatispayGatewaysRequestData) GetRelationships() POSTSatispayGatewaysRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTSatispayGatewaysRequestData) GetRelationshipsOk() (*POSTSatispayGatewaysRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTSatispayGatewaysRequestData) SetRelationships(v POSTSatispayGatewaysRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTSatispayGatewaysRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


