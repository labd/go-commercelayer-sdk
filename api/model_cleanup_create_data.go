/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CleanupCreateData struct for CleanupCreateData
type CleanupCreateData struct {
	// The resource's type
	Type          string                                `json:"type"`
	Attributes    POSTCleanups201ResponseDataAttributes `json:"attributes"`
	Relationships map[string]interface{}                `json:"relationships,omitempty"`
}

// NewCleanupCreateData instantiates a new CleanupCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCleanupCreateData(type_ string, attributes POSTCleanups201ResponseDataAttributes) *CleanupCreateData {
	this := CleanupCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCleanupCreateDataWithDefaults instantiates a new CleanupCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCleanupCreateDataWithDefaults() *CleanupCreateData {
	this := CleanupCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *CleanupCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CleanupCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CleanupCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CleanupCreateData) GetAttributes() POSTCleanups201ResponseDataAttributes {
	if o == nil {
		var ret POSTCleanups201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CleanupCreateData) GetAttributesOk() (*POSTCleanups201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CleanupCreateData) SetAttributes(v POSTCleanups201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CleanupCreateData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CleanupCreateData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CleanupCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *CleanupCreateData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o CleanupCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableCleanupCreateData struct {
	value *CleanupCreateData
	isSet bool
}

func (v NullableCleanupCreateData) Get() *CleanupCreateData {
	return v.value
}

func (v *NullableCleanupCreateData) Set(val *CleanupCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableCleanupCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableCleanupCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCleanupCreateData(val *CleanupCreateData) *NullableCleanupCreateData {
	return &NullableCleanupCreateData{value: val, isSet: true}
}

func (v NullableCleanupCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCleanupCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
