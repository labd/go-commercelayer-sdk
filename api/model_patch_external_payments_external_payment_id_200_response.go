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

// PATCHExternalPaymentsExternalPaymentId200Response struct for PATCHExternalPaymentsExternalPaymentId200Response
type PATCHExternalPaymentsExternalPaymentId200Response struct {
	Data *PATCHExternalPaymentsExternalPaymentId200ResponseData `json:"data,omitempty"`
}

// NewPATCHExternalPaymentsExternalPaymentId200Response instantiates a new PATCHExternalPaymentsExternalPaymentId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHExternalPaymentsExternalPaymentId200Response() *PATCHExternalPaymentsExternalPaymentId200Response {
	this := PATCHExternalPaymentsExternalPaymentId200Response{}
	return &this
}

// NewPATCHExternalPaymentsExternalPaymentId200ResponseWithDefaults instantiates a new PATCHExternalPaymentsExternalPaymentId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHExternalPaymentsExternalPaymentId200ResponseWithDefaults() *PATCHExternalPaymentsExternalPaymentId200Response {
	this := PATCHExternalPaymentsExternalPaymentId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHExternalPaymentsExternalPaymentId200Response) GetData() PATCHExternalPaymentsExternalPaymentId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHExternalPaymentsExternalPaymentId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHExternalPaymentsExternalPaymentId200Response) GetDataOk() (*PATCHExternalPaymentsExternalPaymentId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHExternalPaymentsExternalPaymentId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHExternalPaymentsExternalPaymentId200ResponseData and assigns it to the Data field.
func (o *PATCHExternalPaymentsExternalPaymentId200Response) SetData(v PATCHExternalPaymentsExternalPaymentId200ResponseData) {
	o.Data = &v
}

func (o PATCHExternalPaymentsExternalPaymentId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHExternalPaymentsExternalPaymentId200Response struct {
	value *PATCHExternalPaymentsExternalPaymentId200Response
	isSet bool
}

func (v NullablePATCHExternalPaymentsExternalPaymentId200Response) Get() *PATCHExternalPaymentsExternalPaymentId200Response {
	return v.value
}

func (v *NullablePATCHExternalPaymentsExternalPaymentId200Response) Set(val *PATCHExternalPaymentsExternalPaymentId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHExternalPaymentsExternalPaymentId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHExternalPaymentsExternalPaymentId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHExternalPaymentsExternalPaymentId200Response(val *PATCHExternalPaymentsExternalPaymentId200Response) *NullablePATCHExternalPaymentsExternalPaymentId200Response {
	return &NullablePATCHExternalPaymentsExternalPaymentId200Response{value: val, isSet: true}
}

func (v NullablePATCHExternalPaymentsExternalPaymentId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHExternalPaymentsExternalPaymentId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
