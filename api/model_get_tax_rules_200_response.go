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

// GETTaxRules200Response struct for GETTaxRules200Response
type GETTaxRules200Response struct {
	Data []GETTaxRules200ResponseDataInner `json:"data,omitempty"`
}

// NewGETTaxRules200Response instantiates a new GETTaxRules200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETTaxRules200Response() *GETTaxRules200Response {
	this := GETTaxRules200Response{}
	return &this
}

// NewGETTaxRules200ResponseWithDefaults instantiates a new GETTaxRules200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETTaxRules200ResponseWithDefaults() *GETTaxRules200Response {
	this := GETTaxRules200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETTaxRules200Response) GetData() []GETTaxRules200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETTaxRules200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTaxRules200Response) GetDataOk() ([]GETTaxRules200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETTaxRules200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETTaxRules200ResponseDataInner and assigns it to the Data field.
func (o *GETTaxRules200Response) SetData(v []GETTaxRules200ResponseDataInner) {
	o.Data = v
}

func (o GETTaxRules200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETTaxRules200Response struct {
	value *GETTaxRules200Response
	isSet bool
}

func (v NullableGETTaxRules200Response) Get() *GETTaxRules200Response {
	return v.value
}

func (v *NullableGETTaxRules200Response) Set(val *GETTaxRules200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETTaxRules200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETTaxRules200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETTaxRules200Response(val *GETTaxRules200Response) *NullableGETTaxRules200Response {
	return &NullableGETTaxRules200Response{value: val, isSet: true}
}

func (v NullableGETTaxRules200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETTaxRules200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
