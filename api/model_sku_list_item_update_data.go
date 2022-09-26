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

// SkuListItemUpdateData struct for SkuListItemUpdateData
type SkuListItemUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                          `json:"id"`
	Attributes    SkuListItemCreateDataAttributes `json:"attributes"`
	Relationships map[string]interface{}          `json:"relationships,omitempty"`
}

// NewSkuListItemUpdateData instantiates a new SkuListItemUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuListItemUpdateData(type_ string, id string, attributes SkuListItemCreateDataAttributes) *SkuListItemUpdateData {
	this := SkuListItemUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewSkuListItemUpdateDataWithDefaults instantiates a new SkuListItemUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListItemUpdateDataWithDefaults() *SkuListItemUpdateData {
	this := SkuListItemUpdateData{}
	return &this
}

// GetType returns the Type field value
func (o *SkuListItemUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SkuListItemUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SkuListItemUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *SkuListItemUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SkuListItemUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SkuListItemUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *SkuListItemUpdateData) GetAttributes() SkuListItemCreateDataAttributes {
	if o == nil {
		var ret SkuListItemCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SkuListItemUpdateData) GetAttributesOk() (*SkuListItemCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SkuListItemUpdateData) SetAttributes(v SkuListItemCreateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SkuListItemUpdateData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListItemUpdateData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SkuListItemUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *SkuListItemUpdateData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o SkuListItemUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableSkuListItemUpdateData struct {
	value *SkuListItemUpdateData
	isSet bool
}

func (v NullableSkuListItemUpdateData) Get() *SkuListItemUpdateData {
	return v.value
}

func (v *NullableSkuListItemUpdateData) Set(val *SkuListItemUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuListItemUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuListItemUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuListItemUpdateData(val *SkuListItemUpdateData) *NullableSkuListItemUpdateData {
	return &NullableSkuListItemUpdateData{value: val, isSet: true}
}

func (v NullableSkuListItemUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuListItemUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
