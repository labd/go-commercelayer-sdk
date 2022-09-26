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

// CustomerPasswordResetCreate struct for CustomerPasswordResetCreate
type CustomerPasswordResetCreate struct {
	Data CustomerPasswordResetCreateData `json:"data"`
}

// NewCustomerPasswordResetCreate instantiates a new CustomerPasswordResetCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPasswordResetCreate(data CustomerPasswordResetCreateData) *CustomerPasswordResetCreate {
	this := CustomerPasswordResetCreate{}
	this.Data = data
	return &this
}

// NewCustomerPasswordResetCreateWithDefaults instantiates a new CustomerPasswordResetCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPasswordResetCreateWithDefaults() *CustomerPasswordResetCreate {
	this := CustomerPasswordResetCreate{}
	return &this
}

// GetData returns the Data field value
func (o *CustomerPasswordResetCreate) GetData() CustomerPasswordResetCreateData {
	if o == nil {
		var ret CustomerPasswordResetCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CustomerPasswordResetCreate) GetDataOk() (*CustomerPasswordResetCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CustomerPasswordResetCreate) SetData(v CustomerPasswordResetCreateData) {
	o.Data = v
}

func (o CustomerPasswordResetCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerPasswordResetCreate struct {
	value *CustomerPasswordResetCreate
	isSet bool
}

func (v NullableCustomerPasswordResetCreate) Get() *CustomerPasswordResetCreate {
	return v.value
}

func (v *NullableCustomerPasswordResetCreate) Set(val *CustomerPasswordResetCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPasswordResetCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPasswordResetCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPasswordResetCreate(val *CustomerPasswordResetCreate) *NullableCustomerPasswordResetCreate {
	return &NullableCustomerPasswordResetCreate{value: val, isSet: true}
}

func (v NullableCustomerPasswordResetCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPasswordResetCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
