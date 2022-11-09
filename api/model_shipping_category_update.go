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

// ShippingCategoryUpdate struct for ShippingCategoryUpdate
type ShippingCategoryUpdate struct {
	Data ShippingCategoryUpdateData `json:"data"`
}

// NewShippingCategoryUpdate instantiates a new ShippingCategoryUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingCategoryUpdate(data ShippingCategoryUpdateData) *ShippingCategoryUpdate {
	this := ShippingCategoryUpdate{}
	this.Data = data
	return &this
}

// NewShippingCategoryUpdateWithDefaults instantiates a new ShippingCategoryUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingCategoryUpdateWithDefaults() *ShippingCategoryUpdate {
	this := ShippingCategoryUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *ShippingCategoryUpdate) GetData() ShippingCategoryUpdateData {
	if o == nil {
		var ret ShippingCategoryUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ShippingCategoryUpdate) GetDataOk() (*ShippingCategoryUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ShippingCategoryUpdate) SetData(v ShippingCategoryUpdateData) {
	o.Data = v
}

func (o ShippingCategoryUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableShippingCategoryUpdate struct {
	value *ShippingCategoryUpdate
	isSet bool
}

func (v NullableShippingCategoryUpdate) Get() *ShippingCategoryUpdate {
	return v.value
}

func (v *NullableShippingCategoryUpdate) Set(val *ShippingCategoryUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingCategoryUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingCategoryUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingCategoryUpdate(val *ShippingCategoryUpdate) *NullableShippingCategoryUpdate {
	return &NullableShippingCategoryUpdate{value: val, isSet: true}
}

func (v NullableShippingCategoryUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingCategoryUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
