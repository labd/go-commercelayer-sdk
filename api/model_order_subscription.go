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

// checks if the OrderSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSubscription{}

// OrderSubscription struct for OrderSubscription
type OrderSubscription struct {
	Data *OrderSubscriptionData `json:"data,omitempty"`
}

// NewOrderSubscription instantiates a new OrderSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSubscription() *OrderSubscription {
	this := OrderSubscription{}
	return &this
}

// NewOrderSubscriptionWithDefaults instantiates a new OrderSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSubscriptionWithDefaults() *OrderSubscription {
	this := OrderSubscription{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OrderSubscription) GetData() OrderSubscriptionData {
	if o == nil || IsNil(o.Data) {
		var ret OrderSubscriptionData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscription) GetDataOk() (*OrderSubscriptionData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OrderSubscription) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given OrderSubscriptionData and assigns it to the Data field.
func (o *OrderSubscription) SetData(v OrderSubscriptionData) {
	o.Data = &v
}

func (o OrderSubscription) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableOrderSubscription struct {
	value *OrderSubscription
	isSet bool
}

func (v NullableOrderSubscription) Get() *OrderSubscription {
	return v.value
}

func (v *NullableOrderSubscription) Set(val *OrderSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSubscription(val *OrderSubscription) *NullableOrderSubscription {
	return &NullableOrderSubscription{value: val, isSet: true}
}

func (v NullableOrderSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
