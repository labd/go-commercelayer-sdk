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

// PriceListResponseList struct for PriceListResponseList
type PriceListResponseList struct {
	Data []Data `json:"data,omitempty"`
}

// NewPriceListResponseList instantiates a new PriceListResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceListResponseList() *PriceListResponseList {
	this := PriceListResponseList{}
	return &this
}

// NewPriceListResponseListWithDefaults instantiates a new PriceListResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceListResponseListWithDefaults() *PriceListResponseList {
	this := PriceListResponseList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PriceListResponseList) GetData() []Data {
	if o == nil || o.Data == nil {
		var ret []Data
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceListResponseList) GetDataOk() ([]Data, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PriceListResponseList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Data and assigns it to the Data field.
func (o *PriceListResponseList) SetData(v []Data) {
	o.Data = v
}

func (o PriceListResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePriceListResponseList struct {
	value *PriceListResponseList
	isSet bool
}

func (v NullablePriceListResponseList) Get() *PriceListResponseList {
	return v.value
}

func (v *NullablePriceListResponseList) Set(val *PriceListResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceListResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceListResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceListResponseList(val *PriceListResponseList) *NullablePriceListResponseList {
	return &NullablePriceListResponseList{value: val, isSet: true}
}

func (v NullablePriceListResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceListResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
