/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GiftCardRecipientResponseList struct for GiftCardRecipientResponseList
type GiftCardRecipientResponseList struct {
	Data []Data `json:"data,omitempty"`
}

// NewGiftCardRecipientResponseList instantiates a new GiftCardRecipientResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftCardRecipientResponseList() *GiftCardRecipientResponseList {
	this := GiftCardRecipientResponseList{}
	return &this
}

// NewGiftCardRecipientResponseListWithDefaults instantiates a new GiftCardRecipientResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftCardRecipientResponseListWithDefaults() *GiftCardRecipientResponseList {
	this := GiftCardRecipientResponseList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GiftCardRecipientResponseList) GetData() []Data {
	if o == nil || o.Data == nil {
		var ret []Data
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardRecipientResponseList) GetDataOk() ([]Data, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GiftCardRecipientResponseList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Data and assigns it to the Data field.
func (o *GiftCardRecipientResponseList) SetData(v []Data) {
	o.Data = v
}

func (o GiftCardRecipientResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGiftCardRecipientResponseList struct {
	value *GiftCardRecipientResponseList
	isSet bool
}

func (v NullableGiftCardRecipientResponseList) Get() *GiftCardRecipientResponseList {
	return v.value
}

func (v *NullableGiftCardRecipientResponseList) Set(val *GiftCardRecipientResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftCardRecipientResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftCardRecipientResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftCardRecipientResponseList(val *GiftCardRecipientResponseList) *NullableGiftCardRecipientResponseList {
	return &NullableGiftCardRecipientResponseList{value: val, isSet: true}
}

func (v NullableGiftCardRecipientResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftCardRecipientResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
