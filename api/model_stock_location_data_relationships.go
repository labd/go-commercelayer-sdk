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

// StockLocationDataRelationships struct for StockLocationDataRelationships
type StockLocationDataRelationships struct {
	Address                  *BingGeocoderDataRelationshipsAddresses                  `json:"address,omitempty"`
	InventoryStockLocations  *InventoryModelDataRelationshipsInventoryStockLocations  `json:"inventory_stock_locations,omitempty"`
	InventoryReturnLocations *InventoryModelDataRelationshipsInventoryReturnLocations `json:"inventory_return_locations,omitempty"`
	StockItems               *SkuDataRelationshipsStockItems                          `json:"stock_items,omitempty"`
	StockTransfers           *LineItemDataRelationshipsStockTransfers                 `json:"stock_transfers,omitempty"`
	Attachments              *AvalaraAccountDataRelationshipsAttachments              `json:"attachments,omitempty"`
}

// NewStockLocationDataRelationships instantiates a new StockLocationDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockLocationDataRelationships() *StockLocationDataRelationships {
	this := StockLocationDataRelationships{}
	return &this
}

// NewStockLocationDataRelationshipsWithDefaults instantiates a new StockLocationDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockLocationDataRelationshipsWithDefaults() *StockLocationDataRelationships {
	this := StockLocationDataRelationships{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *StockLocationDataRelationships) GetAddress() BingGeocoderDataRelationshipsAddresses {
	if o == nil || o.Address == nil {
		var ret BingGeocoderDataRelationshipsAddresses
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockLocationDataRelationships) GetAddressOk() (*BingGeocoderDataRelationshipsAddresses, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *StockLocationDataRelationships) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given BingGeocoderDataRelationshipsAddresses and assigns it to the Address field.
func (o *StockLocationDataRelationships) SetAddress(v BingGeocoderDataRelationshipsAddresses) {
	o.Address = &v
}

// GetInventoryStockLocations returns the InventoryStockLocations field value if set, zero value otherwise.
func (o *StockLocationDataRelationships) GetInventoryStockLocations() InventoryModelDataRelationshipsInventoryStockLocations {
	if o == nil || o.InventoryStockLocations == nil {
		var ret InventoryModelDataRelationshipsInventoryStockLocations
		return ret
	}
	return *o.InventoryStockLocations
}

// GetInventoryStockLocationsOk returns a tuple with the InventoryStockLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockLocationDataRelationships) GetInventoryStockLocationsOk() (*InventoryModelDataRelationshipsInventoryStockLocations, bool) {
	if o == nil || o.InventoryStockLocations == nil {
		return nil, false
	}
	return o.InventoryStockLocations, true
}

// HasInventoryStockLocations returns a boolean if a field has been set.
func (o *StockLocationDataRelationships) HasInventoryStockLocations() bool {
	if o != nil && o.InventoryStockLocations != nil {
		return true
	}

	return false
}

// SetInventoryStockLocations gets a reference to the given InventoryModelDataRelationshipsInventoryStockLocations and assigns it to the InventoryStockLocations field.
func (o *StockLocationDataRelationships) SetInventoryStockLocations(v InventoryModelDataRelationshipsInventoryStockLocations) {
	o.InventoryStockLocations = &v
}

// GetInventoryReturnLocations returns the InventoryReturnLocations field value if set, zero value otherwise.
func (o *StockLocationDataRelationships) GetInventoryReturnLocations() InventoryModelDataRelationshipsInventoryReturnLocations {
	if o == nil || o.InventoryReturnLocations == nil {
		var ret InventoryModelDataRelationshipsInventoryReturnLocations
		return ret
	}
	return *o.InventoryReturnLocations
}

// GetInventoryReturnLocationsOk returns a tuple with the InventoryReturnLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockLocationDataRelationships) GetInventoryReturnLocationsOk() (*InventoryModelDataRelationshipsInventoryReturnLocations, bool) {
	if o == nil || o.InventoryReturnLocations == nil {
		return nil, false
	}
	return o.InventoryReturnLocations, true
}

// HasInventoryReturnLocations returns a boolean if a field has been set.
func (o *StockLocationDataRelationships) HasInventoryReturnLocations() bool {
	if o != nil && o.InventoryReturnLocations != nil {
		return true
	}

	return false
}

// SetInventoryReturnLocations gets a reference to the given InventoryModelDataRelationshipsInventoryReturnLocations and assigns it to the InventoryReturnLocations field.
func (o *StockLocationDataRelationships) SetInventoryReturnLocations(v InventoryModelDataRelationshipsInventoryReturnLocations) {
	o.InventoryReturnLocations = &v
}

// GetStockItems returns the StockItems field value if set, zero value otherwise.
func (o *StockLocationDataRelationships) GetStockItems() SkuDataRelationshipsStockItems {
	if o == nil || o.StockItems == nil {
		var ret SkuDataRelationshipsStockItems
		return ret
	}
	return *o.StockItems
}

// GetStockItemsOk returns a tuple with the StockItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockLocationDataRelationships) GetStockItemsOk() (*SkuDataRelationshipsStockItems, bool) {
	if o == nil || o.StockItems == nil {
		return nil, false
	}
	return o.StockItems, true
}

// HasStockItems returns a boolean if a field has been set.
func (o *StockLocationDataRelationships) HasStockItems() bool {
	if o != nil && o.StockItems != nil {
		return true
	}

	return false
}

// SetStockItems gets a reference to the given SkuDataRelationshipsStockItems and assigns it to the StockItems field.
func (o *StockLocationDataRelationships) SetStockItems(v SkuDataRelationshipsStockItems) {
	o.StockItems = &v
}

// GetStockTransfers returns the StockTransfers field value if set, zero value otherwise.
func (o *StockLocationDataRelationships) GetStockTransfers() LineItemDataRelationshipsStockTransfers {
	if o == nil || o.StockTransfers == nil {
		var ret LineItemDataRelationshipsStockTransfers
		return ret
	}
	return *o.StockTransfers
}

// GetStockTransfersOk returns a tuple with the StockTransfers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockLocationDataRelationships) GetStockTransfersOk() (*LineItemDataRelationshipsStockTransfers, bool) {
	if o == nil || o.StockTransfers == nil {
		return nil, false
	}
	return o.StockTransfers, true
}

// HasStockTransfers returns a boolean if a field has been set.
func (o *StockLocationDataRelationships) HasStockTransfers() bool {
	if o != nil && o.StockTransfers != nil {
		return true
	}

	return false
}

// SetStockTransfers gets a reference to the given LineItemDataRelationshipsStockTransfers and assigns it to the StockTransfers field.
func (o *StockLocationDataRelationships) SetStockTransfers(v LineItemDataRelationshipsStockTransfers) {
	o.StockTransfers = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *StockLocationDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret AvalaraAccountDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockLocationDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *StockLocationDataRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *StockLocationDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments) {
	o.Attachments = &v
}

func (o StockLocationDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.InventoryStockLocations != nil {
		toSerialize["inventory_stock_locations"] = o.InventoryStockLocations
	}
	if o.InventoryReturnLocations != nil {
		toSerialize["inventory_return_locations"] = o.InventoryReturnLocations
	}
	if o.StockItems != nil {
		toSerialize["stock_items"] = o.StockItems
	}
	if o.StockTransfers != nil {
		toSerialize["stock_transfers"] = o.StockTransfers
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableStockLocationDataRelationships struct {
	value *StockLocationDataRelationships
	isSet bool
}

func (v NullableStockLocationDataRelationships) Get() *StockLocationDataRelationships {
	return v.value
}

func (v *NullableStockLocationDataRelationships) Set(val *StockLocationDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableStockLocationDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableStockLocationDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockLocationDataRelationships(val *StockLocationDataRelationships) *NullableStockLocationDataRelationships {
	return &NullableStockLocationDataRelationships{value: val, isSet: true}
}

func (v NullableStockLocationDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockLocationDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
