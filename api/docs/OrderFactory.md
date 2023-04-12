# OrderFactory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**OrderFactoryData**](OrderFactoryData.md) |  | [optional] 

## Methods

### NewOrderFactory

`func NewOrderFactory() *OrderFactory`

NewOrderFactory instantiates a new OrderFactory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderFactoryWithDefaults

`func NewOrderFactoryWithDefaults() *OrderFactory`

NewOrderFactoryWithDefaults instantiates a new OrderFactory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *OrderFactory) GetData() OrderFactoryData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OrderFactory) GetDataOk() (*OrderFactoryData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OrderFactory) SetData(v OrderFactoryData)`

SetData sets Data field to given value.

### HasData

`func (o *OrderFactory) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


