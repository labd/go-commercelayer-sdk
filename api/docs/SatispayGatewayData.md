# SatispayGatewayData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETSatispayGateways200ResponseDataInnerAttributes**](GETSatispayGateways200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**SatispayGatewayDataRelationships**](SatispayGatewayDataRelationships.md) |  | [optional] 

## Methods

### NewSatispayGatewayData

`func NewSatispayGatewayData(type_ interface{}, attributes GETSatispayGateways200ResponseDataInnerAttributes, ) *SatispayGatewayData`

NewSatispayGatewayData instantiates a new SatispayGatewayData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSatispayGatewayDataWithDefaults

`func NewSatispayGatewayDataWithDefaults() *SatispayGatewayData`

NewSatispayGatewayDataWithDefaults instantiates a new SatispayGatewayData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SatispayGatewayData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SatispayGatewayData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SatispayGatewayData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *SatispayGatewayData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SatispayGatewayData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *SatispayGatewayData) GetAttributes() GETSatispayGateways200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SatispayGatewayData) GetAttributesOk() (*GETSatispayGateways200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SatispayGatewayData) SetAttributes(v GETSatispayGateways200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SatispayGatewayData) GetRelationships() SatispayGatewayDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SatispayGatewayData) GetRelationshipsOk() (*SatispayGatewayDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SatispayGatewayData) SetRelationships(v SatispayGatewayDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SatispayGatewayData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


