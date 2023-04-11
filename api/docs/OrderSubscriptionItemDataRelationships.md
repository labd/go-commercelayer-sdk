# OrderSubscriptionItemDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderSubscription** | Pointer to [**CustomerDataRelationshipsOrderSubscriptions**](CustomerDataRelationshipsOrderSubscriptions.md) |  | [optional] 
**Item** | Pointer to [**OrderSubscriptionItemDataRelationshipsItem**](OrderSubscriptionItemDataRelationshipsItem.md) |  | [optional] 

## Methods

### NewOrderSubscriptionItemDataRelationships

`func NewOrderSubscriptionItemDataRelationships() *OrderSubscriptionItemDataRelationships`

NewOrderSubscriptionItemDataRelationships instantiates a new OrderSubscriptionItemDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSubscriptionItemDataRelationshipsWithDefaults

`func NewOrderSubscriptionItemDataRelationshipsWithDefaults() *OrderSubscriptionItemDataRelationships`

NewOrderSubscriptionItemDataRelationshipsWithDefaults instantiates a new OrderSubscriptionItemDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderSubscription

`func (o *OrderSubscriptionItemDataRelationships) GetOrderSubscription() CustomerDataRelationshipsOrderSubscriptions`

GetOrderSubscription returns the OrderSubscription field if non-nil, zero value otherwise.

### GetOrderSubscriptionOk

`func (o *OrderSubscriptionItemDataRelationships) GetOrderSubscriptionOk() (*CustomerDataRelationshipsOrderSubscriptions, bool)`

GetOrderSubscriptionOk returns a tuple with the OrderSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSubscription

`func (o *OrderSubscriptionItemDataRelationships) SetOrderSubscription(v CustomerDataRelationshipsOrderSubscriptions)`

SetOrderSubscription sets OrderSubscription field to given value.

### HasOrderSubscription

`func (o *OrderSubscriptionItemDataRelationships) HasOrderSubscription() bool`

HasOrderSubscription returns a boolean if a field has been set.

### GetItem

`func (o *OrderSubscriptionItemDataRelationships) GetItem() OrderSubscriptionItemDataRelationshipsItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *OrderSubscriptionItemDataRelationships) GetItemOk() (*OrderSubscriptionItemDataRelationshipsItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *OrderSubscriptionItemDataRelationships) SetItem(v OrderSubscriptionItemDataRelationshipsItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *OrderSubscriptionItemDataRelationships) HasItem() bool`

HasItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


