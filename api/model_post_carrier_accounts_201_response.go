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

// checks if the POSTCarrierAccounts201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCarrierAccounts201Response{}

// POSTCarrierAccounts201Response struct for POSTCarrierAccounts201Response
type POSTCarrierAccounts201Response struct {
	Data *POSTCarrierAccounts201ResponseData `json:"data,omitempty"`
}

// NewPOSTCarrierAccounts201Response instantiates a new POSTCarrierAccounts201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCarrierAccounts201Response() *POSTCarrierAccounts201Response {
	this := POSTCarrierAccounts201Response{}
	return &this
}

// NewPOSTCarrierAccounts201ResponseWithDefaults instantiates a new POSTCarrierAccounts201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCarrierAccounts201ResponseWithDefaults() *POSTCarrierAccounts201Response {
	this := POSTCarrierAccounts201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTCarrierAccounts201Response) GetData() POSTCarrierAccounts201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTCarrierAccounts201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCarrierAccounts201Response) GetDataOk() (*POSTCarrierAccounts201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTCarrierAccounts201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTCarrierAccounts201ResponseData and assigns it to the Data field.
func (o *POSTCarrierAccounts201Response) SetData(v POSTCarrierAccounts201ResponseData) {
	o.Data = &v
}

func (o POSTCarrierAccounts201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCarrierAccounts201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTCarrierAccounts201Response struct {
	value *POSTCarrierAccounts201Response
	isSet bool
}

func (v NullablePOSTCarrierAccounts201Response) Get() *POSTCarrierAccounts201Response {
	return v.value
}

func (v *NullablePOSTCarrierAccounts201Response) Set(val *POSTCarrierAccounts201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCarrierAccounts201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCarrierAccounts201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCarrierAccounts201Response(val *POSTCarrierAccounts201Response) *NullablePOSTCarrierAccounts201Response {
	return &NullablePOSTCarrierAccounts201Response{value: val, isSet: true}
}

func (v NullablePOSTCarrierAccounts201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCarrierAccounts201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
