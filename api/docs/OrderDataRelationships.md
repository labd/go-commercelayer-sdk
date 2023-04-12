# OrderDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**AvalaraAccountDataRelationshipsMarkets**](AvalaraAccountDataRelationshipsMarkets.md) |  | [optional] 
**Customer** | Pointer to [**CouponRecipientDataRelationshipsCustomer**](CouponRecipientDataRelationshipsCustomer.md) |  | [optional] 
**ShippingAddress** | Pointer to [**BingGeocoderDataRelationshipsAddresses**](BingGeocoderDataRelationshipsAddresses.md) |  | [optional] 
**BillingAddress** | Pointer to [**BingGeocoderDataRelationshipsAddresses**](BingGeocoderDataRelationshipsAddresses.md) |  | [optional] 
**AvailablePaymentMethods** | Pointer to [**AdyenGatewayDataRelationshipsPaymentMethods**](AdyenGatewayDataRelationshipsPaymentMethods.md) |  | [optional] 
**AvailableCustomerPaymentSources** | Pointer to [**CustomerDataRelationshipsCustomerPaymentSources**](CustomerDataRelationshipsCustomerPaymentSources.md) |  | [optional] 
**AvailableFreeSkus** | Pointer to [**BundleDataRelationshipsSkus**](BundleDataRelationshipsSkus.md) |  | [optional] 
**AvailableFreeBundles** | Pointer to [**OrderDataRelationshipsAvailableFreeBundles**](OrderDataRelationshipsAvailableFreeBundles.md) |  | [optional] 
**PaymentMethod** | Pointer to [**AdyenGatewayDataRelationshipsPaymentMethods**](AdyenGatewayDataRelationshipsPaymentMethods.md) |  | [optional] 
**PaymentSource** | Pointer to [**CustomerPaymentSourceDataRelationshipsPaymentSource**](CustomerPaymentSourceDataRelationshipsPaymentSource.md) |  | [optional] 
**LineItems** | Pointer to [**LineItemOptionDataRelationshipsLineItem**](LineItemOptionDataRelationshipsLineItem.md) |  | [optional] 
**Shipments** | Pointer to [**OrderDataRelationshipsShipments**](OrderDataRelationshipsShipments.md) |  | [optional] 
**Transactions** | Pointer to [**OrderDataRelationshipsTransactions**](OrderDataRelationshipsTransactions.md) |  | [optional] 
**Authorizations** | Pointer to [**CaptureDataRelationshipsReferenceAuthorization**](CaptureDataRelationshipsReferenceAuthorization.md) |  | [optional] 
**Captures** | Pointer to [**AuthorizationDataRelationshipsCaptures**](AuthorizationDataRelationshipsCaptures.md) |  | [optional] 
**Voids** | Pointer to [**AuthorizationDataRelationshipsVoids**](AuthorizationDataRelationshipsVoids.md) |  | [optional] 
**Refunds** | Pointer to [**CaptureDataRelationshipsRefunds**](CaptureDataRelationshipsRefunds.md) |  | [optional] 
**Returns** | Pointer to [**CustomerDataRelationshipsReturns**](CustomerDataRelationshipsReturns.md) |  | [optional] 
**OrderSubscriptions** | Pointer to [**CustomerDataRelationshipsOrderSubscriptions**](CustomerDataRelationshipsOrderSubscriptions.md) |  | [optional] 
**OrderFactories** | Pointer to [**OrderSubscriptionDataRelationshipsOrderFactories**](OrderSubscriptionDataRelationshipsOrderFactories.md) |  | [optional] 
**OrderCopies** | Pointer to [**OrderDataRelationshipsOrderCopies**](OrderDataRelationshipsOrderCopies.md) |  | [optional] 
**RecurringOrderCopies** | Pointer to [**OrderSubscriptionDataRelationshipsRecurringOrderCopies**](OrderSubscriptionDataRelationshipsRecurringOrderCopies.md) |  | [optional] 
**Attachments** | Pointer to [**AvalaraAccountDataRelationshipsAttachments**](AvalaraAccountDataRelationshipsAttachments.md) |  | [optional] 
**Events** | Pointer to [**AuthorizationDataRelationshipsEvents**](AuthorizationDataRelationshipsEvents.md) |  | [optional] 

## Methods

### NewOrderDataRelationships

`func NewOrderDataRelationships() *OrderDataRelationships`

NewOrderDataRelationships instantiates a new OrderDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDataRelationshipsWithDefaults

`func NewOrderDataRelationshipsWithDefaults() *OrderDataRelationships`

NewOrderDataRelationshipsWithDefaults instantiates a new OrderDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *OrderDataRelationships) GetMarket() AvalaraAccountDataRelationshipsMarkets`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *OrderDataRelationships) GetMarketOk() (*AvalaraAccountDataRelationshipsMarkets, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *OrderDataRelationships) SetMarket(v AvalaraAccountDataRelationshipsMarkets)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *OrderDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetCustomer

`func (o *OrderDataRelationships) GetCustomer() CouponRecipientDataRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *OrderDataRelationships) GetCustomerOk() (*CouponRecipientDataRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *OrderDataRelationships) SetCustomer(v CouponRecipientDataRelationshipsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *OrderDataRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetShippingAddress

`func (o *OrderDataRelationships) GetShippingAddress() BingGeocoderDataRelationshipsAddresses`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *OrderDataRelationships) GetShippingAddressOk() (*BingGeocoderDataRelationshipsAddresses, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *OrderDataRelationships) SetShippingAddress(v BingGeocoderDataRelationshipsAddresses)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *OrderDataRelationships) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetBillingAddress

`func (o *OrderDataRelationships) GetBillingAddress() BingGeocoderDataRelationshipsAddresses`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *OrderDataRelationships) GetBillingAddressOk() (*BingGeocoderDataRelationshipsAddresses, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *OrderDataRelationships) SetBillingAddress(v BingGeocoderDataRelationshipsAddresses)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *OrderDataRelationships) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetAvailablePaymentMethods

`func (o *OrderDataRelationships) GetAvailablePaymentMethods() AdyenGatewayDataRelationshipsPaymentMethods`

GetAvailablePaymentMethods returns the AvailablePaymentMethods field if non-nil, zero value otherwise.

### GetAvailablePaymentMethodsOk

`func (o *OrderDataRelationships) GetAvailablePaymentMethodsOk() (*AdyenGatewayDataRelationshipsPaymentMethods, bool)`

GetAvailablePaymentMethodsOk returns a tuple with the AvailablePaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePaymentMethods

`func (o *OrderDataRelationships) SetAvailablePaymentMethods(v AdyenGatewayDataRelationshipsPaymentMethods)`

SetAvailablePaymentMethods sets AvailablePaymentMethods field to given value.

### HasAvailablePaymentMethods

`func (o *OrderDataRelationships) HasAvailablePaymentMethods() bool`

HasAvailablePaymentMethods returns a boolean if a field has been set.

### GetAvailableCustomerPaymentSources

`func (o *OrderDataRelationships) GetAvailableCustomerPaymentSources() CustomerDataRelationshipsCustomerPaymentSources`

GetAvailableCustomerPaymentSources returns the AvailableCustomerPaymentSources field if non-nil, zero value otherwise.

### GetAvailableCustomerPaymentSourcesOk

`func (o *OrderDataRelationships) GetAvailableCustomerPaymentSourcesOk() (*CustomerDataRelationshipsCustomerPaymentSources, bool)`

GetAvailableCustomerPaymentSourcesOk returns a tuple with the AvailableCustomerPaymentSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCustomerPaymentSources

`func (o *OrderDataRelationships) SetAvailableCustomerPaymentSources(v CustomerDataRelationshipsCustomerPaymentSources)`

SetAvailableCustomerPaymentSources sets AvailableCustomerPaymentSources field to given value.

### HasAvailableCustomerPaymentSources

`func (o *OrderDataRelationships) HasAvailableCustomerPaymentSources() bool`

HasAvailableCustomerPaymentSources returns a boolean if a field has been set.

### GetAvailableFreeSkus

`func (o *OrderDataRelationships) GetAvailableFreeSkus() BundleDataRelationshipsSkus`

GetAvailableFreeSkus returns the AvailableFreeSkus field if non-nil, zero value otherwise.

### GetAvailableFreeSkusOk

`func (o *OrderDataRelationships) GetAvailableFreeSkusOk() (*BundleDataRelationshipsSkus, bool)`

GetAvailableFreeSkusOk returns a tuple with the AvailableFreeSkus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFreeSkus

`func (o *OrderDataRelationships) SetAvailableFreeSkus(v BundleDataRelationshipsSkus)`

SetAvailableFreeSkus sets AvailableFreeSkus field to given value.

### HasAvailableFreeSkus

`func (o *OrderDataRelationships) HasAvailableFreeSkus() bool`

HasAvailableFreeSkus returns a boolean if a field has been set.

### GetAvailableFreeBundles

`func (o *OrderDataRelationships) GetAvailableFreeBundles() OrderDataRelationshipsAvailableFreeBundles`

GetAvailableFreeBundles returns the AvailableFreeBundles field if non-nil, zero value otherwise.

### GetAvailableFreeBundlesOk

`func (o *OrderDataRelationships) GetAvailableFreeBundlesOk() (*OrderDataRelationshipsAvailableFreeBundles, bool)`

GetAvailableFreeBundlesOk returns a tuple with the AvailableFreeBundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFreeBundles

`func (o *OrderDataRelationships) SetAvailableFreeBundles(v OrderDataRelationshipsAvailableFreeBundles)`

SetAvailableFreeBundles sets AvailableFreeBundles field to given value.

### HasAvailableFreeBundles

`func (o *OrderDataRelationships) HasAvailableFreeBundles() bool`

HasAvailableFreeBundles returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *OrderDataRelationships) GetPaymentMethod() AdyenGatewayDataRelationshipsPaymentMethods`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *OrderDataRelationships) GetPaymentMethodOk() (*AdyenGatewayDataRelationshipsPaymentMethods, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *OrderDataRelationships) SetPaymentMethod(v AdyenGatewayDataRelationshipsPaymentMethods)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *OrderDataRelationships) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPaymentSource

`func (o *OrderDataRelationships) GetPaymentSource() CustomerPaymentSourceDataRelationshipsPaymentSource`

GetPaymentSource returns the PaymentSource field if non-nil, zero value otherwise.

### GetPaymentSourceOk

`func (o *OrderDataRelationships) GetPaymentSourceOk() (*CustomerPaymentSourceDataRelationshipsPaymentSource, bool)`

GetPaymentSourceOk returns a tuple with the PaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSource

`func (o *OrderDataRelationships) SetPaymentSource(v CustomerPaymentSourceDataRelationshipsPaymentSource)`

SetPaymentSource sets PaymentSource field to given value.

### HasPaymentSource

`func (o *OrderDataRelationships) HasPaymentSource() bool`

HasPaymentSource returns a boolean if a field has been set.

### GetLineItems

`func (o *OrderDataRelationships) GetLineItems() LineItemOptionDataRelationshipsLineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *OrderDataRelationships) GetLineItemsOk() (*LineItemOptionDataRelationshipsLineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *OrderDataRelationships) SetLineItems(v LineItemOptionDataRelationshipsLineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *OrderDataRelationships) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetShipments

`func (o *OrderDataRelationships) GetShipments() OrderDataRelationshipsShipments`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *OrderDataRelationships) GetShipmentsOk() (*OrderDataRelationshipsShipments, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *OrderDataRelationships) SetShipments(v OrderDataRelationshipsShipments)`

SetShipments sets Shipments field to given value.

### HasShipments

`func (o *OrderDataRelationships) HasShipments() bool`

HasShipments returns a boolean if a field has been set.

### GetTransactions

`func (o *OrderDataRelationships) GetTransactions() OrderDataRelationshipsTransactions`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *OrderDataRelationships) GetTransactionsOk() (*OrderDataRelationshipsTransactions, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *OrderDataRelationships) SetTransactions(v OrderDataRelationshipsTransactions)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *OrderDataRelationships) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetAuthorizations

`func (o *OrderDataRelationships) GetAuthorizations() CaptureDataRelationshipsReferenceAuthorization`

GetAuthorizations returns the Authorizations field if non-nil, zero value otherwise.

### GetAuthorizationsOk

`func (o *OrderDataRelationships) GetAuthorizationsOk() (*CaptureDataRelationshipsReferenceAuthorization, bool)`

GetAuthorizationsOk returns a tuple with the Authorizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizations

`func (o *OrderDataRelationships) SetAuthorizations(v CaptureDataRelationshipsReferenceAuthorization)`

SetAuthorizations sets Authorizations field to given value.

### HasAuthorizations

`func (o *OrderDataRelationships) HasAuthorizations() bool`

HasAuthorizations returns a boolean if a field has been set.

### GetCaptures

`func (o *OrderDataRelationships) GetCaptures() AuthorizationDataRelationshipsCaptures`

GetCaptures returns the Captures field if non-nil, zero value otherwise.

### GetCapturesOk

`func (o *OrderDataRelationships) GetCapturesOk() (*AuthorizationDataRelationshipsCaptures, bool)`

GetCapturesOk returns a tuple with the Captures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptures

`func (o *OrderDataRelationships) SetCaptures(v AuthorizationDataRelationshipsCaptures)`

SetCaptures sets Captures field to given value.

### HasCaptures

`func (o *OrderDataRelationships) HasCaptures() bool`

HasCaptures returns a boolean if a field has been set.

### GetVoids

`func (o *OrderDataRelationships) GetVoids() AuthorizationDataRelationshipsVoids`

GetVoids returns the Voids field if non-nil, zero value otherwise.

### GetVoidsOk

`func (o *OrderDataRelationships) GetVoidsOk() (*AuthorizationDataRelationshipsVoids, bool)`

GetVoidsOk returns a tuple with the Voids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoids

`func (o *OrderDataRelationships) SetVoids(v AuthorizationDataRelationshipsVoids)`

SetVoids sets Voids field to given value.

### HasVoids

`func (o *OrderDataRelationships) HasVoids() bool`

HasVoids returns a boolean if a field has been set.

### GetRefunds

`func (o *OrderDataRelationships) GetRefunds() CaptureDataRelationshipsRefunds`

GetRefunds returns the Refunds field if non-nil, zero value otherwise.

### GetRefundsOk

`func (o *OrderDataRelationships) GetRefundsOk() (*CaptureDataRelationshipsRefunds, bool)`

GetRefundsOk returns a tuple with the Refunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunds

`func (o *OrderDataRelationships) SetRefunds(v CaptureDataRelationshipsRefunds)`

SetRefunds sets Refunds field to given value.

### HasRefunds

`func (o *OrderDataRelationships) HasRefunds() bool`

HasRefunds returns a boolean if a field has been set.

### GetReturns

`func (o *OrderDataRelationships) GetReturns() CustomerDataRelationshipsReturns`

GetReturns returns the Returns field if non-nil, zero value otherwise.

### GetReturnsOk

`func (o *OrderDataRelationships) GetReturnsOk() (*CustomerDataRelationshipsReturns, bool)`

GetReturnsOk returns a tuple with the Returns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturns

`func (o *OrderDataRelationships) SetReturns(v CustomerDataRelationshipsReturns)`

SetReturns sets Returns field to given value.

### HasReturns

`func (o *OrderDataRelationships) HasReturns() bool`

HasReturns returns a boolean if a field has been set.

### GetOrderSubscriptions

`func (o *OrderDataRelationships) GetOrderSubscriptions() CustomerDataRelationshipsOrderSubscriptions`

GetOrderSubscriptions returns the OrderSubscriptions field if non-nil, zero value otherwise.

### GetOrderSubscriptionsOk

`func (o *OrderDataRelationships) GetOrderSubscriptionsOk() (*CustomerDataRelationshipsOrderSubscriptions, bool)`

GetOrderSubscriptionsOk returns a tuple with the OrderSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSubscriptions

`func (o *OrderDataRelationships) SetOrderSubscriptions(v CustomerDataRelationshipsOrderSubscriptions)`

SetOrderSubscriptions sets OrderSubscriptions field to given value.

### HasOrderSubscriptions

`func (o *OrderDataRelationships) HasOrderSubscriptions() bool`

HasOrderSubscriptions returns a boolean if a field has been set.

### GetOrderFactories

`func (o *OrderDataRelationships) GetOrderFactories() OrderSubscriptionDataRelationshipsOrderFactories`

GetOrderFactories returns the OrderFactories field if non-nil, zero value otherwise.

### GetOrderFactoriesOk

`func (o *OrderDataRelationships) GetOrderFactoriesOk() (*OrderSubscriptionDataRelationshipsOrderFactories, bool)`

GetOrderFactoriesOk returns a tuple with the OrderFactories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderFactories

`func (o *OrderDataRelationships) SetOrderFactories(v OrderSubscriptionDataRelationshipsOrderFactories)`

SetOrderFactories sets OrderFactories field to given value.

### HasOrderFactories

`func (o *OrderDataRelationships) HasOrderFactories() bool`

HasOrderFactories returns a boolean if a field has been set.

### GetOrderCopies

`func (o *OrderDataRelationships) GetOrderCopies() OrderDataRelationshipsOrderCopies`

GetOrderCopies returns the OrderCopies field if non-nil, zero value otherwise.

### GetOrderCopiesOk

`func (o *OrderDataRelationships) GetOrderCopiesOk() (*OrderDataRelationshipsOrderCopies, bool)`

GetOrderCopiesOk returns a tuple with the OrderCopies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCopies

`func (o *OrderDataRelationships) SetOrderCopies(v OrderDataRelationshipsOrderCopies)`

SetOrderCopies sets OrderCopies field to given value.

### HasOrderCopies

`func (o *OrderDataRelationships) HasOrderCopies() bool`

HasOrderCopies returns a boolean if a field has been set.

### GetRecurringOrderCopies

`func (o *OrderDataRelationships) GetRecurringOrderCopies() OrderSubscriptionDataRelationshipsRecurringOrderCopies`

GetRecurringOrderCopies returns the RecurringOrderCopies field if non-nil, zero value otherwise.

### GetRecurringOrderCopiesOk

`func (o *OrderDataRelationships) GetRecurringOrderCopiesOk() (*OrderSubscriptionDataRelationshipsRecurringOrderCopies, bool)`

GetRecurringOrderCopiesOk returns a tuple with the RecurringOrderCopies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringOrderCopies

`func (o *OrderDataRelationships) SetRecurringOrderCopies(v OrderSubscriptionDataRelationshipsRecurringOrderCopies)`

SetRecurringOrderCopies sets RecurringOrderCopies field to given value.

### HasRecurringOrderCopies

`func (o *OrderDataRelationships) HasRecurringOrderCopies() bool`

HasRecurringOrderCopies returns a boolean if a field has been set.

### GetAttachments

`func (o *OrderDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *OrderDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *OrderDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *OrderDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *OrderDataRelationships) GetEvents() AuthorizationDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *OrderDataRelationships) GetEventsOk() (*AuthorizationDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *OrderDataRelationships) SetEvents(v AuthorizationDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *OrderDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


