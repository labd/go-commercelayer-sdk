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

// checks if the LineItemDataRelationshipsAdjustment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LineItemDataRelationshipsAdjustment{}

// LineItemDataRelationshipsAdjustment struct for LineItemDataRelationshipsAdjustment
type LineItemDataRelationshipsAdjustment struct {
	Data *LineItemDataRelationshipsAdjustmentData `json:"data,omitempty"`
}

// NewLineItemDataRelationshipsAdjustment instantiates a new LineItemDataRelationshipsAdjustment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemDataRelationshipsAdjustment() *LineItemDataRelationshipsAdjustment {
	this := LineItemDataRelationshipsAdjustment{}
	return &this
}

// NewLineItemDataRelationshipsAdjustmentWithDefaults instantiates a new LineItemDataRelationshipsAdjustment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemDataRelationshipsAdjustmentWithDefaults() *LineItemDataRelationshipsAdjustment {
	this := LineItemDataRelationshipsAdjustment{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LineItemDataRelationshipsAdjustment) GetData() LineItemDataRelationshipsAdjustmentData {
	if o == nil || IsNil(o.Data) {
		var ret LineItemDataRelationshipsAdjustmentData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemDataRelationshipsAdjustment) GetDataOk() (*LineItemDataRelationshipsAdjustmentData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LineItemDataRelationshipsAdjustment) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given LineItemDataRelationshipsAdjustmentData and assigns it to the Data field.
func (o *LineItemDataRelationshipsAdjustment) SetData(v LineItemDataRelationshipsAdjustmentData) {
	o.Data = &v
}

func (o LineItemDataRelationshipsAdjustment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LineItemDataRelationshipsAdjustment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableLineItemDataRelationshipsAdjustment struct {
	value *LineItemDataRelationshipsAdjustment
	isSet bool
}

func (v NullableLineItemDataRelationshipsAdjustment) Get() *LineItemDataRelationshipsAdjustment {
	return v.value
}

func (v *NullableLineItemDataRelationshipsAdjustment) Set(val *LineItemDataRelationshipsAdjustment) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemDataRelationshipsAdjustment) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemDataRelationshipsAdjustment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemDataRelationshipsAdjustment(val *LineItemDataRelationshipsAdjustment) *NullableLineItemDataRelationshipsAdjustment {
	return &NullableLineItemDataRelationshipsAdjustment{value: val, isSet: true}
}

func (v NullableLineItemDataRelationshipsAdjustment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemDataRelationshipsAdjustment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
