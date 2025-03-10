# StockLineItemUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHStockLineItemsStockLineItemId200ResponseDataAttributes**](PATCHStockLineItemsStockLineItemId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**StockLineItemCreateDataRelationships**](StockLineItemCreateDataRelationships.md) |  | [optional] 

## Methods

### NewStockLineItemUpdateData

`func NewStockLineItemUpdateData(type_ interface{}, id interface{}, attributes PATCHStockLineItemsStockLineItemId200ResponseDataAttributes, ) *StockLineItemUpdateData`

NewStockLineItemUpdateData instantiates a new StockLineItemUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockLineItemUpdateDataWithDefaults

`func NewStockLineItemUpdateDataWithDefaults() *StockLineItemUpdateData`

NewStockLineItemUpdateDataWithDefaults instantiates a new StockLineItemUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StockLineItemUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StockLineItemUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StockLineItemUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *StockLineItemUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *StockLineItemUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *StockLineItemUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StockLineItemUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StockLineItemUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *StockLineItemUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *StockLineItemUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *StockLineItemUpdateData) GetAttributes() PATCHStockLineItemsStockLineItemId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *StockLineItemUpdateData) GetAttributesOk() (*PATCHStockLineItemsStockLineItemId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *StockLineItemUpdateData) SetAttributes(v PATCHStockLineItemsStockLineItemId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *StockLineItemUpdateData) GetRelationships() StockLineItemCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *StockLineItemUpdateData) GetRelationshipsOk() (*StockLineItemCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *StockLineItemUpdateData) SetRelationships(v StockLineItemCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *StockLineItemUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


