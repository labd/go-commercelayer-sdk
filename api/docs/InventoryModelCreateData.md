# InventoryModelCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**InventoryModelCreateDataAttributes**](InventoryModelCreateDataAttributes.md) |  | 
**Relationships** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewInventoryModelCreateData

`func NewInventoryModelCreateData(type_ string, attributes InventoryModelCreateDataAttributes, ) *InventoryModelCreateData`

NewInventoryModelCreateData instantiates a new InventoryModelCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryModelCreateDataWithDefaults

`func NewInventoryModelCreateDataWithDefaults() *InventoryModelCreateData`

NewInventoryModelCreateDataWithDefaults instantiates a new InventoryModelCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InventoryModelCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InventoryModelCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InventoryModelCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *InventoryModelCreateData) GetAttributes() InventoryModelCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InventoryModelCreateData) GetAttributesOk() (*InventoryModelCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *InventoryModelCreateData) SetAttributes(v InventoryModelCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *InventoryModelCreateData) GetRelationships() map[string]interface{}`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *InventoryModelCreateData) GetRelationshipsOk() (*map[string]interface{}, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *InventoryModelCreateData) SetRelationships(v map[string]interface{})`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *InventoryModelCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


