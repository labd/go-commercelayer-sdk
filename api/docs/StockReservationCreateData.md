# StockReservationCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTStockReservations201ResponseDataAttributes**](POSTStockReservations201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**StockReservationCreateDataRelationships**](StockReservationCreateDataRelationships.md) |  | [optional] 

## Methods

### NewStockReservationCreateData

`func NewStockReservationCreateData(type_ interface{}, attributes POSTStockReservations201ResponseDataAttributes, ) *StockReservationCreateData`

NewStockReservationCreateData instantiates a new StockReservationCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockReservationCreateDataWithDefaults

`func NewStockReservationCreateDataWithDefaults() *StockReservationCreateData`

NewStockReservationCreateDataWithDefaults instantiates a new StockReservationCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StockReservationCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StockReservationCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StockReservationCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *StockReservationCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *StockReservationCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *StockReservationCreateData) GetAttributes() POSTStockReservations201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *StockReservationCreateData) GetAttributesOk() (*POSTStockReservations201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *StockReservationCreateData) SetAttributes(v POSTStockReservations201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *StockReservationCreateData) GetRelationships() StockReservationCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *StockReservationCreateData) GetRelationshipsOk() (*StockReservationCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *StockReservationCreateData) SetRelationships(v StockReservationCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *StockReservationCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


