# AdyenGatewayCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "adyen_gateways"]
**Attributes** | [**AdyenGatewayCreateDataAttributes**](AdyenGatewayCreateDataAttributes.md) |  | 
**Relationships** | Pointer to [**AdyenGatewayCreateDataRelationships**](AdyenGatewayCreateDataRelationships.md) |  | [optional] 

## Methods

### NewAdyenGatewayCreateData

`func NewAdyenGatewayCreateData(type_ string, attributes AdyenGatewayCreateDataAttributes, ) *AdyenGatewayCreateData`

NewAdyenGatewayCreateData instantiates a new AdyenGatewayCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdyenGatewayCreateDataWithDefaults

`func NewAdyenGatewayCreateDataWithDefaults() *AdyenGatewayCreateData`

NewAdyenGatewayCreateDataWithDefaults instantiates a new AdyenGatewayCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AdyenGatewayCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdyenGatewayCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdyenGatewayCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *AdyenGatewayCreateData) GetAttributes() AdyenGatewayCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AdyenGatewayCreateData) GetAttributesOk() (*AdyenGatewayCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AdyenGatewayCreateData) SetAttributes(v AdyenGatewayCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *AdyenGatewayCreateData) GetRelationships() AdyenGatewayCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AdyenGatewayCreateData) GetRelationshipsOk() (*AdyenGatewayCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AdyenGatewayCreateData) SetRelationships(v AdyenGatewayCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AdyenGatewayCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


