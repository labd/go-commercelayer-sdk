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

// CouponCreateDataRelationships struct for CouponCreateDataRelationships
type CouponCreateDataRelationships struct {
	PromotionRule CouponDataRelationshipsPromotionRule `json:"promotion_rule"`
}

// NewCouponCreateDataRelationships instantiates a new CouponCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponCreateDataRelationships(promotionRule CouponDataRelationshipsPromotionRule) *CouponCreateDataRelationships {
	this := CouponCreateDataRelationships{}
	this.PromotionRule = promotionRule
	return &this
}

// NewCouponCreateDataRelationshipsWithDefaults instantiates a new CouponCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponCreateDataRelationshipsWithDefaults() *CouponCreateDataRelationships {
	this := CouponCreateDataRelationships{}
	return &this
}

// GetPromotionRule returns the PromotionRule field value
func (o *CouponCreateDataRelationships) GetPromotionRule() CouponDataRelationshipsPromotionRule {
	if o == nil {
		var ret CouponDataRelationshipsPromotionRule
		return ret
	}

	return o.PromotionRule
}

// GetPromotionRuleOk returns a tuple with the PromotionRule field value
// and a boolean to check if the value has been set.
func (o *CouponCreateDataRelationships) GetPromotionRuleOk() (*CouponDataRelationshipsPromotionRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PromotionRule, true
}

// SetPromotionRule sets field value
func (o *CouponCreateDataRelationships) SetPromotionRule(v CouponDataRelationshipsPromotionRule) {
	o.PromotionRule = v
}

func (o CouponCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["promotion_rule"] = o.PromotionRule
	}
	return json.Marshal(toSerialize)
}

type NullableCouponCreateDataRelationships struct {
	value *CouponCreateDataRelationships
	isSet bool
}

func (v NullableCouponCreateDataRelationships) Get() *CouponCreateDataRelationships {
	return v.value
}

func (v *NullableCouponCreateDataRelationships) Set(val *CouponCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponCreateDataRelationships(val *CouponCreateDataRelationships) *NullableCouponCreateDataRelationships {
	return &NullableCouponCreateDataRelationships{value: val, isSet: true}
}

func (v NullableCouponCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
