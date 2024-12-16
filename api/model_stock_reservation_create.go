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

// checks if the StockReservationCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StockReservationCreate{}

// StockReservationCreate struct for StockReservationCreate
type StockReservationCreate struct {
	Data StockReservationCreateData `json:"data"`
}

// NewStockReservationCreate instantiates a new StockReservationCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockReservationCreate(data StockReservationCreateData) *StockReservationCreate {
	this := StockReservationCreate{}
	this.Data = data
	return &this
}

// NewStockReservationCreateWithDefaults instantiates a new StockReservationCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockReservationCreateWithDefaults() *StockReservationCreate {
	this := StockReservationCreate{}
	return &this
}

// GetData returns the Data field value
func (o *StockReservationCreate) GetData() StockReservationCreateData {
	if o == nil {
		var ret StockReservationCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *StockReservationCreate) GetDataOk() (*StockReservationCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *StockReservationCreate) SetData(v StockReservationCreateData) {
	o.Data = v
}

func (o StockReservationCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StockReservationCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableStockReservationCreate struct {
	value *StockReservationCreate
	isSet bool
}

func (v NullableStockReservationCreate) Get() *StockReservationCreate {
	return v.value
}

func (v *NullableStockReservationCreate) Set(val *StockReservationCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableStockReservationCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableStockReservationCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockReservationCreate(val *StockReservationCreate) *NullableStockReservationCreate {
	return &NullableStockReservationCreate{value: val, isSet: true}
}

func (v NullableStockReservationCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockReservationCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
