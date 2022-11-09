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

// AvalaraAccountDataRelationships struct for AvalaraAccountDataRelationships
type AvalaraAccountDataRelationships struct {
	TaxCategories *AvalaraAccountDataRelationshipsTaxCategories `json:"tax_categories,omitempty"`
	Markets       *AvalaraAccountDataRelationshipsMarkets       `json:"markets,omitempty"`
	Attachments   *AvalaraAccountDataRelationshipsAttachments   `json:"attachments,omitempty"`
}

// NewAvalaraAccountDataRelationships instantiates a new AvalaraAccountDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvalaraAccountDataRelationships() *AvalaraAccountDataRelationships {
	this := AvalaraAccountDataRelationships{}
	return &this
}

// NewAvalaraAccountDataRelationshipsWithDefaults instantiates a new AvalaraAccountDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvalaraAccountDataRelationshipsWithDefaults() *AvalaraAccountDataRelationships {
	this := AvalaraAccountDataRelationships{}
	return &this
}

// GetTaxCategories returns the TaxCategories field value if set, zero value otherwise.
func (o *AvalaraAccountDataRelationships) GetTaxCategories() AvalaraAccountDataRelationshipsTaxCategories {
	if o == nil || o.TaxCategories == nil {
		var ret AvalaraAccountDataRelationshipsTaxCategories
		return ret
	}
	return *o.TaxCategories
}

// GetTaxCategoriesOk returns a tuple with the TaxCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvalaraAccountDataRelationships) GetTaxCategoriesOk() (*AvalaraAccountDataRelationshipsTaxCategories, bool) {
	if o == nil || o.TaxCategories == nil {
		return nil, false
	}
	return o.TaxCategories, true
}

// HasTaxCategories returns a boolean if a field has been set.
func (o *AvalaraAccountDataRelationships) HasTaxCategories() bool {
	if o != nil && o.TaxCategories != nil {
		return true
	}

	return false
}

// SetTaxCategories gets a reference to the given AvalaraAccountDataRelationshipsTaxCategories and assigns it to the TaxCategories field.
func (o *AvalaraAccountDataRelationships) SetTaxCategories(v AvalaraAccountDataRelationshipsTaxCategories) {
	o.TaxCategories = &v
}

// GetMarkets returns the Markets field value if set, zero value otherwise.
func (o *AvalaraAccountDataRelationships) GetMarkets() AvalaraAccountDataRelationshipsMarkets {
	if o == nil || o.Markets == nil {
		var ret AvalaraAccountDataRelationshipsMarkets
		return ret
	}
	return *o.Markets
}

// GetMarketsOk returns a tuple with the Markets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvalaraAccountDataRelationships) GetMarketsOk() (*AvalaraAccountDataRelationshipsMarkets, bool) {
	if o == nil || o.Markets == nil {
		return nil, false
	}
	return o.Markets, true
}

// HasMarkets returns a boolean if a field has been set.
func (o *AvalaraAccountDataRelationships) HasMarkets() bool {
	if o != nil && o.Markets != nil {
		return true
	}

	return false
}

// SetMarkets gets a reference to the given AvalaraAccountDataRelationshipsMarkets and assigns it to the Markets field.
func (o *AvalaraAccountDataRelationships) SetMarkets(v AvalaraAccountDataRelationshipsMarkets) {
	o.Markets = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *AvalaraAccountDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret AvalaraAccountDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvalaraAccountDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *AvalaraAccountDataRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *AvalaraAccountDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments) {
	o.Attachments = &v
}

func (o AvalaraAccountDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TaxCategories != nil {
		toSerialize["tax_categories"] = o.TaxCategories
	}
	if o.Markets != nil {
		toSerialize["markets"] = o.Markets
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableAvalaraAccountDataRelationships struct {
	value *AvalaraAccountDataRelationships
	isSet bool
}

func (v NullableAvalaraAccountDataRelationships) Get() *AvalaraAccountDataRelationships {
	return v.value
}

func (v *NullableAvalaraAccountDataRelationships) Set(val *AvalaraAccountDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAvalaraAccountDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAvalaraAccountDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvalaraAccountDataRelationships(val *AvalaraAccountDataRelationships) *NullableAvalaraAccountDataRelationships {
	return &NullableAvalaraAccountDataRelationships{value: val, isSet: true}
}

func (v NullableAvalaraAccountDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvalaraAccountDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
