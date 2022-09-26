/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CustomerAddressResponseDataRelationships struct for CustomerAddressResponseDataRelationships
type CustomerAddressResponseDataRelationships struct {
	Customer *CouponRecipientResponseDataRelationshipsCustomer `json:"customer,omitempty"`
	Address  *CustomerAddressResponseDataRelationshipsAddress  `json:"address,omitempty"`
	Events   *CustomerAddressResponseDataRelationshipsEvents   `json:"events,omitempty"`
}

// NewCustomerAddressResponseDataRelationships instantiates a new CustomerAddressResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerAddressResponseDataRelationships() *CustomerAddressResponseDataRelationships {
	this := CustomerAddressResponseDataRelationships{}
	return &this
}

// NewCustomerAddressResponseDataRelationshipsWithDefaults instantiates a new CustomerAddressResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerAddressResponseDataRelationshipsWithDefaults() *CustomerAddressResponseDataRelationships {
	this := CustomerAddressResponseDataRelationships{}
	return &this
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *CustomerAddressResponseDataRelationships) GetCustomer() CouponRecipientResponseDataRelationshipsCustomer {
	if o == nil || o.Customer == nil {
		var ret CouponRecipientResponseDataRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressResponseDataRelationships) GetCustomerOk() (*CouponRecipientResponseDataRelationshipsCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *CustomerAddressResponseDataRelationships) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given CouponRecipientResponseDataRelationshipsCustomer and assigns it to the Customer field.
func (o *CustomerAddressResponseDataRelationships) SetCustomer(v CouponRecipientResponseDataRelationshipsCustomer) {
	o.Customer = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *CustomerAddressResponseDataRelationships) GetAddress() CustomerAddressResponseDataRelationshipsAddress {
	if o == nil || o.Address == nil {
		var ret CustomerAddressResponseDataRelationshipsAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressResponseDataRelationships) GetAddressOk() (*CustomerAddressResponseDataRelationshipsAddress, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *CustomerAddressResponseDataRelationships) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given CustomerAddressResponseDataRelationshipsAddress and assigns it to the Address field.
func (o *CustomerAddressResponseDataRelationships) SetAddress(v CustomerAddressResponseDataRelationshipsAddress) {
	o.Address = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *CustomerAddressResponseDataRelationships) GetEvents() CustomerAddressResponseDataRelationshipsEvents {
	if o == nil || o.Events == nil {
		var ret CustomerAddressResponseDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressResponseDataRelationships) GetEventsOk() (*CustomerAddressResponseDataRelationshipsEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *CustomerAddressResponseDataRelationships) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given CustomerAddressResponseDataRelationshipsEvents and assigns it to the Events field.
func (o *CustomerAddressResponseDataRelationships) SetEvents(v CustomerAddressResponseDataRelationshipsEvents) {
	o.Events = &v
}

func (o CustomerAddressResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerAddressResponseDataRelationships struct {
	value *CustomerAddressResponseDataRelationships
	isSet bool
}

func (v NullableCustomerAddressResponseDataRelationships) Get() *CustomerAddressResponseDataRelationships {
	return v.value
}

func (v *NullableCustomerAddressResponseDataRelationships) Set(val *CustomerAddressResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerAddressResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerAddressResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerAddressResponseDataRelationships(val *CustomerAddressResponseDataRelationships) *NullableCustomerAddressResponseDataRelationships {
	return &NullableCustomerAddressResponseDataRelationships{value: val, isSet: true}
}

func (v NullableCustomerAddressResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerAddressResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
