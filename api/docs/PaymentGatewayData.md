# PaymentGatewayData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes**](GETKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**ManualGatewayDataRelationships**](ManualGatewayDataRelationships.md) |  | [optional] 

## Methods

### NewPaymentGatewayData

`func NewPaymentGatewayData(type_ interface{}, attributes GETKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes, ) *PaymentGatewayData`

NewPaymentGatewayData instantiates a new PaymentGatewayData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentGatewayDataWithDefaults

`func NewPaymentGatewayDataWithDefaults() *PaymentGatewayData`

NewPaymentGatewayDataWithDefaults instantiates a new PaymentGatewayData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentGatewayData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentGatewayData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentGatewayData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *PaymentGatewayData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PaymentGatewayData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *PaymentGatewayData) GetAttributes() GETKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PaymentGatewayData) GetAttributesOk() (*GETKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PaymentGatewayData) SetAttributes(v GETKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PaymentGatewayData) GetRelationships() ManualGatewayDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PaymentGatewayData) GetRelationshipsOk() (*ManualGatewayDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PaymentGatewayData) SetRelationships(v ManualGatewayDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PaymentGatewayData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


