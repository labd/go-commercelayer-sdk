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

// checks if the AttachmentDataRelationshipsAttachable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttachmentDataRelationshipsAttachable{}

// AttachmentDataRelationshipsAttachable struct for AttachmentDataRelationshipsAttachable
type AttachmentDataRelationshipsAttachable struct {
	Data *AttachmentDataRelationshipsAttachableData `json:"data,omitempty"`
}

// NewAttachmentDataRelationshipsAttachable instantiates a new AttachmentDataRelationshipsAttachable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentDataRelationshipsAttachable() *AttachmentDataRelationshipsAttachable {
	this := AttachmentDataRelationshipsAttachable{}
	return &this
}

// NewAttachmentDataRelationshipsAttachableWithDefaults instantiates a new AttachmentDataRelationshipsAttachable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentDataRelationshipsAttachableWithDefaults() *AttachmentDataRelationshipsAttachable {
	this := AttachmentDataRelationshipsAttachable{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AttachmentDataRelationshipsAttachable) GetData() AttachmentDataRelationshipsAttachableData {
	if o == nil || IsNil(o.Data) {
		var ret AttachmentDataRelationshipsAttachableData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentDataRelationshipsAttachable) GetDataOk() (*AttachmentDataRelationshipsAttachableData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AttachmentDataRelationshipsAttachable) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AttachmentDataRelationshipsAttachableData and assigns it to the Data field.
func (o *AttachmentDataRelationshipsAttachable) SetData(v AttachmentDataRelationshipsAttachableData) {
	o.Data = &v
}

func (o AttachmentDataRelationshipsAttachable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttachmentDataRelationshipsAttachable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAttachmentDataRelationshipsAttachable struct {
	value *AttachmentDataRelationshipsAttachable
	isSet bool
}

func (v NullableAttachmentDataRelationshipsAttachable) Get() *AttachmentDataRelationshipsAttachable {
	return v.value
}

func (v *NullableAttachmentDataRelationshipsAttachable) Set(val *AttachmentDataRelationshipsAttachable) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentDataRelationshipsAttachable) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentDataRelationshipsAttachable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentDataRelationshipsAttachable(val *AttachmentDataRelationshipsAttachable) *NullableAttachmentDataRelationshipsAttachable {
	return &NullableAttachmentDataRelationshipsAttachable{value: val, isSet: true}
}

func (v NullableAttachmentDataRelationshipsAttachable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentDataRelationshipsAttachable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
