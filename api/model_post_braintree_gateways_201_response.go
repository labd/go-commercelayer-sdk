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

// checks if the POSTBraintreeGateways201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTBraintreeGateways201Response{}

// POSTBraintreeGateways201Response struct for POSTBraintreeGateways201Response
type POSTBraintreeGateways201Response struct {
	Data *POSTBraintreeGateways201ResponseData `json:"data,omitempty"`
}

// NewPOSTBraintreeGateways201Response instantiates a new POSTBraintreeGateways201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTBraintreeGateways201Response() *POSTBraintreeGateways201Response {
	this := POSTBraintreeGateways201Response{}
	return &this
}

// NewPOSTBraintreeGateways201ResponseWithDefaults instantiates a new POSTBraintreeGateways201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTBraintreeGateways201ResponseWithDefaults() *POSTBraintreeGateways201Response {
	this := POSTBraintreeGateways201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTBraintreeGateways201Response) GetData() POSTBraintreeGateways201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTBraintreeGateways201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBraintreeGateways201Response) GetDataOk() (*POSTBraintreeGateways201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTBraintreeGateways201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTBraintreeGateways201ResponseData and assigns it to the Data field.
func (o *POSTBraintreeGateways201Response) SetData(v POSTBraintreeGateways201ResponseData) {
	o.Data = &v
}

func (o POSTBraintreeGateways201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTBraintreeGateways201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTBraintreeGateways201Response struct {
	value *POSTBraintreeGateways201Response
	isSet bool
}

func (v NullablePOSTBraintreeGateways201Response) Get() *POSTBraintreeGateways201Response {
	return v.value
}

func (v *NullablePOSTBraintreeGateways201Response) Set(val *POSTBraintreeGateways201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTBraintreeGateways201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTBraintreeGateways201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTBraintreeGateways201Response(val *POSTBraintreeGateways201Response) *NullablePOSTBraintreeGateways201Response {
	return &NullablePOSTBraintreeGateways201Response{value: val, isSet: true}
}

func (v NullablePOSTBraintreeGateways201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTBraintreeGateways201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
