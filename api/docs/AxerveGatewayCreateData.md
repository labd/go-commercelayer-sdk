# AxerveGatewayCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTAxerveGateways201ResponseDataAttributes**](POSTAxerveGateways201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**AxerveGatewayCreateDataRelationships**](AxerveGatewayCreateDataRelationships.md) |  | [optional] 

## Methods

### NewAxerveGatewayCreateData

`func NewAxerveGatewayCreateData(type_ interface{}, attributes POSTAxerveGateways201ResponseDataAttributes, ) *AxerveGatewayCreateData`

NewAxerveGatewayCreateData instantiates a new AxerveGatewayCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAxerveGatewayCreateDataWithDefaults

`func NewAxerveGatewayCreateDataWithDefaults() *AxerveGatewayCreateData`

NewAxerveGatewayCreateDataWithDefaults instantiates a new AxerveGatewayCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AxerveGatewayCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AxerveGatewayCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AxerveGatewayCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *AxerveGatewayCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AxerveGatewayCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *AxerveGatewayCreateData) GetAttributes() POSTAxerveGateways201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AxerveGatewayCreateData) GetAttributesOk() (*POSTAxerveGateways201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AxerveGatewayCreateData) SetAttributes(v POSTAxerveGateways201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *AxerveGatewayCreateData) GetRelationships() AxerveGatewayCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AxerveGatewayCreateData) GetRelationshipsOk() (*AxerveGatewayCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AxerveGatewayCreateData) SetRelationships(v AxerveGatewayCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AxerveGatewayCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


