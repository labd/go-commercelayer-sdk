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

// checks if the CustomPromotionRuleDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomPromotionRuleDataRelationships{}

// CustomPromotionRuleDataRelationships struct for CustomPromotionRuleDataRelationships
type CustomPromotionRuleDataRelationships struct {
	Promotion *CouponCodesPromotionRuleDataRelationshipsPromotion `json:"promotion,omitempty"`
	Versions  *AddressDataRelationshipsVersions                   `json:"versions,omitempty"`
}

// NewCustomPromotionRuleDataRelationships instantiates a new CustomPromotionRuleDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomPromotionRuleDataRelationships() *CustomPromotionRuleDataRelationships {
	this := CustomPromotionRuleDataRelationships{}
	return &this
}

// NewCustomPromotionRuleDataRelationshipsWithDefaults instantiates a new CustomPromotionRuleDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomPromotionRuleDataRelationshipsWithDefaults() *CustomPromotionRuleDataRelationships {
	this := CustomPromotionRuleDataRelationships{}
	return &this
}

// GetPromotion returns the Promotion field value if set, zero value otherwise.
func (o *CustomPromotionRuleDataRelationships) GetPromotion() CouponCodesPromotionRuleDataRelationshipsPromotion {
	if o == nil || IsNil(o.Promotion) {
		var ret CouponCodesPromotionRuleDataRelationshipsPromotion
		return ret
	}
	return *o.Promotion
}

// GetPromotionOk returns a tuple with the Promotion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomPromotionRuleDataRelationships) GetPromotionOk() (*CouponCodesPromotionRuleDataRelationshipsPromotion, bool) {
	if o == nil || IsNil(o.Promotion) {
		return nil, false
	}
	return o.Promotion, true
}

// HasPromotion returns a boolean if a field has been set.
func (o *CustomPromotionRuleDataRelationships) HasPromotion() bool {
	if o != nil && !IsNil(o.Promotion) {
		return true
	}

	return false
}

// SetPromotion gets a reference to the given CouponCodesPromotionRuleDataRelationshipsPromotion and assigns it to the Promotion field.
func (o *CustomPromotionRuleDataRelationships) SetPromotion(v CouponCodesPromotionRuleDataRelationshipsPromotion) {
	o.Promotion = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *CustomPromotionRuleDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomPromotionRuleDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *CustomPromotionRuleDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *CustomPromotionRuleDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

func (o CustomPromotionRuleDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomPromotionRuleDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Promotion) {
		toSerialize["promotion"] = o.Promotion
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullableCustomPromotionRuleDataRelationships struct {
	value *CustomPromotionRuleDataRelationships
	isSet bool
}

func (v NullableCustomPromotionRuleDataRelationships) Get() *CustomPromotionRuleDataRelationships {
	return v.value
}

func (v *NullableCustomPromotionRuleDataRelationships) Set(val *CustomPromotionRuleDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomPromotionRuleDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomPromotionRuleDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomPromotionRuleDataRelationships(val *CustomPromotionRuleDataRelationships) *NullableCustomPromotionRuleDataRelationships {
	return &NullableCustomPromotionRuleDataRelationships{value: val, isSet: true}
}

func (v NullableCustomPromotionRuleDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomPromotionRuleDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
