# OrderResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**BillingInfoValidationRuleResponseDataRelationshipsMarket**](BillingInfoValidationRuleResponseDataRelationshipsMarket.md) |  | [optional] 
**Customer** | Pointer to [**CouponRecipientResponseDataRelationshipsCustomer**](CouponRecipientResponseDataRelationshipsCustomer.md) |  | [optional] 
**ShippingAddress** | Pointer to [**OrderResponseDataRelationshipsShippingAddress**](OrderResponseDataRelationshipsShippingAddress.md) |  | [optional] 
**BillingAddress** | Pointer to [**OrderResponseDataRelationshipsBillingAddress**](OrderResponseDataRelationshipsBillingAddress.md) |  | [optional] 
**AvailablePaymentMethods** | Pointer to [**OrderResponseDataRelationshipsAvailablePaymentMethods**](OrderResponseDataRelationshipsAvailablePaymentMethods.md) |  | [optional] 
**AvailableCustomerPaymentSources** | Pointer to [**OrderResponseDataRelationshipsAvailableCustomerPaymentSources**](OrderResponseDataRelationshipsAvailableCustomerPaymentSources.md) |  | [optional] 
**AvailableFreeSkus** | Pointer to [**OrderResponseDataRelationshipsAvailableFreeSkus**](OrderResponseDataRelationshipsAvailableFreeSkus.md) |  | [optional] 
**AvailableFreeBundles** | Pointer to [**OrderResponseDataRelationshipsAvailableFreeBundles**](OrderResponseDataRelationshipsAvailableFreeBundles.md) |  | [optional] 
**PaymentMethod** | Pointer to [**OrderResponseDataRelationshipsPaymentMethod**](OrderResponseDataRelationshipsPaymentMethod.md) |  | [optional] 
**PaymentSource** | Pointer to [**CustomerPaymentSourceResponseDataRelationshipsPaymentSource**](CustomerPaymentSourceResponseDataRelationshipsPaymentSource.md) |  | [optional] 
**LineItems** | Pointer to [**OrderResponseDataRelationshipsLineItems**](OrderResponseDataRelationshipsLineItems.md) |  | [optional] 
**Shipments** | Pointer to [**OrderResponseDataRelationshipsShipments**](OrderResponseDataRelationshipsShipments.md) |  | [optional] 
**Transactions** | Pointer to [**OrderResponseDataRelationshipsTransactions**](OrderResponseDataRelationshipsTransactions.md) |  | [optional] 
**Authorizations** | Pointer to [**OrderResponseDataRelationshipsAuthorizations**](OrderResponseDataRelationshipsAuthorizations.md) |  | [optional] 
**Captures** | Pointer to [**AuthorizationResponseDataRelationshipsCaptures**](AuthorizationResponseDataRelationshipsCaptures.md) |  | [optional] 
**Voids** | Pointer to [**AuthorizationResponseDataRelationshipsVoids**](AuthorizationResponseDataRelationshipsVoids.md) |  | [optional] 
**Refunds** | Pointer to [**CaptureResponseDataRelationshipsRefunds**](CaptureResponseDataRelationshipsRefunds.md) |  | [optional] 
**OrderSubscriptions** | Pointer to [**CustomerResponseDataRelationshipsOrderSubscriptions**](CustomerResponseDataRelationshipsOrderSubscriptions.md) |  | [optional] 
**OrderCopies** | Pointer to [**OrderSubscriptionResponseDataRelationshipsOrderCopies**](OrderSubscriptionResponseDataRelationshipsOrderCopies.md) |  | [optional] 
**Attachments** | Pointer to [**AvalaraAccountResponseDataRelationshipsAttachments**](AvalaraAccountResponseDataRelationshipsAttachments.md) |  | [optional] 
**Events** | Pointer to [**CustomerAddressResponseDataRelationshipsEvents**](CustomerAddressResponseDataRelationshipsEvents.md) |  | [optional] 

## Methods

### NewOrderResponseDataRelationships

`func NewOrderResponseDataRelationships() *OrderResponseDataRelationships`

NewOrderResponseDataRelationships instantiates a new OrderResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderResponseDataRelationshipsWithDefaults

`func NewOrderResponseDataRelationshipsWithDefaults() *OrderResponseDataRelationships`

NewOrderResponseDataRelationshipsWithDefaults instantiates a new OrderResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *OrderResponseDataRelationships) GetMarket() BillingInfoValidationRuleResponseDataRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *OrderResponseDataRelationships) GetMarketOk() (*BillingInfoValidationRuleResponseDataRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *OrderResponseDataRelationships) SetMarket(v BillingInfoValidationRuleResponseDataRelationshipsMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *OrderResponseDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetCustomer

`func (o *OrderResponseDataRelationships) GetCustomer() CouponRecipientResponseDataRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *OrderResponseDataRelationships) GetCustomerOk() (*CouponRecipientResponseDataRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *OrderResponseDataRelationships) SetCustomer(v CouponRecipientResponseDataRelationshipsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *OrderResponseDataRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetShippingAddress

`func (o *OrderResponseDataRelationships) GetShippingAddress() OrderResponseDataRelationshipsShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *OrderResponseDataRelationships) GetShippingAddressOk() (*OrderResponseDataRelationshipsShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *OrderResponseDataRelationships) SetShippingAddress(v OrderResponseDataRelationshipsShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *OrderResponseDataRelationships) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetBillingAddress

`func (o *OrderResponseDataRelationships) GetBillingAddress() OrderResponseDataRelationshipsBillingAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *OrderResponseDataRelationships) GetBillingAddressOk() (*OrderResponseDataRelationshipsBillingAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *OrderResponseDataRelationships) SetBillingAddress(v OrderResponseDataRelationshipsBillingAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *OrderResponseDataRelationships) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetAvailablePaymentMethods

`func (o *OrderResponseDataRelationships) GetAvailablePaymentMethods() OrderResponseDataRelationshipsAvailablePaymentMethods`

GetAvailablePaymentMethods returns the AvailablePaymentMethods field if non-nil, zero value otherwise.

### GetAvailablePaymentMethodsOk

`func (o *OrderResponseDataRelationships) GetAvailablePaymentMethodsOk() (*OrderResponseDataRelationshipsAvailablePaymentMethods, bool)`

GetAvailablePaymentMethodsOk returns a tuple with the AvailablePaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePaymentMethods

`func (o *OrderResponseDataRelationships) SetAvailablePaymentMethods(v OrderResponseDataRelationshipsAvailablePaymentMethods)`

SetAvailablePaymentMethods sets AvailablePaymentMethods field to given value.

### HasAvailablePaymentMethods

`func (o *OrderResponseDataRelationships) HasAvailablePaymentMethods() bool`

HasAvailablePaymentMethods returns a boolean if a field has been set.

### GetAvailableCustomerPaymentSources

`func (o *OrderResponseDataRelationships) GetAvailableCustomerPaymentSources() OrderResponseDataRelationshipsAvailableCustomerPaymentSources`

GetAvailableCustomerPaymentSources returns the AvailableCustomerPaymentSources field if non-nil, zero value otherwise.

### GetAvailableCustomerPaymentSourcesOk

`func (o *OrderResponseDataRelationships) GetAvailableCustomerPaymentSourcesOk() (*OrderResponseDataRelationshipsAvailableCustomerPaymentSources, bool)`

GetAvailableCustomerPaymentSourcesOk returns a tuple with the AvailableCustomerPaymentSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCustomerPaymentSources

`func (o *OrderResponseDataRelationships) SetAvailableCustomerPaymentSources(v OrderResponseDataRelationshipsAvailableCustomerPaymentSources)`

SetAvailableCustomerPaymentSources sets AvailableCustomerPaymentSources field to given value.

### HasAvailableCustomerPaymentSources

`func (o *OrderResponseDataRelationships) HasAvailableCustomerPaymentSources() bool`

HasAvailableCustomerPaymentSources returns a boolean if a field has been set.

### GetAvailableFreeSkus

`func (o *OrderResponseDataRelationships) GetAvailableFreeSkus() OrderResponseDataRelationshipsAvailableFreeSkus`

GetAvailableFreeSkus returns the AvailableFreeSkus field if non-nil, zero value otherwise.

### GetAvailableFreeSkusOk

`func (o *OrderResponseDataRelationships) GetAvailableFreeSkusOk() (*OrderResponseDataRelationshipsAvailableFreeSkus, bool)`

GetAvailableFreeSkusOk returns a tuple with the AvailableFreeSkus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFreeSkus

`func (o *OrderResponseDataRelationships) SetAvailableFreeSkus(v OrderResponseDataRelationshipsAvailableFreeSkus)`

SetAvailableFreeSkus sets AvailableFreeSkus field to given value.

### HasAvailableFreeSkus

`func (o *OrderResponseDataRelationships) HasAvailableFreeSkus() bool`

HasAvailableFreeSkus returns a boolean if a field has been set.

### GetAvailableFreeBundles

`func (o *OrderResponseDataRelationships) GetAvailableFreeBundles() OrderResponseDataRelationshipsAvailableFreeBundles`

GetAvailableFreeBundles returns the AvailableFreeBundles field if non-nil, zero value otherwise.

### GetAvailableFreeBundlesOk

`func (o *OrderResponseDataRelationships) GetAvailableFreeBundlesOk() (*OrderResponseDataRelationshipsAvailableFreeBundles, bool)`

GetAvailableFreeBundlesOk returns a tuple with the AvailableFreeBundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFreeBundles

`func (o *OrderResponseDataRelationships) SetAvailableFreeBundles(v OrderResponseDataRelationshipsAvailableFreeBundles)`

SetAvailableFreeBundles sets AvailableFreeBundles field to given value.

### HasAvailableFreeBundles

`func (o *OrderResponseDataRelationships) HasAvailableFreeBundles() bool`

HasAvailableFreeBundles returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *OrderResponseDataRelationships) GetPaymentMethod() OrderResponseDataRelationshipsPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *OrderResponseDataRelationships) GetPaymentMethodOk() (*OrderResponseDataRelationshipsPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *OrderResponseDataRelationships) SetPaymentMethod(v OrderResponseDataRelationshipsPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *OrderResponseDataRelationships) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPaymentSource

`func (o *OrderResponseDataRelationships) GetPaymentSource() CustomerPaymentSourceResponseDataRelationshipsPaymentSource`

GetPaymentSource returns the PaymentSource field if non-nil, zero value otherwise.

### GetPaymentSourceOk

`func (o *OrderResponseDataRelationships) GetPaymentSourceOk() (*CustomerPaymentSourceResponseDataRelationshipsPaymentSource, bool)`

GetPaymentSourceOk returns a tuple with the PaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSource

`func (o *OrderResponseDataRelationships) SetPaymentSource(v CustomerPaymentSourceResponseDataRelationshipsPaymentSource)`

SetPaymentSource sets PaymentSource field to given value.

### HasPaymentSource

`func (o *OrderResponseDataRelationships) HasPaymentSource() bool`

HasPaymentSource returns a boolean if a field has been set.

### GetLineItems

`func (o *OrderResponseDataRelationships) GetLineItems() OrderResponseDataRelationshipsLineItems`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *OrderResponseDataRelationships) GetLineItemsOk() (*OrderResponseDataRelationshipsLineItems, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *OrderResponseDataRelationships) SetLineItems(v OrderResponseDataRelationshipsLineItems)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *OrderResponseDataRelationships) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetShipments

`func (o *OrderResponseDataRelationships) GetShipments() OrderResponseDataRelationshipsShipments`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *OrderResponseDataRelationships) GetShipmentsOk() (*OrderResponseDataRelationshipsShipments, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *OrderResponseDataRelationships) SetShipments(v OrderResponseDataRelationshipsShipments)`

SetShipments sets Shipments field to given value.

### HasShipments

`func (o *OrderResponseDataRelationships) HasShipments() bool`

HasShipments returns a boolean if a field has been set.

### GetTransactions

`func (o *OrderResponseDataRelationships) GetTransactions() OrderResponseDataRelationshipsTransactions`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *OrderResponseDataRelationships) GetTransactionsOk() (*OrderResponseDataRelationshipsTransactions, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *OrderResponseDataRelationships) SetTransactions(v OrderResponseDataRelationshipsTransactions)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *OrderResponseDataRelationships) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetAuthorizations

`func (o *OrderResponseDataRelationships) GetAuthorizations() OrderResponseDataRelationshipsAuthorizations`

GetAuthorizations returns the Authorizations field if non-nil, zero value otherwise.

### GetAuthorizationsOk

`func (o *OrderResponseDataRelationships) GetAuthorizationsOk() (*OrderResponseDataRelationshipsAuthorizations, bool)`

GetAuthorizationsOk returns a tuple with the Authorizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizations

`func (o *OrderResponseDataRelationships) SetAuthorizations(v OrderResponseDataRelationshipsAuthorizations)`

SetAuthorizations sets Authorizations field to given value.

### HasAuthorizations

`func (o *OrderResponseDataRelationships) HasAuthorizations() bool`

HasAuthorizations returns a boolean if a field has been set.

### GetCaptures

`func (o *OrderResponseDataRelationships) GetCaptures() AuthorizationResponseDataRelationshipsCaptures`

GetCaptures returns the Captures field if non-nil, zero value otherwise.

### GetCapturesOk

`func (o *OrderResponseDataRelationships) GetCapturesOk() (*AuthorizationResponseDataRelationshipsCaptures, bool)`

GetCapturesOk returns a tuple with the Captures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptures

`func (o *OrderResponseDataRelationships) SetCaptures(v AuthorizationResponseDataRelationshipsCaptures)`

SetCaptures sets Captures field to given value.

### HasCaptures

`func (o *OrderResponseDataRelationships) HasCaptures() bool`

HasCaptures returns a boolean if a field has been set.

### GetVoids

`func (o *OrderResponseDataRelationships) GetVoids() AuthorizationResponseDataRelationshipsVoids`

GetVoids returns the Voids field if non-nil, zero value otherwise.

### GetVoidsOk

`func (o *OrderResponseDataRelationships) GetVoidsOk() (*AuthorizationResponseDataRelationshipsVoids, bool)`

GetVoidsOk returns a tuple with the Voids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoids

`func (o *OrderResponseDataRelationships) SetVoids(v AuthorizationResponseDataRelationshipsVoids)`

SetVoids sets Voids field to given value.

### HasVoids

`func (o *OrderResponseDataRelationships) HasVoids() bool`

HasVoids returns a boolean if a field has been set.

### GetRefunds

`func (o *OrderResponseDataRelationships) GetRefunds() CaptureResponseDataRelationshipsRefunds`

GetRefunds returns the Refunds field if non-nil, zero value otherwise.

### GetRefundsOk

`func (o *OrderResponseDataRelationships) GetRefundsOk() (*CaptureResponseDataRelationshipsRefunds, bool)`

GetRefundsOk returns a tuple with the Refunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunds

`func (o *OrderResponseDataRelationships) SetRefunds(v CaptureResponseDataRelationshipsRefunds)`

SetRefunds sets Refunds field to given value.

### HasRefunds

`func (o *OrderResponseDataRelationships) HasRefunds() bool`

HasRefunds returns a boolean if a field has been set.

### GetOrderSubscriptions

`func (o *OrderResponseDataRelationships) GetOrderSubscriptions() CustomerResponseDataRelationshipsOrderSubscriptions`

GetOrderSubscriptions returns the OrderSubscriptions field if non-nil, zero value otherwise.

### GetOrderSubscriptionsOk

`func (o *OrderResponseDataRelationships) GetOrderSubscriptionsOk() (*CustomerResponseDataRelationshipsOrderSubscriptions, bool)`

GetOrderSubscriptionsOk returns a tuple with the OrderSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSubscriptions

`func (o *OrderResponseDataRelationships) SetOrderSubscriptions(v CustomerResponseDataRelationshipsOrderSubscriptions)`

SetOrderSubscriptions sets OrderSubscriptions field to given value.

### HasOrderSubscriptions

`func (o *OrderResponseDataRelationships) HasOrderSubscriptions() bool`

HasOrderSubscriptions returns a boolean if a field has been set.

### GetOrderCopies

`func (o *OrderResponseDataRelationships) GetOrderCopies() OrderSubscriptionResponseDataRelationshipsOrderCopies`

GetOrderCopies returns the OrderCopies field if non-nil, zero value otherwise.

### GetOrderCopiesOk

`func (o *OrderResponseDataRelationships) GetOrderCopiesOk() (*OrderSubscriptionResponseDataRelationshipsOrderCopies, bool)`

GetOrderCopiesOk returns a tuple with the OrderCopies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCopies

`func (o *OrderResponseDataRelationships) SetOrderCopies(v OrderSubscriptionResponseDataRelationshipsOrderCopies)`

SetOrderCopies sets OrderCopies field to given value.

### HasOrderCopies

`func (o *OrderResponseDataRelationships) HasOrderCopies() bool`

HasOrderCopies returns a boolean if a field has been set.

### GetAttachments

`func (o *OrderResponseDataRelationships) GetAttachments() AvalaraAccountResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *OrderResponseDataRelationships) GetAttachmentsOk() (*AvalaraAccountResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *OrderResponseDataRelationships) SetAttachments(v AvalaraAccountResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *OrderResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *OrderResponseDataRelationships) GetEvents() CustomerAddressResponseDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *OrderResponseDataRelationships) GetEventsOk() (*CustomerAddressResponseDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *OrderResponseDataRelationships) SetEvents(v CustomerAddressResponseDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *OrderResponseDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


