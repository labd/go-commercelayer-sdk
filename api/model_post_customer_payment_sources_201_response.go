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

// checks if the POSTCustomerPaymentSources201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCustomerPaymentSources201Response{}

// POSTCustomerPaymentSources201Response struct for POSTCustomerPaymentSources201Response
type POSTCustomerPaymentSources201Response struct {
	Data *POSTCustomerPaymentSources201ResponseData `json:"data,omitempty"`
}

// NewPOSTCustomerPaymentSources201Response instantiates a new POSTCustomerPaymentSources201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCustomerPaymentSources201Response() *POSTCustomerPaymentSources201Response {
	this := POSTCustomerPaymentSources201Response{}
	return &this
}

// NewPOSTCustomerPaymentSources201ResponseWithDefaults instantiates a new POSTCustomerPaymentSources201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCustomerPaymentSources201ResponseWithDefaults() *POSTCustomerPaymentSources201Response {
	this := POSTCustomerPaymentSources201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTCustomerPaymentSources201Response) GetData() POSTCustomerPaymentSources201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTCustomerPaymentSources201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCustomerPaymentSources201Response) GetDataOk() (*POSTCustomerPaymentSources201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTCustomerPaymentSources201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTCustomerPaymentSources201ResponseData and assigns it to the Data field.
func (o *POSTCustomerPaymentSources201Response) SetData(v POSTCustomerPaymentSources201ResponseData) {
	o.Data = &v
}

func (o POSTCustomerPaymentSources201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCustomerPaymentSources201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTCustomerPaymentSources201Response struct {
	value *POSTCustomerPaymentSources201Response
	isSet bool
}

func (v NullablePOSTCustomerPaymentSources201Response) Get() *POSTCustomerPaymentSources201Response {
	return v.value
}

func (v *NullablePOSTCustomerPaymentSources201Response) Set(val *POSTCustomerPaymentSources201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCustomerPaymentSources201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCustomerPaymentSources201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCustomerPaymentSources201Response(val *POSTCustomerPaymentSources201Response) *NullablePOSTCustomerPaymentSources201Response {
	return &NullablePOSTCustomerPaymentSources201Response{value: val, isSet: true}
}

func (v NullablePOSTCustomerPaymentSources201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCustomerPaymentSources201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
