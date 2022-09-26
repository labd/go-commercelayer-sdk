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

// CouponCodesPromotionRuleDataRelationshipsCoupons struct for CouponCodesPromotionRuleDataRelationshipsCoupons
type CouponCodesPromotionRuleDataRelationshipsCoupons struct {
	Data CouponCodesPromotionRuleDataRelationshipsCouponsData `json:"data"`
}

// NewCouponCodesPromotionRuleDataRelationshipsCoupons instantiates a new CouponCodesPromotionRuleDataRelationshipsCoupons object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponCodesPromotionRuleDataRelationshipsCoupons(data CouponCodesPromotionRuleDataRelationshipsCouponsData) *CouponCodesPromotionRuleDataRelationshipsCoupons {
	this := CouponCodesPromotionRuleDataRelationshipsCoupons{}
	this.Data = data
	return &this
}

// NewCouponCodesPromotionRuleDataRelationshipsCouponsWithDefaults instantiates a new CouponCodesPromotionRuleDataRelationshipsCoupons object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponCodesPromotionRuleDataRelationshipsCouponsWithDefaults() *CouponCodesPromotionRuleDataRelationshipsCoupons {
	this := CouponCodesPromotionRuleDataRelationshipsCoupons{}
	return &this
}

// GetData returns the Data field value
func (o *CouponCodesPromotionRuleDataRelationshipsCoupons) GetData() CouponCodesPromotionRuleDataRelationshipsCouponsData {
	if o == nil {
		var ret CouponCodesPromotionRuleDataRelationshipsCouponsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CouponCodesPromotionRuleDataRelationshipsCoupons) GetDataOk() (*CouponCodesPromotionRuleDataRelationshipsCouponsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CouponCodesPromotionRuleDataRelationshipsCoupons) SetData(v CouponCodesPromotionRuleDataRelationshipsCouponsData) {
	o.Data = v
}

func (o CouponCodesPromotionRuleDataRelationshipsCoupons) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCouponCodesPromotionRuleDataRelationshipsCoupons struct {
	value *CouponCodesPromotionRuleDataRelationshipsCoupons
	isSet bool
}

func (v NullableCouponCodesPromotionRuleDataRelationshipsCoupons) Get() *CouponCodesPromotionRuleDataRelationshipsCoupons {
	return v.value
}

func (v *NullableCouponCodesPromotionRuleDataRelationshipsCoupons) Set(val *CouponCodesPromotionRuleDataRelationshipsCoupons) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponCodesPromotionRuleDataRelationshipsCoupons) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponCodesPromotionRuleDataRelationshipsCoupons) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponCodesPromotionRuleDataRelationshipsCoupons(val *CouponCodesPromotionRuleDataRelationshipsCoupons) *NullableCouponCodesPromotionRuleDataRelationshipsCoupons {
	return &NullableCouponCodesPromotionRuleDataRelationshipsCoupons{value: val, isSet: true}
}

func (v NullableCouponCodesPromotionRuleDataRelationshipsCoupons) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponCodesPromotionRuleDataRelationshipsCoupons) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
