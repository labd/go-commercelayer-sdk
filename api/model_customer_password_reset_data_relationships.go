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

// checks if the CustomerPasswordResetDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerPasswordResetDataRelationships{}

// CustomerPasswordResetDataRelationships struct for CustomerPasswordResetDataRelationships
type CustomerPasswordResetDataRelationships struct {
	Customer *CouponRecipientDataRelationshipsCustomer `json:"customer,omitempty"`
	Events   *AddressDataRelationshipsEvents           `json:"events,omitempty"`
}

// NewCustomerPasswordResetDataRelationships instantiates a new CustomerPasswordResetDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPasswordResetDataRelationships() *CustomerPasswordResetDataRelationships {
	this := CustomerPasswordResetDataRelationships{}
	return &this
}

// NewCustomerPasswordResetDataRelationshipsWithDefaults instantiates a new CustomerPasswordResetDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPasswordResetDataRelationshipsWithDefaults() *CustomerPasswordResetDataRelationships {
	this := CustomerPasswordResetDataRelationships{}
	return &this
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *CustomerPasswordResetDataRelationships) GetCustomer() CouponRecipientDataRelationshipsCustomer {
	if o == nil || IsNil(o.Customer) {
		var ret CouponRecipientDataRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPasswordResetDataRelationships) GetCustomerOk() (*CouponRecipientDataRelationshipsCustomer, bool) {
	if o == nil || IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *CustomerPasswordResetDataRelationships) HasCustomer() bool {
	if o != nil && !IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given CouponRecipientDataRelationshipsCustomer and assigns it to the Customer field.
func (o *CustomerPasswordResetDataRelationships) SetCustomer(v CouponRecipientDataRelationshipsCustomer) {
	o.Customer = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *CustomerPasswordResetDataRelationships) GetEvents() AddressDataRelationshipsEvents {
	if o == nil || IsNil(o.Events) {
		var ret AddressDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPasswordResetDataRelationships) GetEventsOk() (*AddressDataRelationshipsEvents, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *CustomerPasswordResetDataRelationships) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given AddressDataRelationshipsEvents and assigns it to the Events field.
func (o *CustomerPasswordResetDataRelationships) SetEvents(v AddressDataRelationshipsEvents) {
	o.Events = &v
}

func (o CustomerPasswordResetDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerPasswordResetDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Customer) {
		toSerialize["customer"] = o.Customer
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	return toSerialize, nil
}

type NullableCustomerPasswordResetDataRelationships struct {
	value *CustomerPasswordResetDataRelationships
	isSet bool
}

func (v NullableCustomerPasswordResetDataRelationships) Get() *CustomerPasswordResetDataRelationships {
	return v.value
}

func (v *NullableCustomerPasswordResetDataRelationships) Set(val *CustomerPasswordResetDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPasswordResetDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPasswordResetDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPasswordResetDataRelationships(val *CustomerPasswordResetDataRelationships) *NullableCustomerPasswordResetDataRelationships {
	return &NullableCustomerPasswordResetDataRelationships{value: val, isSet: true}
}

func (v NullableCustomerPasswordResetDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPasswordResetDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
