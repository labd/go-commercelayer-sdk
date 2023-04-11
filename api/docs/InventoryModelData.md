# InventoryModelData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETInventoryModels200ResponseDataInnerAttributes**](GETInventoryModels200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**InventoryModelDataRelationships**](InventoryModelDataRelationships.md) |  | [optional] 

## Methods

### NewInventoryModelData

`func NewInventoryModelData(type_ interface{}, attributes GETInventoryModels200ResponseDataInnerAttributes, ) *InventoryModelData`

NewInventoryModelData instantiates a new InventoryModelData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryModelDataWithDefaults

`func NewInventoryModelDataWithDefaults() *InventoryModelData`

NewInventoryModelDataWithDefaults instantiates a new InventoryModelData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InventoryModelData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InventoryModelData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InventoryModelData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *InventoryModelData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *InventoryModelData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *InventoryModelData) GetAttributes() GETInventoryModels200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InventoryModelData) GetAttributesOk() (*GETInventoryModels200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *InventoryModelData) SetAttributes(v GETInventoryModels200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *InventoryModelData) GetRelationships() InventoryModelDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *InventoryModelData) GetRelationshipsOk() (*InventoryModelDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *InventoryModelData) SetRelationships(v InventoryModelDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *InventoryModelData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


