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

// OrderAmountPromotionRuleResponse struct for OrderAmountPromotionRuleResponse
type OrderAmountPromotionRuleResponse struct {
	Data *OrderAmountPromotionRuleResponseData `json:"data,omitempty"`
}

// NewOrderAmountPromotionRuleResponse instantiates a new OrderAmountPromotionRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderAmountPromotionRuleResponse() *OrderAmountPromotionRuleResponse {
	this := OrderAmountPromotionRuleResponse{}
	return &this
}

// NewOrderAmountPromotionRuleResponseWithDefaults instantiates a new OrderAmountPromotionRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderAmountPromotionRuleResponseWithDefaults() *OrderAmountPromotionRuleResponse {
	this := OrderAmountPromotionRuleResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OrderAmountPromotionRuleResponse) GetData() OrderAmountPromotionRuleResponseData {
	if o == nil || o.Data == nil {
		var ret OrderAmountPromotionRuleResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmountPromotionRuleResponse) GetDataOk() (*OrderAmountPromotionRuleResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OrderAmountPromotionRuleResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given OrderAmountPromotionRuleResponseData and assigns it to the Data field.
func (o *OrderAmountPromotionRuleResponse) SetData(v OrderAmountPromotionRuleResponseData) {
	o.Data = &v
}

func (o OrderAmountPromotionRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableOrderAmountPromotionRuleResponse struct {
	value *OrderAmountPromotionRuleResponse
	isSet bool
}

func (v NullableOrderAmountPromotionRuleResponse) Get() *OrderAmountPromotionRuleResponse {
	return v.value
}

func (v *NullableOrderAmountPromotionRuleResponse) Set(val *OrderAmountPromotionRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderAmountPromotionRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderAmountPromotionRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderAmountPromotionRuleResponse(val *OrderAmountPromotionRuleResponse) *NullableOrderAmountPromotionRuleResponse {
	return &NullableOrderAmountPromotionRuleResponse{value: val, isSet: true}
}

func (v NullableOrderAmountPromotionRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderAmountPromotionRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
