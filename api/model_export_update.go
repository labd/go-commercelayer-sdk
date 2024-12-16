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

// checks if the ExportUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportUpdate{}

// ExportUpdate struct for ExportUpdate
type ExportUpdate struct {
	Data ExportUpdateData `json:"data"`
}

// NewExportUpdate instantiates a new ExportUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportUpdate(data ExportUpdateData) *ExportUpdate {
	this := ExportUpdate{}
	this.Data = data
	return &this
}

// NewExportUpdateWithDefaults instantiates a new ExportUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportUpdateWithDefaults() *ExportUpdate {
	this := ExportUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *ExportUpdate) GetData() ExportUpdateData {
	if o == nil {
		var ret ExportUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ExportUpdate) GetDataOk() (*ExportUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ExportUpdate) SetData(v ExportUpdateData) {
	o.Data = v
}

func (o ExportUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableExportUpdate struct {
	value *ExportUpdate
	isSet bool
}

func (v NullableExportUpdate) Get() *ExportUpdate {
	return v.value
}

func (v *NullableExportUpdate) Set(val *ExportUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableExportUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableExportUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportUpdate(val *ExportUpdate) *NullableExportUpdate {
	return &NullableExportUpdate{value: val, isSet: true}
}

func (v NullableExportUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
