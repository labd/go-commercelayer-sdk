# AdyenGatewayData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETAdyenGatewaysAdyenGatewayId200ResponseDataAttributes**](GETAdyenGatewaysAdyenGatewayId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**AdyenGatewayDataRelationships**](AdyenGatewayDataRelationships.md) |  | [optional] 

## Methods

### NewAdyenGatewayData

`func NewAdyenGatewayData(type_ interface{}, attributes GETAdyenGatewaysAdyenGatewayId200ResponseDataAttributes, ) *AdyenGatewayData`

NewAdyenGatewayData instantiates a new AdyenGatewayData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdyenGatewayDataWithDefaults

`func NewAdyenGatewayDataWithDefaults() *AdyenGatewayData`

NewAdyenGatewayDataWithDefaults instantiates a new AdyenGatewayData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AdyenGatewayData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdyenGatewayData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdyenGatewayData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *AdyenGatewayData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AdyenGatewayData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *AdyenGatewayData) GetAttributes() GETAdyenGatewaysAdyenGatewayId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AdyenGatewayData) GetAttributesOk() (*GETAdyenGatewaysAdyenGatewayId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AdyenGatewayData) SetAttributes(v GETAdyenGatewaysAdyenGatewayId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *AdyenGatewayData) GetRelationships() AdyenGatewayDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AdyenGatewayData) GetRelationshipsOk() (*AdyenGatewayDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AdyenGatewayData) SetRelationships(v AdyenGatewayDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AdyenGatewayData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


