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

// DeliveryLeadTime struct for DeliveryLeadTime
type DeliveryLeadTime struct {
	Data *DeliveryLeadTimeData `json:"data,omitempty"`
}

// NewDeliveryLeadTime instantiates a new DeliveryLeadTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryLeadTime() *DeliveryLeadTime {
	this := DeliveryLeadTime{}
	return &this
}

// NewDeliveryLeadTimeWithDefaults instantiates a new DeliveryLeadTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryLeadTimeWithDefaults() *DeliveryLeadTime {
	this := DeliveryLeadTime{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DeliveryLeadTime) GetData() DeliveryLeadTimeData {
	if o == nil || o.Data == nil {
		var ret DeliveryLeadTimeData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTime) GetDataOk() (*DeliveryLeadTimeData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DeliveryLeadTime) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given DeliveryLeadTimeData and assigns it to the Data field.
func (o *DeliveryLeadTime) SetData(v DeliveryLeadTimeData) {
	o.Data = &v
}

func (o DeliveryLeadTime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableDeliveryLeadTime struct {
	value *DeliveryLeadTime
	isSet bool
}

func (v NullableDeliveryLeadTime) Get() *DeliveryLeadTime {
	return v.value
}

func (v *NullableDeliveryLeadTime) Set(val *DeliveryLeadTime) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryLeadTime) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryLeadTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryLeadTime(val *DeliveryLeadTime) *NullableDeliveryLeadTime {
	return &NullableDeliveryLeadTime{value: val, isSet: true}
}

func (v NullableDeliveryLeadTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryLeadTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
