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

// GETBraintreeGatewaysBraintreeGatewayId200Response struct for GETBraintreeGatewaysBraintreeGatewayId200Response
type GETBraintreeGatewaysBraintreeGatewayId200Response struct {
	Data *GETBraintreeGateways200ResponseDataInner `json:"data,omitempty"`
}

// NewGETBraintreeGatewaysBraintreeGatewayId200Response instantiates a new GETBraintreeGatewaysBraintreeGatewayId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETBraintreeGatewaysBraintreeGatewayId200Response() *GETBraintreeGatewaysBraintreeGatewayId200Response {
	this := GETBraintreeGatewaysBraintreeGatewayId200Response{}
	return &this
}

// NewGETBraintreeGatewaysBraintreeGatewayId200ResponseWithDefaults instantiates a new GETBraintreeGatewaysBraintreeGatewayId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETBraintreeGatewaysBraintreeGatewayId200ResponseWithDefaults() *GETBraintreeGatewaysBraintreeGatewayId200Response {
	this := GETBraintreeGatewaysBraintreeGatewayId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETBraintreeGatewaysBraintreeGatewayId200Response) GetData() GETBraintreeGateways200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETBraintreeGateways200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBraintreeGatewaysBraintreeGatewayId200Response) GetDataOk() (*GETBraintreeGateways200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETBraintreeGatewaysBraintreeGatewayId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETBraintreeGateways200ResponseDataInner and assigns it to the Data field.
func (o *GETBraintreeGatewaysBraintreeGatewayId200Response) SetData(v GETBraintreeGateways200ResponseDataInner) {
	o.Data = &v
}

func (o GETBraintreeGatewaysBraintreeGatewayId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETBraintreeGatewaysBraintreeGatewayId200Response struct {
	value *GETBraintreeGatewaysBraintreeGatewayId200Response
	isSet bool
}

func (v NullableGETBraintreeGatewaysBraintreeGatewayId200Response) Get() *GETBraintreeGatewaysBraintreeGatewayId200Response {
	return v.value
}

func (v *NullableGETBraintreeGatewaysBraintreeGatewayId200Response) Set(val *GETBraintreeGatewaysBraintreeGatewayId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETBraintreeGatewaysBraintreeGatewayId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETBraintreeGatewaysBraintreeGatewayId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETBraintreeGatewaysBraintreeGatewayId200Response(val *GETBraintreeGatewaysBraintreeGatewayId200Response) *NullableGETBraintreeGatewaysBraintreeGatewayId200Response {
	return &NullableGETBraintreeGatewaysBraintreeGatewayId200Response{value: val, isSet: true}
}

func (v NullableGETBraintreeGatewaysBraintreeGatewayId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETBraintreeGatewaysBraintreeGatewayId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
