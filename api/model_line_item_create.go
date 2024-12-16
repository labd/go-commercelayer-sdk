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

// checks if the LineItemCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LineItemCreate{}

// LineItemCreate struct for LineItemCreate
type LineItemCreate struct {
	Data LineItemCreateData `json:"data"`
}

// NewLineItemCreate instantiates a new LineItemCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemCreate(data LineItemCreateData) *LineItemCreate {
	this := LineItemCreate{}
	this.Data = data
	return &this
}

// NewLineItemCreateWithDefaults instantiates a new LineItemCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemCreateWithDefaults() *LineItemCreate {
	this := LineItemCreate{}
	return &this
}

// GetData returns the Data field value
func (o *LineItemCreate) GetData() LineItemCreateData {
	if o == nil {
		var ret LineItemCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *LineItemCreate) GetDataOk() (*LineItemCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *LineItemCreate) SetData(v LineItemCreateData) {
	o.Data = v
}

func (o LineItemCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LineItemCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableLineItemCreate struct {
	value *LineItemCreate
	isSet bool
}

func (v NullableLineItemCreate) Get() *LineItemCreate {
	return v.value
}

func (v *NullableLineItemCreate) Set(val *LineItemCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemCreate(val *LineItemCreate) *NullableLineItemCreate {
	return &NullableLineItemCreate{value: val, isSet: true}
}

func (v NullableLineItemCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
