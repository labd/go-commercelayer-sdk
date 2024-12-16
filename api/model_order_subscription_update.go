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

// checks if the OrderSubscriptionUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSubscriptionUpdate{}

// OrderSubscriptionUpdate struct for OrderSubscriptionUpdate
type OrderSubscriptionUpdate struct {
	Data OrderSubscriptionUpdateData `json:"data"`
}

// NewOrderSubscriptionUpdate instantiates a new OrderSubscriptionUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSubscriptionUpdate(data OrderSubscriptionUpdateData) *OrderSubscriptionUpdate {
	this := OrderSubscriptionUpdate{}
	this.Data = data
	return &this
}

// NewOrderSubscriptionUpdateWithDefaults instantiates a new OrderSubscriptionUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSubscriptionUpdateWithDefaults() *OrderSubscriptionUpdate {
	this := OrderSubscriptionUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *OrderSubscriptionUpdate) GetData() OrderSubscriptionUpdateData {
	if o == nil {
		var ret OrderSubscriptionUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionUpdate) GetDataOk() (*OrderSubscriptionUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *OrderSubscriptionUpdate) SetData(v OrderSubscriptionUpdateData) {
	o.Data = v
}

func (o OrderSubscriptionUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSubscriptionUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableOrderSubscriptionUpdate struct {
	value *OrderSubscriptionUpdate
	isSet bool
}

func (v NullableOrderSubscriptionUpdate) Get() *OrderSubscriptionUpdate {
	return v.value
}

func (v *NullableOrderSubscriptionUpdate) Set(val *OrderSubscriptionUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSubscriptionUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSubscriptionUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSubscriptionUpdate(val *OrderSubscriptionUpdate) *NullableOrderSubscriptionUpdate {
	return &NullableOrderSubscriptionUpdate{value: val, isSet: true}
}

func (v NullableOrderSubscriptionUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSubscriptionUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
