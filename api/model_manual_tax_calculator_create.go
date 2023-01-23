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

// ManualTaxCalculatorCreate struct for ManualTaxCalculatorCreate
type ManualTaxCalculatorCreate struct {
	Data ManualTaxCalculatorCreateData `json:"data"`
}

// NewManualTaxCalculatorCreate instantiates a new ManualTaxCalculatorCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualTaxCalculatorCreate(data ManualTaxCalculatorCreateData) *ManualTaxCalculatorCreate {
	this := ManualTaxCalculatorCreate{}
	this.Data = data
	return &this
}

// NewManualTaxCalculatorCreateWithDefaults instantiates a new ManualTaxCalculatorCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualTaxCalculatorCreateWithDefaults() *ManualTaxCalculatorCreate {
	this := ManualTaxCalculatorCreate{}
	return &this
}

// GetData returns the Data field value
func (o *ManualTaxCalculatorCreate) GetData() ManualTaxCalculatorCreateData {
	if o == nil {
		var ret ManualTaxCalculatorCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ManualTaxCalculatorCreate) GetDataOk() (*ManualTaxCalculatorCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ManualTaxCalculatorCreate) SetData(v ManualTaxCalculatorCreateData) {
	o.Data = v
}

func (o ManualTaxCalculatorCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableManualTaxCalculatorCreate struct {
	value *ManualTaxCalculatorCreate
	isSet bool
}

func (v NullableManualTaxCalculatorCreate) Get() *ManualTaxCalculatorCreate {
	return v.value
}

func (v *NullableManualTaxCalculatorCreate) Set(val *ManualTaxCalculatorCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableManualTaxCalculatorCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableManualTaxCalculatorCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualTaxCalculatorCreate(val *ManualTaxCalculatorCreate) *NullableManualTaxCalculatorCreate {
	return &NullableManualTaxCalculatorCreate{value: val, isSet: true}
}

func (v NullableManualTaxCalculatorCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualTaxCalculatorCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
