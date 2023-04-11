# POSTReturnsRequestDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | [**POSTAdyenPaymentsRequestDataRelationshipsOrder**](POSTAdyenPaymentsRequestDataRelationshipsOrder.md) |  | 
**StockLocation** | Pointer to [**POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation**](POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation.md) |  | [optional] 

## Methods

### NewPOSTReturnsRequestDataRelationships

`func NewPOSTReturnsRequestDataRelationships(order POSTAdyenPaymentsRequestDataRelationshipsOrder, ) *POSTReturnsRequestDataRelationships`

NewPOSTReturnsRequestDataRelationships instantiates a new POSTReturnsRequestDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTReturnsRequestDataRelationshipsWithDefaults

`func NewPOSTReturnsRequestDataRelationshipsWithDefaults() *POSTReturnsRequestDataRelationships`

NewPOSTReturnsRequestDataRelationshipsWithDefaults instantiates a new POSTReturnsRequestDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *POSTReturnsRequestDataRelationships) GetOrder() POSTAdyenPaymentsRequestDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *POSTReturnsRequestDataRelationships) GetOrderOk() (*POSTAdyenPaymentsRequestDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *POSTReturnsRequestDataRelationships) SetOrder(v POSTAdyenPaymentsRequestDataRelationshipsOrder)`

SetOrder sets Order field to given value.


### GetStockLocation

`func (o *POSTReturnsRequestDataRelationships) GetStockLocation() POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *POSTReturnsRequestDataRelationships) GetStockLocationOk() (*POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *POSTReturnsRequestDataRelationships) SetStockLocation(v POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation)`

SetStockLocation sets StockLocation field to given value.

### HasStockLocation

`func (o *POSTReturnsRequestDataRelationships) HasStockLocation() bool`

HasStockLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


