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

// GETShippingMethodTiersShippingMethodTierId200Response struct for GETShippingMethodTiersShippingMethodTierId200Response
type GETShippingMethodTiersShippingMethodTierId200Response struct {
	Data *GETShippingMethodTiers200ResponseDataInner `json:"data,omitempty"`
}

// NewGETShippingMethodTiersShippingMethodTierId200Response instantiates a new GETShippingMethodTiersShippingMethodTierId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETShippingMethodTiersShippingMethodTierId200Response() *GETShippingMethodTiersShippingMethodTierId200Response {
	this := GETShippingMethodTiersShippingMethodTierId200Response{}
	return &this
}

// NewGETShippingMethodTiersShippingMethodTierId200ResponseWithDefaults instantiates a new GETShippingMethodTiersShippingMethodTierId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETShippingMethodTiersShippingMethodTierId200ResponseWithDefaults() *GETShippingMethodTiersShippingMethodTierId200Response {
	this := GETShippingMethodTiersShippingMethodTierId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETShippingMethodTiersShippingMethodTierId200Response) GetData() GETShippingMethodTiers200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETShippingMethodTiers200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShippingMethodTiersShippingMethodTierId200Response) GetDataOk() (*GETShippingMethodTiers200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETShippingMethodTiersShippingMethodTierId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETShippingMethodTiers200ResponseDataInner and assigns it to the Data field.
func (o *GETShippingMethodTiersShippingMethodTierId200Response) SetData(v GETShippingMethodTiers200ResponseDataInner) {
	o.Data = &v
}

func (o GETShippingMethodTiersShippingMethodTierId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETShippingMethodTiersShippingMethodTierId200Response struct {
	value *GETShippingMethodTiersShippingMethodTierId200Response
	isSet bool
}

func (v NullableGETShippingMethodTiersShippingMethodTierId200Response) Get() *GETShippingMethodTiersShippingMethodTierId200Response {
	return v.value
}

func (v *NullableGETShippingMethodTiersShippingMethodTierId200Response) Set(val *GETShippingMethodTiersShippingMethodTierId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETShippingMethodTiersShippingMethodTierId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETShippingMethodTiersShippingMethodTierId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETShippingMethodTiersShippingMethodTierId200Response(val *GETShippingMethodTiersShippingMethodTierId200Response) *NullableGETShippingMethodTiersShippingMethodTierId200Response {
	return &NullableGETShippingMethodTiersShippingMethodTierId200Response{value: val, isSet: true}
}

func (v NullableGETShippingMethodTiersShippingMethodTierId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETShippingMethodTiersShippingMethodTierId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}