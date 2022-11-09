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

// AdyenGatewayDataRelationships struct for AdyenGatewayDataRelationships
type AdyenGatewayDataRelationships struct {
	PaymentMethods *AdyenGatewayDataRelationshipsPaymentMethods `json:"payment_methods,omitempty"`
	AdyenPayments  *AdyenGatewayDataRelationshipsAdyenPayments  `json:"adyen_payments,omitempty"`
}

// NewAdyenGatewayDataRelationships instantiates a new AdyenGatewayDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenGatewayDataRelationships() *AdyenGatewayDataRelationships {
	this := AdyenGatewayDataRelationships{}
	return &this
}

// NewAdyenGatewayDataRelationshipsWithDefaults instantiates a new AdyenGatewayDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenGatewayDataRelationshipsWithDefaults() *AdyenGatewayDataRelationships {
	this := AdyenGatewayDataRelationships{}
	return &this
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *AdyenGatewayDataRelationships) GetPaymentMethods() AdyenGatewayDataRelationshipsPaymentMethods {
	if o == nil || o.PaymentMethods == nil {
		var ret AdyenGatewayDataRelationshipsPaymentMethods
		return ret
	}
	return *o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenGatewayDataRelationships) GetPaymentMethodsOk() (*AdyenGatewayDataRelationshipsPaymentMethods, bool) {
	if o == nil || o.PaymentMethods == nil {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *AdyenGatewayDataRelationships) HasPaymentMethods() bool {
	if o != nil && o.PaymentMethods != nil {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given AdyenGatewayDataRelationshipsPaymentMethods and assigns it to the PaymentMethods field.
func (o *AdyenGatewayDataRelationships) SetPaymentMethods(v AdyenGatewayDataRelationshipsPaymentMethods) {
	o.PaymentMethods = &v
}

// GetAdyenPayments returns the AdyenPayments field value if set, zero value otherwise.
func (o *AdyenGatewayDataRelationships) GetAdyenPayments() AdyenGatewayDataRelationshipsAdyenPayments {
	if o == nil || o.AdyenPayments == nil {
		var ret AdyenGatewayDataRelationshipsAdyenPayments
		return ret
	}
	return *o.AdyenPayments
}

// GetAdyenPaymentsOk returns a tuple with the AdyenPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenGatewayDataRelationships) GetAdyenPaymentsOk() (*AdyenGatewayDataRelationshipsAdyenPayments, bool) {
	if o == nil || o.AdyenPayments == nil {
		return nil, false
	}
	return o.AdyenPayments, true
}

// HasAdyenPayments returns a boolean if a field has been set.
func (o *AdyenGatewayDataRelationships) HasAdyenPayments() bool {
	if o != nil && o.AdyenPayments != nil {
		return true
	}

	return false
}

// SetAdyenPayments gets a reference to the given AdyenGatewayDataRelationshipsAdyenPayments and assigns it to the AdyenPayments field.
func (o *AdyenGatewayDataRelationships) SetAdyenPayments(v AdyenGatewayDataRelationshipsAdyenPayments) {
	o.AdyenPayments = &v
}

func (o AdyenGatewayDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentMethods != nil {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	if o.AdyenPayments != nil {
		toSerialize["adyen_payments"] = o.AdyenPayments
	}
	return json.Marshal(toSerialize)
}

type NullableAdyenGatewayDataRelationships struct {
	value *AdyenGatewayDataRelationships
	isSet bool
}

func (v NullableAdyenGatewayDataRelationships) Get() *AdyenGatewayDataRelationships {
	return v.value
}

func (v *NullableAdyenGatewayDataRelationships) Set(val *AdyenGatewayDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAdyenGatewayDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAdyenGatewayDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdyenGatewayDataRelationships(val *AdyenGatewayDataRelationships) *NullableAdyenGatewayDataRelationships {
	return &NullableAdyenGatewayDataRelationships{value: val, isSet: true}
}

func (v NullableAdyenGatewayDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdyenGatewayDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
