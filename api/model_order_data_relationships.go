/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.2
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// OrderDataRelationships struct for OrderDataRelationships
type OrderDataRelationships struct {
	Market                          *AvalaraAccountDataRelationshipsMarkets              `json:"market,omitempty"`
	Customer                        *CouponRecipientDataRelationshipsCustomer            `json:"customer,omitempty"`
	ShippingAddress                 *BingGeocoderDataRelationshipsAddresses              `json:"shipping_address,omitempty"`
	BillingAddress                  *BingGeocoderDataRelationshipsAddresses              `json:"billing_address,omitempty"`
	AvailablePaymentMethods         *AdyenGatewayDataRelationshipsPaymentMethods         `json:"available_payment_methods,omitempty"`
	AvailableCustomerPaymentSources *CustomerDataRelationshipsCustomerPaymentSources     `json:"available_customer_payment_sources,omitempty"`
	AvailableFreeSkus               *BundleDataRelationshipsSkus                         `json:"available_free_skus,omitempty"`
	AvailableFreeBundles            *OrderDataRelationshipsAvailableFreeBundles          `json:"available_free_bundles,omitempty"`
	PaymentMethod                   *AdyenGatewayDataRelationshipsPaymentMethods         `json:"payment_method,omitempty"`
	PaymentSource                   *CustomerPaymentSourceDataRelationshipsPaymentSource `json:"payment_source,omitempty"`
	LineItems                       *LineItemOptionDataRelationshipsLineItem             `json:"line_items,omitempty"`
	Shipments                       *OrderDataRelationshipsShipments                     `json:"shipments,omitempty"`
	Transactions                    *OrderDataRelationshipsTransactions                  `json:"transactions,omitempty"`
	Authorizations                  *CaptureDataRelationshipsReferenceAuthorization      `json:"authorizations,omitempty"`
	Captures                        *AuthorizationDataRelationshipsCaptures              `json:"captures,omitempty"`
	Voids                           *AuthorizationDataRelationshipsVoids                 `json:"voids,omitempty"`
	Refunds                         *CaptureDataRelationshipsRefunds                     `json:"refunds,omitempty"`
	OrderSubscriptions              *CustomerDataRelationshipsOrderSubscriptions         `json:"order_subscriptions,omitempty"`
	OrderCopies                     *OrderSubscriptionDataRelationshipsOrderCopies       `json:"order_copies,omitempty"`
	Attachments                     *AvalaraAccountDataRelationshipsAttachments          `json:"attachments,omitempty"`
	Events                          *CustomerAddressDataRelationshipsEvents              `json:"events,omitempty"`
}

// NewOrderDataRelationships instantiates a new OrderDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDataRelationships() *OrderDataRelationships {
	this := OrderDataRelationships{}
	return &this
}

// NewOrderDataRelationshipsWithDefaults instantiates a new OrderDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDataRelationshipsWithDefaults() *OrderDataRelationships {
	this := OrderDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *OrderDataRelationships) GetMarket() AvalaraAccountDataRelationshipsMarkets {
	if o == nil || o.Market == nil {
		var ret AvalaraAccountDataRelationshipsMarkets
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataRelationships) GetMarketOk() (*AvalaraAccountDataRelationshipsMarkets, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *OrderDataRelationships) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given AvalaraAccountDataRelationshipsMarkets and assigns it to the Market field.
func (o *OrderDataRelationships) SetMarket(v AvalaraAccountDataRelationshipsMarkets) {
	o.Market = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *OrderDataRelationships) GetCustomer() CouponRecipientDataRelationshipsCustomer {
	if o == nil || o.Customer == nil {
		var ret CouponRecipientDataRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataRelationships) GetCustomerOk() (*CouponRecipientDataRelationshipsCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *OrderDataRelationships) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given CouponRecipientDataRelationshipsCustomer and assigns it to the Customer field.
func (o *OrderDataRelationships) SetCustomer(v CouponRecipientDataRelationshipsCustomer) {
	o.Customer = &v
}

// GetShippingAddress returns the ShippingAddress field value if set, zero value otherwise.
func (o *OrderDataRelationships) GetShippingAddress() BingGeocoderDataRelationshipsAddresses {
	if o == nil || o.ShippingAddress == nil {
		var ret BingGeocoderDataRelationshipsAddresses
		return ret
	}
	return *o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataRelationships) GetShippingAddressOk() (*BingGeocoderDataRelationshipsAddresses, bool) {
	if o == nil || o.ShippingAddress == nil {
		return nil, false
	}
	return o.ShippingAddress, true
}

// HasShippingAddress returns a boolean if a field has been set.
func (o *OrderDataRelationships) HasShippingAddress() bool {
	if o != nil && o.ShippingAddress != nil {
		return true
	}

	return false
}

// SetShippingAddress gets a reference to the given BingGeocoderDataRelationshipsAddresses and assigns it to the ShippingAddress field.
func (o *OrderDataRelationships) SetShippingAddress(v BingGeocoderDataRelationshipsAddresses) {
	o.ShippingAddress = &v
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *OrderDataRelationships) GetBillingAddress() BingGeocoderDataRelationshipsAddresses {
	if o == nil || o.BillingAddress == nil {
		var ret BingGeocoderDataRelationshipsAddresses
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataRelationships) GetBillingAddressOk() (*BingGeocoderDataRelationshipsAddresses, bool) {
	if o == nil || o.BillingAddress == nil {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *OrderDataRelationships) HasBillingAddress() bool {
	if o != nil && o.BillingAddress != nil {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given BingGeocoderDataRelationshipsAddresses and assigns it to the BillingAddress field.
func (o *OrderDataRelationships) SetBillingAddress(v BingGeocoderDataRelationshipsAddresses) {
	o.BillingAddress = &v
}

// GetAvailablePaymentMethods returns the AvailablePaymentMethods field value if set, zero value otherwise.
func (o *OrderDataRelationships) GetAvailablePaymentMethods() AdyenGatewayDataRelationshipsPaymentMethods {
	if o == nil || o.AvailablePaymentMethods == nil {
		var ret AdyenGatewayDataRelationshipsPaymentMethods
		return ret
	}
	return *o.AvailablePaymentMethods
}

// GetAvailablePaymentMethodsOk returns a tuple with the AvailablePaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataRelationships) GetAvailablePaymentMethodsOk() (*AdyenGatewayDataRelationshipsPaymentMethods, bool) {
	if o == nil || o.AvailablePaymentMethods == nil {
		return nil, false
	}
	return o.AvailablePaymentMethods, true
}

// HasAvailablePaymentMethods returns a boolean if a field has been set.
func (o *OrderDataRelationships) HasAvailablePaymentMethods() bool {
	if o != nil && o.AvailablePaymentMethods != nil {
		return true
	}

	return false
}

// SetAvailablePaymentMethods gets a reference to the given AdyenGatewayDataRelationshipsPaymentMethods and assigns it to the AvailablePaymentMethods field.
func (o *OrderDataRelationships) SetAvailablePaymentMethods(v AdyenGatewayDataRelationshipsPaymentMethods) {
	o.AvailablePaymentMethods = &v
}

// GetAvailableCustomerPaymentSources returns the AvailableCustomerPaymentSources field value if set, zero value otherwise.
func (o *OrderDataRelationships) GetAvailableCustomerPaymentSources() CustomerDataRelationshipsCustomerPaymentSources {
	if o == nil || o.AvailableCustomerPaymentSources == nil {
		var ret CustomerDataRelationshipsCustomerPaymentSources
		return ret
	}
	return *o.AvailableCustomerPaymentSources
}

// GetAvailableCustomerPaymentSourcesOk returns a tuple with the AvailableCustomerPaymentSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataRelationships) GetAvailableCustomerPaymentSourcesOk() (*CustomerDataRelationshipsCustomerPaymentSources, bool) {
	if o == nil || o.AvailableCustomerPaymentSources == nil {
		return nil, false
	}
	return o.AvailableCustomerPaymentSources, true
}

// HasAvailableCustomerPaymentSources returns a boolean if a field has been set.
func (o *OrderDataRelationships) HasAvailableCustomerPaymentSources() bool {
	if o != nil && o.AvailableCustomerPaymentSources != nil {
		return true
	}

	return false
}

// SetAvailableCustomerPaymentSources gets a reference to the given CustomerDataRelationshipsCustomerPaymentSources and assigns it to the AvailableCustomerPaymentSources field.
func (o *OrderDataRelationships) SetAvailableCustomerPaymentSources(v CustomerDataRelationshipsCustomerPaymentSources) {
	o.AvailableCustomerPaymentSources = &v
}

// GetAvailableFreeSkus returns the AvailableFreeSkus field value if set, zero value otherwise.
func (o *OrderDataRelationships) GetAvailableFreeSkus() BundleDataRelationshipsSkus {
	if o == nil || o.AvailableFreeSkus == nil {
		var ret BundleDataRelationshipsSkus
		return ret
	}
	return *o.AvailableFreeSkus
}

// GetAvailableFreeSkusOk returns a tuple with the AvailableFreeSkus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataRelationships) GetAvailableFreeSkusOk() (*BundleDataRelationshipsSkus, bool) {
	if o == nil || o.AvailableFreeSkus == nil {
		return nil, false
	}
	return o.AvailableFreeSkus, true
}

// HasAvailableFreeSkus returns a boolean if a field has been set.
func (o *OrderDataRelationships) HasAvailableFreeSkus() bool {
	if o != nil && o.AvailableFreeSkus != nil {
		return true
	}

	return false
}

// SetAvailableFreeSkus gets a reference to the given BundleDataRelationshipsSkus and assigns it to the AvailableFreeSkus field.
func (o *OrderDataRelationships) SetAvailableFreeSkus(v BundleDataRelationshipsSkus) {
	o.AvailableFreeSkus = &v
}

// GetAvailableFreeBundles returns the AvailableFreeBundles field value if set, zero value otherwise.
func (o *OrderDataRelationships) GetAvailableFreeBundles() OrderDataRelationshipsAvailableFreeBundles {
	if o == nil || o.AvailableFreeBundles == nil {
		var ret OrderDataRelationshipsAvailableFreeBundles
		return ret
	}
	return *o.AvailableFreeBundles
}

// GetAvailableFreeBundlesOk returns a tuple with the AvailableFreeBundles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataRelationships) GetAvailableFreeBundlesOk() (*OrderDataRelationshipsAvailableFreeBundles, bool) {
	if o == nil || o.AvailableFreeBundles == nil {
		return nil, false
	}
	return o.AvailableFreeBundles, true
}

// HasAvailableFreeBundles returns a boolean if a field has been set.
func (o *OrderDataRelationships) HasAvailableFreeBundles() bool {
	if o != nil && o.AvailableFreeBundles != nil {
		return true
	}

	return false
}

// SetAvailableFreeBundles gets a reference to the given OrderDataRelationshipsAvailableFreeBundles and assigns it to the AvailableFreeBundles field.
func (o *OrderDataRelationships) SetAvailableFreeBundles(v OrderDataRelationshipsAvailableFreeBundles) {
	o.AvailableFreeBundles = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *OrderDataRelationships) GetPaymentMethod() AdyenGatewayDataRelationshipsPaymentMethods {
	if o == nil || o.PaymentMethod == nil {
		var ret AdyenGatewayDataRelationshipsPaymentMethods
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataRelationships) GetPaymentMethodOk() (*AdyenGatewayDataRelationshipsPaymentMethods, bool) {
	if o == nil || o.PaymentMethod == nil {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *OrderDataRelationships) HasPaymentMethod() bool {
	if o != nil && o.PaymentMethod != nil {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given AdyenGatewayDataRelationshipsPaymentMethods and assigns it to the PaymentMethod field.
func (o *OrderDataRelationships) SetPaymentMethod(v AdyenGatewayDataRelationshipsPaymentMethods) {
	o.PaymentMethod = &v
}

// GetPaymentSource returns the PaymentSource field value if set, zero value otherwise.
func (o *OrderDataRelationships) GetPaymentSource() CustomerPaymentSourceDataRelationshipsPaymentSource {
	if o == nil || o.PaymentSource == nil {
		var ret CustomerPaymentSourceDataRelationshipsPaymentSource
		return ret
	}
	return *o.PaymentSource
}

// GetPaymentSourceOk returns a tuple with the PaymentSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataRelationships) GetPaymentSourceOk() (*CustomerPaymentSourceDataRelationshipsPaymentSource, bool) {
	if o == nil || o.PaymentSource == nil {
		return nil, false
	}
	return o.PaymentSource, true
}

// HasPaymentSource returns a boolean if a field has been set.
func (o *OrderDataRelationships) HasPaymentSource() bool {
	if o != nil && o.PaymentSource != nil {
		return true
	}

	return false
}

// SetPaymentSource gets a reference to the given CustomerPaymentSourceDataRelationshipsPaymentSource and assigns it to the PaymentSource field.
func (o *OrderDataRelationships) SetPaymentSource(v CustomerPaymentSourceDataRelationshipsPaymentSource) {
	o.PaymentSource = &v
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *OrderDataRelationships) GetLineItems() LineItemOptionDataRelationshipsLineItem {
	if o == nil || o.LineItems == nil {
		var ret LineItemOptionDataRelationshipsLineItem
		return ret
	}
	return *o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataRelationships) GetLineItemsOk() (*LineItemOptionDataRelationshipsLineItem, bool) {
	if o == nil || o.LineItems == nil {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItems returns a boolean if a field has been set.
func (o *OrderDataRelationships) HasLineItems() bool {
	if o != nil && o.LineItems != nil {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given LineItemOptionDataRelationshipsLineItem and assigns it to the LineItems field.
func (o *OrderDataRelationships) SetLineItems(v LineItemOptionDataRelationshipsLineItem) {
	o.LineItems = &v
}

// GetShipments returns the Shipments field value if set, zero value otherwise.
func (o *OrderDataRelationships) GetShipments() OrderDataRelationshipsShipments {
	if o == nil || o.Shipments == nil {
		var ret OrderDataRelationshipsShipments
		return ret
	}
	return *o.Shipments
}

// GetShipmentsOk returns a tuple with the Shipments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataRelationships) GetShipmentsOk() (*OrderDataRelationshipsShipments, bool) {
	if o == nil || o.Shipments == nil {
		return nil, false
	}
	return o.Shipments, true
}

// HasShipments returns a boolean if a field has been set.
func (o *OrderDataRelationships) HasShipments() bool {
	if o != nil && o.Shipments != nil {
		return true
	}

	return false
}

// SetShipments gets a reference to the given OrderDataRelationshipsShipments and assigns it to the Shipments field.
func (o *OrderDataRelationships) SetShipments(v OrderDataRelationshipsShipments) {
	o.Shipments = &v
}

// GetTransactions returns the Transactions field value if set, zero value otherwise.
func (o *OrderDataRelationships) GetTransactions() OrderDataRelationshipsTransactions {
	if o == nil || o.Transactions == nil {
		var ret OrderDataRelationshipsTransactions
		return ret
	}
	return *o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataRelationships) GetTransactionsOk() (*OrderDataRelationshipsTransactions, bool) {
	if o == nil || o.Transactions == nil {
		return nil, false
	}
	return o.Transactions, true
}

// HasTransactions returns a boolean if a field has been set.
func (o *OrderDataRelationships) HasTransactions() bool {
	if o != nil && o.Transactions != nil {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given OrderDataRelationshipsTransactions and assigns it to the Transactions field.
func (o *OrderDataRelationships) SetTransactions(v OrderDataRelationshipsTransactions) {
	o.Transactions = &v
}

// GetAuthorizations returns the Authorizations field value if set, zero value otherwise.
func (o *OrderDataRelationships) GetAuthorizations() CaptureDataRelationshipsReferenceAuthorization {
	if o == nil || o.Authorizations == nil {
		var ret CaptureDataRelationshipsReferenceAuthorization
		return ret
	}
	return *o.Authorizations
}

// GetAuthorizationsOk returns a tuple with the Authorizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataRelationships) GetAuthorizationsOk() (*CaptureDataRelationshipsReferenceAuthorization, bool) {
	if o == nil || o.Authorizations == nil {
		return nil, false
	}
	return o.Authorizations, true
}

// HasAuthorizations returns a boolean if a field has been set.
func (o *OrderDataRelationships) HasAuthorizations() bool {
	if o != nil && o.Authorizations != nil {
		return true
	}

	return false
}

// SetAuthorizations gets a reference to the given CaptureDataRelationshipsReferenceAuthorization and assigns it to the Authorizations field.
func (o *OrderDataRelationships) SetAuthorizations(v CaptureDataRelationshipsReferenceAuthorization) {
	o.Authorizations = &v
}

// GetCaptures returns the Captures field value if set, zero value otherwise.
func (o *OrderDataRelationships) GetCaptures() AuthorizationDataRelationshipsCaptures {
	if o == nil || o.Captures == nil {
		var ret AuthorizationDataRelationshipsCaptures
		return ret
	}
	return *o.Captures
}

// GetCapturesOk returns a tuple with the Captures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataRelationships) GetCapturesOk() (*AuthorizationDataRelationshipsCaptures, bool) {
	if o == nil || o.Captures == nil {
		return nil, false
	}
	return o.Captures, true
}

// HasCaptures returns a boolean if a field has been set.
func (o *OrderDataRelationships) HasCaptures() bool {
	if o != nil && o.Captures != nil {
		return true
	}

	return false
}

// SetCaptures gets a reference to the given AuthorizationDataRelationshipsCaptures and assigns it to the Captures field.
func (o *OrderDataRelationships) SetCaptures(v AuthorizationDataRelationshipsCaptures) {
	o.Captures = &v
}

// GetVoids returns the Voids field value if set, zero value otherwise.
func (o *OrderDataRelationships) GetVoids() AuthorizationDataRelationshipsVoids {
	if o == nil || o.Voids == nil {
		var ret AuthorizationDataRelationshipsVoids
		return ret
	}
	return *o.Voids
}

// GetVoidsOk returns a tuple with the Voids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataRelationships) GetVoidsOk() (*AuthorizationDataRelationshipsVoids, bool) {
	if o == nil || o.Voids == nil {
		return nil, false
	}
	return o.Voids, true
}

// HasVoids returns a boolean if a field has been set.
func (o *OrderDataRelationships) HasVoids() bool {
	if o != nil && o.Voids != nil {
		return true
	}

	return false
}

// SetVoids gets a reference to the given AuthorizationDataRelationshipsVoids and assigns it to the Voids field.
func (o *OrderDataRelationships) SetVoids(v AuthorizationDataRelationshipsVoids) {
	o.Voids = &v
}

// GetRefunds returns the Refunds field value if set, zero value otherwise.
func (o *OrderDataRelationships) GetRefunds() CaptureDataRelationshipsRefunds {
	if o == nil || o.Refunds == nil {
		var ret CaptureDataRelationshipsRefunds
		return ret
	}
	return *o.Refunds
}

// GetRefundsOk returns a tuple with the Refunds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataRelationships) GetRefundsOk() (*CaptureDataRelationshipsRefunds, bool) {
	if o == nil || o.Refunds == nil {
		return nil, false
	}
	return o.Refunds, true
}

// HasRefunds returns a boolean if a field has been set.
func (o *OrderDataRelationships) HasRefunds() bool {
	if o != nil && o.Refunds != nil {
		return true
	}

	return false
}

// SetRefunds gets a reference to the given CaptureDataRelationshipsRefunds and assigns it to the Refunds field.
func (o *OrderDataRelationships) SetRefunds(v CaptureDataRelationshipsRefunds) {
	o.Refunds = &v
}

// GetOrderSubscriptions returns the OrderSubscriptions field value if set, zero value otherwise.
func (o *OrderDataRelationships) GetOrderSubscriptions() CustomerDataRelationshipsOrderSubscriptions {
	if o == nil || o.OrderSubscriptions == nil {
		var ret CustomerDataRelationshipsOrderSubscriptions
		return ret
	}
	return *o.OrderSubscriptions
}

// GetOrderSubscriptionsOk returns a tuple with the OrderSubscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataRelationships) GetOrderSubscriptionsOk() (*CustomerDataRelationshipsOrderSubscriptions, bool) {
	if o == nil || o.OrderSubscriptions == nil {
		return nil, false
	}
	return o.OrderSubscriptions, true
}

// HasOrderSubscriptions returns a boolean if a field has been set.
func (o *OrderDataRelationships) HasOrderSubscriptions() bool {
	if o != nil && o.OrderSubscriptions != nil {
		return true
	}

	return false
}

// SetOrderSubscriptions gets a reference to the given CustomerDataRelationshipsOrderSubscriptions and assigns it to the OrderSubscriptions field.
func (o *OrderDataRelationships) SetOrderSubscriptions(v CustomerDataRelationshipsOrderSubscriptions) {
	o.OrderSubscriptions = &v
}

// GetOrderCopies returns the OrderCopies field value if set, zero value otherwise.
func (o *OrderDataRelationships) GetOrderCopies() OrderSubscriptionDataRelationshipsOrderCopies {
	if o == nil || o.OrderCopies == nil {
		var ret OrderSubscriptionDataRelationshipsOrderCopies
		return ret
	}
	return *o.OrderCopies
}

// GetOrderCopiesOk returns a tuple with the OrderCopies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataRelationships) GetOrderCopiesOk() (*OrderSubscriptionDataRelationshipsOrderCopies, bool) {
	if o == nil || o.OrderCopies == nil {
		return nil, false
	}
	return o.OrderCopies, true
}

// HasOrderCopies returns a boolean if a field has been set.
func (o *OrderDataRelationships) HasOrderCopies() bool {
	if o != nil && o.OrderCopies != nil {
		return true
	}

	return false
}

// SetOrderCopies gets a reference to the given OrderSubscriptionDataRelationshipsOrderCopies and assigns it to the OrderCopies field.
func (o *OrderDataRelationships) SetOrderCopies(v OrderSubscriptionDataRelationshipsOrderCopies) {
	o.OrderCopies = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *OrderDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret AvalaraAccountDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *OrderDataRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *OrderDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *OrderDataRelationships) GetEvents() CustomerAddressDataRelationshipsEvents {
	if o == nil || o.Events == nil {
		var ret CustomerAddressDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataRelationships) GetEventsOk() (*CustomerAddressDataRelationshipsEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *OrderDataRelationships) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given CustomerAddressDataRelationshipsEvents and assigns it to the Events field.
func (o *OrderDataRelationships) SetEvents(v CustomerAddressDataRelationshipsEvents) {
	o.Events = &v
}

func (o OrderDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Market != nil {
		toSerialize["market"] = o.Market
	}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.ShippingAddress != nil {
		toSerialize["shipping_address"] = o.ShippingAddress
	}
	if o.BillingAddress != nil {
		toSerialize["billing_address"] = o.BillingAddress
	}
	if o.AvailablePaymentMethods != nil {
		toSerialize["available_payment_methods"] = o.AvailablePaymentMethods
	}
	if o.AvailableCustomerPaymentSources != nil {
		toSerialize["available_customer_payment_sources"] = o.AvailableCustomerPaymentSources
	}
	if o.AvailableFreeSkus != nil {
		toSerialize["available_free_skus"] = o.AvailableFreeSkus
	}
	if o.AvailableFreeBundles != nil {
		toSerialize["available_free_bundles"] = o.AvailableFreeBundles
	}
	if o.PaymentMethod != nil {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	if o.PaymentSource != nil {
		toSerialize["payment_source"] = o.PaymentSource
	}
	if o.LineItems != nil {
		toSerialize["line_items"] = o.LineItems
	}
	if o.Shipments != nil {
		toSerialize["shipments"] = o.Shipments
	}
	if o.Transactions != nil {
		toSerialize["transactions"] = o.Transactions
	}
	if o.Authorizations != nil {
		toSerialize["authorizations"] = o.Authorizations
	}
	if o.Captures != nil {
		toSerialize["captures"] = o.Captures
	}
	if o.Voids != nil {
		toSerialize["voids"] = o.Voids
	}
	if o.Refunds != nil {
		toSerialize["refunds"] = o.Refunds
	}
	if o.OrderSubscriptions != nil {
		toSerialize["order_subscriptions"] = o.OrderSubscriptions
	}
	if o.OrderCopies != nil {
		toSerialize["order_copies"] = o.OrderCopies
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableOrderDataRelationships struct {
	value *OrderDataRelationships
	isSet bool
}

func (v NullableOrderDataRelationships) Get() *OrderDataRelationships {
	return v.value
}

func (v *NullableOrderDataRelationships) Set(val *OrderDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDataRelationships(val *OrderDataRelationships) *NullableOrderDataRelationships {
	return &NullableOrderDataRelationships{value: val, isSet: true}
}

func (v NullableOrderDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
