# AdyenGatewayUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes**](PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**AdyenGatewayCreateDataRelationships**](AdyenGatewayCreateDataRelationships.md) |  | [optional] 

## Methods

### NewAdyenGatewayUpdateData

`func NewAdyenGatewayUpdateData(type_ interface{}, id interface{}, attributes PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes, ) *AdyenGatewayUpdateData`

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

`func (o *AdyenGatewayUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdyenGatewayUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdyenGatewayUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *AdyenGatewayUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AdyenGatewayUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *AdyenGatewayUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdyenGatewayUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdyenGatewayUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *AdyenGatewayUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AdyenGatewayUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
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


