/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// CouponCodesPromotionRuleCreateDataRelationshipsPromotion - struct for CouponCodesPromotionRuleCreateDataRelationshipsPromotion
type CouponCodesPromotionRuleCreateDataRelationshipsPromotion struct {
	ExternalPromotion           *ExternalPromotion
	FixedAmountPromotion        *FixedAmountPromotion
	FixedPricePromotion         *FixedPricePromotion
	FreeGiftPromotion           *FreeGiftPromotion
	FreeShippingPromotion       *FreeShippingPromotion
	PercentageDiscountPromotion *PercentageDiscountPromotion
}

// ExternalPromotionAsCouponCodesPromotionRuleCreateDataRelationshipsPromotion is a convenience function that returns ExternalPromotion wrapped in CouponCodesPromotionRuleCreateDataRelationshipsPromotion
func ExternalPromotionAsCouponCodesPromotionRuleCreateDataRelationshipsPromotion(v *ExternalPromotion) CouponCodesPromotionRuleCreateDataRelationshipsPromotion {
	return CouponCodesPromotionRuleCreateDataRelationshipsPromotion{
		ExternalPromotion: v,
	}
}

// FixedAmountPromotionAsCouponCodesPromotionRuleCreateDataRelationshipsPromotion is a convenience function that returns FixedAmountPromotion wrapped in CouponCodesPromotionRuleCreateDataRelationshipsPromotion
func FixedAmountPromotionAsCouponCodesPromotionRuleCreateDataRelationshipsPromotion(v *FixedAmountPromotion) CouponCodesPromotionRuleCreateDataRelationshipsPromotion {
	return CouponCodesPromotionRuleCreateDataRelationshipsPromotion{
		FixedAmountPromotion: v,
	}
}

// FixedPricePromotionAsCouponCodesPromotionRuleCreateDataRelationshipsPromotion is a convenience function that returns FixedPricePromotion wrapped in CouponCodesPromotionRuleCreateDataRelationshipsPromotion
func FixedPricePromotionAsCouponCodesPromotionRuleCreateDataRelationshipsPromotion(v *FixedPricePromotion) CouponCodesPromotionRuleCreateDataRelationshipsPromotion {
	return CouponCodesPromotionRuleCreateDataRelationshipsPromotion{
		FixedPricePromotion: v,
	}
}

// FreeGiftPromotionAsCouponCodesPromotionRuleCreateDataRelationshipsPromotion is a convenience function that returns FreeGiftPromotion wrapped in CouponCodesPromotionRuleCreateDataRelationshipsPromotion
func FreeGiftPromotionAsCouponCodesPromotionRuleCreateDataRelationshipsPromotion(v *FreeGiftPromotion) CouponCodesPromotionRuleCreateDataRelationshipsPromotion {
	return CouponCodesPromotionRuleCreateDataRelationshipsPromotion{
		FreeGiftPromotion: v,
	}
}

// FreeShippingPromotionAsCouponCodesPromotionRuleCreateDataRelationshipsPromotion is a convenience function that returns FreeShippingPromotion wrapped in CouponCodesPromotionRuleCreateDataRelationshipsPromotion
func FreeShippingPromotionAsCouponCodesPromotionRuleCreateDataRelationshipsPromotion(v *FreeShippingPromotion) CouponCodesPromotionRuleCreateDataRelationshipsPromotion {
	return CouponCodesPromotionRuleCreateDataRelationshipsPromotion{
		FreeShippingPromotion: v,
	}
}

// PercentageDiscountPromotionAsCouponCodesPromotionRuleCreateDataRelationshipsPromotion is a convenience function that returns PercentageDiscountPromotion wrapped in CouponCodesPromotionRuleCreateDataRelationshipsPromotion
func PercentageDiscountPromotionAsCouponCodesPromotionRuleCreateDataRelationshipsPromotion(v *PercentageDiscountPromotion) CouponCodesPromotionRuleCreateDataRelationshipsPromotion {
	return CouponCodesPromotionRuleCreateDataRelationshipsPromotion{
		PercentageDiscountPromotion: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CouponCodesPromotionRuleCreateDataRelationshipsPromotion) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ExternalPromotion
	err = newStrictDecoder(data).Decode(&dst.ExternalPromotion)
	if err == nil {
		jsonExternalPromotion, _ := json.Marshal(dst.ExternalPromotion)
		if string(jsonExternalPromotion) == "{}" { // empty struct
			dst.ExternalPromotion = nil
		} else {
			match++
		}
	} else {
		dst.ExternalPromotion = nil
	}

	// try to unmarshal data into FixedAmountPromotion
	err = newStrictDecoder(data).Decode(&dst.FixedAmountPromotion)
	if err == nil {
		jsonFixedAmountPromotion, _ := json.Marshal(dst.FixedAmountPromotion)
		if string(jsonFixedAmountPromotion) == "{}" { // empty struct
			dst.FixedAmountPromotion = nil
		} else {
			match++
		}
	} else {
		dst.FixedAmountPromotion = nil
	}

	// try to unmarshal data into FixedPricePromotion
	err = newStrictDecoder(data).Decode(&dst.FixedPricePromotion)
	if err == nil {
		jsonFixedPricePromotion, _ := json.Marshal(dst.FixedPricePromotion)
		if string(jsonFixedPricePromotion) == "{}" { // empty struct
			dst.FixedPricePromotion = nil
		} else {
			match++
		}
	} else {
		dst.FixedPricePromotion = nil
	}

	// try to unmarshal data into FreeGiftPromotion
	err = newStrictDecoder(data).Decode(&dst.FreeGiftPromotion)
	if err == nil {
		jsonFreeGiftPromotion, _ := json.Marshal(dst.FreeGiftPromotion)
		if string(jsonFreeGiftPromotion) == "{}" { // empty struct
			dst.FreeGiftPromotion = nil
		} else {
			match++
		}
	} else {
		dst.FreeGiftPromotion = nil
	}

	// try to unmarshal data into FreeShippingPromotion
	err = newStrictDecoder(data).Decode(&dst.FreeShippingPromotion)
	if err == nil {
		jsonFreeShippingPromotion, _ := json.Marshal(dst.FreeShippingPromotion)
		if string(jsonFreeShippingPromotion) == "{}" { // empty struct
			dst.FreeShippingPromotion = nil
		} else {
			match++
		}
	} else {
		dst.FreeShippingPromotion = nil
	}

	// try to unmarshal data into PercentageDiscountPromotion
	err = newStrictDecoder(data).Decode(&dst.PercentageDiscountPromotion)
	if err == nil {
		jsonPercentageDiscountPromotion, _ := json.Marshal(dst.PercentageDiscountPromotion)
		if string(jsonPercentageDiscountPromotion) == "{}" { // empty struct
			dst.PercentageDiscountPromotion = nil
		} else {
			match++
		}
	} else {
		dst.PercentageDiscountPromotion = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ExternalPromotion = nil
		dst.FixedAmountPromotion = nil
		dst.FixedPricePromotion = nil
		dst.FreeGiftPromotion = nil
		dst.FreeShippingPromotion = nil
		dst.PercentageDiscountPromotion = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(CouponCodesPromotionRuleCreateDataRelationshipsPromotion)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(CouponCodesPromotionRuleCreateDataRelationshipsPromotion)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CouponCodesPromotionRuleCreateDataRelationshipsPromotion) MarshalJSON() ([]byte, error) {
	if src.ExternalPromotion != nil {
		return json.Marshal(&src.ExternalPromotion)
	}

	if src.FixedAmountPromotion != nil {
		return json.Marshal(&src.FixedAmountPromotion)
	}

	if src.FixedPricePromotion != nil {
		return json.Marshal(&src.FixedPricePromotion)
	}

	if src.FreeGiftPromotion != nil {
		return json.Marshal(&src.FreeGiftPromotion)
	}

	if src.FreeShippingPromotion != nil {
		return json.Marshal(&src.FreeShippingPromotion)
	}

	if src.PercentageDiscountPromotion != nil {
		return json.Marshal(&src.PercentageDiscountPromotion)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CouponCodesPromotionRuleCreateDataRelationshipsPromotion) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ExternalPromotion != nil {
		return obj.ExternalPromotion
	}

	if obj.FixedAmountPromotion != nil {
		return obj.FixedAmountPromotion
	}

	if obj.FixedPricePromotion != nil {
		return obj.FixedPricePromotion
	}

	if obj.FreeGiftPromotion != nil {
		return obj.FreeGiftPromotion
	}

	if obj.FreeShippingPromotion != nil {
		return obj.FreeShippingPromotion
	}

	if obj.PercentageDiscountPromotion != nil {
		return obj.PercentageDiscountPromotion
	}

	// all schemas are nil
	return nil
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
