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

// GETReturnsReturnId200Response struct for GETReturnsReturnId200Response
type GETReturnsReturnId200Response struct {
	Data *GETReturns200ResponseDataInner `json:"data,omitempty"`
}

// NewGETReturnsReturnId200Response instantiates a new GETReturnsReturnId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETReturnsReturnId200Response() *GETReturnsReturnId200Response {
	this := GETReturnsReturnId200Response{}
	return &this
}

// NewGETReturnsReturnId200ResponseWithDefaults instantiates a new GETReturnsReturnId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETReturnsReturnId200ResponseWithDefaults() *GETReturnsReturnId200Response {
	this := GETReturnsReturnId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETReturnsReturnId200Response) GetData() GETReturns200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETReturns200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturnsReturnId200Response) GetDataOk() (*GETReturns200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETReturnsReturnId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETReturns200ResponseDataInner and assigns it to the Data field.
func (o *GETReturnsReturnId200Response) SetData(v GETReturns200ResponseDataInner) {
	o.Data = &v
}

func (o GETReturnsReturnId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETReturnsReturnId200Response struct {
	value *GETReturnsReturnId200Response
	isSet bool
}

func (v NullableGETReturnsReturnId200Response) Get() *GETReturnsReturnId200Response {
	return v.value
}

func (v *NullableGETReturnsReturnId200Response) Set(val *GETReturnsReturnId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETReturnsReturnId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETReturnsReturnId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETReturnsReturnId200Response(val *GETReturnsReturnId200Response) *NullableGETReturnsReturnId200Response {
	return &NullableGETReturnsReturnId200Response{value: val, isSet: true}
}

func (v NullableGETReturnsReturnId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETReturnsReturnId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
