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

// GETParcelLineItems200ResponseDataInnerRelationshipsParcelData struct for GETParcelLineItems200ResponseDataInnerRelationshipsParcelData
type GETParcelLineItems200ResponseDataInnerRelationshipsParcelData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETParcelLineItems200ResponseDataInnerRelationshipsParcelData instantiates a new GETParcelLineItems200ResponseDataInnerRelationshipsParcelData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETParcelLineItems200ResponseDataInnerRelationshipsParcelData() *GETParcelLineItems200ResponseDataInnerRelationshipsParcelData {
	this := GETParcelLineItems200ResponseDataInnerRelationshipsParcelData{}
	return &this
}

// NewGETParcelLineItems200ResponseDataInnerRelationshipsParcelDataWithDefaults instantiates a new GETParcelLineItems200ResponseDataInnerRelationshipsParcelData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETParcelLineItems200ResponseDataInnerRelationshipsParcelDataWithDefaults() *GETParcelLineItems200ResponseDataInnerRelationshipsParcelData {
	this := GETParcelLineItems200ResponseDataInnerRelationshipsParcelData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETParcelLineItems200ResponseDataInnerRelationshipsParcelData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETParcelLineItems200ResponseDataInnerRelationshipsParcelData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETParcelLineItems200ResponseDataInnerRelationshipsParcelData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETParcelLineItems200ResponseDataInnerRelationshipsParcelData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETParcelLineItems200ResponseDataInnerRelationshipsParcelData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETParcelLineItems200ResponseDataInnerRelationshipsParcelData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETParcelLineItems200ResponseDataInnerRelationshipsParcelData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETParcelLineItems200ResponseDataInnerRelationshipsParcelData) SetId(v string) {
	o.Id = &v
}

func (o GETParcelLineItems200ResponseDataInnerRelationshipsParcelData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETParcelLineItems200ResponseDataInnerRelationshipsParcelData struct {
	value *GETParcelLineItems200ResponseDataInnerRelationshipsParcelData
	isSet bool
}

func (v NullableGETParcelLineItems200ResponseDataInnerRelationshipsParcelData) Get() *GETParcelLineItems200ResponseDataInnerRelationshipsParcelData {
	return v.value
}

func (v *NullableGETParcelLineItems200ResponseDataInnerRelationshipsParcelData) Set(val *GETParcelLineItems200ResponseDataInnerRelationshipsParcelData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETParcelLineItems200ResponseDataInnerRelationshipsParcelData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETParcelLineItems200ResponseDataInnerRelationshipsParcelData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETParcelLineItems200ResponseDataInnerRelationshipsParcelData(val *GETParcelLineItems200ResponseDataInnerRelationshipsParcelData) *NullableGETParcelLineItems200ResponseDataInnerRelationshipsParcelData {
	return &NullableGETParcelLineItems200ResponseDataInnerRelationshipsParcelData{value: val, isSet: true}
}

func (v NullableGETParcelLineItems200ResponseDataInnerRelationshipsParcelData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETParcelLineItems200ResponseDataInnerRelationshipsParcelData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
