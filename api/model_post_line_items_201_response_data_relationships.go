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

// checks if the POSTLineItems201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTLineItems201ResponseDataRelationships{}

// POSTLineItems201ResponseDataRelationships struct for POSTLineItems201ResponseDataRelationships
type POSTLineItems201ResponseDataRelationships struct {
	Order             *POSTAdyenPayments201ResponseDataRelationshipsOrder                  `json:"order,omitempty"`
	Item              *POSTLineItems201ResponseDataRelationshipsItem                       `json:"item,omitempty"`
	Sku               *POSTInStockSubscriptions201ResponseDataRelationshipsSku             `json:"sku,omitempty"`
	Bundle            *POSTLineItems201ResponseDataRelationshipsBundle                     `json:"bundle,omitempty"`
	Adjustment        *POSTLineItems201ResponseDataRelationshipsAdjustment                 `json:"adjustment,omitempty"`
	GiftCard          *POSTLineItems201ResponseDataRelationshipsGiftCard                   `json:"gift_card,omitempty"`
	Shipment          *POSTLineItems201ResponseDataRelationshipsShipment                   `json:"shipment,omitempty"`
	PaymentMethod     *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentMethod `json:"payment_method,omitempty"`
	LineItemOptions   *POSTLineItems201ResponseDataRelationshipsLineItemOptions            `json:"line_item_options,omitempty"`
	ReturnLineItems   *POSTLineItems201ResponseDataRelationshipsReturnLineItems            `json:"return_line_items,omitempty"`
	StockReservations *POSTLineItems201ResponseDataRelationshipsStockReservations          `json:"stock_reservations,omitempty"`
	StockLineItems    *POSTLineItems201ResponseDataRelationshipsStockLineItems             `json:"stock_line_items,omitempty"`
	StockTransfers    *POSTLineItems201ResponseDataRelationshipsStockTransfers             `json:"stock_transfers,omitempty"`
	Notifications     *POSTLineItems201ResponseDataRelationshipsNotifications              `json:"notifications,omitempty"`
	Events            *POSTAddresses201ResponseDataRelationshipsEvents                     `json:"events,omitempty"`
	Tags              *POSTAddresses201ResponseDataRelationshipsTags                       `json:"tags,omitempty"`
}

// NewPOSTLineItems201ResponseDataRelationships instantiates a new POSTLineItems201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTLineItems201ResponseDataRelationships() *POSTLineItems201ResponseDataRelationships {
	this := POSTLineItems201ResponseDataRelationships{}
	return &this
}

// NewPOSTLineItems201ResponseDataRelationshipsWithDefaults instantiates a new POSTLineItems201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTLineItems201ResponseDataRelationshipsWithDefaults() *POSTLineItems201ResponseDataRelationships {
	this := POSTLineItems201ResponseDataRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataRelationships) GetOrder() POSTAdyenPayments201ResponseDataRelationshipsOrder {
	if o == nil || IsNil(o.Order) {
		var ret POSTAdyenPayments201ResponseDataRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataRelationships) GetOrderOk() (*POSTAdyenPayments201ResponseDataRelationshipsOrder, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationships) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given POSTAdyenPayments201ResponseDataRelationshipsOrder and assigns it to the Order field.
func (o *POSTLineItems201ResponseDataRelationships) SetOrder(v POSTAdyenPayments201ResponseDataRelationshipsOrder) {
	o.Order = &v
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataRelationships) GetItem() POSTLineItems201ResponseDataRelationshipsItem {
	if o == nil || IsNil(o.Item) {
		var ret POSTLineItems201ResponseDataRelationshipsItem
		return ret
	}
	return *o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataRelationships) GetItemOk() (*POSTLineItems201ResponseDataRelationshipsItem, bool) {
	if o == nil || IsNil(o.Item) {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationships) HasItem() bool {
	if o != nil && !IsNil(o.Item) {
		return true
	}

	return false
}

// SetItem gets a reference to the given POSTLineItems201ResponseDataRelationshipsItem and assigns it to the Item field.
func (o *POSTLineItems201ResponseDataRelationships) SetItem(v POSTLineItems201ResponseDataRelationshipsItem) {
	o.Item = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataRelationships) GetSku() POSTInStockSubscriptions201ResponseDataRelationshipsSku {
	if o == nil || IsNil(o.Sku) {
		var ret POSTInStockSubscriptions201ResponseDataRelationshipsSku
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataRelationships) GetSkuOk() (*POSTInStockSubscriptions201ResponseDataRelationshipsSku, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationships) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given POSTInStockSubscriptions201ResponseDataRelationshipsSku and assigns it to the Sku field.
func (o *POSTLineItems201ResponseDataRelationships) SetSku(v POSTInStockSubscriptions201ResponseDataRelationshipsSku) {
	o.Sku = &v
}

// GetBundle returns the Bundle field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataRelationships) GetBundle() POSTLineItems201ResponseDataRelationshipsBundle {
	if o == nil || IsNil(o.Bundle) {
		var ret POSTLineItems201ResponseDataRelationshipsBundle
		return ret
	}
	return *o.Bundle
}

// GetBundleOk returns a tuple with the Bundle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataRelationships) GetBundleOk() (*POSTLineItems201ResponseDataRelationshipsBundle, bool) {
	if o == nil || IsNil(o.Bundle) {
		return nil, false
	}
	return o.Bundle, true
}

// HasBundle returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationships) HasBundle() bool {
	if o != nil && !IsNil(o.Bundle) {
		return true
	}

	return false
}

// SetBundle gets a reference to the given POSTLineItems201ResponseDataRelationshipsBundle and assigns it to the Bundle field.
func (o *POSTLineItems201ResponseDataRelationships) SetBundle(v POSTLineItems201ResponseDataRelationshipsBundle) {
	o.Bundle = &v
}

// GetAdjustment returns the Adjustment field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataRelationships) GetAdjustment() POSTLineItems201ResponseDataRelationshipsAdjustment {
	if o == nil || IsNil(o.Adjustment) {
		var ret POSTLineItems201ResponseDataRelationshipsAdjustment
		return ret
	}
	return *o.Adjustment
}

// GetAdjustmentOk returns a tuple with the Adjustment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataRelationships) GetAdjustmentOk() (*POSTLineItems201ResponseDataRelationshipsAdjustment, bool) {
	if o == nil || IsNil(o.Adjustment) {
		return nil, false
	}
	return o.Adjustment, true
}

// HasAdjustment returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationships) HasAdjustment() bool {
	if o != nil && !IsNil(o.Adjustment) {
		return true
	}

	return false
}

// SetAdjustment gets a reference to the given POSTLineItems201ResponseDataRelationshipsAdjustment and assigns it to the Adjustment field.
func (o *POSTLineItems201ResponseDataRelationships) SetAdjustment(v POSTLineItems201ResponseDataRelationshipsAdjustment) {
	o.Adjustment = &v
}

// GetGiftCard returns the GiftCard field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataRelationships) GetGiftCard() POSTLineItems201ResponseDataRelationshipsGiftCard {
	if o == nil || IsNil(o.GiftCard) {
		var ret POSTLineItems201ResponseDataRelationshipsGiftCard
		return ret
	}
	return *o.GiftCard
}

// GetGiftCardOk returns a tuple with the GiftCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataRelationships) GetGiftCardOk() (*POSTLineItems201ResponseDataRelationshipsGiftCard, bool) {
	if o == nil || IsNil(o.GiftCard) {
		return nil, false
	}
	return o.GiftCard, true
}

// HasGiftCard returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationships) HasGiftCard() bool {
	if o != nil && !IsNil(o.GiftCard) {
		return true
	}

	return false
}

// SetGiftCard gets a reference to the given POSTLineItems201ResponseDataRelationshipsGiftCard and assigns it to the GiftCard field.
func (o *POSTLineItems201ResponseDataRelationships) SetGiftCard(v POSTLineItems201ResponseDataRelationshipsGiftCard) {
	o.GiftCard = &v
}

// GetShipment returns the Shipment field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataRelationships) GetShipment() POSTLineItems201ResponseDataRelationshipsShipment {
	if o == nil || IsNil(o.Shipment) {
		var ret POSTLineItems201ResponseDataRelationshipsShipment
		return ret
	}
	return *o.Shipment
}

// GetShipmentOk returns a tuple with the Shipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataRelationships) GetShipmentOk() (*POSTLineItems201ResponseDataRelationshipsShipment, bool) {
	if o == nil || IsNil(o.Shipment) {
		return nil, false
	}
	return o.Shipment, true
}

// HasShipment returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationships) HasShipment() bool {
	if o != nil && !IsNil(o.Shipment) {
		return true
	}

	return false
}

// SetShipment gets a reference to the given POSTLineItems201ResponseDataRelationshipsShipment and assigns it to the Shipment field.
func (o *POSTLineItems201ResponseDataRelationships) SetShipment(v POSTLineItems201ResponseDataRelationshipsShipment) {
	o.Shipment = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataRelationships) GetPaymentMethod() POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentMethod {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentMethod
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataRelationships) GetPaymentMethodOk() (*POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentMethod, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationships) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentMethod and assigns it to the PaymentMethod field.
func (o *POSTLineItems201ResponseDataRelationships) SetPaymentMethod(v POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentMethod) {
	o.PaymentMethod = &v
}

// GetLineItemOptions returns the LineItemOptions field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataRelationships) GetLineItemOptions() POSTLineItems201ResponseDataRelationshipsLineItemOptions {
	if o == nil || IsNil(o.LineItemOptions) {
		var ret POSTLineItems201ResponseDataRelationshipsLineItemOptions
		return ret
	}
	return *o.LineItemOptions
}

// GetLineItemOptionsOk returns a tuple with the LineItemOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataRelationships) GetLineItemOptionsOk() (*POSTLineItems201ResponseDataRelationshipsLineItemOptions, bool) {
	if o == nil || IsNil(o.LineItemOptions) {
		return nil, false
	}
	return o.LineItemOptions, true
}

// HasLineItemOptions returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationships) HasLineItemOptions() bool {
	if o != nil && !IsNil(o.LineItemOptions) {
		return true
	}

	return false
}

// SetLineItemOptions gets a reference to the given POSTLineItems201ResponseDataRelationshipsLineItemOptions and assigns it to the LineItemOptions field.
func (o *POSTLineItems201ResponseDataRelationships) SetLineItemOptions(v POSTLineItems201ResponseDataRelationshipsLineItemOptions) {
	o.LineItemOptions = &v
}

// GetReturnLineItems returns the ReturnLineItems field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataRelationships) GetReturnLineItems() POSTLineItems201ResponseDataRelationshipsReturnLineItems {
	if o == nil || IsNil(o.ReturnLineItems) {
		var ret POSTLineItems201ResponseDataRelationshipsReturnLineItems
		return ret
	}
	return *o.ReturnLineItems
}

// GetReturnLineItemsOk returns a tuple with the ReturnLineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataRelationships) GetReturnLineItemsOk() (*POSTLineItems201ResponseDataRelationshipsReturnLineItems, bool) {
	if o == nil || IsNil(o.ReturnLineItems) {
		return nil, false
	}
	return o.ReturnLineItems, true
}

// HasReturnLineItems returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationships) HasReturnLineItems() bool {
	if o != nil && !IsNil(o.ReturnLineItems) {
		return true
	}

	return false
}

// SetReturnLineItems gets a reference to the given POSTLineItems201ResponseDataRelationshipsReturnLineItems and assigns it to the ReturnLineItems field.
func (o *POSTLineItems201ResponseDataRelationships) SetReturnLineItems(v POSTLineItems201ResponseDataRelationshipsReturnLineItems) {
	o.ReturnLineItems = &v
}

// GetStockReservations returns the StockReservations field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataRelationships) GetStockReservations() POSTLineItems201ResponseDataRelationshipsStockReservations {
	if o == nil || IsNil(o.StockReservations) {
		var ret POSTLineItems201ResponseDataRelationshipsStockReservations
		return ret
	}
	return *o.StockReservations
}

// GetStockReservationsOk returns a tuple with the StockReservations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataRelationships) GetStockReservationsOk() (*POSTLineItems201ResponseDataRelationshipsStockReservations, bool) {
	if o == nil || IsNil(o.StockReservations) {
		return nil, false
	}
	return o.StockReservations, true
}

// HasStockReservations returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationships) HasStockReservations() bool {
	if o != nil && !IsNil(o.StockReservations) {
		return true
	}

	return false
}

// SetStockReservations gets a reference to the given POSTLineItems201ResponseDataRelationshipsStockReservations and assigns it to the StockReservations field.
func (o *POSTLineItems201ResponseDataRelationships) SetStockReservations(v POSTLineItems201ResponseDataRelationshipsStockReservations) {
	o.StockReservations = &v
}

// GetStockLineItems returns the StockLineItems field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataRelationships) GetStockLineItems() POSTLineItems201ResponseDataRelationshipsStockLineItems {
	if o == nil || IsNil(o.StockLineItems) {
		var ret POSTLineItems201ResponseDataRelationshipsStockLineItems
		return ret
	}
	return *o.StockLineItems
}

// GetStockLineItemsOk returns a tuple with the StockLineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataRelationships) GetStockLineItemsOk() (*POSTLineItems201ResponseDataRelationshipsStockLineItems, bool) {
	if o == nil || IsNil(o.StockLineItems) {
		return nil, false
	}
	return o.StockLineItems, true
}

// HasStockLineItems returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationships) HasStockLineItems() bool {
	if o != nil && !IsNil(o.StockLineItems) {
		return true
	}

	return false
}

// SetStockLineItems gets a reference to the given POSTLineItems201ResponseDataRelationshipsStockLineItems and assigns it to the StockLineItems field.
func (o *POSTLineItems201ResponseDataRelationships) SetStockLineItems(v POSTLineItems201ResponseDataRelationshipsStockLineItems) {
	o.StockLineItems = &v
}

// GetStockTransfers returns the StockTransfers field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataRelationships) GetStockTransfers() POSTLineItems201ResponseDataRelationshipsStockTransfers {
	if o == nil || IsNil(o.StockTransfers) {
		var ret POSTLineItems201ResponseDataRelationshipsStockTransfers
		return ret
	}
	return *o.StockTransfers
}

// GetStockTransfersOk returns a tuple with the StockTransfers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataRelationships) GetStockTransfersOk() (*POSTLineItems201ResponseDataRelationshipsStockTransfers, bool) {
	if o == nil || IsNil(o.StockTransfers) {
		return nil, false
	}
	return o.StockTransfers, true
}

// HasStockTransfers returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationships) HasStockTransfers() bool {
	if o != nil && !IsNil(o.StockTransfers) {
		return true
	}

	return false
}

// SetStockTransfers gets a reference to the given POSTLineItems201ResponseDataRelationshipsStockTransfers and assigns it to the StockTransfers field.
func (o *POSTLineItems201ResponseDataRelationships) SetStockTransfers(v POSTLineItems201ResponseDataRelationshipsStockTransfers) {
	o.StockTransfers = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataRelationships) GetNotifications() POSTLineItems201ResponseDataRelationshipsNotifications {
	if o == nil || IsNil(o.Notifications) {
		var ret POSTLineItems201ResponseDataRelationshipsNotifications
		return ret
	}
	return *o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataRelationships) GetNotificationsOk() (*POSTLineItems201ResponseDataRelationshipsNotifications, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationships) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given POSTLineItems201ResponseDataRelationshipsNotifications and assigns it to the Notifications field.
func (o *POSTLineItems201ResponseDataRelationships) SetNotifications(v POSTLineItems201ResponseDataRelationshipsNotifications) {
	o.Notifications = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataRelationships) GetEvents() POSTAddresses201ResponseDataRelationshipsEvents {
	if o == nil || IsNil(o.Events) {
		var ret POSTAddresses201ResponseDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataRelationships) GetEventsOk() (*POSTAddresses201ResponseDataRelationshipsEvents, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationships) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given POSTAddresses201ResponseDataRelationshipsEvents and assigns it to the Events field.
func (o *POSTLineItems201ResponseDataRelationships) SetEvents(v POSTAddresses201ResponseDataRelationshipsEvents) {
	o.Events = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataRelationships) GetTags() POSTAddresses201ResponseDataRelationshipsTags {
	if o == nil || IsNil(o.Tags) {
		var ret POSTAddresses201ResponseDataRelationshipsTags
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataRelationships) GetTagsOk() (*POSTAddresses201ResponseDataRelationshipsTags, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationships) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given POSTAddresses201ResponseDataRelationshipsTags and assigns it to the Tags field.
func (o *POSTLineItems201ResponseDataRelationships) SetTags(v POSTAddresses201ResponseDataRelationshipsTags) {
	o.Tags = &v
}

func (o POSTLineItems201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTLineItems201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Item) {
		toSerialize["item"] = o.Item
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.Bundle) {
		toSerialize["bundle"] = o.Bundle
	}
	if !IsNil(o.Adjustment) {
		toSerialize["adjustment"] = o.Adjustment
	}
	if !IsNil(o.GiftCard) {
		toSerialize["gift_card"] = o.GiftCard
	}
	if !IsNil(o.Shipment) {
		toSerialize["shipment"] = o.Shipment
	}
	if !IsNil(o.PaymentMethod) {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	if !IsNil(o.LineItemOptions) {
		toSerialize["line_item_options"] = o.LineItemOptions
	}
	if !IsNil(o.ReturnLineItems) {
		toSerialize["return_line_items"] = o.ReturnLineItems
	}
	if !IsNil(o.StockReservations) {
		toSerialize["stock_reservations"] = o.StockReservations
	}
	if !IsNil(o.StockLineItems) {
		toSerialize["stock_line_items"] = o.StockLineItems
	}
	if !IsNil(o.StockTransfers) {
		toSerialize["stock_transfers"] = o.StockTransfers
	}
	if !IsNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullablePOSTLineItems201ResponseDataRelationships struct {
	value *POSTLineItems201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTLineItems201ResponseDataRelationships) Get() *POSTLineItems201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTLineItems201ResponseDataRelationships) Set(val *POSTLineItems201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTLineItems201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTLineItems201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTLineItems201ResponseDataRelationships(val *POSTLineItems201ResponseDataRelationships) *NullablePOSTLineItems201ResponseDataRelationships {
	return &NullablePOSTLineItems201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTLineItems201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTLineItems201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
