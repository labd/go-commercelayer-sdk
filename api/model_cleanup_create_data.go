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

// checks if the CleanupCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CleanupCreateData{}

// CleanupCreateData struct for CleanupCreateData
type CleanupCreateData struct {
	// The resource's type
	Type          interface{}                           `json:"type"`
	Attributes    POSTCleanups201ResponseDataAttributes `json:"attributes"`
	Relationships interface{}                           `json:"relationships,omitempty"`
}

// NewCleanupCreateData instantiates a new CleanupCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCleanupCreateData(type_ interface{}, attributes POSTCleanups201ResponseDataAttributes) *CleanupCreateData {
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
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CleanupCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CleanupCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CleanupCreateData) SetType(v interface{}) {
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

// GetRelationships returns the Relationships field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CleanupCreateData) GetRelationships() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CleanupCreateData) GetRelationshipsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return &o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CleanupCreateData) HasRelationships() bool {
	if o != nil && IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given interface{} and assigns it to the Relationships field.
func (o *CleanupCreateData) SetRelationships(v interface{}) {
	o.Relationships = v
}

func (o CleanupCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CleanupCreateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
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
