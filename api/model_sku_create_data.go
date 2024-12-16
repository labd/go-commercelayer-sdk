/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.5.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the SkuCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkuCreateData{}

// SkuCreateData struct for SkuCreateData
type SkuCreateData struct {
	// The resource's type
	Type          interface{}                       `json:"type"`
	Attributes    POSTSkus201ResponseDataAttributes `json:"attributes"`
	Relationships *SkuCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewSkuCreateData instantiates a new SkuCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuCreateData(type_ interface{}, attributes POSTSkus201ResponseDataAttributes) *SkuCreateData {
	this := SkuCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewSkuCreateDataWithDefaults instantiates a new SkuCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuCreateDataWithDefaults() *SkuCreateData {
	this := SkuCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *SkuCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SkuCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SkuCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *SkuCreateData) GetAttributes() POSTSkus201ResponseDataAttributes {
	if o == nil {
		var ret POSTSkus201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SkuCreateData) GetAttributesOk() (*POSTSkus201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SkuCreateData) SetAttributes(v POSTSkus201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SkuCreateData) GetRelationships() SkuCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret SkuCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuCreateData) GetRelationshipsOk() (*SkuCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SkuCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given SkuCreateDataRelationships and assigns it to the Relationships field.
func (o *SkuCreateData) SetRelationships(v SkuCreateDataRelationships) {
	o.Relationships = &v
}

func (o SkuCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkuCreateData) ToMap() (map[string]interface{}, error) {
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

type NullableSkuCreateData struct {
	value *SkuCreateData
	isSet bool
}

func (v NullableSkuCreateData) Get() *SkuCreateData {
	return v.value
}

func (v *NullableSkuCreateData) Set(val *SkuCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuCreateData(val *SkuCreateData) *NullableSkuCreateData {
	return &NullableSkuCreateData{value: val, isSet: true}
}

func (v NullableSkuCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
