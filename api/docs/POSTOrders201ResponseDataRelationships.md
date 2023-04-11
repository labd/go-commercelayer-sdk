# POSTOrders201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket**](POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket.md) |  | [optional] 
**Customer** | Pointer to [**POSTCouponRecipients201ResponseDataRelationshipsCustomer**](POSTCouponRecipients201ResponseDataRelationshipsCustomer.md) |  | [optional] 
**ShippingAddress** | Pointer to [**POSTOrders201ResponseDataRelationshipsShippingAddress**](POSTOrders201ResponseDataRelationshipsShippingAddress.md) |  | [optional] 
**BillingAddress** | Pointer to [**POSTOrders201ResponseDataRelationshipsBillingAddress**](POSTOrders201ResponseDataRelationshipsBillingAddress.md) |  | [optional] 
**AvailablePaymentMethods** | Pointer to [**POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods**](POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods.md) |  | [optional] 
**AvailableCustomerPaymentSources** | Pointer to [**POSTOrders201ResponseDataRelationshipsAvailableCustomerPaymentSources**](POSTOrders201ResponseDataRelationshipsAvailableCustomerPaymentSources.md) |  | [optional] 
**AvailableFreeSkus** | Pointer to [**POSTOrders201ResponseDataRelationshipsAvailableFreeSkus**](POSTOrders201ResponseDataRelationshipsAvailableFreeSkus.md) |  | [optional] 
**AvailableFreeBundles** | Pointer to [**POSTOrders201ResponseDataRelationshipsAvailableFreeBundles**](POSTOrders201ResponseDataRelationshipsAvailableFreeBundles.md) |  | [optional] 
**PaymentMethod** | Pointer to [**POSTOrders201ResponseDataRelationshipsPaymentMethod**](POSTOrders201ResponseDataRelationshipsPaymentMethod.md) |  | [optional] 
**PaymentSource** | Pointer to [**POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource**](POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource.md) |  | [optional] 
**LineItems** | Pointer to [**POSTOrders201ResponseDataRelationshipsLineItems**](POSTOrders201ResponseDataRelationshipsLineItems.md) |  | [optional] 
**Shipments** | Pointer to [**POSTOrders201ResponseDataRelationshipsShipments**](POSTOrders201ResponseDataRelationshipsShipments.md) |  | [optional] 
**Transactions** | Pointer to [**POSTOrders201ResponseDataRelationshipsTransactions**](POSTOrders201ResponseDataRelationshipsTransactions.md) |  | [optional] 
**Authorizations** | Pointer to [**POSTOrders201ResponseDataRelationshipsAuthorizations**](POSTOrders201ResponseDataRelationshipsAuthorizations.md) |  | [optional] 
**Captures** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsCaptures**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsCaptures.md) |  | [optional] 
**Voids** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids.md) |  | [optional] 
**Refunds** | Pointer to [**GETCapturesCaptureId200ResponseDataRelationshipsRefunds**](GETCapturesCaptureId200ResponseDataRelationshipsRefunds.md) |  | [optional] 
**Returns** | Pointer to [**POSTCustomers201ResponseDataRelationshipsReturns**](POSTCustomers201ResponseDataRelationshipsReturns.md) |  | [optional] 
**OrderSubscriptions** | Pointer to [**POSTCustomers201ResponseDataRelationshipsOrderSubscriptions**](POSTCustomers201ResponseDataRelationshipsOrderSubscriptions.md) |  | [optional] 
**OrderFactories** | Pointer to [**POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactories**](POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactories.md) |  | [optional] 
**OrderCopies** | Pointer to [**POSTOrderSubscriptions201ResponseDataRelationshipsOrderCopies**](POSTOrderSubscriptions201ResponseDataRelationshipsOrderCopies.md) |  | [optional] 
**RecurringOrderCopies** | Pointer to [**POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopies**](POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopies.md) |  | [optional] 
**Attachments** | Pointer to [**POSTAvalaraAccounts201ResponseDataRelationshipsAttachments**](POSTAvalaraAccounts201ResponseDataRelationshipsAttachments.md) |  | [optional] 
**Events** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents.md) |  | [optional] 

## Methods

### NewPOSTOrders201ResponseDataRelationships

`func NewPOSTOrders201ResponseDataRelationships() *POSTOrders201ResponseDataRelationships`

NewPOSTOrders201ResponseDataRelationships instantiates a new POSTOrders201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTOrders201ResponseDataRelationshipsWithDefaults

`func NewPOSTOrders201ResponseDataRelationshipsWithDefaults() *POSTOrders201ResponseDataRelationships`

NewPOSTOrders201ResponseDataRelationshipsWithDefaults instantiates a new POSTOrders201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *POSTOrders201ResponseDataRelationships) GetMarket() POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *POSTOrders201ResponseDataRelationships) GetMarketOk() (*POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *POSTOrders201ResponseDataRelationships) SetMarket(v POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *POSTOrders201ResponseDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetCustomer

`func (o *POSTOrders201ResponseDataRelationships) GetCustomer() POSTCouponRecipients201ResponseDataRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *POSTOrders201ResponseDataRelationships) GetCustomerOk() (*POSTCouponRecipients201ResponseDataRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *POSTOrders201ResponseDataRelationships) SetCustomer(v POSTCouponRecipients201ResponseDataRelationshipsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *POSTOrders201ResponseDataRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetShippingAddress

`func (o *POSTOrders201ResponseDataRelationships) GetShippingAddress() POSTOrders201ResponseDataRelationshipsShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *POSTOrders201ResponseDataRelationships) GetShippingAddressOk() (*POSTOrders201ResponseDataRelationshipsShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *POSTOrders201ResponseDataRelationships) SetShippingAddress(v POSTOrders201ResponseDataRelationshipsShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *POSTOrders201ResponseDataRelationships) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetBillingAddress

`func (o *POSTOrders201ResponseDataRelationships) GetBillingAddress() POSTOrders201ResponseDataRelationshipsBillingAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *POSTOrders201ResponseDataRelationships) GetBillingAddressOk() (*POSTOrders201ResponseDataRelationshipsBillingAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *POSTOrders201ResponseDataRelationships) SetBillingAddress(v POSTOrders201ResponseDataRelationshipsBillingAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *POSTOrders201ResponseDataRelationships) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetAvailablePaymentMethods

`func (o *POSTOrders201ResponseDataRelationships) GetAvailablePaymentMethods() POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods`

GetAvailablePaymentMethods returns the AvailablePaymentMethods field if non-nil, zero value otherwise.

### GetAvailablePaymentMethodsOk

`func (o *POSTOrders201ResponseDataRelationships) GetAvailablePaymentMethodsOk() (*POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods, bool)`

GetAvailablePaymentMethodsOk returns a tuple with the AvailablePaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePaymentMethods

`func (o *POSTOrders201ResponseDataRelationships) SetAvailablePaymentMethods(v POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods)`

SetAvailablePaymentMethods sets AvailablePaymentMethods field to given value.

### HasAvailablePaymentMethods

`func (o *POSTOrders201ResponseDataRelationships) HasAvailablePaymentMethods() bool`

HasAvailablePaymentMethods returns a boolean if a field has been set.

### GetAvailableCustomerPaymentSources

`func (o *POSTOrders201ResponseDataRelationships) GetAvailableCustomerPaymentSources() POSTOrders201ResponseDataRelationshipsAvailableCustomerPaymentSources`

GetAvailableCustomerPaymentSources returns the AvailableCustomerPaymentSources field if non-nil, zero value otherwise.

### GetAvailableCustomerPaymentSourcesOk

`func (o *POSTOrders201ResponseDataRelationships) GetAvailableCustomerPaymentSourcesOk() (*POSTOrders201ResponseDataRelationshipsAvailableCustomerPaymentSources, bool)`

GetAvailableCustomerPaymentSourcesOk returns a tuple with the AvailableCustomerPaymentSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCustomerPaymentSources

`func (o *POSTOrders201ResponseDataRelationships) SetAvailableCustomerPaymentSources(v POSTOrders201ResponseDataRelationshipsAvailableCustomerPaymentSources)`

SetAvailableCustomerPaymentSources sets AvailableCustomerPaymentSources field to given value.

### HasAvailableCustomerPaymentSources

`func (o *POSTOrders201ResponseDataRelationships) HasAvailableCustomerPaymentSources() bool`

HasAvailableCustomerPaymentSources returns a boolean if a field has been set.

### GetAvailableFreeSkus

`func (o *POSTOrders201ResponseDataRelationships) GetAvailableFreeSkus() POSTOrders201ResponseDataRelationshipsAvailableFreeSkus`

GetAvailableFreeSkus returns the AvailableFreeSkus field if non-nil, zero value otherwise.

### GetAvailableFreeSkusOk

`func (o *POSTOrders201ResponseDataRelationships) GetAvailableFreeSkusOk() (*POSTOrders201ResponseDataRelationshipsAvailableFreeSkus, bool)`

GetAvailableFreeSkusOk returns a tuple with the AvailableFreeSkus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFreeSkus

`func (o *POSTOrders201ResponseDataRelationships) SetAvailableFreeSkus(v POSTOrders201ResponseDataRelationshipsAvailableFreeSkus)`

SetAvailableFreeSkus sets AvailableFreeSkus field to given value.

### HasAvailableFreeSkus

`func (o *POSTOrders201ResponseDataRelationships) HasAvailableFreeSkus() bool`

HasAvailableFreeSkus returns a boolean if a field has been set.

### GetAvailableFreeBundles

`func (o *POSTOrders201ResponseDataRelationships) GetAvailableFreeBundles() POSTOrders201ResponseDataRelationshipsAvailableFreeBundles`

GetAvailableFreeBundles returns the AvailableFreeBundles field if non-nil, zero value otherwise.

### GetAvailableFreeBundlesOk

`func (o *POSTOrders201ResponseDataRelationships) GetAvailableFreeBundlesOk() (*POSTOrders201ResponseDataRelationshipsAvailableFreeBundles, bool)`

GetAvailableFreeBundlesOk returns a tuple with the AvailableFreeBundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFreeBundles

`func (o *POSTOrders201ResponseDataRelationships) SetAvailableFreeBundles(v POSTOrders201ResponseDataRelationshipsAvailableFreeBundles)`

SetAvailableFreeBundles sets AvailableFreeBundles field to given value.

### HasAvailableFreeBundles

`func (o *POSTOrders201ResponseDataRelationships) HasAvailableFreeBundles() bool`

HasAvailableFreeBundles returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *POSTOrders201ResponseDataRelationships) GetPaymentMethod() POSTOrders201ResponseDataRelationshipsPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *POSTOrders201ResponseDataRelationships) GetPaymentMethodOk() (*POSTOrders201ResponseDataRelationshipsPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *POSTOrders201ResponseDataRelationships) SetPaymentMethod(v POSTOrders201ResponseDataRelationshipsPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *POSTOrders201ResponseDataRelationships) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPaymentSource

`func (o *POSTOrders201ResponseDataRelationships) GetPaymentSource() POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource`

GetPaymentSource returns the PaymentSource field if non-nil, zero value otherwise.

### GetPaymentSourceOk

`func (o *POSTOrders201ResponseDataRelationships) GetPaymentSourceOk() (*POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource, bool)`

GetPaymentSourceOk returns a tuple with the PaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSource

`func (o *POSTOrders201ResponseDataRelationships) SetPaymentSource(v POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource)`

SetPaymentSource sets PaymentSource field to given value.

### HasPaymentSource

`func (o *POSTOrders201ResponseDataRelationships) HasPaymentSource() bool`

HasPaymentSource returns a boolean if a field has been set.

### GetLineItems

`func (o *POSTOrders201ResponseDataRelationships) GetLineItems() POSTOrders201ResponseDataRelationshipsLineItems`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *POSTOrders201ResponseDataRelationships) GetLineItemsOk() (*POSTOrders201ResponseDataRelationshipsLineItems, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *POSTOrders201ResponseDataRelationships) SetLineItems(v POSTOrders201ResponseDataRelationshipsLineItems)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *POSTOrders201ResponseDataRelationships) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetShipments

`func (o *POSTOrders201ResponseDataRelationships) GetShipments() POSTOrders201ResponseDataRelationshipsShipments`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *POSTOrders201ResponseDataRelationships) GetShipmentsOk() (*POSTOrders201ResponseDataRelationshipsShipments, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *POSTOrders201ResponseDataRelationships) SetShipments(v POSTOrders201ResponseDataRelationshipsShipments)`

SetShipments sets Shipments field to given value.

### HasShipments

`func (o *POSTOrders201ResponseDataRelationships) HasShipments() bool`

HasShipments returns a boolean if a field has been set.

### GetTransactions

`func (o *POSTOrders201ResponseDataRelationships) GetTransactions() POSTOrders201ResponseDataRelationshipsTransactions`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *POSTOrders201ResponseDataRelationships) GetTransactionsOk() (*POSTOrders201ResponseDataRelationshipsTransactions, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *POSTOrders201ResponseDataRelationships) SetTransactions(v POSTOrders201ResponseDataRelationshipsTransactions)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *POSTOrders201ResponseDataRelationships) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetAuthorizations

`func (o *POSTOrders201ResponseDataRelationships) GetAuthorizations() POSTOrders201ResponseDataRelationshipsAuthorizations`

GetAuthorizations returns the Authorizations field if non-nil, zero value otherwise.

### GetAuthorizationsOk

`func (o *POSTOrders201ResponseDataRelationships) GetAuthorizationsOk() (*POSTOrders201ResponseDataRelationshipsAuthorizations, bool)`

GetAuthorizationsOk returns a tuple with the Authorizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizations

`func (o *POSTOrders201ResponseDataRelationships) SetAuthorizations(v POSTOrders201ResponseDataRelationshipsAuthorizations)`

SetAuthorizations sets Authorizations field to given value.

### HasAuthorizations

`func (o *POSTOrders201ResponseDataRelationships) HasAuthorizations() bool`

HasAuthorizations returns a boolean if a field has been set.

### GetCaptures

`func (o *POSTOrders201ResponseDataRelationships) GetCaptures() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsCaptures`

GetCaptures returns the Captures field if non-nil, zero value otherwise.

### GetCapturesOk

`func (o *POSTOrders201ResponseDataRelationships) GetCapturesOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsCaptures, bool)`

GetCapturesOk returns a tuple with the Captures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptures

`func (o *POSTOrders201ResponseDataRelationships) SetCaptures(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsCaptures)`

SetCaptures sets Captures field to given value.

### HasCaptures

`func (o *POSTOrders201ResponseDataRelationships) HasCaptures() bool`

HasCaptures returns a boolean if a field has been set.

### GetVoids

`func (o *POSTOrders201ResponseDataRelationships) GetVoids() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids`

GetVoids returns the Voids field if non-nil, zero value otherwise.

### GetVoidsOk

`func (o *POSTOrders201ResponseDataRelationships) GetVoidsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids, bool)`

GetVoidsOk returns a tuple with the Voids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoids

`func (o *POSTOrders201ResponseDataRelationships) SetVoids(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids)`

SetVoids sets Voids field to given value.

### HasVoids

`func (o *POSTOrders201ResponseDataRelationships) HasVoids() bool`

HasVoids returns a boolean if a field has been set.

### GetRefunds

`func (o *POSTOrders201ResponseDataRelationships) GetRefunds() GETCapturesCaptureId200ResponseDataRelationshipsRefunds`

GetRefunds returns the Refunds field if non-nil, zero value otherwise.

### GetRefundsOk

`func (o *POSTOrders201ResponseDataRelationships) GetRefundsOk() (*GETCapturesCaptureId200ResponseDataRelationshipsRefunds, bool)`

GetRefundsOk returns a tuple with the Refunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunds

`func (o *POSTOrders201ResponseDataRelationships) SetRefunds(v GETCapturesCaptureId200ResponseDataRelationshipsRefunds)`

SetRefunds sets Refunds field to given value.

### HasRefunds

`func (o *POSTOrders201ResponseDataRelationships) HasRefunds() bool`

HasRefunds returns a boolean if a field has been set.

### GetReturns

`func (o *POSTOrders201ResponseDataRelationships) GetReturns() POSTCustomers201ResponseDataRelationshipsReturns`

GetReturns returns the Returns field if non-nil, zero value otherwise.

### GetReturnsOk

`func (o *POSTOrders201ResponseDataRelationships) GetReturnsOk() (*POSTCustomers201ResponseDataRelationshipsReturns, bool)`

GetReturnsOk returns a tuple with the Returns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturns

`func (o *POSTOrders201ResponseDataRelationships) SetReturns(v POSTCustomers201ResponseDataRelationshipsReturns)`

SetReturns sets Returns field to given value.

### HasReturns

`func (o *POSTOrders201ResponseDataRelationships) HasReturns() bool`

HasReturns returns a boolean if a field has been set.

### GetOrderSubscriptions

`func (o *POSTOrders201ResponseDataRelationships) GetOrderSubscriptions() POSTCustomers201ResponseDataRelationshipsOrderSubscriptions`

GetOrderSubscriptions returns the OrderSubscriptions field if non-nil, zero value otherwise.

### GetOrderSubscriptionsOk

`func (o *POSTOrders201ResponseDataRelationships) GetOrderSubscriptionsOk() (*POSTCustomers201ResponseDataRelationshipsOrderSubscriptions, bool)`

GetOrderSubscriptionsOk returns a tuple with the OrderSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSubscriptions

`func (o *POSTOrders201ResponseDataRelationships) SetOrderSubscriptions(v POSTCustomers201ResponseDataRelationshipsOrderSubscriptions)`

SetOrderSubscriptions sets OrderSubscriptions field to given value.

### HasOrderSubscriptions

`func (o *POSTOrders201ResponseDataRelationships) HasOrderSubscriptions() bool`

HasOrderSubscriptions returns a boolean if a field has been set.

### GetOrderFactories

`func (o *POSTOrders201ResponseDataRelationships) GetOrderFactories() POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactories`

GetOrderFactories returns the OrderFactories field if non-nil, zero value otherwise.

### GetOrderFactoriesOk

`func (o *POSTOrders201ResponseDataRelationships) GetOrderFactoriesOk() (*POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactories, bool)`

GetOrderFactoriesOk returns a tuple with the OrderFactories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderFactories

`func (o *POSTOrders201ResponseDataRelationships) SetOrderFactories(v POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactories)`

SetOrderFactories sets OrderFactories field to given value.

### HasOrderFactories

`func (o *POSTOrders201ResponseDataRelationships) HasOrderFactories() bool`

HasOrderFactories returns a boolean if a field has been set.

### GetOrderCopies

`func (o *POSTOrders201ResponseDataRelationships) GetOrderCopies() POSTOrderSubscriptions201ResponseDataRelationshipsOrderCopies`

GetOrderCopies returns the OrderCopies field if non-nil, zero value otherwise.

### GetOrderCopiesOk

`func (o *POSTOrders201ResponseDataRelationships) GetOrderCopiesOk() (*POSTOrderSubscriptions201ResponseDataRelationshipsOrderCopies, bool)`

GetOrderCopiesOk returns a tuple with the OrderCopies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCopies

`func (o *POSTOrders201ResponseDataRelationships) SetOrderCopies(v POSTOrderSubscriptions201ResponseDataRelationshipsOrderCopies)`

SetOrderCopies sets OrderCopies field to given value.

### HasOrderCopies

`func (o *POSTOrders201ResponseDataRelationships) HasOrderCopies() bool`

HasOrderCopies returns a boolean if a field has been set.

### GetRecurringOrderCopies

`func (o *POSTOrders201ResponseDataRelationships) GetRecurringOrderCopies() POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopies`

GetRecurringOrderCopies returns the RecurringOrderCopies field if non-nil, zero value otherwise.

### GetRecurringOrderCopiesOk

`func (o *POSTOrders201ResponseDataRelationships) GetRecurringOrderCopiesOk() (*POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopies, bool)`

GetRecurringOrderCopiesOk returns a tuple with the RecurringOrderCopies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringOrderCopies

`func (o *POSTOrders201ResponseDataRelationships) SetRecurringOrderCopies(v POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopies)`

SetRecurringOrderCopies sets RecurringOrderCopies field to given value.

### HasRecurringOrderCopies

`func (o *POSTOrders201ResponseDataRelationships) HasRecurringOrderCopies() bool`

HasRecurringOrderCopies returns a boolean if a field has been set.

### GetAttachments

`func (o *POSTOrders201ResponseDataRelationships) GetAttachments() POSTAvalaraAccounts201ResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *POSTOrders201ResponseDataRelationships) GetAttachmentsOk() (*POSTAvalaraAccounts201ResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *POSTOrders201ResponseDataRelationships) SetAttachments(v POSTAvalaraAccounts201ResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *POSTOrders201ResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *POSTOrders201ResponseDataRelationships) GetEvents() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *POSTOrders201ResponseDataRelationships) GetEventsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *POSTOrders201ResponseDataRelationships) SetEvents(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *POSTOrders201ResponseDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


