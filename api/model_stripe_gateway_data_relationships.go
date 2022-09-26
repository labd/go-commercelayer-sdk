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

// StripeGatewayDataRelationships struct for StripeGatewayDataRelationships
type StripeGatewayDataRelationships struct {
	PaymentMethods *AdyenGatewayDataRelationshipsPaymentMethods  `json:"payment_methods,omitempty"`
	StripePayments *StripeGatewayDataRelationshipsStripePayments `json:"stripe_payments,omitempty"`
}

// NewStripeGatewayDataRelationships instantiates a new StripeGatewayDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripeGatewayDataRelationships() *StripeGatewayDataRelationships {
	this := StripeGatewayDataRelationships{}
	return &this
}

// NewStripeGatewayDataRelationshipsWithDefaults instantiates a new StripeGatewayDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripeGatewayDataRelationshipsWithDefaults() *StripeGatewayDataRelationships {
	this := StripeGatewayDataRelationships{}
	return &this
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *StripeGatewayDataRelationships) GetPaymentMethods() AdyenGatewayDataRelationshipsPaymentMethods {
	if o == nil || o.PaymentMethods == nil {
		var ret AdyenGatewayDataRelationshipsPaymentMethods
		return ret
	}
	return *o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeGatewayDataRelationships) GetPaymentMethodsOk() (*AdyenGatewayDataRelationshipsPaymentMethods, bool) {
	if o == nil || o.PaymentMethods == nil {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *StripeGatewayDataRelationships) HasPaymentMethods() bool {
	if o != nil && o.PaymentMethods != nil {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given AdyenGatewayDataRelationshipsPaymentMethods and assigns it to the PaymentMethods field.
func (o *StripeGatewayDataRelationships) SetPaymentMethods(v AdyenGatewayDataRelationshipsPaymentMethods) {
	o.PaymentMethods = &v
}

// GetStripePayments returns the StripePayments field value if set, zero value otherwise.
func (o *StripeGatewayDataRelationships) GetStripePayments() StripeGatewayDataRelationshipsStripePayments {
	if o == nil || o.StripePayments == nil {
		var ret StripeGatewayDataRelationshipsStripePayments
		return ret
	}
	return *o.StripePayments
}

// GetStripePaymentsOk returns a tuple with the StripePayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeGatewayDataRelationships) GetStripePaymentsOk() (*StripeGatewayDataRelationshipsStripePayments, bool) {
	if o == nil || o.StripePayments == nil {
		return nil, false
	}
	return o.StripePayments, true
}

// HasStripePayments returns a boolean if a field has been set.
func (o *StripeGatewayDataRelationships) HasStripePayments() bool {
	if o != nil && o.StripePayments != nil {
		return true
	}

	return false
}

// SetStripePayments gets a reference to the given StripeGatewayDataRelationshipsStripePayments and assigns it to the StripePayments field.
func (o *StripeGatewayDataRelationships) SetStripePayments(v StripeGatewayDataRelationshipsStripePayments) {
	o.StripePayments = &v
}

func (o StripeGatewayDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentMethods != nil {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	if o.StripePayments != nil {
		toSerialize["stripe_payments"] = o.StripePayments
	}
	return json.Marshal(toSerialize)
}

type NullableStripeGatewayDataRelationships struct {
	value *StripeGatewayDataRelationships
	isSet bool
}

func (v NullableStripeGatewayDataRelationships) Get() *StripeGatewayDataRelationships {
	return v.value
}

func (v *NullableStripeGatewayDataRelationships) Set(val *StripeGatewayDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableStripeGatewayDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableStripeGatewayDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripeGatewayDataRelationships(val *StripeGatewayDataRelationships) *NullableStripeGatewayDataRelationships {
	return &NullableStripeGatewayDataRelationships{value: val, isSet: true}
}

func (v NullableStripeGatewayDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripeGatewayDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
