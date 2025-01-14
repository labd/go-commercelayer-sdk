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

// checks if the GETPaypalGatewaysPaypalGatewayId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETPaypalGatewaysPaypalGatewayId200Response{}

// GETPaypalGatewaysPaypalGatewayId200Response struct for GETPaypalGatewaysPaypalGatewayId200Response
type GETPaypalGatewaysPaypalGatewayId200Response struct {
	Data *GETPaypalGatewaysPaypalGatewayId200ResponseData `json:"data,omitempty"`
}

// NewGETPaypalGatewaysPaypalGatewayId200Response instantiates a new GETPaypalGatewaysPaypalGatewayId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPaypalGatewaysPaypalGatewayId200Response() *GETPaypalGatewaysPaypalGatewayId200Response {
	this := GETPaypalGatewaysPaypalGatewayId200Response{}
	return &this
}

// NewGETPaypalGatewaysPaypalGatewayId200ResponseWithDefaults instantiates a new GETPaypalGatewaysPaypalGatewayId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPaypalGatewaysPaypalGatewayId200ResponseWithDefaults() *GETPaypalGatewaysPaypalGatewayId200Response {
	this := GETPaypalGatewaysPaypalGatewayId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETPaypalGatewaysPaypalGatewayId200Response) GetData() GETPaypalGatewaysPaypalGatewayId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETPaypalGatewaysPaypalGatewayId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPaypalGatewaysPaypalGatewayId200Response) GetDataOk() (*GETPaypalGatewaysPaypalGatewayId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETPaypalGatewaysPaypalGatewayId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETPaypalGatewaysPaypalGatewayId200ResponseData and assigns it to the Data field.
func (o *GETPaypalGatewaysPaypalGatewayId200Response) SetData(v GETPaypalGatewaysPaypalGatewayId200ResponseData) {
	o.Data = &v
}

func (o GETPaypalGatewaysPaypalGatewayId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETPaypalGatewaysPaypalGatewayId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETPaypalGatewaysPaypalGatewayId200Response struct {
	value *GETPaypalGatewaysPaypalGatewayId200Response
	isSet bool
}

func (v NullableGETPaypalGatewaysPaypalGatewayId200Response) Get() *GETPaypalGatewaysPaypalGatewayId200Response {
	return v.value
}

func (v *NullableGETPaypalGatewaysPaypalGatewayId200Response) Set(val *GETPaypalGatewaysPaypalGatewayId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPaypalGatewaysPaypalGatewayId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPaypalGatewaysPaypalGatewayId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPaypalGatewaysPaypalGatewayId200Response(val *GETPaypalGatewaysPaypalGatewayId200Response) *NullableGETPaypalGatewaysPaypalGatewayId200Response {
	return &NullableGETPaypalGatewaysPaypalGatewayId200Response{value: val, isSet: true}
}

func (v NullableGETPaypalGatewaysPaypalGatewayId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPaypalGatewaysPaypalGatewayId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
