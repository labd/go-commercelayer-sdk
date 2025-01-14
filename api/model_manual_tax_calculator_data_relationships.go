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

// checks if the ManualTaxCalculatorDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManualTaxCalculatorDataRelationships{}

// ManualTaxCalculatorDataRelationships struct for ManualTaxCalculatorDataRelationships
type ManualTaxCalculatorDataRelationships struct {
	Markets     *AvalaraAccountDataRelationshipsMarkets       `json:"markets,omitempty"`
	Attachments *AuthorizationDataRelationshipsAttachments    `json:"attachments,omitempty"`
	Versions    *AddressDataRelationshipsVersions             `json:"versions,omitempty"`
	TaxRules    *ManualTaxCalculatorDataRelationshipsTaxRules `json:"tax_rules,omitempty"`
}

// NewManualTaxCalculatorDataRelationships instantiates a new ManualTaxCalculatorDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualTaxCalculatorDataRelationships() *ManualTaxCalculatorDataRelationships {
	this := ManualTaxCalculatorDataRelationships{}
	return &this
}

// NewManualTaxCalculatorDataRelationshipsWithDefaults instantiates a new ManualTaxCalculatorDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualTaxCalculatorDataRelationshipsWithDefaults() *ManualTaxCalculatorDataRelationships {
	this := ManualTaxCalculatorDataRelationships{}
	return &this
}

// GetMarkets returns the Markets field value if set, zero value otherwise.
func (o *ManualTaxCalculatorDataRelationships) GetMarkets() AvalaraAccountDataRelationshipsMarkets {
	if o == nil || IsNil(o.Markets) {
		var ret AvalaraAccountDataRelationshipsMarkets
		return ret
	}
	return *o.Markets
}

// GetMarketsOk returns a tuple with the Markets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualTaxCalculatorDataRelationships) GetMarketsOk() (*AvalaraAccountDataRelationshipsMarkets, bool) {
	if o == nil || IsNil(o.Markets) {
		return nil, false
	}
	return o.Markets, true
}

// HasMarkets returns a boolean if a field has been set.
func (o *ManualTaxCalculatorDataRelationships) HasMarkets() bool {
	if o != nil && !IsNil(o.Markets) {
		return true
	}

	return false
}

// SetMarkets gets a reference to the given AvalaraAccountDataRelationshipsMarkets and assigns it to the Markets field.
func (o *ManualTaxCalculatorDataRelationships) SetMarkets(v AvalaraAccountDataRelationshipsMarkets) {
	o.Markets = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *ManualTaxCalculatorDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret AuthorizationDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualTaxCalculatorDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *ManualTaxCalculatorDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AuthorizationDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *ManualTaxCalculatorDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *ManualTaxCalculatorDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualTaxCalculatorDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *ManualTaxCalculatorDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *ManualTaxCalculatorDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

// GetTaxRules returns the TaxRules field value if set, zero value otherwise.
func (o *ManualTaxCalculatorDataRelationships) GetTaxRules() ManualTaxCalculatorDataRelationshipsTaxRules {
	if o == nil || IsNil(o.TaxRules) {
		var ret ManualTaxCalculatorDataRelationshipsTaxRules
		return ret
	}
	return *o.TaxRules
}

// GetTaxRulesOk returns a tuple with the TaxRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualTaxCalculatorDataRelationships) GetTaxRulesOk() (*ManualTaxCalculatorDataRelationshipsTaxRules, bool) {
	if o == nil || IsNil(o.TaxRules) {
		return nil, false
	}
	return o.TaxRules, true
}

// HasTaxRules returns a boolean if a field has been set.
func (o *ManualTaxCalculatorDataRelationships) HasTaxRules() bool {
	if o != nil && !IsNil(o.TaxRules) {
		return true
	}

	return false
}

// SetTaxRules gets a reference to the given ManualTaxCalculatorDataRelationshipsTaxRules and assigns it to the TaxRules field.
func (o *ManualTaxCalculatorDataRelationships) SetTaxRules(v ManualTaxCalculatorDataRelationshipsTaxRules) {
	o.TaxRules = &v
}

func (o ManualTaxCalculatorDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManualTaxCalculatorDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Markets) {
		toSerialize["markets"] = o.Markets
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	if !IsNil(o.TaxRules) {
		toSerialize["tax_rules"] = o.TaxRules
	}
	return toSerialize, nil
}

type NullableManualTaxCalculatorDataRelationships struct {
	value *ManualTaxCalculatorDataRelationships
	isSet bool
}

func (v NullableManualTaxCalculatorDataRelationships) Get() *ManualTaxCalculatorDataRelationships {
	return v.value
}

func (v *NullableManualTaxCalculatorDataRelationships) Set(val *ManualTaxCalculatorDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableManualTaxCalculatorDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableManualTaxCalculatorDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualTaxCalculatorDataRelationships(val *ManualTaxCalculatorDataRelationships) *NullableManualTaxCalculatorDataRelationships {
	return &NullableManualTaxCalculatorDataRelationships{value: val, isSet: true}
}

func (v NullableManualTaxCalculatorDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualTaxCalculatorDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
