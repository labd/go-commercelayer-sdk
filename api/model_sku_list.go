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

// SkuList struct for SkuList
type SkuList struct {
	Data SkuListData `json:"data"`
}

// NewSkuList instantiates a new SkuList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuList(data SkuListData) *SkuList {
	this := SkuList{}
	this.Data = data
	return &this
}

// NewSkuListWithDefaults instantiates a new SkuList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListWithDefaults() *SkuList {
	this := SkuList{}
	return &this
}

// GetData returns the Data field value
func (o *SkuList) GetData() SkuListData {
	if o == nil {
		var ret SkuListData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SkuList) GetDataOk() (*SkuListData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SkuList) SetData(v SkuListData) {
	o.Data = v
}

func (o SkuList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableSkuList struct {
	value *SkuList
	isSet bool
}

func (v NullableSkuList) Get() *SkuList {
	return v.value
}

func (v *NullableSkuList) Set(val *SkuList) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuList) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuList(val *SkuList) *NullableSkuList {
	return &NullableSkuList{value: val, isSet: true}
}

func (v NullableSkuList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
