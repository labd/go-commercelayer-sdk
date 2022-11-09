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

// StockTransfer struct for StockTransfer
type StockTransfer struct {
	Data StockTransferData `json:"data"`
}

// NewStockTransfer instantiates a new StockTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockTransfer(data StockTransferData) *StockTransfer {
	this := StockTransfer{}
	this.Data = data
	return &this
}

// NewStockTransferWithDefaults instantiates a new StockTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockTransferWithDefaults() *StockTransfer {
	this := StockTransfer{}
	return &this
}

// GetData returns the Data field value
func (o *StockTransfer) GetData() StockTransferData {
	if o == nil {
		var ret StockTransferData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *StockTransfer) GetDataOk() (*StockTransferData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *StockTransfer) SetData(v StockTransferData) {
	o.Data = v
}

func (o StockTransfer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableStockTransfer struct {
	value *StockTransfer
	isSet bool
}

func (v NullableStockTransfer) Get() *StockTransfer {
	return v.value
}

func (v *NullableStockTransfer) Set(val *StockTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullableStockTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableStockTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockTransfer(val *StockTransfer) *NullableStockTransfer {
	return &NullableStockTransfer{value: val, isSet: true}
}

func (v NullableStockTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
