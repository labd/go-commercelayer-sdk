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

// checks if the PATCHPriceListsPriceListId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHPriceListsPriceListId200ResponseDataAttributes{}

// PATCHPriceListsPriceListId200ResponseDataAttributes struct for PATCHPriceListsPriceListId200ResponseDataAttributes
type PATCHPriceListsPriceListId200ResponseDataAttributes struct {
	// The price list's internal name.
	Name interface{} `json:"name,omitempty"`
	// A string that you can use to identify the price list (must be unique within the environment).
	Code interface{} `json:"code,omitempty"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode interface{} `json:"currency_code,omitempty"`
	// Indicates if the associated prices include taxes.
	TaxIncluded interface{} `json:"tax_included,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// The rules (using Rules Engine) to be applied.
	Rules interface{} `json:"rules,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
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

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) GetCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) GetCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return &o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) HasCode() bool {
	if o != nil && IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given interface{} and assigns it to the Code field.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) SetCode(v interface{}) {
	o.Code = v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) GetCurrencyCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) HasCurrencyCode() bool {
	if o != nil && IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given interface{} and assigns it to the CurrencyCode field.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) SetCurrencyCode(v interface{}) {
	o.CurrencyCode = v
}

// GetTaxIncluded returns the TaxIncluded field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) GetTaxIncluded() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TaxIncluded
}

// GetTaxIncludedOk returns a tuple with the TaxIncluded field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) GetTaxIncludedOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TaxIncluded) {
		return nil, false
	}
	return &o.TaxIncluded, true
}

// HasTaxIncluded returns a boolean if a field has been set.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) HasTaxIncluded() bool {
	if o != nil && IsNil(o.TaxIncluded) {
		return true
	}

	return false
}

// SetTaxIncluded gets a reference to the given interface{} and assigns it to the TaxIncluded field.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) SetTaxIncluded(v interface{}) {
	o.TaxIncluded = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetRules returns the Rules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) GetRules() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) GetRulesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return &o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) HasRules() bool {
	if o != nil && IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given interface{} and assigns it to the Rules field.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) SetRules(v interface{}) {
	o.Rules = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHPriceListsPriceListId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o PATCHPriceListsPriceListId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHPriceListsPriceListId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
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
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
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
