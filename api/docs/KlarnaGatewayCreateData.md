# KlarnaGatewayCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "klarna_gateways"]
**Attributes** | [**KlarnaGatewayCreateDataAttributes**](KlarnaGatewayCreateDataAttributes.md) |  | 
**Relationships** | Pointer to [**KlarnaGatewayCreateDataRelationships**](KlarnaGatewayCreateDataRelationships.md) |  | [optional] 

## Methods

### NewKlarnaGatewayCreateData

`func NewKlarnaGatewayCreateData(type_ string, attributes KlarnaGatewayCreateDataAttributes, ) *KlarnaGatewayCreateData`

NewKlarnaGatewayCreateData instantiates a new KlarnaGatewayCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKlarnaGatewayCreateDataWithDefaults

`func NewKlarnaGatewayCreateDataWithDefaults() *KlarnaGatewayCreateData`

NewKlarnaGatewayCreateDataWithDefaults instantiates a new KlarnaGatewayCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *KlarnaGatewayCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KlarnaGatewayCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KlarnaGatewayCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *KlarnaGatewayCreateData) GetAttributes() KlarnaGatewayCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *KlarnaGatewayCreateData) GetAttributesOk() (*KlarnaGatewayCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *KlarnaGatewayCreateData) SetAttributes(v KlarnaGatewayCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *KlarnaGatewayCreateData) GetRelationships() KlarnaGatewayCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *KlarnaGatewayCreateData) GetRelationshipsOk() (*KlarnaGatewayCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *KlarnaGatewayCreateData) SetRelationships(v KlarnaGatewayCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *KlarnaGatewayCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


