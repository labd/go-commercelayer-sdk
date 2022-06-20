# CustomerDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerGroup** | Pointer to [**CustomerDataRelationshipsCustomerGroup**](CustomerDataRelationshipsCustomerGroup.md) |  | [optional] 
**CustomerAddresses** | Pointer to [**CustomerDataRelationshipsCustomerAddresses**](CustomerDataRelationshipsCustomerAddresses.md) |  | [optional] 
**CustomerPaymentSources** | Pointer to [**CustomerDataRelationshipsCustomerPaymentSources**](CustomerDataRelationshipsCustomerPaymentSources.md) |  | [optional] 
**CustomerSubscriptions** | Pointer to [**CustomerDataRelationshipsCustomerSubscriptions**](CustomerDataRelationshipsCustomerSubscriptions.md) |  | [optional] 
**Orders** | Pointer to [**AdyenPaymentDataRelationshipsOrder**](AdyenPaymentDataRelationshipsOrder.md) |  | [optional] 
**OrderSubscriptions** | Pointer to [**CustomerDataRelationshipsOrderSubscriptions**](CustomerDataRelationshipsOrderSubscriptions.md) |  | [optional] 
**Returns** | Pointer to [**CustomerDataRelationshipsReturns**](CustomerDataRelationshipsReturns.md) |  | [optional] 
**Attachments** | Pointer to [**AvalaraAccountDataRelationshipsAttachments**](AvalaraAccountDataRelationshipsAttachments.md) |  | [optional] 

## Methods

### NewCustomerDataRelationships

`func NewCustomerDataRelationships() *CustomerDataRelationships`

NewCustomerDataRelationships instantiates a new CustomerDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerDataRelationshipsWithDefaults

`func NewCustomerDataRelationshipsWithDefaults() *CustomerDataRelationships`

NewCustomerDataRelationshipsWithDefaults instantiates a new CustomerDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerGroup

`func (o *CustomerDataRelationships) GetCustomerGroup() CustomerDataRelationshipsCustomerGroup`

GetCustomerGroup returns the CustomerGroup field if non-nil, zero value otherwise.

### GetCustomerGroupOk

`func (o *CustomerDataRelationships) GetCustomerGroupOk() (*CustomerDataRelationshipsCustomerGroup, bool)`

GetCustomerGroupOk returns a tuple with the CustomerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerGroup

`func (o *CustomerDataRelationships) SetCustomerGroup(v CustomerDataRelationshipsCustomerGroup)`

SetCustomerGroup sets CustomerGroup field to given value.

### HasCustomerGroup

`func (o *CustomerDataRelationships) HasCustomerGroup() bool`

HasCustomerGroup returns a boolean if a field has been set.

### GetCustomerAddresses

`func (o *CustomerDataRelationships) GetCustomerAddresses() CustomerDataRelationshipsCustomerAddresses`

GetCustomerAddresses returns the CustomerAddresses field if non-nil, zero value otherwise.

### GetCustomerAddressesOk

`func (o *CustomerDataRelationships) GetCustomerAddressesOk() (*CustomerDataRelationshipsCustomerAddresses, bool)`

GetCustomerAddressesOk returns a tuple with the CustomerAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAddresses

`func (o *CustomerDataRelationships) SetCustomerAddresses(v CustomerDataRelationshipsCustomerAddresses)`

SetCustomerAddresses sets CustomerAddresses field to given value.

### HasCustomerAddresses

`func (o *CustomerDataRelationships) HasCustomerAddresses() bool`

HasCustomerAddresses returns a boolean if a field has been set.

### GetCustomerPaymentSources

`func (o *CustomerDataRelationships) GetCustomerPaymentSources() CustomerDataRelationshipsCustomerPaymentSources`

GetCustomerPaymentSources returns the CustomerPaymentSources field if non-nil, zero value otherwise.

### GetCustomerPaymentSourcesOk

`func (o *CustomerDataRelationships) GetCustomerPaymentSourcesOk() (*CustomerDataRelationshipsCustomerPaymentSources, bool)`

GetCustomerPaymentSourcesOk returns a tuple with the CustomerPaymentSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPaymentSources

`func (o *CustomerDataRelationships) SetCustomerPaymentSources(v CustomerDataRelationshipsCustomerPaymentSources)`

SetCustomerPaymentSources sets CustomerPaymentSources field to given value.

### HasCustomerPaymentSources

`func (o *CustomerDataRelationships) HasCustomerPaymentSources() bool`

HasCustomerPaymentSources returns a boolean if a field has been set.

### GetCustomerSubscriptions

`func (o *CustomerDataRelationships) GetCustomerSubscriptions() CustomerDataRelationshipsCustomerSubscriptions`

GetCustomerSubscriptions returns the CustomerSubscriptions field if non-nil, zero value otherwise.

### GetCustomerSubscriptionsOk

`func (o *CustomerDataRelationships) GetCustomerSubscriptionsOk() (*CustomerDataRelationshipsCustomerSubscriptions, bool)`

GetCustomerSubscriptionsOk returns a tuple with the CustomerSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerSubscriptions

`func (o *CustomerDataRelationships) SetCustomerSubscriptions(v CustomerDataRelationshipsCustomerSubscriptions)`

SetCustomerSubscriptions sets CustomerSubscriptions field to given value.

### HasCustomerSubscriptions

`func (o *CustomerDataRelationships) HasCustomerSubscriptions() bool`

HasCustomerSubscriptions returns a boolean if a field has been set.

### GetOrders

`func (o *CustomerDataRelationships) GetOrders() AdyenPaymentDataRelationshipsOrder`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *CustomerDataRelationships) GetOrdersOk() (*AdyenPaymentDataRelationshipsOrder, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *CustomerDataRelationships) SetOrders(v AdyenPaymentDataRelationshipsOrder)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *CustomerDataRelationships) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetOrderSubscriptions

`func (o *CustomerDataRelationships) GetOrderSubscriptions() CustomerDataRelationshipsOrderSubscriptions`

GetOrderSubscriptions returns the OrderSubscriptions field if non-nil, zero value otherwise.

### GetOrderSubscriptionsOk

`func (o *CustomerDataRelationships) GetOrderSubscriptionsOk() (*CustomerDataRelationshipsOrderSubscriptions, bool)`

GetOrderSubscriptionsOk returns a tuple with the OrderSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSubscriptions

`func (o *CustomerDataRelationships) SetOrderSubscriptions(v CustomerDataRelationshipsOrderSubscriptions)`

SetOrderSubscriptions sets OrderSubscriptions field to given value.

### HasOrderSubscriptions

`func (o *CustomerDataRelationships) HasOrderSubscriptions() bool`

HasOrderSubscriptions returns a boolean if a field has been set.

### GetReturns

`func (o *CustomerDataRelationships) GetReturns() CustomerDataRelationshipsReturns`

GetReturns returns the Returns field if non-nil, zero value otherwise.

### GetReturnsOk

`func (o *CustomerDataRelationships) GetReturnsOk() (*CustomerDataRelationshipsReturns, bool)`

GetReturnsOk returns a tuple with the Returns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturns

`func (o *CustomerDataRelationships) SetReturns(v CustomerDataRelationshipsReturns)`

SetReturns sets Returns field to given value.

### HasReturns

`func (o *CustomerDataRelationships) HasReturns() bool`

HasReturns returns a boolean if a field has been set.

### GetAttachments

`func (o *CustomerDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *CustomerDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *CustomerDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *CustomerDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


