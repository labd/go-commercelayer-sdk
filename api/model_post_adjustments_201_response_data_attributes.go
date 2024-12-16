/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.5.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTAdjustments201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTAdjustments201ResponseDataAttributes{}

// POSTAdjustments201ResponseDataAttributes struct for POSTAdjustments201ResponseDataAttributes
type POSTAdjustments201ResponseDataAttributes struct {
	// The adjustment name.
	Name interface{} `json:"name"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode interface{} `json:"currency_code"`
	// The adjustment amount, in cents.
	AmountCents interface{} `json:"amount_cents"`
	// Indicates if negative adjustment amount is distributed for tax calculation.
	DistributeDiscount interface{} `json:"distribute_discount,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTAdjustments201ResponseDataAttributes instantiates a new POSTAdjustments201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAdjustments201ResponseDataAttributes(name interface{}, currencyCode interface{}, amountCents interface{}) *POSTAdjustments201ResponseDataAttributes {
	this := POSTAdjustments201ResponseDataAttributes{}
	this.Name = name
	this.CurrencyCode = currencyCode
	this.AmountCents = amountCents
	return &this
}

// NewPOSTAdjustments201ResponseDataAttributesWithDefaults instantiates a new POSTAdjustments201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAdjustments201ResponseDataAttributesWithDefaults() *POSTAdjustments201ResponseDataAttributes {
	this := POSTAdjustments201ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTAdjustments201ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAdjustments201ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *POSTAdjustments201ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetCurrencyCode returns the CurrencyCode field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTAdjustments201ResponseDataAttributes) GetCurrencyCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAdjustments201ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *POSTAdjustments201ResponseDataAttributes) SetCurrencyCode(v interface{}) {
	o.CurrencyCode = v
}

// GetAmountCents returns the AmountCents field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTAdjustments201ResponseDataAttributes) GetAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.AmountCents
}

// GetAmountCentsOk returns a tuple with the AmountCents field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAdjustments201ResponseDataAttributes) GetAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AmountCents) {
		return nil, false
	}
	return &o.AmountCents, true
}

// SetAmountCents sets field value
func (o *POSTAdjustments201ResponseDataAttributes) SetAmountCents(v interface{}) {
	o.AmountCents = v
}

// GetDistributeDiscount returns the DistributeDiscount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAdjustments201ResponseDataAttributes) GetDistributeDiscount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DistributeDiscount
}

// GetDistributeDiscountOk returns a tuple with the DistributeDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAdjustments201ResponseDataAttributes) GetDistributeDiscountOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DistributeDiscount) {
		return nil, false
	}
	return &o.DistributeDiscount, true
}

// HasDistributeDiscount returns a boolean if a field has been set.
func (o *POSTAdjustments201ResponseDataAttributes) HasDistributeDiscount() bool {
	if o != nil && IsNil(o.DistributeDiscount) {
		return true
	}

	return false
}

// SetDistributeDiscount gets a reference to the given interface{} and assigns it to the DistributeDiscount field.
func (o *POSTAdjustments201ResponseDataAttributes) SetDistributeDiscount(v interface{}) {
	o.DistributeDiscount = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAdjustments201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAdjustments201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTAdjustments201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTAdjustments201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAdjustments201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAdjustments201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTAdjustments201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTAdjustments201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAdjustments201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAdjustments201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTAdjustments201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTAdjustments201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTAdjustments201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTAdjustments201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.AmountCents != nil {
		toSerialize["amount_cents"] = o.AmountCents
	}
	if o.DistributeDiscount != nil {
		toSerialize["distribute_discount"] = o.DistributeDiscount
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
	return toSerialize, nil
}

type NullablePOSTAdjustments201ResponseDataAttributes struct {
	value *POSTAdjustments201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTAdjustments201ResponseDataAttributes) Get() *POSTAdjustments201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTAdjustments201ResponseDataAttributes) Set(val *POSTAdjustments201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAdjustments201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAdjustments201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAdjustments201ResponseDataAttributes(val *POSTAdjustments201ResponseDataAttributes) *NullablePOSTAdjustments201ResponseDataAttributes {
	return &NullablePOSTAdjustments201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTAdjustments201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAdjustments201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
