# POSTLineItemsRequestDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | [**POSTAdyenPaymentsRequestDataRelationshipsOrder**](POSTAdyenPaymentsRequestDataRelationshipsOrder.md) |  | 
**Item** | Pointer to [**POSTLineItemsRequestDataRelationshipsItem**](POSTLineItemsRequestDataRelationshipsItem.md) |  | [optional] 

## Methods

### NewPOSTLineItemsRequestDataRelationships

`func NewPOSTLineItemsRequestDataRelationships(order POSTAdyenPaymentsRequestDataRelationshipsOrder, ) *POSTLineItemsRequestDataRelationships`

NewPOSTLineItemsRequestDataRelationships instantiates a new POSTLineItemsRequestDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTLineItemsRequestDataRelationshipsWithDefaults

`func NewPOSTLineItemsRequestDataRelationshipsWithDefaults() *POSTLineItemsRequestDataRelationships`

NewPOSTLineItemsRequestDataRelationshipsWithDefaults instantiates a new POSTLineItemsRequestDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *POSTLineItemsRequestDataRelationships) GetOrder() POSTAdyenPaymentsRequestDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *POSTLineItemsRequestDataRelationships) GetOrderOk() (*POSTAdyenPaymentsRequestDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *POSTLineItemsRequestDataRelationships) SetOrder(v POSTAdyenPaymentsRequestDataRelationshipsOrder)`

SetOrder sets Order field to given value.


### GetItem

`func (o *POSTLineItemsRequestDataRelationships) GetItem() POSTLineItemsRequestDataRelationshipsItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *POSTLineItemsRequestDataRelationships) GetItemOk() (*POSTLineItemsRequestDataRelationshipsItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *POSTLineItemsRequestDataRelationships) SetItem(v POSTLineItemsRequestDataRelationshipsItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *POSTLineItemsRequestDataRelationships) HasItem() bool`

HasItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


