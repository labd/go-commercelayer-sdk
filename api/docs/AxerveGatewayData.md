# AxerveGatewayData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETAxerveGateways200ResponseDataInnerAttributes**](GETAxerveGateways200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**AxerveGatewayDataRelationships**](AxerveGatewayDataRelationships.md) |  | [optional] 

## Methods

### NewAxerveGatewayData

`func NewAxerveGatewayData(type_ interface{}, attributes GETAxerveGateways200ResponseDataInnerAttributes, ) *AxerveGatewayData`

NewAxerveGatewayData instantiates a new AxerveGatewayData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAxerveGatewayDataWithDefaults

`func NewAxerveGatewayDataWithDefaults() *AxerveGatewayData`

NewAxerveGatewayDataWithDefaults instantiates a new AxerveGatewayData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AxerveGatewayData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AxerveGatewayData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AxerveGatewayData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *AxerveGatewayData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AxerveGatewayData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *AxerveGatewayData) GetAttributes() GETAxerveGateways200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AxerveGatewayData) GetAttributesOk() (*GETAxerveGateways200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AxerveGatewayData) SetAttributes(v GETAxerveGateways200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *AxerveGatewayData) GetRelationships() AxerveGatewayDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AxerveGatewayData) GetRelationshipsOk() (*AxerveGatewayDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AxerveGatewayData) SetRelationships(v AxerveGatewayDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AxerveGatewayData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


