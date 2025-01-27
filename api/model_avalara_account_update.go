/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.6.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AvalaraAccountUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvalaraAccountUpdate{}

// AvalaraAccountUpdate struct for AvalaraAccountUpdate
type AvalaraAccountUpdate struct {
	Data AvalaraAccountUpdateData `json:"data"`
}

// NewAvalaraAccountUpdate instantiates a new AvalaraAccountUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvalaraAccountUpdate(data AvalaraAccountUpdateData) *AvalaraAccountUpdate {
	this := AvalaraAccountUpdate{}
	this.Data = data
	return &this
}

// NewAvalaraAccountUpdateWithDefaults instantiates a new AvalaraAccountUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvalaraAccountUpdateWithDefaults() *AvalaraAccountUpdate {
	this := AvalaraAccountUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *AvalaraAccountUpdate) GetData() AvalaraAccountUpdateData {
	if o == nil {
		var ret AvalaraAccountUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AvalaraAccountUpdate) GetDataOk() (*AvalaraAccountUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AvalaraAccountUpdate) SetData(v AvalaraAccountUpdateData) {
	o.Data = v
}

func (o AvalaraAccountUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvalaraAccountUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAvalaraAccountUpdate struct {
	value *AvalaraAccountUpdate
	isSet bool
}

func (v NullableAvalaraAccountUpdate) Get() *AvalaraAccountUpdate {
	return v.value
}

func (v *NullableAvalaraAccountUpdate) Set(val *AvalaraAccountUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAvalaraAccountUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAvalaraAccountUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvalaraAccountUpdate(val *AvalaraAccountUpdate) *NullableAvalaraAccountUpdate {
	return &NullableAvalaraAccountUpdate{value: val, isSet: true}
}

func (v NullableAvalaraAccountUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvalaraAccountUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
