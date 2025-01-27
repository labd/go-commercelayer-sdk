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

// checks if the AttachmentCreateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttachmentCreateDataRelationships{}

// AttachmentCreateDataRelationships struct for AttachmentCreateDataRelationships
type AttachmentCreateDataRelationships struct {
	Attachable AttachmentCreateDataRelationshipsAttachable `json:"attachable"`
}

// NewAttachmentCreateDataRelationships instantiates a new AttachmentCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentCreateDataRelationships(attachable AttachmentCreateDataRelationshipsAttachable) *AttachmentCreateDataRelationships {
	this := AttachmentCreateDataRelationships{}
	this.Attachable = attachable
	return &this
}

// NewAttachmentCreateDataRelationshipsWithDefaults instantiates a new AttachmentCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentCreateDataRelationshipsWithDefaults() *AttachmentCreateDataRelationships {
	this := AttachmentCreateDataRelationships{}
	return &this
}

// GetAttachable returns the Attachable field value
func (o *AttachmentCreateDataRelationships) GetAttachable() AttachmentCreateDataRelationshipsAttachable {
	if o == nil {
		var ret AttachmentCreateDataRelationshipsAttachable
		return ret
	}

	return o.Attachable
}

// GetAttachableOk returns a tuple with the Attachable field value
// and a boolean to check if the value has been set.
func (o *AttachmentCreateDataRelationships) GetAttachableOk() (*AttachmentCreateDataRelationshipsAttachable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attachable, true
}

// SetAttachable sets field value
func (o *AttachmentCreateDataRelationships) SetAttachable(v AttachmentCreateDataRelationshipsAttachable) {
	o.Attachable = v
}

func (o AttachmentCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttachmentCreateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attachable"] = o.Attachable
	return toSerialize, nil
}

type NullableAttachmentCreateDataRelationships struct {
	value *AttachmentCreateDataRelationships
	isSet bool
}

func (v NullableAttachmentCreateDataRelationships) Get() *AttachmentCreateDataRelationships {
	return v.value
}

func (v *NullableAttachmentCreateDataRelationships) Set(val *AttachmentCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentCreateDataRelationships(val *AttachmentCreateDataRelationships) *NullableAttachmentCreateDataRelationships {
	return &NullableAttachmentCreateDataRelationships{value: val, isSet: true}
}

func (v NullableAttachmentCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
