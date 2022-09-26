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

// PriceVolumeTierCreateDataAttributes struct for PriceVolumeTierCreateDataAttributes
type PriceVolumeTierCreateDataAttributes struct {
	// The price tier's name
	Name string `json:"name"`
	// The tier upper limit. When 'null' it means infinity (useful to have an always matching tier).
	UpTo *float32 `json:"up_to,omitempty"`
	// The price of this price tier, in cents.
	PriceAmountCents int32 `json:"price_amount_cents"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPriceVolumeTierCreateDataAttributes instantiates a new PriceVolumeTierCreateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceVolumeTierCreateDataAttributes(name string, priceAmountCents int32) *PriceVolumeTierCreateDataAttributes {
	this := PriceVolumeTierCreateDataAttributes{}
	this.Name = name
	this.PriceAmountCents = priceAmountCents
	return &this
}

// NewPriceVolumeTierCreateDataAttributesWithDefaults instantiates a new PriceVolumeTierCreateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceVolumeTierCreateDataAttributesWithDefaults() *PriceVolumeTierCreateDataAttributes {
	this := PriceVolumeTierCreateDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *PriceVolumeTierCreateDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PriceVolumeTierCreateDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PriceVolumeTierCreateDataAttributes) SetName(v string) {
	o.Name = v
}

// GetUpTo returns the UpTo field value if set, zero value otherwise.
func (o *PriceVolumeTierCreateDataAttributes) GetUpTo() float32 {
	if o == nil || o.UpTo == nil {
		var ret float32
		return ret
	}
	return *o.UpTo
}

// GetUpToOk returns a tuple with the UpTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceVolumeTierCreateDataAttributes) GetUpToOk() (*float32, bool) {
	if o == nil || o.UpTo == nil {
		return nil, false
	}
	return o.UpTo, true
}

// HasUpTo returns a boolean if a field has been set.
func (o *PriceVolumeTierCreateDataAttributes) HasUpTo() bool {
	if o != nil && o.UpTo != nil {
		return true
	}

	return false
}

// SetUpTo gets a reference to the given float32 and assigns it to the UpTo field.
func (o *PriceVolumeTierCreateDataAttributes) SetUpTo(v float32) {
	o.UpTo = &v
}

// GetPriceAmountCents returns the PriceAmountCents field value
func (o *PriceVolumeTierCreateDataAttributes) GetPriceAmountCents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PriceAmountCents
}

// GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field value
// and a boolean to check if the value has been set.
func (o *PriceVolumeTierCreateDataAttributes) GetPriceAmountCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceAmountCents, true
}

// SetPriceAmountCents sets field value
func (o *PriceVolumeTierCreateDataAttributes) SetPriceAmountCents(v int32) {
	o.PriceAmountCents = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PriceVolumeTierCreateDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceVolumeTierCreateDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PriceVolumeTierCreateDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PriceVolumeTierCreateDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PriceVolumeTierCreateDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceVolumeTierCreateDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PriceVolumeTierCreateDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PriceVolumeTierCreateDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PriceVolumeTierCreateDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceVolumeTierCreateDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PriceVolumeTierCreateDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PriceVolumeTierCreateDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PriceVolumeTierCreateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.UpTo != nil {
		toSerialize["up_to"] = o.UpTo
	}
	if true {
		toSerialize["price_amount_cents"] = o.PriceAmountCents
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

type NullablePriceVolumeTierCreateDataAttributes struct {
	value *PriceVolumeTierCreateDataAttributes
	isSet bool
}

func (v NullablePriceVolumeTierCreateDataAttributes) Get() *PriceVolumeTierCreateDataAttributes {
	return v.value
}

func (v *NullablePriceVolumeTierCreateDataAttributes) Set(val *PriceVolumeTierCreateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceVolumeTierCreateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceVolumeTierCreateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceVolumeTierCreateDataAttributes(val *PriceVolumeTierCreateDataAttributes) *NullablePriceVolumeTierCreateDataAttributes {
	return &NullablePriceVolumeTierCreateDataAttributes{value: val, isSet: true}
}

func (v NullablePriceVolumeTierCreateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceVolumeTierCreateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
