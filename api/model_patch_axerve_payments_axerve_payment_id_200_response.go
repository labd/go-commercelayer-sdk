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

// checks if the PATCHAxervePaymentsAxervePaymentId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHAxervePaymentsAxervePaymentId200Response{}

// PATCHAxervePaymentsAxervePaymentId200Response struct for PATCHAxervePaymentsAxervePaymentId200Response
type PATCHAxervePaymentsAxervePaymentId200Response struct {
	Data *PATCHAxervePaymentsAxervePaymentId200ResponseData `json:"data,omitempty"`
}

// NewPATCHAxervePaymentsAxervePaymentId200Response instantiates a new PATCHAxervePaymentsAxervePaymentId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHAxervePaymentsAxervePaymentId200Response() *PATCHAxervePaymentsAxervePaymentId200Response {
	this := PATCHAxervePaymentsAxervePaymentId200Response{}
	return &this
}

// NewPATCHAxervePaymentsAxervePaymentId200ResponseWithDefaults instantiates a new PATCHAxervePaymentsAxervePaymentId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHAxervePaymentsAxervePaymentId200ResponseWithDefaults() *PATCHAxervePaymentsAxervePaymentId200Response {
	this := PATCHAxervePaymentsAxervePaymentId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHAxervePaymentsAxervePaymentId200Response) GetData() PATCHAxervePaymentsAxervePaymentId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHAxervePaymentsAxervePaymentId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAxervePaymentsAxervePaymentId200Response) GetDataOk() (*PATCHAxervePaymentsAxervePaymentId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHAxervePaymentsAxervePaymentId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHAxervePaymentsAxervePaymentId200ResponseData and assigns it to the Data field.
func (o *PATCHAxervePaymentsAxervePaymentId200Response) SetData(v PATCHAxervePaymentsAxervePaymentId200ResponseData) {
	o.Data = &v
}

func (o PATCHAxervePaymentsAxervePaymentId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHAxervePaymentsAxervePaymentId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHAxervePaymentsAxervePaymentId200Response struct {
	value *PATCHAxervePaymentsAxervePaymentId200Response
	isSet bool
}

func (v NullablePATCHAxervePaymentsAxervePaymentId200Response) Get() *PATCHAxervePaymentsAxervePaymentId200Response {
	return v.value
}

func (v *NullablePATCHAxervePaymentsAxervePaymentId200Response) Set(val *PATCHAxervePaymentsAxervePaymentId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHAxervePaymentsAxervePaymentId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHAxervePaymentsAxervePaymentId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHAxervePaymentsAxervePaymentId200Response(val *PATCHAxervePaymentsAxervePaymentId200Response) *NullablePATCHAxervePaymentsAxervePaymentId200Response {
	return &NullablePATCHAxervePaymentsAxervePaymentId200Response{value: val, isSet: true}
}

func (v NullablePATCHAxervePaymentsAxervePaymentId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHAxervePaymentsAxervePaymentId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
