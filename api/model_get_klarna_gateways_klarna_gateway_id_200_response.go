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

// GETKlarnaGatewaysKlarnaGatewayId200Response struct for GETKlarnaGatewaysKlarnaGatewayId200Response
type GETKlarnaGatewaysKlarnaGatewayId200Response struct {
	Data *GETKlarnaGateways200ResponseDataInner `json:"data,omitempty"`
}

// NewGETKlarnaGatewaysKlarnaGatewayId200Response instantiates a new GETKlarnaGatewaysKlarnaGatewayId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETKlarnaGatewaysKlarnaGatewayId200Response() *GETKlarnaGatewaysKlarnaGatewayId200Response {
	this := GETKlarnaGatewaysKlarnaGatewayId200Response{}
	return &this
}

// NewGETKlarnaGatewaysKlarnaGatewayId200ResponseWithDefaults instantiates a new GETKlarnaGatewaysKlarnaGatewayId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETKlarnaGatewaysKlarnaGatewayId200ResponseWithDefaults() *GETKlarnaGatewaysKlarnaGatewayId200Response {
	this := GETKlarnaGatewaysKlarnaGatewayId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETKlarnaGatewaysKlarnaGatewayId200Response) GetData() GETKlarnaGateways200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETKlarnaGateways200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETKlarnaGatewaysKlarnaGatewayId200Response) GetDataOk() (*GETKlarnaGateways200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETKlarnaGatewaysKlarnaGatewayId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETKlarnaGateways200ResponseDataInner and assigns it to the Data field.
func (o *GETKlarnaGatewaysKlarnaGatewayId200Response) SetData(v GETKlarnaGateways200ResponseDataInner) {
	o.Data = &v
}

func (o GETKlarnaGatewaysKlarnaGatewayId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETKlarnaGatewaysKlarnaGatewayId200Response struct {
	value *GETKlarnaGatewaysKlarnaGatewayId200Response
	isSet bool
}

func (v NullableGETKlarnaGatewaysKlarnaGatewayId200Response) Get() *GETKlarnaGatewaysKlarnaGatewayId200Response {
	return v.value
}

func (v *NullableGETKlarnaGatewaysKlarnaGatewayId200Response) Set(val *GETKlarnaGatewaysKlarnaGatewayId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETKlarnaGatewaysKlarnaGatewayId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETKlarnaGatewaysKlarnaGatewayId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETKlarnaGatewaysKlarnaGatewayId200Response(val *GETKlarnaGatewaysKlarnaGatewayId200Response) *NullableGETKlarnaGatewaysKlarnaGatewayId200Response {
	return &NullableGETKlarnaGatewaysKlarnaGatewayId200Response{value: val, isSet: true}
}

func (v NullableGETKlarnaGatewaysKlarnaGatewayId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETKlarnaGatewaysKlarnaGatewayId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}