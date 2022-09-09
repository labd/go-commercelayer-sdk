# GETOrders200ResponseDataInnerRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**GETAddresses200ResponseDataInnerRelationshipsGeocoder**](GETAddresses200ResponseDataInnerRelationshipsGeocoder.md) |  | [optional] 
**Customer** | Pointer to [**GETAddresses200ResponseDataInnerRelationshipsGeocoder**](GETAddresses200ResponseDataInnerRelationshipsGeocoder.md) |  | [optional] 
**ShippingAddress** | Pointer to [**GETAddresses200ResponseDataInnerRelationshipsGeocoder**](GETAddresses200ResponseDataInnerRelationshipsGeocoder.md) |  | [optional] 
**BillingAddress** | Pointer to [**GETAddresses200ResponseDataInnerRelationshipsGeocoder**](GETAddresses200ResponseDataInnerRelationshipsGeocoder.md) |  | [optional] 
**AvailablePaymentMethods** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**AvailableCustomerPaymentSources** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**AvailableFreeSkus** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**AvailableFreeBundles** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**PaymentMethod** | Pointer to [**GETAddresses200ResponseDataInnerRelationshipsGeocoder**](GETAddresses200ResponseDataInnerRelationshipsGeocoder.md) |  | [optional] 
**PaymentSource** | Pointer to [**GETAddresses200ResponseDataInnerRelationshipsGeocoder**](GETAddresses200ResponseDataInnerRelationshipsGeocoder.md) |  | [optional] 
**LineItems** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**Shipments** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**Transactions** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**Authorizations** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**Captures** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**Voids** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**Refunds** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**OrderSubscriptions** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**OrderCopies** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**Attachments** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**Events** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 

## Methods

### NewGETOrders200ResponseDataInnerRelationships

`func NewGETOrders200ResponseDataInnerRelationships() *GETOrders200ResponseDataInnerRelationships`

NewGETOrders200ResponseDataInnerRelationships instantiates a new GETOrders200ResponseDataInnerRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETOrders200ResponseDataInnerRelationshipsWithDefaults

`func NewGETOrders200ResponseDataInnerRelationshipsWithDefaults() *GETOrders200ResponseDataInnerRelationships`

NewGETOrders200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETOrders200ResponseDataInnerRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *GETOrders200ResponseDataInnerRelationships) GetMarket() GETAddresses200ResponseDataInnerRelationshipsGeocoder`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetMarketOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *GETOrders200ResponseDataInnerRelationships) SetMarket(v GETAddresses200ResponseDataInnerRelationshipsGeocoder)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *GETOrders200ResponseDataInnerRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetCustomer

`func (o *GETOrders200ResponseDataInnerRelationships) GetCustomer() GETAddresses200ResponseDataInnerRelationshipsGeocoder`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetCustomerOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *GETOrders200ResponseDataInnerRelationships) SetCustomer(v GETAddresses200ResponseDataInnerRelationshipsGeocoder)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *GETOrders200ResponseDataInnerRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetShippingAddress

`func (o *GETOrders200ResponseDataInnerRelationships) GetShippingAddress() GETAddresses200ResponseDataInnerRelationshipsGeocoder`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetShippingAddressOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *GETOrders200ResponseDataInnerRelationships) SetShippingAddress(v GETAddresses200ResponseDataInnerRelationshipsGeocoder)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *GETOrders200ResponseDataInnerRelationships) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetBillingAddress

`func (o *GETOrders200ResponseDataInnerRelationships) GetBillingAddress() GETAddresses200ResponseDataInnerRelationshipsGeocoder`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetBillingAddressOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *GETOrders200ResponseDataInnerRelationships) SetBillingAddress(v GETAddresses200ResponseDataInnerRelationshipsGeocoder)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *GETOrders200ResponseDataInnerRelationships) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetAvailablePaymentMethods

`func (o *GETOrders200ResponseDataInnerRelationships) GetAvailablePaymentMethods() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetAvailablePaymentMethods returns the AvailablePaymentMethods field if non-nil, zero value otherwise.

### GetAvailablePaymentMethodsOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetAvailablePaymentMethodsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetAvailablePaymentMethodsOk returns a tuple with the AvailablePaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePaymentMethods

`func (o *GETOrders200ResponseDataInnerRelationships) SetAvailablePaymentMethods(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetAvailablePaymentMethods sets AvailablePaymentMethods field to given value.

### HasAvailablePaymentMethods

`func (o *GETOrders200ResponseDataInnerRelationships) HasAvailablePaymentMethods() bool`

HasAvailablePaymentMethods returns a boolean if a field has been set.

### GetAvailableCustomerPaymentSources

`func (o *GETOrders200ResponseDataInnerRelationships) GetAvailableCustomerPaymentSources() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetAvailableCustomerPaymentSources returns the AvailableCustomerPaymentSources field if non-nil, zero value otherwise.

### GetAvailableCustomerPaymentSourcesOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetAvailableCustomerPaymentSourcesOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetAvailableCustomerPaymentSourcesOk returns a tuple with the AvailableCustomerPaymentSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCustomerPaymentSources

`func (o *GETOrders200ResponseDataInnerRelationships) SetAvailableCustomerPaymentSources(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetAvailableCustomerPaymentSources sets AvailableCustomerPaymentSources field to given value.

### HasAvailableCustomerPaymentSources

`func (o *GETOrders200ResponseDataInnerRelationships) HasAvailableCustomerPaymentSources() bool`

HasAvailableCustomerPaymentSources returns a boolean if a field has been set.

### GetAvailableFreeSkus

`func (o *GETOrders200ResponseDataInnerRelationships) GetAvailableFreeSkus() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetAvailableFreeSkus returns the AvailableFreeSkus field if non-nil, zero value otherwise.

### GetAvailableFreeSkusOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetAvailableFreeSkusOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetAvailableFreeSkusOk returns a tuple with the AvailableFreeSkus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFreeSkus

`func (o *GETOrders200ResponseDataInnerRelationships) SetAvailableFreeSkus(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetAvailableFreeSkus sets AvailableFreeSkus field to given value.

### HasAvailableFreeSkus

`func (o *GETOrders200ResponseDataInnerRelationships) HasAvailableFreeSkus() bool`

HasAvailableFreeSkus returns a boolean if a field has been set.

### GetAvailableFreeBundles

`func (o *GETOrders200ResponseDataInnerRelationships) GetAvailableFreeBundles() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetAvailableFreeBundles returns the AvailableFreeBundles field if non-nil, zero value otherwise.

### GetAvailableFreeBundlesOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetAvailableFreeBundlesOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetAvailableFreeBundlesOk returns a tuple with the AvailableFreeBundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFreeBundles

`func (o *GETOrders200ResponseDataInnerRelationships) SetAvailableFreeBundles(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetAvailableFreeBundles sets AvailableFreeBundles field to given value.

### HasAvailableFreeBundles

`func (o *GETOrders200ResponseDataInnerRelationships) HasAvailableFreeBundles() bool`

HasAvailableFreeBundles returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *GETOrders200ResponseDataInnerRelationships) GetPaymentMethod() GETAddresses200ResponseDataInnerRelationshipsGeocoder`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetPaymentMethodOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *GETOrders200ResponseDataInnerRelationships) SetPaymentMethod(v GETAddresses200ResponseDataInnerRelationshipsGeocoder)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *GETOrders200ResponseDataInnerRelationships) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPaymentSource

`func (o *GETOrders200ResponseDataInnerRelationships) GetPaymentSource() GETAddresses200ResponseDataInnerRelationshipsGeocoder`

GetPaymentSource returns the PaymentSource field if non-nil, zero value otherwise.

### GetPaymentSourceOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetPaymentSourceOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool)`

GetPaymentSourceOk returns a tuple with the PaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSource

`func (o *GETOrders200ResponseDataInnerRelationships) SetPaymentSource(v GETAddresses200ResponseDataInnerRelationshipsGeocoder)`

SetPaymentSource sets PaymentSource field to given value.

### HasPaymentSource

`func (o *GETOrders200ResponseDataInnerRelationships) HasPaymentSource() bool`

HasPaymentSource returns a boolean if a field has been set.

### GetLineItems

`func (o *GETOrders200ResponseDataInnerRelationships) GetLineItems() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetLineItemsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *GETOrders200ResponseDataInnerRelationships) SetLineItems(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *GETOrders200ResponseDataInnerRelationships) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetShipments

`func (o *GETOrders200ResponseDataInnerRelationships) GetShipments() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetShipmentsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *GETOrders200ResponseDataInnerRelationships) SetShipments(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetShipments sets Shipments field to given value.

### HasShipments

`func (o *GETOrders200ResponseDataInnerRelationships) HasShipments() bool`

HasShipments returns a boolean if a field has been set.

### GetTransactions

`func (o *GETOrders200ResponseDataInnerRelationships) GetTransactions() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetTransactionsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *GETOrders200ResponseDataInnerRelationships) SetTransactions(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *GETOrders200ResponseDataInnerRelationships) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetAuthorizations

`func (o *GETOrders200ResponseDataInnerRelationships) GetAuthorizations() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetAuthorizations returns the Authorizations field if non-nil, zero value otherwise.

### GetAuthorizationsOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetAuthorizationsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetAuthorizationsOk returns a tuple with the Authorizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizations

`func (o *GETOrders200ResponseDataInnerRelationships) SetAuthorizations(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetAuthorizations sets Authorizations field to given value.

### HasAuthorizations

`func (o *GETOrders200ResponseDataInnerRelationships) HasAuthorizations() bool`

HasAuthorizations returns a boolean if a field has been set.

### GetCaptures

`func (o *GETOrders200ResponseDataInnerRelationships) GetCaptures() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetCaptures returns the Captures field if non-nil, zero value otherwise.

### GetCapturesOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetCapturesOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetCapturesOk returns a tuple with the Captures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptures

`func (o *GETOrders200ResponseDataInnerRelationships) SetCaptures(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetCaptures sets Captures field to given value.

### HasCaptures

`func (o *GETOrders200ResponseDataInnerRelationships) HasCaptures() bool`

HasCaptures returns a boolean if a field has been set.

### GetVoids

`func (o *GETOrders200ResponseDataInnerRelationships) GetVoids() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetVoids returns the Voids field if non-nil, zero value otherwise.

### GetVoidsOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetVoidsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetVoidsOk returns a tuple with the Voids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoids

`func (o *GETOrders200ResponseDataInnerRelationships) SetVoids(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetVoids sets Voids field to given value.

### HasVoids

`func (o *GETOrders200ResponseDataInnerRelationships) HasVoids() bool`

HasVoids returns a boolean if a field has been set.

### GetRefunds

`func (o *GETOrders200ResponseDataInnerRelationships) GetRefunds() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetRefunds returns the Refunds field if non-nil, zero value otherwise.

### GetRefundsOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetRefundsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetRefundsOk returns a tuple with the Refunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunds

`func (o *GETOrders200ResponseDataInnerRelationships) SetRefunds(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetRefunds sets Refunds field to given value.

### HasRefunds

`func (o *GETOrders200ResponseDataInnerRelationships) HasRefunds() bool`

HasRefunds returns a boolean if a field has been set.

### GetOrderSubscriptions

`func (o *GETOrders200ResponseDataInnerRelationships) GetOrderSubscriptions() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetOrderSubscriptions returns the OrderSubscriptions field if non-nil, zero value otherwise.

### GetOrderSubscriptionsOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetOrderSubscriptionsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetOrderSubscriptionsOk returns a tuple with the OrderSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSubscriptions

`func (o *GETOrders200ResponseDataInnerRelationships) SetOrderSubscriptions(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetOrderSubscriptions sets OrderSubscriptions field to given value.

### HasOrderSubscriptions

`func (o *GETOrders200ResponseDataInnerRelationships) HasOrderSubscriptions() bool`

HasOrderSubscriptions returns a boolean if a field has been set.

### GetOrderCopies

`func (o *GETOrders200ResponseDataInnerRelationships) GetOrderCopies() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetOrderCopies returns the OrderCopies field if non-nil, zero value otherwise.

### GetOrderCopiesOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetOrderCopiesOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetOrderCopiesOk returns a tuple with the OrderCopies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCopies

`func (o *GETOrders200ResponseDataInnerRelationships) SetOrderCopies(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetOrderCopies sets OrderCopies field to given value.

### HasOrderCopies

`func (o *GETOrders200ResponseDataInnerRelationships) HasOrderCopies() bool`

HasOrderCopies returns a boolean if a field has been set.

### GetAttachments

`func (o *GETOrders200ResponseDataInnerRelationships) GetAttachments() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GETOrders200ResponseDataInnerRelationships) SetAttachments(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GETOrders200ResponseDataInnerRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *GETOrders200ResponseDataInnerRelationships) GetEvents() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetEventsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GETOrders200ResponseDataInnerRelationships) SetEvents(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *GETOrders200ResponseDataInnerRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


