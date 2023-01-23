/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CheckoutComPaymentUpdate struct for CheckoutComPaymentUpdate
type CheckoutComPaymentUpdate struct {
	Data CheckoutComPaymentUpdateData `json:"data"`
}

// NewCheckoutComPaymentUpdate instantiates a new CheckoutComPaymentUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutComPaymentUpdate(data CheckoutComPaymentUpdateData) *CheckoutComPaymentUpdate {
	this := CheckoutComPaymentUpdate{}
	this.Data = data
	return &this
}

// NewCheckoutComPaymentUpdateWithDefaults instantiates a new CheckoutComPaymentUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutComPaymentUpdateWithDefaults() *CheckoutComPaymentUpdate {
	this := CheckoutComPaymentUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *CheckoutComPaymentUpdate) GetData() CheckoutComPaymentUpdateData {
	if o == nil {
		var ret CheckoutComPaymentUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CheckoutComPaymentUpdate) GetDataOk() (*CheckoutComPaymentUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CheckoutComPaymentUpdate) SetData(v CheckoutComPaymentUpdateData) {
	o.Data = v
}

func (o CheckoutComPaymentUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCheckoutComPaymentUpdate struct {
	value *CheckoutComPaymentUpdate
	isSet bool
}

func (v NullableCheckoutComPaymentUpdate) Get() *CheckoutComPaymentUpdate {
	return v.value
}

func (v *NullableCheckoutComPaymentUpdate) Set(val *CheckoutComPaymentUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutComPaymentUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutComPaymentUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutComPaymentUpdate(val *CheckoutComPaymentUpdate) *NullableCheckoutComPaymentUpdate {
	return &NullableCheckoutComPaymentUpdate{value: val, isSet: true}
}

func (v NullableCheckoutComPaymentUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutComPaymentUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
