/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the StockItemDataRelationshipsReservedStock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StockItemDataRelationshipsReservedStock{}

// StockItemDataRelationshipsReservedStock struct for StockItemDataRelationshipsReservedStock
type StockItemDataRelationshipsReservedStock struct {
	Data *StockItemDataRelationshipsReservedStockData `json:"data,omitempty"`
}

// NewStockItemDataRelationshipsReservedStock instantiates a new StockItemDataRelationshipsReservedStock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockItemDataRelationshipsReservedStock() *StockItemDataRelationshipsReservedStock {
	this := StockItemDataRelationshipsReservedStock{}
	return &this
}

// NewStockItemDataRelationshipsReservedStockWithDefaults instantiates a new StockItemDataRelationshipsReservedStock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockItemDataRelationshipsReservedStockWithDefaults() *StockItemDataRelationshipsReservedStock {
	this := StockItemDataRelationshipsReservedStock{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *StockItemDataRelationshipsReservedStock) GetData() StockItemDataRelationshipsReservedStockData {
	if o == nil || IsNil(o.Data) {
		var ret StockItemDataRelationshipsReservedStockData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockItemDataRelationshipsReservedStock) GetDataOk() (*StockItemDataRelationshipsReservedStockData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *StockItemDataRelationshipsReservedStock) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given StockItemDataRelationshipsReservedStockData and assigns it to the Data field.
func (o *StockItemDataRelationshipsReservedStock) SetData(v StockItemDataRelationshipsReservedStockData) {
	o.Data = &v
}

func (o StockItemDataRelationshipsReservedStock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StockItemDataRelationshipsReservedStock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableStockItemDataRelationshipsReservedStock struct {
	value *StockItemDataRelationshipsReservedStock
	isSet bool
}

func (v NullableStockItemDataRelationshipsReservedStock) Get() *StockItemDataRelationshipsReservedStock {
	return v.value
}

func (v *NullableStockItemDataRelationshipsReservedStock) Set(val *StockItemDataRelationshipsReservedStock) {
	v.value = val
	v.isSet = true
}

func (v NullableStockItemDataRelationshipsReservedStock) IsSet() bool {
	return v.isSet
}

func (v *NullableStockItemDataRelationshipsReservedStock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockItemDataRelationshipsReservedStock(val *StockItemDataRelationshipsReservedStock) *NullableStockItemDataRelationshipsReservedStock {
	return &NullableStockItemDataRelationshipsReservedStock{value: val, isSet: true}
}

func (v NullableStockItemDataRelationshipsReservedStock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockItemDataRelationshipsReservedStock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}