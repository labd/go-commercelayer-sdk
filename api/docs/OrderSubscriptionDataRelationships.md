# OrderSubscriptionDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**AvalaraAccountDataRelationshipsMarkets**](AvalaraAccountDataRelationshipsMarkets.md) |  | [optional] 
**SourceOrder** | Pointer to [**AdyenPaymentDataRelationshipsOrder**](AdyenPaymentDataRelationshipsOrder.md) |  | [optional] 
**Customer** | Pointer to [**CouponRecipientDataRelationshipsCustomer**](CouponRecipientDataRelationshipsCustomer.md) |  | [optional] 
**OrderCopies** | Pointer to [**OrderSubscriptionDataRelationshipsOrderCopies**](OrderSubscriptionDataRelationshipsOrderCopies.md) |  | [optional] 
**Orders** | Pointer to [**AdyenPaymentDataRelationshipsOrder**](AdyenPaymentDataRelationshipsOrder.md) |  | [optional] 

## Methods

### NewOrderSubscriptionDataRelationships

`func NewOrderSubscriptionDataRelationships() *OrderSubscriptionDataRelationships`

NewOrderSubscriptionDataRelationships instantiates a new OrderSubscriptionDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSubscriptionDataRelationshipsWithDefaults

`func NewOrderSubscriptionDataRelationshipsWithDefaults() *OrderSubscriptionDataRelationships`

NewOrderSubscriptionDataRelationshipsWithDefaults instantiates a new OrderSubscriptionDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *OrderSubscriptionDataRelationships) GetMarket() AvalaraAccountDataRelationshipsMarkets`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *OrderSubscriptionDataRelationships) GetMarketOk() (*AvalaraAccountDataRelationshipsMarkets, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *OrderSubscriptionDataRelationships) SetMarket(v AvalaraAccountDataRelationshipsMarkets)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *OrderSubscriptionDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetSourceOrder

`func (o *OrderSubscriptionDataRelationships) GetSourceOrder() AdyenPaymentDataRelationshipsOrder`

GetSourceOrder returns the SourceOrder field if non-nil, zero value otherwise.

### GetSourceOrderOk

`func (o *OrderSubscriptionDataRelationships) GetSourceOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool)`

GetSourceOrderOk returns a tuple with the SourceOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOrder

`func (o *OrderSubscriptionDataRelationships) SetSourceOrder(v AdyenPaymentDataRelationshipsOrder)`

SetSourceOrder sets SourceOrder field to given value.

### HasSourceOrder

`func (o *OrderSubscriptionDataRelationships) HasSourceOrder() bool`

HasSourceOrder returns a boolean if a field has been set.

### GetCustomer

`func (o *OrderSubscriptionDataRelationships) GetCustomer() CouponRecipientDataRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *OrderSubscriptionDataRelationships) GetCustomerOk() (*CouponRecipientDataRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *OrderSubscriptionDataRelationships) SetCustomer(v CouponRecipientDataRelationshipsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *OrderSubscriptionDataRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetOrderCopies

`func (o *OrderSubscriptionDataRelationships) GetOrderCopies() OrderSubscriptionDataRelationshipsOrderCopies`

GetOrderCopies returns the OrderCopies field if non-nil, zero value otherwise.

### GetOrderCopiesOk

`func (o *OrderSubscriptionDataRelationships) GetOrderCopiesOk() (*OrderSubscriptionDataRelationshipsOrderCopies, bool)`

GetOrderCopiesOk returns a tuple with the OrderCopies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCopies

`func (o *OrderSubscriptionDataRelationships) SetOrderCopies(v OrderSubscriptionDataRelationshipsOrderCopies)`

SetOrderCopies sets OrderCopies field to given value.

### HasOrderCopies

`func (o *OrderSubscriptionDataRelationships) HasOrderCopies() bool`

HasOrderCopies returns a boolean if a field has been set.

### GetOrders

`func (o *OrderSubscriptionDataRelationships) GetOrders() AdyenPaymentDataRelationshipsOrder`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *OrderSubscriptionDataRelationships) GetOrdersOk() (*AdyenPaymentDataRelationshipsOrder, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *OrderSubscriptionDataRelationships) SetOrders(v AdyenPaymentDataRelationshipsOrder)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *OrderSubscriptionDataRelationships) HasOrders() bool`

HasOrders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


