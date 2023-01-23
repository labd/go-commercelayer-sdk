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

// OrderCopy struct for OrderCopy
type OrderCopy struct {
	Data *OrderCopyData `json:"data,omitempty"`
}

// NewOrderCopy instantiates a new OrderCopy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCopy() *OrderCopy {
	this := OrderCopy{}
	return &this
}

// NewOrderCopyWithDefaults instantiates a new OrderCopy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCopyWithDefaults() *OrderCopy {
	this := OrderCopy{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OrderCopy) GetData() OrderCopyData {
	if o == nil || o.Data == nil {
		var ret OrderCopyData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCopy) GetDataOk() (*OrderCopyData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OrderCopy) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given OrderCopyData and assigns it to the Data field.
func (o *OrderCopy) SetData(v OrderCopyData) {
	o.Data = &v
}

func (o OrderCopy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableOrderCopy struct {
	value *OrderCopy
	isSet bool
}

func (v NullableOrderCopy) Get() *OrderCopy {
	return v.value
}

func (v *NullableOrderCopy) Set(val *OrderCopy) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCopy) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCopy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCopy(val *OrderCopy) *NullableOrderCopy {
	return &NullableOrderCopy{value: val, isSet: true}
}

func (v NullableOrderCopy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCopy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
