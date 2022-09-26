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

// StockTransferResponse struct for StockTransferResponse
type StockTransferResponse struct {
	Data *StockTransferResponseData `json:"data,omitempty"`
}

// NewStockTransferResponse instantiates a new StockTransferResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockTransferResponse() *StockTransferResponse {
	this := StockTransferResponse{}
	return &this
}

// NewStockTransferResponseWithDefaults instantiates a new StockTransferResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockTransferResponseWithDefaults() *StockTransferResponse {
	this := StockTransferResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *StockTransferResponse) GetData() StockTransferResponseData {
	if o == nil || o.Data == nil {
		var ret StockTransferResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockTransferResponse) GetDataOk() (*StockTransferResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *StockTransferResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given StockTransferResponseData and assigns it to the Data field.
func (o *StockTransferResponse) SetData(v StockTransferResponseData) {
	o.Data = &v
}

func (o StockTransferResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableStockTransferResponse struct {
	value *StockTransferResponse
	isSet bool
}

func (v NullableStockTransferResponse) Get() *StockTransferResponse {
	return v.value
}

func (v *NullableStockTransferResponse) Set(val *StockTransferResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStockTransferResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStockTransferResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockTransferResponse(val *StockTransferResponse) *NullableStockTransferResponse {
	return &NullableStockTransferResponse{value: val, isSet: true}
}

func (v NullableStockTransferResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockTransferResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
