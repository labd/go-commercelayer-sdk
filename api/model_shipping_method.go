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

// checks if the ShippingMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingMethod{}

// ShippingMethod struct for ShippingMethod
type ShippingMethod struct {
	Data *ShippingMethodData `json:"data,omitempty"`
}

// NewShippingMethod instantiates a new ShippingMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingMethod() *ShippingMethod {
	this := ShippingMethod{}
	return &this
}

// NewShippingMethodWithDefaults instantiates a new ShippingMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingMethodWithDefaults() *ShippingMethod {
	this := ShippingMethod{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ShippingMethod) GetData() ShippingMethodData {
	if o == nil || IsNil(o.Data) {
		var ret ShippingMethodData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethod) GetDataOk() (*ShippingMethodData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ShippingMethod) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ShippingMethodData and assigns it to the Data field.
func (o *ShippingMethod) SetData(v ShippingMethodData) {
	o.Data = &v
}

func (o ShippingMethod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShippingMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableShippingMethod struct {
	value *ShippingMethod
	isSet bool
}

func (v NullableShippingMethod) Get() *ShippingMethod {
	return v.value
}

func (v *NullableShippingMethod) Set(val *ShippingMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingMethod(val *ShippingMethod) *NullableShippingMethod {
	return &NullableShippingMethod{value: val, isSet: true}
}

func (v NullableShippingMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
