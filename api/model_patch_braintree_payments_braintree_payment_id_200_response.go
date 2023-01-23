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

// PATCHBraintreePaymentsBraintreePaymentId200Response struct for PATCHBraintreePaymentsBraintreePaymentId200Response
type PATCHBraintreePaymentsBraintreePaymentId200Response struct {
	Data *PATCHBraintreePaymentsBraintreePaymentId200ResponseData `json:"data,omitempty"`
}

// NewPATCHBraintreePaymentsBraintreePaymentId200Response instantiates a new PATCHBraintreePaymentsBraintreePaymentId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHBraintreePaymentsBraintreePaymentId200Response() *PATCHBraintreePaymentsBraintreePaymentId200Response {
	this := PATCHBraintreePaymentsBraintreePaymentId200Response{}
	return &this
}

// NewPATCHBraintreePaymentsBraintreePaymentId200ResponseWithDefaults instantiates a new PATCHBraintreePaymentsBraintreePaymentId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHBraintreePaymentsBraintreePaymentId200ResponseWithDefaults() *PATCHBraintreePaymentsBraintreePaymentId200Response {
	this := PATCHBraintreePaymentsBraintreePaymentId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHBraintreePaymentsBraintreePaymentId200Response) GetData() PATCHBraintreePaymentsBraintreePaymentId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHBraintreePaymentsBraintreePaymentId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHBraintreePaymentsBraintreePaymentId200Response) GetDataOk() (*PATCHBraintreePaymentsBraintreePaymentId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHBraintreePaymentsBraintreePaymentId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHBraintreePaymentsBraintreePaymentId200ResponseData and assigns it to the Data field.
func (o *PATCHBraintreePaymentsBraintreePaymentId200Response) SetData(v PATCHBraintreePaymentsBraintreePaymentId200ResponseData) {
	o.Data = &v
}

func (o PATCHBraintreePaymentsBraintreePaymentId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHBraintreePaymentsBraintreePaymentId200Response struct {
	value *PATCHBraintreePaymentsBraintreePaymentId200Response
	isSet bool
}

func (v NullablePATCHBraintreePaymentsBraintreePaymentId200Response) Get() *PATCHBraintreePaymentsBraintreePaymentId200Response {
	return v.value
}

func (v *NullablePATCHBraintreePaymentsBraintreePaymentId200Response) Set(val *PATCHBraintreePaymentsBraintreePaymentId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHBraintreePaymentsBraintreePaymentId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHBraintreePaymentsBraintreePaymentId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHBraintreePaymentsBraintreePaymentId200Response(val *PATCHBraintreePaymentsBraintreePaymentId200Response) *NullablePATCHBraintreePaymentsBraintreePaymentId200Response {
	return &NullablePATCHBraintreePaymentsBraintreePaymentId200Response{value: val, isSet: true}
}

func (v NullablePATCHBraintreePaymentsBraintreePaymentId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHBraintreePaymentsBraintreePaymentId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
