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

// checks if the POSTManualGateways201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTManualGateways201Response{}

// POSTManualGateways201Response struct for POSTManualGateways201Response
type POSTManualGateways201Response struct {
	Data *POSTManualGateways201ResponseData `json:"data,omitempty"`
}

// NewPOSTManualGateways201Response instantiates a new POSTManualGateways201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTManualGateways201Response() *POSTManualGateways201Response {
	this := POSTManualGateways201Response{}
	return &this
}

// NewPOSTManualGateways201ResponseWithDefaults instantiates a new POSTManualGateways201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTManualGateways201ResponseWithDefaults() *POSTManualGateways201Response {
	this := POSTManualGateways201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTManualGateways201Response) GetData() POSTManualGateways201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTManualGateways201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTManualGateways201Response) GetDataOk() (*POSTManualGateways201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTManualGateways201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTManualGateways201ResponseData and assigns it to the Data field.
func (o *POSTManualGateways201Response) SetData(v POSTManualGateways201ResponseData) {
	o.Data = &v
}

func (o POSTManualGateways201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTManualGateways201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTManualGateways201Response struct {
	value *POSTManualGateways201Response
	isSet bool
}

func (v NullablePOSTManualGateways201Response) Get() *POSTManualGateways201Response {
	return v.value
}

func (v *NullablePOSTManualGateways201Response) Set(val *POSTManualGateways201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTManualGateways201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTManualGateways201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTManualGateways201Response(val *POSTManualGateways201Response) *NullablePOSTManualGateways201Response {
	return &NullablePOSTManualGateways201Response{value: val, isSet: true}
}

func (v NullablePOSTManualGateways201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTManualGateways201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
