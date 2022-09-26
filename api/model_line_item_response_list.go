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

// LineItemResponseList struct for LineItemResponseList
type LineItemResponseList struct {
	Data []Data `json:"data,omitempty"`
}

// NewLineItemResponseList instantiates a new LineItemResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemResponseList() *LineItemResponseList {
	this := LineItemResponseList{}
	return &this
}

// NewLineItemResponseListWithDefaults instantiates a new LineItemResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemResponseListWithDefaults() *LineItemResponseList {
	this := LineItemResponseList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LineItemResponseList) GetData() []Data {
	if o == nil || o.Data == nil {
		var ret []Data
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemResponseList) GetDataOk() ([]Data, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LineItemResponseList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Data and assigns it to the Data field.
func (o *LineItemResponseList) SetData(v []Data) {
	o.Data = v
}

func (o LineItemResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableLineItemResponseList struct {
	value *LineItemResponseList
	isSet bool
}

func (v NullableLineItemResponseList) Get() *LineItemResponseList {
	return v.value
}

func (v *NullableLineItemResponseList) Set(val *LineItemResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemResponseList(val *LineItemResponseList) *NullableLineItemResponseList {
	return &NullableLineItemResponseList{value: val, isSet: true}
}

func (v NullableLineItemResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
