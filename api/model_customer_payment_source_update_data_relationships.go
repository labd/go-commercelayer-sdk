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

// CustomerPaymentSourceUpdateDataRelationships struct for CustomerPaymentSourceUpdateDataRelationships
type CustomerPaymentSourceUpdateDataRelationships struct {
	Customer      *CouponRecipientCreateDataRelationshipsCustomer            `json:"customer,omitempty"`
	PaymentSource *CustomerPaymentSourceCreateDataRelationshipsPaymentSource `json:"payment_source,omitempty"`
}

// NewCustomerPaymentSourceUpdateDataRelationships instantiates a new CustomerPaymentSourceUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPaymentSourceUpdateDataRelationships() *CustomerPaymentSourceUpdateDataRelationships {
	this := CustomerPaymentSourceUpdateDataRelationships{}
	return &this
}

// NewCustomerPaymentSourceUpdateDataRelationshipsWithDefaults instantiates a new CustomerPaymentSourceUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPaymentSourceUpdateDataRelationshipsWithDefaults() *CustomerPaymentSourceUpdateDataRelationships {
	this := CustomerPaymentSourceUpdateDataRelationships{}
	return &this
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *CustomerPaymentSourceUpdateDataRelationships) GetCustomer() CouponRecipientCreateDataRelationshipsCustomer {
	if o == nil || o.Customer == nil {
		var ret CouponRecipientCreateDataRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceUpdateDataRelationships) GetCustomerOk() (*CouponRecipientCreateDataRelationshipsCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *CustomerPaymentSourceUpdateDataRelationships) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given CouponRecipientCreateDataRelationshipsCustomer and assigns it to the Customer field.
func (o *CustomerPaymentSourceUpdateDataRelationships) SetCustomer(v CouponRecipientCreateDataRelationshipsCustomer) {
	o.Customer = &v
}

// GetPaymentSource returns the PaymentSource field value if set, zero value otherwise.
func (o *CustomerPaymentSourceUpdateDataRelationships) GetPaymentSource() CustomerPaymentSourceCreateDataRelationshipsPaymentSource {
	if o == nil || o.PaymentSource == nil {
		var ret CustomerPaymentSourceCreateDataRelationshipsPaymentSource
		return ret
	}
	return *o.PaymentSource
}

// GetPaymentSourceOk returns a tuple with the PaymentSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceUpdateDataRelationships) GetPaymentSourceOk() (*CustomerPaymentSourceCreateDataRelationshipsPaymentSource, bool) {
	if o == nil || o.PaymentSource == nil {
		return nil, false
	}
	return o.PaymentSource, true
}

// HasPaymentSource returns a boolean if a field has been set.
func (o *CustomerPaymentSourceUpdateDataRelationships) HasPaymentSource() bool {
	if o != nil && o.PaymentSource != nil {
		return true
	}

	return false
}

// SetPaymentSource gets a reference to the given CustomerPaymentSourceCreateDataRelationshipsPaymentSource and assigns it to the PaymentSource field.
func (o *CustomerPaymentSourceUpdateDataRelationships) SetPaymentSource(v CustomerPaymentSourceCreateDataRelationshipsPaymentSource) {
	o.PaymentSource = &v
}

func (o CustomerPaymentSourceUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.PaymentSource != nil {
		toSerialize["payment_source"] = o.PaymentSource
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerPaymentSourceUpdateDataRelationships struct {
	value *CustomerPaymentSourceUpdateDataRelationships
	isSet bool
}

func (v NullableCustomerPaymentSourceUpdateDataRelationships) Get() *CustomerPaymentSourceUpdateDataRelationships {
	return v.value
}

func (v *NullableCustomerPaymentSourceUpdateDataRelationships) Set(val *CustomerPaymentSourceUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPaymentSourceUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPaymentSourceUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPaymentSourceUpdateDataRelationships(val *CustomerPaymentSourceUpdateDataRelationships) *NullableCustomerPaymentSourceUpdateDataRelationships {
	return &NullableCustomerPaymentSourceUpdateDataRelationships{value: val, isSet: true}
}

func (v NullableCustomerPaymentSourceUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPaymentSourceUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
