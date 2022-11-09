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

// POSTAdyenPayments201Response struct for POSTAdyenPayments201Response
type POSTAdyenPayments201Response struct {
	Data *POSTAdyenPayments201ResponseData `json:"data,omitempty"`
}

// NewPOSTAdyenPayments201Response instantiates a new POSTAdyenPayments201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAdyenPayments201Response() *POSTAdyenPayments201Response {
	this := POSTAdyenPayments201Response{}
	return &this
}

// NewPOSTAdyenPayments201ResponseWithDefaults instantiates a new POSTAdyenPayments201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAdyenPayments201ResponseWithDefaults() *POSTAdyenPayments201Response {
	this := POSTAdyenPayments201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTAdyenPayments201Response) GetData() POSTAdyenPayments201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTAdyenPayments201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAdyenPayments201Response) GetDataOk() (*POSTAdyenPayments201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTAdyenPayments201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTAdyenPayments201ResponseData and assigns it to the Data field.
func (o *POSTAdyenPayments201Response) SetData(v POSTAdyenPayments201ResponseData) {
	o.Data = &v
}

func (o POSTAdyenPayments201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTAdyenPayments201Response struct {
	value *POSTAdyenPayments201Response
	isSet bool
}

func (v NullablePOSTAdyenPayments201Response) Get() *POSTAdyenPayments201Response {
	return v.value
}

func (v *NullablePOSTAdyenPayments201Response) Set(val *POSTAdyenPayments201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAdyenPayments201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAdyenPayments201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAdyenPayments201Response(val *POSTAdyenPayments201Response) *NullablePOSTAdyenPayments201Response {
	return &NullablePOSTAdyenPayments201Response{value: val, isSet: true}
}

func (v NullablePOSTAdyenPayments201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAdyenPayments201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}