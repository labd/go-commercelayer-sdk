# StockLineItemCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTStockLineItems201ResponseDataAttributes**](POSTStockLineItems201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**StockLineItemCreateDataRelationships**](StockLineItemCreateDataRelationships.md) |  | [optional] 

## Methods

### NewStockLineItemCreateData

`func NewStockLineItemCreateData(type_ interface{}, attributes POSTStockLineItems201ResponseDataAttributes, ) *StockLineItemCreateData`

NewStockLineItemCreateData instantiates a new StockLineItemCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockLineItemCreateDataWithDefaults

`func NewStockLineItemCreateDataWithDefaults() *StockLineItemCreateData`

NewStockLineItemCreateDataWithDefaults instantiates a new StockLineItemCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StockLineItemCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StockLineItemCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StockLineItemCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *StockLineItemCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *StockLineItemCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *StockLineItemCreateData) GetAttributes() POSTStockLineItems201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *StockLineItemCreateData) GetAttributesOk() (*POSTStockLineItems201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *StockLineItemCreateData) SetAttributes(v POSTStockLineItems201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *StockLineItemCreateData) GetRelationships() StockLineItemCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *StockLineItemCreateData) GetRelationshipsOk() (*StockLineItemCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *StockLineItemCreateData) SetRelationships(v StockLineItemCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *StockLineItemCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


