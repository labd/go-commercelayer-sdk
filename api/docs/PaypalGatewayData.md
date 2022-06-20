# PaypalGatewayData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "paypal_gateways"]
**Attributes** | [**KlarnaGatewayDataAttributes**](KlarnaGatewayDataAttributes.md) |  | 
**Relationships** | Pointer to [**PaypalGatewayDataRelationships**](PaypalGatewayDataRelationships.md) |  | [optional] 

## Methods

### NewPaypalGatewayData

`func NewPaypalGatewayData(type_ string, attributes KlarnaGatewayDataAttributes, ) *PaypalGatewayData`

NewPaypalGatewayData instantiates a new PaypalGatewayData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaypalGatewayDataWithDefaults

`func NewPaypalGatewayDataWithDefaults() *PaypalGatewayData`

NewPaypalGatewayDataWithDefaults instantiates a new PaypalGatewayData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaypalGatewayData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaypalGatewayData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaypalGatewayData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *PaypalGatewayData) GetAttributes() KlarnaGatewayDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PaypalGatewayData) GetAttributesOk() (*KlarnaGatewayDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PaypalGatewayData) SetAttributes(v KlarnaGatewayDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PaypalGatewayData) GetRelationships() PaypalGatewayDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PaypalGatewayData) GetRelationshipsOk() (*PaypalGatewayDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PaypalGatewayData) SetRelationships(v PaypalGatewayDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PaypalGatewayData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


