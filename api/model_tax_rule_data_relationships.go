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

// checks if the TaxRuleDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxRuleDataRelationships{}

// TaxRuleDataRelationships struct for TaxRuleDataRelationships
type TaxRuleDataRelationships struct {
	ManualTaxCalculator *TaxRuleDataRelationshipsManualTaxCalculator `json:"manual_tax_calculator,omitempty"`
	Versions            *AddressDataRelationshipsVersions            `json:"versions,omitempty"`
}

// NewTaxRuleDataRelationships instantiates a new TaxRuleDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxRuleDataRelationships() *TaxRuleDataRelationships {
	this := TaxRuleDataRelationships{}
	return &this
}

// NewTaxRuleDataRelationshipsWithDefaults instantiates a new TaxRuleDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxRuleDataRelationshipsWithDefaults() *TaxRuleDataRelationships {
	this := TaxRuleDataRelationships{}
	return &this
}

// GetManualTaxCalculator returns the ManualTaxCalculator field value if set, zero value otherwise.
func (o *TaxRuleDataRelationships) GetManualTaxCalculator() TaxRuleDataRelationshipsManualTaxCalculator {
	if o == nil || IsNil(o.ManualTaxCalculator) {
		var ret TaxRuleDataRelationshipsManualTaxCalculator
		return ret
	}
	return *o.ManualTaxCalculator
}

// GetManualTaxCalculatorOk returns a tuple with the ManualTaxCalculator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxRuleDataRelationships) GetManualTaxCalculatorOk() (*TaxRuleDataRelationshipsManualTaxCalculator, bool) {
	if o == nil || IsNil(o.ManualTaxCalculator) {
		return nil, false
	}
	return o.ManualTaxCalculator, true
}

// HasManualTaxCalculator returns a boolean if a field has been set.
func (o *TaxRuleDataRelationships) HasManualTaxCalculator() bool {
	if o != nil && !IsNil(o.ManualTaxCalculator) {
		return true
	}

	return false
}

// SetManualTaxCalculator gets a reference to the given TaxRuleDataRelationshipsManualTaxCalculator and assigns it to the ManualTaxCalculator field.
func (o *TaxRuleDataRelationships) SetManualTaxCalculator(v TaxRuleDataRelationshipsManualTaxCalculator) {
	o.ManualTaxCalculator = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *TaxRuleDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxRuleDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *TaxRuleDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *TaxRuleDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

func (o TaxRuleDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxRuleDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ManualTaxCalculator) {
		toSerialize["manual_tax_calculator"] = o.ManualTaxCalculator
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullableTaxRuleDataRelationships struct {
	value *TaxRuleDataRelationships
	isSet bool
}

func (v NullableTaxRuleDataRelationships) Get() *TaxRuleDataRelationships {
	return v.value
}

func (v *NullableTaxRuleDataRelationships) Set(val *TaxRuleDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxRuleDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxRuleDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxRuleDataRelationships(val *TaxRuleDataRelationships) *NullableTaxRuleDataRelationships {
	return &NullableTaxRuleDataRelationships{value: val, isSet: true}
}

func (v NullableTaxRuleDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxRuleDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
