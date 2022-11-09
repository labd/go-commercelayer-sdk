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

// KlarnaPaymentUpdate struct for KlarnaPaymentUpdate
type KlarnaPaymentUpdate struct {
	Data KlarnaPaymentUpdateData `json:"data"`
}

// NewKlarnaPaymentUpdate instantiates a new KlarnaPaymentUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlarnaPaymentUpdate(data KlarnaPaymentUpdateData) *KlarnaPaymentUpdate {
	this := KlarnaPaymentUpdate{}
	this.Data = data
	return &this
}

// NewKlarnaPaymentUpdateWithDefaults instantiates a new KlarnaPaymentUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlarnaPaymentUpdateWithDefaults() *KlarnaPaymentUpdate {
	this := KlarnaPaymentUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *KlarnaPaymentUpdate) GetData() KlarnaPaymentUpdateData {
	if o == nil {
		var ret KlarnaPaymentUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *KlarnaPaymentUpdate) GetDataOk() (*KlarnaPaymentUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *KlarnaPaymentUpdate) SetData(v KlarnaPaymentUpdateData) {
	o.Data = v
}

func (o KlarnaPaymentUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableKlarnaPaymentUpdate struct {
	value *KlarnaPaymentUpdate
	isSet bool
}

func (v NullableKlarnaPaymentUpdate) Get() *KlarnaPaymentUpdate {
	return v.value
}

func (v *NullableKlarnaPaymentUpdate) Set(val *KlarnaPaymentUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableKlarnaPaymentUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableKlarnaPaymentUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlarnaPaymentUpdate(val *KlarnaPaymentUpdate) *NullableKlarnaPaymentUpdate {
	return &NullableKlarnaPaymentUpdate{value: val, isSet: true}
}

func (v NullableKlarnaPaymentUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlarnaPaymentUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
