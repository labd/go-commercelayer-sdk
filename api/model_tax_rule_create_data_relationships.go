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

// checks if the TaxRuleCreateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxRuleCreateDataRelationships{}

// TaxRuleCreateDataRelationships struct for TaxRuleCreateDataRelationships
type TaxRuleCreateDataRelationships struct {
	ManualTaxCalculator TaxRuleCreateDataRelationshipsManualTaxCalculator `json:"manual_tax_calculator"`
}

// NewTaxRuleCreateDataRelationships instantiates a new TaxRuleCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxRuleCreateDataRelationships(manualTaxCalculator TaxRuleCreateDataRelationshipsManualTaxCalculator) *TaxRuleCreateDataRelationships {
	this := TaxRuleCreateDataRelationships{}
	this.ManualTaxCalculator = manualTaxCalculator
	return &this
}

// NewTaxRuleCreateDataRelationshipsWithDefaults instantiates a new TaxRuleCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxRuleCreateDataRelationshipsWithDefaults() *TaxRuleCreateDataRelationships {
	this := TaxRuleCreateDataRelationships{}
	return &this
}

// GetManualTaxCalculator returns the ManualTaxCalculator field value
func (o *TaxRuleCreateDataRelationships) GetManualTaxCalculator() TaxRuleCreateDataRelationshipsManualTaxCalculator {
	if o == nil {
		var ret TaxRuleCreateDataRelationshipsManualTaxCalculator
		return ret
	}

	return o.ManualTaxCalculator
}

// GetManualTaxCalculatorOk returns a tuple with the ManualTaxCalculator field value
// and a boolean to check if the value has been set.
func (o *TaxRuleCreateDataRelationships) GetManualTaxCalculatorOk() (*TaxRuleCreateDataRelationshipsManualTaxCalculator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManualTaxCalculator, true
}

// SetManualTaxCalculator sets field value
func (o *TaxRuleCreateDataRelationships) SetManualTaxCalculator(v TaxRuleCreateDataRelationshipsManualTaxCalculator) {
	o.ManualTaxCalculator = v
}

func (o TaxRuleCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxRuleCreateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["manual_tax_calculator"] = o.ManualTaxCalculator
	return toSerialize, nil
}

type NullableTaxRuleCreateDataRelationships struct {
	value *TaxRuleCreateDataRelationships
	isSet bool
}

func (v NullableTaxRuleCreateDataRelationships) Get() *TaxRuleCreateDataRelationships {
	return v.value
}

func (v *NullableTaxRuleCreateDataRelationships) Set(val *TaxRuleCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxRuleCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxRuleCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxRuleCreateDataRelationships(val *TaxRuleCreateDataRelationships) *NullableTaxRuleCreateDataRelationships {
	return &NullableTaxRuleCreateDataRelationships{value: val, isSet: true}
}

func (v NullableTaxRuleCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxRuleCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
