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

// PATCHSkusSkuId200Response struct for PATCHSkusSkuId200Response
type PATCHSkusSkuId200Response struct {
	Data *PATCHSkusSkuId200ResponseData `json:"data,omitempty"`
}

// NewPATCHSkusSkuId200Response instantiates a new PATCHSkusSkuId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHSkusSkuId200Response() *PATCHSkusSkuId200Response {
	this := PATCHSkusSkuId200Response{}
	return &this
}

// NewPATCHSkusSkuId200ResponseWithDefaults instantiates a new PATCHSkusSkuId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHSkusSkuId200ResponseWithDefaults() *PATCHSkusSkuId200Response {
	this := PATCHSkusSkuId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHSkusSkuId200Response) GetData() PATCHSkusSkuId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHSkusSkuId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHSkusSkuId200Response) GetDataOk() (*PATCHSkusSkuId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHSkusSkuId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHSkusSkuId200ResponseData and assigns it to the Data field.
func (o *PATCHSkusSkuId200Response) SetData(v PATCHSkusSkuId200ResponseData) {
	o.Data = &v
}

func (o PATCHSkusSkuId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHSkusSkuId200Response struct {
	value *PATCHSkusSkuId200Response
	isSet bool
}

func (v NullablePATCHSkusSkuId200Response) Get() *PATCHSkusSkuId200Response {
	return v.value
}

func (v *NullablePATCHSkusSkuId200Response) Set(val *PATCHSkusSkuId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHSkusSkuId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHSkusSkuId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHSkusSkuId200Response(val *PATCHSkusSkuId200Response) *NullablePATCHSkusSkuId200Response {
	return &NullablePATCHSkusSkuId200Response{value: val, isSet: true}
}

func (v NullablePATCHSkusSkuId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHSkusSkuId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
