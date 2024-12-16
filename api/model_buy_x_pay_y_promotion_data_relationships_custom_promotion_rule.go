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

// checks if the BuyXPayYPromotionDataRelationshipsCustomPromotionRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuyXPayYPromotionDataRelationshipsCustomPromotionRule{}

// BuyXPayYPromotionDataRelationshipsCustomPromotionRule struct for BuyXPayYPromotionDataRelationshipsCustomPromotionRule
type BuyXPayYPromotionDataRelationshipsCustomPromotionRule struct {
	Data *BuyXPayYPromotionDataRelationshipsCustomPromotionRuleData `json:"data,omitempty"`
}

// NewBuyXPayYPromotionDataRelationshipsCustomPromotionRule instantiates a new BuyXPayYPromotionDataRelationshipsCustomPromotionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuyXPayYPromotionDataRelationshipsCustomPromotionRule() *BuyXPayYPromotionDataRelationshipsCustomPromotionRule {
	this := BuyXPayYPromotionDataRelationshipsCustomPromotionRule{}
	return &this
}

// NewBuyXPayYPromotionDataRelationshipsCustomPromotionRuleWithDefaults instantiates a new BuyXPayYPromotionDataRelationshipsCustomPromotionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuyXPayYPromotionDataRelationshipsCustomPromotionRuleWithDefaults() *BuyXPayYPromotionDataRelationshipsCustomPromotionRule {
	this := BuyXPayYPromotionDataRelationshipsCustomPromotionRule{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BuyXPayYPromotionDataRelationshipsCustomPromotionRule) GetData() BuyXPayYPromotionDataRelationshipsCustomPromotionRuleData {
	if o == nil || IsNil(o.Data) {
		var ret BuyXPayYPromotionDataRelationshipsCustomPromotionRuleData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyXPayYPromotionDataRelationshipsCustomPromotionRule) GetDataOk() (*BuyXPayYPromotionDataRelationshipsCustomPromotionRuleData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BuyXPayYPromotionDataRelationshipsCustomPromotionRule) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given BuyXPayYPromotionDataRelationshipsCustomPromotionRuleData and assigns it to the Data field.
func (o *BuyXPayYPromotionDataRelationshipsCustomPromotionRule) SetData(v BuyXPayYPromotionDataRelationshipsCustomPromotionRuleData) {
	o.Data = &v
}

func (o BuyXPayYPromotionDataRelationshipsCustomPromotionRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuyXPayYPromotionDataRelationshipsCustomPromotionRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableBuyXPayYPromotionDataRelationshipsCustomPromotionRule struct {
	value *BuyXPayYPromotionDataRelationshipsCustomPromotionRule
	isSet bool
}

func (v NullableBuyXPayYPromotionDataRelationshipsCustomPromotionRule) Get() *BuyXPayYPromotionDataRelationshipsCustomPromotionRule {
	return v.value
}

func (v *NullableBuyXPayYPromotionDataRelationshipsCustomPromotionRule) Set(val *BuyXPayYPromotionDataRelationshipsCustomPromotionRule) {
	v.value = val
	v.isSet = true
}

func (v NullableBuyXPayYPromotionDataRelationshipsCustomPromotionRule) IsSet() bool {
	return v.isSet
}

func (v *NullableBuyXPayYPromotionDataRelationshipsCustomPromotionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuyXPayYPromotionDataRelationshipsCustomPromotionRule(val *BuyXPayYPromotionDataRelationshipsCustomPromotionRule) *NullableBuyXPayYPromotionDataRelationshipsCustomPromotionRule {
	return &NullableBuyXPayYPromotionDataRelationshipsCustomPromotionRule{value: val, isSet: true}
}

func (v NullableBuyXPayYPromotionDataRelationshipsCustomPromotionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuyXPayYPromotionDataRelationshipsCustomPromotionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
