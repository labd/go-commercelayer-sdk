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

// ShippingZone struct for ShippingZone
type ShippingZone struct {
	Data *ShippingZoneData `json:"data,omitempty"`
}

// NewShippingZone instantiates a new ShippingZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingZone() *ShippingZone {
	this := ShippingZone{}
	return &this
}

// NewShippingZoneWithDefaults instantiates a new ShippingZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingZoneWithDefaults() *ShippingZone {
	this := ShippingZone{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ShippingZone) GetData() ShippingZoneData {
	if o == nil || o.Data == nil {
		var ret ShippingZoneData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingZone) GetDataOk() (*ShippingZoneData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ShippingZone) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ShippingZoneData and assigns it to the Data field.
func (o *ShippingZone) SetData(v ShippingZoneData) {
	o.Data = &v
}

func (o ShippingZone) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableShippingZone struct {
	value *ShippingZone
	isSet bool
}

func (v NullableShippingZone) Get() *ShippingZone {
	return v.value
}

func (v *NullableShippingZone) Set(val *ShippingZone) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingZone) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingZone(val *ShippingZone) *NullableShippingZone {
	return &NullableShippingZone{value: val, isSet: true}
}

func (v NullableShippingZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
