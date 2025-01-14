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

// checks if the POSTManualTaxCalculators201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTManualTaxCalculators201ResponseDataRelationships{}

// POSTManualTaxCalculators201ResponseDataRelationships struct for POSTManualTaxCalculators201ResponseDataRelationships
type POSTManualTaxCalculators201ResponseDataRelationships struct {
	Markets     *POSTAvalaraAccounts201ResponseDataRelationshipsMarkets                  `json:"markets,omitempty"`
	Attachments *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments `json:"attachments,omitempty"`
	Versions    *POSTAddresses201ResponseDataRelationshipsVersions                       `json:"versions,omitempty"`
	TaxRules    *POSTManualTaxCalculators201ResponseDataRelationshipsTaxRules            `json:"tax_rules,omitempty"`
}

// NewPOSTManualTaxCalculators201ResponseDataRelationships instantiates a new POSTManualTaxCalculators201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTManualTaxCalculators201ResponseDataRelationships() *POSTManualTaxCalculators201ResponseDataRelationships {
	this := POSTManualTaxCalculators201ResponseDataRelationships{}
	return &this
}

// NewPOSTManualTaxCalculators201ResponseDataRelationshipsWithDefaults instantiates a new POSTManualTaxCalculators201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTManualTaxCalculators201ResponseDataRelationshipsWithDefaults() *POSTManualTaxCalculators201ResponseDataRelationships {
	this := POSTManualTaxCalculators201ResponseDataRelationships{}
	return &this
}

// GetMarkets returns the Markets field value if set, zero value otherwise.
func (o *POSTManualTaxCalculators201ResponseDataRelationships) GetMarkets() POSTAvalaraAccounts201ResponseDataRelationshipsMarkets {
	if o == nil || IsNil(o.Markets) {
		var ret POSTAvalaraAccounts201ResponseDataRelationshipsMarkets
		return ret
	}
	return *o.Markets
}

// GetMarketsOk returns a tuple with the Markets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTManualTaxCalculators201ResponseDataRelationships) GetMarketsOk() (*POSTAvalaraAccounts201ResponseDataRelationshipsMarkets, bool) {
	if o == nil || IsNil(o.Markets) {
		return nil, false
	}
	return o.Markets, true
}

// HasMarkets returns a boolean if a field has been set.
func (o *POSTManualTaxCalculators201ResponseDataRelationships) HasMarkets() bool {
	if o != nil && !IsNil(o.Markets) {
		return true
	}

	return false
}

// SetMarkets gets a reference to the given POSTAvalaraAccounts201ResponseDataRelationshipsMarkets and assigns it to the Markets field.
func (o *POSTManualTaxCalculators201ResponseDataRelationships) SetMarkets(v POSTAvalaraAccounts201ResponseDataRelationshipsMarkets) {
	o.Markets = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *POSTManualTaxCalculators201ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTManualTaxCalculators201ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *POSTManualTaxCalculators201ResponseDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *POSTManualTaxCalculators201ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *POSTManualTaxCalculators201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTManualTaxCalculators201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *POSTManualTaxCalculators201ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *POSTManualTaxCalculators201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

// GetTaxRules returns the TaxRules field value if set, zero value otherwise.
func (o *POSTManualTaxCalculators201ResponseDataRelationships) GetTaxRules() POSTManualTaxCalculators201ResponseDataRelationshipsTaxRules {
	if o == nil || IsNil(o.TaxRules) {
		var ret POSTManualTaxCalculators201ResponseDataRelationshipsTaxRules
		return ret
	}
	return *o.TaxRules
}

// GetTaxRulesOk returns a tuple with the TaxRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTManualTaxCalculators201ResponseDataRelationships) GetTaxRulesOk() (*POSTManualTaxCalculators201ResponseDataRelationshipsTaxRules, bool) {
	if o == nil || IsNil(o.TaxRules) {
		return nil, false
	}
	return o.TaxRules, true
}

// HasTaxRules returns a boolean if a field has been set.
func (o *POSTManualTaxCalculators201ResponseDataRelationships) HasTaxRules() bool {
	if o != nil && !IsNil(o.TaxRules) {
		return true
	}

	return false
}

// SetTaxRules gets a reference to the given POSTManualTaxCalculators201ResponseDataRelationshipsTaxRules and assigns it to the TaxRules field.
func (o *POSTManualTaxCalculators201ResponseDataRelationships) SetTaxRules(v POSTManualTaxCalculators201ResponseDataRelationshipsTaxRules) {
	o.TaxRules = &v
}

func (o POSTManualTaxCalculators201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTManualTaxCalculators201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
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

type NullablePOSTManualTaxCalculators201ResponseDataRelationships struct {
	value *POSTManualTaxCalculators201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTManualTaxCalculators201ResponseDataRelationships) Get() *POSTManualTaxCalculators201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTManualTaxCalculators201ResponseDataRelationships) Set(val *POSTManualTaxCalculators201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTManualTaxCalculators201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTManualTaxCalculators201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTManualTaxCalculators201ResponseDataRelationships(val *POSTManualTaxCalculators201ResponseDataRelationships) *NullablePOSTManualTaxCalculators201ResponseDataRelationships {
	return &NullablePOSTManualTaxCalculators201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTManualTaxCalculators201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTManualTaxCalculators201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
