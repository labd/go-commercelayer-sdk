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

// checks if the LineItemCreateDataRelationshipsItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LineItemCreateDataRelationshipsItem{}

// LineItemCreateDataRelationshipsItem struct for LineItemCreateDataRelationshipsItem
type LineItemCreateDataRelationshipsItem struct {
	Data LineItemDataRelationshipsItemData `json:"data"`
}

// NewLineItemCreateDataRelationshipsItem instantiates a new LineItemCreateDataRelationshipsItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemCreateDataRelationshipsItem(data LineItemDataRelationshipsItemData) *LineItemCreateDataRelationshipsItem {
	this := LineItemCreateDataRelationshipsItem{}
	this.Data = data
	return &this
}

// NewLineItemCreateDataRelationshipsItemWithDefaults instantiates a new LineItemCreateDataRelationshipsItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemCreateDataRelationshipsItemWithDefaults() *LineItemCreateDataRelationshipsItem {
	this := LineItemCreateDataRelationshipsItem{}
	return &this
}

// GetData returns the Data field value
func (o *LineItemCreateDataRelationshipsItem) GetData() LineItemDataRelationshipsItemData {
	if o == nil {
		var ret LineItemDataRelationshipsItemData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *LineItemCreateDataRelationshipsItem) GetDataOk() (*LineItemDataRelationshipsItemData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *LineItemCreateDataRelationshipsItem) SetData(v LineItemDataRelationshipsItemData) {
	o.Data = v
}

func (o LineItemCreateDataRelationshipsItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LineItemCreateDataRelationshipsItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableLineItemCreateDataRelationshipsItem struct {
	value *LineItemCreateDataRelationshipsItem
	isSet bool
}

func (v NullableLineItemCreateDataRelationshipsItem) Get() *LineItemCreateDataRelationshipsItem {
	return v.value
}

func (v *NullableLineItemCreateDataRelationshipsItem) Set(val *LineItemCreateDataRelationshipsItem) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemCreateDataRelationshipsItem) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemCreateDataRelationshipsItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemCreateDataRelationshipsItem(val *LineItemCreateDataRelationshipsItem) *NullableLineItemCreateDataRelationshipsItem {
	return &NullableLineItemCreateDataRelationshipsItem{value: val, isSet: true}
}

func (v NullableLineItemCreateDataRelationshipsItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemCreateDataRelationshipsItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
