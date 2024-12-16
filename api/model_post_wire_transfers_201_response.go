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

// checks if the POSTWireTransfers201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTWireTransfers201Response{}

// POSTWireTransfers201Response struct for POSTWireTransfers201Response
type POSTWireTransfers201Response struct {
	Data *POSTWireTransfers201ResponseData `json:"data,omitempty"`
}

// NewPOSTWireTransfers201Response instantiates a new POSTWireTransfers201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTWireTransfers201Response() *POSTWireTransfers201Response {
	this := POSTWireTransfers201Response{}
	return &this
}

// NewPOSTWireTransfers201ResponseWithDefaults instantiates a new POSTWireTransfers201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTWireTransfers201ResponseWithDefaults() *POSTWireTransfers201Response {
	this := POSTWireTransfers201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTWireTransfers201Response) GetData() POSTWireTransfers201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTWireTransfers201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTWireTransfers201Response) GetDataOk() (*POSTWireTransfers201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTWireTransfers201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTWireTransfers201ResponseData and assigns it to the Data field.
func (o *POSTWireTransfers201Response) SetData(v POSTWireTransfers201ResponseData) {
	o.Data = &v
}

func (o POSTWireTransfers201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTWireTransfers201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTWireTransfers201Response struct {
	value *POSTWireTransfers201Response
	isSet bool
}

func (v NullablePOSTWireTransfers201Response) Get() *POSTWireTransfers201Response {
	return v.value
}

func (v *NullablePOSTWireTransfers201Response) Set(val *POSTWireTransfers201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTWireTransfers201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTWireTransfers201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTWireTransfers201Response(val *POSTWireTransfers201Response) *NullablePOSTWireTransfers201Response {
	return &NullablePOSTWireTransfers201Response{value: val, isSet: true}
}

func (v NullablePOSTWireTransfers201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTWireTransfers201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
