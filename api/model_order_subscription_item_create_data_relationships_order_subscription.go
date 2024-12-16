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

// checks if the OrderSubscriptionItemCreateDataRelationshipsOrderSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSubscriptionItemCreateDataRelationshipsOrderSubscription{}

// OrderSubscriptionItemCreateDataRelationshipsOrderSubscription struct for OrderSubscriptionItemCreateDataRelationshipsOrderSubscription
type OrderSubscriptionItemCreateDataRelationshipsOrderSubscription struct {
	Data CustomerDataRelationshipsOrderSubscriptionsData `json:"data"`
}

// NewOrderSubscriptionItemCreateDataRelationshipsOrderSubscription instantiates a new OrderSubscriptionItemCreateDataRelationshipsOrderSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSubscriptionItemCreateDataRelationshipsOrderSubscription(data CustomerDataRelationshipsOrderSubscriptionsData) *OrderSubscriptionItemCreateDataRelationshipsOrderSubscription {
	this := OrderSubscriptionItemCreateDataRelationshipsOrderSubscription{}
	this.Data = data
	return &this
}

// NewOrderSubscriptionItemCreateDataRelationshipsOrderSubscriptionWithDefaults instantiates a new OrderSubscriptionItemCreateDataRelationshipsOrderSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSubscriptionItemCreateDataRelationshipsOrderSubscriptionWithDefaults() *OrderSubscriptionItemCreateDataRelationshipsOrderSubscription {
	this := OrderSubscriptionItemCreateDataRelationshipsOrderSubscription{}
	return &this
}

// GetData returns the Data field value
func (o *OrderSubscriptionItemCreateDataRelationshipsOrderSubscription) GetData() CustomerDataRelationshipsOrderSubscriptionsData {
	if o == nil {
		var ret CustomerDataRelationshipsOrderSubscriptionsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionItemCreateDataRelationshipsOrderSubscription) GetDataOk() (*CustomerDataRelationshipsOrderSubscriptionsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *OrderSubscriptionItemCreateDataRelationshipsOrderSubscription) SetData(v CustomerDataRelationshipsOrderSubscriptionsData) {
	o.Data = v
}

func (o OrderSubscriptionItemCreateDataRelationshipsOrderSubscription) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSubscriptionItemCreateDataRelationshipsOrderSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableOrderSubscriptionItemCreateDataRelationshipsOrderSubscription struct {
	value *OrderSubscriptionItemCreateDataRelationshipsOrderSubscription
	isSet bool
}

func (v NullableOrderSubscriptionItemCreateDataRelationshipsOrderSubscription) Get() *OrderSubscriptionItemCreateDataRelationshipsOrderSubscription {
	return v.value
}

func (v *NullableOrderSubscriptionItemCreateDataRelationshipsOrderSubscription) Set(val *OrderSubscriptionItemCreateDataRelationshipsOrderSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSubscriptionItemCreateDataRelationshipsOrderSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSubscriptionItemCreateDataRelationshipsOrderSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSubscriptionItemCreateDataRelationshipsOrderSubscription(val *OrderSubscriptionItemCreateDataRelationshipsOrderSubscription) *NullableOrderSubscriptionItemCreateDataRelationshipsOrderSubscription {
	return &NullableOrderSubscriptionItemCreateDataRelationshipsOrderSubscription{value: val, isSet: true}
}

func (v NullableOrderSubscriptionItemCreateDataRelationshipsOrderSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSubscriptionItemCreateDataRelationshipsOrderSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
