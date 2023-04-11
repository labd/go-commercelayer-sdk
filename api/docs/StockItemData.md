# StockItemData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETStockItems200ResponseDataInnerAttributes**](GETStockItems200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**StockItemDataRelationships**](StockItemDataRelationships.md) |  | [optional] 

## Methods

### NewStockItemData

`func NewStockItemData(type_ interface{}, attributes GETStockItems200ResponseDataInnerAttributes, ) *StockItemData`

NewStockItemData instantiates a new StockItemData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockItemDataWithDefaults

`func NewStockItemDataWithDefaults() *StockItemData`

NewStockItemDataWithDefaults instantiates a new StockItemData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StockItemData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StockItemData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StockItemData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *StockItemData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *StockItemData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *StockItemData) GetAttributes() GETStockItems200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *StockItemData) GetAttributesOk() (*GETStockItems200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *StockItemData) SetAttributes(v GETStockItems200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *StockItemData) GetRelationships() StockItemDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *StockItemData) GetRelationshipsOk() (*StockItemDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *StockItemData) SetRelationships(v StockItemDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *StockItemData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


