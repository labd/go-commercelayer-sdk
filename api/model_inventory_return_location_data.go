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

// InventoryReturnLocationData struct for InventoryReturnLocationData
type InventoryReturnLocationData struct {
	// The resource's type
	Type          string                                                    `json:"type"`
	Attributes    GETInventoryReturnLocations200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *InventoryReturnLocationDataRelationships                 `json:"relationships,omitempty"`
}

// NewInventoryReturnLocationData instantiates a new InventoryReturnLocationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryReturnLocationData(type_ string, attributes GETInventoryReturnLocations200ResponseDataInnerAttributes) *InventoryReturnLocationData {
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
func (o *InventoryReturnLocationData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InventoryReturnLocationData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InventoryReturnLocationData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *InventoryReturnLocationData) GetAttributes() GETInventoryReturnLocations200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETInventoryReturnLocations200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *InventoryReturnLocationData) GetAttributesOk() (*GETInventoryReturnLocations200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *InventoryReturnLocationData) SetAttributes(v GETInventoryReturnLocations200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *InventoryReturnLocationData) GetRelationships() InventoryReturnLocationDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret InventoryReturnLocationDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryReturnLocationData) GetRelationshipsOk() (*InventoryReturnLocationDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *InventoryReturnLocationData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given InventoryReturnLocationDataRelationships and assigns it to the Relationships field.
func (o *InventoryReturnLocationData) SetRelationships(v InventoryReturnLocationDataRelationships) {
	o.Relationships = &v
}

func (o InventoryReturnLocationData) MarshalJSON() ([]byte, error) {
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
