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

// checks if the POSTAttachments201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTAttachments201ResponseDataRelationships{}

// POSTAttachments201ResponseDataRelationships struct for POSTAttachments201ResponseDataRelationships
type POSTAttachments201ResponseDataRelationships struct {
	Attachable *POSTAttachments201ResponseDataRelationshipsAttachable `json:"attachable,omitempty"`
}

// NewPOSTAttachments201ResponseDataRelationships instantiates a new POSTAttachments201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAttachments201ResponseDataRelationships() *POSTAttachments201ResponseDataRelationships {
	this := POSTAttachments201ResponseDataRelationships{}
	return &this
}

// NewPOSTAttachments201ResponseDataRelationshipsWithDefaults instantiates a new POSTAttachments201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAttachments201ResponseDataRelationshipsWithDefaults() *POSTAttachments201ResponseDataRelationships {
	this := POSTAttachments201ResponseDataRelationships{}
	return &this
}

// GetAttachable returns the Attachable field value if set, zero value otherwise.
func (o *POSTAttachments201ResponseDataRelationships) GetAttachable() POSTAttachments201ResponseDataRelationshipsAttachable {
	if o == nil || IsNil(o.Attachable) {
		var ret POSTAttachments201ResponseDataRelationshipsAttachable
		return ret
	}
	return *o.Attachable
}

// GetAttachableOk returns a tuple with the Attachable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAttachments201ResponseDataRelationships) GetAttachableOk() (*POSTAttachments201ResponseDataRelationshipsAttachable, bool) {
	if o == nil || IsNil(o.Attachable) {
		return nil, false
	}
	return o.Attachable, true
}

// HasAttachable returns a boolean if a field has been set.
func (o *POSTAttachments201ResponseDataRelationships) HasAttachable() bool {
	if o != nil && !IsNil(o.Attachable) {
		return true
	}

	return false
}

// SetAttachable gets a reference to the given POSTAttachments201ResponseDataRelationshipsAttachable and assigns it to the Attachable field.
func (o *POSTAttachments201ResponseDataRelationships) SetAttachable(v POSTAttachments201ResponseDataRelationshipsAttachable) {
	o.Attachable = &v
}

func (o POSTAttachments201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTAttachments201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attachable) {
		toSerialize["attachable"] = o.Attachable
	}
	return toSerialize, nil
}

type NullablePOSTAttachments201ResponseDataRelationships struct {
	value *POSTAttachments201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTAttachments201ResponseDataRelationships) Get() *POSTAttachments201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTAttachments201ResponseDataRelationships) Set(val *POSTAttachments201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAttachments201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAttachments201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAttachments201ResponseDataRelationships(val *POSTAttachments201ResponseDataRelationships) *NullablePOSTAttachments201ResponseDataRelationships {
	return &NullablePOSTAttachments201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTAttachments201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAttachments201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
