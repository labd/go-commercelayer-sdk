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

// SkuData struct for SkuData
type SkuData struct {
	// The resource's type
	Type          string                `json:"type"`
	Attributes    SkuDataAttributes     `json:"attributes"`
	Relationships *SkuDataRelationships `json:"relationships,omitempty"`
}

// NewSkuData instantiates a new SkuData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuData(type_ string, attributes SkuDataAttributes) *SkuData {
	this := SkuData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewSkuDataWithDefaults instantiates a new SkuData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuDataWithDefaults() *SkuData {
	this := SkuData{}
	return &this
}

// GetType returns the Type field value
func (o *SkuData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SkuData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SkuData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *SkuData) GetAttributes() SkuDataAttributes {
	if o == nil {
		var ret SkuDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SkuData) GetAttributesOk() (*SkuDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SkuData) SetAttributes(v SkuDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SkuData) GetRelationships() SkuDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret SkuDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuData) GetRelationshipsOk() (*SkuDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SkuData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given SkuDataRelationships and assigns it to the Relationships field.
func (o *SkuData) SetRelationships(v SkuDataRelationships) {
	o.Relationships = &v
}

func (o SkuData) MarshalJSON() ([]byte, error) {
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

type NullableSkuData struct {
	value *SkuData
	isSet bool
}

func (v NullableSkuData) Get() *SkuData {
	return v.value
}

func (v *NullableSkuData) Set(val *SkuData) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuData) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuData(val *SkuData) *NullableSkuData {
	return &NullableSkuData{value: val, isSet: true}
}

func (v NullableSkuData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
