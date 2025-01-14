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

// checks if the OrderSubscriptionDataRelationshipsRecurringOrderCopies type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSubscriptionDataRelationshipsRecurringOrderCopies{}

// OrderSubscriptionDataRelationshipsRecurringOrderCopies struct for OrderSubscriptionDataRelationshipsRecurringOrderCopies
type OrderSubscriptionDataRelationshipsRecurringOrderCopies struct {
	Data *OrderSubscriptionDataRelationshipsRecurringOrderCopiesData `json:"data,omitempty"`
}

// NewOrderSubscriptionDataRelationshipsRecurringOrderCopies instantiates a new OrderSubscriptionDataRelationshipsRecurringOrderCopies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSubscriptionDataRelationshipsRecurringOrderCopies() *OrderSubscriptionDataRelationshipsRecurringOrderCopies {
	this := OrderSubscriptionDataRelationshipsRecurringOrderCopies{}
	return &this
}

// NewOrderSubscriptionDataRelationshipsRecurringOrderCopiesWithDefaults instantiates a new OrderSubscriptionDataRelationshipsRecurringOrderCopies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSubscriptionDataRelationshipsRecurringOrderCopiesWithDefaults() *OrderSubscriptionDataRelationshipsRecurringOrderCopies {
	this := OrderSubscriptionDataRelationshipsRecurringOrderCopies{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OrderSubscriptionDataRelationshipsRecurringOrderCopies) GetData() OrderSubscriptionDataRelationshipsRecurringOrderCopiesData {
	if o == nil || IsNil(o.Data) {
		var ret OrderSubscriptionDataRelationshipsRecurringOrderCopiesData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionDataRelationshipsRecurringOrderCopies) GetDataOk() (*OrderSubscriptionDataRelationshipsRecurringOrderCopiesData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OrderSubscriptionDataRelationshipsRecurringOrderCopies) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given OrderSubscriptionDataRelationshipsRecurringOrderCopiesData and assigns it to the Data field.
func (o *OrderSubscriptionDataRelationshipsRecurringOrderCopies) SetData(v OrderSubscriptionDataRelationshipsRecurringOrderCopiesData) {
	o.Data = &v
}

func (o OrderSubscriptionDataRelationshipsRecurringOrderCopies) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSubscriptionDataRelationshipsRecurringOrderCopies) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableOrderSubscriptionDataRelationshipsRecurringOrderCopies struct {
	value *OrderSubscriptionDataRelationshipsRecurringOrderCopies
	isSet bool
}

func (v NullableOrderSubscriptionDataRelationshipsRecurringOrderCopies) Get() *OrderSubscriptionDataRelationshipsRecurringOrderCopies {
	return v.value
}

func (v *NullableOrderSubscriptionDataRelationshipsRecurringOrderCopies) Set(val *OrderSubscriptionDataRelationshipsRecurringOrderCopies) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSubscriptionDataRelationshipsRecurringOrderCopies) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSubscriptionDataRelationshipsRecurringOrderCopies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSubscriptionDataRelationshipsRecurringOrderCopies(val *OrderSubscriptionDataRelationshipsRecurringOrderCopies) *NullableOrderSubscriptionDataRelationshipsRecurringOrderCopies {
	return &NullableOrderSubscriptionDataRelationshipsRecurringOrderCopies{value: val, isSet: true}
}

func (v NullableOrderSubscriptionDataRelationshipsRecurringOrderCopies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSubscriptionDataRelationshipsRecurringOrderCopies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
