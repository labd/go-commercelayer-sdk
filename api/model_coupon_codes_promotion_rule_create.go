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

// checks if the CouponCodesPromotionRuleCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CouponCodesPromotionRuleCreate{}

// CouponCodesPromotionRuleCreate struct for CouponCodesPromotionRuleCreate
type CouponCodesPromotionRuleCreate struct {
	Data CouponCodesPromotionRuleCreateData `json:"data"`
}

// NewCouponCodesPromotionRuleCreate instantiates a new CouponCodesPromotionRuleCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponCodesPromotionRuleCreate(data CouponCodesPromotionRuleCreateData) *CouponCodesPromotionRuleCreate {
	this := CouponCodesPromotionRuleCreate{}
	this.Data = data
	return &this
}

// NewCouponCodesPromotionRuleCreateWithDefaults instantiates a new CouponCodesPromotionRuleCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponCodesPromotionRuleCreateWithDefaults() *CouponCodesPromotionRuleCreate {
	this := CouponCodesPromotionRuleCreate{}
	return &this
}

// GetData returns the Data field value
func (o *CouponCodesPromotionRuleCreate) GetData() CouponCodesPromotionRuleCreateData {
	if o == nil {
		var ret CouponCodesPromotionRuleCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CouponCodesPromotionRuleCreate) GetDataOk() (*CouponCodesPromotionRuleCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CouponCodesPromotionRuleCreate) SetData(v CouponCodesPromotionRuleCreateData) {
	o.Data = v
}

func (o CouponCodesPromotionRuleCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CouponCodesPromotionRuleCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableCouponCodesPromotionRuleCreate struct {
	value *CouponCodesPromotionRuleCreate
	isSet bool
}

func (v NullableCouponCodesPromotionRuleCreate) Get() *CouponCodesPromotionRuleCreate {
	return v.value
}

func (v *NullableCouponCodesPromotionRuleCreate) Set(val *CouponCodesPromotionRuleCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponCodesPromotionRuleCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponCodesPromotionRuleCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponCodesPromotionRuleCreate(val *CouponCodesPromotionRuleCreate) *NullableCouponCodesPromotionRuleCreate {
	return &NullableCouponCodesPromotionRuleCreate{value: val, isSet: true}
}

func (v NullableCouponCodesPromotionRuleCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponCodesPromotionRuleCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
