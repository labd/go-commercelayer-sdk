/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CheckoutComGatewayDataRelationships struct for CheckoutComGatewayDataRelationships
type CheckoutComGatewayDataRelationships struct {
	PaymentMethods      *AdyenGatewayDataRelationshipsPaymentMethods            `json:"payment_methods,omitempty"`
	CheckoutComPayments *CheckoutComGatewayDataRelationshipsCheckoutComPayments `json:"checkout_com_payments,omitempty"`
}

// NewCheckoutComGatewayDataRelationships instantiates a new CheckoutComGatewayDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutComGatewayDataRelationships() *CheckoutComGatewayDataRelationships {
	this := CheckoutComGatewayDataRelationships{}
	return &this
}

// NewCheckoutComGatewayDataRelationshipsWithDefaults instantiates a new CheckoutComGatewayDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutComGatewayDataRelationshipsWithDefaults() *CheckoutComGatewayDataRelationships {
	this := CheckoutComGatewayDataRelationships{}
	return &this
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *CheckoutComGatewayDataRelationships) GetPaymentMethods() AdyenGatewayDataRelationshipsPaymentMethods {
	if o == nil || o.PaymentMethods == nil {
		var ret AdyenGatewayDataRelationshipsPaymentMethods
		return ret
	}
	return *o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutComGatewayDataRelationships) GetPaymentMethodsOk() (*AdyenGatewayDataRelationshipsPaymentMethods, bool) {
	if o == nil || o.PaymentMethods == nil {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *CheckoutComGatewayDataRelationships) HasPaymentMethods() bool {
	if o != nil && o.PaymentMethods != nil {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given AdyenGatewayDataRelationshipsPaymentMethods and assigns it to the PaymentMethods field.
func (o *CheckoutComGatewayDataRelationships) SetPaymentMethods(v AdyenGatewayDataRelationshipsPaymentMethods) {
	o.PaymentMethods = &v
}

// GetCheckoutComPayments returns the CheckoutComPayments field value if set, zero value otherwise.
func (o *CheckoutComGatewayDataRelationships) GetCheckoutComPayments() CheckoutComGatewayDataRelationshipsCheckoutComPayments {
	if o == nil || o.CheckoutComPayments == nil {
		var ret CheckoutComGatewayDataRelationshipsCheckoutComPayments
		return ret
	}
	return *o.CheckoutComPayments
}

// GetCheckoutComPaymentsOk returns a tuple with the CheckoutComPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutComGatewayDataRelationships) GetCheckoutComPaymentsOk() (*CheckoutComGatewayDataRelationshipsCheckoutComPayments, bool) {
	if o == nil || o.CheckoutComPayments == nil {
		return nil, false
	}
	return o.CheckoutComPayments, true
}

// HasCheckoutComPayments returns a boolean if a field has been set.
func (o *CheckoutComGatewayDataRelationships) HasCheckoutComPayments() bool {
	if o != nil && o.CheckoutComPayments != nil {
		return true
	}

	return false
}

// SetCheckoutComPayments gets a reference to the given CheckoutComGatewayDataRelationshipsCheckoutComPayments and assigns it to the CheckoutComPayments field.
func (o *CheckoutComGatewayDataRelationships) SetCheckoutComPayments(v CheckoutComGatewayDataRelationshipsCheckoutComPayments) {
	o.CheckoutComPayments = &v
}

func (o CheckoutComGatewayDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentMethods != nil {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	if o.CheckoutComPayments != nil {
		toSerialize["checkout_com_payments"] = o.CheckoutComPayments
	}
	return json.Marshal(toSerialize)
}

type NullableCheckoutComGatewayDataRelationships struct {
	value *CheckoutComGatewayDataRelationships
	isSet bool
}

func (v NullableCheckoutComGatewayDataRelationships) Get() *CheckoutComGatewayDataRelationships {
	return v.value
}

func (v *NullableCheckoutComGatewayDataRelationships) Set(val *CheckoutComGatewayDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutComGatewayDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutComGatewayDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutComGatewayDataRelationships(val *CheckoutComGatewayDataRelationships) *NullableCheckoutComGatewayDataRelationships {
	return &NullableCheckoutComGatewayDataRelationships{value: val, isSet: true}
}

func (v NullableCheckoutComGatewayDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutComGatewayDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
