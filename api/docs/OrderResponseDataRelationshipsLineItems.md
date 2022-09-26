# OrderResponseDataRelationshipsLineItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AddressResponseDataRelationshipsGeocoderLinks**](AddressResponseDataRelationshipsGeocoderLinks.md) |  | [optional] 
**Data** | Pointer to [**[]OrderResponseDataRelationshipsLineItemsDataInner**](OrderResponseDataRelationshipsLineItemsDataInner.md) |  | [optional] 

## Methods

### NewOrderResponseDataRelationshipsLineItems

`func NewOrderResponseDataRelationshipsLineItems() *OrderResponseDataRelationshipsLineItems`

NewOrderResponseDataRelationshipsLineItems instantiates a new OrderResponseDataRelationshipsLineItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderResponseDataRelationshipsLineItemsWithDefaults

`func NewOrderResponseDataRelationshipsLineItemsWithDefaults() *OrderResponseDataRelationshipsLineItems`

NewOrderResponseDataRelationshipsLineItemsWithDefaults instantiates a new OrderResponseDataRelationshipsLineItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *OrderResponseDataRelationshipsLineItems) GetLinks() AddressResponseDataRelationshipsGeocoderLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrderResponseDataRelationshipsLineItems) GetLinksOk() (*AddressResponseDataRelationshipsGeocoderLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrderResponseDataRelationshipsLineItems) SetLinks(v AddressResponseDataRelationshipsGeocoderLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrderResponseDataRelationshipsLineItems) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetData

`func (o *OrderResponseDataRelationshipsLineItems) GetData() []OrderResponseDataRelationshipsLineItemsDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OrderResponseDataRelationshipsLineItems) GetDataOk() (*[]OrderResponseDataRelationshipsLineItemsDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OrderResponseDataRelationshipsLineItems) SetData(v []OrderResponseDataRelationshipsLineItemsDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *OrderResponseDataRelationshipsLineItems) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


