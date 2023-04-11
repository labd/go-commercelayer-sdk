# POSTCustomers201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerGroup** | Pointer to [**POSTCustomers201ResponseDataRelationshipsCustomerGroup**](POSTCustomers201ResponseDataRelationshipsCustomerGroup.md) |  | [optional] 
**CustomerAddresses** | Pointer to [**POSTCustomers201ResponseDataRelationshipsCustomerAddresses**](POSTCustomers201ResponseDataRelationshipsCustomerAddresses.md) |  | [optional] 
**CustomerPaymentSources** | Pointer to [**POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources**](POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources.md) |  | [optional] 
**CustomerSubscriptions** | Pointer to [**POSTCustomers201ResponseDataRelationshipsCustomerSubscriptions**](POSTCustomers201ResponseDataRelationshipsCustomerSubscriptions.md) |  | [optional] 
**Orders** | Pointer to [**POSTCustomers201ResponseDataRelationshipsOrders**](POSTCustomers201ResponseDataRelationshipsOrders.md) |  | [optional] 
**OrderSubscriptions** | Pointer to [**POSTCustomers201ResponseDataRelationshipsOrderSubscriptions**](POSTCustomers201ResponseDataRelationshipsOrderSubscriptions.md) |  | [optional] 
**Returns** | Pointer to [**POSTCustomers201ResponseDataRelationshipsReturns**](POSTCustomers201ResponseDataRelationshipsReturns.md) |  | [optional] 
**SkuLists** | Pointer to [**POSTCustomers201ResponseDataRelationshipsSkuLists**](POSTCustomers201ResponseDataRelationshipsSkuLists.md) |  | [optional] 
**Attachments** | Pointer to [**POSTAvalaraAccounts201ResponseDataRelationshipsAttachments**](POSTAvalaraAccounts201ResponseDataRelationshipsAttachments.md) |  | [optional] 
**Events** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents.md) |  | [optional] 

## Methods

### NewPOSTCustomers201ResponseDataRelationships

`func NewPOSTCustomers201ResponseDataRelationships() *POSTCustomers201ResponseDataRelationships`

NewPOSTCustomers201ResponseDataRelationships instantiates a new POSTCustomers201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTCustomers201ResponseDataRelationshipsWithDefaults

`func NewPOSTCustomers201ResponseDataRelationshipsWithDefaults() *POSTCustomers201ResponseDataRelationships`

NewPOSTCustomers201ResponseDataRelationshipsWithDefaults instantiates a new POSTCustomers201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerGroup

`func (o *POSTCustomers201ResponseDataRelationships) GetCustomerGroup() POSTCustomers201ResponseDataRelationshipsCustomerGroup`

GetCustomerGroup returns the CustomerGroup field if non-nil, zero value otherwise.

### GetCustomerGroupOk

`func (o *POSTCustomers201ResponseDataRelationships) GetCustomerGroupOk() (*POSTCustomers201ResponseDataRelationshipsCustomerGroup, bool)`

GetCustomerGroupOk returns a tuple with the CustomerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerGroup

`func (o *POSTCustomers201ResponseDataRelationships) SetCustomerGroup(v POSTCustomers201ResponseDataRelationshipsCustomerGroup)`

SetCustomerGroup sets CustomerGroup field to given value.

### HasCustomerGroup

`func (o *POSTCustomers201ResponseDataRelationships) HasCustomerGroup() bool`

HasCustomerGroup returns a boolean if a field has been set.

### GetCustomerAddresses

`func (o *POSTCustomers201ResponseDataRelationships) GetCustomerAddresses() POSTCustomers201ResponseDataRelationshipsCustomerAddresses`

GetCustomerAddresses returns the CustomerAddresses field if non-nil, zero value otherwise.

### GetCustomerAddressesOk

`func (o *POSTCustomers201ResponseDataRelationships) GetCustomerAddressesOk() (*POSTCustomers201ResponseDataRelationshipsCustomerAddresses, bool)`

GetCustomerAddressesOk returns a tuple with the CustomerAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAddresses

`func (o *POSTCustomers201ResponseDataRelationships) SetCustomerAddresses(v POSTCustomers201ResponseDataRelationshipsCustomerAddresses)`

SetCustomerAddresses sets CustomerAddresses field to given value.

### HasCustomerAddresses

`func (o *POSTCustomers201ResponseDataRelationships) HasCustomerAddresses() bool`

HasCustomerAddresses returns a boolean if a field has been set.

### GetCustomerPaymentSources

`func (o *POSTCustomers201ResponseDataRelationships) GetCustomerPaymentSources() POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources`

GetCustomerPaymentSources returns the CustomerPaymentSources field if non-nil, zero value otherwise.

### GetCustomerPaymentSourcesOk

`func (o *POSTCustomers201ResponseDataRelationships) GetCustomerPaymentSourcesOk() (*POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources, bool)`

GetCustomerPaymentSourcesOk returns a tuple with the CustomerPaymentSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPaymentSources

`func (o *POSTCustomers201ResponseDataRelationships) SetCustomerPaymentSources(v POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources)`

SetCustomerPaymentSources sets CustomerPaymentSources field to given value.

### HasCustomerPaymentSources

`func (o *POSTCustomers201ResponseDataRelationships) HasCustomerPaymentSources() bool`

HasCustomerPaymentSources returns a boolean if a field has been set.

### GetCustomerSubscriptions

`func (o *POSTCustomers201ResponseDataRelationships) GetCustomerSubscriptions() POSTCustomers201ResponseDataRelationshipsCustomerSubscriptions`

GetCustomerSubscriptions returns the CustomerSubscriptions field if non-nil, zero value otherwise.

### GetCustomerSubscriptionsOk

`func (o *POSTCustomers201ResponseDataRelationships) GetCustomerSubscriptionsOk() (*POSTCustomers201ResponseDataRelationshipsCustomerSubscriptions, bool)`

GetCustomerSubscriptionsOk returns a tuple with the CustomerSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerSubscriptions

`func (o *POSTCustomers201ResponseDataRelationships) SetCustomerSubscriptions(v POSTCustomers201ResponseDataRelationshipsCustomerSubscriptions)`

SetCustomerSubscriptions sets CustomerSubscriptions field to given value.

### HasCustomerSubscriptions

`func (o *POSTCustomers201ResponseDataRelationships) HasCustomerSubscriptions() bool`

HasCustomerSubscriptions returns a boolean if a field has been set.

### GetOrders

`func (o *POSTCustomers201ResponseDataRelationships) GetOrders() POSTCustomers201ResponseDataRelationshipsOrders`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *POSTCustomers201ResponseDataRelationships) GetOrdersOk() (*POSTCustomers201ResponseDataRelationshipsOrders, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *POSTCustomers201ResponseDataRelationships) SetOrders(v POSTCustomers201ResponseDataRelationshipsOrders)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *POSTCustomers201ResponseDataRelationships) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetOrderSubscriptions

`func (o *POSTCustomers201ResponseDataRelationships) GetOrderSubscriptions() POSTCustomers201ResponseDataRelationshipsOrderSubscriptions`

GetOrderSubscriptions returns the OrderSubscriptions field if non-nil, zero value otherwise.

### GetOrderSubscriptionsOk

`func (o *POSTCustomers201ResponseDataRelationships) GetOrderSubscriptionsOk() (*POSTCustomers201ResponseDataRelationshipsOrderSubscriptions, bool)`

GetOrderSubscriptionsOk returns a tuple with the OrderSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSubscriptions

`func (o *POSTCustomers201ResponseDataRelationships) SetOrderSubscriptions(v POSTCustomers201ResponseDataRelationshipsOrderSubscriptions)`

SetOrderSubscriptions sets OrderSubscriptions field to given value.

### HasOrderSubscriptions

`func (o *POSTCustomers201ResponseDataRelationships) HasOrderSubscriptions() bool`

HasOrderSubscriptions returns a boolean if a field has been set.

### GetReturns

`func (o *POSTCustomers201ResponseDataRelationships) GetReturns() POSTCustomers201ResponseDataRelationshipsReturns`

GetReturns returns the Returns field if non-nil, zero value otherwise.

### GetReturnsOk

`func (o *POSTCustomers201ResponseDataRelationships) GetReturnsOk() (*POSTCustomers201ResponseDataRelationshipsReturns, bool)`

GetReturnsOk returns a tuple with the Returns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturns

`func (o *POSTCustomers201ResponseDataRelationships) SetReturns(v POSTCustomers201ResponseDataRelationshipsReturns)`

SetReturns sets Returns field to given value.

### HasReturns

`func (o *POSTCustomers201ResponseDataRelationships) HasReturns() bool`

HasReturns returns a boolean if a field has been set.

### GetSkuLists

`func (o *POSTCustomers201ResponseDataRelationships) GetSkuLists() POSTCustomers201ResponseDataRelationshipsSkuLists`

GetSkuLists returns the SkuLists field if non-nil, zero value otherwise.

### GetSkuListsOk

`func (o *POSTCustomers201ResponseDataRelationships) GetSkuListsOk() (*POSTCustomers201ResponseDataRelationshipsSkuLists, bool)`

GetSkuListsOk returns a tuple with the SkuLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuLists

`func (o *POSTCustomers201ResponseDataRelationships) SetSkuLists(v POSTCustomers201ResponseDataRelationshipsSkuLists)`

SetSkuLists sets SkuLists field to given value.

### HasSkuLists

`func (o *POSTCustomers201ResponseDataRelationships) HasSkuLists() bool`

HasSkuLists returns a boolean if a field has been set.

### GetAttachments

`func (o *POSTCustomers201ResponseDataRelationships) GetAttachments() POSTAvalaraAccounts201ResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *POSTCustomers201ResponseDataRelationships) GetAttachmentsOk() (*POSTAvalaraAccounts201ResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *POSTCustomers201ResponseDataRelationships) SetAttachments(v POSTAvalaraAccounts201ResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *POSTCustomers201ResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *POSTCustomers201ResponseDataRelationships) GetEvents() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *POSTCustomers201ResponseDataRelationships) GetEventsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *POSTCustomers201ResponseDataRelationships) SetEvents(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *POSTCustomers201ResponseDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


