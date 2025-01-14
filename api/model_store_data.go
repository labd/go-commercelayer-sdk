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

// checks if the StoreData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StoreData{}

// StoreData struct for StoreData
type StoreData struct {
	// The resource's type
	Type          interface{}                               `json:"type"`
	Attributes    GETStoresStoreId200ResponseDataAttributes `json:"attributes"`
	Relationships *StoreDataRelationships                   `json:"relationships,omitempty"`
}

// NewStoreData instantiates a new StoreData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreData(type_ interface{}, attributes GETStoresStoreId200ResponseDataAttributes) *StoreData {
	this := StoreData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewStoreDataWithDefaults instantiates a new StoreData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreDataWithDefaults() *StoreData {
	this := StoreData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *StoreData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoreData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StoreData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *StoreData) GetAttributes() GETStoresStoreId200ResponseDataAttributes {
	if o == nil {
		var ret GETStoresStoreId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *StoreData) GetAttributesOk() (*GETStoresStoreId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *StoreData) SetAttributes(v GETStoresStoreId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *StoreData) GetRelationships() StoreDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret StoreDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreData) GetRelationshipsOk() (*StoreDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *StoreData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given StoreDataRelationships and assigns it to the Relationships field.
func (o *StoreData) SetRelationships(v StoreDataRelationships) {
	o.Relationships = &v
}

func (o StoreData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoreData) ToMap() (map[string]interface{}, error) {
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

type NullableStoreData struct {
	value *StoreData
	isSet bool
}

func (v NullableStoreData) Get() *StoreData {
	return v.value
}

func (v *NullableStoreData) Set(val *StoreData) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreData) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreData(val *StoreData) *NullableStoreData {
	return &NullableStoreData{value: val, isSet: true}
}

func (v NullableStoreData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}