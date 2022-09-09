# GETCustomers200ResponseDataInnerRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerGroup** | Pointer to [**GETAddresses200ResponseDataInnerRelationshipsGeocoder**](GETAddresses200ResponseDataInnerRelationshipsGeocoder.md) |  | [optional] 
**CustomerAddresses** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**CustomerPaymentSources** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**CustomerSubscriptions** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**Orders** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**OrderSubscriptions** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**Returns** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**Attachments** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**Events** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 

## Methods

### NewGETCustomers200ResponseDataInnerRelationships

`func NewGETCustomers200ResponseDataInnerRelationships() *GETCustomers200ResponseDataInnerRelationships`

NewGETCustomers200ResponseDataInnerRelationships instantiates a new GETCustomers200ResponseDataInnerRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETCustomers200ResponseDataInnerRelationshipsWithDefaults

`func NewGETCustomers200ResponseDataInnerRelationshipsWithDefaults() *GETCustomers200ResponseDataInnerRelationships`

NewGETCustomers200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETCustomers200ResponseDataInnerRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerGroup

`func (o *GETCustomers200ResponseDataInnerRelationships) GetCustomerGroup() GETAddresses200ResponseDataInnerRelationshipsGeocoder`

GetCustomerGroup returns the CustomerGroup field if non-nil, zero value otherwise.

### GetCustomerGroupOk

`func (o *GETCustomers200ResponseDataInnerRelationships) GetCustomerGroupOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool)`

GetCustomerGroupOk returns a tuple with the CustomerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerGroup

`func (o *GETCustomers200ResponseDataInnerRelationships) SetCustomerGroup(v GETAddresses200ResponseDataInnerRelationshipsGeocoder)`

SetCustomerGroup sets CustomerGroup field to given value.

### HasCustomerGroup

`func (o *GETCustomers200ResponseDataInnerRelationships) HasCustomerGroup() bool`

HasCustomerGroup returns a boolean if a field has been set.

### GetCustomerAddresses

`func (o *GETCustomers200ResponseDataInnerRelationships) GetCustomerAddresses() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetCustomerAddresses returns the CustomerAddresses field if non-nil, zero value otherwise.

### GetCustomerAddressesOk

`func (o *GETCustomers200ResponseDataInnerRelationships) GetCustomerAddressesOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetCustomerAddressesOk returns a tuple with the CustomerAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAddresses

`func (o *GETCustomers200ResponseDataInnerRelationships) SetCustomerAddresses(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetCustomerAddresses sets CustomerAddresses field to given value.

### HasCustomerAddresses

`func (o *GETCustomers200ResponseDataInnerRelationships) HasCustomerAddresses() bool`

HasCustomerAddresses returns a boolean if a field has been set.

### GetCustomerPaymentSources

`func (o *GETCustomers200ResponseDataInnerRelationships) GetCustomerPaymentSources() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetCustomerPaymentSources returns the CustomerPaymentSources field if non-nil, zero value otherwise.

### GetCustomerPaymentSourcesOk

`func (o *GETCustomers200ResponseDataInnerRelationships) GetCustomerPaymentSourcesOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetCustomerPaymentSourcesOk returns a tuple with the CustomerPaymentSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPaymentSources

`func (o *GETCustomers200ResponseDataInnerRelationships) SetCustomerPaymentSources(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetCustomerPaymentSources sets CustomerPaymentSources field to given value.

### HasCustomerPaymentSources

`func (o *GETCustomers200ResponseDataInnerRelationships) HasCustomerPaymentSources() bool`

HasCustomerPaymentSources returns a boolean if a field has been set.

### GetCustomerSubscriptions

`func (o *GETCustomers200ResponseDataInnerRelationships) GetCustomerSubscriptions() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetCustomerSubscriptions returns the CustomerSubscriptions field if non-nil, zero value otherwise.

### GetCustomerSubscriptionsOk

`func (o *GETCustomers200ResponseDataInnerRelationships) GetCustomerSubscriptionsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetCustomerSubscriptionsOk returns a tuple with the CustomerSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerSubscriptions

`func (o *GETCustomers200ResponseDataInnerRelationships) SetCustomerSubscriptions(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetCustomerSubscriptions sets CustomerSubscriptions field to given value.

### HasCustomerSubscriptions

`func (o *GETCustomers200ResponseDataInnerRelationships) HasCustomerSubscriptions() bool`

HasCustomerSubscriptions returns a boolean if a field has been set.

### GetOrders

`func (o *GETCustomers200ResponseDataInnerRelationships) GetOrders() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *GETCustomers200ResponseDataInnerRelationships) GetOrdersOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *GETCustomers200ResponseDataInnerRelationships) SetOrders(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *GETCustomers200ResponseDataInnerRelationships) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetOrderSubscriptions

`func (o *GETCustomers200ResponseDataInnerRelationships) GetOrderSubscriptions() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetOrderSubscriptions returns the OrderSubscriptions field if non-nil, zero value otherwise.

### GetOrderSubscriptionsOk

`func (o *GETCustomers200ResponseDataInnerRelationships) GetOrderSubscriptionsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetOrderSubscriptionsOk returns a tuple with the OrderSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSubscriptions

`func (o *GETCustomers200ResponseDataInnerRelationships) SetOrderSubscriptions(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetOrderSubscriptions sets OrderSubscriptions field to given value.

### HasOrderSubscriptions

`func (o *GETCustomers200ResponseDataInnerRelationships) HasOrderSubscriptions() bool`

HasOrderSubscriptions returns a boolean if a field has been set.

### GetReturns

`func (o *GETCustomers200ResponseDataInnerRelationships) GetReturns() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetReturns returns the Returns field if non-nil, zero value otherwise.

### GetReturnsOk

`func (o *GETCustomers200ResponseDataInnerRelationships) GetReturnsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetReturnsOk returns a tuple with the Returns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturns

`func (o *GETCustomers200ResponseDataInnerRelationships) SetReturns(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetReturns sets Returns field to given value.

### HasReturns

`func (o *GETCustomers200ResponseDataInnerRelationships) HasReturns() bool`

HasReturns returns a boolean if a field has been set.

### GetAttachments

`func (o *GETCustomers200ResponseDataInnerRelationships) GetAttachments() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GETCustomers200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GETCustomers200ResponseDataInnerRelationships) SetAttachments(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GETCustomers200ResponseDataInnerRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *GETCustomers200ResponseDataInnerRelationships) GetEvents() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GETCustomers200ResponseDataInnerRelationships) GetEventsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GETCustomers200ResponseDataInnerRelationships) SetEvents(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *GETCustomers200ResponseDataInnerRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


