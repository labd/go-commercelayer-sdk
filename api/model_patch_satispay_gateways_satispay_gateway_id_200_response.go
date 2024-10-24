/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PATCHSatispayGatewaysSatispayGatewayId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHSatispayGatewaysSatispayGatewayId200Response{}

// PATCHSatispayGatewaysSatispayGatewayId200Response struct for PATCHSatispayGatewaysSatispayGatewayId200Response
type PATCHSatispayGatewaysSatispayGatewayId200Response struct {
	Data *PATCHSatispayGatewaysSatispayGatewayId200ResponseData `json:"data,omitempty"`
}

// NewPATCHSatispayGatewaysSatispayGatewayId200Response instantiates a new PATCHSatispayGatewaysSatispayGatewayId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHSatispayGatewaysSatispayGatewayId200Response() *PATCHSatispayGatewaysSatispayGatewayId200Response {
	this := PATCHSatispayGatewaysSatispayGatewayId200Response{}
	return &this
}

// NewPATCHSatispayGatewaysSatispayGatewayId200ResponseWithDefaults instantiates a new PATCHSatispayGatewaysSatispayGatewayId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHSatispayGatewaysSatispayGatewayId200ResponseWithDefaults() *PATCHSatispayGatewaysSatispayGatewayId200Response {
	this := PATCHSatispayGatewaysSatispayGatewayId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHSatispayGatewaysSatispayGatewayId200Response) GetData() PATCHSatispayGatewaysSatispayGatewayId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHSatispayGatewaysSatispayGatewayId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHSatispayGatewaysSatispayGatewayId200Response) GetDataOk() (*PATCHSatispayGatewaysSatispayGatewayId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHSatispayGatewaysSatispayGatewayId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHSatispayGatewaysSatispayGatewayId200ResponseData and assigns it to the Data field.
func (o *PATCHSatispayGatewaysSatispayGatewayId200Response) SetData(v PATCHSatispayGatewaysSatispayGatewayId200ResponseData) {
	o.Data = &v
}

func (o PATCHSatispayGatewaysSatispayGatewayId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHSatispayGatewaysSatispayGatewayId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHSatispayGatewaysSatispayGatewayId200Response struct {
	value *PATCHSatispayGatewaysSatispayGatewayId200Response
	isSet bool
}

func (v NullablePATCHSatispayGatewaysSatispayGatewayId200Response) Get() *PATCHSatispayGatewaysSatispayGatewayId200Response {
	return v.value
}

func (v *NullablePATCHSatispayGatewaysSatispayGatewayId200Response) Set(val *PATCHSatispayGatewaysSatispayGatewayId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHSatispayGatewaysSatispayGatewayId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHSatispayGatewaysSatispayGatewayId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHSatispayGatewaysSatispayGatewayId200Response(val *PATCHSatispayGatewaysSatispayGatewayId200Response) *NullablePATCHSatispayGatewaysSatispayGatewayId200Response {
	return &NullablePATCHSatispayGatewaysSatispayGatewayId200Response{value: val, isSet: true}
}

func (v NullablePATCHSatispayGatewaysSatispayGatewayId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHSatispayGatewaysSatispayGatewayId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}