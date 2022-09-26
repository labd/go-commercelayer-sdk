# KlarnaGatewayData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**GETKlarnaGateways200ResponseDataInnerAttributes**](GETKlarnaGateways200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**KlarnaGatewayDataRelationships**](KlarnaGatewayDataRelationships.md) |  | [optional] 

## Methods

### NewKlarnaGatewayData

`func NewKlarnaGatewayData(type_ string, attributes GETKlarnaGateways200ResponseDataInnerAttributes, ) *KlarnaGatewayData`

NewKlarnaGatewayData instantiates a new KlarnaGatewayData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKlarnaGatewayDataWithDefaults

`func NewKlarnaGatewayDataWithDefaults() *KlarnaGatewayData`

NewKlarnaGatewayDataWithDefaults instantiates a new KlarnaGatewayData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *KlarnaGatewayData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KlarnaGatewayData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KlarnaGatewayData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *KlarnaGatewayData) GetAttributes() GETKlarnaGateways200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *KlarnaGatewayData) GetAttributesOk() (*GETKlarnaGateways200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *KlarnaGatewayData) SetAttributes(v GETKlarnaGateways200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *KlarnaGatewayData) GetRelationships() KlarnaGatewayDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *KlarnaGatewayData) GetRelationshipsOk() (*KlarnaGatewayDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *KlarnaGatewayData) SetRelationships(v KlarnaGatewayDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *KlarnaGatewayData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


