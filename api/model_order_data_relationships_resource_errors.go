/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the OrderDataRelationshipsResourceErrors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDataRelationshipsResourceErrors{}

// OrderDataRelationshipsResourceErrors struct for OrderDataRelationshipsResourceErrors
type OrderDataRelationshipsResourceErrors struct {
	Data *OrderDataRelationshipsResourceErrorsData `json:"data,omitempty"`
}

// NewOrderDataRelationshipsResourceErrors instantiates a new OrderDataRelationshipsResourceErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDataRelationshipsResourceErrors() *OrderDataRelationshipsResourceErrors {
	this := OrderDataRelationshipsResourceErrors{}
	return &this
}

// NewOrderDataRelationshipsResourceErrorsWithDefaults instantiates a new OrderDataRelationshipsResourceErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDataRelationshipsResourceErrorsWithDefaults() *OrderDataRelationshipsResourceErrors {
	this := OrderDataRelationshipsResourceErrors{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OrderDataRelationshipsResourceErrors) GetData() OrderDataRelationshipsResourceErrorsData {
	if o == nil || IsNil(o.Data) {
		var ret OrderDataRelationshipsResourceErrorsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataRelationshipsResourceErrors) GetDataOk() (*OrderDataRelationshipsResourceErrorsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OrderDataRelationshipsResourceErrors) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given OrderDataRelationshipsResourceErrorsData and assigns it to the Data field.
func (o *OrderDataRelationshipsResourceErrors) SetData(v OrderDataRelationshipsResourceErrorsData) {
	o.Data = &v
}

func (o OrderDataRelationshipsResourceErrors) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDataRelationshipsResourceErrors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableOrderDataRelationshipsResourceErrors struct {
	value *OrderDataRelationshipsResourceErrors
	isSet bool
}

func (v NullableOrderDataRelationshipsResourceErrors) Get() *OrderDataRelationshipsResourceErrors {
	return v.value
}

func (v *NullableOrderDataRelationshipsResourceErrors) Set(val *OrderDataRelationshipsResourceErrors) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDataRelationshipsResourceErrors) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDataRelationshipsResourceErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDataRelationshipsResourceErrors(val *OrderDataRelationshipsResourceErrors) *NullableOrderDataRelationshipsResourceErrors {
	return &NullableOrderDataRelationshipsResourceErrors{value: val, isSet: true}
}

func (v NullableOrderDataRelationshipsResourceErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDataRelationshipsResourceErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}