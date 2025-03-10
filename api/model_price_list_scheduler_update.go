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

// checks if the PriceListSchedulerUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceListSchedulerUpdate{}

// PriceListSchedulerUpdate struct for PriceListSchedulerUpdate
type PriceListSchedulerUpdate struct {
	Data PriceListSchedulerUpdateData `json:"data"`
}

// NewPriceListSchedulerUpdate instantiates a new PriceListSchedulerUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceListSchedulerUpdate(data PriceListSchedulerUpdateData) *PriceListSchedulerUpdate {
	this := PriceListSchedulerUpdate{}
	this.Data = data
	return &this
}

// NewPriceListSchedulerUpdateWithDefaults instantiates a new PriceListSchedulerUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceListSchedulerUpdateWithDefaults() *PriceListSchedulerUpdate {
	this := PriceListSchedulerUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *PriceListSchedulerUpdate) GetData() PriceListSchedulerUpdateData {
	if o == nil {
		var ret PriceListSchedulerUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PriceListSchedulerUpdate) GetDataOk() (*PriceListSchedulerUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PriceListSchedulerUpdate) SetData(v PriceListSchedulerUpdateData) {
	o.Data = v
}

func (o PriceListSchedulerUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceListSchedulerUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePriceListSchedulerUpdate struct {
	value *PriceListSchedulerUpdate
	isSet bool
}

func (v NullablePriceListSchedulerUpdate) Get() *PriceListSchedulerUpdate {
	return v.value
}

func (v *NullablePriceListSchedulerUpdate) Set(val *PriceListSchedulerUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceListSchedulerUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceListSchedulerUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceListSchedulerUpdate(val *PriceListSchedulerUpdate) *NullablePriceListSchedulerUpdate {
	return &NullablePriceListSchedulerUpdate{value: val, isSet: true}
}

func (v NullablePriceListSchedulerUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceListSchedulerUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
