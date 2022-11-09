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

// AvalaraAccountDataRelationshipsMarkets struct for AvalaraAccountDataRelationshipsMarkets
type AvalaraAccountDataRelationshipsMarkets struct {
	Data AvalaraAccountDataRelationshipsMarketsData `json:"data"`
}

// NewAvalaraAccountDataRelationshipsMarkets instantiates a new AvalaraAccountDataRelationshipsMarkets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvalaraAccountDataRelationshipsMarkets(data AvalaraAccountDataRelationshipsMarketsData) *AvalaraAccountDataRelationshipsMarkets {
	this := AvalaraAccountDataRelationshipsMarkets{}
	this.Data = data
	return &this
}

// NewAvalaraAccountDataRelationshipsMarketsWithDefaults instantiates a new AvalaraAccountDataRelationshipsMarkets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvalaraAccountDataRelationshipsMarketsWithDefaults() *AvalaraAccountDataRelationshipsMarkets {
	this := AvalaraAccountDataRelationshipsMarkets{}
	return &this
}

// GetData returns the Data field value
func (o *AvalaraAccountDataRelationshipsMarkets) GetData() AvalaraAccountDataRelationshipsMarketsData {
	if o == nil {
		var ret AvalaraAccountDataRelationshipsMarketsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AvalaraAccountDataRelationshipsMarkets) GetDataOk() (*AvalaraAccountDataRelationshipsMarketsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AvalaraAccountDataRelationshipsMarkets) SetData(v AvalaraAccountDataRelationshipsMarketsData) {
	o.Data = v
}

func (o AvalaraAccountDataRelationshipsMarkets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAvalaraAccountDataRelationshipsMarkets struct {
	value *AvalaraAccountDataRelationshipsMarkets
	isSet bool
}

func (v NullableAvalaraAccountDataRelationshipsMarkets) Get() *AvalaraAccountDataRelationshipsMarkets {
	return v.value
}

func (v *NullableAvalaraAccountDataRelationshipsMarkets) Set(val *AvalaraAccountDataRelationshipsMarkets) {
	v.value = val
	v.isSet = true
}

func (v NullableAvalaraAccountDataRelationshipsMarkets) IsSet() bool {
	return v.isSet
}

func (v *NullableAvalaraAccountDataRelationshipsMarkets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvalaraAccountDataRelationshipsMarkets(val *AvalaraAccountDataRelationshipsMarkets) *NullableAvalaraAccountDataRelationshipsMarkets {
	return &NullableAvalaraAccountDataRelationshipsMarkets{value: val, isSet: true}
}

func (v NullableAvalaraAccountDataRelationshipsMarkets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvalaraAccountDataRelationshipsMarkets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
