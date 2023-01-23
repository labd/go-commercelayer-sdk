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

// GETCheckoutComGatewaysCheckoutComGatewayId200Response struct for GETCheckoutComGatewaysCheckoutComGatewayId200Response
type GETCheckoutComGatewaysCheckoutComGatewayId200Response struct {
	Data *GETCheckoutComGateways200ResponseDataInner `json:"data,omitempty"`
}

// NewGETCheckoutComGatewaysCheckoutComGatewayId200Response instantiates a new GETCheckoutComGatewaysCheckoutComGatewayId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCheckoutComGatewaysCheckoutComGatewayId200Response() *GETCheckoutComGatewaysCheckoutComGatewayId200Response {
	this := GETCheckoutComGatewaysCheckoutComGatewayId200Response{}
	return &this
}

// NewGETCheckoutComGatewaysCheckoutComGatewayId200ResponseWithDefaults instantiates a new GETCheckoutComGatewaysCheckoutComGatewayId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCheckoutComGatewaysCheckoutComGatewayId200ResponseWithDefaults() *GETCheckoutComGatewaysCheckoutComGatewayId200Response {
	this := GETCheckoutComGatewaysCheckoutComGatewayId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200Response) GetData() GETCheckoutComGateways200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETCheckoutComGateways200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200Response) GetDataOk() (*GETCheckoutComGateways200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETCheckoutComGateways200ResponseDataInner and assigns it to the Data field.
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200Response) SetData(v GETCheckoutComGateways200ResponseDataInner) {
	o.Data = &v
}

func (o GETCheckoutComGatewaysCheckoutComGatewayId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETCheckoutComGatewaysCheckoutComGatewayId200Response struct {
	value *GETCheckoutComGatewaysCheckoutComGatewayId200Response
	isSet bool
}

func (v NullableGETCheckoutComGatewaysCheckoutComGatewayId200Response) Get() *GETCheckoutComGatewaysCheckoutComGatewayId200Response {
	return v.value
}

func (v *NullableGETCheckoutComGatewaysCheckoutComGatewayId200Response) Set(val *GETCheckoutComGatewaysCheckoutComGatewayId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCheckoutComGatewaysCheckoutComGatewayId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCheckoutComGatewaysCheckoutComGatewayId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCheckoutComGatewaysCheckoutComGatewayId200Response(val *GETCheckoutComGatewaysCheckoutComGatewayId200Response) *NullableGETCheckoutComGatewaysCheckoutComGatewayId200Response {
	return &NullableGETCheckoutComGatewaysCheckoutComGatewayId200Response{value: val, isSet: true}
}

func (v NullableGETCheckoutComGatewaysCheckoutComGatewayId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCheckoutComGatewaysCheckoutComGatewayId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
