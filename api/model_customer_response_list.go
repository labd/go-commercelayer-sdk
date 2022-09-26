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

// CustomerResponseList struct for CustomerResponseList
type CustomerResponseList struct {
	Data []Data `json:"data,omitempty"`
}

// NewCustomerResponseList instantiates a new CustomerResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerResponseList() *CustomerResponseList {
	this := CustomerResponseList{}
	return &this
}

// NewCustomerResponseListWithDefaults instantiates a new CustomerResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerResponseListWithDefaults() *CustomerResponseList {
	this := CustomerResponseList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CustomerResponseList) GetData() []Data {
	if o == nil || o.Data == nil {
		var ret []Data
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponseList) GetDataOk() ([]Data, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CustomerResponseList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Data and assigns it to the Data field.
func (o *CustomerResponseList) SetData(v []Data) {
	o.Data = v
}

func (o CustomerResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerResponseList struct {
	value *CustomerResponseList
	isSet bool
}

func (v NullableCustomerResponseList) Get() *CustomerResponseList {
	return v.value
}

func (v *NullableCustomerResponseList) Set(val *CustomerResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerResponseList(val *CustomerResponseList) *NullableCustomerResponseList {
	return &NullableCustomerResponseList{value: val, isSet: true}
}

func (v NullableCustomerResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
