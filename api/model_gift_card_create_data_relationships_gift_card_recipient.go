/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GiftCardCreateDataRelationshipsGiftCardRecipient struct for GiftCardCreateDataRelationshipsGiftCardRecipient
type GiftCardCreateDataRelationshipsGiftCardRecipient struct {
	Data GiftCardDataRelationshipsGiftCardRecipientData `json:"data"`
}

// NewGiftCardCreateDataRelationshipsGiftCardRecipient instantiates a new GiftCardCreateDataRelationshipsGiftCardRecipient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftCardCreateDataRelationshipsGiftCardRecipient(data GiftCardDataRelationshipsGiftCardRecipientData) *GiftCardCreateDataRelationshipsGiftCardRecipient {
	this := GiftCardCreateDataRelationshipsGiftCardRecipient{}
	this.Data = data
	return &this
}

// NewGiftCardCreateDataRelationshipsGiftCardRecipientWithDefaults instantiates a new GiftCardCreateDataRelationshipsGiftCardRecipient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftCardCreateDataRelationshipsGiftCardRecipientWithDefaults() *GiftCardCreateDataRelationshipsGiftCardRecipient {
	this := GiftCardCreateDataRelationshipsGiftCardRecipient{}
	return &this
}

// GetData returns the Data field value
func (o *GiftCardCreateDataRelationshipsGiftCardRecipient) GetData() GiftCardDataRelationshipsGiftCardRecipientData {
	if o == nil {
		var ret GiftCardDataRelationshipsGiftCardRecipientData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GiftCardCreateDataRelationshipsGiftCardRecipient) GetDataOk() (*GiftCardDataRelationshipsGiftCardRecipientData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GiftCardCreateDataRelationshipsGiftCardRecipient) SetData(v GiftCardDataRelationshipsGiftCardRecipientData) {
	o.Data = v
}

func (o GiftCardCreateDataRelationshipsGiftCardRecipient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGiftCardCreateDataRelationshipsGiftCardRecipient struct {
	value *GiftCardCreateDataRelationshipsGiftCardRecipient
	isSet bool
}

func (v NullableGiftCardCreateDataRelationshipsGiftCardRecipient) Get() *GiftCardCreateDataRelationshipsGiftCardRecipient {
	return v.value
}

func (v *NullableGiftCardCreateDataRelationshipsGiftCardRecipient) Set(val *GiftCardCreateDataRelationshipsGiftCardRecipient) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftCardCreateDataRelationshipsGiftCardRecipient) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftCardCreateDataRelationshipsGiftCardRecipient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftCardCreateDataRelationshipsGiftCardRecipient(val *GiftCardCreateDataRelationshipsGiftCardRecipient) *NullableGiftCardCreateDataRelationshipsGiftCardRecipient {
	return &NullableGiftCardCreateDataRelationshipsGiftCardRecipient{value: val, isSet: true}
}

func (v NullableGiftCardCreateDataRelationshipsGiftCardRecipient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftCardCreateDataRelationshipsGiftCardRecipient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
