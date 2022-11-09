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

// StockLineItem struct for StockLineItem
type StockLineItem struct {
	Data StockLineItemData `json:"data"`
}

// NewStockLineItem instantiates a new StockLineItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockLineItem(data StockLineItemData) *StockLineItem {
	this := StockLineItem{}
	this.Data = data
	return &this
}

// NewStockLineItemWithDefaults instantiates a new StockLineItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockLineItemWithDefaults() *StockLineItem {
	this := StockLineItem{}
	return &this
}

// GetData returns the Data field value
func (o *StockLineItem) GetData() StockLineItemData {
	if o == nil {
		var ret StockLineItemData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *StockLineItem) GetDataOk() (*StockLineItemData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *StockLineItem) SetData(v StockLineItemData) {
	o.Data = v
}

func (o StockLineItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableStockLineItem struct {
	value *StockLineItem
	isSet bool
}

func (v NullableStockLineItem) Get() *StockLineItem {
	return v.value
}

func (v *NullableStockLineItem) Set(val *StockLineItem) {
	v.value = val
	v.isSet = true
}

func (v NullableStockLineItem) IsSet() bool {
	return v.isSet
}

func (v *NullableStockLineItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockLineItem(val *StockLineItem) *NullableStockLineItem {
	return &NullableStockLineItem{value: val, isSet: true}
}

func (v NullableStockLineItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockLineItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
