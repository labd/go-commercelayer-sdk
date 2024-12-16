/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.5.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GETExternalGatewaysExternalGatewayId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETExternalGatewaysExternalGatewayId200Response{}

// GETExternalGatewaysExternalGatewayId200Response struct for GETExternalGatewaysExternalGatewayId200Response
type GETExternalGatewaysExternalGatewayId200Response struct {
	Data *GETExternalGatewaysExternalGatewayId200ResponseData `json:"data,omitempty"`
}

// NewGETExternalGatewaysExternalGatewayId200Response instantiates a new GETExternalGatewaysExternalGatewayId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETExternalGatewaysExternalGatewayId200Response() *GETExternalGatewaysExternalGatewayId200Response {
	this := GETExternalGatewaysExternalGatewayId200Response{}
	return &this
}

// NewGETExternalGatewaysExternalGatewayId200ResponseWithDefaults instantiates a new GETExternalGatewaysExternalGatewayId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETExternalGatewaysExternalGatewayId200ResponseWithDefaults() *GETExternalGatewaysExternalGatewayId200Response {
	this := GETExternalGatewaysExternalGatewayId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETExternalGatewaysExternalGatewayId200Response) GetData() GETExternalGatewaysExternalGatewayId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETExternalGatewaysExternalGatewayId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalGatewaysExternalGatewayId200Response) GetDataOk() (*GETExternalGatewaysExternalGatewayId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETExternalGatewaysExternalGatewayId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETExternalGatewaysExternalGatewayId200ResponseData and assigns it to the Data field.
func (o *GETExternalGatewaysExternalGatewayId200Response) SetData(v GETExternalGatewaysExternalGatewayId200ResponseData) {
	o.Data = &v
}

func (o GETExternalGatewaysExternalGatewayId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETExternalGatewaysExternalGatewayId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETExternalGatewaysExternalGatewayId200Response struct {
	value *GETExternalGatewaysExternalGatewayId200Response
	isSet bool
}

func (v NullableGETExternalGatewaysExternalGatewayId200Response) Get() *GETExternalGatewaysExternalGatewayId200Response {
	return v.value
}

func (v *NullableGETExternalGatewaysExternalGatewayId200Response) Set(val *GETExternalGatewaysExternalGatewayId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETExternalGatewaysExternalGatewayId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETExternalGatewaysExternalGatewayId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETExternalGatewaysExternalGatewayId200Response(val *GETExternalGatewaysExternalGatewayId200Response) *NullableGETExternalGatewaysExternalGatewayId200Response {
	return &NullableGETExternalGatewaysExternalGatewayId200Response{value: val, isSet: true}
}

func (v NullableGETExternalGatewaysExternalGatewayId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETExternalGatewaysExternalGatewayId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
