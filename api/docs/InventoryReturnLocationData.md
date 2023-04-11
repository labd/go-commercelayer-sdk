# InventoryReturnLocationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETInventoryReturnLocations200ResponseDataInnerAttributes**](GETInventoryReturnLocations200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**InventoryReturnLocationDataRelationships**](InventoryReturnLocationDataRelationships.md) |  | [optional] 

## Methods

### NewInventoryReturnLocationData

`func NewInventoryReturnLocationData(type_ interface{}, attributes GETInventoryReturnLocations200ResponseDataInnerAttributes, ) *InventoryReturnLocationData`

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

`func (o *InventoryReturnLocationData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InventoryReturnLocationData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InventoryReturnLocationData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *InventoryReturnLocationData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *InventoryReturnLocationData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *InventoryReturnLocationData) GetAttributes() GETInventoryReturnLocations200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InventoryReturnLocationData) GetAttributesOk() (*GETInventoryReturnLocations200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *InventoryReturnLocationData) SetAttributes(v GETInventoryReturnLocations200ResponseDataInnerAttributes)`

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


