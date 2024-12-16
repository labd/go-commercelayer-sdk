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

// checks if the CustomPromotionRuleUpdateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomPromotionRuleUpdateDataRelationships{}

// CustomPromotionRuleUpdateDataRelationships struct for CustomPromotionRuleUpdateDataRelationships
type CustomPromotionRuleUpdateDataRelationships struct {
	Promotion *CouponCodesPromotionRuleCreateDataRelationshipsPromotion `json:"promotion,omitempty"`
}

// NewCustomPromotionRuleUpdateDataRelationships instantiates a new CustomPromotionRuleUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomPromotionRuleUpdateDataRelationships() *CustomPromotionRuleUpdateDataRelationships {
	this := CustomPromotionRuleUpdateDataRelationships{}
	return &this
}

// NewCustomPromotionRuleUpdateDataRelationshipsWithDefaults instantiates a new CustomPromotionRuleUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomPromotionRuleUpdateDataRelationshipsWithDefaults() *CustomPromotionRuleUpdateDataRelationships {
	this := CustomPromotionRuleUpdateDataRelationships{}
	return &this
}

// GetPromotion returns the Promotion field value if set, zero value otherwise.
func (o *CustomPromotionRuleUpdateDataRelationships) GetPromotion() CouponCodesPromotionRuleCreateDataRelationshipsPromotion {
	if o == nil || IsNil(o.Promotion) {
		var ret CouponCodesPromotionRuleCreateDataRelationshipsPromotion
		return ret
	}
	return *o.Promotion
}

// GetPromotionOk returns a tuple with the Promotion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomPromotionRuleUpdateDataRelationships) GetPromotionOk() (*CouponCodesPromotionRuleCreateDataRelationshipsPromotion, bool) {
	if o == nil || IsNil(o.Promotion) {
		return nil, false
	}
	return o.Promotion, true
}

// HasPromotion returns a boolean if a field has been set.
func (o *CustomPromotionRuleUpdateDataRelationships) HasPromotion() bool {
	if o != nil && !IsNil(o.Promotion) {
		return true
	}

	return false
}

// SetPromotion gets a reference to the given CouponCodesPromotionRuleCreateDataRelationshipsPromotion and assigns it to the Promotion field.
func (o *CustomPromotionRuleUpdateDataRelationships) SetPromotion(v CouponCodesPromotionRuleCreateDataRelationshipsPromotion) {
	o.Promotion = &v
}

func (o CustomPromotionRuleUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomPromotionRuleUpdateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Promotion) {
		toSerialize["promotion"] = o.Promotion
	}
	return toSerialize, nil
}

type NullableCustomPromotionRuleUpdateDataRelationships struct {
	value *CustomPromotionRuleUpdateDataRelationships
	isSet bool
}

func (v NullableCustomPromotionRuleUpdateDataRelationships) Get() *CustomPromotionRuleUpdateDataRelationships {
	return v.value
}

func (v *NullableCustomPromotionRuleUpdateDataRelationships) Set(val *CustomPromotionRuleUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomPromotionRuleUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomPromotionRuleUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomPromotionRuleUpdateDataRelationships(val *CustomPromotionRuleUpdateDataRelationships) *NullableCustomPromotionRuleUpdateDataRelationships {
	return &NullableCustomPromotionRuleUpdateDataRelationships{value: val, isSet: true}
}

func (v NullableCustomPromotionRuleUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomPromotionRuleUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
