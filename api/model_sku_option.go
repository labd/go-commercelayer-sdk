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

// checks if the SkuOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkuOption{}

// SkuOption struct for SkuOption
type SkuOption struct {
	Data *SkuOptionData `json:"data,omitempty"`
}

// NewSkuOption instantiates a new SkuOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuOption() *SkuOption {
	this := SkuOption{}
	return &this
}

// NewSkuOptionWithDefaults instantiates a new SkuOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuOptionWithDefaults() *SkuOption {
	this := SkuOption{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SkuOption) GetData() SkuOptionData {
	if o == nil || IsNil(o.Data) {
		var ret SkuOptionData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuOption) GetDataOk() (*SkuOptionData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SkuOption) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given SkuOptionData and assigns it to the Data field.
func (o *SkuOption) SetData(v SkuOptionData) {
	o.Data = &v
}

func (o SkuOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkuOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableSkuOption struct {
	value *SkuOption
	isSet bool
}

func (v NullableSkuOption) Get() *SkuOption {
	return v.value
}

func (v *NullableSkuOption) Set(val *SkuOption) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuOption) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuOption(val *SkuOption) *NullableSkuOption {
	return &NullableSkuOption{value: val, isSet: true}
}

func (v NullableSkuOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
