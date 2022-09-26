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

// ParcelLineItemResponseDataRelationships struct for ParcelLineItemResponseDataRelationships
type ParcelLineItemResponseDataRelationships struct {
	Parcel           *ParcelLineItemResponseDataRelationshipsParcel           `json:"parcel,omitempty"`
	StockLineItem    *ParcelLineItemResponseDataRelationshipsStockLineItem    `json:"stock_line_item,omitempty"`
	ShipmentLineItem *ParcelLineItemResponseDataRelationshipsShipmentLineItem `json:"shipment_line_item,omitempty"`
}

// NewParcelLineItemResponseDataRelationships instantiates a new ParcelLineItemResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParcelLineItemResponseDataRelationships() *ParcelLineItemResponseDataRelationships {
	this := ParcelLineItemResponseDataRelationships{}
	return &this
}

// NewParcelLineItemResponseDataRelationshipsWithDefaults instantiates a new ParcelLineItemResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParcelLineItemResponseDataRelationshipsWithDefaults() *ParcelLineItemResponseDataRelationships {
	this := ParcelLineItemResponseDataRelationships{}
	return &this
}

// GetParcel returns the Parcel field value if set, zero value otherwise.
func (o *ParcelLineItemResponseDataRelationships) GetParcel() ParcelLineItemResponseDataRelationshipsParcel {
	if o == nil || o.Parcel == nil {
		var ret ParcelLineItemResponseDataRelationshipsParcel
		return ret
	}
	return *o.Parcel
}

// GetParcelOk returns a tuple with the Parcel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelLineItemResponseDataRelationships) GetParcelOk() (*ParcelLineItemResponseDataRelationshipsParcel, bool) {
	if o == nil || o.Parcel == nil {
		return nil, false
	}
	return o.Parcel, true
}

// HasParcel returns a boolean if a field has been set.
func (o *ParcelLineItemResponseDataRelationships) HasParcel() bool {
	if o != nil && o.Parcel != nil {
		return true
	}

	return false
}

// SetParcel gets a reference to the given ParcelLineItemResponseDataRelationshipsParcel and assigns it to the Parcel field.
func (o *ParcelLineItemResponseDataRelationships) SetParcel(v ParcelLineItemResponseDataRelationshipsParcel) {
	o.Parcel = &v
}

// GetStockLineItem returns the StockLineItem field value if set, zero value otherwise.
func (o *ParcelLineItemResponseDataRelationships) GetStockLineItem() ParcelLineItemResponseDataRelationshipsStockLineItem {
	if o == nil || o.StockLineItem == nil {
		var ret ParcelLineItemResponseDataRelationshipsStockLineItem
		return ret
	}
	return *o.StockLineItem
}

// GetStockLineItemOk returns a tuple with the StockLineItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelLineItemResponseDataRelationships) GetStockLineItemOk() (*ParcelLineItemResponseDataRelationshipsStockLineItem, bool) {
	if o == nil || o.StockLineItem == nil {
		return nil, false
	}
	return o.StockLineItem, true
}

// HasStockLineItem returns a boolean if a field has been set.
func (o *ParcelLineItemResponseDataRelationships) HasStockLineItem() bool {
	if o != nil && o.StockLineItem != nil {
		return true
	}

	return false
}

// SetStockLineItem gets a reference to the given ParcelLineItemResponseDataRelationshipsStockLineItem and assigns it to the StockLineItem field.
func (o *ParcelLineItemResponseDataRelationships) SetStockLineItem(v ParcelLineItemResponseDataRelationshipsStockLineItem) {
	o.StockLineItem = &v
}

// GetShipmentLineItem returns the ShipmentLineItem field value if set, zero value otherwise.
func (o *ParcelLineItemResponseDataRelationships) GetShipmentLineItem() ParcelLineItemResponseDataRelationshipsShipmentLineItem {
	if o == nil || o.ShipmentLineItem == nil {
		var ret ParcelLineItemResponseDataRelationshipsShipmentLineItem
		return ret
	}
	return *o.ShipmentLineItem
}

// GetShipmentLineItemOk returns a tuple with the ShipmentLineItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelLineItemResponseDataRelationships) GetShipmentLineItemOk() (*ParcelLineItemResponseDataRelationshipsShipmentLineItem, bool) {
	if o == nil || o.ShipmentLineItem == nil {
		return nil, false
	}
	return o.ShipmentLineItem, true
}

// HasShipmentLineItem returns a boolean if a field has been set.
func (o *ParcelLineItemResponseDataRelationships) HasShipmentLineItem() bool {
	if o != nil && o.ShipmentLineItem != nil {
		return true
	}

	return false
}

// SetShipmentLineItem gets a reference to the given ParcelLineItemResponseDataRelationshipsShipmentLineItem and assigns it to the ShipmentLineItem field.
func (o *ParcelLineItemResponseDataRelationships) SetShipmentLineItem(v ParcelLineItemResponseDataRelationshipsShipmentLineItem) {
	o.ShipmentLineItem = &v
}

func (o ParcelLineItemResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Parcel != nil {
		toSerialize["parcel"] = o.Parcel
	}
	if o.StockLineItem != nil {
		toSerialize["stock_line_item"] = o.StockLineItem
	}
	if o.ShipmentLineItem != nil {
		toSerialize["shipment_line_item"] = o.ShipmentLineItem
	}
	return json.Marshal(toSerialize)
}

type NullableParcelLineItemResponseDataRelationships struct {
	value *ParcelLineItemResponseDataRelationships
	isSet bool
}

func (v NullableParcelLineItemResponseDataRelationships) Get() *ParcelLineItemResponseDataRelationships {
	return v.value
}

func (v *NullableParcelLineItemResponseDataRelationships) Set(val *ParcelLineItemResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableParcelLineItemResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableParcelLineItemResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParcelLineItemResponseDataRelationships(val *ParcelLineItemResponseDataRelationships) *NullableParcelLineItemResponseDataRelationships {
	return &NullableParcelLineItemResponseDataRelationships{value: val, isSet: true}
}

func (v NullableParcelLineItemResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParcelLineItemResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
