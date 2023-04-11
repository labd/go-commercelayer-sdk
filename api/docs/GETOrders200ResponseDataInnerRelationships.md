# GETOrders200ResponseDataInnerRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket**](GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket.md) |  | [optional] 
**Customer** | Pointer to [**GETCouponRecipients200ResponseDataInnerRelationshipsCustomer**](GETCouponRecipients200ResponseDataInnerRelationshipsCustomer.md) |  | [optional] 
**ShippingAddress** | Pointer to [**GETOrders200ResponseDataInnerRelationshipsShippingAddress**](GETOrders200ResponseDataInnerRelationshipsShippingAddress.md) |  | [optional] 
**BillingAddress** | Pointer to [**GETOrders200ResponseDataInnerRelationshipsBillingAddress**](GETOrders200ResponseDataInnerRelationshipsBillingAddress.md) |  | [optional] 
**AvailablePaymentMethods** | Pointer to [**GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods**](GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods.md) |  | [optional] 
**AvailableCustomerPaymentSources** | Pointer to [**GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources**](GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources.md) |  | [optional] 
**AvailableFreeSkus** | Pointer to [**GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus**](GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus.md) |  | [optional] 
**AvailableFreeBundles** | Pointer to [**GETOrders200ResponseDataInnerRelationshipsAvailableFreeBundles**](GETOrders200ResponseDataInnerRelationshipsAvailableFreeBundles.md) |  | [optional] 
**PaymentMethod** | Pointer to [**GETOrders200ResponseDataInnerRelationshipsPaymentMethod**](GETOrders200ResponseDataInnerRelationshipsPaymentMethod.md) |  | [optional] 
**PaymentSource** | Pointer to [**GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource**](GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource.md) |  | [optional] 
**LineItems** | Pointer to [**GETOrders200ResponseDataInnerRelationshipsLineItems**](GETOrders200ResponseDataInnerRelationshipsLineItems.md) |  | [optional] 
**Shipments** | Pointer to [**GETOrders200ResponseDataInnerRelationshipsShipments**](GETOrders200ResponseDataInnerRelationshipsShipments.md) |  | [optional] 
**Transactions** | Pointer to [**GETOrders200ResponseDataInnerRelationshipsTransactions**](GETOrders200ResponseDataInnerRelationshipsTransactions.md) |  | [optional] 
**Authorizations** | Pointer to [**GETOrders200ResponseDataInnerRelationshipsAuthorizations**](GETOrders200ResponseDataInnerRelationshipsAuthorizations.md) |  | [optional] 
**Captures** | Pointer to [**GETAuthorizations200ResponseDataInnerRelationshipsCaptures**](GETAuthorizations200ResponseDataInnerRelationshipsCaptures.md) |  | [optional] 
**Voids** | Pointer to [**GETAuthorizations200ResponseDataInnerRelationshipsVoids**](GETAuthorizations200ResponseDataInnerRelationshipsVoids.md) |  | [optional] 
**Refunds** | Pointer to [**GETCaptures200ResponseDataInnerRelationshipsRefunds**](GETCaptures200ResponseDataInnerRelationshipsRefunds.md) |  | [optional] 
**Returns** | Pointer to [**GETCustomers200ResponseDataInnerRelationshipsReturns**](GETCustomers200ResponseDataInnerRelationshipsReturns.md) |  | [optional] 
**OrderSubscriptions** | Pointer to [**GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions**](GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions.md) |  | [optional] 
**OrderFactories** | Pointer to [**GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderFactories**](GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderFactories.md) |  | [optional] 
**OrderCopies** | Pointer to [**GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies**](GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies.md) |  | [optional] 
**RecurringOrderCopies** | Pointer to [**GETOrderSubscriptions200ResponseDataInnerRelationshipsRecurringOrderCopies**](GETOrderSubscriptions200ResponseDataInnerRelationshipsRecurringOrderCopies.md) |  | [optional] 
**Attachments** | Pointer to [**GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments**](GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments.md) |  | [optional] 
**Events** | Pointer to [**GETAuthorizations200ResponseDataInnerRelationshipsEvents**](GETAuthorizations200ResponseDataInnerRelationshipsEvents.md) |  | [optional] 

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

`func (o *GETOrders200ResponseDataInnerRelationships) GetMarket() GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetMarketOk() (*GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *GETOrders200ResponseDataInnerRelationships) SetMarket(v GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *GETOrders200ResponseDataInnerRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetCustomer

`func (o *GETOrders200ResponseDataInnerRelationships) GetCustomer() GETCouponRecipients200ResponseDataInnerRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetCustomerOk() (*GETCouponRecipients200ResponseDataInnerRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *GETOrders200ResponseDataInnerRelationships) SetCustomer(v GETCouponRecipients200ResponseDataInnerRelationshipsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *GETOrders200ResponseDataInnerRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetShippingAddress

`func (o *GETOrders200ResponseDataInnerRelationships) GetShippingAddress() GETOrders200ResponseDataInnerRelationshipsShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetShippingAddressOk() (*GETOrders200ResponseDataInnerRelationshipsShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *GETOrders200ResponseDataInnerRelationships) SetShippingAddress(v GETOrders200ResponseDataInnerRelationshipsShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *GETOrders200ResponseDataInnerRelationships) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetBillingAddress

`func (o *GETOrders200ResponseDataInnerRelationships) GetBillingAddress() GETOrders200ResponseDataInnerRelationshipsBillingAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetBillingAddressOk() (*GETOrders200ResponseDataInnerRelationshipsBillingAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *GETOrders200ResponseDataInnerRelationships) SetBillingAddress(v GETOrders200ResponseDataInnerRelationshipsBillingAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *GETOrders200ResponseDataInnerRelationships) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetAvailablePaymentMethods

`func (o *GETOrders200ResponseDataInnerRelationships) GetAvailablePaymentMethods() GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods`

GetAvailablePaymentMethods returns the AvailablePaymentMethods field if non-nil, zero value otherwise.

### GetAvailablePaymentMethodsOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetAvailablePaymentMethodsOk() (*GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods, bool)`

GetAvailablePaymentMethodsOk returns a tuple with the AvailablePaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePaymentMethods

`func (o *GETOrders200ResponseDataInnerRelationships) SetAvailablePaymentMethods(v GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethods)`

SetAvailablePaymentMethods sets AvailablePaymentMethods field to given value.

### HasAvailablePaymentMethods

`func (o *GETOrders200ResponseDataInnerRelationships) HasAvailablePaymentMethods() bool`

HasAvailablePaymentMethods returns a boolean if a field has been set.

### GetAvailableCustomerPaymentSources

`func (o *GETOrders200ResponseDataInnerRelationships) GetAvailableCustomerPaymentSources() GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources`

GetAvailableCustomerPaymentSources returns the AvailableCustomerPaymentSources field if non-nil, zero value otherwise.

### GetAvailableCustomerPaymentSourcesOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetAvailableCustomerPaymentSourcesOk() (*GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources, bool)`

GetAvailableCustomerPaymentSourcesOk returns a tuple with the AvailableCustomerPaymentSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCustomerPaymentSources

`func (o *GETOrders200ResponseDataInnerRelationships) SetAvailableCustomerPaymentSources(v GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources)`

SetAvailableCustomerPaymentSources sets AvailableCustomerPaymentSources field to given value.

### HasAvailableCustomerPaymentSources

`func (o *GETOrders200ResponseDataInnerRelationships) HasAvailableCustomerPaymentSources() bool`

HasAvailableCustomerPaymentSources returns a boolean if a field has been set.

### GetAvailableFreeSkus

`func (o *GETOrders200ResponseDataInnerRelationships) GetAvailableFreeSkus() GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus`

GetAvailableFreeSkus returns the AvailableFreeSkus field if non-nil, zero value otherwise.

### GetAvailableFreeSkusOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetAvailableFreeSkusOk() (*GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus, bool)`

GetAvailableFreeSkusOk returns a tuple with the AvailableFreeSkus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFreeSkus

`func (o *GETOrders200ResponseDataInnerRelationships) SetAvailableFreeSkus(v GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus)`

SetAvailableFreeSkus sets AvailableFreeSkus field to given value.

### HasAvailableFreeSkus

`func (o *GETOrders200ResponseDataInnerRelationships) HasAvailableFreeSkus() bool`

HasAvailableFreeSkus returns a boolean if a field has been set.

### GetAvailableFreeBundles

`func (o *GETOrders200ResponseDataInnerRelationships) GetAvailableFreeBundles() GETOrders200ResponseDataInnerRelationshipsAvailableFreeBundles`

GetAvailableFreeBundles returns the AvailableFreeBundles field if non-nil, zero value otherwise.

### GetAvailableFreeBundlesOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetAvailableFreeBundlesOk() (*GETOrders200ResponseDataInnerRelationshipsAvailableFreeBundles, bool)`

GetAvailableFreeBundlesOk returns a tuple with the AvailableFreeBundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFreeBundles

`func (o *GETOrders200ResponseDataInnerRelationships) SetAvailableFreeBundles(v GETOrders200ResponseDataInnerRelationshipsAvailableFreeBundles)`

SetAvailableFreeBundles sets AvailableFreeBundles field to given value.

### HasAvailableFreeBundles

`func (o *GETOrders200ResponseDataInnerRelationships) HasAvailableFreeBundles() bool`

HasAvailableFreeBundles returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *GETOrders200ResponseDataInnerRelationships) GetPaymentMethod() GETOrders200ResponseDataInnerRelationshipsPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetPaymentMethodOk() (*GETOrders200ResponseDataInnerRelationshipsPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *GETOrders200ResponseDataInnerRelationships) SetPaymentMethod(v GETOrders200ResponseDataInnerRelationshipsPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *GETOrders200ResponseDataInnerRelationships) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPaymentSource

`func (o *GETOrders200ResponseDataInnerRelationships) GetPaymentSource() GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource`

GetPaymentSource returns the PaymentSource field if non-nil, zero value otherwise.

### GetPaymentSourceOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetPaymentSourceOk() (*GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource, bool)`

GetPaymentSourceOk returns a tuple with the PaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSource

`func (o *GETOrders200ResponseDataInnerRelationships) SetPaymentSource(v GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource)`

SetPaymentSource sets PaymentSource field to given value.

### HasPaymentSource

`func (o *GETOrders200ResponseDataInnerRelationships) HasPaymentSource() bool`

HasPaymentSource returns a boolean if a field has been set.

### GetLineItems

`func (o *GETOrders200ResponseDataInnerRelationships) GetLineItems() GETOrders200ResponseDataInnerRelationshipsLineItems`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetLineItemsOk() (*GETOrders200ResponseDataInnerRelationshipsLineItems, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *GETOrders200ResponseDataInnerRelationships) SetLineItems(v GETOrders200ResponseDataInnerRelationshipsLineItems)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *GETOrders200ResponseDataInnerRelationships) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetShipments

`func (o *GETOrders200ResponseDataInnerRelationships) GetShipments() GETOrders200ResponseDataInnerRelationshipsShipments`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetShipmentsOk() (*GETOrders200ResponseDataInnerRelationshipsShipments, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *GETOrders200ResponseDataInnerRelationships) SetShipments(v GETOrders200ResponseDataInnerRelationshipsShipments)`

SetShipments sets Shipments field to given value.

### HasShipments

`func (o *GETOrders200ResponseDataInnerRelationships) HasShipments() bool`

HasShipments returns a boolean if a field has been set.

### GetTransactions

`func (o *GETOrders200ResponseDataInnerRelationships) GetTransactions() GETOrders200ResponseDataInnerRelationshipsTransactions`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetTransactionsOk() (*GETOrders200ResponseDataInnerRelationshipsTransactions, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *GETOrders200ResponseDataInnerRelationships) SetTransactions(v GETOrders200ResponseDataInnerRelationshipsTransactions)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *GETOrders200ResponseDataInnerRelationships) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetAuthorizations

`func (o *GETOrders200ResponseDataInnerRelationships) GetAuthorizations() GETOrders200ResponseDataInnerRelationshipsAuthorizations`

GetAuthorizations returns the Authorizations field if non-nil, zero value otherwise.

### GetAuthorizationsOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetAuthorizationsOk() (*GETOrders200ResponseDataInnerRelationshipsAuthorizations, bool)`

GetAuthorizationsOk returns a tuple with the Authorizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizations

`func (o *GETOrders200ResponseDataInnerRelationships) SetAuthorizations(v GETOrders200ResponseDataInnerRelationshipsAuthorizations)`

SetAuthorizations sets Authorizations field to given value.

### HasAuthorizations

`func (o *GETOrders200ResponseDataInnerRelationships) HasAuthorizations() bool`

HasAuthorizations returns a boolean if a field has been set.

### GetCaptures

`func (o *GETOrders200ResponseDataInnerRelationships) GetCaptures() GETAuthorizations200ResponseDataInnerRelationshipsCaptures`

GetCaptures returns the Captures field if non-nil, zero value otherwise.

### GetCapturesOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetCapturesOk() (*GETAuthorizations200ResponseDataInnerRelationshipsCaptures, bool)`

GetCapturesOk returns a tuple with the Captures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptures

`func (o *GETOrders200ResponseDataInnerRelationships) SetCaptures(v GETAuthorizations200ResponseDataInnerRelationshipsCaptures)`

SetCaptures sets Captures field to given value.

### HasCaptures

`func (o *GETOrders200ResponseDataInnerRelationships) HasCaptures() bool`

HasCaptures returns a boolean if a field has been set.

### GetVoids

`func (o *GETOrders200ResponseDataInnerRelationships) GetVoids() GETAuthorizations200ResponseDataInnerRelationshipsVoids`

GetVoids returns the Voids field if non-nil, zero value otherwise.

### GetVoidsOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetVoidsOk() (*GETAuthorizations200ResponseDataInnerRelationshipsVoids, bool)`

GetVoidsOk returns a tuple with the Voids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoids

`func (o *GETOrders200ResponseDataInnerRelationships) SetVoids(v GETAuthorizations200ResponseDataInnerRelationshipsVoids)`

SetVoids sets Voids field to given value.

### HasVoids

`func (o *GETOrders200ResponseDataInnerRelationships) HasVoids() bool`

HasVoids returns a boolean if a field has been set.

### GetRefunds

`func (o *GETOrders200ResponseDataInnerRelationships) GetRefunds() GETCaptures200ResponseDataInnerRelationshipsRefunds`

GetRefunds returns the Refunds field if non-nil, zero value otherwise.

### GetRefundsOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetRefundsOk() (*GETCaptures200ResponseDataInnerRelationshipsRefunds, bool)`

GetRefundsOk returns a tuple with the Refunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunds

`func (o *GETOrders200ResponseDataInnerRelationships) SetRefunds(v GETCaptures200ResponseDataInnerRelationshipsRefunds)`

SetRefunds sets Refunds field to given value.

### HasRefunds

`func (o *GETOrders200ResponseDataInnerRelationships) HasRefunds() bool`

HasRefunds returns a boolean if a field has been set.

### GetReturns

`func (o *GETOrders200ResponseDataInnerRelationships) GetReturns() GETCustomers200ResponseDataInnerRelationshipsReturns`

GetReturns returns the Returns field if non-nil, zero value otherwise.

### GetReturnsOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetReturnsOk() (*GETCustomers200ResponseDataInnerRelationshipsReturns, bool)`

GetReturnsOk returns a tuple with the Returns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturns

`func (o *GETOrders200ResponseDataInnerRelationships) SetReturns(v GETCustomers200ResponseDataInnerRelationshipsReturns)`

SetReturns sets Returns field to given value.

### HasReturns

`func (o *GETOrders200ResponseDataInnerRelationships) HasReturns() bool`

HasReturns returns a boolean if a field has been set.

### GetOrderSubscriptions

`func (o *GETOrders200ResponseDataInnerRelationships) GetOrderSubscriptions() GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions`

GetOrderSubscriptions returns the OrderSubscriptions field if non-nil, zero value otherwise.

### GetOrderSubscriptionsOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetOrderSubscriptionsOk() (*GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions, bool)`

GetOrderSubscriptionsOk returns a tuple with the OrderSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSubscriptions

`func (o *GETOrders200ResponseDataInnerRelationships) SetOrderSubscriptions(v GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions)`

SetOrderSubscriptions sets OrderSubscriptions field to given value.

### HasOrderSubscriptions

`func (o *GETOrders200ResponseDataInnerRelationships) HasOrderSubscriptions() bool`

HasOrderSubscriptions returns a boolean if a field has been set.

### GetOrderFactories

`func (o *GETOrders200ResponseDataInnerRelationships) GetOrderFactories() GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderFactories`

GetOrderFactories returns the OrderFactories field if non-nil, zero value otherwise.

### GetOrderFactoriesOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetOrderFactoriesOk() (*GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderFactories, bool)`

GetOrderFactoriesOk returns a tuple with the OrderFactories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderFactories

`func (o *GETOrders200ResponseDataInnerRelationships) SetOrderFactories(v GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderFactories)`

SetOrderFactories sets OrderFactories field to given value.

### HasOrderFactories

`func (o *GETOrders200ResponseDataInnerRelationships) HasOrderFactories() bool`

HasOrderFactories returns a boolean if a field has been set.

### GetOrderCopies

`func (o *GETOrders200ResponseDataInnerRelationships) GetOrderCopies() GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies`

GetOrderCopies returns the OrderCopies field if non-nil, zero value otherwise.

### GetOrderCopiesOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetOrderCopiesOk() (*GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies, bool)`

GetOrderCopiesOk returns a tuple with the OrderCopies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCopies

`func (o *GETOrders200ResponseDataInnerRelationships) SetOrderCopies(v GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies)`

SetOrderCopies sets OrderCopies field to given value.

### HasOrderCopies

`func (o *GETOrders200ResponseDataInnerRelationships) HasOrderCopies() bool`

HasOrderCopies returns a boolean if a field has been set.

### GetRecurringOrderCopies

`func (o *GETOrders200ResponseDataInnerRelationships) GetRecurringOrderCopies() GETOrderSubscriptions200ResponseDataInnerRelationshipsRecurringOrderCopies`

GetRecurringOrderCopies returns the RecurringOrderCopies field if non-nil, zero value otherwise.

### GetRecurringOrderCopiesOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetRecurringOrderCopiesOk() (*GETOrderSubscriptions200ResponseDataInnerRelationshipsRecurringOrderCopies, bool)`

GetRecurringOrderCopiesOk returns a tuple with the RecurringOrderCopies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringOrderCopies

`func (o *GETOrders200ResponseDataInnerRelationships) SetRecurringOrderCopies(v GETOrderSubscriptions200ResponseDataInnerRelationshipsRecurringOrderCopies)`

SetRecurringOrderCopies sets RecurringOrderCopies field to given value.

### HasRecurringOrderCopies

`func (o *GETOrders200ResponseDataInnerRelationships) HasRecurringOrderCopies() bool`

HasRecurringOrderCopies returns a boolean if a field has been set.

### GetAttachments

`func (o *GETOrders200ResponseDataInnerRelationships) GetAttachments() GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GETOrders200ResponseDataInnerRelationships) SetAttachments(v GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GETOrders200ResponseDataInnerRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *GETOrders200ResponseDataInnerRelationships) GetEvents() GETAuthorizations200ResponseDataInnerRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GETOrders200ResponseDataInnerRelationships) GetEventsOk() (*GETAuthorizations200ResponseDataInnerRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GETOrders200ResponseDataInnerRelationships) SetEvents(v GETAuthorizations200ResponseDataInnerRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *GETOrders200ResponseDataInnerRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


