# StockTransferUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**StockTransferUpdateDataAttributes**](StockTransferUpdateDataAttributes.md) |  | 
**Relationships** | Pointer to [**StockTransferUpdateDataRelationships**](StockTransferUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewStockTransferUpdateData

`func NewStockTransferUpdateData(type_ string, id string, attributes StockTransferUpdateDataAttributes, ) *StockTransferUpdateData`

NewStockTransferUpdateData instantiates a new StockTransferUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockTransferUpdateDataWithDefaults

`func NewStockTransferUpdateDataWithDefaults() *StockTransferUpdateData`

NewStockTransferUpdateDataWithDefaults instantiates a new StockTransferUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StockTransferUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StockTransferUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StockTransferUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *StockTransferUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StockTransferUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StockTransferUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *StockTransferUpdateData) GetAttributes() StockTransferUpdateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *StockTransferUpdateData) GetAttributesOk() (*StockTransferUpdateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *StockTransferUpdateData) SetAttributes(v StockTransferUpdateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *StockTransferUpdateData) GetRelationships() StockTransferUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *StockTransferUpdateData) GetRelationshipsOk() (*StockTransferUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *StockTransferUpdateData) SetRelationships(v StockTransferUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *StockTransferUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


