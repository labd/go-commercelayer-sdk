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

// checks if the InventoryReturnLocationData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryReturnLocationData{}

// InventoryReturnLocationData struct for InventoryReturnLocationData
type InventoryReturnLocationData struct {
	// The resource's type
	Type          interface{}                                                                   `json:"type"`
	Attributes    GETInventoryReturnLocationsInventoryReturnLocationId200ResponseDataAttributes `json:"attributes"`
	Relationships *InventoryReturnLocationDataRelationships                                     `json:"relationships,omitempty"`
}

// NewInventoryReturnLocationData instantiates a new InventoryReturnLocationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryReturnLocationData(type_ interface{}, attributes GETInventoryReturnLocationsInventoryReturnLocationId200ResponseDataAttributes) *InventoryReturnLocationData {
	this := InventoryReturnLocationData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewInventoryReturnLocationDataWithDefaults instantiates a new InventoryReturnLocationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryReturnLocationDataWithDefaults() *InventoryReturnLocationData {
	this := InventoryReturnLocationData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *InventoryReturnLocationData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryReturnLocationData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InventoryReturnLocationData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *InventoryReturnLocationData) GetAttributes() GETInventoryReturnLocationsInventoryReturnLocationId200ResponseDataAttributes {
	if o == nil {
		var ret GETInventoryReturnLocationsInventoryReturnLocationId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *InventoryReturnLocationData) GetAttributesOk() (*GETInventoryReturnLocationsInventoryReturnLocationId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *InventoryReturnLocationData) SetAttributes(v GETInventoryReturnLocationsInventoryReturnLocationId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *InventoryReturnLocationData) GetRelationships() InventoryReturnLocationDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret InventoryReturnLocationDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryReturnLocationData) GetRelationshipsOk() (*InventoryReturnLocationDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *InventoryReturnLocationData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given InventoryReturnLocationDataRelationships and assigns it to the Relationships field.
func (o *InventoryReturnLocationData) SetRelationships(v InventoryReturnLocationDataRelationships) {
	o.Relationships = &v
}

func (o InventoryReturnLocationData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryReturnLocationData) ToMap() (map[string]interface{}, error) {
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

type NullableInventoryReturnLocationData struct {
	value *InventoryReturnLocationData
	isSet bool
}

func (v NullableInventoryReturnLocationData) Get() *InventoryReturnLocationData {
	return v.value
}

func (v *NullableInventoryReturnLocationData) Set(val *InventoryReturnLocationData) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryReturnLocationData) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryReturnLocationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryReturnLocationData(val *InventoryReturnLocationData) *NullableInventoryReturnLocationData {
	return &NullableInventoryReturnLocationData{value: val, isSet: true}
}

func (v NullableInventoryReturnLocationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryReturnLocationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
