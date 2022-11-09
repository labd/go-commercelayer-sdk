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

// SkuListItemUpdate struct for SkuListItemUpdate
type SkuListItemUpdate struct {
	Data SkuListItemUpdateData `json:"data"`
}

// NewSkuListItemUpdate instantiates a new SkuListItemUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuListItemUpdate(data SkuListItemUpdateData) *SkuListItemUpdate {
	this := SkuListItemUpdate{}
	this.Data = data
	return &this
}

// NewSkuListItemUpdateWithDefaults instantiates a new SkuListItemUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListItemUpdateWithDefaults() *SkuListItemUpdate {
	this := SkuListItemUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *SkuListItemUpdate) GetData() SkuListItemUpdateData {
	if o == nil {
		var ret SkuListItemUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SkuListItemUpdate) GetDataOk() (*SkuListItemUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SkuListItemUpdate) SetData(v SkuListItemUpdateData) {
	o.Data = v
}

func (o SkuListItemUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableSkuListItemUpdate struct {
	value *SkuListItemUpdate
	isSet bool
}

func (v NullableSkuListItemUpdate) Get() *SkuListItemUpdate {
	return v.value
}

func (v *NullableSkuListItemUpdate) Set(val *SkuListItemUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuListItemUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuListItemUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuListItemUpdate(val *SkuListItemUpdate) *NullableSkuListItemUpdate {
	return &NullableSkuListItemUpdate{value: val, isSet: true}
}

func (v NullableSkuListItemUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuListItemUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
