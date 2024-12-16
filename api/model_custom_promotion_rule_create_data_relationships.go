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

// checks if the CustomPromotionRuleCreateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomPromotionRuleCreateDataRelationships{}

// CustomPromotionRuleCreateDataRelationships struct for CustomPromotionRuleCreateDataRelationships
type CustomPromotionRuleCreateDataRelationships struct {
	Promotion CouponCodesPromotionRuleCreateDataRelationshipsPromotion `json:"promotion"`
}

// NewCustomPromotionRuleCreateDataRelationships instantiates a new CustomPromotionRuleCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomPromotionRuleCreateDataRelationships(promotion CouponCodesPromotionRuleCreateDataRelationshipsPromotion) *CustomPromotionRuleCreateDataRelationships {
	this := CustomPromotionRuleCreateDataRelationships{}
	this.Promotion = promotion
	return &this
}

// NewCustomPromotionRuleCreateDataRelationshipsWithDefaults instantiates a new CustomPromotionRuleCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomPromotionRuleCreateDataRelationshipsWithDefaults() *CustomPromotionRuleCreateDataRelationships {
	this := CustomPromotionRuleCreateDataRelationships{}
	return &this
}

// GetPromotion returns the Promotion field value
func (o *CustomPromotionRuleCreateDataRelationships) GetPromotion() CouponCodesPromotionRuleCreateDataRelationshipsPromotion {
	if o == nil {
		var ret CouponCodesPromotionRuleCreateDataRelationshipsPromotion
		return ret
	}

	return o.Promotion
}

// GetPromotionOk returns a tuple with the Promotion field value
// and a boolean to check if the value has been set.
func (o *CustomPromotionRuleCreateDataRelationships) GetPromotionOk() (*CouponCodesPromotionRuleCreateDataRelationshipsPromotion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Promotion, true
}

// SetPromotion sets field value
func (o *CustomPromotionRuleCreateDataRelationships) SetPromotion(v CouponCodesPromotionRuleCreateDataRelationshipsPromotion) {
	o.Promotion = v
}

func (o CustomPromotionRuleCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomPromotionRuleCreateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["promotion"] = o.Promotion
	return toSerialize, nil
}

type NullableCustomPromotionRuleCreateDataRelationships struct {
	value *CustomPromotionRuleCreateDataRelationships
	isSet bool
}

func (v NullableCustomPromotionRuleCreateDataRelationships) Get() *CustomPromotionRuleCreateDataRelationships {
	return v.value
}

func (v *NullableCustomPromotionRuleCreateDataRelationships) Set(val *CustomPromotionRuleCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomPromotionRuleCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomPromotionRuleCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomPromotionRuleCreateDataRelationships(val *CustomPromotionRuleCreateDataRelationships) *NullableCustomPromotionRuleCreateDataRelationships {
	return &NullableCustomPromotionRuleCreateDataRelationships{value: val, isSet: true}
}

func (v NullableCustomPromotionRuleCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomPromotionRuleCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
