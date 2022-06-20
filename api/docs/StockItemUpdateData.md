# StockItemUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "stock_items"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**StockItemUpdateDataAttributes**](StockItemUpdateDataAttributes.md) |  | 
**Relationships** | Pointer to [**StockItemUpdateDataRelationships**](StockItemUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewStockItemUpdateData

`func NewStockItemUpdateData(type_ string, id string, attributes StockItemUpdateDataAttributes, ) *StockItemUpdateData`

NewStockItemUpdateData instantiates a new StockItemUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockItemUpdateDataWithDefaults

`func NewStockItemUpdateDataWithDefaults() *StockItemUpdateData`

NewStockItemUpdateDataWithDefaults instantiates a new StockItemUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StockItemUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StockItemUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StockItemUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *StockItemUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StockItemUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StockItemUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *StockItemUpdateData) GetAttributes() StockItemUpdateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *StockItemUpdateData) GetAttributesOk() (*StockItemUpdateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *StockItemUpdateData) SetAttributes(v StockItemUpdateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *StockItemUpdateData) GetRelationships() StockItemUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *StockItemUpdateData) GetRelationshipsOk() (*StockItemUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *StockItemUpdateData) SetRelationships(v StockItemUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *StockItemUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


