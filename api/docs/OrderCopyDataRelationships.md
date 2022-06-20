# OrderCopyDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceOrder** | Pointer to [**AdyenPaymentDataRelationshipsOrder**](AdyenPaymentDataRelationshipsOrder.md) |  | [optional] 
**TargetOrder** | Pointer to [**AdyenPaymentDataRelationshipsOrder**](AdyenPaymentDataRelationshipsOrder.md) |  | [optional] 
**OrderSubscription** | Pointer to [**CustomerDataRelationshipsOrderSubscriptions**](CustomerDataRelationshipsOrderSubscriptions.md) |  | [optional] 

## Methods

### NewOrderCopyDataRelationships

`func NewOrderCopyDataRelationships() *OrderCopyDataRelationships`

NewOrderCopyDataRelationships instantiates a new OrderCopyDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCopyDataRelationshipsWithDefaults

`func NewOrderCopyDataRelationshipsWithDefaults() *OrderCopyDataRelationships`

NewOrderCopyDataRelationshipsWithDefaults instantiates a new OrderCopyDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceOrder

`func (o *OrderCopyDataRelationships) GetSourceOrder() AdyenPaymentDataRelationshipsOrder`

GetSourceOrder returns the SourceOrder field if non-nil, zero value otherwise.

### GetSourceOrderOk

`func (o *OrderCopyDataRelationships) GetSourceOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool)`

GetSourceOrderOk returns a tuple with the SourceOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOrder

`func (o *OrderCopyDataRelationships) SetSourceOrder(v AdyenPaymentDataRelationshipsOrder)`

SetSourceOrder sets SourceOrder field to given value.

### HasSourceOrder

`func (o *OrderCopyDataRelationships) HasSourceOrder() bool`

HasSourceOrder returns a boolean if a field has been set.

### GetTargetOrder

`func (o *OrderCopyDataRelationships) GetTargetOrder() AdyenPaymentDataRelationshipsOrder`

GetTargetOrder returns the TargetOrder field if non-nil, zero value otherwise.

### GetTargetOrderOk

`func (o *OrderCopyDataRelationships) GetTargetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool)`

GetTargetOrderOk returns a tuple with the TargetOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetOrder

`func (o *OrderCopyDataRelationships) SetTargetOrder(v AdyenPaymentDataRelationshipsOrder)`

SetTargetOrder sets TargetOrder field to given value.

### HasTargetOrder

`func (o *OrderCopyDataRelationships) HasTargetOrder() bool`

HasTargetOrder returns a boolean if a field has been set.

### GetOrderSubscription

`func (o *OrderCopyDataRelationships) GetOrderSubscription() CustomerDataRelationshipsOrderSubscriptions`

GetOrderSubscription returns the OrderSubscription field if non-nil, zero value otherwise.

### GetOrderSubscriptionOk

`func (o *OrderCopyDataRelationships) GetOrderSubscriptionOk() (*CustomerDataRelationshipsOrderSubscriptions, bool)`

GetOrderSubscriptionOk returns a tuple with the OrderSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSubscription

`func (o *OrderCopyDataRelationships) SetOrderSubscription(v CustomerDataRelationshipsOrderSubscriptions)`

SetOrderSubscription sets OrderSubscription field to given value.

### HasOrderSubscription

`func (o *OrderCopyDataRelationships) HasOrderSubscription() bool`

HasOrderSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


