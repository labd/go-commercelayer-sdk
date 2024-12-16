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

// checks if the GiftCardCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftCardCreate{}

// GiftCardCreate struct for GiftCardCreate
type GiftCardCreate struct {
	Data GiftCardCreateData `json:"data"`
}

// NewGiftCardCreate instantiates a new GiftCardCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftCardCreate(data GiftCardCreateData) *GiftCardCreate {
	this := GiftCardCreate{}
	this.Data = data
	return &this
}

// NewGiftCardCreateWithDefaults instantiates a new GiftCardCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftCardCreateWithDefaults() *GiftCardCreate {
	this := GiftCardCreate{}
	return &this
}

// GetData returns the Data field value
func (o *GiftCardCreate) GetData() GiftCardCreateData {
	if o == nil {
		var ret GiftCardCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GiftCardCreate) GetDataOk() (*GiftCardCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GiftCardCreate) SetData(v GiftCardCreateData) {
	o.Data = v
}

func (o GiftCardCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftCardCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGiftCardCreate struct {
	value *GiftCardCreate
	isSet bool
}

func (v NullableGiftCardCreate) Get() *GiftCardCreate {
	return v.value
}

func (v *NullableGiftCardCreate) Set(val *GiftCardCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftCardCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftCardCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftCardCreate(val *GiftCardCreate) *NullableGiftCardCreate {
	return &NullableGiftCardCreate{value: val, isSet: true}
}

func (v NullableGiftCardCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftCardCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
