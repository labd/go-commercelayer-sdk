# OrderSubscriptionItemCreateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderSubscription** | [**OrderSubscriptionItemCreateDataRelationshipsOrderSubscription**](OrderSubscriptionItemCreateDataRelationshipsOrderSubscription.md) |  | 
**Item** | [**OrderSubscriptionItemCreateDataRelationshipsItem**](OrderSubscriptionItemCreateDataRelationshipsItem.md) |  | 
**Sku** | Pointer to [**InStockSubscriptionCreateDataRelationshipsSku**](InStockSubscriptionCreateDataRelationshipsSku.md) |  | [optional] 
**Bundle** | Pointer to [**OrderSubscriptionItemCreateDataRelationshipsBundle**](OrderSubscriptionItemCreateDataRelationshipsBundle.md) |  | [optional] 
**Adjustment** | Pointer to [**OrderSubscriptionItemCreateDataRelationshipsAdjustment**](OrderSubscriptionItemCreateDataRelationshipsAdjustment.md) |  | [optional] 

## Methods

### NewOrderSubscriptionItemCreateDataRelationships

`func NewOrderSubscriptionItemCreateDataRelationships(orderSubscription OrderSubscriptionItemCreateDataRelationshipsOrderSubscription, item OrderSubscriptionItemCreateDataRelationshipsItem, ) *OrderSubscriptionItemCreateDataRelationships`

NewOrderSubscriptionItemCreateDataRelationships instantiates a new OrderSubscriptionItemCreateDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSubscriptionItemCreateDataRelationshipsWithDefaults

`func NewOrderSubscriptionItemCreateDataRelationshipsWithDefaults() *OrderSubscriptionItemCreateDataRelationships`

NewOrderSubscriptionItemCreateDataRelationshipsWithDefaults instantiates a new OrderSubscriptionItemCreateDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderSubscription

`func (o *OrderSubscriptionItemCreateDataRelationships) GetOrderSubscription() OrderSubscriptionItemCreateDataRelationshipsOrderSubscription`

GetOrderSubscription returns the OrderSubscription field if non-nil, zero value otherwise.

### GetOrderSubscriptionOk

`func (o *OrderSubscriptionItemCreateDataRelationships) GetOrderSubscriptionOk() (*OrderSubscriptionItemCreateDataRelationshipsOrderSubscription, bool)`

GetOrderSubscriptionOk returns a tuple with the OrderSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSubscription

`func (o *OrderSubscriptionItemCreateDataRelationships) SetOrderSubscription(v OrderSubscriptionItemCreateDataRelationshipsOrderSubscription)`

SetOrderSubscription sets OrderSubscription field to given value.


### GetItem

`func (o *OrderSubscriptionItemCreateDataRelationships) GetItem() OrderSubscriptionItemCreateDataRelationshipsItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *OrderSubscriptionItemCreateDataRelationships) GetItemOk() (*OrderSubscriptionItemCreateDataRelationshipsItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *OrderSubscriptionItemCreateDataRelationships) SetItem(v OrderSubscriptionItemCreateDataRelationshipsItem)`

SetItem sets Item field to given value.


### GetSku

`func (o *OrderSubscriptionItemCreateDataRelationships) GetSku() InStockSubscriptionCreateDataRelationshipsSku`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *OrderSubscriptionItemCreateDataRelationships) GetSkuOk() (*InStockSubscriptionCreateDataRelationshipsSku, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *OrderSubscriptionItemCreateDataRelationships) SetSku(v InStockSubscriptionCreateDataRelationshipsSku)`

SetSku sets Sku field to given value.

### HasSku

`func (o *OrderSubscriptionItemCreateDataRelationships) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetBundle

`func (o *OrderSubscriptionItemCreateDataRelationships) GetBundle() OrderSubscriptionItemCreateDataRelationshipsBundle`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *OrderSubscriptionItemCreateDataRelationships) GetBundleOk() (*OrderSubscriptionItemCreateDataRelationshipsBundle, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *OrderSubscriptionItemCreateDataRelationships) SetBundle(v OrderSubscriptionItemCreateDataRelationshipsBundle)`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *OrderSubscriptionItemCreateDataRelationships) HasBundle() bool`

HasBundle returns a boolean if a field has been set.

### GetAdjustment

`func (o *OrderSubscriptionItemCreateDataRelationships) GetAdjustment() OrderSubscriptionItemCreateDataRelationshipsAdjustment`

GetAdjustment returns the Adjustment field if non-nil, zero value otherwise.

### GetAdjustmentOk

`func (o *OrderSubscriptionItemCreateDataRelationships) GetAdjustmentOk() (*OrderSubscriptionItemCreateDataRelationshipsAdjustment, bool)`

GetAdjustmentOk returns a tuple with the Adjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustment

`func (o *OrderSubscriptionItemCreateDataRelationships) SetAdjustment(v OrderSubscriptionItemCreateDataRelationshipsAdjustment)`

SetAdjustment sets Adjustment field to given value.

### HasAdjustment

`func (o *OrderSubscriptionItemCreateDataRelationships) HasAdjustment() bool`

HasAdjustment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


