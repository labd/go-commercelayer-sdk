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

// checks if the POSTPaypalGateways201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTPaypalGateways201ResponseDataRelationships{}

// POSTPaypalGateways201ResponseDataRelationships struct for POSTPaypalGateways201ResponseDataRelationships
type POSTPaypalGateways201ResponseDataRelationships struct {
	PaymentMethods *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods  `json:"payment_methods,omitempty"`
	Versions       *POSTAddresses201ResponseDataRelationshipsVersions            `json:"versions,omitempty"`
	PaypalPayments *POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments `json:"paypal_payments,omitempty"`
}

// NewPOSTPaypalGateways201ResponseDataRelationships instantiates a new POSTPaypalGateways201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPaypalGateways201ResponseDataRelationships() *POSTPaypalGateways201ResponseDataRelationships {
	this := POSTPaypalGateways201ResponseDataRelationships{}
	return &this
}

// NewPOSTPaypalGateways201ResponseDataRelationshipsWithDefaults instantiates a new POSTPaypalGateways201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPaypalGateways201ResponseDataRelationshipsWithDefaults() *POSTPaypalGateways201ResponseDataRelationships {
	this := POSTPaypalGateways201ResponseDataRelationships{}
	return &this
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *POSTPaypalGateways201ResponseDataRelationships) GetPaymentMethods() POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods {
	if o == nil || IsNil(o.PaymentMethods) {
		var ret POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods
		return ret
	}
	return *o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPaypalGateways201ResponseDataRelationships) GetPaymentMethodsOk() (*POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods, bool) {
	if o == nil || IsNil(o.PaymentMethods) {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *POSTPaypalGateways201ResponseDataRelationships) HasPaymentMethods() bool {
	if o != nil && !IsNil(o.PaymentMethods) {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods and assigns it to the PaymentMethods field.
func (o *POSTPaypalGateways201ResponseDataRelationships) SetPaymentMethods(v POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods) {
	o.PaymentMethods = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *POSTPaypalGateways201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPaypalGateways201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *POSTPaypalGateways201ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *POSTPaypalGateways201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

// GetPaypalPayments returns the PaypalPayments field value if set, zero value otherwise.
func (o *POSTPaypalGateways201ResponseDataRelationships) GetPaypalPayments() POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments {
	if o == nil || IsNil(o.PaypalPayments) {
		var ret POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments
		return ret
	}
	return *o.PaypalPayments
}

// GetPaypalPaymentsOk returns a tuple with the PaypalPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPaypalGateways201ResponseDataRelationships) GetPaypalPaymentsOk() (*POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments, bool) {
	if o == nil || IsNil(o.PaypalPayments) {
		return nil, false
	}
	return o.PaypalPayments, true
}

// HasPaypalPayments returns a boolean if a field has been set.
func (o *POSTPaypalGateways201ResponseDataRelationships) HasPaypalPayments() bool {
	if o != nil && !IsNil(o.PaypalPayments) {
		return true
	}

	return false
}

// SetPaypalPayments gets a reference to the given POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments and assigns it to the PaypalPayments field.
func (o *POSTPaypalGateways201ResponseDataRelationships) SetPaypalPayments(v POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments) {
	o.PaypalPayments = &v
}

func (o POSTPaypalGateways201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTPaypalGateways201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentMethods) {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	if !IsNil(o.PaypalPayments) {
		toSerialize["paypal_payments"] = o.PaypalPayments
	}
	return toSerialize, nil
}

type NullablePOSTPaypalGateways201ResponseDataRelationships struct {
	value *POSTPaypalGateways201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTPaypalGateways201ResponseDataRelationships) Get() *POSTPaypalGateways201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTPaypalGateways201ResponseDataRelationships) Set(val *POSTPaypalGateways201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPaypalGateways201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPaypalGateways201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPaypalGateways201ResponseDataRelationships(val *POSTPaypalGateways201ResponseDataRelationships) *NullablePOSTPaypalGateways201ResponseDataRelationships {
	return &NullablePOSTPaypalGateways201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTPaypalGateways201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPaypalGateways201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
