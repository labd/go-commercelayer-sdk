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

// checks if the LineItemDataRelationshipsReturnLineItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LineItemDataRelationshipsReturnLineItems{}

// LineItemDataRelationshipsReturnLineItems struct for LineItemDataRelationshipsReturnLineItems
type LineItemDataRelationshipsReturnLineItems struct {
	Data *LineItemDataRelationshipsReturnLineItemsData `json:"data,omitempty"`
}

// NewLineItemDataRelationshipsReturnLineItems instantiates a new LineItemDataRelationshipsReturnLineItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemDataRelationshipsReturnLineItems() *LineItemDataRelationshipsReturnLineItems {
	this := LineItemDataRelationshipsReturnLineItems{}
	return &this
}

// NewLineItemDataRelationshipsReturnLineItemsWithDefaults instantiates a new LineItemDataRelationshipsReturnLineItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemDataRelationshipsReturnLineItemsWithDefaults() *LineItemDataRelationshipsReturnLineItems {
	this := LineItemDataRelationshipsReturnLineItems{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LineItemDataRelationshipsReturnLineItems) GetData() LineItemDataRelationshipsReturnLineItemsData {
	if o == nil || IsNil(o.Data) {
		var ret LineItemDataRelationshipsReturnLineItemsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemDataRelationshipsReturnLineItems) GetDataOk() (*LineItemDataRelationshipsReturnLineItemsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LineItemDataRelationshipsReturnLineItems) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given LineItemDataRelationshipsReturnLineItemsData and assigns it to the Data field.
func (o *LineItemDataRelationshipsReturnLineItems) SetData(v LineItemDataRelationshipsReturnLineItemsData) {
	o.Data = &v
}

func (o LineItemDataRelationshipsReturnLineItems) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LineItemDataRelationshipsReturnLineItems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableLineItemDataRelationshipsReturnLineItems struct {
	value *LineItemDataRelationshipsReturnLineItems
	isSet bool
}

func (v NullableLineItemDataRelationshipsReturnLineItems) Get() *LineItemDataRelationshipsReturnLineItems {
	return v.value
}

func (v *NullableLineItemDataRelationshipsReturnLineItems) Set(val *LineItemDataRelationshipsReturnLineItems) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemDataRelationshipsReturnLineItems) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemDataRelationshipsReturnLineItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemDataRelationshipsReturnLineItems(val *LineItemDataRelationshipsReturnLineItems) *NullableLineItemDataRelationshipsReturnLineItems {
	return &NullableLineItemDataRelationshipsReturnLineItems{value: val, isSet: true}
}

func (v NullableLineItemDataRelationshipsReturnLineItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemDataRelationshipsReturnLineItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
