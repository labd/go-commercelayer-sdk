/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETPackages200ResponseDataInnerRelationshipsParcelsData struct for GETPackages200ResponseDataInnerRelationshipsParcelsData
type GETPackages200ResponseDataInnerRelationshipsParcelsData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETPackages200ResponseDataInnerRelationshipsParcelsData instantiates a new GETPackages200ResponseDataInnerRelationshipsParcelsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPackages200ResponseDataInnerRelationshipsParcelsData() *GETPackages200ResponseDataInnerRelationshipsParcelsData {
	this := GETPackages200ResponseDataInnerRelationshipsParcelsData{}
	return &this
}

// NewGETPackages200ResponseDataInnerRelationshipsParcelsDataWithDefaults instantiates a new GETPackages200ResponseDataInnerRelationshipsParcelsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPackages200ResponseDataInnerRelationshipsParcelsDataWithDefaults() *GETPackages200ResponseDataInnerRelationshipsParcelsData {
	this := GETPackages200ResponseDataInnerRelationshipsParcelsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETPackages200ResponseDataInnerRelationshipsParcelsData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPackages200ResponseDataInnerRelationshipsParcelsData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETPackages200ResponseDataInnerRelationshipsParcelsData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETPackages200ResponseDataInnerRelationshipsParcelsData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETPackages200ResponseDataInnerRelationshipsParcelsData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPackages200ResponseDataInnerRelationshipsParcelsData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETPackages200ResponseDataInnerRelationshipsParcelsData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETPackages200ResponseDataInnerRelationshipsParcelsData) SetId(v string) {
	o.Id = &v
}

func (o GETPackages200ResponseDataInnerRelationshipsParcelsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETPackages200ResponseDataInnerRelationshipsParcelsData struct {
	value *GETPackages200ResponseDataInnerRelationshipsParcelsData
	isSet bool
}

func (v NullableGETPackages200ResponseDataInnerRelationshipsParcelsData) Get() *GETPackages200ResponseDataInnerRelationshipsParcelsData {
	return v.value
}

func (v *NullableGETPackages200ResponseDataInnerRelationshipsParcelsData) Set(val *GETPackages200ResponseDataInnerRelationshipsParcelsData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPackages200ResponseDataInnerRelationshipsParcelsData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPackages200ResponseDataInnerRelationshipsParcelsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPackages200ResponseDataInnerRelationshipsParcelsData(val *GETPackages200ResponseDataInnerRelationshipsParcelsData) *NullableGETPackages200ResponseDataInnerRelationshipsParcelsData {
	return &NullableGETPackages200ResponseDataInnerRelationshipsParcelsData{value: val, isSet: true}
}

func (v NullableGETPackages200ResponseDataInnerRelationshipsParcelsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPackages200ResponseDataInnerRelationshipsParcelsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
