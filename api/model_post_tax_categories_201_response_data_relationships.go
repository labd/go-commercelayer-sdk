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

// checks if the POSTTaxCategories201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTTaxCategories201ResponseDataRelationships{}

// POSTTaxCategories201ResponseDataRelationships struct for POSTTaxCategories201ResponseDataRelationships
type POSTTaxCategories201ResponseDataRelationships struct {
	Sku           *POSTInStockSubscriptions201ResponseDataRelationshipsSku                 `json:"sku,omitempty"`
	TaxCalculator *POSTMarkets201ResponseDataRelationshipsTaxCalculator                    `json:"tax_calculator,omitempty"`
	Attachments   *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments `json:"attachments,omitempty"`
	Versions      *POSTAddresses201ResponseDataRelationshipsVersions                       `json:"versions,omitempty"`
}

// NewPOSTTaxCategories201ResponseDataRelationships instantiates a new POSTTaxCategories201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTTaxCategories201ResponseDataRelationships() *POSTTaxCategories201ResponseDataRelationships {
	this := POSTTaxCategories201ResponseDataRelationships{}
	return &this
}

// NewPOSTTaxCategories201ResponseDataRelationshipsWithDefaults instantiates a new POSTTaxCategories201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTTaxCategories201ResponseDataRelationshipsWithDefaults() *POSTTaxCategories201ResponseDataRelationships {
	this := POSTTaxCategories201ResponseDataRelationships{}
	return &this
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *POSTTaxCategories201ResponseDataRelationships) GetSku() POSTInStockSubscriptions201ResponseDataRelationshipsSku {
	if o == nil || IsNil(o.Sku) {
		var ret POSTInStockSubscriptions201ResponseDataRelationshipsSku
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTTaxCategories201ResponseDataRelationships) GetSkuOk() (*POSTInStockSubscriptions201ResponseDataRelationshipsSku, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *POSTTaxCategories201ResponseDataRelationships) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given POSTInStockSubscriptions201ResponseDataRelationshipsSku and assigns it to the Sku field.
func (o *POSTTaxCategories201ResponseDataRelationships) SetSku(v POSTInStockSubscriptions201ResponseDataRelationshipsSku) {
	o.Sku = &v
}

// GetTaxCalculator returns the TaxCalculator field value if set, zero value otherwise.
func (o *POSTTaxCategories201ResponseDataRelationships) GetTaxCalculator() POSTMarkets201ResponseDataRelationshipsTaxCalculator {
	if o == nil || IsNil(o.TaxCalculator) {
		var ret POSTMarkets201ResponseDataRelationshipsTaxCalculator
		return ret
	}
	return *o.TaxCalculator
}

// GetTaxCalculatorOk returns a tuple with the TaxCalculator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTTaxCategories201ResponseDataRelationships) GetTaxCalculatorOk() (*POSTMarkets201ResponseDataRelationshipsTaxCalculator, bool) {
	if o == nil || IsNil(o.TaxCalculator) {
		return nil, false
	}
	return o.TaxCalculator, true
}

// HasTaxCalculator returns a boolean if a field has been set.
func (o *POSTTaxCategories201ResponseDataRelationships) HasTaxCalculator() bool {
	if o != nil && !IsNil(o.TaxCalculator) {
		return true
	}

	return false
}

// SetTaxCalculator gets a reference to the given POSTMarkets201ResponseDataRelationshipsTaxCalculator and assigns it to the TaxCalculator field.
func (o *POSTTaxCategories201ResponseDataRelationships) SetTaxCalculator(v POSTMarkets201ResponseDataRelationshipsTaxCalculator) {
	o.TaxCalculator = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *POSTTaxCategories201ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTTaxCategories201ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *POSTTaxCategories201ResponseDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *POSTTaxCategories201ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *POSTTaxCategories201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTTaxCategories201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *POSTTaxCategories201ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *POSTTaxCategories201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

func (o POSTTaxCategories201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTTaxCategories201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.TaxCalculator) {
		toSerialize["tax_calculator"] = o.TaxCalculator
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullablePOSTTaxCategories201ResponseDataRelationships struct {
	value *POSTTaxCategories201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTTaxCategories201ResponseDataRelationships) Get() *POSTTaxCategories201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTTaxCategories201ResponseDataRelationships) Set(val *POSTTaxCategories201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTTaxCategories201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTTaxCategories201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTTaxCategories201ResponseDataRelationships(val *POSTTaxCategories201ResponseDataRelationships) *NullablePOSTTaxCategories201ResponseDataRelationships {
	return &NullablePOSTTaxCategories201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTTaxCategories201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTTaxCategories201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
