# AdyenGatewayUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "adyen_gateways"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes**](PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**AdyenGatewayCreateDataRelationships**](AdyenGatewayCreateDataRelationships.md) |  | [optional] 

## Methods

### NewAdyenGatewayUpdateData

`func NewAdyenGatewayUpdateData(type_ string, id string, attributes PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes, ) *AdyenGatewayUpdateData`

NewAdyenGatewayUpdateData instantiates a new AdyenGatewayUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdyenGatewayUpdateDataWithDefaults

`func NewAdyenGatewayUpdateDataWithDefaults() *AdyenGatewayUpdateData`

NewAdyenGatewayUpdateDataWithDefaults instantiates a new AdyenGatewayUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AdyenGatewayUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdyenGatewayUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdyenGatewayUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *AdyenGatewayUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdyenGatewayUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdyenGatewayUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *AdyenGatewayUpdateData) GetAttributes() PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AdyenGatewayUpdateData) GetAttributesOk() (*PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AdyenGatewayUpdateData) SetAttributes(v PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *AdyenGatewayUpdateData) GetRelationships() AdyenGatewayCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AdyenGatewayUpdateData) GetRelationshipsOk() (*AdyenGatewayCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AdyenGatewayUpdateData) SetRelationships(v AdyenGatewayCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AdyenGatewayUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


