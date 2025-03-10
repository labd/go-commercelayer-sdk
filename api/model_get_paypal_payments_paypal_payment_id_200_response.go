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

// checks if the GETPaypalPaymentsPaypalPaymentId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETPaypalPaymentsPaypalPaymentId200Response{}

// GETPaypalPaymentsPaypalPaymentId200Response struct for GETPaypalPaymentsPaypalPaymentId200Response
type GETPaypalPaymentsPaypalPaymentId200Response struct {
	Data *GETPaypalPaymentsPaypalPaymentId200ResponseData `json:"data,omitempty"`
}

// NewGETPaypalPaymentsPaypalPaymentId200Response instantiates a new GETPaypalPaymentsPaypalPaymentId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPaypalPaymentsPaypalPaymentId200Response() *GETPaypalPaymentsPaypalPaymentId200Response {
	this := GETPaypalPaymentsPaypalPaymentId200Response{}
	return &this
}

// NewGETPaypalPaymentsPaypalPaymentId200ResponseWithDefaults instantiates a new GETPaypalPaymentsPaypalPaymentId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPaypalPaymentsPaypalPaymentId200ResponseWithDefaults() *GETPaypalPaymentsPaypalPaymentId200Response {
	this := GETPaypalPaymentsPaypalPaymentId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETPaypalPaymentsPaypalPaymentId200Response) GetData() GETPaypalPaymentsPaypalPaymentId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETPaypalPaymentsPaypalPaymentId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPaypalPaymentsPaypalPaymentId200Response) GetDataOk() (*GETPaypalPaymentsPaypalPaymentId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETPaypalPaymentsPaypalPaymentId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETPaypalPaymentsPaypalPaymentId200ResponseData and assigns it to the Data field.
func (o *GETPaypalPaymentsPaypalPaymentId200Response) SetData(v GETPaypalPaymentsPaypalPaymentId200ResponseData) {
	o.Data = &v
}

func (o GETPaypalPaymentsPaypalPaymentId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETPaypalPaymentsPaypalPaymentId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETPaypalPaymentsPaypalPaymentId200Response struct {
	value *GETPaypalPaymentsPaypalPaymentId200Response
	isSet bool
}

func (v NullableGETPaypalPaymentsPaypalPaymentId200Response) Get() *GETPaypalPaymentsPaypalPaymentId200Response {
	return v.value
}

func (v *NullableGETPaypalPaymentsPaypalPaymentId200Response) Set(val *GETPaypalPaymentsPaypalPaymentId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPaypalPaymentsPaypalPaymentId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPaypalPaymentsPaypalPaymentId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPaypalPaymentsPaypalPaymentId200Response(val *GETPaypalPaymentsPaypalPaymentId200Response) *NullableGETPaypalPaymentsPaypalPaymentId200Response {
	return &NullableGETPaypalPaymentsPaypalPaymentId200Response{value: val, isSet: true}
}

func (v NullableGETPaypalPaymentsPaypalPaymentId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPaypalPaymentsPaypalPaymentId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
