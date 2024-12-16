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

// checks if the GETCarrierAccountsCarrierAccountId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETCarrierAccountsCarrierAccountId200Response{}

// GETCarrierAccountsCarrierAccountId200Response struct for GETCarrierAccountsCarrierAccountId200Response
type GETCarrierAccountsCarrierAccountId200Response struct {
	Data *GETCarrierAccountsCarrierAccountId200ResponseData `json:"data,omitempty"`
}

// NewGETCarrierAccountsCarrierAccountId200Response instantiates a new GETCarrierAccountsCarrierAccountId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCarrierAccountsCarrierAccountId200Response() *GETCarrierAccountsCarrierAccountId200Response {
	this := GETCarrierAccountsCarrierAccountId200Response{}
	return &this
}

// NewGETCarrierAccountsCarrierAccountId200ResponseWithDefaults instantiates a new GETCarrierAccountsCarrierAccountId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCarrierAccountsCarrierAccountId200ResponseWithDefaults() *GETCarrierAccountsCarrierAccountId200Response {
	this := GETCarrierAccountsCarrierAccountId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCarrierAccountsCarrierAccountId200Response) GetData() GETCarrierAccountsCarrierAccountId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETCarrierAccountsCarrierAccountId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCarrierAccountsCarrierAccountId200Response) GetDataOk() (*GETCarrierAccountsCarrierAccountId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCarrierAccountsCarrierAccountId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETCarrierAccountsCarrierAccountId200ResponseData and assigns it to the Data field.
func (o *GETCarrierAccountsCarrierAccountId200Response) SetData(v GETCarrierAccountsCarrierAccountId200ResponseData) {
	o.Data = &v
}

func (o GETCarrierAccountsCarrierAccountId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETCarrierAccountsCarrierAccountId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETCarrierAccountsCarrierAccountId200Response struct {
	value *GETCarrierAccountsCarrierAccountId200Response
	isSet bool
}

func (v NullableGETCarrierAccountsCarrierAccountId200Response) Get() *GETCarrierAccountsCarrierAccountId200Response {
	return v.value
}

func (v *NullableGETCarrierAccountsCarrierAccountId200Response) Set(val *GETCarrierAccountsCarrierAccountId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCarrierAccountsCarrierAccountId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCarrierAccountsCarrierAccountId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCarrierAccountsCarrierAccountId200Response(val *GETCarrierAccountsCarrierAccountId200Response) *NullableGETCarrierAccountsCarrierAccountId200Response {
	return &NullableGETCarrierAccountsCarrierAccountId200Response{value: val, isSet: true}
}

func (v NullableGETCarrierAccountsCarrierAccountId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCarrierAccountsCarrierAccountId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
