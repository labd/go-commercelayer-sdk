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

// checks if the POSTKlarnaGateways201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTKlarnaGateways201ResponseDataRelationships{}

// POSTKlarnaGateways201ResponseDataRelationships struct for POSTKlarnaGateways201ResponseDataRelationships
type POSTKlarnaGateways201ResponseDataRelationships struct {
	PaymentMethods *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods  `json:"payment_methods,omitempty"`
	Versions       *POSTAddresses201ResponseDataRelationshipsVersions            `json:"versions,omitempty"`
	KlarnaPayments *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments `json:"klarna_payments,omitempty"`
}

// NewPOSTKlarnaGateways201ResponseDataRelationships instantiates a new POSTKlarnaGateways201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTKlarnaGateways201ResponseDataRelationships() *POSTKlarnaGateways201ResponseDataRelationships {
	this := POSTKlarnaGateways201ResponseDataRelationships{}
	return &this
}

// NewPOSTKlarnaGateways201ResponseDataRelationshipsWithDefaults instantiates a new POSTKlarnaGateways201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTKlarnaGateways201ResponseDataRelationshipsWithDefaults() *POSTKlarnaGateways201ResponseDataRelationships {
	this := POSTKlarnaGateways201ResponseDataRelationships{}
	return &this
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *POSTKlarnaGateways201ResponseDataRelationships) GetPaymentMethods() POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods {
	if o == nil || IsNil(o.PaymentMethods) {
		var ret POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods
		return ret
	}
	return *o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTKlarnaGateways201ResponseDataRelationships) GetPaymentMethodsOk() (*POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods, bool) {
	if o == nil || IsNil(o.PaymentMethods) {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *POSTKlarnaGateways201ResponseDataRelationships) HasPaymentMethods() bool {
	if o != nil && !IsNil(o.PaymentMethods) {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods and assigns it to the PaymentMethods field.
func (o *POSTKlarnaGateways201ResponseDataRelationships) SetPaymentMethods(v POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods) {
	o.PaymentMethods = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *POSTKlarnaGateways201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTKlarnaGateways201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *POSTKlarnaGateways201ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *POSTKlarnaGateways201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

// GetKlarnaPayments returns the KlarnaPayments field value if set, zero value otherwise.
func (o *POSTKlarnaGateways201ResponseDataRelationships) GetKlarnaPayments() POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments {
	if o == nil || IsNil(o.KlarnaPayments) {
		var ret POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments
		return ret
	}
	return *o.KlarnaPayments
}

// GetKlarnaPaymentsOk returns a tuple with the KlarnaPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTKlarnaGateways201ResponseDataRelationships) GetKlarnaPaymentsOk() (*POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments, bool) {
	if o == nil || IsNil(o.KlarnaPayments) {
		return nil, false
	}
	return o.KlarnaPayments, true
}

// HasKlarnaPayments returns a boolean if a field has been set.
func (o *POSTKlarnaGateways201ResponseDataRelationships) HasKlarnaPayments() bool {
	if o != nil && !IsNil(o.KlarnaPayments) {
		return true
	}

	return false
}

// SetKlarnaPayments gets a reference to the given POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments and assigns it to the KlarnaPayments field.
func (o *POSTKlarnaGateways201ResponseDataRelationships) SetKlarnaPayments(v POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments) {
	o.KlarnaPayments = &v
}

func (o POSTKlarnaGateways201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTKlarnaGateways201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentMethods) {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	if !IsNil(o.KlarnaPayments) {
		toSerialize["klarna_payments"] = o.KlarnaPayments
	}
	return toSerialize, nil
}

type NullablePOSTKlarnaGateways201ResponseDataRelationships struct {
	value *POSTKlarnaGateways201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTKlarnaGateways201ResponseDataRelationships) Get() *POSTKlarnaGateways201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTKlarnaGateways201ResponseDataRelationships) Set(val *POSTKlarnaGateways201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTKlarnaGateways201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTKlarnaGateways201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTKlarnaGateways201ResponseDataRelationships(val *POSTKlarnaGateways201ResponseDataRelationships) *NullablePOSTKlarnaGateways201ResponseDataRelationships {
	return &NullablePOSTKlarnaGateways201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTKlarnaGateways201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTKlarnaGateways201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
