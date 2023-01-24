/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CleanupData struct for CleanupData
type CleanupData struct {
	// The resource's type
	Type          string                                    `json:"type"`
	Attributes    GETCleanups200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *CleanupDataRelationships                 `json:"relationships,omitempty"`
}

// NewCleanupData instantiates a new CleanupData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCleanupData(type_ string, attributes GETCleanups200ResponseDataInnerAttributes) *CleanupData {
	this := CleanupData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCleanupDataWithDefaults instantiates a new CleanupData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCleanupDataWithDefaults() *CleanupData {
	this := CleanupData{}
	return &this
}

// GetType returns the Type field value
func (o *CleanupData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CleanupData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CleanupData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CleanupData) GetAttributes() GETCleanups200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETCleanups200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CleanupData) GetAttributesOk() (*GETCleanups200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CleanupData) SetAttributes(v GETCleanups200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CleanupData) GetRelationships() CleanupDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret CleanupDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CleanupData) GetRelationshipsOk() (*CleanupDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CleanupData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CleanupDataRelationships and assigns it to the Relationships field.
func (o *CleanupData) SetRelationships(v CleanupDataRelationships) {
	o.Relationships = &v
}

func (o CleanupData) MarshalJSON() ([]byte, error) {
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

type NullableCleanupData struct {
	value *CleanupData
	isSet bool
}

func (v NullableCleanupData) Get() *CleanupData {
	return v.value
}

func (v *NullableCleanupData) Set(val *CleanupData) {
	v.value = val
	v.isSet = true
}

func (v NullableCleanupData) IsSet() bool {
	return v.isSet
}

func (v *NullableCleanupData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCleanupData(val *CleanupData) *NullableCleanupData {
	return &NullableCleanupData{value: val, isSet: true}
}

func (v NullableCleanupData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCleanupData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}