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

// checks if the POSTPriceListSchedulers201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTPriceListSchedulers201Response{}

// POSTPriceListSchedulers201Response struct for POSTPriceListSchedulers201Response
type POSTPriceListSchedulers201Response struct {
	Data *POSTPriceListSchedulers201ResponseData `json:"data,omitempty"`
}

// NewPOSTPriceListSchedulers201Response instantiates a new POSTPriceListSchedulers201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPriceListSchedulers201Response() *POSTPriceListSchedulers201Response {
	this := POSTPriceListSchedulers201Response{}
	return &this
}

// NewPOSTPriceListSchedulers201ResponseWithDefaults instantiates a new POSTPriceListSchedulers201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPriceListSchedulers201ResponseWithDefaults() *POSTPriceListSchedulers201Response {
	this := POSTPriceListSchedulers201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTPriceListSchedulers201Response) GetData() POSTPriceListSchedulers201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTPriceListSchedulers201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPriceListSchedulers201Response) GetDataOk() (*POSTPriceListSchedulers201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTPriceListSchedulers201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTPriceListSchedulers201ResponseData and assigns it to the Data field.
func (o *POSTPriceListSchedulers201Response) SetData(v POSTPriceListSchedulers201ResponseData) {
	o.Data = &v
}

func (o POSTPriceListSchedulers201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTPriceListSchedulers201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTPriceListSchedulers201Response struct {
	value *POSTPriceListSchedulers201Response
	isSet bool
}

func (v NullablePOSTPriceListSchedulers201Response) Get() *POSTPriceListSchedulers201Response {
	return v.value
}

func (v *NullablePOSTPriceListSchedulers201Response) Set(val *POSTPriceListSchedulers201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPriceListSchedulers201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPriceListSchedulers201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPriceListSchedulers201Response(val *POSTPriceListSchedulers201Response) *NullablePOSTPriceListSchedulers201Response {
	return &NullablePOSTPriceListSchedulers201Response{value: val, isSet: true}
}

func (v NullablePOSTPriceListSchedulers201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPriceListSchedulers201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
