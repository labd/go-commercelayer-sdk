# POSTOrderSubscriptionItems201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderSubscription** | Pointer to [**POSTOrderCopies201ResponseDataRelationshipsOrderSubscription**](POSTOrderCopies201ResponseDataRelationshipsOrderSubscription.md) |  | [optional] 
**Item** | Pointer to [**POSTLineItems201ResponseDataRelationshipsItem**](POSTLineItems201ResponseDataRelationshipsItem.md) |  | [optional] 
**Sku** | Pointer to [**POSTInStockSubscriptions201ResponseDataRelationshipsSku**](POSTInStockSubscriptions201ResponseDataRelationshipsSku.md) |  | [optional] 
**Bundle** | Pointer to [**POSTLineItems201ResponseDataRelationshipsBundle**](POSTLineItems201ResponseDataRelationshipsBundle.md) |  | [optional] 
**Adjustment** | Pointer to [**POSTLineItems201ResponseDataRelationshipsAdjustment**](POSTLineItems201ResponseDataRelationshipsAdjustment.md) |  | [optional] 
**SourceLineItem** | Pointer to [**POSTOrderSubscriptionItems201ResponseDataRelationshipsSourceLineItem**](POSTOrderSubscriptionItems201ResponseDataRelationshipsSourceLineItem.md) |  | [optional] 

## Methods

### NewPOSTOrderSubscriptionItems201ResponseDataRelationships

`func NewPOSTOrderSubscriptionItems201ResponseDataRelationships() *POSTOrderSubscriptionItems201ResponseDataRelationships`

NewPOSTOrderSubscriptionItems201ResponseDataRelationships instantiates a new POSTOrderSubscriptionItems201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTOrderSubscriptionItems201ResponseDataRelationshipsWithDefaults

`func NewPOSTOrderSubscriptionItems201ResponseDataRelationshipsWithDefaults() *POSTOrderSubscriptionItems201ResponseDataRelationships`

NewPOSTOrderSubscriptionItems201ResponseDataRelationshipsWithDefaults instantiates a new POSTOrderSubscriptionItems201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderSubscription

`func (o *POSTOrderSubscriptionItems201ResponseDataRelationships) GetOrderSubscription() POSTOrderCopies201ResponseDataRelationshipsOrderSubscription`

GetOrderSubscription returns the OrderSubscription field if non-nil, zero value otherwise.

### GetOrderSubscriptionOk

`func (o *POSTOrderSubscriptionItems201ResponseDataRelationships) GetOrderSubscriptionOk() (*POSTOrderCopies201ResponseDataRelationshipsOrderSubscription, bool)`

GetOrderSubscriptionOk returns a tuple with the OrderSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSubscription

`func (o *POSTOrderSubscriptionItems201ResponseDataRelationships) SetOrderSubscription(v POSTOrderCopies201ResponseDataRelationshipsOrderSubscription)`

SetOrderSubscription sets OrderSubscription field to given value.

### HasOrderSubscription

`func (o *POSTOrderSubscriptionItems201ResponseDataRelationships) HasOrderSubscription() bool`

HasOrderSubscription returns a boolean if a field has been set.

### GetItem

`func (o *POSTOrderSubscriptionItems201ResponseDataRelationships) GetItem() POSTLineItems201ResponseDataRelationshipsItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *POSTOrderSubscriptionItems201ResponseDataRelationships) GetItemOk() (*POSTLineItems201ResponseDataRelationshipsItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *POSTOrderSubscriptionItems201ResponseDataRelationships) SetItem(v POSTLineItems201ResponseDataRelationshipsItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *POSTOrderSubscriptionItems201ResponseDataRelationships) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetSku

`func (o *POSTOrderSubscriptionItems201ResponseDataRelationships) GetSku() POSTInStockSubscriptions201ResponseDataRelationshipsSku`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *POSTOrderSubscriptionItems201ResponseDataRelationships) GetSkuOk() (*POSTInStockSubscriptions201ResponseDataRelationshipsSku, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *POSTOrderSubscriptionItems201ResponseDataRelationships) SetSku(v POSTInStockSubscriptions201ResponseDataRelationshipsSku)`

SetSku sets Sku field to given value.

### HasSku

`func (o *POSTOrderSubscriptionItems201ResponseDataRelationships) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetBundle

`func (o *POSTOrderSubscriptionItems201ResponseDataRelationships) GetBundle() POSTLineItems201ResponseDataRelationshipsBundle`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *POSTOrderSubscriptionItems201ResponseDataRelationships) GetBundleOk() (*POSTLineItems201ResponseDataRelationshipsBundle, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *POSTOrderSubscriptionItems201ResponseDataRelationships) SetBundle(v POSTLineItems201ResponseDataRelationshipsBundle)`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *POSTOrderSubscriptionItems201ResponseDataRelationships) HasBundle() bool`

HasBundle returns a boolean if a field has been set.

### GetAdjustment

`func (o *POSTOrderSubscriptionItems201ResponseDataRelationships) GetAdjustment() POSTLineItems201ResponseDataRelationshipsAdjustment`

GetAdjustment returns the Adjustment field if non-nil, zero value otherwise.

### GetAdjustmentOk

`func (o *POSTOrderSubscriptionItems201ResponseDataRelationships) GetAdjustmentOk() (*POSTLineItems201ResponseDataRelationshipsAdjustment, bool)`

GetAdjustmentOk returns a tuple with the Adjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustment

`func (o *POSTOrderSubscriptionItems201ResponseDataRelationships) SetAdjustment(v POSTLineItems201ResponseDataRelationshipsAdjustment)`

SetAdjustment sets Adjustment field to given value.

### HasAdjustment

`func (o *POSTOrderSubscriptionItems201ResponseDataRelationships) HasAdjustment() bool`

HasAdjustment returns a boolean if a field has been set.

### GetSourceLineItem

`func (o *POSTOrderSubscriptionItems201ResponseDataRelationships) GetSourceLineItem() POSTOrderSubscriptionItems201ResponseDataRelationshipsSourceLineItem`

GetSourceLineItem returns the SourceLineItem field if non-nil, zero value otherwise.

### GetSourceLineItemOk

`func (o *POSTOrderSubscriptionItems201ResponseDataRelationships) GetSourceLineItemOk() (*POSTOrderSubscriptionItems201ResponseDataRelationshipsSourceLineItem, bool)`

GetSourceLineItemOk returns a tuple with the SourceLineItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLineItem

`func (o *POSTOrderSubscriptionItems201ResponseDataRelationships) SetSourceLineItem(v POSTOrderSubscriptionItems201ResponseDataRelationshipsSourceLineItem)`

SetSourceLineItem sets SourceLineItem field to given value.

### HasSourceLineItem

`func (o *POSTOrderSubscriptionItems201ResponseDataRelationships) HasSourceLineItem() bool`

HasSourceLineItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


