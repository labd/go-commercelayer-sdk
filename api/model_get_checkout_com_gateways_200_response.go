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

// GETCheckoutComGateways200Response struct for GETCheckoutComGateways200Response
type GETCheckoutComGateways200Response struct {
	Data []GETCheckoutComGateways200ResponseDataInner `json:"data,omitempty"`
}

// NewGETCheckoutComGateways200Response instantiates a new GETCheckoutComGateways200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCheckoutComGateways200Response() *GETCheckoutComGateways200Response {
	this := GETCheckoutComGateways200Response{}
	return &this
}

// NewGETCheckoutComGateways200ResponseWithDefaults instantiates a new GETCheckoutComGateways200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCheckoutComGateways200ResponseWithDefaults() *GETCheckoutComGateways200Response {
	this := GETCheckoutComGateways200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCheckoutComGateways200Response) GetData() []GETCheckoutComGateways200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETCheckoutComGateways200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComGateways200Response) GetDataOk() ([]GETCheckoutComGateways200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCheckoutComGateways200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETCheckoutComGateways200ResponseDataInner and assigns it to the Data field.
func (o *GETCheckoutComGateways200Response) SetData(v []GETCheckoutComGateways200ResponseDataInner) {
	o.Data = v
}

func (o GETCheckoutComGateways200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETCheckoutComGateways200Response struct {
	value *GETCheckoutComGateways200Response
	isSet bool
}

func (v NullableGETCheckoutComGateways200Response) Get() *GETCheckoutComGateways200Response {
	return v.value
}

func (v *NullableGETCheckoutComGateways200Response) Set(val *GETCheckoutComGateways200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCheckoutComGateways200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCheckoutComGateways200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCheckoutComGateways200Response(val *GETCheckoutComGateways200Response) *NullableGETCheckoutComGateways200Response {
	return &NullableGETCheckoutComGateways200Response{value: val, isSet: true}
}

func (v NullableGETCheckoutComGateways200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCheckoutComGateways200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
