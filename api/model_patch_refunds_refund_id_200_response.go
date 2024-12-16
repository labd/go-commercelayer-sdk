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

// checks if the PATCHRefundsRefundId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHRefundsRefundId200Response{}

// PATCHRefundsRefundId200Response struct for PATCHRefundsRefundId200Response
type PATCHRefundsRefundId200Response struct {
	Data *PATCHRefundsRefundId200ResponseData `json:"data,omitempty"`
}

// NewPATCHRefundsRefundId200Response instantiates a new PATCHRefundsRefundId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHRefundsRefundId200Response() *PATCHRefundsRefundId200Response {
	this := PATCHRefundsRefundId200Response{}
	return &this
}

// NewPATCHRefundsRefundId200ResponseWithDefaults instantiates a new PATCHRefundsRefundId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHRefundsRefundId200ResponseWithDefaults() *PATCHRefundsRefundId200Response {
	this := PATCHRefundsRefundId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHRefundsRefundId200Response) GetData() PATCHRefundsRefundId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHRefundsRefundId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHRefundsRefundId200Response) GetDataOk() (*PATCHRefundsRefundId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHRefundsRefundId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHRefundsRefundId200ResponseData and assigns it to the Data field.
func (o *PATCHRefundsRefundId200Response) SetData(v PATCHRefundsRefundId200ResponseData) {
	o.Data = &v
}

func (o PATCHRefundsRefundId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHRefundsRefundId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHRefundsRefundId200Response struct {
	value *PATCHRefundsRefundId200Response
	isSet bool
}

func (v NullablePATCHRefundsRefundId200Response) Get() *PATCHRefundsRefundId200Response {
	return v.value
}

func (v *NullablePATCHRefundsRefundId200Response) Set(val *PATCHRefundsRefundId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHRefundsRefundId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHRefundsRefundId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHRefundsRefundId200Response(val *PATCHRefundsRefundId200Response) *NullablePATCHRefundsRefundId200Response {
	return &NullablePATCHRefundsRefundId200Response{value: val, isSet: true}
}

func (v NullablePATCHRefundsRefundId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHRefundsRefundId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
