# OrderCopyResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceOrder** | Pointer to [**OrderCopyResponseDataRelationshipsSourceOrder**](OrderCopyResponseDataRelationshipsSourceOrder.md) |  | [optional] 
**TargetOrder** | Pointer to [**OrderCopyResponseDataRelationshipsTargetOrder**](OrderCopyResponseDataRelationshipsTargetOrder.md) |  | [optional] 
**OrderSubscription** | Pointer to [**OrderCopyResponseDataRelationshipsOrderSubscription**](OrderCopyResponseDataRelationshipsOrderSubscription.md) |  | [optional] 
**Events** | Pointer to [**CustomerAddressResponseDataRelationshipsEvents**](CustomerAddressResponseDataRelationshipsEvents.md) |  | [optional] 

## Methods

### NewOrderCopyResponseDataRelationships

`func NewOrderCopyResponseDataRelationships() *OrderCopyResponseDataRelationships`

NewOrderCopyResponseDataRelationships instantiates a new OrderCopyResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCopyResponseDataRelationshipsWithDefaults

`func NewOrderCopyResponseDataRelationshipsWithDefaults() *OrderCopyResponseDataRelationships`

NewOrderCopyResponseDataRelationshipsWithDefaults instantiates a new OrderCopyResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceOrder

`func (o *OrderCopyResponseDataRelationships) GetSourceOrder() OrderCopyResponseDataRelationshipsSourceOrder`

GetSourceOrder returns the SourceOrder field if non-nil, zero value otherwise.

### GetSourceOrderOk

`func (o *OrderCopyResponseDataRelationships) GetSourceOrderOk() (*OrderCopyResponseDataRelationshipsSourceOrder, bool)`

GetSourceOrderOk returns a tuple with the SourceOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOrder

`func (o *OrderCopyResponseDataRelationships) SetSourceOrder(v OrderCopyResponseDataRelationshipsSourceOrder)`

SetSourceOrder sets SourceOrder field to given value.

### HasSourceOrder

`func (o *OrderCopyResponseDataRelationships) HasSourceOrder() bool`

HasSourceOrder returns a boolean if a field has been set.

### GetTargetOrder

`func (o *OrderCopyResponseDataRelationships) GetTargetOrder() OrderCopyResponseDataRelationshipsTargetOrder`

GetTargetOrder returns the TargetOrder field if non-nil, zero value otherwise.

### GetTargetOrderOk

`func (o *OrderCopyResponseDataRelationships) GetTargetOrderOk() (*OrderCopyResponseDataRelationshipsTargetOrder, bool)`

GetTargetOrderOk returns a tuple with the TargetOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetOrder

`func (o *OrderCopyResponseDataRelationships) SetTargetOrder(v OrderCopyResponseDataRelationshipsTargetOrder)`

SetTargetOrder sets TargetOrder field to given value.

### HasTargetOrder

`func (o *OrderCopyResponseDataRelationships) HasTargetOrder() bool`

HasTargetOrder returns a boolean if a field has been set.

### GetOrderSubscription

`func (o *OrderCopyResponseDataRelationships) GetOrderSubscription() OrderCopyResponseDataRelationshipsOrderSubscription`

GetOrderSubscription returns the OrderSubscription field if non-nil, zero value otherwise.

### GetOrderSubscriptionOk

`func (o *OrderCopyResponseDataRelationships) GetOrderSubscriptionOk() (*OrderCopyResponseDataRelationshipsOrderSubscription, bool)`

GetOrderSubscriptionOk returns a tuple with the OrderSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSubscription

`func (o *OrderCopyResponseDataRelationships) SetOrderSubscription(v OrderCopyResponseDataRelationshipsOrderSubscription)`

SetOrderSubscription sets OrderSubscription field to given value.

### HasOrderSubscription

`func (o *OrderCopyResponseDataRelationships) HasOrderSubscription() bool`

HasOrderSubscription returns a boolean if a field has been set.

### GetEvents

`func (o *OrderCopyResponseDataRelationships) GetEvents() CustomerAddressResponseDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *OrderCopyResponseDataRelationships) GetEventsOk() (*CustomerAddressResponseDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *OrderCopyResponseDataRelationships) SetEvents(v CustomerAddressResponseDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *OrderCopyResponseDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


