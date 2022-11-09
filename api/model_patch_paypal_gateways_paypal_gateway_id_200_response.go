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

// PATCHPaypalGatewaysPaypalGatewayId200Response struct for PATCHPaypalGatewaysPaypalGatewayId200Response
type PATCHPaypalGatewaysPaypalGatewayId200Response struct {
	Data *PATCHPaypalGatewaysPaypalGatewayId200ResponseData `json:"data,omitempty"`
}

// NewPATCHPaypalGatewaysPaypalGatewayId200Response instantiates a new PATCHPaypalGatewaysPaypalGatewayId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHPaypalGatewaysPaypalGatewayId200Response() *PATCHPaypalGatewaysPaypalGatewayId200Response {
	this := PATCHPaypalGatewaysPaypalGatewayId200Response{}
	return &this
}

// NewPATCHPaypalGatewaysPaypalGatewayId200ResponseWithDefaults instantiates a new PATCHPaypalGatewaysPaypalGatewayId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHPaypalGatewaysPaypalGatewayId200ResponseWithDefaults() *PATCHPaypalGatewaysPaypalGatewayId200Response {
	this := PATCHPaypalGatewaysPaypalGatewayId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHPaypalGatewaysPaypalGatewayId200Response) GetData() PATCHPaypalGatewaysPaypalGatewayId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHPaypalGatewaysPaypalGatewayId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPaypalGatewaysPaypalGatewayId200Response) GetDataOk() (*PATCHPaypalGatewaysPaypalGatewayId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHPaypalGatewaysPaypalGatewayId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHPaypalGatewaysPaypalGatewayId200ResponseData and assigns it to the Data field.
func (o *PATCHPaypalGatewaysPaypalGatewayId200Response) SetData(v PATCHPaypalGatewaysPaypalGatewayId200ResponseData) {
	o.Data = &v
}

func (o PATCHPaypalGatewaysPaypalGatewayId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHPaypalGatewaysPaypalGatewayId200Response struct {
	value *PATCHPaypalGatewaysPaypalGatewayId200Response
	isSet bool
}

func (v NullablePATCHPaypalGatewaysPaypalGatewayId200Response) Get() *PATCHPaypalGatewaysPaypalGatewayId200Response {
	return v.value
}

func (v *NullablePATCHPaypalGatewaysPaypalGatewayId200Response) Set(val *PATCHPaypalGatewaysPaypalGatewayId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHPaypalGatewaysPaypalGatewayId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHPaypalGatewaysPaypalGatewayId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHPaypalGatewaysPaypalGatewayId200Response(val *PATCHPaypalGatewaysPaypalGatewayId200Response) *NullablePATCHPaypalGatewaysPaypalGatewayId200Response {
	return &NullablePATCHPaypalGatewaysPaypalGatewayId200Response{value: val, isSet: true}
}

func (v NullablePATCHPaypalGatewaysPaypalGatewayId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHPaypalGatewaysPaypalGatewayId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}