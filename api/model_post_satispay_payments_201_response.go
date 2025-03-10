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

// checks if the POSTSatispayPayments201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTSatispayPayments201Response{}

// POSTSatispayPayments201Response struct for POSTSatispayPayments201Response
type POSTSatispayPayments201Response struct {
	Data *POSTSatispayPayments201ResponseData `json:"data,omitempty"`
}

// NewPOSTSatispayPayments201Response instantiates a new POSTSatispayPayments201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTSatispayPayments201Response() *POSTSatispayPayments201Response {
	this := POSTSatispayPayments201Response{}
	return &this
}

// NewPOSTSatispayPayments201ResponseWithDefaults instantiates a new POSTSatispayPayments201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTSatispayPayments201ResponseWithDefaults() *POSTSatispayPayments201Response {
	this := POSTSatispayPayments201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTSatispayPayments201Response) GetData() POSTSatispayPayments201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTSatispayPayments201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTSatispayPayments201Response) GetDataOk() (*POSTSatispayPayments201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTSatispayPayments201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTSatispayPayments201ResponseData and assigns it to the Data field.
func (o *POSTSatispayPayments201Response) SetData(v POSTSatispayPayments201ResponseData) {
	o.Data = &v
}

func (o POSTSatispayPayments201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTSatispayPayments201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTSatispayPayments201Response struct {
	value *POSTSatispayPayments201Response
	isSet bool
}

func (v NullablePOSTSatispayPayments201Response) Get() *POSTSatispayPayments201Response {
	return v.value
}

func (v *NullablePOSTSatispayPayments201Response) Set(val *POSTSatispayPayments201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTSatispayPayments201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTSatispayPayments201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTSatispayPayments201Response(val *POSTSatispayPayments201Response) *NullablePOSTSatispayPayments201Response {
	return &NullablePOSTSatispayPayments201Response{value: val, isSet: true}
}

func (v NullablePOSTSatispayPayments201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTSatispayPayments201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
