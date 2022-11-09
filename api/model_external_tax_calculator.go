/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.2
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ExternalTaxCalculator struct for ExternalTaxCalculator
type ExternalTaxCalculator struct {
	Data ExternalTaxCalculatorData `json:"data"`
}

// NewExternalTaxCalculator instantiates a new ExternalTaxCalculator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalTaxCalculator(data ExternalTaxCalculatorData) *ExternalTaxCalculator {
	this := ExternalTaxCalculator{}
	this.Data = data
	return &this
}

// NewExternalTaxCalculatorWithDefaults instantiates a new ExternalTaxCalculator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalTaxCalculatorWithDefaults() *ExternalTaxCalculator {
	this := ExternalTaxCalculator{}
	return &this
}

// GetData returns the Data field value
func (o *ExternalTaxCalculator) GetData() ExternalTaxCalculatorData {
	if o == nil {
		var ret ExternalTaxCalculatorData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ExternalTaxCalculator) GetDataOk() (*ExternalTaxCalculatorData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ExternalTaxCalculator) SetData(v ExternalTaxCalculatorData) {
	o.Data = v
}

func (o ExternalTaxCalculator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableExternalTaxCalculator struct {
	value *ExternalTaxCalculator
	isSet bool
}

func (v NullableExternalTaxCalculator) Get() *ExternalTaxCalculator {
	return v.value
}

func (v *NullableExternalTaxCalculator) Set(val *ExternalTaxCalculator) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalTaxCalculator) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalTaxCalculator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalTaxCalculator(val *ExternalTaxCalculator) *NullableExternalTaxCalculator {
	return &NullableExternalTaxCalculator{value: val, isSet: true}
}

func (v NullableExternalTaxCalculator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalTaxCalculator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
