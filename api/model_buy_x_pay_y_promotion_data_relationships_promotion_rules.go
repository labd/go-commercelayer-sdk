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

// checks if the BuyXPayYPromotionDataRelationshipsPromotionRules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuyXPayYPromotionDataRelationshipsPromotionRules{}

// BuyXPayYPromotionDataRelationshipsPromotionRules struct for BuyXPayYPromotionDataRelationshipsPromotionRules
type BuyXPayYPromotionDataRelationshipsPromotionRules struct {
	Data *BuyXPayYPromotionDataRelationshipsPromotionRulesData `json:"data,omitempty"`
}

// NewBuyXPayYPromotionDataRelationshipsPromotionRules instantiates a new BuyXPayYPromotionDataRelationshipsPromotionRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuyXPayYPromotionDataRelationshipsPromotionRules() *BuyXPayYPromotionDataRelationshipsPromotionRules {
	this := BuyXPayYPromotionDataRelationshipsPromotionRules{}
	return &this
}

// NewBuyXPayYPromotionDataRelationshipsPromotionRulesWithDefaults instantiates a new BuyXPayYPromotionDataRelationshipsPromotionRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuyXPayYPromotionDataRelationshipsPromotionRulesWithDefaults() *BuyXPayYPromotionDataRelationshipsPromotionRules {
	this := BuyXPayYPromotionDataRelationshipsPromotionRules{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BuyXPayYPromotionDataRelationshipsPromotionRules) GetData() BuyXPayYPromotionDataRelationshipsPromotionRulesData {
	if o == nil || IsNil(o.Data) {
		var ret BuyXPayYPromotionDataRelationshipsPromotionRulesData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyXPayYPromotionDataRelationshipsPromotionRules) GetDataOk() (*BuyXPayYPromotionDataRelationshipsPromotionRulesData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BuyXPayYPromotionDataRelationshipsPromotionRules) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given BuyXPayYPromotionDataRelationshipsPromotionRulesData and assigns it to the Data field.
func (o *BuyXPayYPromotionDataRelationshipsPromotionRules) SetData(v BuyXPayYPromotionDataRelationshipsPromotionRulesData) {
	o.Data = &v
}

func (o BuyXPayYPromotionDataRelationshipsPromotionRules) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuyXPayYPromotionDataRelationshipsPromotionRules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableBuyXPayYPromotionDataRelationshipsPromotionRules struct {
	value *BuyXPayYPromotionDataRelationshipsPromotionRules
	isSet bool
}

func (v NullableBuyXPayYPromotionDataRelationshipsPromotionRules) Get() *BuyXPayYPromotionDataRelationshipsPromotionRules {
	return v.value
}

func (v *NullableBuyXPayYPromotionDataRelationshipsPromotionRules) Set(val *BuyXPayYPromotionDataRelationshipsPromotionRules) {
	v.value = val
	v.isSet = true
}

func (v NullableBuyXPayYPromotionDataRelationshipsPromotionRules) IsSet() bool {
	return v.isSet
}

func (v *NullableBuyXPayYPromotionDataRelationshipsPromotionRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuyXPayYPromotionDataRelationshipsPromotionRules(val *BuyXPayYPromotionDataRelationshipsPromotionRules) *NullableBuyXPayYPromotionDataRelationshipsPromotionRules {
	return &NullableBuyXPayYPromotionDataRelationshipsPromotionRules{value: val, isSet: true}
}

func (v NullableBuyXPayYPromotionDataRelationshipsPromotionRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuyXPayYPromotionDataRelationshipsPromotionRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
