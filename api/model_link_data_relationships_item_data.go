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

// checks if the LinkDataRelationshipsItemData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkDataRelationshipsItemData{}

// LinkDataRelationshipsItemData struct for LinkDataRelationshipsItemData
type LinkDataRelationshipsItemData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource's id
	Id interface{} `json:"id,omitempty"`
}

// NewLinkDataRelationshipsItemData instantiates a new LinkDataRelationshipsItemData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkDataRelationshipsItemData() *LinkDataRelationshipsItemData {
	this := LinkDataRelationshipsItemData{}
	return &this
}

// NewLinkDataRelationshipsItemDataWithDefaults instantiates a new LinkDataRelationshipsItemData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkDataRelationshipsItemDataWithDefaults() *LinkDataRelationshipsItemData {
	this := LinkDataRelationshipsItemData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LinkDataRelationshipsItemData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinkDataRelationshipsItemData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LinkDataRelationshipsItemData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *LinkDataRelationshipsItemData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LinkDataRelationshipsItemData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinkDataRelationshipsItemData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LinkDataRelationshipsItemData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *LinkDataRelationshipsItemData) SetId(v interface{}) {
	o.Id = v
}

func (o LinkDataRelationshipsItemData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkDataRelationshipsItemData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableLinkDataRelationshipsItemData struct {
	value *LinkDataRelationshipsItemData
	isSet bool
}

func (v NullableLinkDataRelationshipsItemData) Get() *LinkDataRelationshipsItemData {
	return v.value
}

func (v *NullableLinkDataRelationshipsItemData) Set(val *LinkDataRelationshipsItemData) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkDataRelationshipsItemData) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkDataRelationshipsItemData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkDataRelationshipsItemData(val *LinkDataRelationshipsItemData) *NullableLinkDataRelationshipsItemData {
	return &NullableLinkDataRelationshipsItemData{value: val, isSet: true}
}

func (v NullableLinkDataRelationshipsItemData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkDataRelationshipsItemData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
