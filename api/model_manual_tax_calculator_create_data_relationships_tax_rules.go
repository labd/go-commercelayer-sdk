/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ManualTaxCalculatorCreateDataRelationshipsTaxRules struct for ManualTaxCalculatorCreateDataRelationshipsTaxRules
type ManualTaxCalculatorCreateDataRelationshipsTaxRules struct {
	Data ManualTaxCalculatorDataRelationshipsTaxRulesData `json:"data"`
}

// NewManualTaxCalculatorCreateDataRelationshipsTaxRules instantiates a new ManualTaxCalculatorCreateDataRelationshipsTaxRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualTaxCalculatorCreateDataRelationshipsTaxRules(data ManualTaxCalculatorDataRelationshipsTaxRulesData) *ManualTaxCalculatorCreateDataRelationshipsTaxRules {
	this := ManualTaxCalculatorCreateDataRelationshipsTaxRules{}
	this.Data = data
	return &this
}

// NewManualTaxCalculatorCreateDataRelationshipsTaxRulesWithDefaults instantiates a new ManualTaxCalculatorCreateDataRelationshipsTaxRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualTaxCalculatorCreateDataRelationshipsTaxRulesWithDefaults() *ManualTaxCalculatorCreateDataRelationshipsTaxRules {
	this := ManualTaxCalculatorCreateDataRelationshipsTaxRules{}
	return &this
}

// GetData returns the Data field value
func (o *ManualTaxCalculatorCreateDataRelationshipsTaxRules) GetData() ManualTaxCalculatorDataRelationshipsTaxRulesData {
	if o == nil {
		var ret ManualTaxCalculatorDataRelationshipsTaxRulesData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ManualTaxCalculatorCreateDataRelationshipsTaxRules) GetDataOk() (*ManualTaxCalculatorDataRelationshipsTaxRulesData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ManualTaxCalculatorCreateDataRelationshipsTaxRules) SetData(v ManualTaxCalculatorDataRelationshipsTaxRulesData) {
	o.Data = v
}

func (o ManualTaxCalculatorCreateDataRelationshipsTaxRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableManualTaxCalculatorCreateDataRelationshipsTaxRules struct {
	value *ManualTaxCalculatorCreateDataRelationshipsTaxRules
	isSet bool
}

func (v NullableManualTaxCalculatorCreateDataRelationshipsTaxRules) Get() *ManualTaxCalculatorCreateDataRelationshipsTaxRules {
	return v.value
}

func (v *NullableManualTaxCalculatorCreateDataRelationshipsTaxRules) Set(val *ManualTaxCalculatorCreateDataRelationshipsTaxRules) {
	v.value = val
	v.isSet = true
}

func (v NullableManualTaxCalculatorCreateDataRelationshipsTaxRules) IsSet() bool {
	return v.isSet
}

func (v *NullableManualTaxCalculatorCreateDataRelationshipsTaxRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualTaxCalculatorCreateDataRelationshipsTaxRules(val *ManualTaxCalculatorCreateDataRelationshipsTaxRules) *NullableManualTaxCalculatorCreateDataRelationshipsTaxRules {
	return &NullableManualTaxCalculatorCreateDataRelationshipsTaxRules{value: val, isSet: true}
}

func (v NullableManualTaxCalculatorCreateDataRelationshipsTaxRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualTaxCalculatorCreateDataRelationshipsTaxRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
