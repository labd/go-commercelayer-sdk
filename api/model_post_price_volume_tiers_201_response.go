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

// POSTPriceVolumeTiers201Response struct for POSTPriceVolumeTiers201Response
type POSTPriceVolumeTiers201Response struct {
	Data *POSTPriceVolumeTiers201ResponseData `json:"data,omitempty"`
}

// NewPOSTPriceVolumeTiers201Response instantiates a new POSTPriceVolumeTiers201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPriceVolumeTiers201Response() *POSTPriceVolumeTiers201Response {
	this := POSTPriceVolumeTiers201Response{}
	return &this
}

// NewPOSTPriceVolumeTiers201ResponseWithDefaults instantiates a new POSTPriceVolumeTiers201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPriceVolumeTiers201ResponseWithDefaults() *POSTPriceVolumeTiers201Response {
	this := POSTPriceVolumeTiers201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTPriceVolumeTiers201Response) GetData() POSTPriceVolumeTiers201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTPriceVolumeTiers201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPriceVolumeTiers201Response) GetDataOk() (*POSTPriceVolumeTiers201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTPriceVolumeTiers201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTPriceVolumeTiers201ResponseData and assigns it to the Data field.
func (o *POSTPriceVolumeTiers201Response) SetData(v POSTPriceVolumeTiers201ResponseData) {
	o.Data = &v
}

func (o POSTPriceVolumeTiers201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTPriceVolumeTiers201Response struct {
	value *POSTPriceVolumeTiers201Response
	isSet bool
}

func (v NullablePOSTPriceVolumeTiers201Response) Get() *POSTPriceVolumeTiers201Response {
	return v.value
}

func (v *NullablePOSTPriceVolumeTiers201Response) Set(val *POSTPriceVolumeTiers201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPriceVolumeTiers201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPriceVolumeTiers201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPriceVolumeTiers201Response(val *POSTPriceVolumeTiers201Response) *NullablePOSTPriceVolumeTiers201Response {
	return &NullablePOSTPriceVolumeTiers201Response{value: val, isSet: true}
}

func (v NullablePOSTPriceVolumeTiers201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPriceVolumeTiers201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
