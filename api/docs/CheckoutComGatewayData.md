# CheckoutComGatewayData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes**](GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**CheckoutComGatewayDataRelationships**](CheckoutComGatewayDataRelationships.md) |  | [optional] 

## Methods

### NewCheckoutComGatewayData

`func NewCheckoutComGatewayData(type_ interface{}, attributes GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes, ) *CheckoutComGatewayData`

NewCheckoutComGatewayData instantiates a new CheckoutComGatewayData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutComGatewayDataWithDefaults

`func NewCheckoutComGatewayDataWithDefaults() *CheckoutComGatewayData`

NewCheckoutComGatewayDataWithDefaults instantiates a new CheckoutComGatewayData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CheckoutComGatewayData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckoutComGatewayData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckoutComGatewayData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CheckoutComGatewayData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CheckoutComGatewayData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *CheckoutComGatewayData) GetAttributes() GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CheckoutComGatewayData) GetAttributesOk() (*GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CheckoutComGatewayData) SetAttributes(v GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CheckoutComGatewayData) GetRelationships() CheckoutComGatewayDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CheckoutComGatewayData) GetRelationshipsOk() (*CheckoutComGatewayDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CheckoutComGatewayData) SetRelationships(v CheckoutComGatewayDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CheckoutComGatewayData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


