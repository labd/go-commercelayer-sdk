# ExternalGatewayData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETExternalGatewaysExternalGatewayId200ResponseDataAttributes**](GETExternalGatewaysExternalGatewayId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**ExternalGatewayDataRelationships**](ExternalGatewayDataRelationships.md) |  | [optional] 

## Methods

### NewExternalGatewayData

`func NewExternalGatewayData(type_ interface{}, attributes GETExternalGatewaysExternalGatewayId200ResponseDataAttributes, ) *ExternalGatewayData`

NewExternalGatewayData instantiates a new ExternalGatewayData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalGatewayDataWithDefaults

`func NewExternalGatewayDataWithDefaults() *ExternalGatewayData`

NewExternalGatewayDataWithDefaults instantiates a new ExternalGatewayData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExternalGatewayData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalGatewayData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalGatewayData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ExternalGatewayData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ExternalGatewayData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *ExternalGatewayData) GetAttributes() GETExternalGatewaysExternalGatewayId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ExternalGatewayData) GetAttributesOk() (*GETExternalGatewaysExternalGatewayId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ExternalGatewayData) SetAttributes(v GETExternalGatewaysExternalGatewayId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ExternalGatewayData) GetRelationships() ExternalGatewayDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ExternalGatewayData) GetRelationshipsOk() (*ExternalGatewayDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ExternalGatewayData) SetRelationships(v ExternalGatewayDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ExternalGatewayData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


