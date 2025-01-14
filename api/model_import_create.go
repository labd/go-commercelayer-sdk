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

// checks if the ImportCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportCreate{}

// ImportCreate struct for ImportCreate
type ImportCreate struct {
	Data ImportCreateData `json:"data"`
}

// NewImportCreate instantiates a new ImportCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportCreate(data ImportCreateData) *ImportCreate {
	this := ImportCreate{}
	this.Data = data
	return &this
}

// NewImportCreateWithDefaults instantiates a new ImportCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportCreateWithDefaults() *ImportCreate {
	this := ImportCreate{}
	return &this
}

// GetData returns the Data field value
func (o *ImportCreate) GetData() ImportCreateData {
	if o == nil {
		var ret ImportCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ImportCreate) GetDataOk() (*ImportCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ImportCreate) SetData(v ImportCreateData) {
	o.Data = v
}

func (o ImportCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableImportCreate struct {
	value *ImportCreate
	isSet bool
}

func (v NullableImportCreate) Get() *ImportCreate {
	return v.value
}

func (v *NullableImportCreate) Set(val *ImportCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableImportCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableImportCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportCreate(val *ImportCreate) *NullableImportCreate {
	return &NullableImportCreate{value: val, isSet: true}
}

func (v NullableImportCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
