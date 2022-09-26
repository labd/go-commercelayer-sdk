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

// ShippingWeightTierCreate struct for ShippingWeightTierCreate
type ShippingWeightTierCreate struct {
	Data ShippingWeightTierCreateData `json:"data"`
}

// NewShippingWeightTierCreate instantiates a new ShippingWeightTierCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingWeightTierCreate(data ShippingWeightTierCreateData) *ShippingWeightTierCreate {
	this := ShippingWeightTierCreate{}
	this.Data = data
	return &this
}

// NewShippingWeightTierCreateWithDefaults instantiates a new ShippingWeightTierCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingWeightTierCreateWithDefaults() *ShippingWeightTierCreate {
	this := ShippingWeightTierCreate{}
	return &this
}

// GetData returns the Data field value
func (o *ShippingWeightTierCreate) GetData() ShippingWeightTierCreateData {
	if o == nil {
		var ret ShippingWeightTierCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ShippingWeightTierCreate) GetDataOk() (*ShippingWeightTierCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ShippingWeightTierCreate) SetData(v ShippingWeightTierCreateData) {
	o.Data = v
}

func (o ShippingWeightTierCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableShippingWeightTierCreate struct {
	value *ShippingWeightTierCreate
	isSet bool
}

func (v NullableShippingWeightTierCreate) Get() *ShippingWeightTierCreate {
	return v.value
}

func (v *NullableShippingWeightTierCreate) Set(val *ShippingWeightTierCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingWeightTierCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingWeightTierCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingWeightTierCreate(val *ShippingWeightTierCreate) *NullableShippingWeightTierCreate {
	return &NullableShippingWeightTierCreate{value: val, isSet: true}
}

func (v NullableShippingWeightTierCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingWeightTierCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
