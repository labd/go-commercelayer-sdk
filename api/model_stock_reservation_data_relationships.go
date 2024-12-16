/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.5.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the StockReservationDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StockReservationDataRelationships{}

// StockReservationDataRelationships struct for StockReservationDataRelationships
type StockReservationDataRelationships struct {
	LineItem      *LineItemOptionDataRelationshipsLineItem `json:"line_item,omitempty"`
	Order         *AdyenPaymentDataRelationshipsOrder      `json:"order,omitempty"`
	StockLineItem *LineItemDataRelationshipsStockLineItems `json:"stock_line_item,omitempty"`
	StockTransfer *LineItemDataRelationshipsStockTransfers `json:"stock_transfer,omitempty"`
	StockItem     *ReservedStockDataRelationshipsStockItem `json:"stock_item,omitempty"`
	ReservedStock *StockItemDataRelationshipsReservedStock `json:"reserved_stock,omitempty"`
	Sku           *BundleDataRelationshipsSkus             `json:"sku,omitempty"`
}

// NewStockReservationDataRelationships instantiates a new StockReservationDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockReservationDataRelationships() *StockReservationDataRelationships {
	this := StockReservationDataRelationships{}
	return &this
}

// NewStockReservationDataRelationshipsWithDefaults instantiates a new StockReservationDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockReservationDataRelationshipsWithDefaults() *StockReservationDataRelationships {
	this := StockReservationDataRelationships{}
	return &this
}

// GetLineItem returns the LineItem field value if set, zero value otherwise.
func (o *StockReservationDataRelationships) GetLineItem() LineItemOptionDataRelationshipsLineItem {
	if o == nil || IsNil(o.LineItem) {
		var ret LineItemOptionDataRelationshipsLineItem
		return ret
	}
	return *o.LineItem
}

// GetLineItemOk returns a tuple with the LineItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockReservationDataRelationships) GetLineItemOk() (*LineItemOptionDataRelationshipsLineItem, bool) {
	if o == nil || IsNil(o.LineItem) {
		return nil, false
	}
	return o.LineItem, true
}

// HasLineItem returns a boolean if a field has been set.
func (o *StockReservationDataRelationships) HasLineItem() bool {
	if o != nil && !IsNil(o.LineItem) {
		return true
	}

	return false
}

// SetLineItem gets a reference to the given LineItemOptionDataRelationshipsLineItem and assigns it to the LineItem field.
func (o *StockReservationDataRelationships) SetLineItem(v LineItemOptionDataRelationshipsLineItem) {
	o.LineItem = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *StockReservationDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder {
	if o == nil || IsNil(o.Order) {
		var ret AdyenPaymentDataRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockReservationDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *StockReservationDataRelationships) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given AdyenPaymentDataRelationshipsOrder and assigns it to the Order field.
func (o *StockReservationDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder) {
	o.Order = &v
}

// GetStockLineItem returns the StockLineItem field value if set, zero value otherwise.
func (o *StockReservationDataRelationships) GetStockLineItem() LineItemDataRelationshipsStockLineItems {
	if o == nil || IsNil(o.StockLineItem) {
		var ret LineItemDataRelationshipsStockLineItems
		return ret
	}
	return *o.StockLineItem
}

// GetStockLineItemOk returns a tuple with the StockLineItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockReservationDataRelationships) GetStockLineItemOk() (*LineItemDataRelationshipsStockLineItems, bool) {
	if o == nil || IsNil(o.StockLineItem) {
		return nil, false
	}
	return o.StockLineItem, true
}

// HasStockLineItem returns a boolean if a field has been set.
func (o *StockReservationDataRelationships) HasStockLineItem() bool {
	if o != nil && !IsNil(o.StockLineItem) {
		return true
	}

	return false
}

// SetStockLineItem gets a reference to the given LineItemDataRelationshipsStockLineItems and assigns it to the StockLineItem field.
func (o *StockReservationDataRelationships) SetStockLineItem(v LineItemDataRelationshipsStockLineItems) {
	o.StockLineItem = &v
}

// GetStockTransfer returns the StockTransfer field value if set, zero value otherwise.
func (o *StockReservationDataRelationships) GetStockTransfer() LineItemDataRelationshipsStockTransfers {
	if o == nil || IsNil(o.StockTransfer) {
		var ret LineItemDataRelationshipsStockTransfers
		return ret
	}
	return *o.StockTransfer
}

// GetStockTransferOk returns a tuple with the StockTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockReservationDataRelationships) GetStockTransferOk() (*LineItemDataRelationshipsStockTransfers, bool) {
	if o == nil || IsNil(o.StockTransfer) {
		return nil, false
	}
	return o.StockTransfer, true
}

// HasStockTransfer returns a boolean if a field has been set.
func (o *StockReservationDataRelationships) HasStockTransfer() bool {
	if o != nil && !IsNil(o.StockTransfer) {
		return true
	}

	return false
}

// SetStockTransfer gets a reference to the given LineItemDataRelationshipsStockTransfers and assigns it to the StockTransfer field.
func (o *StockReservationDataRelationships) SetStockTransfer(v LineItemDataRelationshipsStockTransfers) {
	o.StockTransfer = &v
}

// GetStockItem returns the StockItem field value if set, zero value otherwise.
func (o *StockReservationDataRelationships) GetStockItem() ReservedStockDataRelationshipsStockItem {
	if o == nil || IsNil(o.StockItem) {
		var ret ReservedStockDataRelationshipsStockItem
		return ret
	}
	return *o.StockItem
}

// GetStockItemOk returns a tuple with the StockItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockReservationDataRelationships) GetStockItemOk() (*ReservedStockDataRelationshipsStockItem, bool) {
	if o == nil || IsNil(o.StockItem) {
		return nil, false
	}
	return o.StockItem, true
}

// HasStockItem returns a boolean if a field has been set.
func (o *StockReservationDataRelationships) HasStockItem() bool {
	if o != nil && !IsNil(o.StockItem) {
		return true
	}

	return false
}

// SetStockItem gets a reference to the given ReservedStockDataRelationshipsStockItem and assigns it to the StockItem field.
func (o *StockReservationDataRelationships) SetStockItem(v ReservedStockDataRelationshipsStockItem) {
	o.StockItem = &v
}

// GetReservedStock returns the ReservedStock field value if set, zero value otherwise.
func (o *StockReservationDataRelationships) GetReservedStock() StockItemDataRelationshipsReservedStock {
	if o == nil || IsNil(o.ReservedStock) {
		var ret StockItemDataRelationshipsReservedStock
		return ret
	}
	return *o.ReservedStock
}

// GetReservedStockOk returns a tuple with the ReservedStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockReservationDataRelationships) GetReservedStockOk() (*StockItemDataRelationshipsReservedStock, bool) {
	if o == nil || IsNil(o.ReservedStock) {
		return nil, false
	}
	return o.ReservedStock, true
}

// HasReservedStock returns a boolean if a field has been set.
func (o *StockReservationDataRelationships) HasReservedStock() bool {
	if o != nil && !IsNil(o.ReservedStock) {
		return true
	}

	return false
}

// SetReservedStock gets a reference to the given StockItemDataRelationshipsReservedStock and assigns it to the ReservedStock field.
func (o *StockReservationDataRelationships) SetReservedStock(v StockItemDataRelationshipsReservedStock) {
	o.ReservedStock = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *StockReservationDataRelationships) GetSku() BundleDataRelationshipsSkus {
	if o == nil || IsNil(o.Sku) {
		var ret BundleDataRelationshipsSkus
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockReservationDataRelationships) GetSkuOk() (*BundleDataRelationshipsSkus, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *StockReservationDataRelationships) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given BundleDataRelationshipsSkus and assigns it to the Sku field.
func (o *StockReservationDataRelationships) SetSku(v BundleDataRelationshipsSkus) {
	o.Sku = &v
}

func (o StockReservationDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StockReservationDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LineItem) {
		toSerialize["line_item"] = o.LineItem
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.StockLineItem) {
		toSerialize["stock_line_item"] = o.StockLineItem
	}
	if !IsNil(o.StockTransfer) {
		toSerialize["stock_transfer"] = o.StockTransfer
	}
	if !IsNil(o.StockItem) {
		toSerialize["stock_item"] = o.StockItem
	}
	if !IsNil(o.ReservedStock) {
		toSerialize["reserved_stock"] = o.ReservedStock
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	return toSerialize, nil
}

type NullableStockReservationDataRelationships struct {
	value *StockReservationDataRelationships
	isSet bool
}

func (v NullableStockReservationDataRelationships) Get() *StockReservationDataRelationships {
	return v.value
}

func (v *NullableStockReservationDataRelationships) Set(val *StockReservationDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableStockReservationDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableStockReservationDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockReservationDataRelationships(val *StockReservationDataRelationships) *NullableStockReservationDataRelationships {
	return &NullableStockReservationDataRelationships{value: val, isSet: true}
}

func (v NullableStockReservationDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockReservationDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
