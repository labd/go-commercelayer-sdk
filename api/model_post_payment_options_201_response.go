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

// checks if the POSTPaymentOptions201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTPaymentOptions201Response{}

// POSTPaymentOptions201Response struct for POSTPaymentOptions201Response
type POSTPaymentOptions201Response struct {
	Data *POSTPaymentOptions201ResponseData `json:"data,omitempty"`
}

// NewPOSTPaymentOptions201Response instantiates a new POSTPaymentOptions201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPaymentOptions201Response() *POSTPaymentOptions201Response {
	this := POSTPaymentOptions201Response{}
	return &this
}

// NewPOSTPaymentOptions201ResponseWithDefaults instantiates a new POSTPaymentOptions201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPaymentOptions201ResponseWithDefaults() *POSTPaymentOptions201Response {
	this := POSTPaymentOptions201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTPaymentOptions201Response) GetData() POSTPaymentOptions201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTPaymentOptions201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPaymentOptions201Response) GetDataOk() (*POSTPaymentOptions201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTPaymentOptions201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTPaymentOptions201ResponseData and assigns it to the Data field.
func (o *POSTPaymentOptions201Response) SetData(v POSTPaymentOptions201ResponseData) {
	o.Data = &v
}

func (o POSTPaymentOptions201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTPaymentOptions201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTPaymentOptions201Response struct {
	value *POSTPaymentOptions201Response
	isSet bool
}

func (v NullablePOSTPaymentOptions201Response) Get() *POSTPaymentOptions201Response {
	return v.value
}

func (v *NullablePOSTPaymentOptions201Response) Set(val *POSTPaymentOptions201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPaymentOptions201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPaymentOptions201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPaymentOptions201Response(val *POSTPaymentOptions201Response) *NullablePOSTPaymentOptions201Response {
	return &NullablePOSTPaymentOptions201Response{value: val, isSet: true}
}

func (v NullablePOSTPaymentOptions201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPaymentOptions201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
