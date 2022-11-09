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

// GETExternalPayments200Response struct for GETExternalPayments200Response
type GETExternalPayments200Response struct {
	Data []GETExternalPayments200ResponseDataInner `json:"data,omitempty"`
}

// NewGETExternalPayments200Response instantiates a new GETExternalPayments200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETExternalPayments200Response() *GETExternalPayments200Response {
	this := GETExternalPayments200Response{}
	return &this
}

// NewGETExternalPayments200ResponseWithDefaults instantiates a new GETExternalPayments200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETExternalPayments200ResponseWithDefaults() *GETExternalPayments200Response {
	this := GETExternalPayments200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETExternalPayments200Response) GetData() []GETExternalPayments200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETExternalPayments200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalPayments200Response) GetDataOk() ([]GETExternalPayments200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETExternalPayments200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETExternalPayments200ResponseDataInner and assigns it to the Data field.
func (o *GETExternalPayments200Response) SetData(v []GETExternalPayments200ResponseDataInner) {
	o.Data = v
}

func (o GETExternalPayments200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETExternalPayments200Response struct {
	value *GETExternalPayments200Response
	isSet bool
}

func (v NullableGETExternalPayments200Response) Get() *GETExternalPayments200Response {
	return v.value
}

func (v *NullableGETExternalPayments200Response) Set(val *GETExternalPayments200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETExternalPayments200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETExternalPayments200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETExternalPayments200Response(val *GETExternalPayments200Response) *NullableGETExternalPayments200Response {
	return &NullableGETExternalPayments200Response{value: val, isSet: true}
}

func (v NullableGETExternalPayments200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETExternalPayments200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
