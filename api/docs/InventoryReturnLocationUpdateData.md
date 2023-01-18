# InventoryReturnLocationUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**PATCHInventoryReturnLocationsInventoryReturnLocationId200ResponseDataAttributes**](PATCHInventoryReturnLocationsInventoryReturnLocationId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**InventoryReturnLocationUpdateDataRelationships**](InventoryReturnLocationUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewInventoryReturnLocationUpdateData

`func NewInventoryReturnLocationUpdateData(type_ string, id string, attributes PATCHInventoryReturnLocationsInventoryReturnLocationId200ResponseDataAttributes, ) *InventoryReturnLocationUpdateData`

NewInventoryReturnLocationUpdateData instantiates a new InventoryReturnLocationUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryReturnLocationUpdateDataWithDefaults

`func NewInventoryReturnLocationUpdateDataWithDefaults() *InventoryReturnLocationUpdateData`

NewInventoryReturnLocationUpdateDataWithDefaults instantiates a new InventoryReturnLocationUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InventoryReturnLocationUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InventoryReturnLocationUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InventoryReturnLocationUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *InventoryReturnLocationUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InventoryReturnLocationUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InventoryReturnLocationUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *InventoryReturnLocationUpdateData) GetAttributes() PATCHInventoryReturnLocationsInventoryReturnLocationId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InventoryReturnLocationUpdateData) GetAttributesOk() (*PATCHInventoryReturnLocationsInventoryReturnLocationId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *InventoryReturnLocationUpdateData) SetAttributes(v PATCHInventoryReturnLocationsInventoryReturnLocationId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *InventoryReturnLocationUpdateData) GetRelationships() InventoryReturnLocationUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *InventoryReturnLocationUpdateData) GetRelationshipsOk() (*InventoryReturnLocationUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *InventoryReturnLocationUpdateData) SetRelationships(v InventoryReturnLocationUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *InventoryReturnLocationUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


