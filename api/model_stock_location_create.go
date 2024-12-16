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

// checks if the StockLocationCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StockLocationCreate{}

// StockLocationCreate struct for StockLocationCreate
type StockLocationCreate struct {
	Data StockLocationCreateData `json:"data"`
}

// NewStockLocationCreate instantiates a new StockLocationCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockLocationCreate(data StockLocationCreateData) *StockLocationCreate {
	this := StockLocationCreate{}
	this.Data = data
	return &this
}

// NewStockLocationCreateWithDefaults instantiates a new StockLocationCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockLocationCreateWithDefaults() *StockLocationCreate {
	this := StockLocationCreate{}
	return &this
}

// GetData returns the Data field value
func (o *StockLocationCreate) GetData() StockLocationCreateData {
	if o == nil {
		var ret StockLocationCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *StockLocationCreate) GetDataOk() (*StockLocationCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *StockLocationCreate) SetData(v StockLocationCreateData) {
	o.Data = v
}

func (o StockLocationCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StockLocationCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableStockLocationCreate struct {
	value *StockLocationCreate
	isSet bool
}

func (v NullableStockLocationCreate) Get() *StockLocationCreate {
	return v.value
}

func (v *NullableStockLocationCreate) Set(val *StockLocationCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableStockLocationCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableStockLocationCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockLocationCreate(val *StockLocationCreate) *NullableStockLocationCreate {
	return &NullableStockLocationCreate{value: val, isSet: true}
}

func (v NullableStockLocationCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockLocationCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
