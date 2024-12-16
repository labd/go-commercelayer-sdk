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

// checks if the AdyenGatewayDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdyenGatewayDataRelationships{}

// AdyenGatewayDataRelationships struct for AdyenGatewayDataRelationships
type AdyenGatewayDataRelationships struct {
	PaymentMethods *AdyenGatewayDataRelationshipsPaymentMethods `json:"payment_methods,omitempty"`
	Versions       *AddressDataRelationshipsVersions            `json:"versions,omitempty"`
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
	if o == nil || IsNil(o.PaymentMethods) {
		var ret AdyenGatewayDataRelationshipsPaymentMethods
		return ret
	}
	return *o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenGatewayDataRelationships) GetPaymentMethodsOk() (*AdyenGatewayDataRelationshipsPaymentMethods, bool) {
	if o == nil || IsNil(o.PaymentMethods) {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *AdyenGatewayDataRelationships) HasPaymentMethods() bool {
	if o != nil && !IsNil(o.PaymentMethods) {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given AdyenGatewayDataRelationshipsPaymentMethods and assigns it to the PaymentMethods field.
func (o *AdyenGatewayDataRelationships) SetPaymentMethods(v AdyenGatewayDataRelationshipsPaymentMethods) {
	o.PaymentMethods = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *AdyenGatewayDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenGatewayDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *AdyenGatewayDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *AdyenGatewayDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

// GetAdyenPayments returns the AdyenPayments field value if set, zero value otherwise.
func (o *AdyenGatewayDataRelationships) GetAdyenPayments() AdyenGatewayDataRelationshipsAdyenPayments {
	if o == nil || IsNil(o.AdyenPayments) {
		var ret AdyenGatewayDataRelationshipsAdyenPayments
		return ret
	}
	return *o.AdyenPayments
}

// GetAdyenPaymentsOk returns a tuple with the AdyenPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenGatewayDataRelationships) GetAdyenPaymentsOk() (*AdyenGatewayDataRelationshipsAdyenPayments, bool) {
	if o == nil || IsNil(o.AdyenPayments) {
		return nil, false
	}
	return o.AdyenPayments, true
}

// HasAdyenPayments returns a boolean if a field has been set.
func (o *AdyenGatewayDataRelationships) HasAdyenPayments() bool {
	if o != nil && !IsNil(o.AdyenPayments) {
		return true
	}

	return false
}

// SetAdyenPayments gets a reference to the given AdyenGatewayDataRelationshipsAdyenPayments and assigns it to the AdyenPayments field.
func (o *AdyenGatewayDataRelationships) SetAdyenPayments(v AdyenGatewayDataRelationshipsAdyenPayments) {
	o.AdyenPayments = &v
}

func (o AdyenGatewayDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdyenGatewayDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentMethods) {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	if !IsNil(o.AdyenPayments) {
		toSerialize["adyen_payments"] = o.AdyenPayments
	}
	return toSerialize, nil
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
