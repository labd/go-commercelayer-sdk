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

// checks if the PATCHBraintreeGatewaysBraintreeGatewayId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHBraintreeGatewaysBraintreeGatewayId200Response{}

// PATCHBraintreeGatewaysBraintreeGatewayId200Response struct for PATCHBraintreeGatewaysBraintreeGatewayId200Response
type PATCHBraintreeGatewaysBraintreeGatewayId200Response struct {
	Data *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData `json:"data,omitempty"`
}

// NewPATCHBraintreeGatewaysBraintreeGatewayId200Response instantiates a new PATCHBraintreeGatewaysBraintreeGatewayId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHBraintreeGatewaysBraintreeGatewayId200Response() *PATCHBraintreeGatewaysBraintreeGatewayId200Response {
	this := PATCHBraintreeGatewaysBraintreeGatewayId200Response{}
	return &this
}

// NewPATCHBraintreeGatewaysBraintreeGatewayId200ResponseWithDefaults instantiates a new PATCHBraintreeGatewaysBraintreeGatewayId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHBraintreeGatewaysBraintreeGatewayId200ResponseWithDefaults() *PATCHBraintreeGatewaysBraintreeGatewayId200Response {
	this := PATCHBraintreeGatewaysBraintreeGatewayId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHBraintreeGatewaysBraintreeGatewayId200Response) GetData() PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHBraintreeGatewaysBraintreeGatewayId200Response) GetDataOk() (*PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHBraintreeGatewaysBraintreeGatewayId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData and assigns it to the Data field.
func (o *PATCHBraintreeGatewaysBraintreeGatewayId200Response) SetData(v PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) {
	o.Data = &v
}

func (o PATCHBraintreeGatewaysBraintreeGatewayId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHBraintreeGatewaysBraintreeGatewayId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHBraintreeGatewaysBraintreeGatewayId200Response struct {
	value *PATCHBraintreeGatewaysBraintreeGatewayId200Response
	isSet bool
}

func (v NullablePATCHBraintreeGatewaysBraintreeGatewayId200Response) Get() *PATCHBraintreeGatewaysBraintreeGatewayId200Response {
	return v.value
}

func (v *NullablePATCHBraintreeGatewaysBraintreeGatewayId200Response) Set(val *PATCHBraintreeGatewaysBraintreeGatewayId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHBraintreeGatewaysBraintreeGatewayId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHBraintreeGatewaysBraintreeGatewayId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHBraintreeGatewaysBraintreeGatewayId200Response(val *PATCHBraintreeGatewaysBraintreeGatewayId200Response) *NullablePATCHBraintreeGatewaysBraintreeGatewayId200Response {
	return &NullablePATCHBraintreeGatewaysBraintreeGatewayId200Response{value: val, isSet: true}
}

func (v NullablePATCHBraintreeGatewaysBraintreeGatewayId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHBraintreeGatewaysBraintreeGatewayId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
