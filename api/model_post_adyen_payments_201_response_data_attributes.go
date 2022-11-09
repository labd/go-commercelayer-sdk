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

// POSTAdyenPayments201ResponseDataAttributes struct for POSTAdyenPayments201ResponseDataAttributes
type POSTAdyenPayments201ResponseDataAttributes struct {
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPOSTAdyenPayments201ResponseDataAttributes instantiates a new POSTAdyenPayments201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAdyenPayments201ResponseDataAttributes() *POSTAdyenPayments201ResponseDataAttributes {
	this := POSTAdyenPayments201ResponseDataAttributes{}
	return &this
}

// NewPOSTAdyenPayments201ResponseDataAttributesWithDefaults instantiates a new POSTAdyenPayments201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAdyenPayments201ResponseDataAttributesWithDefaults() *POSTAdyenPayments201ResponseDataAttributes {
	this := POSTAdyenPayments201ResponseDataAttributes{}
	return &this
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *POSTAdyenPayments201ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAdyenPayments201ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTAdyenPayments201ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *POSTAdyenPayments201ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *POSTAdyenPayments201ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAdyenPayments201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTAdyenPayments201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *POSTAdyenPayments201ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *POSTAdyenPayments201ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAdyenPayments201ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTAdyenPayments201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *POSTAdyenPayments201ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o POSTAdyenPayments201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullablePOSTAdyenPayments201ResponseDataAttributes struct {
	value *POSTAdyenPayments201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTAdyenPayments201ResponseDataAttributes) Get() *POSTAdyenPayments201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTAdyenPayments201ResponseDataAttributes) Set(val *POSTAdyenPayments201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAdyenPayments201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAdyenPayments201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAdyenPayments201ResponseDataAttributes(val *POSTAdyenPayments201ResponseDataAttributes) *NullablePOSTAdyenPayments201ResponseDataAttributes {
	return &NullablePOSTAdyenPayments201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTAdyenPayments201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAdyenPayments201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
