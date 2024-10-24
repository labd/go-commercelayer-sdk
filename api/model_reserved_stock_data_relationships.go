/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ReservedStockDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReservedStockDataRelationships{}

// ReservedStockDataRelationships struct for ReservedStockDataRelationships
type ReservedStockDataRelationships struct {
	StockItem         *ReservedStockDataRelationshipsStockItem    `json:"stock_item,omitempty"`
	Sku               *BundleDataRelationshipsSkus                `json:"sku,omitempty"`
	StockReservations *LineItemDataRelationshipsStockReservations `json:"stock_reservations,omitempty"`
	Versions          *AddressDataRelationshipsVersions           `json:"versions,omitempty"`
}

// NewReservedStockDataRelationships instantiates a new ReservedStockDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReservedStockDataRelationships() *ReservedStockDataRelationships {
	this := ReservedStockDataRelationships{}
	return &this
}

// NewReservedStockDataRelationshipsWithDefaults instantiates a new ReservedStockDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReservedStockDataRelationshipsWithDefaults() *ReservedStockDataRelationships {
	this := ReservedStockDataRelationships{}
	return &this
}

// GetStockItem returns the StockItem field value if set, zero value otherwise.
func (o *ReservedStockDataRelationships) GetStockItem() ReservedStockDataRelationshipsStockItem {
	if o == nil || IsNil(o.StockItem) {
		var ret ReservedStockDataRelationshipsStockItem
		return ret
	}
	return *o.StockItem
}

// GetStockItemOk returns a tuple with the StockItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedStockDataRelationships) GetStockItemOk() (*ReservedStockDataRelationshipsStockItem, bool) {
	if o == nil || IsNil(o.StockItem) {
		return nil, false
	}
	return o.StockItem, true
}

// HasStockItem returns a boolean if a field has been set.
func (o *ReservedStockDataRelationships) HasStockItem() bool {
	if o != nil && !IsNil(o.StockItem) {
		return true
	}

	return false
}

// SetStockItem gets a reference to the given ReservedStockDataRelationshipsStockItem and assigns it to the StockItem field.
func (o *ReservedStockDataRelationships) SetStockItem(v ReservedStockDataRelationshipsStockItem) {
	o.StockItem = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *ReservedStockDataRelationships) GetSku() BundleDataRelationshipsSkus {
	if o == nil || IsNil(o.Sku) {
		var ret BundleDataRelationshipsSkus
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedStockDataRelationships) GetSkuOk() (*BundleDataRelationshipsSkus, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *ReservedStockDataRelationships) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given BundleDataRelationshipsSkus and assigns it to the Sku field.
func (o *ReservedStockDataRelationships) SetSku(v BundleDataRelationshipsSkus) {
	o.Sku = &v
}

// GetStockReservations returns the StockReservations field value if set, zero value otherwise.
func (o *ReservedStockDataRelationships) GetStockReservations() LineItemDataRelationshipsStockReservations {
	if o == nil || IsNil(o.StockReservations) {
		var ret LineItemDataRelationshipsStockReservations
		return ret
	}
	return *o.StockReservations
}

// GetStockReservationsOk returns a tuple with the StockReservations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedStockDataRelationships) GetStockReservationsOk() (*LineItemDataRelationshipsStockReservations, bool) {
	if o == nil || IsNil(o.StockReservations) {
		return nil, false
	}
	return o.StockReservations, true
}

// HasStockReservations returns a boolean if a field has been set.
func (o *ReservedStockDataRelationships) HasStockReservations() bool {
	if o != nil && !IsNil(o.StockReservations) {
		return true
	}

	return false
}

// SetStockReservations gets a reference to the given LineItemDataRelationshipsStockReservations and assigns it to the StockReservations field.
func (o *ReservedStockDataRelationships) SetStockReservations(v LineItemDataRelationshipsStockReservations) {
	o.StockReservations = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *ReservedStockDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedStockDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *ReservedStockDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *ReservedStockDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

func (o ReservedStockDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReservedStockDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StockItem) {
		toSerialize["stock_item"] = o.StockItem
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.StockReservations) {
		toSerialize["stock_reservations"] = o.StockReservations
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullableReservedStockDataRelationships struct {
	value *ReservedStockDataRelationships
	isSet bool
}

func (v NullableReservedStockDataRelationships) Get() *ReservedStockDataRelationships {
	return v.value
}

func (v *NullableReservedStockDataRelationships) Set(val *ReservedStockDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableReservedStockDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableReservedStockDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservedStockDataRelationships(val *ReservedStockDataRelationships) *NullableReservedStockDataRelationships {
	return &NullableReservedStockDataRelationships{value: val, isSet: true}
}

func (v NullableReservedStockDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservedStockDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}