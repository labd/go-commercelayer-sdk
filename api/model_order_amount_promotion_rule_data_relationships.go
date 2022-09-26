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

// OrderAmountPromotionRuleDataRelationships struct for OrderAmountPromotionRuleDataRelationships
type OrderAmountPromotionRuleDataRelationships struct {
	Promotion *CouponCodesPromotionRuleDataRelationshipsPromotion `json:"promotion,omitempty"`
}

// NewOrderAmountPromotionRuleDataRelationships instantiates a new OrderAmountPromotionRuleDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderAmountPromotionRuleDataRelationships() *OrderAmountPromotionRuleDataRelationships {
	this := OrderAmountPromotionRuleDataRelationships{}
	return &this
}

// NewOrderAmountPromotionRuleDataRelationshipsWithDefaults instantiates a new OrderAmountPromotionRuleDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderAmountPromotionRuleDataRelationshipsWithDefaults() *OrderAmountPromotionRuleDataRelationships {
	this := OrderAmountPromotionRuleDataRelationships{}
	return &this
}

// GetPromotion returns the Promotion field value if set, zero value otherwise.
func (o *OrderAmountPromotionRuleDataRelationships) GetPromotion() CouponCodesPromotionRuleDataRelationshipsPromotion {
	if o == nil || o.Promotion == nil {
		var ret CouponCodesPromotionRuleDataRelationshipsPromotion
		return ret
	}
	return *o.Promotion
}

// GetPromotionOk returns a tuple with the Promotion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmountPromotionRuleDataRelationships) GetPromotionOk() (*CouponCodesPromotionRuleDataRelationshipsPromotion, bool) {
	if o == nil || o.Promotion == nil {
		return nil, false
	}
	return o.Promotion, true
}

// HasPromotion returns a boolean if a field has been set.
func (o *OrderAmountPromotionRuleDataRelationships) HasPromotion() bool {
	if o != nil && o.Promotion != nil {
		return true
	}

	return false
}

// SetPromotion gets a reference to the given CouponCodesPromotionRuleDataRelationshipsPromotion and assigns it to the Promotion field.
func (o *OrderAmountPromotionRuleDataRelationships) SetPromotion(v CouponCodesPromotionRuleDataRelationshipsPromotion) {
	o.Promotion = &v
}

func (o OrderAmountPromotionRuleDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Promotion != nil {
		toSerialize["promotion"] = o.Promotion
	}
	return json.Marshal(toSerialize)
}

type NullableOrderAmountPromotionRuleDataRelationships struct {
	value *OrderAmountPromotionRuleDataRelationships
	isSet bool
}

func (v NullableOrderAmountPromotionRuleDataRelationships) Get() *OrderAmountPromotionRuleDataRelationships {
	return v.value
}

func (v *NullableOrderAmountPromotionRuleDataRelationships) Set(val *OrderAmountPromotionRuleDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderAmountPromotionRuleDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderAmountPromotionRuleDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderAmountPromotionRuleDataRelationships(val *OrderAmountPromotionRuleDataRelationships) *NullableOrderAmountPromotionRuleDataRelationships {
	return &NullableOrderAmountPromotionRuleDataRelationships{value: val, isSet: true}
}

func (v NullableOrderAmountPromotionRuleDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderAmountPromotionRuleDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
