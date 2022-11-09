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

// GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData struct for GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData
type GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData instantiates a new GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData() *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData {
	this := GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData{}
	return &this
}

// NewGETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelDataWithDefaults instantiates a new GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelDataWithDefaults() *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData {
	this := GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData) SetId(v string) {
	o.Id = &v
}

func (o GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData struct {
	value *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData
	isSet bool
}

func (v NullableGETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData) Get() *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData {
	return v.value
}

func (v *NullableGETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData) Set(val *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData(val *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData) *NullableGETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData {
	return &NullableGETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData{value: val, isSet: true}
}

func (v NullableGETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}