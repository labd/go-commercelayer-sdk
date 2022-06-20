# BraintreeGatewayUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "braintree_gateways"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**BraintreeGatewayUpdateDataAttributes**](BraintreeGatewayUpdateDataAttributes.md) |  | 
**Relationships** | Pointer to [**BraintreeGatewayCreateDataRelationships**](BraintreeGatewayCreateDataRelationships.md) |  | [optional] 

## Methods

### NewBraintreeGatewayUpdateData

`func NewBraintreeGatewayUpdateData(type_ string, id string, attributes BraintreeGatewayUpdateDataAttributes, ) *BraintreeGatewayUpdateData`

NewBraintreeGatewayUpdateData instantiates a new BraintreeGatewayUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBraintreeGatewayUpdateDataWithDefaults

`func NewBraintreeGatewayUpdateDataWithDefaults() *BraintreeGatewayUpdateData`

NewBraintreeGatewayUpdateDataWithDefaults instantiates a new BraintreeGatewayUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BraintreeGatewayUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BraintreeGatewayUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BraintreeGatewayUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *BraintreeGatewayUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BraintreeGatewayUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BraintreeGatewayUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *BraintreeGatewayUpdateData) GetAttributes() BraintreeGatewayUpdateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BraintreeGatewayUpdateData) GetAttributesOk() (*BraintreeGatewayUpdateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BraintreeGatewayUpdateData) SetAttributes(v BraintreeGatewayUpdateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *BraintreeGatewayUpdateData) GetRelationships() BraintreeGatewayCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BraintreeGatewayUpdateData) GetRelationshipsOk() (*BraintreeGatewayCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BraintreeGatewayUpdateData) SetRelationships(v BraintreeGatewayCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BraintreeGatewayUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


