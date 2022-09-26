# ManualGatewayData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**GETManualGateways200ResponseDataInnerAttributes**](GETManualGateways200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**ManualGatewayDataRelationships**](ManualGatewayDataRelationships.md) |  | [optional] 

## Methods

### NewManualGatewayData

`func NewManualGatewayData(type_ string, attributes GETManualGateways200ResponseDataInnerAttributes, ) *ManualGatewayData`

NewManualGatewayData instantiates a new ManualGatewayData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualGatewayDataWithDefaults

`func NewManualGatewayDataWithDefaults() *ManualGatewayData`

NewManualGatewayDataWithDefaults instantiates a new ManualGatewayData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ManualGatewayData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManualGatewayData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManualGatewayData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ManualGatewayData) GetAttributes() GETManualGateways200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ManualGatewayData) GetAttributesOk() (*GETManualGateways200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ManualGatewayData) SetAttributes(v GETManualGateways200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ManualGatewayData) GetRelationships() ManualGatewayDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ManualGatewayData) GetRelationshipsOk() (*ManualGatewayDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ManualGatewayData) SetRelationships(v ManualGatewayDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ManualGatewayData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


