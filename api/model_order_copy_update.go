/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the OrderCopyUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCopyUpdate{}

// OrderCopyUpdate struct for OrderCopyUpdate
type OrderCopyUpdate struct {
	Data OrderCopyUpdateData `json:"data"`
}

// NewOrderCopyUpdate instantiates a new OrderCopyUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCopyUpdate(data OrderCopyUpdateData) *OrderCopyUpdate {
	this := OrderCopyUpdate{}
	this.Data = data
	return &this
}

// NewOrderCopyUpdateWithDefaults instantiates a new OrderCopyUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCopyUpdateWithDefaults() *OrderCopyUpdate {
	this := OrderCopyUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *OrderCopyUpdate) GetData() OrderCopyUpdateData {
	if o == nil {
		var ret OrderCopyUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OrderCopyUpdate) GetDataOk() (*OrderCopyUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *OrderCopyUpdate) SetData(v OrderCopyUpdateData) {
	o.Data = v
}

func (o OrderCopyUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCopyUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableOrderCopyUpdate struct {
	value *OrderCopyUpdate
	isSet bool
}

func (v NullableOrderCopyUpdate) Get() *OrderCopyUpdate {
	return v.value
}

func (v *NullableOrderCopyUpdate) Set(val *OrderCopyUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCopyUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCopyUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCopyUpdate(val *OrderCopyUpdate) *NullableOrderCopyUpdate {
	return &NullableOrderCopyUpdate{value: val, isSet: true}
}

func (v NullableOrderCopyUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCopyUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}