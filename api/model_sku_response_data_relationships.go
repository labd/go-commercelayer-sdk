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

// SkuResponseDataRelationships struct for SkuResponseDataRelationships
type SkuResponseDataRelationships struct {
	ShippingCategory  *ShipmentResponseDataRelationshipsShippingCategory  `json:"shipping_category,omitempty"`
	Prices            *PriceListResponseDataRelationshipsPrices           `json:"prices,omitempty"`
	StockItems        *SkuResponseDataRelationshipsStockItems             `json:"stock_items,omitempty"`
	DeliveryLeadTimes *SkuResponseDataRelationshipsDeliveryLeadTimes      `json:"delivery_lead_times,omitempty"`
	SkuOptions        *SkuResponseDataRelationshipsSkuOptions             `json:"sku_options,omitempty"`
	Attachments       *AvalaraAccountResponseDataRelationshipsAttachments `json:"attachments,omitempty"`
}

// NewSkuResponseDataRelationships instantiates a new SkuResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuResponseDataRelationships() *SkuResponseDataRelationships {
	this := SkuResponseDataRelationships{}
	return &this
}

// NewSkuResponseDataRelationshipsWithDefaults instantiates a new SkuResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuResponseDataRelationshipsWithDefaults() *SkuResponseDataRelationships {
	this := SkuResponseDataRelationships{}
	return &this
}

// GetShippingCategory returns the ShippingCategory field value if set, zero value otherwise.
func (o *SkuResponseDataRelationships) GetShippingCategory() ShipmentResponseDataRelationshipsShippingCategory {
	if o == nil || o.ShippingCategory == nil {
		var ret ShipmentResponseDataRelationshipsShippingCategory
		return ret
	}
	return *o.ShippingCategory
}

// GetShippingCategoryOk returns a tuple with the ShippingCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuResponseDataRelationships) GetShippingCategoryOk() (*ShipmentResponseDataRelationshipsShippingCategory, bool) {
	if o == nil || o.ShippingCategory == nil {
		return nil, false
	}
	return o.ShippingCategory, true
}

// HasShippingCategory returns a boolean if a field has been set.
func (o *SkuResponseDataRelationships) HasShippingCategory() bool {
	if o != nil && o.ShippingCategory != nil {
		return true
	}

	return false
}

// SetShippingCategory gets a reference to the given ShipmentResponseDataRelationshipsShippingCategory and assigns it to the ShippingCategory field.
func (o *SkuResponseDataRelationships) SetShippingCategory(v ShipmentResponseDataRelationshipsShippingCategory) {
	o.ShippingCategory = &v
}

// GetPrices returns the Prices field value if set, zero value otherwise.
func (o *SkuResponseDataRelationships) GetPrices() PriceListResponseDataRelationshipsPrices {
	if o == nil || o.Prices == nil {
		var ret PriceListResponseDataRelationshipsPrices
		return ret
	}
	return *o.Prices
}

// GetPricesOk returns a tuple with the Prices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuResponseDataRelationships) GetPricesOk() (*PriceListResponseDataRelationshipsPrices, bool) {
	if o == nil || o.Prices == nil {
		return nil, false
	}
	return o.Prices, true
}

// HasPrices returns a boolean if a field has been set.
func (o *SkuResponseDataRelationships) HasPrices() bool {
	if o != nil && o.Prices != nil {
		return true
	}

	return false
}

// SetPrices gets a reference to the given PriceListResponseDataRelationshipsPrices and assigns it to the Prices field.
func (o *SkuResponseDataRelationships) SetPrices(v PriceListResponseDataRelationshipsPrices) {
	o.Prices = &v
}

// GetStockItems returns the StockItems field value if set, zero value otherwise.
func (o *SkuResponseDataRelationships) GetStockItems() SkuResponseDataRelationshipsStockItems {
	if o == nil || o.StockItems == nil {
		var ret SkuResponseDataRelationshipsStockItems
		return ret
	}
	return *o.StockItems
}

// GetStockItemsOk returns a tuple with the StockItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuResponseDataRelationships) GetStockItemsOk() (*SkuResponseDataRelationshipsStockItems, bool) {
	if o == nil || o.StockItems == nil {
		return nil, false
	}
	return o.StockItems, true
}

// HasStockItems returns a boolean if a field has been set.
func (o *SkuResponseDataRelationships) HasStockItems() bool {
	if o != nil && o.StockItems != nil {
		return true
	}

	return false
}

// SetStockItems gets a reference to the given SkuResponseDataRelationshipsStockItems and assigns it to the StockItems field.
func (o *SkuResponseDataRelationships) SetStockItems(v SkuResponseDataRelationshipsStockItems) {
	o.StockItems = &v
}

// GetDeliveryLeadTimes returns the DeliveryLeadTimes field value if set, zero value otherwise.
func (o *SkuResponseDataRelationships) GetDeliveryLeadTimes() SkuResponseDataRelationshipsDeliveryLeadTimes {
	if o == nil || o.DeliveryLeadTimes == nil {
		var ret SkuResponseDataRelationshipsDeliveryLeadTimes
		return ret
	}
	return *o.DeliveryLeadTimes
}

// GetDeliveryLeadTimesOk returns a tuple with the DeliveryLeadTimes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuResponseDataRelationships) GetDeliveryLeadTimesOk() (*SkuResponseDataRelationshipsDeliveryLeadTimes, bool) {
	if o == nil || o.DeliveryLeadTimes == nil {
		return nil, false
	}
	return o.DeliveryLeadTimes, true
}

// HasDeliveryLeadTimes returns a boolean if a field has been set.
func (o *SkuResponseDataRelationships) HasDeliveryLeadTimes() bool {
	if o != nil && o.DeliveryLeadTimes != nil {
		return true
	}

	return false
}

// SetDeliveryLeadTimes gets a reference to the given SkuResponseDataRelationshipsDeliveryLeadTimes and assigns it to the DeliveryLeadTimes field.
func (o *SkuResponseDataRelationships) SetDeliveryLeadTimes(v SkuResponseDataRelationshipsDeliveryLeadTimes) {
	o.DeliveryLeadTimes = &v
}

// GetSkuOptions returns the SkuOptions field value if set, zero value otherwise.
func (o *SkuResponseDataRelationships) GetSkuOptions() SkuResponseDataRelationshipsSkuOptions {
	if o == nil || o.SkuOptions == nil {
		var ret SkuResponseDataRelationshipsSkuOptions
		return ret
	}
	return *o.SkuOptions
}

// GetSkuOptionsOk returns a tuple with the SkuOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuResponseDataRelationships) GetSkuOptionsOk() (*SkuResponseDataRelationshipsSkuOptions, bool) {
	if o == nil || o.SkuOptions == nil {
		return nil, false
	}
	return o.SkuOptions, true
}

// HasSkuOptions returns a boolean if a field has been set.
func (o *SkuResponseDataRelationships) HasSkuOptions() bool {
	if o != nil && o.SkuOptions != nil {
		return true
	}

	return false
}

// SetSkuOptions gets a reference to the given SkuResponseDataRelationshipsSkuOptions and assigns it to the SkuOptions field.
func (o *SkuResponseDataRelationships) SetSkuOptions(v SkuResponseDataRelationshipsSkuOptions) {
	o.SkuOptions = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *SkuResponseDataRelationships) GetAttachments() AvalaraAccountResponseDataRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret AvalaraAccountResponseDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuResponseDataRelationships) GetAttachmentsOk() (*AvalaraAccountResponseDataRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *SkuResponseDataRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountResponseDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *SkuResponseDataRelationships) SetAttachments(v AvalaraAccountResponseDataRelationshipsAttachments) {
	o.Attachments = &v
}

func (o SkuResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ShippingCategory != nil {
		toSerialize["shipping_category"] = o.ShippingCategory
	}
	if o.Prices != nil {
		toSerialize["prices"] = o.Prices
	}
	if o.StockItems != nil {
		toSerialize["stock_items"] = o.StockItems
	}
	if o.DeliveryLeadTimes != nil {
		toSerialize["delivery_lead_times"] = o.DeliveryLeadTimes
	}
	if o.SkuOptions != nil {
		toSerialize["sku_options"] = o.SkuOptions
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableSkuResponseDataRelationships struct {
	value *SkuResponseDataRelationships
	isSet bool
}

func (v NullableSkuResponseDataRelationships) Get() *SkuResponseDataRelationships {
	return v.value
}

func (v *NullableSkuResponseDataRelationships) Set(val *SkuResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuResponseDataRelationships(val *SkuResponseDataRelationships) *NullableSkuResponseDataRelationships {
	return &NullableSkuResponseDataRelationships{value: val, isSet: true}
}

func (v NullableSkuResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
