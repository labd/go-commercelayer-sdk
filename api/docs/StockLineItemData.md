# StockLineItemData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETStockLineItemsStockLineItemId200ResponseDataAttributes**](GETStockLineItemsStockLineItemId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**StockLineItemDataRelationships**](StockLineItemDataRelationships.md) |  | [optional] 

## Methods

### NewStockLineItemData

`func NewStockLineItemData(type_ interface{}, attributes GETStockLineItemsStockLineItemId200ResponseDataAttributes, ) *StockLineItemData`

NewStockLineItemData instantiates a new StockLineItemData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockLineItemDataWithDefaults

`func NewStockLineItemDataWithDefaults() *StockLineItemData`

NewStockLineItemDataWithDefaults instantiates a new StockLineItemData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StockLineItemData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StockLineItemData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StockLineItemData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *StockLineItemData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *StockLineItemData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *StockLineItemData) GetAttributes() GETStockLineItemsStockLineItemId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *StockLineItemData) GetAttributesOk() (*GETStockLineItemsStockLineItemId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *StockLineItemData) SetAttributes(v GETStockLineItemsStockLineItemId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *StockLineItemData) GetRelationships() StockLineItemDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *StockLineItemData) GetRelationshipsOk() (*StockLineItemDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *StockLineItemData) SetRelationships(v StockLineItemDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *StockLineItemData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


