/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.5.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTOrders201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrders201ResponseDataAttributes{}

// POSTOrders201ResponseDataAttributes struct for POSTOrders201ResponseDataAttributes
type POSTOrders201ResponseDataAttributes struct {
	// The order identifier. Can be specified if unique within the organization (for enterprise plans only), default to numeric ID otherwise. Cannot be passed by sales channels.
	Number interface{} `json:"number,omitempty"`
	// The affiliate code, if any, to track commissions using any third party services.
	AffiliateCode interface{} `json:"affiliate_code,omitempty"`
	// Save this attribute as 'false' if you want prevent the order to be refreshed automatically at each change (much faster).
	Autorefresh interface{} `json:"autorefresh,omitempty"`
	// Save this attribute as 'true' if you want perform the place asynchronously. Payment errors, if any, will be collected afterwards.
	PlaceAsync interface{} `json:"place_async,omitempty"`
	// Indicates if the order has been placed as guest.
	Guest interface{} `json:"guest,omitempty"`
	// The email address of the associated customer. When creating or updating an order, this is a shortcut to find or create the associated customer by email.
	CustomerEmail interface{} `json:"customer_email,omitempty"`
	// The password of the associated customer. When creating or updating an order, this is a shortcut to sign up the associated customer.
	CustomerPassword interface{} `json:"customer_password,omitempty"`
	// The preferred language code (ISO 639-1) to be used when communicating with the customer. This can be useful when sending the order to 3rd party marketing tools and CRMs. If the language is supported, the hosted checkout will be localized accordingly.
	LanguageCode interface{} `json:"language_code,omitempty"`
	// Indicates if taxes are applied to shipping costs.
	FreightTaxable interface{} `json:"freight_taxable,omitempty"`
	// Indicates if taxes are applied to payment methods costs.
	PaymentMethodTaxable interface{} `json:"payment_method_taxable,omitempty"`
	// Indicates if taxes are applied to positive adjustments.
	AdjustmentTaxable interface{} `json:"adjustment_taxable,omitempty"`
	// Indicates if taxes are applied to purchased gift cards.
	GiftCardTaxable interface{} `json:"gift_card_taxable,omitempty"`
	// The country code that you want the shipping address to be locked to. This can be useful to make sure the shipping address belongs to a given shipping country, e.g. the one selected in a country selector page. Not relevant if order contains only digital products.
	ShippingCountryCodeLock interface{} `json:"shipping_country_code_lock,omitempty"`
	// The coupon code to be used for the order. If valid, it triggers a promotion adding a discount line item to the order.
	CouponCode interface{} `json:"coupon_code,omitempty"`
	// The gift card code (at least the first 8 characters) to be used for the order. If valid, it uses the gift card balance to pay for the order.
	GiftCardCode interface{} `json:"gift_card_code,omitempty"`
	// The cart url on your site. If present, it will be used on our hosted checkout application.
	CartUrl interface{} `json:"cart_url,omitempty"`
	// The return url on your site. If present, it will be used on our hosted checkout application.
	ReturnUrl interface{} `json:"return_url,omitempty"`
	// The terms and conditions url on your site. If present, it will be used on our hosted checkout application.
	TermsUrl interface{} `json:"terms_url,omitempty"`
	// The privacy policy url on your site. If present, it will be used on our hosted checkout application.
	PrivacyUrl interface{} `json:"privacy_url,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTOrders201ResponseDataAttributes instantiates a new POSTOrders201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrders201ResponseDataAttributes() *POSTOrders201ResponseDataAttributes {
	this := POSTOrders201ResponseDataAttributes{}
	return &this
}

// NewPOSTOrders201ResponseDataAttributesWithDefaults instantiates a new POSTOrders201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrders201ResponseDataAttributesWithDefaults() *POSTOrders201ResponseDataAttributes {
	this := POSTOrders201ResponseDataAttributes{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataAttributes) GetNumber() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataAttributes) GetNumberOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return &o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataAttributes) HasNumber() bool {
	if o != nil && IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given interface{} and assigns it to the Number field.
func (o *POSTOrders201ResponseDataAttributes) SetNumber(v interface{}) {
	o.Number = v
}

// GetAffiliateCode returns the AffiliateCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataAttributes) GetAffiliateCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AffiliateCode
}

// GetAffiliateCodeOk returns a tuple with the AffiliateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataAttributes) GetAffiliateCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AffiliateCode) {
		return nil, false
	}
	return &o.AffiliateCode, true
}

// HasAffiliateCode returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataAttributes) HasAffiliateCode() bool {
	if o != nil && IsNil(o.AffiliateCode) {
		return true
	}

	return false
}

// SetAffiliateCode gets a reference to the given interface{} and assigns it to the AffiliateCode field.
func (o *POSTOrders201ResponseDataAttributes) SetAffiliateCode(v interface{}) {
	o.AffiliateCode = v
}

// GetAutorefresh returns the Autorefresh field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataAttributes) GetAutorefresh() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Autorefresh
}

// GetAutorefreshOk returns a tuple with the Autorefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataAttributes) GetAutorefreshOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Autorefresh) {
		return nil, false
	}
	return &o.Autorefresh, true
}

// HasAutorefresh returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataAttributes) HasAutorefresh() bool {
	if o != nil && IsNil(o.Autorefresh) {
		return true
	}

	return false
}

// SetAutorefresh gets a reference to the given interface{} and assigns it to the Autorefresh field.
func (o *POSTOrders201ResponseDataAttributes) SetAutorefresh(v interface{}) {
	o.Autorefresh = v
}

// GetPlaceAsync returns the PlaceAsync field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataAttributes) GetPlaceAsync() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PlaceAsync
}

// GetPlaceAsyncOk returns a tuple with the PlaceAsync field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataAttributes) GetPlaceAsyncOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PlaceAsync) {
		return nil, false
	}
	return &o.PlaceAsync, true
}

// HasPlaceAsync returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataAttributes) HasPlaceAsync() bool {
	if o != nil && IsNil(o.PlaceAsync) {
		return true
	}

	return false
}

// SetPlaceAsync gets a reference to the given interface{} and assigns it to the PlaceAsync field.
func (o *POSTOrders201ResponseDataAttributes) SetPlaceAsync(v interface{}) {
	o.PlaceAsync = v
}

// GetGuest returns the Guest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataAttributes) GetGuest() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Guest
}

// GetGuestOk returns a tuple with the Guest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataAttributes) GetGuestOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Guest) {
		return nil, false
	}
	return &o.Guest, true
}

// HasGuest returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataAttributes) HasGuest() bool {
	if o != nil && IsNil(o.Guest) {
		return true
	}

	return false
}

// SetGuest gets a reference to the given interface{} and assigns it to the Guest field.
func (o *POSTOrders201ResponseDataAttributes) SetGuest(v interface{}) {
	o.Guest = v
}

// GetCustomerEmail returns the CustomerEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataAttributes) GetCustomerEmail() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CustomerEmail
}

// GetCustomerEmailOk returns a tuple with the CustomerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataAttributes) GetCustomerEmailOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CustomerEmail) {
		return nil, false
	}
	return &o.CustomerEmail, true
}

// HasCustomerEmail returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataAttributes) HasCustomerEmail() bool {
	if o != nil && IsNil(o.CustomerEmail) {
		return true
	}

	return false
}

// SetCustomerEmail gets a reference to the given interface{} and assigns it to the CustomerEmail field.
func (o *POSTOrders201ResponseDataAttributes) SetCustomerEmail(v interface{}) {
	o.CustomerEmail = v
}

// GetCustomerPassword returns the CustomerPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataAttributes) GetCustomerPassword() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CustomerPassword
}

// GetCustomerPasswordOk returns a tuple with the CustomerPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataAttributes) GetCustomerPasswordOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CustomerPassword) {
		return nil, false
	}
	return &o.CustomerPassword, true
}

// HasCustomerPassword returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataAttributes) HasCustomerPassword() bool {
	if o != nil && IsNil(o.CustomerPassword) {
		return true
	}

	return false
}

// SetCustomerPassword gets a reference to the given interface{} and assigns it to the CustomerPassword field.
func (o *POSTOrders201ResponseDataAttributes) SetCustomerPassword(v interface{}) {
	o.CustomerPassword = v
}

// GetLanguageCode returns the LanguageCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataAttributes) GetLanguageCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.LanguageCode
}

// GetLanguageCodeOk returns a tuple with the LanguageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataAttributes) GetLanguageCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.LanguageCode) {
		return nil, false
	}
	return &o.LanguageCode, true
}

// HasLanguageCode returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataAttributes) HasLanguageCode() bool {
	if o != nil && IsNil(o.LanguageCode) {
		return true
	}

	return false
}

// SetLanguageCode gets a reference to the given interface{} and assigns it to the LanguageCode field.
func (o *POSTOrders201ResponseDataAttributes) SetLanguageCode(v interface{}) {
	o.LanguageCode = v
}

// GetFreightTaxable returns the FreightTaxable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataAttributes) GetFreightTaxable() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FreightTaxable
}

// GetFreightTaxableOk returns a tuple with the FreightTaxable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataAttributes) GetFreightTaxableOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FreightTaxable) {
		return nil, false
	}
	return &o.FreightTaxable, true
}

// HasFreightTaxable returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataAttributes) HasFreightTaxable() bool {
	if o != nil && IsNil(o.FreightTaxable) {
		return true
	}

	return false
}

// SetFreightTaxable gets a reference to the given interface{} and assigns it to the FreightTaxable field.
func (o *POSTOrders201ResponseDataAttributes) SetFreightTaxable(v interface{}) {
	o.FreightTaxable = v
}

// GetPaymentMethodTaxable returns the PaymentMethodTaxable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataAttributes) GetPaymentMethodTaxable() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PaymentMethodTaxable
}

// GetPaymentMethodTaxableOk returns a tuple with the PaymentMethodTaxable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataAttributes) GetPaymentMethodTaxableOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PaymentMethodTaxable) {
		return nil, false
	}
	return &o.PaymentMethodTaxable, true
}

// HasPaymentMethodTaxable returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataAttributes) HasPaymentMethodTaxable() bool {
	if o != nil && IsNil(o.PaymentMethodTaxable) {
		return true
	}

	return false
}

// SetPaymentMethodTaxable gets a reference to the given interface{} and assigns it to the PaymentMethodTaxable field.
func (o *POSTOrders201ResponseDataAttributes) SetPaymentMethodTaxable(v interface{}) {
	o.PaymentMethodTaxable = v
}

// GetAdjustmentTaxable returns the AdjustmentTaxable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataAttributes) GetAdjustmentTaxable() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AdjustmentTaxable
}

// GetAdjustmentTaxableOk returns a tuple with the AdjustmentTaxable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataAttributes) GetAdjustmentTaxableOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AdjustmentTaxable) {
		return nil, false
	}
	return &o.AdjustmentTaxable, true
}

// HasAdjustmentTaxable returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataAttributes) HasAdjustmentTaxable() bool {
	if o != nil && IsNil(o.AdjustmentTaxable) {
		return true
	}

	return false
}

// SetAdjustmentTaxable gets a reference to the given interface{} and assigns it to the AdjustmentTaxable field.
func (o *POSTOrders201ResponseDataAttributes) SetAdjustmentTaxable(v interface{}) {
	o.AdjustmentTaxable = v
}

// GetGiftCardTaxable returns the GiftCardTaxable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataAttributes) GetGiftCardTaxable() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.GiftCardTaxable
}

// GetGiftCardTaxableOk returns a tuple with the GiftCardTaxable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataAttributes) GetGiftCardTaxableOk() (*interface{}, bool) {
	if o == nil || IsNil(o.GiftCardTaxable) {
		return nil, false
	}
	return &o.GiftCardTaxable, true
}

// HasGiftCardTaxable returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataAttributes) HasGiftCardTaxable() bool {
	if o != nil && IsNil(o.GiftCardTaxable) {
		return true
	}

	return false
}

// SetGiftCardTaxable gets a reference to the given interface{} and assigns it to the GiftCardTaxable field.
func (o *POSTOrders201ResponseDataAttributes) SetGiftCardTaxable(v interface{}) {
	o.GiftCardTaxable = v
}

// GetShippingCountryCodeLock returns the ShippingCountryCodeLock field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataAttributes) GetShippingCountryCodeLock() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ShippingCountryCodeLock
}

// GetShippingCountryCodeLockOk returns a tuple with the ShippingCountryCodeLock field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataAttributes) GetShippingCountryCodeLockOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ShippingCountryCodeLock) {
		return nil, false
	}
	return &o.ShippingCountryCodeLock, true
}

// HasShippingCountryCodeLock returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataAttributes) HasShippingCountryCodeLock() bool {
	if o != nil && IsNil(o.ShippingCountryCodeLock) {
		return true
	}

	return false
}

// SetShippingCountryCodeLock gets a reference to the given interface{} and assigns it to the ShippingCountryCodeLock field.
func (o *POSTOrders201ResponseDataAttributes) SetShippingCountryCodeLock(v interface{}) {
	o.ShippingCountryCodeLock = v
}

// GetCouponCode returns the CouponCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataAttributes) GetCouponCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CouponCode
}

// GetCouponCodeOk returns a tuple with the CouponCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataAttributes) GetCouponCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CouponCode) {
		return nil, false
	}
	return &o.CouponCode, true
}

// HasCouponCode returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataAttributes) HasCouponCode() bool {
	if o != nil && IsNil(o.CouponCode) {
		return true
	}

	return false
}

// SetCouponCode gets a reference to the given interface{} and assigns it to the CouponCode field.
func (o *POSTOrders201ResponseDataAttributes) SetCouponCode(v interface{}) {
	o.CouponCode = v
}

// GetGiftCardCode returns the GiftCardCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataAttributes) GetGiftCardCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.GiftCardCode
}

// GetGiftCardCodeOk returns a tuple with the GiftCardCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataAttributes) GetGiftCardCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.GiftCardCode) {
		return nil, false
	}
	return &o.GiftCardCode, true
}

// HasGiftCardCode returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataAttributes) HasGiftCardCode() bool {
	if o != nil && IsNil(o.GiftCardCode) {
		return true
	}

	return false
}

// SetGiftCardCode gets a reference to the given interface{} and assigns it to the GiftCardCode field.
func (o *POSTOrders201ResponseDataAttributes) SetGiftCardCode(v interface{}) {
	o.GiftCardCode = v
}

// GetCartUrl returns the CartUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataAttributes) GetCartUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CartUrl
}

// GetCartUrlOk returns a tuple with the CartUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataAttributes) GetCartUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CartUrl) {
		return nil, false
	}
	return &o.CartUrl, true
}

// HasCartUrl returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataAttributes) HasCartUrl() bool {
	if o != nil && IsNil(o.CartUrl) {
		return true
	}

	return false
}

// SetCartUrl gets a reference to the given interface{} and assigns it to the CartUrl field.
func (o *POSTOrders201ResponseDataAttributes) SetCartUrl(v interface{}) {
	o.CartUrl = v
}

// GetReturnUrl returns the ReturnUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataAttributes) GetReturnUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReturnUrl
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataAttributes) GetReturnUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReturnUrl) {
		return nil, false
	}
	return &o.ReturnUrl, true
}

// HasReturnUrl returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataAttributes) HasReturnUrl() bool {
	if o != nil && IsNil(o.ReturnUrl) {
		return true
	}

	return false
}

// SetReturnUrl gets a reference to the given interface{} and assigns it to the ReturnUrl field.
func (o *POSTOrders201ResponseDataAttributes) SetReturnUrl(v interface{}) {
	o.ReturnUrl = v
}

// GetTermsUrl returns the TermsUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataAttributes) GetTermsUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TermsUrl
}

// GetTermsUrlOk returns a tuple with the TermsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataAttributes) GetTermsUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TermsUrl) {
		return nil, false
	}
	return &o.TermsUrl, true
}

// HasTermsUrl returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataAttributes) HasTermsUrl() bool {
	if o != nil && IsNil(o.TermsUrl) {
		return true
	}

	return false
}

// SetTermsUrl gets a reference to the given interface{} and assigns it to the TermsUrl field.
func (o *POSTOrders201ResponseDataAttributes) SetTermsUrl(v interface{}) {
	o.TermsUrl = v
}

// GetPrivacyUrl returns the PrivacyUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataAttributes) GetPrivacyUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PrivacyUrl
}

// GetPrivacyUrlOk returns a tuple with the PrivacyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataAttributes) GetPrivacyUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PrivacyUrl) {
		return nil, false
	}
	return &o.PrivacyUrl, true
}

// HasPrivacyUrl returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataAttributes) HasPrivacyUrl() bool {
	if o != nil && IsNil(o.PrivacyUrl) {
		return true
	}

	return false
}

// SetPrivacyUrl gets a reference to the given interface{} and assigns it to the PrivacyUrl field.
func (o *POSTOrders201ResponseDataAttributes) SetPrivacyUrl(v interface{}) {
	o.PrivacyUrl = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTOrders201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTOrders201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTOrders201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTOrders201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrders201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.AffiliateCode != nil {
		toSerialize["affiliate_code"] = o.AffiliateCode
	}
	if o.Autorefresh != nil {
		toSerialize["autorefresh"] = o.Autorefresh
	}
	if o.PlaceAsync != nil {
		toSerialize["place_async"] = o.PlaceAsync
	}
	if o.Guest != nil {
		toSerialize["guest"] = o.Guest
	}
	if o.CustomerEmail != nil {
		toSerialize["customer_email"] = o.CustomerEmail
	}
	if o.CustomerPassword != nil {
		toSerialize["customer_password"] = o.CustomerPassword
	}
	if o.LanguageCode != nil {
		toSerialize["language_code"] = o.LanguageCode
	}
	if o.FreightTaxable != nil {
		toSerialize["freight_taxable"] = o.FreightTaxable
	}
	if o.PaymentMethodTaxable != nil {
		toSerialize["payment_method_taxable"] = o.PaymentMethodTaxable
	}
	if o.AdjustmentTaxable != nil {
		toSerialize["adjustment_taxable"] = o.AdjustmentTaxable
	}
	if o.GiftCardTaxable != nil {
		toSerialize["gift_card_taxable"] = o.GiftCardTaxable
	}
	if o.ShippingCountryCodeLock != nil {
		toSerialize["shipping_country_code_lock"] = o.ShippingCountryCodeLock
	}
	if o.CouponCode != nil {
		toSerialize["coupon_code"] = o.CouponCode
	}
	if o.GiftCardCode != nil {
		toSerialize["gift_card_code"] = o.GiftCardCode
	}
	if o.CartUrl != nil {
		toSerialize["cart_url"] = o.CartUrl
	}
	if o.ReturnUrl != nil {
		toSerialize["return_url"] = o.ReturnUrl
	}
	if o.TermsUrl != nil {
		toSerialize["terms_url"] = o.TermsUrl
	}
	if o.PrivacyUrl != nil {
		toSerialize["privacy_url"] = o.PrivacyUrl
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullablePOSTOrders201ResponseDataAttributes struct {
	value *POSTOrders201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTOrders201ResponseDataAttributes) Get() *POSTOrders201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTOrders201ResponseDataAttributes) Set(val *POSTOrders201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrders201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrders201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrders201ResponseDataAttributes(val *POSTOrders201ResponseDataAttributes) *NullablePOSTOrders201ResponseDataAttributes {
	return &NullablePOSTOrders201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTOrders201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrders201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
