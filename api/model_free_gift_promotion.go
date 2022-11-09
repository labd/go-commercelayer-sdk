/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.2
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// FreeGiftPromotion struct for FreeGiftPromotion
type FreeGiftPromotion struct {
	Data FreeGiftPromotionData `json:"data"`
}

// NewFreeGiftPromotion instantiates a new FreeGiftPromotion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFreeGiftPromotion(data FreeGiftPromotionData) *FreeGiftPromotion {
	this := FreeGiftPromotion{}
	this.Data = data
	return &this
}

// NewFreeGiftPromotionWithDefaults instantiates a new FreeGiftPromotion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFreeGiftPromotionWithDefaults() *FreeGiftPromotion {
	this := FreeGiftPromotion{}
	return &this
}

// GetData returns the Data field value
func (o *FreeGiftPromotion) GetData() FreeGiftPromotionData {
	if o == nil {
		var ret FreeGiftPromotionData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FreeGiftPromotion) GetDataOk() (*FreeGiftPromotionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *FreeGiftPromotion) SetData(v FreeGiftPromotionData) {
	o.Data = v
}

func (o FreeGiftPromotion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableFreeGiftPromotion struct {
	value *FreeGiftPromotion
	isSet bool
}

func (v NullableFreeGiftPromotion) Get() *FreeGiftPromotion {
	return v.value
}

func (v *NullableFreeGiftPromotion) Set(val *FreeGiftPromotion) {
	v.value = val
	v.isSet = true
}

func (v NullableFreeGiftPromotion) IsSet() bool {
	return v.isSet
}

func (v *NullableFreeGiftPromotion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFreeGiftPromotion(val *FreeGiftPromotion) *NullableFreeGiftPromotion {
	return &NullableFreeGiftPromotion{value: val, isSet: true}
}

func (v NullableFreeGiftPromotion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFreeGiftPromotion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
