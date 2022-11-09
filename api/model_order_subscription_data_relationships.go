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

// OrderSubscriptionDataRelationships struct for OrderSubscriptionDataRelationships
type OrderSubscriptionDataRelationships struct {
	Market      *AvalaraAccountDataRelationshipsMarkets        `json:"market,omitempty"`
	SourceOrder *AdyenPaymentDataRelationshipsOrder            `json:"source_order,omitempty"`
	Customer    *CouponRecipientDataRelationshipsCustomer      `json:"customer,omitempty"`
	OrderCopies *OrderSubscriptionDataRelationshipsOrderCopies `json:"order_copies,omitempty"`
	Orders      *AdyenPaymentDataRelationshipsOrder            `json:"orders,omitempty"`
	Events      *CustomerAddressDataRelationshipsEvents        `json:"events,omitempty"`
}

// NewOrderSubscriptionDataRelationships instantiates a new OrderSubscriptionDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSubscriptionDataRelationships() *OrderSubscriptionDataRelationships {
	this := OrderSubscriptionDataRelationships{}
	return &this
}

// NewOrderSubscriptionDataRelationshipsWithDefaults instantiates a new OrderSubscriptionDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSubscriptionDataRelationshipsWithDefaults() *OrderSubscriptionDataRelationships {
	this := OrderSubscriptionDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *OrderSubscriptionDataRelationships) GetMarket() AvalaraAccountDataRelationshipsMarkets {
	if o == nil || o.Market == nil {
		var ret AvalaraAccountDataRelationshipsMarkets
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionDataRelationships) GetMarketOk() (*AvalaraAccountDataRelationshipsMarkets, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *OrderSubscriptionDataRelationships) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given AvalaraAccountDataRelationshipsMarkets and assigns it to the Market field.
func (o *OrderSubscriptionDataRelationships) SetMarket(v AvalaraAccountDataRelationshipsMarkets) {
	o.Market = &v
}

// GetSourceOrder returns the SourceOrder field value if set, zero value otherwise.
func (o *OrderSubscriptionDataRelationships) GetSourceOrder() AdyenPaymentDataRelationshipsOrder {
	if o == nil || o.SourceOrder == nil {
		var ret AdyenPaymentDataRelationshipsOrder
		return ret
	}
	return *o.SourceOrder
}

// GetSourceOrderOk returns a tuple with the SourceOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionDataRelationships) GetSourceOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool) {
	if o == nil || o.SourceOrder == nil {
		return nil, false
	}
	return o.SourceOrder, true
}

// HasSourceOrder returns a boolean if a field has been set.
func (o *OrderSubscriptionDataRelationships) HasSourceOrder() bool {
	if o != nil && o.SourceOrder != nil {
		return true
	}

	return false
}

// SetSourceOrder gets a reference to the given AdyenPaymentDataRelationshipsOrder and assigns it to the SourceOrder field.
func (o *OrderSubscriptionDataRelationships) SetSourceOrder(v AdyenPaymentDataRelationshipsOrder) {
	o.SourceOrder = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *OrderSubscriptionDataRelationships) GetCustomer() CouponRecipientDataRelationshipsCustomer {
	if o == nil || o.Customer == nil {
		var ret CouponRecipientDataRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionDataRelationships) GetCustomerOk() (*CouponRecipientDataRelationshipsCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *OrderSubscriptionDataRelationships) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given CouponRecipientDataRelationshipsCustomer and assigns it to the Customer field.
func (o *OrderSubscriptionDataRelationships) SetCustomer(v CouponRecipientDataRelationshipsCustomer) {
	o.Customer = &v
}

// GetOrderCopies returns the OrderCopies field value if set, zero value otherwise.
func (o *OrderSubscriptionDataRelationships) GetOrderCopies() OrderSubscriptionDataRelationshipsOrderCopies {
	if o == nil || o.OrderCopies == nil {
		var ret OrderSubscriptionDataRelationshipsOrderCopies
		return ret
	}
	return *o.OrderCopies
}

// GetOrderCopiesOk returns a tuple with the OrderCopies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionDataRelationships) GetOrderCopiesOk() (*OrderSubscriptionDataRelationshipsOrderCopies, bool) {
	if o == nil || o.OrderCopies == nil {
		return nil, false
	}
	return o.OrderCopies, true
}

// HasOrderCopies returns a boolean if a field has been set.
func (o *OrderSubscriptionDataRelationships) HasOrderCopies() bool {
	if o != nil && o.OrderCopies != nil {
		return true
	}

	return false
}

// SetOrderCopies gets a reference to the given OrderSubscriptionDataRelationshipsOrderCopies and assigns it to the OrderCopies field.
func (o *OrderSubscriptionDataRelationships) SetOrderCopies(v OrderSubscriptionDataRelationshipsOrderCopies) {
	o.OrderCopies = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *OrderSubscriptionDataRelationships) GetOrders() AdyenPaymentDataRelationshipsOrder {
	if o == nil || o.Orders == nil {
		var ret AdyenPaymentDataRelationshipsOrder
		return ret
	}
	return *o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionDataRelationships) GetOrdersOk() (*AdyenPaymentDataRelationshipsOrder, bool) {
	if o == nil || o.Orders == nil {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *OrderSubscriptionDataRelationships) HasOrders() bool {
	if o != nil && o.Orders != nil {
		return true
	}

	return false
}

// SetOrders gets a reference to the given AdyenPaymentDataRelationshipsOrder and assigns it to the Orders field.
func (o *OrderSubscriptionDataRelationships) SetOrders(v AdyenPaymentDataRelationshipsOrder) {
	o.Orders = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *OrderSubscriptionDataRelationships) GetEvents() CustomerAddressDataRelationshipsEvents {
	if o == nil || o.Events == nil {
		var ret CustomerAddressDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionDataRelationships) GetEventsOk() (*CustomerAddressDataRelationshipsEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *OrderSubscriptionDataRelationships) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given CustomerAddressDataRelationshipsEvents and assigns it to the Events field.
func (o *OrderSubscriptionDataRelationships) SetEvents(v CustomerAddressDataRelationshipsEvents) {
	o.Events = &v
}

func (o OrderSubscriptionDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Market != nil {
		toSerialize["market"] = o.Market
	}
	if o.SourceOrder != nil {
		toSerialize["source_order"] = o.SourceOrder
	}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.OrderCopies != nil {
		toSerialize["order_copies"] = o.OrderCopies
	}
	if o.Orders != nil {
		toSerialize["orders"] = o.Orders
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableOrderSubscriptionDataRelationships struct {
	value *OrderSubscriptionDataRelationships
	isSet bool
}

func (v NullableOrderSubscriptionDataRelationships) Get() *OrderSubscriptionDataRelationships {
	return v.value
}

func (v *NullableOrderSubscriptionDataRelationships) Set(val *OrderSubscriptionDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSubscriptionDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSubscriptionDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSubscriptionDataRelationships(val *OrderSubscriptionDataRelationships) *NullableOrderSubscriptionDataRelationships {
	return &NullableOrderSubscriptionDataRelationships{value: val, isSet: true}
}

func (v NullableOrderSubscriptionDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSubscriptionDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
