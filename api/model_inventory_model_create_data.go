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

// checks if the InventoryModelCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryModelCreateData{}

// InventoryModelCreateData struct for InventoryModelCreateData
type InventoryModelCreateData struct {
	// The resource's type
	Type          interface{}                                  `json:"type"`
	Attributes    POSTInventoryModels201ResponseDataAttributes `json:"attributes"`
	Relationships interface{}                                  `json:"relationships,omitempty"`
}

// NewInventoryModelCreateData instantiates a new InventoryModelCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryModelCreateData(type_ interface{}, attributes POSTInventoryModels201ResponseDataAttributes) *InventoryModelCreateData {
	this := InventoryModelCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewInventoryModelCreateDataWithDefaults instantiates a new InventoryModelCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryModelCreateDataWithDefaults() *InventoryModelCreateData {
	this := InventoryModelCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *InventoryModelCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryModelCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InventoryModelCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *InventoryModelCreateData) GetAttributes() POSTInventoryModels201ResponseDataAttributes {
	if o == nil {
		var ret POSTInventoryModels201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *InventoryModelCreateData) GetAttributesOk() (*POSTInventoryModels201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *InventoryModelCreateData) SetAttributes(v POSTInventoryModels201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryModelCreateData) GetRelationships() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryModelCreateData) GetRelationshipsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return &o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *InventoryModelCreateData) HasRelationships() bool {
	if o != nil && IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given interface{} and assigns it to the Relationships field.
func (o *InventoryModelCreateData) SetRelationships(v interface{}) {
	o.Relationships = v
}

func (o InventoryModelCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryModelCreateData) ToMap() (map[string]interface{}, error) {
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

type NullableInventoryModelCreateData struct {
	value *InventoryModelCreateData
	isSet bool
}

func (v NullableInventoryModelCreateData) Get() *InventoryModelCreateData {
	return v.value
}

func (v *NullableInventoryModelCreateData) Set(val *InventoryModelCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryModelCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryModelCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryModelCreateData(val *InventoryModelCreateData) *NullableInventoryModelCreateData {
	return &NullableInventoryModelCreateData{value: val, isSet: true}
}

func (v NullableInventoryModelCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryModelCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
