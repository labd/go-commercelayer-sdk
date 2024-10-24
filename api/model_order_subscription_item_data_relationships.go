/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the OrderSubscriptionItemDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSubscriptionItemDataRelationships{}

// OrderSubscriptionItemDataRelationships struct for OrderSubscriptionItemDataRelationships
type OrderSubscriptionItemDataRelationships struct {
	OrderSubscription *CustomerDataRelationshipsOrderSubscriptions `json:"order_subscription,omitempty"`
	Item              *OrderSubscriptionItemDataRelationshipsItem  `json:"item,omitempty"`
}

// NewOrderSubscriptionItemDataRelationships instantiates a new OrderSubscriptionItemDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSubscriptionItemDataRelationships() *OrderSubscriptionItemDataRelationships {
	this := OrderSubscriptionItemDataRelationships{}
	return &this
}

// NewOrderSubscriptionItemDataRelationshipsWithDefaults instantiates a new OrderSubscriptionItemDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSubscriptionItemDataRelationshipsWithDefaults() *OrderSubscriptionItemDataRelationships {
	this := OrderSubscriptionItemDataRelationships{}
	return &this
}

// GetOrderSubscription returns the OrderSubscription field value if set, zero value otherwise.
func (o *OrderSubscriptionItemDataRelationships) GetOrderSubscription() CustomerDataRelationshipsOrderSubscriptions {
	if o == nil || IsNil(o.OrderSubscription) {
		var ret CustomerDataRelationshipsOrderSubscriptions
		return ret
	}
	return *o.OrderSubscription
}

// GetOrderSubscriptionOk returns a tuple with the OrderSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionItemDataRelationships) GetOrderSubscriptionOk() (*CustomerDataRelationshipsOrderSubscriptions, bool) {
	if o == nil || IsNil(o.OrderSubscription) {
		return nil, false
	}
	return o.OrderSubscription, true
}

// HasOrderSubscription returns a boolean if a field has been set.
func (o *OrderSubscriptionItemDataRelationships) HasOrderSubscription() bool {
	if o != nil && !IsNil(o.OrderSubscription) {
		return true
	}

	return false
}

// SetOrderSubscription gets a reference to the given CustomerDataRelationshipsOrderSubscriptions and assigns it to the OrderSubscription field.
func (o *OrderSubscriptionItemDataRelationships) SetOrderSubscription(v CustomerDataRelationshipsOrderSubscriptions) {
	o.OrderSubscription = &v
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *OrderSubscriptionItemDataRelationships) GetItem() OrderSubscriptionItemDataRelationshipsItem {
	if o == nil || IsNil(o.Item) {
		var ret OrderSubscriptionItemDataRelationshipsItem
		return ret
	}
	return *o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionItemDataRelationships) GetItemOk() (*OrderSubscriptionItemDataRelationshipsItem, bool) {
	if o == nil || IsNil(o.Item) {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *OrderSubscriptionItemDataRelationships) HasItem() bool {
	if o != nil && !IsNil(o.Item) {
		return true
	}

	return false
}

// SetItem gets a reference to the given OrderSubscriptionItemDataRelationshipsItem and assigns it to the Item field.
func (o *OrderSubscriptionItemDataRelationships) SetItem(v OrderSubscriptionItemDataRelationshipsItem) {
	o.Item = &v
}

func (o OrderSubscriptionItemDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSubscriptionItemDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderSubscription) {
		toSerialize["order_subscription"] = o.OrderSubscription
	}
	if !IsNil(o.Item) {
		toSerialize["item"] = o.Item
	}
	return toSerialize, nil
}

type NullableOrderSubscriptionItemDataRelationships struct {
	value *OrderSubscriptionItemDataRelationships
	isSet bool
}

func (v NullableOrderSubscriptionItemDataRelationships) Get() *OrderSubscriptionItemDataRelationships {
	return v.value
}

func (v *NullableOrderSubscriptionItemDataRelationships) Set(val *OrderSubscriptionItemDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSubscriptionItemDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSubscriptionItemDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSubscriptionItemDataRelationships(val *OrderSubscriptionItemDataRelationships) *NullableOrderSubscriptionItemDataRelationships {
	return &NullableOrderSubscriptionItemDataRelationships{value: val, isSet: true}
}

func (v NullableOrderSubscriptionItemDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSubscriptionItemDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}