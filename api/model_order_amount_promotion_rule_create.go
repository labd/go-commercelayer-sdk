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

// checks if the OrderAmountPromotionRuleCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderAmountPromotionRuleCreate{}

// OrderAmountPromotionRuleCreate struct for OrderAmountPromotionRuleCreate
type OrderAmountPromotionRuleCreate struct {
	Data OrderAmountPromotionRuleCreateData `json:"data"`
}

// NewOrderAmountPromotionRuleCreate instantiates a new OrderAmountPromotionRuleCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderAmountPromotionRuleCreate(data OrderAmountPromotionRuleCreateData) *OrderAmountPromotionRuleCreate {
	this := OrderAmountPromotionRuleCreate{}
	this.Data = data
	return &this
}

// NewOrderAmountPromotionRuleCreateWithDefaults instantiates a new OrderAmountPromotionRuleCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderAmountPromotionRuleCreateWithDefaults() *OrderAmountPromotionRuleCreate {
	this := OrderAmountPromotionRuleCreate{}
	return &this
}

// GetData returns the Data field value
func (o *OrderAmountPromotionRuleCreate) GetData() OrderAmountPromotionRuleCreateData {
	if o == nil {
		var ret OrderAmountPromotionRuleCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OrderAmountPromotionRuleCreate) GetDataOk() (*OrderAmountPromotionRuleCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *OrderAmountPromotionRuleCreate) SetData(v OrderAmountPromotionRuleCreateData) {
	o.Data = v
}

func (o OrderAmountPromotionRuleCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderAmountPromotionRuleCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableOrderAmountPromotionRuleCreate struct {
	value *OrderAmountPromotionRuleCreate
	isSet bool
}

func (v NullableOrderAmountPromotionRuleCreate) Get() *OrderAmountPromotionRuleCreate {
	return v.value
}

func (v *NullableOrderAmountPromotionRuleCreate) Set(val *OrderAmountPromotionRuleCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderAmountPromotionRuleCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderAmountPromotionRuleCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderAmountPromotionRuleCreate(val *OrderAmountPromotionRuleCreate) *NullableOrderAmountPromotionRuleCreate {
	return &NullableOrderAmountPromotionRuleCreate{value: val, isSet: true}
}

func (v NullableOrderAmountPromotionRuleCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderAmountPromotionRuleCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
