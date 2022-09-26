# InventoryStockLocationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**InventoryStockLocationDataAttributes**](InventoryStockLocationDataAttributes.md) |  | 
**Relationships** | Pointer to [**InventoryReturnLocationDataRelationships**](InventoryReturnLocationDataRelationships.md) |  | [optional] 

## Methods

### NewInventoryStockLocationData

`func NewInventoryStockLocationData(type_ string, attributes InventoryStockLocationDataAttributes, ) *InventoryStockLocationData`

NewInventoryStockLocationData instantiates a new InventoryStockLocationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryStockLocationDataWithDefaults

`func NewInventoryStockLocationDataWithDefaults() *InventoryStockLocationData`

NewInventoryStockLocationDataWithDefaults instantiates a new InventoryStockLocationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InventoryStockLocationData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InventoryStockLocationData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InventoryStockLocationData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *InventoryStockLocationData) GetAttributes() InventoryStockLocationDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InventoryStockLocationData) GetAttributesOk() (*InventoryStockLocationDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *InventoryStockLocationData) SetAttributes(v InventoryStockLocationDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *InventoryStockLocationData) GetRelationships() InventoryReturnLocationDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *InventoryStockLocationData) GetRelationshipsOk() (*InventoryReturnLocationDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *InventoryStockLocationData) SetRelationships(v InventoryReturnLocationDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *InventoryStockLocationData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


