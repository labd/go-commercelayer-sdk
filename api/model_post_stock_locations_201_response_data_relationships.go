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

// checks if the POSTStockLocations201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTStockLocations201ResponseDataRelationships{}

// POSTStockLocations201ResponseDataRelationships struct for POSTStockLocations201ResponseDataRelationships
type POSTStockLocations201ResponseDataRelationships struct {
	Address                  *POSTCustomerAddresses201ResponseDataRelationshipsAddress                `json:"address,omitempty"`
	InventoryStockLocations  *POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations  `json:"inventory_stock_locations,omitempty"`
	InventoryReturnLocations *POSTInventoryModels201ResponseDataRelationshipsInventoryReturnLocations `json:"inventory_return_locations,omitempty"`
	StockItems               *POSTSkus201ResponseDataRelationshipsStockItems                          `json:"stock_items,omitempty"`
	StockTransfers           *POSTLineItems201ResponseDataRelationshipsStockTransfers                 `json:"stock_transfers,omitempty"`
	Stores                   *POSTMarkets201ResponseDataRelationshipsStores                           `json:"stores,omitempty"`
	Attachments              *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments `json:"attachments,omitempty"`
	Versions                 *POSTAddresses201ResponseDataRelationshipsVersions                       `json:"versions,omitempty"`
}

// NewPOSTStockLocations201ResponseDataRelationships instantiates a new POSTStockLocations201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTStockLocations201ResponseDataRelationships() *POSTStockLocations201ResponseDataRelationships {
	this := POSTStockLocations201ResponseDataRelationships{}
	return &this
}

// NewPOSTStockLocations201ResponseDataRelationshipsWithDefaults instantiates a new POSTStockLocations201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTStockLocations201ResponseDataRelationshipsWithDefaults() *POSTStockLocations201ResponseDataRelationships {
	this := POSTStockLocations201ResponseDataRelationships{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *POSTStockLocations201ResponseDataRelationships) GetAddress() POSTCustomerAddresses201ResponseDataRelationshipsAddress {
	if o == nil || IsNil(o.Address) {
		var ret POSTCustomerAddresses201ResponseDataRelationshipsAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockLocations201ResponseDataRelationships) GetAddressOk() (*POSTCustomerAddresses201ResponseDataRelationshipsAddress, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *POSTStockLocations201ResponseDataRelationships) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given POSTCustomerAddresses201ResponseDataRelationshipsAddress and assigns it to the Address field.
func (o *POSTStockLocations201ResponseDataRelationships) SetAddress(v POSTCustomerAddresses201ResponseDataRelationshipsAddress) {
	o.Address = &v
}

// GetInventoryStockLocations returns the InventoryStockLocations field value if set, zero value otherwise.
func (o *POSTStockLocations201ResponseDataRelationships) GetInventoryStockLocations() POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations {
	if o == nil || IsNil(o.InventoryStockLocations) {
		var ret POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations
		return ret
	}
	return *o.InventoryStockLocations
}

// GetInventoryStockLocationsOk returns a tuple with the InventoryStockLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockLocations201ResponseDataRelationships) GetInventoryStockLocationsOk() (*POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations, bool) {
	if o == nil || IsNil(o.InventoryStockLocations) {
		return nil, false
	}
	return o.InventoryStockLocations, true
}

// HasInventoryStockLocations returns a boolean if a field has been set.
func (o *POSTStockLocations201ResponseDataRelationships) HasInventoryStockLocations() bool {
	if o != nil && !IsNil(o.InventoryStockLocations) {
		return true
	}

	return false
}

// SetInventoryStockLocations gets a reference to the given POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations and assigns it to the InventoryStockLocations field.
func (o *POSTStockLocations201ResponseDataRelationships) SetInventoryStockLocations(v POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations) {
	o.InventoryStockLocations = &v
}

// GetInventoryReturnLocations returns the InventoryReturnLocations field value if set, zero value otherwise.
func (o *POSTStockLocations201ResponseDataRelationships) GetInventoryReturnLocations() POSTInventoryModels201ResponseDataRelationshipsInventoryReturnLocations {
	if o == nil || IsNil(o.InventoryReturnLocations) {
		var ret POSTInventoryModels201ResponseDataRelationshipsInventoryReturnLocations
		return ret
	}
	return *o.InventoryReturnLocations
}

// GetInventoryReturnLocationsOk returns a tuple with the InventoryReturnLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockLocations201ResponseDataRelationships) GetInventoryReturnLocationsOk() (*POSTInventoryModels201ResponseDataRelationshipsInventoryReturnLocations, bool) {
	if o == nil || IsNil(o.InventoryReturnLocations) {
		return nil, false
	}
	return o.InventoryReturnLocations, true
}

// HasInventoryReturnLocations returns a boolean if a field has been set.
func (o *POSTStockLocations201ResponseDataRelationships) HasInventoryReturnLocations() bool {
	if o != nil && !IsNil(o.InventoryReturnLocations) {
		return true
	}

	return false
}

// SetInventoryReturnLocations gets a reference to the given POSTInventoryModels201ResponseDataRelationshipsInventoryReturnLocations and assigns it to the InventoryReturnLocations field.
func (o *POSTStockLocations201ResponseDataRelationships) SetInventoryReturnLocations(v POSTInventoryModels201ResponseDataRelationshipsInventoryReturnLocations) {
	o.InventoryReturnLocations = &v
}

// GetStockItems returns the StockItems field value if set, zero value otherwise.
func (o *POSTStockLocations201ResponseDataRelationships) GetStockItems() POSTSkus201ResponseDataRelationshipsStockItems {
	if o == nil || IsNil(o.StockItems) {
		var ret POSTSkus201ResponseDataRelationshipsStockItems
		return ret
	}
	return *o.StockItems
}

// GetStockItemsOk returns a tuple with the StockItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockLocations201ResponseDataRelationships) GetStockItemsOk() (*POSTSkus201ResponseDataRelationshipsStockItems, bool) {
	if o == nil || IsNil(o.StockItems) {
		return nil, false
	}
	return o.StockItems, true
}

// HasStockItems returns a boolean if a field has been set.
func (o *POSTStockLocations201ResponseDataRelationships) HasStockItems() bool {
	if o != nil && !IsNil(o.StockItems) {
		return true
	}

	return false
}

// SetStockItems gets a reference to the given POSTSkus201ResponseDataRelationshipsStockItems and assigns it to the StockItems field.
func (o *POSTStockLocations201ResponseDataRelationships) SetStockItems(v POSTSkus201ResponseDataRelationshipsStockItems) {
	o.StockItems = &v
}

// GetStockTransfers returns the StockTransfers field value if set, zero value otherwise.
func (o *POSTStockLocations201ResponseDataRelationships) GetStockTransfers() POSTLineItems201ResponseDataRelationshipsStockTransfers {
	if o == nil || IsNil(o.StockTransfers) {
		var ret POSTLineItems201ResponseDataRelationshipsStockTransfers
		return ret
	}
	return *o.StockTransfers
}

// GetStockTransfersOk returns a tuple with the StockTransfers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockLocations201ResponseDataRelationships) GetStockTransfersOk() (*POSTLineItems201ResponseDataRelationshipsStockTransfers, bool) {
	if o == nil || IsNil(o.StockTransfers) {
		return nil, false
	}
	return o.StockTransfers, true
}

// HasStockTransfers returns a boolean if a field has been set.
func (o *POSTStockLocations201ResponseDataRelationships) HasStockTransfers() bool {
	if o != nil && !IsNil(o.StockTransfers) {
		return true
	}

	return false
}

// SetStockTransfers gets a reference to the given POSTLineItems201ResponseDataRelationshipsStockTransfers and assigns it to the StockTransfers field.
func (o *POSTStockLocations201ResponseDataRelationships) SetStockTransfers(v POSTLineItems201ResponseDataRelationshipsStockTransfers) {
	o.StockTransfers = &v
}

// GetStores returns the Stores field value if set, zero value otherwise.
func (o *POSTStockLocations201ResponseDataRelationships) GetStores() POSTMarkets201ResponseDataRelationshipsStores {
	if o == nil || IsNil(o.Stores) {
		var ret POSTMarkets201ResponseDataRelationshipsStores
		return ret
	}
	return *o.Stores
}

// GetStoresOk returns a tuple with the Stores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockLocations201ResponseDataRelationships) GetStoresOk() (*POSTMarkets201ResponseDataRelationshipsStores, bool) {
	if o == nil || IsNil(o.Stores) {
		return nil, false
	}
	return o.Stores, true
}

// HasStores returns a boolean if a field has been set.
func (o *POSTStockLocations201ResponseDataRelationships) HasStores() bool {
	if o != nil && !IsNil(o.Stores) {
		return true
	}

	return false
}

// SetStores gets a reference to the given POSTMarkets201ResponseDataRelationshipsStores and assigns it to the Stores field.
func (o *POSTStockLocations201ResponseDataRelationships) SetStores(v POSTMarkets201ResponseDataRelationshipsStores) {
	o.Stores = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *POSTStockLocations201ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockLocations201ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *POSTStockLocations201ResponseDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *POSTStockLocations201ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *POSTStockLocations201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockLocations201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *POSTStockLocations201ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *POSTStockLocations201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

func (o POSTStockLocations201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTStockLocations201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
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

type NullablePOSTStockLocations201ResponseDataRelationships struct {
	value *POSTStockLocations201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTStockLocations201ResponseDataRelationships) Get() *POSTStockLocations201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTStockLocations201ResponseDataRelationships) Set(val *POSTStockLocations201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTStockLocations201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTStockLocations201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTStockLocations201ResponseDataRelationships(val *POSTStockLocations201ResponseDataRelationships) *NullablePOSTStockLocations201ResponseDataRelationships {
	return &NullablePOSTStockLocations201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTStockLocations201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTStockLocations201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
