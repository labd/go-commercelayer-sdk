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

// InventoryReturnLocationUpdateDataAttributes struct for InventoryReturnLocationUpdateDataAttributes
type InventoryReturnLocationUpdateDataAttributes struct {
	// The inventory return location priority within the associated invetory model.
	Priority *int32 `json:"priority,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewInventoryReturnLocationUpdateDataAttributes instantiates a new InventoryReturnLocationUpdateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryReturnLocationUpdateDataAttributes() *InventoryReturnLocationUpdateDataAttributes {
	this := InventoryReturnLocationUpdateDataAttributes{}
	return &this
}

// NewInventoryReturnLocationUpdateDataAttributesWithDefaults instantiates a new InventoryReturnLocationUpdateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryReturnLocationUpdateDataAttributesWithDefaults() *InventoryReturnLocationUpdateDataAttributes {
	this := InventoryReturnLocationUpdateDataAttributes{}
	return &this
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *InventoryReturnLocationUpdateDataAttributes) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryReturnLocationUpdateDataAttributes) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *InventoryReturnLocationUpdateDataAttributes) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *InventoryReturnLocationUpdateDataAttributes) SetPriority(v int32) {
	o.Priority = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *InventoryReturnLocationUpdateDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryReturnLocationUpdateDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *InventoryReturnLocationUpdateDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *InventoryReturnLocationUpdateDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *InventoryReturnLocationUpdateDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryReturnLocationUpdateDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *InventoryReturnLocationUpdateDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *InventoryReturnLocationUpdateDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *InventoryReturnLocationUpdateDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryReturnLocationUpdateDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *InventoryReturnLocationUpdateDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *InventoryReturnLocationUpdateDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o InventoryReturnLocationUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
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

type NullableInventoryReturnLocationUpdateDataAttributes struct {
	value *InventoryReturnLocationUpdateDataAttributes
	isSet bool
}

func (v NullableInventoryReturnLocationUpdateDataAttributes) Get() *InventoryReturnLocationUpdateDataAttributes {
	return v.value
}

func (v *NullableInventoryReturnLocationUpdateDataAttributes) Set(val *InventoryReturnLocationUpdateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryReturnLocationUpdateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryReturnLocationUpdateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryReturnLocationUpdateDataAttributes(val *InventoryReturnLocationUpdateDataAttributes) *NullableInventoryReturnLocationUpdateDataAttributes {
	return &NullableInventoryReturnLocationUpdateDataAttributes{value: val, isSet: true}
}

func (v NullableInventoryReturnLocationUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryReturnLocationUpdateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
