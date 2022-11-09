/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.2
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETAttachments200ResponseDataInnerRelationshipsAttachableData struct for GETAttachments200ResponseDataInnerRelationshipsAttachableData
type GETAttachments200ResponseDataInnerRelationshipsAttachableData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETAttachments200ResponseDataInnerRelationshipsAttachableData instantiates a new GETAttachments200ResponseDataInnerRelationshipsAttachableData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAttachments200ResponseDataInnerRelationshipsAttachableData() *GETAttachments200ResponseDataInnerRelationshipsAttachableData {
	this := GETAttachments200ResponseDataInnerRelationshipsAttachableData{}
	return &this
}

// NewGETAttachments200ResponseDataInnerRelationshipsAttachableDataWithDefaults instantiates a new GETAttachments200ResponseDataInnerRelationshipsAttachableData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAttachments200ResponseDataInnerRelationshipsAttachableDataWithDefaults() *GETAttachments200ResponseDataInnerRelationshipsAttachableData {
	this := GETAttachments200ResponseDataInnerRelationshipsAttachableData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETAttachments200ResponseDataInnerRelationshipsAttachableData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAttachments200ResponseDataInnerRelationshipsAttachableData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETAttachments200ResponseDataInnerRelationshipsAttachableData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETAttachments200ResponseDataInnerRelationshipsAttachableData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETAttachments200ResponseDataInnerRelationshipsAttachableData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAttachments200ResponseDataInnerRelationshipsAttachableData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETAttachments200ResponseDataInnerRelationshipsAttachableData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETAttachments200ResponseDataInnerRelationshipsAttachableData) SetId(v string) {
	o.Id = &v
}

func (o GETAttachments200ResponseDataInnerRelationshipsAttachableData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETAttachments200ResponseDataInnerRelationshipsAttachableData struct {
	value *GETAttachments200ResponseDataInnerRelationshipsAttachableData
	isSet bool
}

func (v NullableGETAttachments200ResponseDataInnerRelationshipsAttachableData) Get() *GETAttachments200ResponseDataInnerRelationshipsAttachableData {
	return v.value
}

func (v *NullableGETAttachments200ResponseDataInnerRelationshipsAttachableData) Set(val *GETAttachments200ResponseDataInnerRelationshipsAttachableData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAttachments200ResponseDataInnerRelationshipsAttachableData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAttachments200ResponseDataInnerRelationshipsAttachableData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAttachments200ResponseDataInnerRelationshipsAttachableData(val *GETAttachments200ResponseDataInnerRelationshipsAttachableData) *NullableGETAttachments200ResponseDataInnerRelationshipsAttachableData {
	return &NullableGETAttachments200ResponseDataInnerRelationshipsAttachableData{value: val, isSet: true}
}

func (v NullableGETAttachments200ResponseDataInnerRelationshipsAttachableData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAttachments200ResponseDataInnerRelationshipsAttachableData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
