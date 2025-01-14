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

// checks if the InventoryReturnLocationUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryReturnLocationUpdateData{}

// InventoryReturnLocationUpdateData struct for InventoryReturnLocationUpdateData
type InventoryReturnLocationUpdateData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                                                     `json:"id"`
	Attributes    PATCHInventoryReturnLocationsInventoryReturnLocationId200ResponseDataAttributes `json:"attributes"`
	Relationships *InventoryReturnLocationUpdateDataRelationships                                 `json:"relationships,omitempty"`
}

// NewInventoryReturnLocationUpdateData instantiates a new InventoryReturnLocationUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryReturnLocationUpdateData(type_ interface{}, id interface{}, attributes PATCHInventoryReturnLocationsInventoryReturnLocationId200ResponseDataAttributes) *InventoryReturnLocationUpdateData {
	this := InventoryReturnLocationUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewInventoryReturnLocationUpdateDataWithDefaults instantiates a new InventoryReturnLocationUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryReturnLocationUpdateDataWithDefaults() *InventoryReturnLocationUpdateData {
	this := InventoryReturnLocationUpdateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *InventoryReturnLocationUpdateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryReturnLocationUpdateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InventoryReturnLocationUpdateData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *InventoryReturnLocationUpdateData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryReturnLocationUpdateData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InventoryReturnLocationUpdateData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *InventoryReturnLocationUpdateData) GetAttributes() PATCHInventoryReturnLocationsInventoryReturnLocationId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHInventoryReturnLocationsInventoryReturnLocationId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *InventoryReturnLocationUpdateData) GetAttributesOk() (*PATCHInventoryReturnLocationsInventoryReturnLocationId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *InventoryReturnLocationUpdateData) SetAttributes(v PATCHInventoryReturnLocationsInventoryReturnLocationId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *InventoryReturnLocationUpdateData) GetRelationships() InventoryReturnLocationUpdateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret InventoryReturnLocationUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryReturnLocationUpdateData) GetRelationshipsOk() (*InventoryReturnLocationUpdateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *InventoryReturnLocationUpdateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given InventoryReturnLocationUpdateDataRelationships and assigns it to the Relationships field.
func (o *InventoryReturnLocationUpdateData) SetRelationships(v InventoryReturnLocationUpdateDataRelationships) {
	o.Relationships = &v
}

func (o InventoryReturnLocationUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryReturnLocationUpdateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableInventoryReturnLocationUpdateData struct {
	value *InventoryReturnLocationUpdateData
	isSet bool
}

func (v NullableInventoryReturnLocationUpdateData) Get() *InventoryReturnLocationUpdateData {
	return v.value
}

func (v *NullableInventoryReturnLocationUpdateData) Set(val *InventoryReturnLocationUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryReturnLocationUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryReturnLocationUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryReturnLocationUpdateData(val *InventoryReturnLocationUpdateData) *NullableInventoryReturnLocationUpdateData {
	return &NullableInventoryReturnLocationUpdateData{value: val, isSet: true}
}

func (v NullableInventoryReturnLocationUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryReturnLocationUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
