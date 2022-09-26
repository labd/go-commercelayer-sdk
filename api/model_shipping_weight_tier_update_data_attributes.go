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

// ShippingWeightTierUpdateDataAttributes struct for ShippingWeightTierUpdateDataAttributes
type ShippingWeightTierUpdateDataAttributes struct {
	// The shipping method tier's name
	Name *string `json:"name,omitempty"`
	// The tier upper limit. When 'null' it means infinity (useful to have an always matching tier).
	UpTo *float32 `json:"up_to,omitempty"`
	// The price of this shipping method tier, in cents.
	PriceAmountCents *int32 `json:"price_amount_cents,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewShippingWeightTierUpdateDataAttributes instantiates a new ShippingWeightTierUpdateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingWeightTierUpdateDataAttributes() *ShippingWeightTierUpdateDataAttributes {
	this := ShippingWeightTierUpdateDataAttributes{}
	return &this
}

// NewShippingWeightTierUpdateDataAttributesWithDefaults instantiates a new ShippingWeightTierUpdateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingWeightTierUpdateDataAttributesWithDefaults() *ShippingWeightTierUpdateDataAttributes {
	this := ShippingWeightTierUpdateDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ShippingWeightTierUpdateDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingWeightTierUpdateDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ShippingWeightTierUpdateDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ShippingWeightTierUpdateDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetUpTo returns the UpTo field value if set, zero value otherwise.
func (o *ShippingWeightTierUpdateDataAttributes) GetUpTo() float32 {
	if o == nil || o.UpTo == nil {
		var ret float32
		return ret
	}
	return *o.UpTo
}

// GetUpToOk returns a tuple with the UpTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingWeightTierUpdateDataAttributes) GetUpToOk() (*float32, bool) {
	if o == nil || o.UpTo == nil {
		return nil, false
	}
	return o.UpTo, true
}

// HasUpTo returns a boolean if a field has been set.
func (o *ShippingWeightTierUpdateDataAttributes) HasUpTo() bool {
	if o != nil && o.UpTo != nil {
		return true
	}

	return false
}

// SetUpTo gets a reference to the given float32 and assigns it to the UpTo field.
func (o *ShippingWeightTierUpdateDataAttributes) SetUpTo(v float32) {
	o.UpTo = &v
}

// GetPriceAmountCents returns the PriceAmountCents field value if set, zero value otherwise.
func (o *ShippingWeightTierUpdateDataAttributes) GetPriceAmountCents() int32 {
	if o == nil || o.PriceAmountCents == nil {
		var ret int32
		return ret
	}
	return *o.PriceAmountCents
}

// GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingWeightTierUpdateDataAttributes) GetPriceAmountCentsOk() (*int32, bool) {
	if o == nil || o.PriceAmountCents == nil {
		return nil, false
	}
	return o.PriceAmountCents, true
}

// HasPriceAmountCents returns a boolean if a field has been set.
func (o *ShippingWeightTierUpdateDataAttributes) HasPriceAmountCents() bool {
	if o != nil && o.PriceAmountCents != nil {
		return true
	}

	return false
}

// SetPriceAmountCents gets a reference to the given int32 and assigns it to the PriceAmountCents field.
func (o *ShippingWeightTierUpdateDataAttributes) SetPriceAmountCents(v int32) {
	o.PriceAmountCents = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *ShippingWeightTierUpdateDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingWeightTierUpdateDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *ShippingWeightTierUpdateDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *ShippingWeightTierUpdateDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *ShippingWeightTierUpdateDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingWeightTierUpdateDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *ShippingWeightTierUpdateDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *ShippingWeightTierUpdateDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ShippingWeightTierUpdateDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingWeightTierUpdateDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ShippingWeightTierUpdateDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ShippingWeightTierUpdateDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o ShippingWeightTierUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.UpTo != nil {
		toSerialize["up_to"] = o.UpTo
	}
	if o.PriceAmountCents != nil {
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

type NullableShippingWeightTierUpdateDataAttributes struct {
	value *ShippingWeightTierUpdateDataAttributes
	isSet bool
}

func (v NullableShippingWeightTierUpdateDataAttributes) Get() *ShippingWeightTierUpdateDataAttributes {
	return v.value
}

func (v *NullableShippingWeightTierUpdateDataAttributes) Set(val *ShippingWeightTierUpdateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingWeightTierUpdateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingWeightTierUpdateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingWeightTierUpdateDataAttributes(val *ShippingWeightTierUpdateDataAttributes) *NullableShippingWeightTierUpdateDataAttributes {
	return &NullableShippingWeightTierUpdateDataAttributes{value: val, isSet: true}
}

func (v NullableShippingWeightTierUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingWeightTierUpdateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
