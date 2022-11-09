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

// GETLineItems200ResponseDataInnerRelationshipsItemData struct for GETLineItems200ResponseDataInnerRelationshipsItemData
type GETLineItems200ResponseDataInnerRelationshipsItemData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETLineItems200ResponseDataInnerRelationshipsItemData instantiates a new GETLineItems200ResponseDataInnerRelationshipsItemData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETLineItems200ResponseDataInnerRelationshipsItemData() *GETLineItems200ResponseDataInnerRelationshipsItemData {
	this := GETLineItems200ResponseDataInnerRelationshipsItemData{}
	return &this
}

// NewGETLineItems200ResponseDataInnerRelationshipsItemDataWithDefaults instantiates a new GETLineItems200ResponseDataInnerRelationshipsItemData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETLineItems200ResponseDataInnerRelationshipsItemDataWithDefaults() *GETLineItems200ResponseDataInnerRelationshipsItemData {
	this := GETLineItems200ResponseDataInnerRelationshipsItemData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETLineItems200ResponseDataInnerRelationshipsItemData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsItemData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsItemData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETLineItems200ResponseDataInnerRelationshipsItemData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETLineItems200ResponseDataInnerRelationshipsItemData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsItemData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsItemData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETLineItems200ResponseDataInnerRelationshipsItemData) SetId(v string) {
	o.Id = &v
}

func (o GETLineItems200ResponseDataInnerRelationshipsItemData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETLineItems200ResponseDataInnerRelationshipsItemData struct {
	value *GETLineItems200ResponseDataInnerRelationshipsItemData
	isSet bool
}

func (v NullableGETLineItems200ResponseDataInnerRelationshipsItemData) Get() *GETLineItems200ResponseDataInnerRelationshipsItemData {
	return v.value
}

func (v *NullableGETLineItems200ResponseDataInnerRelationshipsItemData) Set(val *GETLineItems200ResponseDataInnerRelationshipsItemData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETLineItems200ResponseDataInnerRelationshipsItemData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETLineItems200ResponseDataInnerRelationshipsItemData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETLineItems200ResponseDataInnerRelationshipsItemData(val *GETLineItems200ResponseDataInnerRelationshipsItemData) *NullableGETLineItems200ResponseDataInnerRelationshipsItemData {
	return &NullableGETLineItems200ResponseDataInnerRelationshipsItemData{value: val, isSet: true}
}

func (v NullableGETLineItems200ResponseDataInnerRelationshipsItemData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETLineItems200ResponseDataInnerRelationshipsItemData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}