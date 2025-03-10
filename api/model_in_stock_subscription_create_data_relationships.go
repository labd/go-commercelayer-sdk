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

// checks if the InStockSubscriptionCreateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InStockSubscriptionCreateDataRelationships{}

// InStockSubscriptionCreateDataRelationships struct for InStockSubscriptionCreateDataRelationships
type InStockSubscriptionCreateDataRelationships struct {
	Market   BundleCreateDataRelationshipsMarket            `json:"market"`
	Customer CouponRecipientCreateDataRelationshipsCustomer `json:"customer"`
	Sku      InStockSubscriptionCreateDataRelationshipsSku  `json:"sku"`
}

// NewInStockSubscriptionCreateDataRelationships instantiates a new InStockSubscriptionCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInStockSubscriptionCreateDataRelationships(market BundleCreateDataRelationshipsMarket, customer CouponRecipientCreateDataRelationshipsCustomer, sku InStockSubscriptionCreateDataRelationshipsSku) *InStockSubscriptionCreateDataRelationships {
	this := InStockSubscriptionCreateDataRelationships{}
	this.Market = market
	this.Customer = customer
	this.Sku = sku
	return &this
}

// NewInStockSubscriptionCreateDataRelationshipsWithDefaults instantiates a new InStockSubscriptionCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInStockSubscriptionCreateDataRelationshipsWithDefaults() *InStockSubscriptionCreateDataRelationships {
	this := InStockSubscriptionCreateDataRelationships{}
	return &this
}

// GetMarket returns the Market field value
func (o *InStockSubscriptionCreateDataRelationships) GetMarket() BundleCreateDataRelationshipsMarket {
	if o == nil {
		var ret BundleCreateDataRelationshipsMarket
		return ret
	}

	return o.Market
}

// GetMarketOk returns a tuple with the Market field value
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionCreateDataRelationships) GetMarketOk() (*BundleCreateDataRelationshipsMarket, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Market, true
}

// SetMarket sets field value
func (o *InStockSubscriptionCreateDataRelationships) SetMarket(v BundleCreateDataRelationshipsMarket) {
	o.Market = v
}

// GetCustomer returns the Customer field value
func (o *InStockSubscriptionCreateDataRelationships) GetCustomer() CouponRecipientCreateDataRelationshipsCustomer {
	if o == nil {
		var ret CouponRecipientCreateDataRelationshipsCustomer
		return ret
	}

	return o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionCreateDataRelationships) GetCustomerOk() (*CouponRecipientCreateDataRelationshipsCustomer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Customer, true
}

// SetCustomer sets field value
func (o *InStockSubscriptionCreateDataRelationships) SetCustomer(v CouponRecipientCreateDataRelationshipsCustomer) {
	o.Customer = v
}

// GetSku returns the Sku field value
func (o *InStockSubscriptionCreateDataRelationships) GetSku() InStockSubscriptionCreateDataRelationshipsSku {
	if o == nil {
		var ret InStockSubscriptionCreateDataRelationshipsSku
		return ret
	}

	return o.Sku
}

// GetSkuOk returns a tuple with the Sku field value
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionCreateDataRelationships) GetSkuOk() (*InStockSubscriptionCreateDataRelationshipsSku, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sku, true
}

// SetSku sets field value
func (o *InStockSubscriptionCreateDataRelationships) SetSku(v InStockSubscriptionCreateDataRelationshipsSku) {
	o.Sku = v
}

func (o InStockSubscriptionCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InStockSubscriptionCreateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["market"] = o.Market
	toSerialize["customer"] = o.Customer
	toSerialize["sku"] = o.Sku
	return toSerialize, nil
}

type NullableInStockSubscriptionCreateDataRelationships struct {
	value *InStockSubscriptionCreateDataRelationships
	isSet bool
}

func (v NullableInStockSubscriptionCreateDataRelationships) Get() *InStockSubscriptionCreateDataRelationships {
	return v.value
}

func (v *NullableInStockSubscriptionCreateDataRelationships) Set(val *InStockSubscriptionCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableInStockSubscriptionCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableInStockSubscriptionCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInStockSubscriptionCreateDataRelationships(val *InStockSubscriptionCreateDataRelationships) *NullableInStockSubscriptionCreateDataRelationships {
	return &NullableInStockSubscriptionCreateDataRelationships{value: val, isSet: true}
}

func (v NullableInStockSubscriptionCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInStockSubscriptionCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
