/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PackageUpdate struct for PackageUpdate
type PackageUpdate struct {
	Data PackageUpdateData `json:"data"`
}

// NewPackageUpdate instantiates a new PackageUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageUpdate(data PackageUpdateData) *PackageUpdate {
	this := PackageUpdate{}
	this.Data = data
	return &this
}

// NewPackageUpdateWithDefaults instantiates a new PackageUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageUpdateWithDefaults() *PackageUpdate {
	this := PackageUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *PackageUpdate) GetData() PackageUpdateData {
	if o == nil {
		var ret PackageUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PackageUpdate) GetDataOk() (*PackageUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PackageUpdate) SetData(v PackageUpdateData) {
	o.Data = v
}

func (o PackageUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePackageUpdate struct {
	value *PackageUpdate
	isSet bool
}

func (v NullablePackageUpdate) Get() *PackageUpdate {
	return v.value
}

func (v *NullablePackageUpdate) Set(val *PackageUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageUpdate(val *PackageUpdate) *NullablePackageUpdate {
	return &NullablePackageUpdate{value: val, isSet: true}
}

func (v NullablePackageUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
