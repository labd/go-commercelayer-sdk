/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.5.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GETStockTransfersStockTransferId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETStockTransfersStockTransferId200Response{}

// GETStockTransfersStockTransferId200Response struct for GETStockTransfersStockTransferId200Response
type GETStockTransfersStockTransferId200Response struct {
	Data *GETStockTransfersStockTransferId200ResponseData `json:"data,omitempty"`
}

// NewGETStockTransfersStockTransferId200Response instantiates a new GETStockTransfersStockTransferId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETStockTransfersStockTransferId200Response() *GETStockTransfersStockTransferId200Response {
	this := GETStockTransfersStockTransferId200Response{}
	return &this
}

// NewGETStockTransfersStockTransferId200ResponseWithDefaults instantiates a new GETStockTransfersStockTransferId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETStockTransfersStockTransferId200ResponseWithDefaults() *GETStockTransfersStockTransferId200Response {
	this := GETStockTransfersStockTransferId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETStockTransfersStockTransferId200Response) GetData() GETStockTransfersStockTransferId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETStockTransfersStockTransferId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockTransfersStockTransferId200Response) GetDataOk() (*GETStockTransfersStockTransferId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETStockTransfersStockTransferId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETStockTransfersStockTransferId200ResponseData and assigns it to the Data field.
func (o *GETStockTransfersStockTransferId200Response) SetData(v GETStockTransfersStockTransferId200ResponseData) {
	o.Data = &v
}

func (o GETStockTransfersStockTransferId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETStockTransfersStockTransferId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETStockTransfersStockTransferId200Response struct {
	value *GETStockTransfersStockTransferId200Response
	isSet bool
}

func (v NullableGETStockTransfersStockTransferId200Response) Get() *GETStockTransfersStockTransferId200Response {
	return v.value
}

func (v *NullableGETStockTransfersStockTransferId200Response) Set(val *GETStockTransfersStockTransferId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETStockTransfersStockTransferId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETStockTransfersStockTransferId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETStockTransfersStockTransferId200Response(val *GETStockTransfersStockTransferId200Response) *NullableGETStockTransfersStockTransferId200Response {
	return &NullableGETStockTransfersStockTransferId200Response{value: val, isSet: true}
}

func (v NullableGETStockTransfersStockTransferId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETStockTransfersStockTransferId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
