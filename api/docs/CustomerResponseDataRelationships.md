# CustomerResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerGroup** | Pointer to [**CustomerResponseDataRelationshipsCustomerGroup**](CustomerResponseDataRelationshipsCustomerGroup.md) |  | [optional] 
**CustomerAddresses** | Pointer to [**CustomerResponseDataRelationshipsCustomerAddresses**](CustomerResponseDataRelationshipsCustomerAddresses.md) |  | [optional] 
**CustomerPaymentSources** | Pointer to [**CustomerResponseDataRelationshipsCustomerPaymentSources**](CustomerResponseDataRelationshipsCustomerPaymentSources.md) |  | [optional] 
**CustomerSubscriptions** | Pointer to [**CustomerResponseDataRelationshipsCustomerSubscriptions**](CustomerResponseDataRelationshipsCustomerSubscriptions.md) |  | [optional] 
**Orders** | Pointer to [**CustomerResponseDataRelationshipsOrders**](CustomerResponseDataRelationshipsOrders.md) |  | [optional] 
**OrderSubscriptions** | Pointer to [**CustomerResponseDataRelationshipsOrderSubscriptions**](CustomerResponseDataRelationshipsOrderSubscriptions.md) |  | [optional] 
**Returns** | Pointer to [**CustomerResponseDataRelationshipsReturns**](CustomerResponseDataRelationshipsReturns.md) |  | [optional] 
**Attachments** | Pointer to [**AvalaraAccountResponseDataRelationshipsAttachments**](AvalaraAccountResponseDataRelationshipsAttachments.md) |  | [optional] 
**Events** | Pointer to [**CustomerAddressResponseDataRelationshipsEvents**](CustomerAddressResponseDataRelationshipsEvents.md) |  | [optional] 

## Methods

### NewCustomerResponseDataRelationships

`func NewCustomerResponseDataRelationships() *CustomerResponseDataRelationships`

NewCustomerResponseDataRelationships instantiates a new CustomerResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerResponseDataRelationshipsWithDefaults

`func NewCustomerResponseDataRelationshipsWithDefaults() *CustomerResponseDataRelationships`

NewCustomerResponseDataRelationshipsWithDefaults instantiates a new CustomerResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerGroup

`func (o *CustomerResponseDataRelationships) GetCustomerGroup() CustomerResponseDataRelationshipsCustomerGroup`

GetCustomerGroup returns the CustomerGroup field if non-nil, zero value otherwise.

### GetCustomerGroupOk

`func (o *CustomerResponseDataRelationships) GetCustomerGroupOk() (*CustomerResponseDataRelationshipsCustomerGroup, bool)`

GetCustomerGroupOk returns a tuple with the CustomerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerGroup

`func (o *CustomerResponseDataRelationships) SetCustomerGroup(v CustomerResponseDataRelationshipsCustomerGroup)`

SetCustomerGroup sets CustomerGroup field to given value.

### HasCustomerGroup

`func (o *CustomerResponseDataRelationships) HasCustomerGroup() bool`

HasCustomerGroup returns a boolean if a field has been set.

### GetCustomerAddresses

`func (o *CustomerResponseDataRelationships) GetCustomerAddresses() CustomerResponseDataRelationshipsCustomerAddresses`

GetCustomerAddresses returns the CustomerAddresses field if non-nil, zero value otherwise.

### GetCustomerAddressesOk

`func (o *CustomerResponseDataRelationships) GetCustomerAddressesOk() (*CustomerResponseDataRelationshipsCustomerAddresses, bool)`

GetCustomerAddressesOk returns a tuple with the CustomerAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAddresses

`func (o *CustomerResponseDataRelationships) SetCustomerAddresses(v CustomerResponseDataRelationshipsCustomerAddresses)`

SetCustomerAddresses sets CustomerAddresses field to given value.

### HasCustomerAddresses

`func (o *CustomerResponseDataRelationships) HasCustomerAddresses() bool`

HasCustomerAddresses returns a boolean if a field has been set.

### GetCustomerPaymentSources

`func (o *CustomerResponseDataRelationships) GetCustomerPaymentSources() CustomerResponseDataRelationshipsCustomerPaymentSources`

GetCustomerPaymentSources returns the CustomerPaymentSources field if non-nil, zero value otherwise.

### GetCustomerPaymentSourcesOk

`func (o *CustomerResponseDataRelationships) GetCustomerPaymentSourcesOk() (*CustomerResponseDataRelationshipsCustomerPaymentSources, bool)`

GetCustomerPaymentSourcesOk returns a tuple with the CustomerPaymentSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPaymentSources

`func (o *CustomerResponseDataRelationships) SetCustomerPaymentSources(v CustomerResponseDataRelationshipsCustomerPaymentSources)`

SetCustomerPaymentSources sets CustomerPaymentSources field to given value.

### HasCustomerPaymentSources

`func (o *CustomerResponseDataRelationships) HasCustomerPaymentSources() bool`

HasCustomerPaymentSources returns a boolean if a field has been set.

### GetCustomerSubscriptions

`func (o *CustomerResponseDataRelationships) GetCustomerSubscriptions() CustomerResponseDataRelationshipsCustomerSubscriptions`

GetCustomerSubscriptions returns the CustomerSubscriptions field if non-nil, zero value otherwise.

### GetCustomerSubscriptionsOk

`func (o *CustomerResponseDataRelationships) GetCustomerSubscriptionsOk() (*CustomerResponseDataRelationshipsCustomerSubscriptions, bool)`

GetCustomerSubscriptionsOk returns a tuple with the CustomerSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerSubscriptions

`func (o *CustomerResponseDataRelationships) SetCustomerSubscriptions(v CustomerResponseDataRelationshipsCustomerSubscriptions)`

SetCustomerSubscriptions sets CustomerSubscriptions field to given value.

### HasCustomerSubscriptions

`func (o *CustomerResponseDataRelationships) HasCustomerSubscriptions() bool`

HasCustomerSubscriptions returns a boolean if a field has been set.

### GetOrders

`func (o *CustomerResponseDataRelationships) GetOrders() CustomerResponseDataRelationshipsOrders`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *CustomerResponseDataRelationships) GetOrdersOk() (*CustomerResponseDataRelationshipsOrders, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *CustomerResponseDataRelationships) SetOrders(v CustomerResponseDataRelationshipsOrders)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *CustomerResponseDataRelationships) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetOrderSubscriptions

`func (o *CustomerResponseDataRelationships) GetOrderSubscriptions() CustomerResponseDataRelationshipsOrderSubscriptions`

GetOrderSubscriptions returns the OrderSubscriptions field if non-nil, zero value otherwise.

### GetOrderSubscriptionsOk

`func (o *CustomerResponseDataRelationships) GetOrderSubscriptionsOk() (*CustomerResponseDataRelationshipsOrderSubscriptions, bool)`

GetOrderSubscriptionsOk returns a tuple with the OrderSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSubscriptions

`func (o *CustomerResponseDataRelationships) SetOrderSubscriptions(v CustomerResponseDataRelationshipsOrderSubscriptions)`

SetOrderSubscriptions sets OrderSubscriptions field to given value.

### HasOrderSubscriptions

`func (o *CustomerResponseDataRelationships) HasOrderSubscriptions() bool`

HasOrderSubscriptions returns a boolean if a field has been set.

### GetReturns

`func (o *CustomerResponseDataRelationships) GetReturns() CustomerResponseDataRelationshipsReturns`

GetReturns returns the Returns field if non-nil, zero value otherwise.

### GetReturnsOk

`func (o *CustomerResponseDataRelationships) GetReturnsOk() (*CustomerResponseDataRelationshipsReturns, bool)`

GetReturnsOk returns a tuple with the Returns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturns

`func (o *CustomerResponseDataRelationships) SetReturns(v CustomerResponseDataRelationshipsReturns)`

SetReturns sets Returns field to given value.

### HasReturns

`func (o *CustomerResponseDataRelationships) HasReturns() bool`

HasReturns returns a boolean if a field has been set.

### GetAttachments

`func (o *CustomerResponseDataRelationships) GetAttachments() AvalaraAccountResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *CustomerResponseDataRelationships) GetAttachmentsOk() (*AvalaraAccountResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *CustomerResponseDataRelationships) SetAttachments(v AvalaraAccountResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *CustomerResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *CustomerResponseDataRelationships) GetEvents() CustomerAddressResponseDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CustomerResponseDataRelationships) GetEventsOk() (*CustomerAddressResponseDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CustomerResponseDataRelationships) SetEvents(v CustomerAddressResponseDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *CustomerResponseDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


