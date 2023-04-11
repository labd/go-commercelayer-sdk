# GETOrderSubscriptions200ResponseDataInnerRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket**](GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket.md) |  | [optional] 
**SubscriptionModel** | Pointer to [**GETMarkets200ResponseDataInnerRelationshipsSubscriptionModel**](GETMarkets200ResponseDataInnerRelationshipsSubscriptionModel.md) |  | [optional] 
**SourceOrder** | Pointer to [**GETOrderCopies200ResponseDataInnerRelationshipsSourceOrder**](GETOrderCopies200ResponseDataInnerRelationshipsSourceOrder.md) |  | [optional] 
**Customer** | Pointer to [**GETCouponRecipients200ResponseDataInnerRelationshipsCustomer**](GETCouponRecipients200ResponseDataInnerRelationshipsCustomer.md) |  | [optional] 
**CustomerPaymentSource** | Pointer to [**GETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSource**](GETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSource.md) |  | [optional] 
**OrderSubscriptionItems** | Pointer to [**GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderSubscriptionItems**](GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderSubscriptionItems.md) |  | [optional] 
**OrderFactories** | Pointer to [**GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderFactories**](GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderFactories.md) |  | [optional] 
**OrderCopies** | Pointer to [**GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies**](GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies.md) |  | [optional] 
**RecurringOrderCopies** | Pointer to [**GETOrderSubscriptions200ResponseDataInnerRelationshipsRecurringOrderCopies**](GETOrderSubscriptions200ResponseDataInnerRelationshipsRecurringOrderCopies.md) |  | [optional] 
**Orders** | Pointer to [**GETCustomers200ResponseDataInnerRelationshipsOrders**](GETCustomers200ResponseDataInnerRelationshipsOrders.md) |  | [optional] 
**Events** | Pointer to [**GETAuthorizations200ResponseDataInnerRelationshipsEvents**](GETAuthorizations200ResponseDataInnerRelationshipsEvents.md) |  | [optional] 

## Methods

### NewGETOrderSubscriptions200ResponseDataInnerRelationships

`func NewGETOrderSubscriptions200ResponseDataInnerRelationships() *GETOrderSubscriptions200ResponseDataInnerRelationships`

NewGETOrderSubscriptions200ResponseDataInnerRelationships instantiates a new GETOrderSubscriptions200ResponseDataInnerRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETOrderSubscriptions200ResponseDataInnerRelationshipsWithDefaults

`func NewGETOrderSubscriptions200ResponseDataInnerRelationshipsWithDefaults() *GETOrderSubscriptions200ResponseDataInnerRelationships`

NewGETOrderSubscriptions200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETOrderSubscriptions200ResponseDataInnerRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetMarket() GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetMarketOk() (*GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) SetMarket(v GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetSubscriptionModel

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetSubscriptionModel() GETMarkets200ResponseDataInnerRelationshipsSubscriptionModel`

GetSubscriptionModel returns the SubscriptionModel field if non-nil, zero value otherwise.

### GetSubscriptionModelOk

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetSubscriptionModelOk() (*GETMarkets200ResponseDataInnerRelationshipsSubscriptionModel, bool)`

GetSubscriptionModelOk returns a tuple with the SubscriptionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionModel

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) SetSubscriptionModel(v GETMarkets200ResponseDataInnerRelationshipsSubscriptionModel)`

SetSubscriptionModel sets SubscriptionModel field to given value.

### HasSubscriptionModel

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) HasSubscriptionModel() bool`

HasSubscriptionModel returns a boolean if a field has been set.

### GetSourceOrder

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetSourceOrder() GETOrderCopies200ResponseDataInnerRelationshipsSourceOrder`

GetSourceOrder returns the SourceOrder field if non-nil, zero value otherwise.

### GetSourceOrderOk

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetSourceOrderOk() (*GETOrderCopies200ResponseDataInnerRelationshipsSourceOrder, bool)`

GetSourceOrderOk returns a tuple with the SourceOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOrder

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) SetSourceOrder(v GETOrderCopies200ResponseDataInnerRelationshipsSourceOrder)`

SetSourceOrder sets SourceOrder field to given value.

### HasSourceOrder

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) HasSourceOrder() bool`

HasSourceOrder returns a boolean if a field has been set.

### GetCustomer

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetCustomer() GETCouponRecipients200ResponseDataInnerRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetCustomerOk() (*GETCouponRecipients200ResponseDataInnerRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) SetCustomer(v GETCouponRecipients200ResponseDataInnerRelationshipsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetCustomerPaymentSource

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetCustomerPaymentSource() GETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSource`

GetCustomerPaymentSource returns the CustomerPaymentSource field if non-nil, zero value otherwise.

### GetCustomerPaymentSourceOk

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetCustomerPaymentSourceOk() (*GETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSource, bool)`

GetCustomerPaymentSourceOk returns a tuple with the CustomerPaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPaymentSource

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) SetCustomerPaymentSource(v GETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSource)`

SetCustomerPaymentSource sets CustomerPaymentSource field to given value.

### HasCustomerPaymentSource

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) HasCustomerPaymentSource() bool`

HasCustomerPaymentSource returns a boolean if a field has been set.

### GetOrderSubscriptionItems

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetOrderSubscriptionItems() GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderSubscriptionItems`

GetOrderSubscriptionItems returns the OrderSubscriptionItems field if non-nil, zero value otherwise.

### GetOrderSubscriptionItemsOk

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetOrderSubscriptionItemsOk() (*GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderSubscriptionItems, bool)`

GetOrderSubscriptionItemsOk returns a tuple with the OrderSubscriptionItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSubscriptionItems

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) SetOrderSubscriptionItems(v GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderSubscriptionItems)`

SetOrderSubscriptionItems sets OrderSubscriptionItems field to given value.

### HasOrderSubscriptionItems

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) HasOrderSubscriptionItems() bool`

HasOrderSubscriptionItems returns a boolean if a field has been set.

### GetOrderFactories

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetOrderFactories() GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderFactories`

GetOrderFactories returns the OrderFactories field if non-nil, zero value otherwise.

### GetOrderFactoriesOk

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetOrderFactoriesOk() (*GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderFactories, bool)`

GetOrderFactoriesOk returns a tuple with the OrderFactories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderFactories

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) SetOrderFactories(v GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderFactories)`

SetOrderFactories sets OrderFactories field to given value.

### HasOrderFactories

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) HasOrderFactories() bool`

HasOrderFactories returns a boolean if a field has been set.

### GetOrderCopies

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetOrderCopies() GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies`

GetOrderCopies returns the OrderCopies field if non-nil, zero value otherwise.

### GetOrderCopiesOk

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetOrderCopiesOk() (*GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies, bool)`

GetOrderCopiesOk returns a tuple with the OrderCopies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCopies

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) SetOrderCopies(v GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies)`

SetOrderCopies sets OrderCopies field to given value.

### HasOrderCopies

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) HasOrderCopies() bool`

HasOrderCopies returns a boolean if a field has been set.

### GetRecurringOrderCopies

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetRecurringOrderCopies() GETOrderSubscriptions200ResponseDataInnerRelationshipsRecurringOrderCopies`

GetRecurringOrderCopies returns the RecurringOrderCopies field if non-nil, zero value otherwise.

### GetRecurringOrderCopiesOk

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetRecurringOrderCopiesOk() (*GETOrderSubscriptions200ResponseDataInnerRelationshipsRecurringOrderCopies, bool)`

GetRecurringOrderCopiesOk returns a tuple with the RecurringOrderCopies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringOrderCopies

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) SetRecurringOrderCopies(v GETOrderSubscriptions200ResponseDataInnerRelationshipsRecurringOrderCopies)`

SetRecurringOrderCopies sets RecurringOrderCopies field to given value.

### HasRecurringOrderCopies

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) HasRecurringOrderCopies() bool`

HasRecurringOrderCopies returns a boolean if a field has been set.

### GetOrders

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetOrders() GETCustomers200ResponseDataInnerRelationshipsOrders`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetOrdersOk() (*GETCustomers200ResponseDataInnerRelationshipsOrders, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) SetOrders(v GETCustomers200ResponseDataInnerRelationshipsOrders)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetEvents

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetEvents() GETAuthorizations200ResponseDataInnerRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetEventsOk() (*GETAuthorizations200ResponseDataInnerRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) SetEvents(v GETAuthorizations200ResponseDataInnerRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


