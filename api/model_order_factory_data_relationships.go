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

// checks if the OrderFactoryDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderFactoryDataRelationships{}

// OrderFactoryDataRelationships struct for OrderFactoryDataRelationships
type OrderFactoryDataRelationships struct {
	SourceOrder *AdyenPaymentDataRelationshipsOrder `json:"source_order,omitempty"`
	TargetOrder *AdyenPaymentDataRelationshipsOrder `json:"target_order,omitempty"`
	Events      *AddressDataRelationshipsEvents     `json:"events,omitempty"`
}

// NewOrderFactoryDataRelationships instantiates a new OrderFactoryDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderFactoryDataRelationships() *OrderFactoryDataRelationships {
	this := OrderFactoryDataRelationships{}
	return &this
}

// NewOrderFactoryDataRelationshipsWithDefaults instantiates a new OrderFactoryDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderFactoryDataRelationshipsWithDefaults() *OrderFactoryDataRelationships {
	this := OrderFactoryDataRelationships{}
	return &this
}

// GetSourceOrder returns the SourceOrder field value if set, zero value otherwise.
func (o *OrderFactoryDataRelationships) GetSourceOrder() AdyenPaymentDataRelationshipsOrder {
	if o == nil || IsNil(o.SourceOrder) {
		var ret AdyenPaymentDataRelationshipsOrder
		return ret
	}
	return *o.SourceOrder
}

// GetSourceOrderOk returns a tuple with the SourceOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFactoryDataRelationships) GetSourceOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool) {
	if o == nil || IsNil(o.SourceOrder) {
		return nil, false
	}
	return o.SourceOrder, true
}

// HasSourceOrder returns a boolean if a field has been set.
func (o *OrderFactoryDataRelationships) HasSourceOrder() bool {
	if o != nil && !IsNil(o.SourceOrder) {
		return true
	}

	return false
}

// SetSourceOrder gets a reference to the given AdyenPaymentDataRelationshipsOrder and assigns it to the SourceOrder field.
func (o *OrderFactoryDataRelationships) SetSourceOrder(v AdyenPaymentDataRelationshipsOrder) {
	o.SourceOrder = &v
}

// GetTargetOrder returns the TargetOrder field value if set, zero value otherwise.
func (o *OrderFactoryDataRelationships) GetTargetOrder() AdyenPaymentDataRelationshipsOrder {
	if o == nil || IsNil(o.TargetOrder) {
		var ret AdyenPaymentDataRelationshipsOrder
		return ret
	}
	return *o.TargetOrder
}

// GetTargetOrderOk returns a tuple with the TargetOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFactoryDataRelationships) GetTargetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool) {
	if o == nil || IsNil(o.TargetOrder) {
		return nil, false
	}
	return o.TargetOrder, true
}

// HasTargetOrder returns a boolean if a field has been set.
func (o *OrderFactoryDataRelationships) HasTargetOrder() bool {
	if o != nil && !IsNil(o.TargetOrder) {
		return true
	}

	return false
}

// SetTargetOrder gets a reference to the given AdyenPaymentDataRelationshipsOrder and assigns it to the TargetOrder field.
func (o *OrderFactoryDataRelationships) SetTargetOrder(v AdyenPaymentDataRelationshipsOrder) {
	o.TargetOrder = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *OrderFactoryDataRelationships) GetEvents() AddressDataRelationshipsEvents {
	if o == nil || IsNil(o.Events) {
		var ret AddressDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFactoryDataRelationships) GetEventsOk() (*AddressDataRelationshipsEvents, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *OrderFactoryDataRelationships) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given AddressDataRelationshipsEvents and assigns it to the Events field.
func (o *OrderFactoryDataRelationships) SetEvents(v AddressDataRelationshipsEvents) {
	o.Events = &v
}

func (o OrderFactoryDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderFactoryDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SourceOrder) {
		toSerialize["source_order"] = o.SourceOrder
	}
	if !IsNil(o.TargetOrder) {
		toSerialize["target_order"] = o.TargetOrder
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	return toSerialize, nil
}

type NullableOrderFactoryDataRelationships struct {
	value *OrderFactoryDataRelationships
	isSet bool
}

func (v NullableOrderFactoryDataRelationships) Get() *OrderFactoryDataRelationships {
	return v.value
}

func (v *NullableOrderFactoryDataRelationships) Set(val *OrderFactoryDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderFactoryDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderFactoryDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderFactoryDataRelationships(val *OrderFactoryDataRelationships) *NullableOrderFactoryDataRelationships {
	return &NullableOrderFactoryDataRelationships{value: val, isSet: true}
}

func (v NullableOrderFactoryDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderFactoryDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
