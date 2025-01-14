/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.6.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the TaxRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxRule{}

// TaxRule struct for TaxRule
type TaxRule struct {
	Data *TaxRuleData `json:"data,omitempty"`
}

// NewTaxRule instantiates a new TaxRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxRule() *TaxRule {
	this := TaxRule{}
	return &this
}

// NewTaxRuleWithDefaults instantiates a new TaxRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxRuleWithDefaults() *TaxRule {
	this := TaxRule{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TaxRule) GetData() TaxRuleData {
	if o == nil || IsNil(o.Data) {
		var ret TaxRuleData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxRule) GetDataOk() (*TaxRuleData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TaxRule) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given TaxRuleData and assigns it to the Data field.
func (o *TaxRule) SetData(v TaxRuleData) {
	o.Data = &v
}

func (o TaxRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
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
