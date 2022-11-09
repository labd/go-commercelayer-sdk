/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.2
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaymentMethodUpdateDataRelationships struct for PaymentMethodUpdateDataRelationships
type PaymentMethodUpdateDataRelationships struct {
	Market         *AvalaraAccountDataRelationshipsMarkets      `json:"market,omitempty"`
	PaymentGateway *AdyenPaymentDataRelationshipsPaymentGateway `json:"payment_gateway,omitempty"`
}

// NewPaymentMethodUpdateDataRelationships instantiates a new PaymentMethodUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodUpdateDataRelationships() *PaymentMethodUpdateDataRelationships {
	this := PaymentMethodUpdateDataRelationships{}
	return &this
}

// NewPaymentMethodUpdateDataRelationshipsWithDefaults instantiates a new PaymentMethodUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodUpdateDataRelationshipsWithDefaults() *PaymentMethodUpdateDataRelationships {
	this := PaymentMethodUpdateDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *PaymentMethodUpdateDataRelationships) GetMarket() AvalaraAccountDataRelationshipsMarkets {
	if o == nil || o.Market == nil {
		var ret AvalaraAccountDataRelationshipsMarkets
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodUpdateDataRelationships) GetMarketOk() (*AvalaraAccountDataRelationshipsMarkets, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *PaymentMethodUpdateDataRelationships) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given AvalaraAccountDataRelationshipsMarkets and assigns it to the Market field.
func (o *PaymentMethodUpdateDataRelationships) SetMarket(v AvalaraAccountDataRelationshipsMarkets) {
	o.Market = &v
}

// GetPaymentGateway returns the PaymentGateway field value if set, zero value otherwise.
func (o *PaymentMethodUpdateDataRelationships) GetPaymentGateway() AdyenPaymentDataRelationshipsPaymentGateway {
	if o == nil || o.PaymentGateway == nil {
		var ret AdyenPaymentDataRelationshipsPaymentGateway
		return ret
	}
	return *o.PaymentGateway
}

// GetPaymentGatewayOk returns a tuple with the PaymentGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodUpdateDataRelationships) GetPaymentGatewayOk() (*AdyenPaymentDataRelationshipsPaymentGateway, bool) {
	if o == nil || o.PaymentGateway == nil {
		return nil, false
	}
	return o.PaymentGateway, true
}

// HasPaymentGateway returns a boolean if a field has been set.
func (o *PaymentMethodUpdateDataRelationships) HasPaymentGateway() bool {
	if o != nil && o.PaymentGateway != nil {
		return true
	}

	return false
}

// SetPaymentGateway gets a reference to the given AdyenPaymentDataRelationshipsPaymentGateway and assigns it to the PaymentGateway field.
func (o *PaymentMethodUpdateDataRelationships) SetPaymentGateway(v AdyenPaymentDataRelationshipsPaymentGateway) {
	o.PaymentGateway = &v
}

func (o PaymentMethodUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Market != nil {
		toSerialize["market"] = o.Market
	}
	if o.PaymentGateway != nil {
		toSerialize["payment_gateway"] = o.PaymentGateway
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentMethodUpdateDataRelationships struct {
	value *PaymentMethodUpdateDataRelationships
	isSet bool
}

func (v NullablePaymentMethodUpdateDataRelationships) Get() *PaymentMethodUpdateDataRelationships {
	return v.value
}

func (v *NullablePaymentMethodUpdateDataRelationships) Set(val *PaymentMethodUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodUpdateDataRelationships(val *PaymentMethodUpdateDataRelationships) *NullablePaymentMethodUpdateDataRelationships {
	return &NullablePaymentMethodUpdateDataRelationships{value: val, isSet: true}
}

func (v NullablePaymentMethodUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
