# StockLocationCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "stock_locations"]
**Attributes** | [**POSTStockLocations201ResponseDataAttributes**](POSTStockLocations201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTMerchants201ResponseDataRelationships**](POSTMerchants201ResponseDataRelationships.md) |  | [optional] 

## Methods

### NewStockLocationCreateData

`func NewStockLocationCreateData(type_ string, attributes POSTStockLocations201ResponseDataAttributes, ) *StockLocationCreateData`

NewStockLocationCreateData instantiates a new StockLocationCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockLocationCreateDataWithDefaults

`func NewStockLocationCreateDataWithDefaults() *StockLocationCreateData`

NewStockLocationCreateDataWithDefaults instantiates a new StockLocationCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StockLocationCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StockLocationCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StockLocationCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *StockLocationCreateData) GetAttributes() POSTStockLocations201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *StockLocationCreateData) GetAttributesOk() (*POSTStockLocations201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *StockLocationCreateData) SetAttributes(v POSTStockLocations201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *StockLocationCreateData) GetRelationships() POSTMerchants201ResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *StockLocationCreateData) GetRelationshipsOk() (*POSTMerchants201ResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *StockLocationCreateData) SetRelationships(v POSTMerchants201ResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *StockLocationCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


