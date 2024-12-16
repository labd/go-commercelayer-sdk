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

// checks if the GETManualTaxCalculatorsManualTaxCalculatorId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETManualTaxCalculatorsManualTaxCalculatorId200Response{}

// GETManualTaxCalculatorsManualTaxCalculatorId200Response struct for GETManualTaxCalculatorsManualTaxCalculatorId200Response
type GETManualTaxCalculatorsManualTaxCalculatorId200Response struct {
	Data *GETManualTaxCalculatorsManualTaxCalculatorId200ResponseData `json:"data,omitempty"`
}

// NewGETManualTaxCalculatorsManualTaxCalculatorId200Response instantiates a new GETManualTaxCalculatorsManualTaxCalculatorId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETManualTaxCalculatorsManualTaxCalculatorId200Response() *GETManualTaxCalculatorsManualTaxCalculatorId200Response {
	this := GETManualTaxCalculatorsManualTaxCalculatorId200Response{}
	return &this
}

// NewGETManualTaxCalculatorsManualTaxCalculatorId200ResponseWithDefaults instantiates a new GETManualTaxCalculatorsManualTaxCalculatorId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETManualTaxCalculatorsManualTaxCalculatorId200ResponseWithDefaults() *GETManualTaxCalculatorsManualTaxCalculatorId200Response {
	this := GETManualTaxCalculatorsManualTaxCalculatorId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETManualTaxCalculatorsManualTaxCalculatorId200Response) GetData() GETManualTaxCalculatorsManualTaxCalculatorId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETManualTaxCalculatorsManualTaxCalculatorId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETManualTaxCalculatorsManualTaxCalculatorId200Response) GetDataOk() (*GETManualTaxCalculatorsManualTaxCalculatorId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETManualTaxCalculatorsManualTaxCalculatorId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETManualTaxCalculatorsManualTaxCalculatorId200ResponseData and assigns it to the Data field.
func (o *GETManualTaxCalculatorsManualTaxCalculatorId200Response) SetData(v GETManualTaxCalculatorsManualTaxCalculatorId200ResponseData) {
	o.Data = &v
}

func (o GETManualTaxCalculatorsManualTaxCalculatorId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETManualTaxCalculatorsManualTaxCalculatorId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETManualTaxCalculatorsManualTaxCalculatorId200Response struct {
	value *GETManualTaxCalculatorsManualTaxCalculatorId200Response
	isSet bool
}

func (v NullableGETManualTaxCalculatorsManualTaxCalculatorId200Response) Get() *GETManualTaxCalculatorsManualTaxCalculatorId200Response {
	return v.value
}

func (v *NullableGETManualTaxCalculatorsManualTaxCalculatorId200Response) Set(val *GETManualTaxCalculatorsManualTaxCalculatorId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETManualTaxCalculatorsManualTaxCalculatorId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETManualTaxCalculatorsManualTaxCalculatorId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETManualTaxCalculatorsManualTaxCalculatorId200Response(val *GETManualTaxCalculatorsManualTaxCalculatorId200Response) *NullableGETManualTaxCalculatorsManualTaxCalculatorId200Response {
	return &NullableGETManualTaxCalculatorsManualTaxCalculatorId200Response{value: val, isSet: true}
}

func (v NullableGETManualTaxCalculatorsManualTaxCalculatorId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETManualTaxCalculatorsManualTaxCalculatorId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
