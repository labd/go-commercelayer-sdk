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

// checks if the POSTOrderAmountPromotionRules201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrderAmountPromotionRules201ResponseDataAttributes{}

// POSTOrderAmountPromotionRules201ResponseDataAttributes struct for POSTOrderAmountPromotionRules201ResponseDataAttributes
type POSTOrderAmountPromotionRules201ResponseDataAttributes struct {
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
	// Apply the promotion only when order is over this amount, in cents.
	OrderAmountCents interface{} `json:"order_amount_cents,omitempty"`
	// Send this attribute if you want to compare the specified amount with order's subtotal (excluding discounts, if any).
	UseSubtotal interface{} `json:"use_subtotal,omitempty"`
}

// NewPOSTOrderAmountPromotionRules201ResponseDataAttributes instantiates a new POSTOrderAmountPromotionRules201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrderAmountPromotionRules201ResponseDataAttributes() *POSTOrderAmountPromotionRules201ResponseDataAttributes {
	this := POSTOrderAmountPromotionRules201ResponseDataAttributes{}
	return &this
}

// NewPOSTOrderAmountPromotionRules201ResponseDataAttributesWithDefaults instantiates a new POSTOrderAmountPromotionRules201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrderAmountPromotionRules201ResponseDataAttributesWithDefaults() *POSTOrderAmountPromotionRules201ResponseDataAttributes {
	this := POSTOrderAmountPromotionRules201ResponseDataAttributes{}
	return &this
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetOrderAmountCents returns the OrderAmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) GetOrderAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.OrderAmountCents
}

// GetOrderAmountCentsOk returns a tuple with the OrderAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) GetOrderAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.OrderAmountCents) {
		return nil, false
	}
	return &o.OrderAmountCents, true
}

// HasOrderAmountCents returns a boolean if a field has been set.
func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) HasOrderAmountCents() bool {
	if o != nil && IsNil(o.OrderAmountCents) {
		return true
	}

	return false
}

// SetOrderAmountCents gets a reference to the given interface{} and assigns it to the OrderAmountCents field.
func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) SetOrderAmountCents(v interface{}) {
	o.OrderAmountCents = v
}

// GetUseSubtotal returns the UseSubtotal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) GetUseSubtotal() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UseSubtotal
}

// GetUseSubtotalOk returns a tuple with the UseSubtotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) GetUseSubtotalOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UseSubtotal) {
		return nil, false
	}
	return &o.UseSubtotal, true
}

// HasUseSubtotal returns a boolean if a field has been set.
func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) HasUseSubtotal() bool {
	if o != nil && IsNil(o.UseSubtotal) {
		return true
	}

	return false
}

// SetUseSubtotal gets a reference to the given interface{} and assigns it to the UseSubtotal field.
func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) SetUseSubtotal(v interface{}) {
	o.UseSubtotal = v
}

func (o POSTOrderAmountPromotionRules201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrderAmountPromotionRules201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.OrderAmountCents != nil {
		toSerialize["order_amount_cents"] = o.OrderAmountCents
	}
	if o.UseSubtotal != nil {
		toSerialize["use_subtotal"] = o.UseSubtotal
	}
	return toSerialize, nil
}

type NullablePOSTOrderAmountPromotionRules201ResponseDataAttributes struct {
	value *POSTOrderAmountPromotionRules201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTOrderAmountPromotionRules201ResponseDataAttributes) Get() *POSTOrderAmountPromotionRules201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTOrderAmountPromotionRules201ResponseDataAttributes) Set(val *POSTOrderAmountPromotionRules201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrderAmountPromotionRules201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrderAmountPromotionRules201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrderAmountPromotionRules201ResponseDataAttributes(val *POSTOrderAmountPromotionRules201ResponseDataAttributes) *NullablePOSTOrderAmountPromotionRules201ResponseDataAttributes {
	return &NullablePOSTOrderAmountPromotionRules201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTOrderAmountPromotionRules201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrderAmountPromotionRules201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
