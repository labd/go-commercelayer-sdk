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

// GETCouponCodesPromotionRules200ResponseDataInnerRelationships struct for GETCouponCodesPromotionRules200ResponseDataInnerRelationships
type GETCouponCodesPromotionRules200ResponseDataInnerRelationships struct {
	Promotion *GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion `json:"promotion,omitempty"`
	Coupons   *GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsCoupons   `json:"coupons,omitempty"`
}

// NewGETCouponCodesPromotionRules200ResponseDataInnerRelationships instantiates a new GETCouponCodesPromotionRules200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCouponCodesPromotionRules200ResponseDataInnerRelationships() *GETCouponCodesPromotionRules200ResponseDataInnerRelationships {
	this := GETCouponCodesPromotionRules200ResponseDataInnerRelationships{}
	return &this
}

// NewGETCouponCodesPromotionRules200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETCouponCodesPromotionRules200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCouponCodesPromotionRules200ResponseDataInnerRelationshipsWithDefaults() *GETCouponCodesPromotionRules200ResponseDataInnerRelationships {
	this := GETCouponCodesPromotionRules200ResponseDataInnerRelationships{}
	return &this
}

// GetPromotion returns the Promotion field value if set, zero value otherwise.
func (o *GETCouponCodesPromotionRules200ResponseDataInnerRelationships) GetPromotion() GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion {
	if o == nil || o.Promotion == nil {
		var ret GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion
		return ret
	}
	return *o.Promotion
}

// GetPromotionOk returns a tuple with the Promotion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCouponCodesPromotionRules200ResponseDataInnerRelationships) GetPromotionOk() (*GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion, bool) {
	if o == nil || o.Promotion == nil {
		return nil, false
	}
	return o.Promotion, true
}

// HasPromotion returns a boolean if a field has been set.
func (o *GETCouponCodesPromotionRules200ResponseDataInnerRelationships) HasPromotion() bool {
	if o != nil && o.Promotion != nil {
		return true
	}

	return false
}

// SetPromotion gets a reference to the given GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion and assigns it to the Promotion field.
func (o *GETCouponCodesPromotionRules200ResponseDataInnerRelationships) SetPromotion(v GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion) {
	o.Promotion = &v
}

// GetCoupons returns the Coupons field value if set, zero value otherwise.
func (o *GETCouponCodesPromotionRules200ResponseDataInnerRelationships) GetCoupons() GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsCoupons {
	if o == nil || o.Coupons == nil {
		var ret GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsCoupons
		return ret
	}
	return *o.Coupons
}

// GetCouponsOk returns a tuple with the Coupons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCouponCodesPromotionRules200ResponseDataInnerRelationships) GetCouponsOk() (*GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsCoupons, bool) {
	if o == nil || o.Coupons == nil {
		return nil, false
	}
	return o.Coupons, true
}

// HasCoupons returns a boolean if a field has been set.
func (o *GETCouponCodesPromotionRules200ResponseDataInnerRelationships) HasCoupons() bool {
	if o != nil && o.Coupons != nil {
		return true
	}

	return false
}

// SetCoupons gets a reference to the given GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsCoupons and assigns it to the Coupons field.
func (o *GETCouponCodesPromotionRules200ResponseDataInnerRelationships) SetCoupons(v GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsCoupons) {
	o.Coupons = &v
}

func (o GETCouponCodesPromotionRules200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Promotion != nil {
		toSerialize["promotion"] = o.Promotion
	}
	if o.Coupons != nil {
		toSerialize["coupons"] = o.Coupons
	}
	return json.Marshal(toSerialize)
}

type NullableGETCouponCodesPromotionRules200ResponseDataInnerRelationships struct {
	value *GETCouponCodesPromotionRules200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETCouponCodesPromotionRules200ResponseDataInnerRelationships) Get() *GETCouponCodesPromotionRules200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETCouponCodesPromotionRules200ResponseDataInnerRelationships) Set(val *GETCouponCodesPromotionRules200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCouponCodesPromotionRules200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCouponCodesPromotionRules200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCouponCodesPromotionRules200ResponseDataInnerRelationships(val *GETCouponCodesPromotionRules200ResponseDataInnerRelationships) *NullableGETCouponCodesPromotionRules200ResponseDataInnerRelationships {
	return &NullableGETCouponCodesPromotionRules200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETCouponCodesPromotionRules200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCouponCodesPromotionRules200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
