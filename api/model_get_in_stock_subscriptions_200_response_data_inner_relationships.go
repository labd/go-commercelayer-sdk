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

// GETInStockSubscriptions200ResponseDataInnerRelationships struct for GETInStockSubscriptions200ResponseDataInnerRelationships
type GETInStockSubscriptions200ResponseDataInnerRelationships struct {
	Market   *GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket `json:"market,omitempty"`
	Customer *GETCouponRecipients200ResponseDataInnerRelationshipsCustomer         `json:"customer,omitempty"`
	Sku      *GETInStockSubscriptions200ResponseDataInnerRelationshipsSku          `json:"sku,omitempty"`
	Events   *GETCustomerAddresses200ResponseDataInnerRelationshipsEvents          `json:"events,omitempty"`
}

// NewGETInStockSubscriptions200ResponseDataInnerRelationships instantiates a new GETInStockSubscriptions200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETInStockSubscriptions200ResponseDataInnerRelationships() *GETInStockSubscriptions200ResponseDataInnerRelationships {
	this := GETInStockSubscriptions200ResponseDataInnerRelationships{}
	return &this
}

// NewGETInStockSubscriptions200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETInStockSubscriptions200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETInStockSubscriptions200ResponseDataInnerRelationshipsWithDefaults() *GETInStockSubscriptions200ResponseDataInnerRelationships {
	this := GETInStockSubscriptions200ResponseDataInnerRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationships) GetMarket() GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket {
	if o == nil || o.Market == nil {
		var ret GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationships) GetMarketOk() (*GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationships) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket and assigns it to the Market field.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationships) SetMarket(v GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket) {
	o.Market = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationships) GetCustomer() GETCouponRecipients200ResponseDataInnerRelationshipsCustomer {
	if o == nil || o.Customer == nil {
		var ret GETCouponRecipients200ResponseDataInnerRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationships) GetCustomerOk() (*GETCouponRecipients200ResponseDataInnerRelationshipsCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationships) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given GETCouponRecipients200ResponseDataInnerRelationshipsCustomer and assigns it to the Customer field.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationships) SetCustomer(v GETCouponRecipients200ResponseDataInnerRelationshipsCustomer) {
	o.Customer = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationships) GetSku() GETInStockSubscriptions200ResponseDataInnerRelationshipsSku {
	if o == nil || o.Sku == nil {
		var ret GETInStockSubscriptions200ResponseDataInnerRelationshipsSku
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationships) GetSkuOk() (*GETInStockSubscriptions200ResponseDataInnerRelationshipsSku, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationships) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given GETInStockSubscriptions200ResponseDataInnerRelationshipsSku and assigns it to the Sku field.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationships) SetSku(v GETInStockSubscriptions200ResponseDataInnerRelationshipsSku) {
	o.Sku = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationships) GetEvents() GETCustomerAddresses200ResponseDataInnerRelationshipsEvents {
	if o == nil || o.Events == nil {
		var ret GETCustomerAddresses200ResponseDataInnerRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationships) GetEventsOk() (*GETCustomerAddresses200ResponseDataInnerRelationshipsEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationships) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given GETCustomerAddresses200ResponseDataInnerRelationshipsEvents and assigns it to the Events field.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationships) SetEvents(v GETCustomerAddresses200ResponseDataInnerRelationshipsEvents) {
	o.Events = &v
}

func (o GETInStockSubscriptions200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Market != nil {
		toSerialize["market"] = o.Market
	}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.Sku != nil {
		toSerialize["sku"] = o.Sku
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableGETInStockSubscriptions200ResponseDataInnerRelationships struct {
	value *GETInStockSubscriptions200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETInStockSubscriptions200ResponseDataInnerRelationships) Get() *GETInStockSubscriptions200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETInStockSubscriptions200ResponseDataInnerRelationships) Set(val *GETInStockSubscriptions200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETInStockSubscriptions200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETInStockSubscriptions200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETInStockSubscriptions200ResponseDataInnerRelationships(val *GETInStockSubscriptions200ResponseDataInnerRelationships) *NullableGETInStockSubscriptions200ResponseDataInnerRelationships {
	return &NullableGETInStockSubscriptions200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETInStockSubscriptions200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETInStockSubscriptions200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}