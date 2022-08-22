# ShippingWeightTierUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "shipping_weight_tiers"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**PATCHShippingWeightTiersShippingWeightTierId200ResponseDataAttributes**](PATCHShippingWeightTiersShippingWeightTierId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**PATCHShipmentsShipmentId200ResponseDataRelationships**](PATCHShipmentsShipmentId200ResponseDataRelationships.md) |  | [optional] 

## Methods

### NewShippingWeightTierUpdateData

`func NewShippingWeightTierUpdateData(type_ string, id string, attributes PATCHShippingWeightTiersShippingWeightTierId200ResponseDataAttributes, ) *ShippingWeightTierUpdateData`

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

`func (o *ShippingWeightTierUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShippingWeightTierUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShippingWeightTierUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ShippingWeightTierUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShippingWeightTierUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShippingWeightTierUpdateData) SetId(v string)`

SetId sets Id field to given value.


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

`func (o *ShippingWeightTierUpdateData) GetRelationships() PATCHShipmentsShipmentId200ResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ShippingWeightTierUpdateData) GetRelationshipsOk() (*PATCHShipmentsShipmentId200ResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ShippingWeightTierUpdateData) SetRelationships(v PATCHShipmentsShipmentId200ResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ShippingWeightTierUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


