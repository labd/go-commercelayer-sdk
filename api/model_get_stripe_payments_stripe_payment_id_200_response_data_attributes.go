/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.6.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GETStripePaymentsStripePaymentId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETStripePaymentsStripePaymentId200ResponseDataAttributes{}

// GETStripePaymentsStripePaymentId200ResponseDataAttributes struct for GETStripePaymentsStripePaymentId200ResponseDataAttributes
type GETStripePaymentsStripePaymentId200ResponseDataAttributes struct {
	// The Stripe payment intent ID. Required to identify a payment session on stripe.
	StripeId interface{} `json:"stripe_id,omitempty"`
	// The Stripe payment intent client secret. Required to create a charge through Stripe.js.
	ClientSecret interface{} `json:"client_secret,omitempty"`
	// The Stripe charge ID. Identifies money movement upon the payment intent confirmation.
	ChargeId interface{} `json:"charge_id,omitempty"`
	// The Stripe publishable API key.
	PublishableKey interface{} `json:"publishable_key,omitempty"`
	// Stripe payment options: 'customer', 'payment_method', 'return_url', etc. Check Stripe payment intent API for more details.
	Options interface{} `json:"options,omitempty"`
	// Stripe 'payment_method', set by webhook.
	PaymentMethod interface{} `json:"payment_method,omitempty"`
	// Indicates if the order current amount differs form the one of the created payment intent.
	MismatchedAmounts interface{} `json:"mismatched_amounts,omitempty"`
	// The URL to return to when a redirect payment is completed.
	ReturnUrl interface{} `json:"return_url,omitempty"`
	// The email address to send the receipt to.
	ReceiptEmail interface{} `json:"receipt_email,omitempty"`
	// Information about the payment instrument used in the transaction.
	PaymentInstrument interface{} `json:"payment_instrument,omitempty"`
	// Time at which the resource was created.
	CreatedAt interface{} `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt interface{} `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewGETStripePaymentsStripePaymentId200ResponseDataAttributes instantiates a new GETStripePaymentsStripePaymentId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETStripePaymentsStripePaymentId200ResponseDataAttributes() *GETStripePaymentsStripePaymentId200ResponseDataAttributes {
	this := GETStripePaymentsStripePaymentId200ResponseDataAttributes{}
	return &this
}

// NewGETStripePaymentsStripePaymentId200ResponseDataAttributesWithDefaults instantiates a new GETStripePaymentsStripePaymentId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETStripePaymentsStripePaymentId200ResponseDataAttributesWithDefaults() *GETStripePaymentsStripePaymentId200ResponseDataAttributes {
	this := GETStripePaymentsStripePaymentId200ResponseDataAttributes{}
	return &this
}

// GetStripeId returns the StripeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetStripeId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.StripeId
}

// GetStripeIdOk returns a tuple with the StripeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetStripeIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.StripeId) {
		return nil, false
	}
	return &o.StripeId, true
}

// HasStripeId returns a boolean if a field has been set.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasStripeId() bool {
	if o != nil && IsNil(o.StripeId) {
		return true
	}

	return false
}

// SetStripeId gets a reference to the given interface{} and assigns it to the StripeId field.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetStripeId(v interface{}) {
	o.StripeId = v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetClientSecret() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetClientSecretOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return &o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasClientSecret() bool {
	if o != nil && IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given interface{} and assigns it to the ClientSecret field.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetClientSecret(v interface{}) {
	o.ClientSecret = v
}

// GetChargeId returns the ChargeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetChargeId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ChargeId
}

// GetChargeIdOk returns a tuple with the ChargeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetChargeIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ChargeId) {
		return nil, false
	}
	return &o.ChargeId, true
}

// HasChargeId returns a boolean if a field has been set.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasChargeId() bool {
	if o != nil && IsNil(o.ChargeId) {
		return true
	}

	return false
}

// SetChargeId gets a reference to the given interface{} and assigns it to the ChargeId field.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetChargeId(v interface{}) {
	o.ChargeId = v
}

// GetPublishableKey returns the PublishableKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetPublishableKey() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PublishableKey
}

// GetPublishableKeyOk returns a tuple with the PublishableKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetPublishableKeyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PublishableKey) {
		return nil, false
	}
	return &o.PublishableKey, true
}

// HasPublishableKey returns a boolean if a field has been set.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasPublishableKey() bool {
	if o != nil && IsNil(o.PublishableKey) {
		return true
	}

	return false
}

// SetPublishableKey gets a reference to the given interface{} and assigns it to the PublishableKey field.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetPublishableKey(v interface{}) {
	o.PublishableKey = v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetOptions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetOptionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return &o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasOptions() bool {
	if o != nil && IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given interface{} and assigns it to the Options field.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetOptions(v interface{}) {
	o.Options = v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetPaymentMethod() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetPaymentMethodOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return &o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasPaymentMethod() bool {
	if o != nil && IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given interface{} and assigns it to the PaymentMethod field.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetPaymentMethod(v interface{}) {
	o.PaymentMethod = v
}

// GetMismatchedAmounts returns the MismatchedAmounts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetMismatchedAmounts() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MismatchedAmounts
}

// GetMismatchedAmountsOk returns a tuple with the MismatchedAmounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetMismatchedAmountsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MismatchedAmounts) {
		return nil, false
	}
	return &o.MismatchedAmounts, true
}

// HasMismatchedAmounts returns a boolean if a field has been set.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasMismatchedAmounts() bool {
	if o != nil && IsNil(o.MismatchedAmounts) {
		return true
	}

	return false
}

// SetMismatchedAmounts gets a reference to the given interface{} and assigns it to the MismatchedAmounts field.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetMismatchedAmounts(v interface{}) {
	o.MismatchedAmounts = v
}

// GetReturnUrl returns the ReturnUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetReturnUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReturnUrl
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetReturnUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReturnUrl) {
		return nil, false
	}
	return &o.ReturnUrl, true
}

// HasReturnUrl returns a boolean if a field has been set.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasReturnUrl() bool {
	if o != nil && IsNil(o.ReturnUrl) {
		return true
	}

	return false
}

// SetReturnUrl gets a reference to the given interface{} and assigns it to the ReturnUrl field.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetReturnUrl(v interface{}) {
	o.ReturnUrl = v
}

// GetReceiptEmail returns the ReceiptEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetReceiptEmail() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReceiptEmail
}

// GetReceiptEmailOk returns a tuple with the ReceiptEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetReceiptEmailOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReceiptEmail) {
		return nil, false
	}
	return &o.ReceiptEmail, true
}

// HasReceiptEmail returns a boolean if a field has been set.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasReceiptEmail() bool {
	if o != nil && IsNil(o.ReceiptEmail) {
		return true
	}

	return false
}

// SetReceiptEmail gets a reference to the given interface{} and assigns it to the ReceiptEmail field.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetReceiptEmail(v interface{}) {
	o.ReceiptEmail = v
}

// GetPaymentInstrument returns the PaymentInstrument field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetPaymentInstrument() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PaymentInstrument
}

// GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetPaymentInstrumentOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PaymentInstrument) {
		return nil, false
	}
	return &o.PaymentInstrument, true
}

// HasPaymentInstrument returns a boolean if a field has been set.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasPaymentInstrument() bool {
	if o != nil && IsNil(o.PaymentInstrument) {
		return true
	}

	return false
}

// SetPaymentInstrument gets a reference to the given interface{} and assigns it to the PaymentInstrument field.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetPaymentInstrument(v interface{}) {
	o.PaymentInstrument = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETStripePaymentsStripePaymentId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETStripePaymentsStripePaymentId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.StripeId != nil {
		toSerialize["stripe_id"] = o.StripeId
	}
	if o.ClientSecret != nil {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if o.ChargeId != nil {
		toSerialize["charge_id"] = o.ChargeId
	}
	if o.PublishableKey != nil {
		toSerialize["publishable_key"] = o.PublishableKey
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.PaymentMethod != nil {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	if o.MismatchedAmounts != nil {
		toSerialize["mismatched_amounts"] = o.MismatchedAmounts
	}
	if o.ReturnUrl != nil {
		toSerialize["return_url"] = o.ReturnUrl
	}
	if o.ReceiptEmail != nil {
		toSerialize["receipt_email"] = o.ReceiptEmail
	}
	if o.PaymentInstrument != nil {
		toSerialize["payment_instrument"] = o.PaymentInstrument
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
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

type NullableGETStripePaymentsStripePaymentId200ResponseDataAttributes struct {
	value *GETStripePaymentsStripePaymentId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETStripePaymentsStripePaymentId200ResponseDataAttributes) Get() *GETStripePaymentsStripePaymentId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETStripePaymentsStripePaymentId200ResponseDataAttributes) Set(val *GETStripePaymentsStripePaymentId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETStripePaymentsStripePaymentId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETStripePaymentsStripePaymentId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETStripePaymentsStripePaymentId200ResponseDataAttributes(val *GETStripePaymentsStripePaymentId200ResponseDataAttributes) *NullableGETStripePaymentsStripePaymentId200ResponseDataAttributes {
	return &NullableGETStripePaymentsStripePaymentId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETStripePaymentsStripePaymentId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETStripePaymentsStripePaymentId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
