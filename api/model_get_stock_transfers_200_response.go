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

// GETStockTransfers200Response struct for GETStockTransfers200Response
type GETStockTransfers200Response struct {
	Data []GETStockTransfers200ResponseDataInner `json:"data,omitempty"`
}

// NewGETStockTransfers200Response instantiates a new GETStockTransfers200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETStockTransfers200Response() *GETStockTransfers200Response {
	this := GETStockTransfers200Response{}
	return &this
}

// NewGETStockTransfers200ResponseWithDefaults instantiates a new GETStockTransfers200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETStockTransfers200ResponseWithDefaults() *GETStockTransfers200Response {
	this := GETStockTransfers200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETStockTransfers200Response) GetData() []GETStockTransfers200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETStockTransfers200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockTransfers200Response) GetDataOk() ([]GETStockTransfers200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETStockTransfers200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETStockTransfers200ResponseDataInner and assigns it to the Data field.
func (o *GETStockTransfers200Response) SetData(v []GETStockTransfers200ResponseDataInner) {
	o.Data = v
}

func (o GETStockTransfers200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETStockTransfers200Response struct {
	value *GETStockTransfers200Response
	isSet bool
}

func (v NullableGETStockTransfers200Response) Get() *GETStockTransfers200Response {
	return v.value
}

func (v *NullableGETStockTransfers200Response) Set(val *GETStockTransfers200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETStockTransfers200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETStockTransfers200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETStockTransfers200Response(val *GETStockTransfers200Response) *NullableGETStockTransfers200Response {
	return &NullableGETStockTransfers200Response{value: val, isSet: true}
}

func (v NullableGETStockTransfers200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETStockTransfers200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}