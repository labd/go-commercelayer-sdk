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

// ShippingWeightTierResponseList struct for ShippingWeightTierResponseList
type ShippingWeightTierResponseList struct {
	Data []Data `json:"data,omitempty"`
}

// NewShippingWeightTierResponseList instantiates a new ShippingWeightTierResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingWeightTierResponseList() *ShippingWeightTierResponseList {
	this := ShippingWeightTierResponseList{}
	return &this
}

// NewShippingWeightTierResponseListWithDefaults instantiates a new ShippingWeightTierResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingWeightTierResponseListWithDefaults() *ShippingWeightTierResponseList {
	this := ShippingWeightTierResponseList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ShippingWeightTierResponseList) GetData() []Data {
	if o == nil || o.Data == nil {
		var ret []Data
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingWeightTierResponseList) GetDataOk() ([]Data, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ShippingWeightTierResponseList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Data and assigns it to the Data field.
func (o *ShippingWeightTierResponseList) SetData(v []Data) {
	o.Data = v
}

func (o ShippingWeightTierResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableShippingWeightTierResponseList struct {
	value *ShippingWeightTierResponseList
	isSet bool
}

func (v NullableShippingWeightTierResponseList) Get() *ShippingWeightTierResponseList {
	return v.value
}

func (v *NullableShippingWeightTierResponseList) Set(val *ShippingWeightTierResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingWeightTierResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingWeightTierResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingWeightTierResponseList(val *ShippingWeightTierResponseList) *NullableShippingWeightTierResponseList {
	return &NullableShippingWeightTierResponseList{value: val, isSet: true}
}

func (v NullableShippingWeightTierResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingWeightTierResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
