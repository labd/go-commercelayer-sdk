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

// ExportData struct for ExportData
type ExportData struct {
	// The resource's type
	Type          string                   `json:"type"`
	Attributes    ExportDataAttributes     `json:"attributes"`
	Relationships *ExportDataRelationships `json:"relationships,omitempty"`
}

// NewExportData instantiates a new ExportData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportData(type_ string, attributes ExportDataAttributes) *ExportData {
	this := ExportData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewExportDataWithDefaults instantiates a new ExportData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportDataWithDefaults() *ExportData {
	this := ExportData{}
	return &this
}

// GetType returns the Type field value
func (o *ExportData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExportData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExportData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ExportData) GetAttributes() ExportDataAttributes {
	if o == nil {
		var ret ExportDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ExportData) GetAttributesOk() (*ExportDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ExportData) SetAttributes(v ExportDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ExportData) GetRelationships() ExportDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ExportDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportData) GetRelationshipsOk() (*ExportDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ExportData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ExportDataRelationships and assigns it to the Relationships field.
func (o *ExportData) SetRelationships(v ExportDataRelationships) {
	o.Relationships = &v
}

func (o ExportData) MarshalJSON() ([]byte, error) {
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

type NullableExportData struct {
	value *ExportData
	isSet bool
}

func (v NullableExportData) Get() *ExportData {
	return v.value
}

func (v *NullableExportData) Set(val *ExportData) {
	v.value = val
	v.isSet = true
}

func (v NullableExportData) IsSet() bool {
	return v.isSet
}

func (v *NullableExportData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportData(val *ExportData) *NullableExportData {
	return &NullableExportData{value: val, isSet: true}
}

func (v NullableExportData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
