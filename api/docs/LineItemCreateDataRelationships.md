# LineItemCreateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | [**AdyenPaymentDataRelationshipsOrder**](AdyenPaymentDataRelationshipsOrder.md) |  | 
**Item** | Pointer to [**LineItemDataRelationshipsItem**](LineItemDataRelationshipsItem.md) |  | [optional] 

## Methods

### NewLineItemCreateDataRelationships

`func NewLineItemCreateDataRelationships(order AdyenPaymentDataRelationshipsOrder, ) *LineItemCreateDataRelationships`

NewLineItemCreateDataRelationships instantiates a new LineItemCreateDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemCreateDataRelationshipsWithDefaults

`func NewLineItemCreateDataRelationshipsWithDefaults() *LineItemCreateDataRelationships`

NewLineItemCreateDataRelationshipsWithDefaults instantiates a new LineItemCreateDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *LineItemCreateDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *LineItemCreateDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *LineItemCreateDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder)`

SetOrder sets Order field to given value.


### GetItem

`func (o *LineItemCreateDataRelationships) GetItem() LineItemDataRelationshipsItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *LineItemCreateDataRelationships) GetItemOk() (*LineItemDataRelationshipsItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *LineItemCreateDataRelationships) SetItem(v LineItemDataRelationshipsItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *LineItemCreateDataRelationships) HasItem() bool`

HasItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


