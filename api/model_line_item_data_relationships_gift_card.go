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

// checks if the LineItemDataRelationshipsGiftCard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LineItemDataRelationshipsGiftCard{}

// LineItemDataRelationshipsGiftCard struct for LineItemDataRelationshipsGiftCard
type LineItemDataRelationshipsGiftCard struct {
	Data *LineItemDataRelationshipsGiftCardData `json:"data,omitempty"`
}

// NewLineItemDataRelationshipsGiftCard instantiates a new LineItemDataRelationshipsGiftCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemDataRelationshipsGiftCard() *LineItemDataRelationshipsGiftCard {
	this := LineItemDataRelationshipsGiftCard{}
	return &this
}

// NewLineItemDataRelationshipsGiftCardWithDefaults instantiates a new LineItemDataRelationshipsGiftCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemDataRelationshipsGiftCardWithDefaults() *LineItemDataRelationshipsGiftCard {
	this := LineItemDataRelationshipsGiftCard{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LineItemDataRelationshipsGiftCard) GetData() LineItemDataRelationshipsGiftCardData {
	if o == nil || IsNil(o.Data) {
		var ret LineItemDataRelationshipsGiftCardData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemDataRelationshipsGiftCard) GetDataOk() (*LineItemDataRelationshipsGiftCardData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LineItemDataRelationshipsGiftCard) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given LineItemDataRelationshipsGiftCardData and assigns it to the Data field.
func (o *LineItemDataRelationshipsGiftCard) SetData(v LineItemDataRelationshipsGiftCardData) {
	o.Data = &v
}

func (o LineItemDataRelationshipsGiftCard) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LineItemDataRelationshipsGiftCard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableLineItemDataRelationshipsGiftCard struct {
	value *LineItemDataRelationshipsGiftCard
	isSet bool
}

func (v NullableLineItemDataRelationshipsGiftCard) Get() *LineItemDataRelationshipsGiftCard {
	return v.value
}

func (v *NullableLineItemDataRelationshipsGiftCard) Set(val *LineItemDataRelationshipsGiftCard) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemDataRelationshipsGiftCard) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemDataRelationshipsGiftCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemDataRelationshipsGiftCard(val *LineItemDataRelationshipsGiftCard) *NullableLineItemDataRelationshipsGiftCard {
	return &NullableLineItemDataRelationshipsGiftCard{value: val, isSet: true}
}

func (v NullableLineItemDataRelationshipsGiftCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemDataRelationshipsGiftCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
