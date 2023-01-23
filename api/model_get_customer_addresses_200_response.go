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

// GETCustomerAddresses200Response struct for GETCustomerAddresses200Response
type GETCustomerAddresses200Response struct {
	Data []GETCustomerAddresses200ResponseDataInner `json:"data,omitempty"`
}

// NewGETCustomerAddresses200Response instantiates a new GETCustomerAddresses200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCustomerAddresses200Response() *GETCustomerAddresses200Response {
	this := GETCustomerAddresses200Response{}
	return &this
}

// NewGETCustomerAddresses200ResponseWithDefaults instantiates a new GETCustomerAddresses200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCustomerAddresses200ResponseWithDefaults() *GETCustomerAddresses200Response {
	this := GETCustomerAddresses200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCustomerAddresses200Response) GetData() []GETCustomerAddresses200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETCustomerAddresses200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomerAddresses200Response) GetDataOk() ([]GETCustomerAddresses200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCustomerAddresses200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETCustomerAddresses200ResponseDataInner and assigns it to the Data field.
func (o *GETCustomerAddresses200Response) SetData(v []GETCustomerAddresses200ResponseDataInner) {
	o.Data = v
}

func (o GETCustomerAddresses200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETCustomerAddresses200Response struct {
	value *GETCustomerAddresses200Response
	isSet bool
}

func (v NullableGETCustomerAddresses200Response) Get() *GETCustomerAddresses200Response {
	return v.value
}

func (v *NullableGETCustomerAddresses200Response) Set(val *GETCustomerAddresses200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCustomerAddresses200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCustomerAddresses200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCustomerAddresses200Response(val *GETCustomerAddresses200Response) *NullableGETCustomerAddresses200Response {
	return &NullableGETCustomerAddresses200Response{value: val, isSet: true}
}

func (v NullableGETCustomerAddresses200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCustomerAddresses200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
