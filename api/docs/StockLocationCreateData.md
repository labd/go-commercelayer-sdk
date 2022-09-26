# StockLocationCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**StockLocationCreateDataAttributes**](StockLocationCreateDataAttributes.md) |  | 
**Relationships** | Pointer to [**MerchantCreateDataRelationships**](MerchantCreateDataRelationships.md) |  | [optional] 

## Methods

### NewStockLocationCreateData

`func NewStockLocationCreateData(type_ string, attributes StockLocationCreateDataAttributes, ) *StockLocationCreateData`

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

`func (o *StockLocationCreateData) GetAttributes() StockLocationCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *StockLocationCreateData) GetAttributesOk() (*StockLocationCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *StockLocationCreateData) SetAttributes(v StockLocationCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *StockLocationCreateData) GetRelationships() MerchantCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *StockLocationCreateData) GetRelationshipsOk() (*MerchantCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *StockLocationCreateData) SetRelationships(v MerchantCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *StockLocationCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


