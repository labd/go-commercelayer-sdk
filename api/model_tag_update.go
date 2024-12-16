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

// checks if the TagUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagUpdate{}

// TagUpdate struct for TagUpdate
type TagUpdate struct {
	Data TagUpdateData `json:"data"`
}

// NewTagUpdate instantiates a new TagUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagUpdate(data TagUpdateData) *TagUpdate {
	this := TagUpdate{}
	this.Data = data
	return &this
}

// NewTagUpdateWithDefaults instantiates a new TagUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagUpdateWithDefaults() *TagUpdate {
	this := TagUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *TagUpdate) GetData() TagUpdateData {
	if o == nil {
		var ret TagUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TagUpdate) GetDataOk() (*TagUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TagUpdate) SetData(v TagUpdateData) {
	o.Data = v
}

func (o TagUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableTagUpdate struct {
	value *TagUpdate
	isSet bool
}

func (v NullableTagUpdate) Get() *TagUpdate {
	return v.value
}

func (v *NullableTagUpdate) Set(val *TagUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableTagUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableTagUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagUpdate(val *TagUpdate) *NullableTagUpdate {
	return &NullableTagUpdate{value: val, isSet: true}
}

func (v NullableTagUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
