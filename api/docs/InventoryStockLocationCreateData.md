# InventoryStockLocationCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**POSTInventoryStockLocations201ResponseDataAttributes**](POSTInventoryStockLocations201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**InventoryReturnLocationCreateDataRelationships**](InventoryReturnLocationCreateDataRelationships.md) |  | [optional] 

## Methods

### NewInventoryStockLocationCreateData

`func NewInventoryStockLocationCreateData(type_ string, attributes POSTInventoryStockLocations201ResponseDataAttributes, ) *InventoryStockLocationCreateData`

NewInventoryStockLocationCreateData instantiates a new InventoryStockLocationCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryStockLocationCreateDataWithDefaults

`func NewInventoryStockLocationCreateDataWithDefaults() *InventoryStockLocationCreateData`

NewInventoryStockLocationCreateDataWithDefaults instantiates a new InventoryStockLocationCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InventoryStockLocationCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InventoryStockLocationCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InventoryStockLocationCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *InventoryStockLocationCreateData) GetAttributes() POSTInventoryStockLocations201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InventoryStockLocationCreateData) GetAttributesOk() (*POSTInventoryStockLocations201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *InventoryStockLocationCreateData) SetAttributes(v POSTInventoryStockLocations201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *InventoryStockLocationCreateData) GetRelationships() InventoryReturnLocationCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *InventoryStockLocationCreateData) GetRelationshipsOk() (*InventoryReturnLocationCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *InventoryStockLocationCreateData) SetRelationships(v InventoryReturnLocationCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *InventoryStockLocationCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


