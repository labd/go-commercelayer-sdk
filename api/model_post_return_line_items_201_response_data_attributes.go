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

// POSTReturnLineItems201ResponseDataAttributes struct for POSTReturnLineItems201ResponseDataAttributes
type POSTReturnLineItems201ResponseDataAttributes struct {
	// The line item quantity.
	Quantity int32 `json:"quantity"`
	// Set of key-value pairs that you can use to add details about return reason.
	ReturnReason map[string]interface{} `json:"return_reason,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPOSTReturnLineItems201ResponseDataAttributes instantiates a new POSTReturnLineItems201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTReturnLineItems201ResponseDataAttributes(quantity int32) *POSTReturnLineItems201ResponseDataAttributes {
	this := POSTReturnLineItems201ResponseDataAttributes{}
	this.Quantity = quantity
	return &this
}

// NewPOSTReturnLineItems201ResponseDataAttributesWithDefaults instantiates a new POSTReturnLineItems201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTReturnLineItems201ResponseDataAttributesWithDefaults() *POSTReturnLineItems201ResponseDataAttributes {
	this := POSTReturnLineItems201ResponseDataAttributes{}
	return &this
}

// GetQuantity returns the Quantity field value
func (o *POSTReturnLineItems201ResponseDataAttributes) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *POSTReturnLineItems201ResponseDataAttributes) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *POSTReturnLineItems201ResponseDataAttributes) SetQuantity(v int32) {
	o.Quantity = v
}

// GetReturnReason returns the ReturnReason field value if set, zero value otherwise.
func (o *POSTReturnLineItems201ResponseDataAttributes) GetReturnReason() map[string]interface{} {
	if o == nil || o.ReturnReason == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ReturnReason
}

// GetReturnReasonOk returns a tuple with the ReturnReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTReturnLineItems201ResponseDataAttributes) GetReturnReasonOk() (map[string]interface{}, bool) {
	if o == nil || o.ReturnReason == nil {
		return nil, false
	}
	return o.ReturnReason, true
}

// HasReturnReason returns a boolean if a field has been set.
func (o *POSTReturnLineItems201ResponseDataAttributes) HasReturnReason() bool {
	if o != nil && o.ReturnReason != nil {
		return true
	}

	return false
}

// SetReturnReason gets a reference to the given map[string]interface{} and assigns it to the ReturnReason field.
func (o *POSTReturnLineItems201ResponseDataAttributes) SetReturnReason(v map[string]interface{}) {
	o.ReturnReason = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *POSTReturnLineItems201ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTReturnLineItems201ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTReturnLineItems201ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *POSTReturnLineItems201ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *POSTReturnLineItems201ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTReturnLineItems201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTReturnLineItems201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *POSTReturnLineItems201ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *POSTReturnLineItems201ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTReturnLineItems201ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTReturnLineItems201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *POSTReturnLineItems201ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o POSTReturnLineItems201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["quantity"] = o.Quantity
	}
	if o.ReturnReason != nil {
		toSerialize["return_reason"] = o.ReturnReason
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

type NullablePOSTReturnLineItems201ResponseDataAttributes struct {
	value *POSTReturnLineItems201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTReturnLineItems201ResponseDataAttributes) Get() *POSTReturnLineItems201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTReturnLineItems201ResponseDataAttributes) Set(val *POSTReturnLineItems201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTReturnLineItems201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTReturnLineItems201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTReturnLineItems201ResponseDataAttributes(val *POSTReturnLineItems201ResponseDataAttributes) *NullablePOSTReturnLineItems201ResponseDataAttributes {
	return &NullablePOSTReturnLineItems201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTReturnLineItems201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTReturnLineItems201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
