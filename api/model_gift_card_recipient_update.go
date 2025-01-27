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

// checks if the GiftCardRecipientUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftCardRecipientUpdate{}

// GiftCardRecipientUpdate struct for GiftCardRecipientUpdate
type GiftCardRecipientUpdate struct {
	Data GiftCardRecipientUpdateData `json:"data"`
}

// NewGiftCardRecipientUpdate instantiates a new GiftCardRecipientUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftCardRecipientUpdate(data GiftCardRecipientUpdateData) *GiftCardRecipientUpdate {
	this := GiftCardRecipientUpdate{}
	this.Data = data
	return &this
}

// NewGiftCardRecipientUpdateWithDefaults instantiates a new GiftCardRecipientUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftCardRecipientUpdateWithDefaults() *GiftCardRecipientUpdate {
	this := GiftCardRecipientUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *GiftCardRecipientUpdate) GetData() GiftCardRecipientUpdateData {
	if o == nil {
		var ret GiftCardRecipientUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GiftCardRecipientUpdate) GetDataOk() (*GiftCardRecipientUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GiftCardRecipientUpdate) SetData(v GiftCardRecipientUpdateData) {
	o.Data = v
}

func (o GiftCardRecipientUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftCardRecipientUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGiftCardRecipientUpdate struct {
	value *GiftCardRecipientUpdate
	isSet bool
}

func (v NullableGiftCardRecipientUpdate) Get() *GiftCardRecipientUpdate {
	return v.value
}

func (v *NullableGiftCardRecipientUpdate) Set(val *GiftCardRecipientUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftCardRecipientUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftCardRecipientUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftCardRecipientUpdate(val *GiftCardRecipientUpdate) *NullableGiftCardRecipientUpdate {
	return &NullableGiftCardRecipientUpdate{value: val, isSet: true}
}

func (v NullableGiftCardRecipientUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftCardRecipientUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
