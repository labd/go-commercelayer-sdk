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

// POSTTaxRules201Response struct for POSTTaxRules201Response
type POSTTaxRules201Response struct {
	Data *POSTTaxRules201ResponseData `json:"data,omitempty"`
}

// NewPOSTTaxRules201Response instantiates a new POSTTaxRules201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTTaxRules201Response() *POSTTaxRules201Response {
	this := POSTTaxRules201Response{}
	return &this
}

// NewPOSTTaxRules201ResponseWithDefaults instantiates a new POSTTaxRules201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTTaxRules201ResponseWithDefaults() *POSTTaxRules201Response {
	this := POSTTaxRules201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTTaxRules201Response) GetData() POSTTaxRules201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTTaxRules201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTTaxRules201Response) GetDataOk() (*POSTTaxRules201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTTaxRules201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTTaxRules201ResponseData and assigns it to the Data field.
func (o *POSTTaxRules201Response) SetData(v POSTTaxRules201ResponseData) {
	o.Data = &v
}

func (o POSTTaxRules201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTTaxRules201Response struct {
	value *POSTTaxRules201Response
	isSet bool
}

func (v NullablePOSTTaxRules201Response) Get() *POSTTaxRules201Response {
	return v.value
}

func (v *NullablePOSTTaxRules201Response) Set(val *POSTTaxRules201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTTaxRules201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTTaxRules201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTTaxRules201Response(val *POSTTaxRules201Response) *NullablePOSTTaxRules201Response {
	return &NullablePOSTTaxRules201Response{value: val, isSet: true}
}

func (v NullablePOSTTaxRules201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTTaxRules201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
