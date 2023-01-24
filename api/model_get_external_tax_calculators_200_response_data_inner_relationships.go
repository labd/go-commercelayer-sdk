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

// GETExternalTaxCalculators200ResponseDataInnerRelationships struct for GETExternalTaxCalculators200ResponseDataInnerRelationships
type GETExternalTaxCalculators200ResponseDataInnerRelationships struct {
	Markets     *GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets     `json:"markets,omitempty"`
	Attachments *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments `json:"attachments,omitempty"`
}

// NewGETExternalTaxCalculators200ResponseDataInnerRelationships instantiates a new GETExternalTaxCalculators200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETExternalTaxCalculators200ResponseDataInnerRelationships() *GETExternalTaxCalculators200ResponseDataInnerRelationships {
	this := GETExternalTaxCalculators200ResponseDataInnerRelationships{}
	return &this
}

// NewGETExternalTaxCalculators200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETExternalTaxCalculators200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETExternalTaxCalculators200ResponseDataInnerRelationshipsWithDefaults() *GETExternalTaxCalculators200ResponseDataInnerRelationships {
	this := GETExternalTaxCalculators200ResponseDataInnerRelationships{}
	return &this
}

// GetMarkets returns the Markets field value if set, zero value otherwise.
func (o *GETExternalTaxCalculators200ResponseDataInnerRelationships) GetMarkets() GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets {
	if o == nil || o.Markets == nil {
		var ret GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets
		return ret
	}
	return *o.Markets
}

// GetMarketsOk returns a tuple with the Markets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalTaxCalculators200ResponseDataInnerRelationships) GetMarketsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets, bool) {
	if o == nil || o.Markets == nil {
		return nil, false
	}
	return o.Markets, true
}

// HasMarkets returns a boolean if a field has been set.
func (o *GETExternalTaxCalculators200ResponseDataInnerRelationships) HasMarkets() bool {
	if o != nil && o.Markets != nil {
		return true
	}

	return false
}

// SetMarkets gets a reference to the given GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets and assigns it to the Markets field.
func (o *GETExternalTaxCalculators200ResponseDataInnerRelationships) SetMarkets(v GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets) {
	o.Markets = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *GETExternalTaxCalculators200ResponseDataInnerRelationships) GetAttachments() GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalTaxCalculators200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *GETExternalTaxCalculators200ResponseDataInnerRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments and assigns it to the Attachments field.
func (o *GETExternalTaxCalculators200ResponseDataInnerRelationships) SetAttachments(v GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) {
	o.Attachments = &v
}

func (o GETExternalTaxCalculators200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Markets != nil {
		toSerialize["markets"] = o.Markets
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableGETExternalTaxCalculators200ResponseDataInnerRelationships struct {
	value *GETExternalTaxCalculators200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETExternalTaxCalculators200ResponseDataInnerRelationships) Get() *GETExternalTaxCalculators200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETExternalTaxCalculators200ResponseDataInnerRelationships) Set(val *GETExternalTaxCalculators200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETExternalTaxCalculators200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETExternalTaxCalculators200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETExternalTaxCalculators200ResponseDataInnerRelationships(val *GETExternalTaxCalculators200ResponseDataInnerRelationships) *NullableGETExternalTaxCalculators200ResponseDataInnerRelationships {
	return &NullableGETExternalTaxCalculators200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETExternalTaxCalculators200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETExternalTaxCalculators200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}