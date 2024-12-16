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

// checks if the POSTAttachments201ResponseDataRelationshipsAttachableData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTAttachments201ResponseDataRelationshipsAttachableData{}

// POSTAttachments201ResponseDataRelationshipsAttachableData struct for POSTAttachments201ResponseDataRelationshipsAttachableData
type POSTAttachments201ResponseDataRelationshipsAttachableData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTAttachments201ResponseDataRelationshipsAttachableData instantiates a new POSTAttachments201ResponseDataRelationshipsAttachableData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAttachments201ResponseDataRelationshipsAttachableData() *POSTAttachments201ResponseDataRelationshipsAttachableData {
	this := POSTAttachments201ResponseDataRelationshipsAttachableData{}
	return &this
}

// NewPOSTAttachments201ResponseDataRelationshipsAttachableDataWithDefaults instantiates a new POSTAttachments201ResponseDataRelationshipsAttachableData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAttachments201ResponseDataRelationshipsAttachableDataWithDefaults() *POSTAttachments201ResponseDataRelationshipsAttachableData {
	this := POSTAttachments201ResponseDataRelationshipsAttachableData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAttachments201ResponseDataRelationshipsAttachableData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAttachments201ResponseDataRelationshipsAttachableData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTAttachments201ResponseDataRelationshipsAttachableData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTAttachments201ResponseDataRelationshipsAttachableData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAttachments201ResponseDataRelationshipsAttachableData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAttachments201ResponseDataRelationshipsAttachableData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTAttachments201ResponseDataRelationshipsAttachableData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTAttachments201ResponseDataRelationshipsAttachableData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTAttachments201ResponseDataRelationshipsAttachableData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTAttachments201ResponseDataRelationshipsAttachableData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTAttachments201ResponseDataRelationshipsAttachableData struct {
	value *POSTAttachments201ResponseDataRelationshipsAttachableData
	isSet bool
}

func (v NullablePOSTAttachments201ResponseDataRelationshipsAttachableData) Get() *POSTAttachments201ResponseDataRelationshipsAttachableData {
	return v.value
}

func (v *NullablePOSTAttachments201ResponseDataRelationshipsAttachableData) Set(val *POSTAttachments201ResponseDataRelationshipsAttachableData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAttachments201ResponseDataRelationshipsAttachableData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAttachments201ResponseDataRelationshipsAttachableData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAttachments201ResponseDataRelationshipsAttachableData(val *POSTAttachments201ResponseDataRelationshipsAttachableData) *NullablePOSTAttachments201ResponseDataRelationshipsAttachableData {
	return &NullablePOSTAttachments201ResponseDataRelationshipsAttachableData{value: val, isSet: true}
}

func (v NullablePOSTAttachments201ResponseDataRelationshipsAttachableData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAttachments201ResponseDataRelationshipsAttachableData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
