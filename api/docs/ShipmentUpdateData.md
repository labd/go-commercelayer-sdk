# ShipmentUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHShipmentsShipmentId200ResponseDataAttributes**](PATCHShipmentsShipmentId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**ShipmentUpdateDataRelationships**](ShipmentUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewShipmentUpdateData

`func NewShipmentUpdateData(type_ interface{}, id interface{}, attributes PATCHShipmentsShipmentId200ResponseDataAttributes, ) *ShipmentUpdateData`

NewShipmentUpdateData instantiates a new ShipmentUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentUpdateDataWithDefaults

`func NewShipmentUpdateDataWithDefaults() *ShipmentUpdateData`

NewShipmentUpdateDataWithDefaults instantiates a new ShipmentUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ShipmentUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShipmentUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShipmentUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ShipmentUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ShipmentUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *ShipmentUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShipmentUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShipmentUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *ShipmentUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ShipmentUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *ShipmentUpdateData) GetAttributes() PATCHShipmentsShipmentId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ShipmentUpdateData) GetAttributesOk() (*PATCHShipmentsShipmentId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ShipmentUpdateData) SetAttributes(v PATCHShipmentsShipmentId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ShipmentUpdateData) GetRelationships() ShipmentUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ShipmentUpdateData) GetRelationshipsOk() (*ShipmentUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ShipmentUpdateData) SetRelationships(v ShipmentUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ShipmentUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


