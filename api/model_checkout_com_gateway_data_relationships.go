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

// checks if the CheckoutComGatewayDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckoutComGatewayDataRelationships{}

// CheckoutComGatewayDataRelationships struct for CheckoutComGatewayDataRelationships
type CheckoutComGatewayDataRelationships struct {
	PaymentMethods      *AdyenGatewayDataRelationshipsPaymentMethods            `json:"payment_methods,omitempty"`
	Versions            *AddressDataRelationshipsVersions                       `json:"versions,omitempty"`
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
	if o == nil || IsNil(o.PaymentMethods) {
		var ret AdyenGatewayDataRelationshipsPaymentMethods
		return ret
	}
	return *o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutComGatewayDataRelationships) GetPaymentMethodsOk() (*AdyenGatewayDataRelationshipsPaymentMethods, bool) {
	if o == nil || IsNil(o.PaymentMethods) {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *CheckoutComGatewayDataRelationships) HasPaymentMethods() bool {
	if o != nil && !IsNil(o.PaymentMethods) {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given AdyenGatewayDataRelationshipsPaymentMethods and assigns it to the PaymentMethods field.
func (o *CheckoutComGatewayDataRelationships) SetPaymentMethods(v AdyenGatewayDataRelationshipsPaymentMethods) {
	o.PaymentMethods = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *CheckoutComGatewayDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutComGatewayDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *CheckoutComGatewayDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *CheckoutComGatewayDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

// GetCheckoutComPayments returns the CheckoutComPayments field value if set, zero value otherwise.
func (o *CheckoutComGatewayDataRelationships) GetCheckoutComPayments() CheckoutComGatewayDataRelationshipsCheckoutComPayments {
	if o == nil || IsNil(o.CheckoutComPayments) {
		var ret CheckoutComGatewayDataRelationshipsCheckoutComPayments
		return ret
	}
	return *o.CheckoutComPayments
}

// GetCheckoutComPaymentsOk returns a tuple with the CheckoutComPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutComGatewayDataRelationships) GetCheckoutComPaymentsOk() (*CheckoutComGatewayDataRelationshipsCheckoutComPayments, bool) {
	if o == nil || IsNil(o.CheckoutComPayments) {
		return nil, false
	}
	return o.CheckoutComPayments, true
}

// HasCheckoutComPayments returns a boolean if a field has been set.
func (o *CheckoutComGatewayDataRelationships) HasCheckoutComPayments() bool {
	if o != nil && !IsNil(o.CheckoutComPayments) {
		return true
	}

	return false
}

// SetCheckoutComPayments gets a reference to the given CheckoutComGatewayDataRelationshipsCheckoutComPayments and assigns it to the CheckoutComPayments field.
func (o *CheckoutComGatewayDataRelationships) SetCheckoutComPayments(v CheckoutComGatewayDataRelationshipsCheckoutComPayments) {
	o.CheckoutComPayments = &v
}

func (o CheckoutComGatewayDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutComGatewayDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentMethods) {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	if !IsNil(o.CheckoutComPayments) {
		toSerialize["checkout_com_payments"] = o.CheckoutComPayments
	}
	return toSerialize, nil
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
