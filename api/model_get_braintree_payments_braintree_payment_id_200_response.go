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

// GETBraintreePaymentsBraintreePaymentId200Response struct for GETBraintreePaymentsBraintreePaymentId200Response
type GETBraintreePaymentsBraintreePaymentId200Response struct {
	Data *GETBraintreePayments200ResponseDataInner `json:"data,omitempty"`
}

// NewGETBraintreePaymentsBraintreePaymentId200Response instantiates a new GETBraintreePaymentsBraintreePaymentId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETBraintreePaymentsBraintreePaymentId200Response() *GETBraintreePaymentsBraintreePaymentId200Response {
	this := GETBraintreePaymentsBraintreePaymentId200Response{}
	return &this
}

// NewGETBraintreePaymentsBraintreePaymentId200ResponseWithDefaults instantiates a new GETBraintreePaymentsBraintreePaymentId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETBraintreePaymentsBraintreePaymentId200ResponseWithDefaults() *GETBraintreePaymentsBraintreePaymentId200Response {
	this := GETBraintreePaymentsBraintreePaymentId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETBraintreePaymentsBraintreePaymentId200Response) GetData() GETBraintreePayments200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETBraintreePayments200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBraintreePaymentsBraintreePaymentId200Response) GetDataOk() (*GETBraintreePayments200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETBraintreePaymentsBraintreePaymentId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETBraintreePayments200ResponseDataInner and assigns it to the Data field.
func (o *GETBraintreePaymentsBraintreePaymentId200Response) SetData(v GETBraintreePayments200ResponseDataInner) {
	o.Data = &v
}

func (o GETBraintreePaymentsBraintreePaymentId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETBraintreePaymentsBraintreePaymentId200Response struct {
	value *GETBraintreePaymentsBraintreePaymentId200Response
	isSet bool
}

func (v NullableGETBraintreePaymentsBraintreePaymentId200Response) Get() *GETBraintreePaymentsBraintreePaymentId200Response {
	return v.value
}

func (v *NullableGETBraintreePaymentsBraintreePaymentId200Response) Set(val *GETBraintreePaymentsBraintreePaymentId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETBraintreePaymentsBraintreePaymentId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETBraintreePaymentsBraintreePaymentId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETBraintreePaymentsBraintreePaymentId200Response(val *GETBraintreePaymentsBraintreePaymentId200Response) *NullableGETBraintreePaymentsBraintreePaymentId200Response {
	return &NullableGETBraintreePaymentsBraintreePaymentId200Response{value: val, isSet: true}
}

func (v NullableGETBraintreePaymentsBraintreePaymentId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETBraintreePaymentsBraintreePaymentId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
