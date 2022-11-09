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

// TaxRule struct for TaxRule
type TaxRule struct {
	Data TaxRuleData `json:"data"`
}

// NewTaxRule instantiates a new TaxRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxRule(data TaxRuleData) *TaxRule {
	this := TaxRule{}
	this.Data = data
	return &this
}

// NewTaxRuleWithDefaults instantiates a new TaxRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxRuleWithDefaults() *TaxRule {
	this := TaxRule{}
	return &this
}

// GetData returns the Data field value
func (o *TaxRule) GetData() TaxRuleData {
	if o == nil {
		var ret TaxRuleData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TaxRule) GetDataOk() (*TaxRuleData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TaxRule) SetData(v TaxRuleData) {
	o.Data = v
}

func (o TaxRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableTaxRule struct {
	value *TaxRule
	isSet bool
}

func (v NullableTaxRule) Get() *TaxRule {
	return v.value
}

func (v *NullableTaxRule) Set(val *TaxRule) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxRule) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxRule(val *TaxRule) *NullableTaxRule {
	return &NullableTaxRule{value: val, isSet: true}
}

func (v NullableTaxRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
