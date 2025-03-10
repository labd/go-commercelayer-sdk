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

// checks if the BuyXPayYPromotion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuyXPayYPromotion{}

// BuyXPayYPromotion struct for BuyXPayYPromotion
type BuyXPayYPromotion struct {
	Data *BuyXPayYPromotionData `json:"data,omitempty"`
}

// NewBuyXPayYPromotion instantiates a new BuyXPayYPromotion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuyXPayYPromotion() *BuyXPayYPromotion {
	this := BuyXPayYPromotion{}
	return &this
}

// NewBuyXPayYPromotionWithDefaults instantiates a new BuyXPayYPromotion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuyXPayYPromotionWithDefaults() *BuyXPayYPromotion {
	this := BuyXPayYPromotion{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BuyXPayYPromotion) GetData() BuyXPayYPromotionData {
	if o == nil || IsNil(o.Data) {
		var ret BuyXPayYPromotionData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyXPayYPromotion) GetDataOk() (*BuyXPayYPromotionData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BuyXPayYPromotion) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given BuyXPayYPromotionData and assigns it to the Data field.
func (o *BuyXPayYPromotion) SetData(v BuyXPayYPromotionData) {
	o.Data = &v
}

func (o BuyXPayYPromotion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuyXPayYPromotion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableBuyXPayYPromotion struct {
	value *BuyXPayYPromotion
	isSet bool
}

func (v NullableBuyXPayYPromotion) Get() *BuyXPayYPromotion {
	return v.value
}

func (v *NullableBuyXPayYPromotion) Set(val *BuyXPayYPromotion) {
	v.value = val
	v.isSet = true
}

func (v NullableBuyXPayYPromotion) IsSet() bool {
	return v.isSet
}

func (v *NullableBuyXPayYPromotion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuyXPayYPromotion(val *BuyXPayYPromotion) *NullableBuyXPayYPromotion {
	return &NullableBuyXPayYPromotion{value: val, isSet: true}
}

func (v NullableBuyXPayYPromotion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuyXPayYPromotion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
