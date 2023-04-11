# SatispayGatewayUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHManualGatewaysManualGatewayId200ResponseDataAttributes**](PATCHManualGatewaysManualGatewayId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**SatispayGatewayCreateDataRelationships**](SatispayGatewayCreateDataRelationships.md) |  | [optional] 

## Methods

### NewSatispayGatewayUpdateData

`func NewSatispayGatewayUpdateData(type_ interface{}, id interface{}, attributes PATCHManualGatewaysManualGatewayId200ResponseDataAttributes, ) *SatispayGatewayUpdateData`

NewSatispayGatewayUpdateData instantiates a new SatispayGatewayUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSatispayGatewayUpdateDataWithDefaults

`func NewSatispayGatewayUpdateDataWithDefaults() *SatispayGatewayUpdateData`

NewSatispayGatewayUpdateDataWithDefaults instantiates a new SatispayGatewayUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SatispayGatewayUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SatispayGatewayUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SatispayGatewayUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *SatispayGatewayUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SatispayGatewayUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *SatispayGatewayUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SatispayGatewayUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SatispayGatewayUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *SatispayGatewayUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SatispayGatewayUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *SatispayGatewayUpdateData) GetAttributes() PATCHManualGatewaysManualGatewayId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SatispayGatewayUpdateData) GetAttributesOk() (*PATCHManualGatewaysManualGatewayId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SatispayGatewayUpdateData) SetAttributes(v PATCHManualGatewaysManualGatewayId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SatispayGatewayUpdateData) GetRelationships() SatispayGatewayCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SatispayGatewayUpdateData) GetRelationshipsOk() (*SatispayGatewayCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SatispayGatewayUpdateData) SetRelationships(v SatispayGatewayCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SatispayGatewayUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


