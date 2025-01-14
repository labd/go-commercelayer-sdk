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

// checks if the StockLocationDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StockLocationDataRelationships{}

// StockLocationDataRelationships struct for StockLocationDataRelationships
type StockLocationDataRelationships struct {
	Address                  *BingGeocoderDataRelationshipsAddresses                  `json:"address,omitempty"`
	InventoryStockLocations  *InventoryModelDataRelationshipsInventoryStockLocations  `json:"inventory_stock_locations,omitempty"`
	InventoryReturnLocations *InventoryModelDataRelationshipsInventoryReturnLocations `json:"inventory_return_locations,omitempty"`
	StockItems               *ReservedStockDataRelationshipsStockItem                 `json:"stock_items,omitempty"`
	StockTransfers           *LineItemDataRelationshipsStockTransfers                 `json:"stock_transfers,omitempty"`
	Stores                   *MarketDataRelationshipsStores                           `json:"stores,omitempty"`
	Attachments              *AuthorizationDataRelationshipsAttachments               `json:"attachments,omitempty"`
	Versions                 *AddressDataRelationshipsVersions                        `json:"versions,omitempty"`
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
	if o == nil || IsNil(o.Address) {
		var ret BingGeocoderDataRelationshipsAddresses
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockLocationDataRelationships) GetAddressOk() (*BingGeocoderDataRelationshipsAddresses, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *StockLocationDataRelationships) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
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
	if o == nil || IsNil(o.InventoryStockLocations) {
		var ret InventoryModelDataRelationshipsInventoryStockLocations
		return ret
	}
	return *o.InventoryStockLocations
}

// GetInventoryStockLocationsOk returns a tuple with the InventoryStockLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockLocationDataRelationships) GetInventoryStockLocationsOk() (*InventoryModelDataRelationshipsInventoryStockLocations, bool) {
	if o == nil || IsNil(o.InventoryStockLocations) {
		return nil, false
	}
	return o.InventoryStockLocations, true
}

// HasInventoryStockLocations returns a boolean if a field has been set.
func (o *StockLocationDataRelationships) HasInventoryStockLocations() bool {
	if o != nil && !IsNil(o.InventoryStockLocations) {
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
	if o == nil || IsNil(o.InventoryReturnLocations) {
		var ret InventoryModelDataRelationshipsInventoryReturnLocations
		return ret
	}
	return *o.InventoryReturnLocations
}

// GetInventoryReturnLocationsOk returns a tuple with the InventoryReturnLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockLocationDataRelationships) GetInventoryReturnLocationsOk() (*InventoryModelDataRelationshipsInventoryReturnLocations, bool) {
	if o == nil || IsNil(o.InventoryReturnLocations) {
		return nil, false
	}
	return o.InventoryReturnLocations, true
}

// HasInventoryReturnLocations returns a boolean if a field has been set.
func (o *StockLocationDataRelationships) HasInventoryReturnLocations() bool {
	if o != nil && !IsNil(o.InventoryReturnLocations) {
		return true
	}

	return false
}

// SetInventoryReturnLocations gets a reference to the given InventoryModelDataRelationshipsInventoryReturnLocations and assigns it to the InventoryReturnLocations field.
func (o *StockLocationDataRelationships) SetInventoryReturnLocations(v InventoryModelDataRelationshipsInventoryReturnLocations) {
	o.InventoryReturnLocations = &v
}

// GetStockItems returns the StockItems field value if set, zero value otherwise.
func (o *StockLocationDataRelationships) GetStockItems() ReservedStockDataRelationshipsStockItem {
	if o == nil || IsNil(o.StockItems) {
		var ret ReservedStockDataRelationshipsStockItem
		return ret
	}
	return *o.StockItems
}

// GetStockItemsOk returns a tuple with the StockItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockLocationDataRelationships) GetStockItemsOk() (*ReservedStockDataRelationshipsStockItem, bool) {
	if o == nil || IsNil(o.StockItems) {
		return nil, false
	}
	return o.StockItems, true
}

// HasStockItems returns a boolean if a field has been set.
func (o *StockLocationDataRelationships) HasStockItems() bool {
	if o != nil && !IsNil(o.StockItems) {
		return true
	}

	return false
}

// SetStockItems gets a reference to the given ReservedStockDataRelationshipsStockItem and assigns it to the StockItems field.
func (o *StockLocationDataRelationships) SetStockItems(v ReservedStockDataRelationshipsStockItem) {
	o.StockItems = &v
}

// GetStockTransfers returns the StockTransfers field value if set, zero value otherwise.
func (o *StockLocationDataRelationships) GetStockTransfers() LineItemDataRelationshipsStockTransfers {
	if o == nil || IsNil(o.StockTransfers) {
		var ret LineItemDataRelationshipsStockTransfers
		return ret
	}
	return *o.StockTransfers
}

// GetStockTransfersOk returns a tuple with the StockTransfers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockLocationDataRelationships) GetStockTransfersOk() (*LineItemDataRelationshipsStockTransfers, bool) {
	if o == nil || IsNil(o.StockTransfers) {
		return nil, false
	}
	return o.StockTransfers, true
}

// HasStockTransfers returns a boolean if a field has been set.
func (o *StockLocationDataRelationships) HasStockTransfers() bool {
	if o != nil && !IsNil(o.StockTransfers) {
		return true
	}

	return false
}

// SetStockTransfers gets a reference to the given LineItemDataRelationshipsStockTransfers and assigns it to the StockTransfers field.
func (o *StockLocationDataRelationships) SetStockTransfers(v LineItemDataRelationshipsStockTransfers) {
	o.StockTransfers = &v
}

// GetStores returns the Stores field value if set, zero value otherwise.
func (o *StockLocationDataRelationships) GetStores() MarketDataRelationshipsStores {
	if o == nil || IsNil(o.Stores) {
		var ret MarketDataRelationshipsStores
		return ret
	}
	return *o.Stores
}

// GetStoresOk returns a tuple with the Stores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockLocationDataRelationships) GetStoresOk() (*MarketDataRelationshipsStores, bool) {
	if o == nil || IsNil(o.Stores) {
		return nil, false
	}
	return o.Stores, true
}

// HasStores returns a boolean if a field has been set.
func (o *StockLocationDataRelationships) HasStores() bool {
	if o != nil && !IsNil(o.Stores) {
		return true
	}

	return false
}

// SetStores gets a reference to the given MarketDataRelationshipsStores and assigns it to the Stores field.
func (o *StockLocationDataRelationships) SetStores(v MarketDataRelationshipsStores) {
	o.Stores = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *StockLocationDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret AuthorizationDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockLocationDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *StockLocationDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AuthorizationDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *StockLocationDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *StockLocationDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockLocationDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *StockLocationDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *StockLocationDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

func (o StockLocationDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StockLocationDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.InventoryStockLocations) {
		toSerialize["inventory_stock_locations"] = o.InventoryStockLocations
	}
	if !IsNil(o.InventoryReturnLocations) {
		toSerialize["inventory_return_locations"] = o.InventoryReturnLocations
	}
	if !IsNil(o.StockItems) {
		toSerialize["stock_items"] = o.StockItems
	}
	if !IsNil(o.StockTransfers) {
		toSerialize["stock_transfers"] = o.StockTransfers
	}
	if !IsNil(o.Stores) {
		toSerialize["stores"] = o.Stores
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
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
