# StripeGatewayData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "stripe_gateways"]
**Attributes** | [**StripeGatewayDataAttributes**](StripeGatewayDataAttributes.md) |  | 
**Relationships** | Pointer to [**StripeGatewayDataRelationships**](StripeGatewayDataRelationships.md) |  | [optional] 

## Methods

### NewStripeGatewayData

`func NewStripeGatewayData(type_ string, attributes StripeGatewayDataAttributes, ) *StripeGatewayData`

NewStripeGatewayData instantiates a new StripeGatewayData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeGatewayDataWithDefaults

`func NewStripeGatewayDataWithDefaults() *StripeGatewayData`

NewStripeGatewayDataWithDefaults instantiates a new StripeGatewayData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StripeGatewayData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StripeGatewayData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StripeGatewayData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *StripeGatewayData) GetAttributes() StripeGatewayDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *StripeGatewayData) GetAttributesOk() (*StripeGatewayDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *StripeGatewayData) SetAttributes(v StripeGatewayDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *StripeGatewayData) GetRelationships() StripeGatewayDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *StripeGatewayData) GetRelationshipsOk() (*StripeGatewayDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *StripeGatewayData) SetRelationships(v StripeGatewayDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *StripeGatewayData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


