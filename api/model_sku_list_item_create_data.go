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

// checks if the SkuListItemCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkuListItemCreateData{}

// SkuListItemCreateData struct for SkuListItemCreateData
type SkuListItemCreateData struct {
	// The resource's type
	Type          interface{}                               `json:"type"`
	Attributes    POSTSkuListItems201ResponseDataAttributes `json:"attributes"`
	Relationships *SkuListItemCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewSkuListItemCreateData instantiates a new SkuListItemCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuListItemCreateData(type_ interface{}, attributes POSTSkuListItems201ResponseDataAttributes) *SkuListItemCreateData {
	this := SkuListItemCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewSkuListItemCreateDataWithDefaults instantiates a new SkuListItemCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListItemCreateDataWithDefaults() *SkuListItemCreateData {
	this := SkuListItemCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *SkuListItemCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SkuListItemCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SkuListItemCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *SkuListItemCreateData) GetAttributes() POSTSkuListItems201ResponseDataAttributes {
	if o == nil {
		var ret POSTSkuListItems201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SkuListItemCreateData) GetAttributesOk() (*POSTSkuListItems201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SkuListItemCreateData) SetAttributes(v POSTSkuListItems201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SkuListItemCreateData) GetRelationships() SkuListItemCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret SkuListItemCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListItemCreateData) GetRelationshipsOk() (*SkuListItemCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SkuListItemCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given SkuListItemCreateDataRelationships and assigns it to the Relationships field.
func (o *SkuListItemCreateData) SetRelationships(v SkuListItemCreateDataRelationships) {
	o.Relationships = &v
}

func (o SkuListItemCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkuListItemCreateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableSkuListItemCreateData struct {
	value *SkuListItemCreateData
	isSet bool
}

func (v NullableSkuListItemCreateData) Get() *SkuListItemCreateData {
	return v.value
}

func (v *NullableSkuListItemCreateData) Set(val *SkuListItemCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuListItemCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuListItemCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuListItemCreateData(val *SkuListItemCreateData) *NullableSkuListItemCreateData {
	return &NullableSkuListItemCreateData{value: val, isSet: true}
}

func (v NullableSkuListItemCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuListItemCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
