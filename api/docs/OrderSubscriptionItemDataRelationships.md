# OrderSubscriptionItemDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderSubscription** | Pointer to [**CustomerDataRelationshipsOrderSubscriptions**](CustomerDataRelationshipsOrderSubscriptions.md) |  | [optional] 
**Item** | Pointer to [**OrderSubscriptionItemDataRelationshipsItem**](OrderSubscriptionItemDataRelationshipsItem.md) |  | [optional] 
**Sku** | Pointer to [**BundleDataRelationshipsSkus**](BundleDataRelationshipsSkus.md) |  | [optional] 
**Bundle** | Pointer to [**LineItemDataRelationshipsBundle**](LineItemDataRelationshipsBundle.md) |  | [optional] 
**Adjustment** | Pointer to [**LineItemDataRelationshipsAdjustment**](LineItemDataRelationshipsAdjustment.md) |  | [optional] 
**SourceLineItem** | Pointer to [**LineItemOptionDataRelationshipsLineItem**](LineItemOptionDataRelationshipsLineItem.md) |  | [optional] 

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

### GetSku

`func (o *OrderSubscriptionItemDataRelationships) GetSku() BundleDataRelationshipsSkus`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *OrderSubscriptionItemDataRelationships) GetSkuOk() (*BundleDataRelationshipsSkus, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *OrderSubscriptionItemDataRelationships) SetSku(v BundleDataRelationshipsSkus)`

SetSku sets Sku field to given value.

### HasSku

`func (o *OrderSubscriptionItemDataRelationships) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetBundle

`func (o *OrderSubscriptionItemDataRelationships) GetBundle() LineItemDataRelationshipsBundle`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *OrderSubscriptionItemDataRelationships) GetBundleOk() (*LineItemDataRelationshipsBundle, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *OrderSubscriptionItemDataRelationships) SetBundle(v LineItemDataRelationshipsBundle)`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *OrderSubscriptionItemDataRelationships) HasBundle() bool`

HasBundle returns a boolean if a field has been set.

### GetAdjustment

`func (o *OrderSubscriptionItemDataRelationships) GetAdjustment() LineItemDataRelationshipsAdjustment`

GetAdjustment returns the Adjustment field if non-nil, zero value otherwise.

### GetAdjustmentOk

`func (o *OrderSubscriptionItemDataRelationships) GetAdjustmentOk() (*LineItemDataRelationshipsAdjustment, bool)`

GetAdjustmentOk returns a tuple with the Adjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustment

`func (o *OrderSubscriptionItemDataRelationships) SetAdjustment(v LineItemDataRelationshipsAdjustment)`

SetAdjustment sets Adjustment field to given value.

### HasAdjustment

`func (o *OrderSubscriptionItemDataRelationships) HasAdjustment() bool`

HasAdjustment returns a boolean if a field has been set.

### GetSourceLineItem

`func (o *OrderSubscriptionItemDataRelationships) GetSourceLineItem() LineItemOptionDataRelationshipsLineItem`

GetSourceLineItem returns the SourceLineItem field if non-nil, zero value otherwise.

### GetSourceLineItemOk

`func (o *OrderSubscriptionItemDataRelationships) GetSourceLineItemOk() (*LineItemOptionDataRelationshipsLineItem, bool)`

GetSourceLineItemOk returns a tuple with the SourceLineItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLineItem

`func (o *OrderSubscriptionItemDataRelationships) SetSourceLineItem(v LineItemOptionDataRelationshipsLineItem)`

SetSourceLineItem sets SourceLineItem field to given value.

### HasSourceLineItem

`func (o *OrderSubscriptionItemDataRelationships) HasSourceLineItem() bool`

HasSourceLineItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


