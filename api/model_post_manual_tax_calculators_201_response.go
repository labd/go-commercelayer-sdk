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

// checks if the POSTManualTaxCalculators201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTManualTaxCalculators201Response{}

// POSTManualTaxCalculators201Response struct for POSTManualTaxCalculators201Response
type POSTManualTaxCalculators201Response struct {
	Data *POSTManualTaxCalculators201ResponseData `json:"data,omitempty"`
}

// NewPOSTManualTaxCalculators201Response instantiates a new POSTManualTaxCalculators201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTManualTaxCalculators201Response() *POSTManualTaxCalculators201Response {
	this := POSTManualTaxCalculators201Response{}
	return &this
}

// NewPOSTManualTaxCalculators201ResponseWithDefaults instantiates a new POSTManualTaxCalculators201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTManualTaxCalculators201ResponseWithDefaults() *POSTManualTaxCalculators201Response {
	this := POSTManualTaxCalculators201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTManualTaxCalculators201Response) GetData() POSTManualTaxCalculators201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTManualTaxCalculators201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTManualTaxCalculators201Response) GetDataOk() (*POSTManualTaxCalculators201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTManualTaxCalculators201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTManualTaxCalculators201ResponseData and assigns it to the Data field.
func (o *POSTManualTaxCalculators201Response) SetData(v POSTManualTaxCalculators201ResponseData) {
	o.Data = &v
}

func (o POSTManualTaxCalculators201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTManualTaxCalculators201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTManualTaxCalculators201Response struct {
	value *POSTManualTaxCalculators201Response
	isSet bool
}

func (v NullablePOSTManualTaxCalculators201Response) Get() *POSTManualTaxCalculators201Response {
	return v.value
}

func (v *NullablePOSTManualTaxCalculators201Response) Set(val *POSTManualTaxCalculators201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTManualTaxCalculators201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTManualTaxCalculators201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTManualTaxCalculators201Response(val *POSTManualTaxCalculators201Response) *NullablePOSTManualTaxCalculators201Response {
	return &NullablePOSTManualTaxCalculators201Response{value: val, isSet: true}
}

func (v NullablePOSTManualTaxCalculators201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTManualTaxCalculators201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
