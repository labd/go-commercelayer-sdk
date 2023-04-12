# SatispayGatewayCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTManualGateways201ResponseDataAttributes**](POSTManualGateways201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**SatispayGatewayCreateDataRelationships**](SatispayGatewayCreateDataRelationships.md) |  | [optional] 

## Methods

### NewSatispayGatewayCreateData

`func NewSatispayGatewayCreateData(type_ interface{}, attributes POSTManualGateways201ResponseDataAttributes, ) *SatispayGatewayCreateData`

NewSatispayGatewayCreateData instantiates a new SatispayGatewayCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSatispayGatewayCreateDataWithDefaults

`func NewSatispayGatewayCreateDataWithDefaults() *SatispayGatewayCreateData`

NewSatispayGatewayCreateDataWithDefaults instantiates a new SatispayGatewayCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SatispayGatewayCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SatispayGatewayCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SatispayGatewayCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *SatispayGatewayCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SatispayGatewayCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *SatispayGatewayCreateData) GetAttributes() POSTManualGateways201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SatispayGatewayCreateData) GetAttributesOk() (*POSTManualGateways201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SatispayGatewayCreateData) SetAttributes(v POSTManualGateways201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SatispayGatewayCreateData) GetRelationships() SatispayGatewayCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SatispayGatewayCreateData) GetRelationshipsOk() (*SatispayGatewayCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SatispayGatewayCreateData) SetRelationships(v SatispayGatewayCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SatispayGatewayCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


