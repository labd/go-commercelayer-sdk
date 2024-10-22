# StockReservationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETStockReservationsStockReservationId200ResponseDataAttributes**](GETStockReservationsStockReservationId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**StockReservationDataRelationships**](StockReservationDataRelationships.md) |  | [optional] 

## Methods

### NewStockReservationData

`func NewStockReservationData(type_ interface{}, attributes GETStockReservationsStockReservationId200ResponseDataAttributes, ) *StockReservationData`

NewStockReservationData instantiates a new StockReservationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockReservationDataWithDefaults

`func NewStockReservationDataWithDefaults() *StockReservationData`

NewStockReservationDataWithDefaults instantiates a new StockReservationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StockReservationData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StockReservationData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StockReservationData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *StockReservationData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *StockReservationData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *StockReservationData) GetAttributes() GETStockReservationsStockReservationId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *StockReservationData) GetAttributesOk() (*GETStockReservationsStockReservationId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *StockReservationData) SetAttributes(v GETStockReservationsStockReservationId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *StockReservationData) GetRelationships() StockReservationDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *StockReservationData) GetRelationshipsOk() (*StockReservationDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *StockReservationData) SetRelationships(v StockReservationDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *StockReservationData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


