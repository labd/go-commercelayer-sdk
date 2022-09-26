# OrderSubscriptionResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**BillingInfoValidationRuleResponseDataRelationshipsMarket**](BillingInfoValidationRuleResponseDataRelationshipsMarket.md) |  | [optional] 
**SourceOrder** | Pointer to [**OrderCopyResponseDataRelationshipsSourceOrder**](OrderCopyResponseDataRelationshipsSourceOrder.md) |  | [optional] 
**Customer** | Pointer to [**CouponRecipientResponseDataRelationshipsCustomer**](CouponRecipientResponseDataRelationshipsCustomer.md) |  | [optional] 
**OrderCopies** | Pointer to [**OrderSubscriptionResponseDataRelationshipsOrderCopies**](OrderSubscriptionResponseDataRelationshipsOrderCopies.md) |  | [optional] 
**Orders** | Pointer to [**CustomerResponseDataRelationshipsOrders**](CustomerResponseDataRelationshipsOrders.md) |  | [optional] 
**Events** | Pointer to [**CustomerAddressResponseDataRelationshipsEvents**](CustomerAddressResponseDataRelationshipsEvents.md) |  | [optional] 

## Methods

### NewOrderSubscriptionResponseDataRelationships

`func NewOrderSubscriptionResponseDataRelationships() *OrderSubscriptionResponseDataRelationships`

NewOrderSubscriptionResponseDataRelationships instantiates a new OrderSubscriptionResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSubscriptionResponseDataRelationshipsWithDefaults

`func NewOrderSubscriptionResponseDataRelationshipsWithDefaults() *OrderSubscriptionResponseDataRelationships`

NewOrderSubscriptionResponseDataRelationshipsWithDefaults instantiates a new OrderSubscriptionResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *OrderSubscriptionResponseDataRelationships) GetMarket() BillingInfoValidationRuleResponseDataRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *OrderSubscriptionResponseDataRelationships) GetMarketOk() (*BillingInfoValidationRuleResponseDataRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *OrderSubscriptionResponseDataRelationships) SetMarket(v BillingInfoValidationRuleResponseDataRelationshipsMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *OrderSubscriptionResponseDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetSourceOrder

`func (o *OrderSubscriptionResponseDataRelationships) GetSourceOrder() OrderCopyResponseDataRelationshipsSourceOrder`

GetSourceOrder returns the SourceOrder field if non-nil, zero value otherwise.

### GetSourceOrderOk

`func (o *OrderSubscriptionResponseDataRelationships) GetSourceOrderOk() (*OrderCopyResponseDataRelationshipsSourceOrder, bool)`

GetSourceOrderOk returns a tuple with the SourceOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOrder

`func (o *OrderSubscriptionResponseDataRelationships) SetSourceOrder(v OrderCopyResponseDataRelationshipsSourceOrder)`

SetSourceOrder sets SourceOrder field to given value.

### HasSourceOrder

`func (o *OrderSubscriptionResponseDataRelationships) HasSourceOrder() bool`

HasSourceOrder returns a boolean if a field has been set.

### GetCustomer

`func (o *OrderSubscriptionResponseDataRelationships) GetCustomer() CouponRecipientResponseDataRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *OrderSubscriptionResponseDataRelationships) GetCustomerOk() (*CouponRecipientResponseDataRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *OrderSubscriptionResponseDataRelationships) SetCustomer(v CouponRecipientResponseDataRelationshipsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *OrderSubscriptionResponseDataRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetOrderCopies

`func (o *OrderSubscriptionResponseDataRelationships) GetOrderCopies() OrderSubscriptionResponseDataRelationshipsOrderCopies`

GetOrderCopies returns the OrderCopies field if non-nil, zero value otherwise.

### GetOrderCopiesOk

`func (o *OrderSubscriptionResponseDataRelationships) GetOrderCopiesOk() (*OrderSubscriptionResponseDataRelationshipsOrderCopies, bool)`

GetOrderCopiesOk returns a tuple with the OrderCopies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCopies

`func (o *OrderSubscriptionResponseDataRelationships) SetOrderCopies(v OrderSubscriptionResponseDataRelationshipsOrderCopies)`

SetOrderCopies sets OrderCopies field to given value.

### HasOrderCopies

`func (o *OrderSubscriptionResponseDataRelationships) HasOrderCopies() bool`

HasOrderCopies returns a boolean if a field has been set.

### GetOrders

`func (o *OrderSubscriptionResponseDataRelationships) GetOrders() CustomerResponseDataRelationshipsOrders`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *OrderSubscriptionResponseDataRelationships) GetOrdersOk() (*CustomerResponseDataRelationshipsOrders, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *OrderSubscriptionResponseDataRelationships) SetOrders(v CustomerResponseDataRelationshipsOrders)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *OrderSubscriptionResponseDataRelationships) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetEvents

`func (o *OrderSubscriptionResponseDataRelationships) GetEvents() CustomerAddressResponseDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *OrderSubscriptionResponseDataRelationships) GetEventsOk() (*CustomerAddressResponseDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *OrderSubscriptionResponseDataRelationships) SetEvents(v CustomerAddressResponseDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *OrderSubscriptionResponseDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


