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

// GETManualGateways200ResponseDataInnerRelationships struct for GETManualGateways200ResponseDataInnerRelationships
type GETManualGateways200ResponseDataInnerRelationships struct {
	PaymentMethods *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods `json:"payment_methods,omitempty"`
}

// NewGETManualGateways200ResponseDataInnerRelationships instantiates a new GETManualGateways200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETManualGateways200ResponseDataInnerRelationships() *GETManualGateways200ResponseDataInnerRelationships {
	this := GETManualGateways200ResponseDataInnerRelationships{}
	return &this
}

// NewGETManualGateways200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETManualGateways200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETManualGateways200ResponseDataInnerRelationshipsWithDefaults() *GETManualGateways200ResponseDataInnerRelationships {
	this := GETManualGateways200ResponseDataInnerRelationships{}
	return &this
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *GETManualGateways200ResponseDataInnerRelationships) GetPaymentMethods() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods {
	if o == nil || o.PaymentMethods == nil {
		var ret GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods
		return ret
	}
	return *o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETManualGateways200ResponseDataInnerRelationships) GetPaymentMethodsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool) {
	if o == nil || o.PaymentMethods == nil {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *GETManualGateways200ResponseDataInnerRelationships) HasPaymentMethods() bool {
	if o != nil && o.PaymentMethods != nil {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods and assigns it to the PaymentMethods field.
func (o *GETManualGateways200ResponseDataInnerRelationships) SetPaymentMethods(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods) {
	o.PaymentMethods = &v
}

func (o GETManualGateways200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentMethods != nil {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	return json.Marshal(toSerialize)
}

type NullableGETManualGateways200ResponseDataInnerRelationships struct {
	value *GETManualGateways200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETManualGateways200ResponseDataInnerRelationships) Get() *GETManualGateways200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETManualGateways200ResponseDataInnerRelationships) Set(val *GETManualGateways200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETManualGateways200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETManualGateways200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETManualGateways200ResponseDataInnerRelationships(val *GETManualGateways200ResponseDataInnerRelationships) *NullableGETManualGateways200ResponseDataInnerRelationships {
	return &NullableGETManualGateways200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETManualGateways200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETManualGateways200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}