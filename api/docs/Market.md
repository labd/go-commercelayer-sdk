# Market

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**MarketData**](MarketData.md) |  | [optional] 

## Methods

### NewMarket

`func NewMarket() *Market`

NewMarket instantiates a new Market object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketWithDefaults

`func NewMarketWithDefaults() *Market`

NewMarketWithDefaults instantiates a new Market object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Market) GetData() MarketData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Market) GetDataOk() (*MarketData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Market) SetData(v MarketData)`

SetData sets Data field to given value.

### HasData

`func (o *Market) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


