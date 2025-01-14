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

// checks if the StripeGatewayUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StripeGatewayUpdate{}

// StripeGatewayUpdate struct for StripeGatewayUpdate
type StripeGatewayUpdate struct {
	Data StripeGatewayUpdateData `json:"data"`
}

// NewStripeGatewayUpdate instantiates a new StripeGatewayUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripeGatewayUpdate(data StripeGatewayUpdateData) *StripeGatewayUpdate {
	this := StripeGatewayUpdate{}
	this.Data = data
	return &this
}

// NewStripeGatewayUpdateWithDefaults instantiates a new StripeGatewayUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripeGatewayUpdateWithDefaults() *StripeGatewayUpdate {
	this := StripeGatewayUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *StripeGatewayUpdate) GetData() StripeGatewayUpdateData {
	if o == nil {
		var ret StripeGatewayUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *StripeGatewayUpdate) GetDataOk() (*StripeGatewayUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *StripeGatewayUpdate) SetData(v StripeGatewayUpdateData) {
	o.Data = v
}

func (o StripeGatewayUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StripeGatewayUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableStripeGatewayUpdate struct {
	value *StripeGatewayUpdate
	isSet bool
}

func (v NullableStripeGatewayUpdate) Get() *StripeGatewayUpdate {
	return v.value
}

func (v *NullableStripeGatewayUpdate) Set(val *StripeGatewayUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableStripeGatewayUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableStripeGatewayUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripeGatewayUpdate(val *StripeGatewayUpdate) *NullableStripeGatewayUpdate {
	return &NullableStripeGatewayUpdate{value: val, isSet: true}
}

func (v NullableStripeGatewayUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripeGatewayUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
