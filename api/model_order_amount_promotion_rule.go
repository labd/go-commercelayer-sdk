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

// OrderAmountPromotionRule struct for OrderAmountPromotionRule
type OrderAmountPromotionRule struct {
	Data OrderAmountPromotionRuleData `json:"data"`
}

// NewOrderAmountPromotionRule instantiates a new OrderAmountPromotionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderAmountPromotionRule(data OrderAmountPromotionRuleData) *OrderAmountPromotionRule {
	this := OrderAmountPromotionRule{}
	this.Data = data
	return &this
}

// NewOrderAmountPromotionRuleWithDefaults instantiates a new OrderAmountPromotionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderAmountPromotionRuleWithDefaults() *OrderAmountPromotionRule {
	this := OrderAmountPromotionRule{}
	return &this
}

// GetData returns the Data field value
func (o *OrderAmountPromotionRule) GetData() OrderAmountPromotionRuleData {
	if o == nil {
		var ret OrderAmountPromotionRuleData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OrderAmountPromotionRule) GetDataOk() (*OrderAmountPromotionRuleData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *OrderAmountPromotionRule) SetData(v OrderAmountPromotionRuleData) {
	o.Data = v
}

func (o OrderAmountPromotionRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableOrderAmountPromotionRule struct {
	value *OrderAmountPromotionRule
	isSet bool
}

func (v NullableOrderAmountPromotionRule) Get() *OrderAmountPromotionRule {
	return v.value
}

func (v *NullableOrderAmountPromotionRule) Set(val *OrderAmountPromotionRule) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderAmountPromotionRule) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderAmountPromotionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderAmountPromotionRule(val *OrderAmountPromotionRule) *NullableOrderAmountPromotionRule {
	return &NullableOrderAmountPromotionRule{value: val, isSet: true}
}

func (v NullableOrderAmountPromotionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderAmountPromotionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
