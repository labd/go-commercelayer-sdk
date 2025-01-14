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

// checks if the CouponCodesPromotionRuleCreateDataRelationshipsPromotion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CouponCodesPromotionRuleCreateDataRelationshipsPromotion{}

// CouponCodesPromotionRuleCreateDataRelationshipsPromotion struct for CouponCodesPromotionRuleCreateDataRelationshipsPromotion
type CouponCodesPromotionRuleCreateDataRelationshipsPromotion struct {
	Data CouponCodesPromotionRuleDataRelationshipsPromotionData `json:"data"`
}

// NewCouponCodesPromotionRuleCreateDataRelationshipsPromotion instantiates a new CouponCodesPromotionRuleCreateDataRelationshipsPromotion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponCodesPromotionRuleCreateDataRelationshipsPromotion(data CouponCodesPromotionRuleDataRelationshipsPromotionData) *CouponCodesPromotionRuleCreateDataRelationshipsPromotion {
	this := CouponCodesPromotionRuleCreateDataRelationshipsPromotion{}
	this.Data = data
	return &this
}

// NewCouponCodesPromotionRuleCreateDataRelationshipsPromotionWithDefaults instantiates a new CouponCodesPromotionRuleCreateDataRelationshipsPromotion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponCodesPromotionRuleCreateDataRelationshipsPromotionWithDefaults() *CouponCodesPromotionRuleCreateDataRelationshipsPromotion {
	this := CouponCodesPromotionRuleCreateDataRelationshipsPromotion{}
	return &this
}

// GetData returns the Data field value
func (o *CouponCodesPromotionRuleCreateDataRelationshipsPromotion) GetData() CouponCodesPromotionRuleDataRelationshipsPromotionData {
	if o == nil {
		var ret CouponCodesPromotionRuleDataRelationshipsPromotionData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CouponCodesPromotionRuleCreateDataRelationshipsPromotion) GetDataOk() (*CouponCodesPromotionRuleDataRelationshipsPromotionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CouponCodesPromotionRuleCreateDataRelationshipsPromotion) SetData(v CouponCodesPromotionRuleDataRelationshipsPromotionData) {
	o.Data = v
}

func (o CouponCodesPromotionRuleCreateDataRelationshipsPromotion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CouponCodesPromotionRuleCreateDataRelationshipsPromotion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableCouponCodesPromotionRuleCreateDataRelationshipsPromotion struct {
	value *CouponCodesPromotionRuleCreateDataRelationshipsPromotion
	isSet bool
}

func (v NullableCouponCodesPromotionRuleCreateDataRelationshipsPromotion) Get() *CouponCodesPromotionRuleCreateDataRelationshipsPromotion {
	return v.value
}

func (v *NullableCouponCodesPromotionRuleCreateDataRelationshipsPromotion) Set(val *CouponCodesPromotionRuleCreateDataRelationshipsPromotion) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponCodesPromotionRuleCreateDataRelationshipsPromotion) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponCodesPromotionRuleCreateDataRelationshipsPromotion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponCodesPromotionRuleCreateDataRelationshipsPromotion(val *CouponCodesPromotionRuleCreateDataRelationshipsPromotion) *NullableCouponCodesPromotionRuleCreateDataRelationshipsPromotion {
	return &NullableCouponCodesPromotionRuleCreateDataRelationshipsPromotion{value: val, isSet: true}
}

func (v NullableCouponCodesPromotionRuleCreateDataRelationshipsPromotion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponCodesPromotionRuleCreateDataRelationshipsPromotion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
