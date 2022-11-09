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

// ShippingCategoryCreateData struct for ShippingCategoryCreateData
type ShippingCategoryCreateData struct {
	// The resource's type
	Type          string                                          `json:"type"`
	Attributes    POSTShippingCategories201ResponseDataAttributes `json:"attributes"`
	Relationships map[string]interface{}                          `json:"relationships,omitempty"`
}

// NewShippingCategoryCreateData instantiates a new ShippingCategoryCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingCategoryCreateData(type_ string, attributes POSTShippingCategories201ResponseDataAttributes) *ShippingCategoryCreateData {
	this := ShippingCategoryCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewShippingCategoryCreateDataWithDefaults instantiates a new ShippingCategoryCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingCategoryCreateDataWithDefaults() *ShippingCategoryCreateData {
	this := ShippingCategoryCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *ShippingCategoryCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ShippingCategoryCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ShippingCategoryCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ShippingCategoryCreateData) GetAttributes() POSTShippingCategories201ResponseDataAttributes {
	if o == nil {
		var ret POSTShippingCategories201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ShippingCategoryCreateData) GetAttributesOk() (*POSTShippingCategories201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ShippingCategoryCreateData) SetAttributes(v POSTShippingCategories201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ShippingCategoryCreateData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingCategoryCreateData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ShippingCategoryCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *ShippingCategoryCreateData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o ShippingCategoryCreateData) MarshalJSON() ([]byte, error) {
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

type NullableShippingCategoryCreateData struct {
	value *ShippingCategoryCreateData
	isSet bool
}

func (v NullableShippingCategoryCreateData) Get() *ShippingCategoryCreateData {
	return v.value
}

func (v *NullableShippingCategoryCreateData) Set(val *ShippingCategoryCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingCategoryCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingCategoryCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingCategoryCreateData(val *ShippingCategoryCreateData) *NullableShippingCategoryCreateData {
	return &NullableShippingCategoryCreateData{value: val, isSet: true}
}

func (v NullableShippingCategoryCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingCategoryCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
