# StockTransferCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTStockTransfers201ResponseDataAttributes**](POSTStockTransfers201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**StockTransferCreateDataRelationships**](StockTransferCreateDataRelationships.md) |  | [optional] 

## Methods

### NewStockTransferCreateData

`func NewStockTransferCreateData(type_ interface{}, attributes POSTStockTransfers201ResponseDataAttributes, ) *StockTransferCreateData`

NewStockTransferCreateData instantiates a new StockTransferCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockTransferCreateDataWithDefaults

`func NewStockTransferCreateDataWithDefaults() *StockTransferCreateData`

NewStockTransferCreateDataWithDefaults instantiates a new StockTransferCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StockTransferCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StockTransferCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StockTransferCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *StockTransferCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *StockTransferCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *StockTransferCreateData) GetAttributes() POSTStockTransfers201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *StockTransferCreateData) GetAttributesOk() (*POSTStockTransfers201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *StockTransferCreateData) SetAttributes(v POSTStockTransfers201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *StockTransferCreateData) GetRelationships() StockTransferCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *StockTransferCreateData) GetRelationshipsOk() (*StockTransferCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *StockTransferCreateData) SetRelationships(v StockTransferCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *StockTransferCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


