# BraintreeGatewayCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "braintree_gateways"]
**Attributes** | [**POSTBraintreeGateways201ResponseDataAttributes**](POSTBraintreeGateways201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTBraintreeGateways201ResponseDataRelationships**](POSTBraintreeGateways201ResponseDataRelationships.md) |  | [optional] 

## Methods

### NewBraintreeGatewayCreateData

`func NewBraintreeGatewayCreateData(type_ string, attributes POSTBraintreeGateways201ResponseDataAttributes, ) *BraintreeGatewayCreateData`

NewBraintreeGatewayCreateData instantiates a new BraintreeGatewayCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBraintreeGatewayCreateDataWithDefaults

`func NewBraintreeGatewayCreateDataWithDefaults() *BraintreeGatewayCreateData`

NewBraintreeGatewayCreateDataWithDefaults instantiates a new BraintreeGatewayCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BraintreeGatewayCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BraintreeGatewayCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BraintreeGatewayCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BraintreeGatewayCreateData) GetAttributes() POSTBraintreeGateways201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BraintreeGatewayCreateData) GetAttributesOk() (*POSTBraintreeGateways201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BraintreeGatewayCreateData) SetAttributes(v POSTBraintreeGateways201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *BraintreeGatewayCreateData) GetRelationships() POSTBraintreeGateways201ResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BraintreeGatewayCreateData) GetRelationshipsOk() (*POSTBraintreeGateways201ResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BraintreeGatewayCreateData) SetRelationships(v POSTBraintreeGateways201ResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BraintreeGatewayCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


