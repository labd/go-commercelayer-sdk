/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// POSTAvalaraAccounts201Response struct for POSTAvalaraAccounts201Response
type POSTAvalaraAccounts201Response struct {
	Data *POSTAvalaraAccounts201ResponseData `json:"data,omitempty"`
}

// NewPOSTAvalaraAccounts201Response instantiates a new POSTAvalaraAccounts201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAvalaraAccounts201Response() *POSTAvalaraAccounts201Response {
	this := POSTAvalaraAccounts201Response{}
	return &this
}

// NewPOSTAvalaraAccounts201ResponseWithDefaults instantiates a new POSTAvalaraAccounts201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAvalaraAccounts201ResponseWithDefaults() *POSTAvalaraAccounts201Response {
	this := POSTAvalaraAccounts201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTAvalaraAccounts201Response) GetData() POSTAvalaraAccounts201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTAvalaraAccounts201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAvalaraAccounts201Response) GetDataOk() (*POSTAvalaraAccounts201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTAvalaraAccounts201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTAvalaraAccounts201ResponseData and assigns it to the Data field.
func (o *POSTAvalaraAccounts201Response) SetData(v POSTAvalaraAccounts201ResponseData) {
	o.Data = &v
}

func (o POSTAvalaraAccounts201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTAvalaraAccounts201Response struct {
	value *POSTAvalaraAccounts201Response
	isSet bool
}

func (v NullablePOSTAvalaraAccounts201Response) Get() *POSTAvalaraAccounts201Response {
	return v.value
}

func (v *NullablePOSTAvalaraAccounts201Response) Set(val *POSTAvalaraAccounts201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAvalaraAccounts201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAvalaraAccounts201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAvalaraAccounts201Response(val *POSTAvalaraAccounts201Response) *NullablePOSTAvalaraAccounts201Response {
	return &NullablePOSTAvalaraAccounts201Response{value: val, isSet: true}
}

func (v NullablePOSTAvalaraAccounts201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAvalaraAccounts201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
