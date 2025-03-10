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

// checks if the GETSatispayGatewaysSatispayGatewayId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETSatispayGatewaysSatispayGatewayId200Response{}

// GETSatispayGatewaysSatispayGatewayId200Response struct for GETSatispayGatewaysSatispayGatewayId200Response
type GETSatispayGatewaysSatispayGatewayId200Response struct {
	Data *GETSatispayGatewaysSatispayGatewayId200ResponseData `json:"data,omitempty"`
}

// NewGETSatispayGatewaysSatispayGatewayId200Response instantiates a new GETSatispayGatewaysSatispayGatewayId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETSatispayGatewaysSatispayGatewayId200Response() *GETSatispayGatewaysSatispayGatewayId200Response {
	this := GETSatispayGatewaysSatispayGatewayId200Response{}
	return &this
}

// NewGETSatispayGatewaysSatispayGatewayId200ResponseWithDefaults instantiates a new GETSatispayGatewaysSatispayGatewayId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETSatispayGatewaysSatispayGatewayId200ResponseWithDefaults() *GETSatispayGatewaysSatispayGatewayId200Response {
	this := GETSatispayGatewaysSatispayGatewayId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETSatispayGatewaysSatispayGatewayId200Response) GetData() GETSatispayGatewaysSatispayGatewayId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETSatispayGatewaysSatispayGatewayId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSatispayGatewaysSatispayGatewayId200Response) GetDataOk() (*GETSatispayGatewaysSatispayGatewayId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETSatispayGatewaysSatispayGatewayId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETSatispayGatewaysSatispayGatewayId200ResponseData and assigns it to the Data field.
func (o *GETSatispayGatewaysSatispayGatewayId200Response) SetData(v GETSatispayGatewaysSatispayGatewayId200ResponseData) {
	o.Data = &v
}

func (o GETSatispayGatewaysSatispayGatewayId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETSatispayGatewaysSatispayGatewayId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETSatispayGatewaysSatispayGatewayId200Response struct {
	value *GETSatispayGatewaysSatispayGatewayId200Response
	isSet bool
}

func (v NullableGETSatispayGatewaysSatispayGatewayId200Response) Get() *GETSatispayGatewaysSatispayGatewayId200Response {
	return v.value
}

func (v *NullableGETSatispayGatewaysSatispayGatewayId200Response) Set(val *GETSatispayGatewaysSatispayGatewayId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETSatispayGatewaysSatispayGatewayId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETSatispayGatewaysSatispayGatewayId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETSatispayGatewaysSatispayGatewayId200Response(val *GETSatispayGatewaysSatispayGatewayId200Response) *NullableGETSatispayGatewaysSatispayGatewayId200Response {
	return &NullableGETSatispayGatewaysSatispayGatewayId200Response{value: val, isSet: true}
}

func (v NullableGETSatispayGatewaysSatispayGatewayId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETSatispayGatewaysSatispayGatewayId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
