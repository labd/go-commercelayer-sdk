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

// checks if the PATCHBundlesBundleId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHBundlesBundleId200ResponseDataAttributes{}

// PATCHBundlesBundleId200ResponseDataAttributes struct for PATCHBundlesBundleId200ResponseDataAttributes
type PATCHBundlesBundleId200ResponseDataAttributes struct {
	// The bundle code, that uniquely identifies the bundle within the market.
	Code interface{} `json:"code,omitempty"`
	// The internal name of the bundle.
	Name interface{} `json:"name,omitempty"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode interface{} `json:"currency_code,omitempty"`
	// An internal description of the bundle.
	Description interface{} `json:"description,omitempty"`
	// The URL of an image that represents the bundle.
	ImageUrl interface{} `json:"image_url,omitempty"`
	// The bundle price amount for the associated market, in cents.
	PriceAmountCents interface{} `json:"price_amount_cents,omitempty"`
	// The compared price amount, in cents. Useful to display a percentage discount.
	CompareAtAmountCents interface{} `json:"compare_at_amount_cents,omitempty"`
	// Send this attribute if you want to compute the price_amount_cents as the sum of the prices of the bundle SKUs for the market.
	ComputePriceAmount interface{} `json:"_compute_price_amount,omitempty"`
	// Send this attribute if you want to compute the compare_at_amount_cents as the sum of the prices of the bundle SKUs for the market.
	ComputeCompareAtAmount interface{} `json:"_compute_compare_at_amount,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPATCHBundlesBundleId200ResponseDataAttributes instantiates a new PATCHBundlesBundleId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHBundlesBundleId200ResponseDataAttributes() *PATCHBundlesBundleId200ResponseDataAttributes {
	this := PATCHBundlesBundleId200ResponseDataAttributes{}
	return &this
}

// NewPATCHBundlesBundleId200ResponseDataAttributesWithDefaults instantiates a new PATCHBundlesBundleId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHBundlesBundleId200ResponseDataAttributesWithDefaults() *PATCHBundlesBundleId200ResponseDataAttributes {
	this := PATCHBundlesBundleId200ResponseDataAttributes{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHBundlesBundleId200ResponseDataAttributes) GetCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHBundlesBundleId200ResponseDataAttributes) GetCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return &o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *PATCHBundlesBundleId200ResponseDataAttributes) HasCode() bool {
	if o != nil && IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given interface{} and assigns it to the Code field.
func (o *PATCHBundlesBundleId200ResponseDataAttributes) SetCode(v interface{}) {
	o.Code = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHBundlesBundleId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHBundlesBundleId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHBundlesBundleId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *PATCHBundlesBundleId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHBundlesBundleId200ResponseDataAttributes) GetCurrencyCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHBundlesBundleId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *PATCHBundlesBundleId200ResponseDataAttributes) HasCurrencyCode() bool {
	if o != nil && IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given interface{} and assigns it to the CurrencyCode field.
func (o *PATCHBundlesBundleId200ResponseDataAttributes) SetCurrencyCode(v interface{}) {
	o.CurrencyCode = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHBundlesBundleId200ResponseDataAttributes) GetDescription() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHBundlesBundleId200ResponseDataAttributes) GetDescriptionOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return &o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PATCHBundlesBundleId200ResponseDataAttributes) HasDescription() bool {
	if o != nil && IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given interface{} and assigns it to the Description field.
func (o *PATCHBundlesBundleId200ResponseDataAttributes) SetDescription(v interface{}) {
	o.Description = v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHBundlesBundleId200ResponseDataAttributes) GetImageUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHBundlesBundleId200ResponseDataAttributes) GetImageUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return &o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *PATCHBundlesBundleId200ResponseDataAttributes) HasImageUrl() bool {
	if o != nil && IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given interface{} and assigns it to the ImageUrl field.
func (o *PATCHBundlesBundleId200ResponseDataAttributes) SetImageUrl(v interface{}) {
	o.ImageUrl = v
}

// GetPriceAmountCents returns the PriceAmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHBundlesBundleId200ResponseDataAttributes) GetPriceAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PriceAmountCents
}

// GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHBundlesBundleId200ResponseDataAttributes) GetPriceAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PriceAmountCents) {
		return nil, false
	}
	return &o.PriceAmountCents, true
}

// HasPriceAmountCents returns a boolean if a field has been set.
func (o *PATCHBundlesBundleId200ResponseDataAttributes) HasPriceAmountCents() bool {
	if o != nil && IsNil(o.PriceAmountCents) {
		return true
	}

	return false
}

// SetPriceAmountCents gets a reference to the given interface{} and assigns it to the PriceAmountCents field.
func (o *PATCHBundlesBundleId200ResponseDataAttributes) SetPriceAmountCents(v interface{}) {
	o.PriceAmountCents = v
}

// GetCompareAtAmountCents returns the CompareAtAmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHBundlesBundleId200ResponseDataAttributes) GetCompareAtAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CompareAtAmountCents
}

// GetCompareAtAmountCentsOk returns a tuple with the CompareAtAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHBundlesBundleId200ResponseDataAttributes) GetCompareAtAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CompareAtAmountCents) {
		return nil, false
	}
	return &o.CompareAtAmountCents, true
}

// HasCompareAtAmountCents returns a boolean if a field has been set.
func (o *PATCHBundlesBundleId200ResponseDataAttributes) HasCompareAtAmountCents() bool {
	if o != nil && IsNil(o.CompareAtAmountCents) {
		return true
	}

	return false
}

// SetCompareAtAmountCents gets a reference to the given interface{} and assigns it to the CompareAtAmountCents field.
func (o *PATCHBundlesBundleId200ResponseDataAttributes) SetCompareAtAmountCents(v interface{}) {
	o.CompareAtAmountCents = v
}

// GetComputePriceAmount returns the ComputePriceAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHBundlesBundleId200ResponseDataAttributes) GetComputePriceAmount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ComputePriceAmount
}

// GetComputePriceAmountOk returns a tuple with the ComputePriceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHBundlesBundleId200ResponseDataAttributes) GetComputePriceAmountOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ComputePriceAmount) {
		return nil, false
	}
	return &o.ComputePriceAmount, true
}

// HasComputePriceAmount returns a boolean if a field has been set.
func (o *PATCHBundlesBundleId200ResponseDataAttributes) HasComputePriceAmount() bool {
	if o != nil && IsNil(o.ComputePriceAmount) {
		return true
	}

	return false
}

// SetComputePriceAmount gets a reference to the given interface{} and assigns it to the ComputePriceAmount field.
func (o *PATCHBundlesBundleId200ResponseDataAttributes) SetComputePriceAmount(v interface{}) {
	o.ComputePriceAmount = v
}

// GetComputeCompareAtAmount returns the ComputeCompareAtAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHBundlesBundleId200ResponseDataAttributes) GetComputeCompareAtAmount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ComputeCompareAtAmount
}

// GetComputeCompareAtAmountOk returns a tuple with the ComputeCompareAtAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHBundlesBundleId200ResponseDataAttributes) GetComputeCompareAtAmountOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ComputeCompareAtAmount) {
		return nil, false
	}
	return &o.ComputeCompareAtAmount, true
}

// HasComputeCompareAtAmount returns a boolean if a field has been set.
func (o *PATCHBundlesBundleId200ResponseDataAttributes) HasComputeCompareAtAmount() bool {
	if o != nil && IsNil(o.ComputeCompareAtAmount) {
		return true
	}

	return false
}

// SetComputeCompareAtAmount gets a reference to the given interface{} and assigns it to the ComputeCompareAtAmount field.
func (o *PATCHBundlesBundleId200ResponseDataAttributes) SetComputeCompareAtAmount(v interface{}) {
	o.ComputeCompareAtAmount = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHBundlesBundleId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHBundlesBundleId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHBundlesBundleId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHBundlesBundleId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHBundlesBundleId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHBundlesBundleId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHBundlesBundleId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHBundlesBundleId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHBundlesBundleId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHBundlesBundleId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHBundlesBundleId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHBundlesBundleId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o PATCHBundlesBundleId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHBundlesBundleId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ImageUrl != nil {
		toSerialize["image_url"] = o.ImageUrl
	}
	if o.PriceAmountCents != nil {
		toSerialize["price_amount_cents"] = o.PriceAmountCents
	}
	if o.CompareAtAmountCents != nil {
		toSerialize["compare_at_amount_cents"] = o.CompareAtAmountCents
	}
	if o.ComputePriceAmount != nil {
		toSerialize["_compute_price_amount"] = o.ComputePriceAmount
	}
	if o.ComputeCompareAtAmount != nil {
		toSerialize["_compute_compare_at_amount"] = o.ComputeCompareAtAmount
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

type NullablePATCHBundlesBundleId200ResponseDataAttributes struct {
	value *PATCHBundlesBundleId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHBundlesBundleId200ResponseDataAttributes) Get() *PATCHBundlesBundleId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHBundlesBundleId200ResponseDataAttributes) Set(val *PATCHBundlesBundleId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHBundlesBundleId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHBundlesBundleId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHBundlesBundleId200ResponseDataAttributes(val *PATCHBundlesBundleId200ResponseDataAttributes) *NullablePATCHBundlesBundleId200ResponseDataAttributes {
	return &NullablePATCHBundlesBundleId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHBundlesBundleId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHBundlesBundleId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
