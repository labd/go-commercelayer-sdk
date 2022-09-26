# InventoryReturnLocationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**InventoryReturnLocationDataAttributes**](InventoryReturnLocationDataAttributes.md) |  | 
**Relationships** | Pointer to [**InventoryReturnLocationDataRelationships**](InventoryReturnLocationDataRelationships.md) |  | [optional] 

## Methods

### NewInventoryReturnLocationData

`func NewInventoryReturnLocationData(type_ string, attributes InventoryReturnLocationDataAttributes, ) *InventoryReturnLocationData`

NewInventoryReturnLocationData instantiates a new InventoryReturnLocationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryReturnLocationDataWithDefaults

`func NewInventoryReturnLocationDataWithDefaults() *InventoryReturnLocationData`

NewInventoryReturnLocationDataWithDefaults instantiates a new InventoryReturnLocationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InventoryReturnLocationData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InventoryReturnLocationData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InventoryReturnLocationData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *InventoryReturnLocationData) GetAttributes() InventoryReturnLocationDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InventoryReturnLocationData) GetAttributesOk() (*InventoryReturnLocationDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *InventoryReturnLocationData) SetAttributes(v InventoryReturnLocationDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *InventoryReturnLocationData) GetRelationships() InventoryReturnLocationDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *InventoryReturnLocationData) GetRelationshipsOk() (*InventoryReturnLocationDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *InventoryReturnLocationData) SetRelationships(v InventoryReturnLocationDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *InventoryReturnLocationData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


