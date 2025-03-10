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

// checks if the StoreUpdateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StoreUpdateDataRelationships{}

// StoreUpdateDataRelationships struct for StoreUpdateDataRelationships
type StoreUpdateDataRelationships struct {
	Market        *BundleCreateDataRelationshipsMarket                  `json:"market,omitempty"`
	Merchant      *MarketCreateDataRelationshipsMerchant                `json:"merchant,omitempty"`
	StockLocation *DeliveryLeadTimeCreateDataRelationshipsStockLocation `json:"stock_location,omitempty"`
}

// NewStoreUpdateDataRelationships instantiates a new StoreUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreUpdateDataRelationships() *StoreUpdateDataRelationships {
	this := StoreUpdateDataRelationships{}
	return &this
}

// NewStoreUpdateDataRelationshipsWithDefaults instantiates a new StoreUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreUpdateDataRelationshipsWithDefaults() *StoreUpdateDataRelationships {
	this := StoreUpdateDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *StoreUpdateDataRelationships) GetMarket() BundleCreateDataRelationshipsMarket {
	if o == nil || IsNil(o.Market) {
		var ret BundleCreateDataRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreUpdateDataRelationships) GetMarketOk() (*BundleCreateDataRelationshipsMarket, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *StoreUpdateDataRelationships) HasMarket() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given BundleCreateDataRelationshipsMarket and assigns it to the Market field.
func (o *StoreUpdateDataRelationships) SetMarket(v BundleCreateDataRelationshipsMarket) {
	o.Market = &v
}

// GetMerchant returns the Merchant field value if set, zero value otherwise.
func (o *StoreUpdateDataRelationships) GetMerchant() MarketCreateDataRelationshipsMerchant {
	if o == nil || IsNil(o.Merchant) {
		var ret MarketCreateDataRelationshipsMerchant
		return ret
	}
	return *o.Merchant
}

// GetMerchantOk returns a tuple with the Merchant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreUpdateDataRelationships) GetMerchantOk() (*MarketCreateDataRelationshipsMerchant, bool) {
	if o == nil || IsNil(o.Merchant) {
		return nil, false
	}
	return o.Merchant, true
}

// HasMerchant returns a boolean if a field has been set.
func (o *StoreUpdateDataRelationships) HasMerchant() bool {
	if o != nil && !IsNil(o.Merchant) {
		return true
	}

	return false
}

// SetMerchant gets a reference to the given MarketCreateDataRelationshipsMerchant and assigns it to the Merchant field.
func (o *StoreUpdateDataRelationships) SetMerchant(v MarketCreateDataRelationshipsMerchant) {
	o.Merchant = &v
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *StoreUpdateDataRelationships) GetStockLocation() DeliveryLeadTimeCreateDataRelationshipsStockLocation {
	if o == nil || IsNil(o.StockLocation) {
		var ret DeliveryLeadTimeCreateDataRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreUpdateDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeCreateDataRelationshipsStockLocation, bool) {
	if o == nil || IsNil(o.StockLocation) {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *StoreUpdateDataRelationships) HasStockLocation() bool {
	if o != nil && !IsNil(o.StockLocation) {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given DeliveryLeadTimeCreateDataRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *StoreUpdateDataRelationships) SetStockLocation(v DeliveryLeadTimeCreateDataRelationshipsStockLocation) {
	o.StockLocation = &v
}

func (o StoreUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoreUpdateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	if !IsNil(o.Merchant) {
		toSerialize["merchant"] = o.Merchant
	}
	if !IsNil(o.StockLocation) {
		toSerialize["stock_location"] = o.StockLocation
	}
	return toSerialize, nil
}

type NullableStoreUpdateDataRelationships struct {
	value *StoreUpdateDataRelationships
	isSet bool
}

func (v NullableStoreUpdateDataRelationships) Get() *StoreUpdateDataRelationships {
	return v.value
}

func (v *NullableStoreUpdateDataRelationships) Set(val *StoreUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreUpdateDataRelationships(val *StoreUpdateDataRelationships) *NullableStoreUpdateDataRelationships {
	return &NullableStoreUpdateDataRelationships{value: val, isSet: true}
}

func (v NullableStoreUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
