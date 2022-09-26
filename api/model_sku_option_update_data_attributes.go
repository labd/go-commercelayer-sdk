/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SkuOptionUpdateDataAttributes struct for SkuOptionUpdateDataAttributes
type SkuOptionUpdateDataAttributes struct {
	// The SKU option's internal name
	Name *string `json:"name,omitempty"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode *string `json:"currency_code,omitempty"`
	// An internal description of the SKU option.
	Description *string `json:"description,omitempty"`
	// The price of this shipping method, in cents.
	PriceAmountCents *int32 `json:"price_amount_cents,omitempty"`
	// The delay time (in hours) that should be added to the delivery lead time when this option is purchased.
	DelayHours *int32 `json:"delay_hours,omitempty"`
	// The regex that will be evaluated to match the SKU codes.
	SkuCodeRegex *string `json:"sku_code_regex,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewSkuOptionUpdateDataAttributes instantiates a new SkuOptionUpdateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuOptionUpdateDataAttributes() *SkuOptionUpdateDataAttributes {
	this := SkuOptionUpdateDataAttributes{}
	return &this
}

// NewSkuOptionUpdateDataAttributesWithDefaults instantiates a new SkuOptionUpdateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuOptionUpdateDataAttributesWithDefaults() *SkuOptionUpdateDataAttributes {
	this := SkuOptionUpdateDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SkuOptionUpdateDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuOptionUpdateDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SkuOptionUpdateDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SkuOptionUpdateDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *SkuOptionUpdateDataAttributes) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuOptionUpdateDataAttributes) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *SkuOptionUpdateDataAttributes) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *SkuOptionUpdateDataAttributes) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SkuOptionUpdateDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuOptionUpdateDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SkuOptionUpdateDataAttributes) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SkuOptionUpdateDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetPriceAmountCents returns the PriceAmountCents field value if set, zero value otherwise.
func (o *SkuOptionUpdateDataAttributes) GetPriceAmountCents() int32 {
	if o == nil || o.PriceAmountCents == nil {
		var ret int32
		return ret
	}
	return *o.PriceAmountCents
}

// GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuOptionUpdateDataAttributes) GetPriceAmountCentsOk() (*int32, bool) {
	if o == nil || o.PriceAmountCents == nil {
		return nil, false
	}
	return o.PriceAmountCents, true
}

// HasPriceAmountCents returns a boolean if a field has been set.
func (o *SkuOptionUpdateDataAttributes) HasPriceAmountCents() bool {
	if o != nil && o.PriceAmountCents != nil {
		return true
	}

	return false
}

// SetPriceAmountCents gets a reference to the given int32 and assigns it to the PriceAmountCents field.
func (o *SkuOptionUpdateDataAttributes) SetPriceAmountCents(v int32) {
	o.PriceAmountCents = &v
}

// GetDelayHours returns the DelayHours field value if set, zero value otherwise.
func (o *SkuOptionUpdateDataAttributes) GetDelayHours() int32 {
	if o == nil || o.DelayHours == nil {
		var ret int32
		return ret
	}
	return *o.DelayHours
}

// GetDelayHoursOk returns a tuple with the DelayHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuOptionUpdateDataAttributes) GetDelayHoursOk() (*int32, bool) {
	if o == nil || o.DelayHours == nil {
		return nil, false
	}
	return o.DelayHours, true
}

// HasDelayHours returns a boolean if a field has been set.
func (o *SkuOptionUpdateDataAttributes) HasDelayHours() bool {
	if o != nil && o.DelayHours != nil {
		return true
	}

	return false
}

// SetDelayHours gets a reference to the given int32 and assigns it to the DelayHours field.
func (o *SkuOptionUpdateDataAttributes) SetDelayHours(v int32) {
	o.DelayHours = &v
}

// GetSkuCodeRegex returns the SkuCodeRegex field value if set, zero value otherwise.
func (o *SkuOptionUpdateDataAttributes) GetSkuCodeRegex() string {
	if o == nil || o.SkuCodeRegex == nil {
		var ret string
		return ret
	}
	return *o.SkuCodeRegex
}

// GetSkuCodeRegexOk returns a tuple with the SkuCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuOptionUpdateDataAttributes) GetSkuCodeRegexOk() (*string, bool) {
	if o == nil || o.SkuCodeRegex == nil {
		return nil, false
	}
	return o.SkuCodeRegex, true
}

// HasSkuCodeRegex returns a boolean if a field has been set.
func (o *SkuOptionUpdateDataAttributes) HasSkuCodeRegex() bool {
	if o != nil && o.SkuCodeRegex != nil {
		return true
	}

	return false
}

// SetSkuCodeRegex gets a reference to the given string and assigns it to the SkuCodeRegex field.
func (o *SkuOptionUpdateDataAttributes) SetSkuCodeRegex(v string) {
	o.SkuCodeRegex = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *SkuOptionUpdateDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuOptionUpdateDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *SkuOptionUpdateDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *SkuOptionUpdateDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *SkuOptionUpdateDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuOptionUpdateDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *SkuOptionUpdateDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *SkuOptionUpdateDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SkuOptionUpdateDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuOptionUpdateDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SkuOptionUpdateDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *SkuOptionUpdateDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o SkuOptionUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.PriceAmountCents != nil {
		toSerialize["price_amount_cents"] = o.PriceAmountCents
	}
	if o.DelayHours != nil {
		toSerialize["delay_hours"] = o.DelayHours
	}
	if o.SkuCodeRegex != nil {
		toSerialize["sku_code_regex"] = o.SkuCodeRegex
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

type NullableSkuOptionUpdateDataAttributes struct {
	value *SkuOptionUpdateDataAttributes
	isSet bool
}

func (v NullableSkuOptionUpdateDataAttributes) Get() *SkuOptionUpdateDataAttributes {
	return v.value
}

func (v *NullableSkuOptionUpdateDataAttributes) Set(val *SkuOptionUpdateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuOptionUpdateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuOptionUpdateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuOptionUpdateDataAttributes(val *SkuOptionUpdateDataAttributes) *NullableSkuOptionUpdateDataAttributes {
	return &NullableSkuOptionUpdateDataAttributes{value: val, isSet: true}
}

func (v NullableSkuOptionUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuOptionUpdateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
