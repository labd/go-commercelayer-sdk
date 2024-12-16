/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.5.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ReturnLineItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReturnLineItem{}

// ReturnLineItem struct for ReturnLineItem
type ReturnLineItem struct {
	Data *ReturnLineItemData `json:"data,omitempty"`
}

// NewReturnLineItem instantiates a new ReturnLineItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnLineItem() *ReturnLineItem {
	this := ReturnLineItem{}
	return &this
}

// NewReturnLineItemWithDefaults instantiates a new ReturnLineItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnLineItemWithDefaults() *ReturnLineItem {
	this := ReturnLineItem{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ReturnLineItem) GetData() ReturnLineItemData {
	if o == nil || IsNil(o.Data) {
		var ret ReturnLineItemData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLineItem) GetDataOk() (*ReturnLineItemData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ReturnLineItem) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ReturnLineItemData and assigns it to the Data field.
func (o *ReturnLineItem) SetData(v ReturnLineItemData) {
	o.Data = &v
}

func (o ReturnLineItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReturnLineItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableReturnLineItem struct {
	value *ReturnLineItem
	isSet bool
}

func (v NullableReturnLineItem) Get() *ReturnLineItem {
	return v.value
}

func (v *NullableReturnLineItem) Set(val *ReturnLineItem) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnLineItem) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnLineItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnLineItem(val *ReturnLineItem) *NullableReturnLineItem {
	return &NullableReturnLineItem{value: val, isSet: true}
}

func (v NullableReturnLineItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnLineItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
