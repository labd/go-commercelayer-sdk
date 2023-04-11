# ReturnCreateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | [**AdyenPaymentCreateDataRelationshipsOrder**](AdyenPaymentCreateDataRelationshipsOrder.md) |  | 
**StockLocation** | Pointer to [**DeliveryLeadTimeCreateDataRelationshipsStockLocation**](DeliveryLeadTimeCreateDataRelationshipsStockLocation.md) |  | [optional] 

## Methods

### NewReturnCreateDataRelationships

`func NewReturnCreateDataRelationships(order AdyenPaymentCreateDataRelationshipsOrder, ) *ReturnCreateDataRelationships`

NewReturnCreateDataRelationships instantiates a new ReturnCreateDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnCreateDataRelationshipsWithDefaults

`func NewReturnCreateDataRelationshipsWithDefaults() *ReturnCreateDataRelationships`

NewReturnCreateDataRelationshipsWithDefaults instantiates a new ReturnCreateDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *ReturnCreateDataRelationships) GetOrder() AdyenPaymentCreateDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ReturnCreateDataRelationships) GetOrderOk() (*AdyenPaymentCreateDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ReturnCreateDataRelationships) SetOrder(v AdyenPaymentCreateDataRelationshipsOrder)`

SetOrder sets Order field to given value.


### GetStockLocation

`func (o *ReturnCreateDataRelationships) GetStockLocation() DeliveryLeadTimeCreateDataRelationshipsStockLocation`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *ReturnCreateDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeCreateDataRelationshipsStockLocation, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *ReturnCreateDataRelationships) SetStockLocation(v DeliveryLeadTimeCreateDataRelationshipsStockLocation)`

SetStockLocation sets StockLocation field to given value.

### HasStockLocation

`func (o *ReturnCreateDataRelationships) HasStockLocation() bool`

HasStockLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


