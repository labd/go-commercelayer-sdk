# AxerveGatewayUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes**](PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**AxerveGatewayCreateDataRelationships**](AxerveGatewayCreateDataRelationships.md) |  | [optional] 

## Methods

### NewAxerveGatewayUpdateData

`func NewAxerveGatewayUpdateData(type_ interface{}, id interface{}, attributes PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes, ) *AxerveGatewayUpdateData`

NewAxerveGatewayUpdateData instantiates a new AxerveGatewayUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAxerveGatewayUpdateDataWithDefaults

`func NewAxerveGatewayUpdateDataWithDefaults() *AxerveGatewayUpdateData`

NewAxerveGatewayUpdateDataWithDefaults instantiates a new AxerveGatewayUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AxerveGatewayUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AxerveGatewayUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AxerveGatewayUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *AxerveGatewayUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AxerveGatewayUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *AxerveGatewayUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AxerveGatewayUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AxerveGatewayUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *AxerveGatewayUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AxerveGatewayUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *AxerveGatewayUpdateData) GetAttributes() PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AxerveGatewayUpdateData) GetAttributesOk() (*PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AxerveGatewayUpdateData) SetAttributes(v PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *AxerveGatewayUpdateData) GetRelationships() AxerveGatewayCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AxerveGatewayUpdateData) GetRelationshipsOk() (*AxerveGatewayCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AxerveGatewayUpdateData) SetRelationships(v AxerveGatewayCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AxerveGatewayUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


