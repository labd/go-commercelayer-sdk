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

// checks if the GiftCardDataRelationshipsGiftCardRecipient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftCardDataRelationshipsGiftCardRecipient{}

// GiftCardDataRelationshipsGiftCardRecipient struct for GiftCardDataRelationshipsGiftCardRecipient
type GiftCardDataRelationshipsGiftCardRecipient struct {
	Data *GiftCardDataRelationshipsGiftCardRecipientData `json:"data,omitempty"`
}

// NewGiftCardDataRelationshipsGiftCardRecipient instantiates a new GiftCardDataRelationshipsGiftCardRecipient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftCardDataRelationshipsGiftCardRecipient() *GiftCardDataRelationshipsGiftCardRecipient {
	this := GiftCardDataRelationshipsGiftCardRecipient{}
	return &this
}

// NewGiftCardDataRelationshipsGiftCardRecipientWithDefaults instantiates a new GiftCardDataRelationshipsGiftCardRecipient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftCardDataRelationshipsGiftCardRecipientWithDefaults() *GiftCardDataRelationshipsGiftCardRecipient {
	this := GiftCardDataRelationshipsGiftCardRecipient{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GiftCardDataRelationshipsGiftCardRecipient) GetData() GiftCardDataRelationshipsGiftCardRecipientData {
	if o == nil || IsNil(o.Data) {
		var ret GiftCardDataRelationshipsGiftCardRecipientData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataRelationshipsGiftCardRecipient) GetDataOk() (*GiftCardDataRelationshipsGiftCardRecipientData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GiftCardDataRelationshipsGiftCardRecipient) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GiftCardDataRelationshipsGiftCardRecipientData and assigns it to the Data field.
func (o *GiftCardDataRelationshipsGiftCardRecipient) SetData(v GiftCardDataRelationshipsGiftCardRecipientData) {
	o.Data = &v
}

func (o GiftCardDataRelationshipsGiftCardRecipient) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftCardDataRelationshipsGiftCardRecipient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGiftCardDataRelationshipsGiftCardRecipient struct {
	value *GiftCardDataRelationshipsGiftCardRecipient
	isSet bool
}

func (v NullableGiftCardDataRelationshipsGiftCardRecipient) Get() *GiftCardDataRelationshipsGiftCardRecipient {
	return v.value
}

func (v *NullableGiftCardDataRelationshipsGiftCardRecipient) Set(val *GiftCardDataRelationshipsGiftCardRecipient) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftCardDataRelationshipsGiftCardRecipient) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftCardDataRelationshipsGiftCardRecipient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftCardDataRelationshipsGiftCardRecipient(val *GiftCardDataRelationshipsGiftCardRecipient) *NullableGiftCardDataRelationshipsGiftCardRecipient {
	return &NullableGiftCardDataRelationshipsGiftCardRecipient{value: val, isSet: true}
}

func (v NullableGiftCardDataRelationshipsGiftCardRecipient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftCardDataRelationshipsGiftCardRecipient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
