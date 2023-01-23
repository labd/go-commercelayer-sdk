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

// GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData struct for GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData
type GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData instantiates a new GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData() *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData {
	this := GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData{}
	return &this
}

// NewGETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsDataWithDefaults instantiates a new GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsDataWithDefaults() *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData {
	this := GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData) SetId(v string) {
	o.Id = &v
}

func (o GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData struct {
	value *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData
	isSet bool
}

func (v NullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData) Get() *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData {
	return v.value
}

func (v *NullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData) Set(val *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData(val *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData) *NullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData {
	return &NullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData{value: val, isSet: true}
}

func (v NullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
