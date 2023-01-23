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

// CustomerPaymentSourceCreate struct for CustomerPaymentSourceCreate
type CustomerPaymentSourceCreate struct {
	Data CustomerPaymentSourceCreateData `json:"data"`
}

// NewCustomerPaymentSourceCreate instantiates a new CustomerPaymentSourceCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPaymentSourceCreate(data CustomerPaymentSourceCreateData) *CustomerPaymentSourceCreate {
	this := CustomerPaymentSourceCreate{}
	this.Data = data
	return &this
}

// NewCustomerPaymentSourceCreateWithDefaults instantiates a new CustomerPaymentSourceCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPaymentSourceCreateWithDefaults() *CustomerPaymentSourceCreate {
	this := CustomerPaymentSourceCreate{}
	return &this
}

// GetData returns the Data field value
func (o *CustomerPaymentSourceCreate) GetData() CustomerPaymentSourceCreateData {
	if o == nil {
		var ret CustomerPaymentSourceCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceCreate) GetDataOk() (*CustomerPaymentSourceCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CustomerPaymentSourceCreate) SetData(v CustomerPaymentSourceCreateData) {
	o.Data = v
}

func (o CustomerPaymentSourceCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerPaymentSourceCreate struct {
	value *CustomerPaymentSourceCreate
	isSet bool
}

func (v NullableCustomerPaymentSourceCreate) Get() *CustomerPaymentSourceCreate {
	return v.value
}

func (v *NullableCustomerPaymentSourceCreate) Set(val *CustomerPaymentSourceCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPaymentSourceCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPaymentSourceCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPaymentSourceCreate(val *CustomerPaymentSourceCreate) *NullableCustomerPaymentSourceCreate {
	return &NullableCustomerPaymentSourceCreate{value: val, isSet: true}
}

func (v NullableCustomerPaymentSourceCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPaymentSourceCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
