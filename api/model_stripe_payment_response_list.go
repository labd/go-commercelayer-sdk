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

// StripePaymentResponseList struct for StripePaymentResponseList
type StripePaymentResponseList struct {
	Data []Data `json:"data,omitempty"`
}

// NewStripePaymentResponseList instantiates a new StripePaymentResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripePaymentResponseList() *StripePaymentResponseList {
	this := StripePaymentResponseList{}
	return &this
}

// NewStripePaymentResponseListWithDefaults instantiates a new StripePaymentResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripePaymentResponseListWithDefaults() *StripePaymentResponseList {
	this := StripePaymentResponseList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *StripePaymentResponseList) GetData() []Data {
	if o == nil || o.Data == nil {
		var ret []Data
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripePaymentResponseList) GetDataOk() ([]Data, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *StripePaymentResponseList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Data and assigns it to the Data field.
func (o *StripePaymentResponseList) SetData(v []Data) {
	o.Data = v
}

func (o StripePaymentResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableStripePaymentResponseList struct {
	value *StripePaymentResponseList
	isSet bool
}

func (v NullableStripePaymentResponseList) Get() *StripePaymentResponseList {
	return v.value
}

func (v *NullableStripePaymentResponseList) Set(val *StripePaymentResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableStripePaymentResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableStripePaymentResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripePaymentResponseList(val *StripePaymentResponseList) *NullableStripePaymentResponseList {
	return &NullableStripePaymentResponseList{value: val, isSet: true}
}

func (v NullableStripePaymentResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripePaymentResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
