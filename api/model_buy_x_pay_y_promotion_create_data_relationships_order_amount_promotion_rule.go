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

// checks if the BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule{}

// BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule struct for BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule
type BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule struct {
	Data BuyXPayYPromotionDataRelationshipsOrderAmountPromotionRuleData `json:"data"`
}

// NewBuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule instantiates a new BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule(data BuyXPayYPromotionDataRelationshipsOrderAmountPromotionRuleData) *BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule {
	this := BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule{}
	this.Data = data
	return &this
}

// NewBuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRuleWithDefaults instantiates a new BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRuleWithDefaults() *BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule {
	this := BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule{}
	return &this
}

// GetData returns the Data field value
func (o *BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule) GetData() BuyXPayYPromotionDataRelationshipsOrderAmountPromotionRuleData {
	if o == nil {
		var ret BuyXPayYPromotionDataRelationshipsOrderAmountPromotionRuleData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule) GetDataOk() (*BuyXPayYPromotionDataRelationshipsOrderAmountPromotionRuleData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule) SetData(v BuyXPayYPromotionDataRelationshipsOrderAmountPromotionRuleData) {
	o.Data = v
}

func (o BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableBuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule struct {
	value *BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule
	isSet bool
}

func (v NullableBuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule) Get() *BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule {
	return v.value
}

func (v *NullableBuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule) Set(val *BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule) {
	v.value = val
	v.isSet = true
}

func (v NullableBuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule) IsSet() bool {
	return v.isSet
}

func (v *NullableBuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule(val *BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule) *NullableBuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule {
	return &NullableBuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule{value: val, isSet: true}
}

func (v NullableBuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
