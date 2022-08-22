# POSTLineItems201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | [**GETAdyenPayments200ResponseDataInnerRelationshipsOrder**](GETAdyenPayments200ResponseDataInnerRelationshipsOrder.md) |  | 
**Item** | Pointer to [**GETLineItems200ResponseDataInnerRelationshipsItem**](GETLineItems200ResponseDataInnerRelationshipsItem.md) |  | [optional] 

## Methods

### NewPOSTLineItems201ResponseDataRelationships

`func NewPOSTLineItems201ResponseDataRelationships(order GETAdyenPayments200ResponseDataInnerRelationshipsOrder, ) *POSTLineItems201ResponseDataRelationships`

NewPOSTLineItems201ResponseDataRelationships instantiates a new POSTLineItems201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTLineItems201ResponseDataRelationshipsWithDefaults

`func NewPOSTLineItems201ResponseDataRelationshipsWithDefaults() *POSTLineItems201ResponseDataRelationships`

NewPOSTLineItems201ResponseDataRelationshipsWithDefaults instantiates a new POSTLineItems201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *POSTLineItems201ResponseDataRelationships) GetOrder() GETAdyenPayments200ResponseDataInnerRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *POSTLineItems201ResponseDataRelationships) GetOrderOk() (*GETAdyenPayments200ResponseDataInnerRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *POSTLineItems201ResponseDataRelationships) SetOrder(v GETAdyenPayments200ResponseDataInnerRelationshipsOrder)`

SetOrder sets Order field to given value.


### GetItem

`func (o *POSTLineItems201ResponseDataRelationships) GetItem() GETLineItems200ResponseDataInnerRelationshipsItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *POSTLineItems201ResponseDataRelationships) GetItemOk() (*GETLineItems200ResponseDataInnerRelationshipsItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *POSTLineItems201ResponseDataRelationships) SetItem(v GETLineItems200ResponseDataInnerRelationshipsItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *POSTLineItems201ResponseDataRelationships) HasItem() bool`

HasItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


