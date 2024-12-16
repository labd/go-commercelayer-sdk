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

// checks if the BuyXPayYPromotionDataRelationshipsCoupons type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuyXPayYPromotionDataRelationshipsCoupons{}

// BuyXPayYPromotionDataRelationshipsCoupons struct for BuyXPayYPromotionDataRelationshipsCoupons
type BuyXPayYPromotionDataRelationshipsCoupons struct {
	Data *BuyXPayYPromotionDataRelationshipsCouponsData `json:"data,omitempty"`
}

// NewBuyXPayYPromotionDataRelationshipsCoupons instantiates a new BuyXPayYPromotionDataRelationshipsCoupons object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuyXPayYPromotionDataRelationshipsCoupons() *BuyXPayYPromotionDataRelationshipsCoupons {
	this := BuyXPayYPromotionDataRelationshipsCoupons{}
	return &this
}

// NewBuyXPayYPromotionDataRelationshipsCouponsWithDefaults instantiates a new BuyXPayYPromotionDataRelationshipsCoupons object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuyXPayYPromotionDataRelationshipsCouponsWithDefaults() *BuyXPayYPromotionDataRelationshipsCoupons {
	this := BuyXPayYPromotionDataRelationshipsCoupons{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BuyXPayYPromotionDataRelationshipsCoupons) GetData() BuyXPayYPromotionDataRelationshipsCouponsData {
	if o == nil || IsNil(o.Data) {
		var ret BuyXPayYPromotionDataRelationshipsCouponsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyXPayYPromotionDataRelationshipsCoupons) GetDataOk() (*BuyXPayYPromotionDataRelationshipsCouponsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BuyXPayYPromotionDataRelationshipsCoupons) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given BuyXPayYPromotionDataRelationshipsCouponsData and assigns it to the Data field.
func (o *BuyXPayYPromotionDataRelationshipsCoupons) SetData(v BuyXPayYPromotionDataRelationshipsCouponsData) {
	o.Data = &v
}

func (o BuyXPayYPromotionDataRelationshipsCoupons) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuyXPayYPromotionDataRelationshipsCoupons) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableBuyXPayYPromotionDataRelationshipsCoupons struct {
	value *BuyXPayYPromotionDataRelationshipsCoupons
	isSet bool
}

func (v NullableBuyXPayYPromotionDataRelationshipsCoupons) Get() *BuyXPayYPromotionDataRelationshipsCoupons {
	return v.value
}

func (v *NullableBuyXPayYPromotionDataRelationshipsCoupons) Set(val *BuyXPayYPromotionDataRelationshipsCoupons) {
	v.value = val
	v.isSet = true
}

func (v NullableBuyXPayYPromotionDataRelationshipsCoupons) IsSet() bool {
	return v.isSet
}

func (v *NullableBuyXPayYPromotionDataRelationshipsCoupons) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuyXPayYPromotionDataRelationshipsCoupons(val *BuyXPayYPromotionDataRelationshipsCoupons) *NullableBuyXPayYPromotionDataRelationshipsCoupons {
	return &NullableBuyXPayYPromotionDataRelationshipsCoupons{value: val, isSet: true}
}

func (v NullableBuyXPayYPromotionDataRelationshipsCoupons) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuyXPayYPromotionDataRelationshipsCoupons) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
