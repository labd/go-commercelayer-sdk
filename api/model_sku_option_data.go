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

// checks if the SkuOptionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkuOptionData{}

// SkuOptionData struct for SkuOptionData
type SkuOptionData struct {
	// The resource's type
	Type          interface{}                                       `json:"type"`
	Attributes    GETSkuOptionsSkuOptionId200ResponseDataAttributes `json:"attributes"`
	Relationships *SkuOptionDataRelationships                       `json:"relationships,omitempty"`
}

// NewSkuOptionData instantiates a new SkuOptionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuOptionData(type_ interface{}, attributes GETSkuOptionsSkuOptionId200ResponseDataAttributes) *SkuOptionData {
	this := SkuOptionData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewSkuOptionDataWithDefaults instantiates a new SkuOptionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuOptionDataWithDefaults() *SkuOptionData {
	this := SkuOptionData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *SkuOptionData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SkuOptionData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SkuOptionData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *SkuOptionData) GetAttributes() GETSkuOptionsSkuOptionId200ResponseDataAttributes {
	if o == nil {
		var ret GETSkuOptionsSkuOptionId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SkuOptionData) GetAttributesOk() (*GETSkuOptionsSkuOptionId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SkuOptionData) SetAttributes(v GETSkuOptionsSkuOptionId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SkuOptionData) GetRelationships() SkuOptionDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret SkuOptionDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuOptionData) GetRelationshipsOk() (*SkuOptionDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SkuOptionData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given SkuOptionDataRelationships and assigns it to the Relationships field.
func (o *SkuOptionData) SetRelationships(v SkuOptionDataRelationships) {
	o.Relationships = &v
}

func (o SkuOptionData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkuOptionData) ToMap() (map[string]interface{}, error) {
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

type NullableSkuOptionData struct {
	value *SkuOptionData
	isSet bool
}

func (v NullableSkuOptionData) Get() *SkuOptionData {
	return v.value
}

func (v *NullableSkuOptionData) Set(val *SkuOptionData) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuOptionData) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuOptionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuOptionData(val *SkuOptionData) *NullableSkuOptionData {
	return &NullableSkuOptionData{value: val, isSet: true}
}

func (v NullableSkuOptionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuOptionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
