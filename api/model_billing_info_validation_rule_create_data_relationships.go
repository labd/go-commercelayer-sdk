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

// BillingInfoValidationRuleCreateDataRelationships struct for BillingInfoValidationRuleCreateDataRelationships
type BillingInfoValidationRuleCreateDataRelationships struct {
	Market AvalaraAccountDataRelationshipsMarkets `json:"market"`
}

// NewBillingInfoValidationRuleCreateDataRelationships instantiates a new BillingInfoValidationRuleCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInfoValidationRuleCreateDataRelationships(market AvalaraAccountDataRelationshipsMarkets) *BillingInfoValidationRuleCreateDataRelationships {
	this := BillingInfoValidationRuleCreateDataRelationships{}
	this.Market = market
	return &this
}

// NewBillingInfoValidationRuleCreateDataRelationshipsWithDefaults instantiates a new BillingInfoValidationRuleCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInfoValidationRuleCreateDataRelationshipsWithDefaults() *BillingInfoValidationRuleCreateDataRelationships {
	this := BillingInfoValidationRuleCreateDataRelationships{}
	return &this
}

// GetMarket returns the Market field value
func (o *BillingInfoValidationRuleCreateDataRelationships) GetMarket() AvalaraAccountDataRelationshipsMarkets {
	if o == nil {
		var ret AvalaraAccountDataRelationshipsMarkets
		return ret
	}

	return o.Market
}

// GetMarketOk returns a tuple with the Market field value
// and a boolean to check if the value has been set.
func (o *BillingInfoValidationRuleCreateDataRelationships) GetMarketOk() (*AvalaraAccountDataRelationshipsMarkets, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Market, true
}

// SetMarket sets field value
func (o *BillingInfoValidationRuleCreateDataRelationships) SetMarket(v AvalaraAccountDataRelationshipsMarkets) {
	o.Market = v
}

func (o BillingInfoValidationRuleCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["market"] = o.Market
	}
	return json.Marshal(toSerialize)
}

type NullableBillingInfoValidationRuleCreateDataRelationships struct {
	value *BillingInfoValidationRuleCreateDataRelationships
	isSet bool
}

func (v NullableBillingInfoValidationRuleCreateDataRelationships) Get() *BillingInfoValidationRuleCreateDataRelationships {
	return v.value
}

func (v *NullableBillingInfoValidationRuleCreateDataRelationships) Set(val *BillingInfoValidationRuleCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingInfoValidationRuleCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingInfoValidationRuleCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingInfoValidationRuleCreateDataRelationships(val *BillingInfoValidationRuleCreateDataRelationships) *NullableBillingInfoValidationRuleCreateDataRelationships {
	return &NullableBillingInfoValidationRuleCreateDataRelationships{value: val, isSet: true}
}

func (v NullableBillingInfoValidationRuleCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingInfoValidationRuleCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
