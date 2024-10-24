/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the BuyXPayYPromotionUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuyXPayYPromotionUpdate{}

// BuyXPayYPromotionUpdate struct for BuyXPayYPromotionUpdate
type BuyXPayYPromotionUpdate struct {
	Data BuyXPayYPromotionUpdateData `json:"data"`
}

// NewBuyXPayYPromotionUpdate instantiates a new BuyXPayYPromotionUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuyXPayYPromotionUpdate(data BuyXPayYPromotionUpdateData) *BuyXPayYPromotionUpdate {
	this := BuyXPayYPromotionUpdate{}
	this.Data = data
	return &this
}

// NewBuyXPayYPromotionUpdateWithDefaults instantiates a new BuyXPayYPromotionUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuyXPayYPromotionUpdateWithDefaults() *BuyXPayYPromotionUpdate {
	this := BuyXPayYPromotionUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *BuyXPayYPromotionUpdate) GetData() BuyXPayYPromotionUpdateData {
	if o == nil {
		var ret BuyXPayYPromotionUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BuyXPayYPromotionUpdate) GetDataOk() (*BuyXPayYPromotionUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BuyXPayYPromotionUpdate) SetData(v BuyXPayYPromotionUpdateData) {
	o.Data = v
}

func (o BuyXPayYPromotionUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuyXPayYPromotionUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableBuyXPayYPromotionUpdate struct {
	value *BuyXPayYPromotionUpdate
	isSet bool
}

func (v NullableBuyXPayYPromotionUpdate) Get() *BuyXPayYPromotionUpdate {
	return v.value
}

func (v *NullableBuyXPayYPromotionUpdate) Set(val *BuyXPayYPromotionUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableBuyXPayYPromotionUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableBuyXPayYPromotionUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuyXPayYPromotionUpdate(val *BuyXPayYPromotionUpdate) *NullableBuyXPayYPromotionUpdate {
	return &NullableBuyXPayYPromotionUpdate{value: val, isSet: true}
}

func (v NullableBuyXPayYPromotionUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuyXPayYPromotionUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}