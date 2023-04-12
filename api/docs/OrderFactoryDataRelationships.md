# OrderFactoryDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceOrder** | Pointer to [**AdyenPaymentDataRelationshipsOrder**](AdyenPaymentDataRelationshipsOrder.md) |  | [optional] 
**TargetOrder** | Pointer to [**AdyenPaymentDataRelationshipsOrder**](AdyenPaymentDataRelationshipsOrder.md) |  | [optional] 
**Events** | Pointer to [**AuthorizationDataRelationshipsEvents**](AuthorizationDataRelationshipsEvents.md) |  | [optional] 

## Methods

### NewOrderFactoryDataRelationships

`func NewOrderFactoryDataRelationships() *OrderFactoryDataRelationships`

NewOrderFactoryDataRelationships instantiates a new OrderFactoryDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderFactoryDataRelationshipsWithDefaults

`func NewOrderFactoryDataRelationshipsWithDefaults() *OrderFactoryDataRelationships`

NewOrderFactoryDataRelationshipsWithDefaults instantiates a new OrderFactoryDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceOrder

`func (o *OrderFactoryDataRelationships) GetSourceOrder() AdyenPaymentDataRelationshipsOrder`

GetSourceOrder returns the SourceOrder field if non-nil, zero value otherwise.

### GetSourceOrderOk

`func (o *OrderFactoryDataRelationships) GetSourceOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool)`

GetSourceOrderOk returns a tuple with the SourceOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOrder

`func (o *OrderFactoryDataRelationships) SetSourceOrder(v AdyenPaymentDataRelationshipsOrder)`

SetSourceOrder sets SourceOrder field to given value.

### HasSourceOrder

`func (o *OrderFactoryDataRelationships) HasSourceOrder() bool`

HasSourceOrder returns a boolean if a field has been set.

### GetTargetOrder

`func (o *OrderFactoryDataRelationships) GetTargetOrder() AdyenPaymentDataRelationshipsOrder`

GetTargetOrder returns the TargetOrder field if non-nil, zero value otherwise.

### GetTargetOrderOk

`func (o *OrderFactoryDataRelationships) GetTargetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool)`

GetTargetOrderOk returns a tuple with the TargetOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetOrder

`func (o *OrderFactoryDataRelationships) SetTargetOrder(v AdyenPaymentDataRelationshipsOrder)`

SetTargetOrder sets TargetOrder field to given value.

### HasTargetOrder

`func (o *OrderFactoryDataRelationships) HasTargetOrder() bool`

HasTargetOrder returns a boolean if a field has been set.

### GetEvents

`func (o *OrderFactoryDataRelationships) GetEvents() AuthorizationDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *OrderFactoryDataRelationships) GetEventsOk() (*AuthorizationDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *OrderFactoryDataRelationships) SetEvents(v AuthorizationDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *OrderFactoryDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


