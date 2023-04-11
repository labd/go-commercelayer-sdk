# ShippingWeightTierUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHShippingWeightTiersShippingWeightTierId200ResponseDataAttributes**](PATCHShippingWeightTiersShippingWeightTierId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**ShipmentUpdateDataRelationships**](ShipmentUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewShippingWeightTierUpdateData

`func NewShippingWeightTierUpdateData(type_ interface{}, id interface{}, attributes PATCHShippingWeightTiersShippingWeightTierId200ResponseDataAttributes, ) *ShippingWeightTierUpdateData`

NewShippingWeightTierUpdateData instantiates a new ShippingWeightTierUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingWeightTierUpdateDataWithDefaults

`func NewShippingWeightTierUpdateDataWithDefaults() *ShippingWeightTierUpdateData`

NewShippingWeightTierUpdateDataWithDefaults instantiates a new ShippingWeightTierUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ShippingWeightTierUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShippingWeightTierUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShippingWeightTierUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ShippingWeightTierUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ShippingWeightTierUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *ShippingWeightTierUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShippingWeightTierUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShippingWeightTierUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *ShippingWeightTierUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ShippingWeightTierUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *ShippingWeightTierUpdateData) GetAttributes() PATCHShippingWeightTiersShippingWeightTierId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ShippingWeightTierUpdateData) GetAttributesOk() (*PATCHShippingWeightTiersShippingWeightTierId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ShippingWeightTierUpdateData) SetAttributes(v PATCHShippingWeightTiersShippingWeightTierId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ShippingWeightTierUpdateData) GetRelationships() ShipmentUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ShippingWeightTierUpdateData) GetRelationshipsOk() (*ShipmentUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ShippingWeightTierUpdateData) SetRelationships(v ShipmentUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ShippingWeightTierUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


