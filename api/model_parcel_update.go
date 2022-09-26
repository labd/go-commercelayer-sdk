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

// ParcelUpdate struct for ParcelUpdate
type ParcelUpdate struct {
	Data ParcelUpdateData `json:"data"`
}

// NewParcelUpdate instantiates a new ParcelUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParcelUpdate(data ParcelUpdateData) *ParcelUpdate {
	this := ParcelUpdate{}
	this.Data = data
	return &this
}

// NewParcelUpdateWithDefaults instantiates a new ParcelUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParcelUpdateWithDefaults() *ParcelUpdate {
	this := ParcelUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *ParcelUpdate) GetData() ParcelUpdateData {
	if o == nil {
		var ret ParcelUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ParcelUpdate) GetDataOk() (*ParcelUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ParcelUpdate) SetData(v ParcelUpdateData) {
	o.Data = v
}

func (o ParcelUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableParcelUpdate struct {
	value *ParcelUpdate
	isSet bool
}

func (v NullableParcelUpdate) Get() *ParcelUpdate {
	return v.value
}

func (v *NullableParcelUpdate) Set(val *ParcelUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableParcelUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableParcelUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParcelUpdate(val *ParcelUpdate) *NullableParcelUpdate {
	return &NullableParcelUpdate{value: val, isSet: true}
}

func (v NullableParcelUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParcelUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
