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

// PATCHKlarnaPaymentsKlarnaPaymentId200Response struct for PATCHKlarnaPaymentsKlarnaPaymentId200Response
type PATCHKlarnaPaymentsKlarnaPaymentId200Response struct {
	Data *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData `json:"data,omitempty"`
}

// NewPATCHKlarnaPaymentsKlarnaPaymentId200Response instantiates a new PATCHKlarnaPaymentsKlarnaPaymentId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHKlarnaPaymentsKlarnaPaymentId200Response() *PATCHKlarnaPaymentsKlarnaPaymentId200Response {
	this := PATCHKlarnaPaymentsKlarnaPaymentId200Response{}
	return &this
}

// NewPATCHKlarnaPaymentsKlarnaPaymentId200ResponseWithDefaults instantiates a new PATCHKlarnaPaymentsKlarnaPaymentId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHKlarnaPaymentsKlarnaPaymentId200ResponseWithDefaults() *PATCHKlarnaPaymentsKlarnaPaymentId200Response {
	this := PATCHKlarnaPaymentsKlarnaPaymentId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200Response) GetData() PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200Response) GetDataOk() (*PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData and assigns it to the Data field.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200Response) SetData(v PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData) {
	o.Data = &v
}

func (o PATCHKlarnaPaymentsKlarnaPaymentId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHKlarnaPaymentsKlarnaPaymentId200Response struct {
	value *PATCHKlarnaPaymentsKlarnaPaymentId200Response
	isSet bool
}

func (v NullablePATCHKlarnaPaymentsKlarnaPaymentId200Response) Get() *PATCHKlarnaPaymentsKlarnaPaymentId200Response {
	return v.value
}

func (v *NullablePATCHKlarnaPaymentsKlarnaPaymentId200Response) Set(val *PATCHKlarnaPaymentsKlarnaPaymentId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHKlarnaPaymentsKlarnaPaymentId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHKlarnaPaymentsKlarnaPaymentId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHKlarnaPaymentsKlarnaPaymentId200Response(val *PATCHKlarnaPaymentsKlarnaPaymentId200Response) *NullablePATCHKlarnaPaymentsKlarnaPaymentId200Response {
	return &NullablePATCHKlarnaPaymentsKlarnaPaymentId200Response{value: val, isSet: true}
}

func (v NullablePATCHKlarnaPaymentsKlarnaPaymentId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHKlarnaPaymentsKlarnaPaymentId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
