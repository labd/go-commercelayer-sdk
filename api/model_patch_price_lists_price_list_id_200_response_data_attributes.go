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

// PATCHPriceListsPriceListId200ResponseDataAttributes struct for PATCHPriceListsPriceListId200ResponseDataAttributes
type PATCHPriceListsPriceListId200ResponseDataAttributes struct {
	// The price list's internal name
	Name *string `json:"name,omitempty"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode *string `json:"currency_code,omitempty"`
	// Indicates if the associated prices include taxes.
	TaxIncluded *bool `json:"tax_included,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPATCHPriceListsPriceListId200ResponseDataAttributes instantiates a new PATCHPriceListsPriceListId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHPriceListsPriceListId200ResponseDataAttributes() *PATCHPriceListsPriceListId200ResponseDataAttributes {
	this := PATCHPriceListsPriceListId200ResponseDataAttributes{}
	return &this
}

// NewPATCHPriceListsPriceListId200ResponseDataAttributesWithDefaults instantiates a new PATCHPriceListsPriceListId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHPriceListsPriceListId200ResponseDataAttributesWithDefaults() *PATCHPriceListsPriceListId200ResponseDataAttributes {
	this := PATCHPriceListsPriceListId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetTaxIncluded returns the TaxIncluded field value if set, zero value otherwise.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) GetTaxIncluded() bool {
	if o == nil || o.TaxIncluded == nil {
		var ret bool
		return ret
	}
	return *o.TaxIncluded
}

// GetTaxIncludedOk returns a tuple with the TaxIncluded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) GetTaxIncludedOk() (*bool, bool) {
	if o == nil || o.TaxIncluded == nil {
		return nil, false
	}
	return o.TaxIncluded, true
}

// HasTaxIncluded returns a boolean if a field has been set.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) HasTaxIncluded() bool {
	if o != nil && o.TaxIncluded != nil {
		return true
	}

	return false
}

// SetTaxIncluded gets a reference to the given bool and assigns it to the TaxIncluded field.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) SetTaxIncluded(v bool) {
	o.TaxIncluded = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PATCHPriceListsPriceListId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.TaxIncluded != nil {
		toSerialize["tax_included"] = o.TaxIncluded
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHPriceListsPriceListId200ResponseDataAttributes struct {
	value *PATCHPriceListsPriceListId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHPriceListsPriceListId200ResponseDataAttributes) Get() *PATCHPriceListsPriceListId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHPriceListsPriceListId200ResponseDataAttributes) Set(val *PATCHPriceListsPriceListId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHPriceListsPriceListId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHPriceListsPriceListId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHPriceListsPriceListId200ResponseDataAttributes(val *PATCHPriceListsPriceListId200ResponseDataAttributes) *NullablePATCHPriceListsPriceListId200ResponseDataAttributes {
	return &NullablePATCHPriceListsPriceListId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHPriceListsPriceListId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHPriceListsPriceListId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
