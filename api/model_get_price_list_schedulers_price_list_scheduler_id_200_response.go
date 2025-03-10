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

// checks if the GETPriceListSchedulersPriceListSchedulerId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETPriceListSchedulersPriceListSchedulerId200Response{}

// GETPriceListSchedulersPriceListSchedulerId200Response struct for GETPriceListSchedulersPriceListSchedulerId200Response
type GETPriceListSchedulersPriceListSchedulerId200Response struct {
	Data *GETPriceListSchedulersPriceListSchedulerId200ResponseData `json:"data,omitempty"`
}

// NewGETPriceListSchedulersPriceListSchedulerId200Response instantiates a new GETPriceListSchedulersPriceListSchedulerId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPriceListSchedulersPriceListSchedulerId200Response() *GETPriceListSchedulersPriceListSchedulerId200Response {
	this := GETPriceListSchedulersPriceListSchedulerId200Response{}
	return &this
}

// NewGETPriceListSchedulersPriceListSchedulerId200ResponseWithDefaults instantiates a new GETPriceListSchedulersPriceListSchedulerId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPriceListSchedulersPriceListSchedulerId200ResponseWithDefaults() *GETPriceListSchedulersPriceListSchedulerId200Response {
	this := GETPriceListSchedulersPriceListSchedulerId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETPriceListSchedulersPriceListSchedulerId200Response) GetData() GETPriceListSchedulersPriceListSchedulerId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETPriceListSchedulersPriceListSchedulerId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPriceListSchedulersPriceListSchedulerId200Response) GetDataOk() (*GETPriceListSchedulersPriceListSchedulerId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETPriceListSchedulersPriceListSchedulerId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETPriceListSchedulersPriceListSchedulerId200ResponseData and assigns it to the Data field.
func (o *GETPriceListSchedulersPriceListSchedulerId200Response) SetData(v GETPriceListSchedulersPriceListSchedulerId200ResponseData) {
	o.Data = &v
}

func (o GETPriceListSchedulersPriceListSchedulerId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETPriceListSchedulersPriceListSchedulerId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETPriceListSchedulersPriceListSchedulerId200Response struct {
	value *GETPriceListSchedulersPriceListSchedulerId200Response
	isSet bool
}

func (v NullableGETPriceListSchedulersPriceListSchedulerId200Response) Get() *GETPriceListSchedulersPriceListSchedulerId200Response {
	return v.value
}

func (v *NullableGETPriceListSchedulersPriceListSchedulerId200Response) Set(val *GETPriceListSchedulersPriceListSchedulerId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPriceListSchedulersPriceListSchedulerId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPriceListSchedulersPriceListSchedulerId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPriceListSchedulersPriceListSchedulerId200Response(val *GETPriceListSchedulersPriceListSchedulerId200Response) *NullableGETPriceListSchedulersPriceListSchedulerId200Response {
	return &NullableGETPriceListSchedulersPriceListSchedulerId200Response{value: val, isSet: true}
}

func (v NullableGETPriceListSchedulersPriceListSchedulerId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPriceListSchedulersPriceListSchedulerId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
