/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PriceResponse struct for PriceResponse
type PriceResponse struct {
	Data *PriceResponseData `json:"data,omitempty"`
}

// NewPriceResponse instantiates a new PriceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceResponse() *PriceResponse {
	this := PriceResponse{}
	return &this
}

// NewPriceResponseWithDefaults instantiates a new PriceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceResponseWithDefaults() *PriceResponse {
	this := PriceResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PriceResponse) GetData() PriceResponseData {
	if o == nil || o.Data == nil {
		var ret PriceResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceResponse) GetDataOk() (*PriceResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PriceResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PriceResponseData and assigns it to the Data field.
func (o *PriceResponse) SetData(v PriceResponseData) {
	o.Data = &v
}

func (o PriceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePriceResponse struct {
	value *PriceResponse
	isSet bool
}

func (v NullablePriceResponse) Get() *PriceResponse {
	return v.value
}

func (v *NullablePriceResponse) Set(val *PriceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceResponse(val *PriceResponse) *NullablePriceResponse {
	return &NullablePriceResponse{value: val, isSet: true}
}

func (v NullablePriceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
