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

// ShippingMethodCreate struct for ShippingMethodCreate
type ShippingMethodCreate struct {
	Data ShippingMethodCreateData `json:"data"`
}

// NewShippingMethodCreate instantiates a new ShippingMethodCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingMethodCreate(data ShippingMethodCreateData) *ShippingMethodCreate {
	this := ShippingMethodCreate{}
	this.Data = data
	return &this
}

// NewShippingMethodCreateWithDefaults instantiates a new ShippingMethodCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingMethodCreateWithDefaults() *ShippingMethodCreate {
	this := ShippingMethodCreate{}
	return &this
}

// GetData returns the Data field value
func (o *ShippingMethodCreate) GetData() ShippingMethodCreateData {
	if o == nil {
		var ret ShippingMethodCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ShippingMethodCreate) GetDataOk() (*ShippingMethodCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ShippingMethodCreate) SetData(v ShippingMethodCreateData) {
	o.Data = v
}

func (o ShippingMethodCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableShippingMethodCreate struct {
	value *ShippingMethodCreate
	isSet bool
}

func (v NullableShippingMethodCreate) Get() *ShippingMethodCreate {
	return v.value
}

func (v *NullableShippingMethodCreate) Set(val *ShippingMethodCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingMethodCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingMethodCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingMethodCreate(val *ShippingMethodCreate) *NullableShippingMethodCreate {
	return &NullableShippingMethodCreate{value: val, isSet: true}
}

func (v NullableShippingMethodCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingMethodCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
