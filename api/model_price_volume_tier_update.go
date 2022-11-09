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

// PriceVolumeTierUpdate struct for PriceVolumeTierUpdate
type PriceVolumeTierUpdate struct {
	Data PriceVolumeTierUpdateData `json:"data"`
}

// NewPriceVolumeTierUpdate instantiates a new PriceVolumeTierUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceVolumeTierUpdate(data PriceVolumeTierUpdateData) *PriceVolumeTierUpdate {
	this := PriceVolumeTierUpdate{}
	this.Data = data
	return &this
}

// NewPriceVolumeTierUpdateWithDefaults instantiates a new PriceVolumeTierUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceVolumeTierUpdateWithDefaults() *PriceVolumeTierUpdate {
	this := PriceVolumeTierUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *PriceVolumeTierUpdate) GetData() PriceVolumeTierUpdateData {
	if o == nil {
		var ret PriceVolumeTierUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PriceVolumeTierUpdate) GetDataOk() (*PriceVolumeTierUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PriceVolumeTierUpdate) SetData(v PriceVolumeTierUpdateData) {
	o.Data = v
}

func (o PriceVolumeTierUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePriceVolumeTierUpdate struct {
	value *PriceVolumeTierUpdate
	isSet bool
}

func (v NullablePriceVolumeTierUpdate) Get() *PriceVolumeTierUpdate {
	return v.value
}

func (v *NullablePriceVolumeTierUpdate) Set(val *PriceVolumeTierUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceVolumeTierUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceVolumeTierUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceVolumeTierUpdate(val *PriceVolumeTierUpdate) *NullablePriceVolumeTierUpdate {
	return &NullablePriceVolumeTierUpdate{value: val, isSet: true}
}

func (v NullablePriceVolumeTierUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceVolumeTierUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
