# LineItemCreateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | [**AdyenPaymentCreateDataRelationshipsOrder**](AdyenPaymentCreateDataRelationshipsOrder.md) |  | 
**Item** | Pointer to [**LineItemCreateDataRelationshipsItem**](LineItemCreateDataRelationshipsItem.md) |  | [optional] 
**Tags** | Pointer to [**AddressCreateDataRelationshipsTags**](AddressCreateDataRelationshipsTags.md) |  | [optional] 

## Methods

### NewLineItemCreateDataRelationships

`func NewLineItemCreateDataRelationships(order AdyenPaymentCreateDataRelationshipsOrder, ) *LineItemCreateDataRelationships`

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

`func (o *LineItemCreateDataRelationships) GetOrder() AdyenPaymentCreateDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *LineItemCreateDataRelationships) GetOrderOk() (*AdyenPaymentCreateDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *LineItemCreateDataRelationships) SetOrder(v AdyenPaymentCreateDataRelationshipsOrder)`

SetOrder sets Order field to given value.


### GetItem

`func (o *LineItemCreateDataRelationships) GetItem() LineItemCreateDataRelationshipsItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *LineItemCreateDataRelationships) GetItemOk() (*LineItemCreateDataRelationshipsItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *LineItemCreateDataRelationships) SetItem(v LineItemCreateDataRelationshipsItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *LineItemCreateDataRelationships) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetTags

`func (o *LineItemCreateDataRelationships) GetTags() AddressCreateDataRelationshipsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LineItemCreateDataRelationships) GetTagsOk() (*AddressCreateDataRelationshipsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LineItemCreateDataRelationships) SetTags(v AddressCreateDataRelationshipsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LineItemCreateDataRelationships) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


