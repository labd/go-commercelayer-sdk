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

// checks if the ShippingMethodCreateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingMethodCreateDataRelationships{}

// ShippingMethodCreateDataRelationships struct for ShippingMethodCreateDataRelationships
type ShippingMethodCreateDataRelationships struct {
	Market              *BundleCreateDataRelationshipsMarket                      `json:"market,omitempty"`
	ShippingZone        *ShippingMethodCreateDataRelationshipsShippingZone        `json:"shipping_zone,omitempty"`
	ShippingCategory    *ShipmentCreateDataRelationshipsShippingCategory          `json:"shipping_category,omitempty"`
	StockLocation       *DeliveryLeadTimeCreateDataRelationshipsStockLocation     `json:"stock_location,omitempty"`
	ShippingMethodTiers *ShippingMethodCreateDataRelationshipsShippingMethodTiers `json:"shipping_method_tiers,omitempty"`
}

// NewShippingMethodCreateDataRelationships instantiates a new ShippingMethodCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingMethodCreateDataRelationships() *ShippingMethodCreateDataRelationships {
	this := ShippingMethodCreateDataRelationships{}
	return &this
}

// NewShippingMethodCreateDataRelationshipsWithDefaults instantiates a new ShippingMethodCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingMethodCreateDataRelationshipsWithDefaults() *ShippingMethodCreateDataRelationships {
	this := ShippingMethodCreateDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *ShippingMethodCreateDataRelationships) GetMarket() BundleCreateDataRelationshipsMarket {
	if o == nil || IsNil(o.Market) {
		var ret BundleCreateDataRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodCreateDataRelationships) GetMarketOk() (*BundleCreateDataRelationshipsMarket, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *ShippingMethodCreateDataRelationships) HasMarket() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given BundleCreateDataRelationshipsMarket and assigns it to the Market field.
func (o *ShippingMethodCreateDataRelationships) SetMarket(v BundleCreateDataRelationshipsMarket) {
	o.Market = &v
}

// GetShippingZone returns the ShippingZone field value if set, zero value otherwise.
func (o *ShippingMethodCreateDataRelationships) GetShippingZone() ShippingMethodCreateDataRelationshipsShippingZone {
	if o == nil || IsNil(o.ShippingZone) {
		var ret ShippingMethodCreateDataRelationshipsShippingZone
		return ret
	}
	return *o.ShippingZone
}

// GetShippingZoneOk returns a tuple with the ShippingZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodCreateDataRelationships) GetShippingZoneOk() (*ShippingMethodCreateDataRelationshipsShippingZone, bool) {
	if o == nil || IsNil(o.ShippingZone) {
		return nil, false
	}
	return o.ShippingZone, true
}

// HasShippingZone returns a boolean if a field has been set.
func (o *ShippingMethodCreateDataRelationships) HasShippingZone() bool {
	if o != nil && !IsNil(o.ShippingZone) {
		return true
	}

	return false
}

// SetShippingZone gets a reference to the given ShippingMethodCreateDataRelationshipsShippingZone and assigns it to the ShippingZone field.
func (o *ShippingMethodCreateDataRelationships) SetShippingZone(v ShippingMethodCreateDataRelationshipsShippingZone) {
	o.ShippingZone = &v
}

// GetShippingCategory returns the ShippingCategory field value if set, zero value otherwise.
func (o *ShippingMethodCreateDataRelationships) GetShippingCategory() ShipmentCreateDataRelationshipsShippingCategory {
	if o == nil || IsNil(o.ShippingCategory) {
		var ret ShipmentCreateDataRelationshipsShippingCategory
		return ret
	}
	return *o.ShippingCategory
}

// GetShippingCategoryOk returns a tuple with the ShippingCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodCreateDataRelationships) GetShippingCategoryOk() (*ShipmentCreateDataRelationshipsShippingCategory, bool) {
	if o == nil || IsNil(o.ShippingCategory) {
		return nil, false
	}
	return o.ShippingCategory, true
}

// HasShippingCategory returns a boolean if a field has been set.
func (o *ShippingMethodCreateDataRelationships) HasShippingCategory() bool {
	if o != nil && !IsNil(o.ShippingCategory) {
		return true
	}

	return false
}

// SetShippingCategory gets a reference to the given ShipmentCreateDataRelationshipsShippingCategory and assigns it to the ShippingCategory field.
func (o *ShippingMethodCreateDataRelationships) SetShippingCategory(v ShipmentCreateDataRelationshipsShippingCategory) {
	o.ShippingCategory = &v
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *ShippingMethodCreateDataRelationships) GetStockLocation() DeliveryLeadTimeCreateDataRelationshipsStockLocation {
	if o == nil || IsNil(o.StockLocation) {
		var ret DeliveryLeadTimeCreateDataRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodCreateDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeCreateDataRelationshipsStockLocation, bool) {
	if o == nil || IsNil(o.StockLocation) {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *ShippingMethodCreateDataRelationships) HasStockLocation() bool {
	if o != nil && !IsNil(o.StockLocation) {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given DeliveryLeadTimeCreateDataRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *ShippingMethodCreateDataRelationships) SetStockLocation(v DeliveryLeadTimeCreateDataRelationshipsStockLocation) {
	o.StockLocation = &v
}

// GetShippingMethodTiers returns the ShippingMethodTiers field value if set, zero value otherwise.
func (o *ShippingMethodCreateDataRelationships) GetShippingMethodTiers() ShippingMethodCreateDataRelationshipsShippingMethodTiers {
	if o == nil || IsNil(o.ShippingMethodTiers) {
		var ret ShippingMethodCreateDataRelationshipsShippingMethodTiers
		return ret
	}
	return *o.ShippingMethodTiers
}

// GetShippingMethodTiersOk returns a tuple with the ShippingMethodTiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodCreateDataRelationships) GetShippingMethodTiersOk() (*ShippingMethodCreateDataRelationshipsShippingMethodTiers, bool) {
	if o == nil || IsNil(o.ShippingMethodTiers) {
		return nil, false
	}
	return o.ShippingMethodTiers, true
}

// HasShippingMethodTiers returns a boolean if a field has been set.
func (o *ShippingMethodCreateDataRelationships) HasShippingMethodTiers() bool {
	if o != nil && !IsNil(o.ShippingMethodTiers) {
		return true
	}

	return false
}

// SetShippingMethodTiers gets a reference to the given ShippingMethodCreateDataRelationshipsShippingMethodTiers and assigns it to the ShippingMethodTiers field.
func (o *ShippingMethodCreateDataRelationships) SetShippingMethodTiers(v ShippingMethodCreateDataRelationshipsShippingMethodTiers) {
	o.ShippingMethodTiers = &v
}

func (o ShippingMethodCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShippingMethodCreateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	if !IsNil(o.ShippingZone) {
		toSerialize["shipping_zone"] = o.ShippingZone
	}
	if !IsNil(o.ShippingCategory) {
		toSerialize["shipping_category"] = o.ShippingCategory
	}
	if !IsNil(o.StockLocation) {
		toSerialize["stock_location"] = o.StockLocation
	}
	if !IsNil(o.ShippingMethodTiers) {
		toSerialize["shipping_method_tiers"] = o.ShippingMethodTiers
	}
	return toSerialize, nil
}

type NullableShippingMethodCreateDataRelationships struct {
	value *ShippingMethodCreateDataRelationships
	isSet bool
}

func (v NullableShippingMethodCreateDataRelationships) Get() *ShippingMethodCreateDataRelationships {
	return v.value
}

func (v *NullableShippingMethodCreateDataRelationships) Set(val *ShippingMethodCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingMethodCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingMethodCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingMethodCreateDataRelationships(val *ShippingMethodCreateDataRelationships) *NullableShippingMethodCreateDataRelationships {
	return &NullableShippingMethodCreateDataRelationships{value: val, isSet: true}
}

func (v NullableShippingMethodCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingMethodCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
