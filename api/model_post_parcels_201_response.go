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

// POSTParcels201Response struct for POSTParcels201Response
type POSTParcels201Response struct {
	Data *POSTParcels201ResponseData `json:"data,omitempty"`
}

// NewPOSTParcels201Response instantiates a new POSTParcels201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTParcels201Response() *POSTParcels201Response {
	this := POSTParcels201Response{}
	return &this
}

// NewPOSTParcels201ResponseWithDefaults instantiates a new POSTParcels201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTParcels201ResponseWithDefaults() *POSTParcels201Response {
	this := POSTParcels201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTParcels201Response) GetData() POSTParcels201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTParcels201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTParcels201Response) GetDataOk() (*POSTParcels201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTParcels201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTParcels201ResponseData and assigns it to the Data field.
func (o *POSTParcels201Response) SetData(v POSTParcels201ResponseData) {
	o.Data = &v
}

func (o POSTParcels201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTParcels201Response struct {
	value *POSTParcels201Response
	isSet bool
}

func (v NullablePOSTParcels201Response) Get() *POSTParcels201Response {
	return v.value
}

func (v *NullablePOSTParcels201Response) Set(val *POSTParcels201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTParcels201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTParcels201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTParcels201Response(val *POSTParcels201Response) *NullablePOSTParcels201Response {
	return &NullablePOSTParcels201Response{value: val, isSet: true}
}

func (v NullablePOSTParcels201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTParcels201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}