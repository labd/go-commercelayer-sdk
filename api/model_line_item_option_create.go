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

// checks if the LineItemOptionCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LineItemOptionCreate{}

// LineItemOptionCreate struct for LineItemOptionCreate
type LineItemOptionCreate struct {
	Data LineItemOptionCreateData `json:"data"`
}

// NewLineItemOptionCreate instantiates a new LineItemOptionCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemOptionCreate(data LineItemOptionCreateData) *LineItemOptionCreate {
	this := LineItemOptionCreate{}
	this.Data = data
	return &this
}

// NewLineItemOptionCreateWithDefaults instantiates a new LineItemOptionCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemOptionCreateWithDefaults() *LineItemOptionCreate {
	this := LineItemOptionCreate{}
	return &this
}

// GetData returns the Data field value
func (o *LineItemOptionCreate) GetData() LineItemOptionCreateData {
	if o == nil {
		var ret LineItemOptionCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *LineItemOptionCreate) GetDataOk() (*LineItemOptionCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *LineItemOptionCreate) SetData(v LineItemOptionCreateData) {
	o.Data = v
}

func (o LineItemOptionCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LineItemOptionCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableLineItemOptionCreate struct {
	value *LineItemOptionCreate
	isSet bool
}

func (v NullableLineItemOptionCreate) Get() *LineItemOptionCreate {
	return v.value
}

func (v *NullableLineItemOptionCreate) Set(val *LineItemOptionCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemOptionCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemOptionCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemOptionCreate(val *LineItemOptionCreate) *NullableLineItemOptionCreate {
	return &NullableLineItemOptionCreate{value: val, isSet: true}
}

func (v NullableLineItemOptionCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemOptionCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
