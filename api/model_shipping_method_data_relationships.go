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

// ShippingMethodDataRelationships struct for ShippingMethodDataRelationships
type ShippingMethodDataRelationships struct {
	Market                      *AvalaraAccountDataRelationshipsMarkets             `json:"market,omitempty"`
	ShippingZone                *ShippingMethodDataRelationshipsShippingZone        `json:"shipping_zone,omitempty"`
	ShippingCategory            *ShipmentDataRelationshipsShippingCategory          `json:"shipping_category,omitempty"`
	StockLocation               *DeliveryLeadTimeDataRelationshipsStockLocation     `json:"stock_location,omitempty"`
	DeliveryLeadTimeForShipment *ShipmentDataRelationshipsDeliveryLeadTime          `json:"delivery_lead_time_for_shipment,omitempty"`
	ShippingMethodTiers         *ShippingMethodDataRelationshipsShippingMethodTiers `json:"shipping_method_tiers,omitempty"`
	ShippingWeightTiers         *ShippingMethodDataRelationshipsShippingWeightTiers `json:"shipping_weight_tiers,omitempty"`
	Attachments                 *AvalaraAccountDataRelationshipsAttachments         `json:"attachments,omitempty"`
}

// NewShippingMethodDataRelationships instantiates a new ShippingMethodDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingMethodDataRelationships() *ShippingMethodDataRelationships {
	this := ShippingMethodDataRelationships{}
	return &this
}

// NewShippingMethodDataRelationshipsWithDefaults instantiates a new ShippingMethodDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingMethodDataRelationshipsWithDefaults() *ShippingMethodDataRelationships {
	this := ShippingMethodDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *ShippingMethodDataRelationships) GetMarket() AvalaraAccountDataRelationshipsMarkets {
	if o == nil || o.Market == nil {
		var ret AvalaraAccountDataRelationshipsMarkets
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodDataRelationships) GetMarketOk() (*AvalaraAccountDataRelationshipsMarkets, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *ShippingMethodDataRelationships) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given AvalaraAccountDataRelationshipsMarkets and assigns it to the Market field.
func (o *ShippingMethodDataRelationships) SetMarket(v AvalaraAccountDataRelationshipsMarkets) {
	o.Market = &v
}

// GetShippingZone returns the ShippingZone field value if set, zero value otherwise.
func (o *ShippingMethodDataRelationships) GetShippingZone() ShippingMethodDataRelationshipsShippingZone {
	if o == nil || o.ShippingZone == nil {
		var ret ShippingMethodDataRelationshipsShippingZone
		return ret
	}
	return *o.ShippingZone
}

// GetShippingZoneOk returns a tuple with the ShippingZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodDataRelationships) GetShippingZoneOk() (*ShippingMethodDataRelationshipsShippingZone, bool) {
	if o == nil || o.ShippingZone == nil {
		return nil, false
	}
	return o.ShippingZone, true
}

// HasShippingZone returns a boolean if a field has been set.
func (o *ShippingMethodDataRelationships) HasShippingZone() bool {
	if o != nil && o.ShippingZone != nil {
		return true
	}

	return false
}

// SetShippingZone gets a reference to the given ShippingMethodDataRelationshipsShippingZone and assigns it to the ShippingZone field.
func (o *ShippingMethodDataRelationships) SetShippingZone(v ShippingMethodDataRelationshipsShippingZone) {
	o.ShippingZone = &v
}

// GetShippingCategory returns the ShippingCategory field value if set, zero value otherwise.
func (o *ShippingMethodDataRelationships) GetShippingCategory() ShipmentDataRelationshipsShippingCategory {
	if o == nil || o.ShippingCategory == nil {
		var ret ShipmentDataRelationshipsShippingCategory
		return ret
	}
	return *o.ShippingCategory
}

// GetShippingCategoryOk returns a tuple with the ShippingCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodDataRelationships) GetShippingCategoryOk() (*ShipmentDataRelationshipsShippingCategory, bool) {
	if o == nil || o.ShippingCategory == nil {
		return nil, false
	}
	return o.ShippingCategory, true
}

// HasShippingCategory returns a boolean if a field has been set.
func (o *ShippingMethodDataRelationships) HasShippingCategory() bool {
	if o != nil && o.ShippingCategory != nil {
		return true
	}

	return false
}

// SetShippingCategory gets a reference to the given ShipmentDataRelationshipsShippingCategory and assigns it to the ShippingCategory field.
func (o *ShippingMethodDataRelationships) SetShippingCategory(v ShipmentDataRelationshipsShippingCategory) {
	o.ShippingCategory = &v
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *ShippingMethodDataRelationships) GetStockLocation() DeliveryLeadTimeDataRelationshipsStockLocation {
	if o == nil || o.StockLocation == nil {
		var ret DeliveryLeadTimeDataRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeDataRelationshipsStockLocation, bool) {
	if o == nil || o.StockLocation == nil {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *ShippingMethodDataRelationships) HasStockLocation() bool {
	if o != nil && o.StockLocation != nil {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given DeliveryLeadTimeDataRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *ShippingMethodDataRelationships) SetStockLocation(v DeliveryLeadTimeDataRelationshipsStockLocation) {
	o.StockLocation = &v
}

// GetDeliveryLeadTimeForShipment returns the DeliveryLeadTimeForShipment field value if set, zero value otherwise.
func (o *ShippingMethodDataRelationships) GetDeliveryLeadTimeForShipment() ShipmentDataRelationshipsDeliveryLeadTime {
	if o == nil || o.DeliveryLeadTimeForShipment == nil {
		var ret ShipmentDataRelationshipsDeliveryLeadTime
		return ret
	}
	return *o.DeliveryLeadTimeForShipment
}

// GetDeliveryLeadTimeForShipmentOk returns a tuple with the DeliveryLeadTimeForShipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodDataRelationships) GetDeliveryLeadTimeForShipmentOk() (*ShipmentDataRelationshipsDeliveryLeadTime, bool) {
	if o == nil || o.DeliveryLeadTimeForShipment == nil {
		return nil, false
	}
	return o.DeliveryLeadTimeForShipment, true
}

// HasDeliveryLeadTimeForShipment returns a boolean if a field has been set.
func (o *ShippingMethodDataRelationships) HasDeliveryLeadTimeForShipment() bool {
	if o != nil && o.DeliveryLeadTimeForShipment != nil {
		return true
	}

	return false
}

// SetDeliveryLeadTimeForShipment gets a reference to the given ShipmentDataRelationshipsDeliveryLeadTime and assigns it to the DeliveryLeadTimeForShipment field.
func (o *ShippingMethodDataRelationships) SetDeliveryLeadTimeForShipment(v ShipmentDataRelationshipsDeliveryLeadTime) {
	o.DeliveryLeadTimeForShipment = &v
}

// GetShippingMethodTiers returns the ShippingMethodTiers field value if set, zero value otherwise.
func (o *ShippingMethodDataRelationships) GetShippingMethodTiers() ShippingMethodDataRelationshipsShippingMethodTiers {
	if o == nil || o.ShippingMethodTiers == nil {
		var ret ShippingMethodDataRelationshipsShippingMethodTiers
		return ret
	}
	return *o.ShippingMethodTiers
}

// GetShippingMethodTiersOk returns a tuple with the ShippingMethodTiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodDataRelationships) GetShippingMethodTiersOk() (*ShippingMethodDataRelationshipsShippingMethodTiers, bool) {
	if o == nil || o.ShippingMethodTiers == nil {
		return nil, false
	}
	return o.ShippingMethodTiers, true
}

// HasShippingMethodTiers returns a boolean if a field has been set.
func (o *ShippingMethodDataRelationships) HasShippingMethodTiers() bool {
	if o != nil && o.ShippingMethodTiers != nil {
		return true
	}

	return false
}

// SetShippingMethodTiers gets a reference to the given ShippingMethodDataRelationshipsShippingMethodTiers and assigns it to the ShippingMethodTiers field.
func (o *ShippingMethodDataRelationships) SetShippingMethodTiers(v ShippingMethodDataRelationshipsShippingMethodTiers) {
	o.ShippingMethodTiers = &v
}

// GetShippingWeightTiers returns the ShippingWeightTiers field value if set, zero value otherwise.
func (o *ShippingMethodDataRelationships) GetShippingWeightTiers() ShippingMethodDataRelationshipsShippingWeightTiers {
	if o == nil || o.ShippingWeightTiers == nil {
		var ret ShippingMethodDataRelationshipsShippingWeightTiers
		return ret
	}
	return *o.ShippingWeightTiers
}

// GetShippingWeightTiersOk returns a tuple with the ShippingWeightTiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodDataRelationships) GetShippingWeightTiersOk() (*ShippingMethodDataRelationshipsShippingWeightTiers, bool) {
	if o == nil || o.ShippingWeightTiers == nil {
		return nil, false
	}
	return o.ShippingWeightTiers, true
}

// HasShippingWeightTiers returns a boolean if a field has been set.
func (o *ShippingMethodDataRelationships) HasShippingWeightTiers() bool {
	if o != nil && o.ShippingWeightTiers != nil {
		return true
	}

	return false
}

// SetShippingWeightTiers gets a reference to the given ShippingMethodDataRelationshipsShippingWeightTiers and assigns it to the ShippingWeightTiers field.
func (o *ShippingMethodDataRelationships) SetShippingWeightTiers(v ShippingMethodDataRelationshipsShippingWeightTiers) {
	o.ShippingWeightTiers = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *ShippingMethodDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret AvalaraAccountDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *ShippingMethodDataRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *ShippingMethodDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments) {
	o.Attachments = &v
}

func (o ShippingMethodDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Market != nil {
		toSerialize["market"] = o.Market
	}
	if o.ShippingZone != nil {
		toSerialize["shipping_zone"] = o.ShippingZone
	}
	if o.ShippingCategory != nil {
		toSerialize["shipping_category"] = o.ShippingCategory
	}
	if o.StockLocation != nil {
		toSerialize["stock_location"] = o.StockLocation
	}
	if o.DeliveryLeadTimeForShipment != nil {
		toSerialize["delivery_lead_time_for_shipment"] = o.DeliveryLeadTimeForShipment
	}
	if o.ShippingMethodTiers != nil {
		toSerialize["shipping_method_tiers"] = o.ShippingMethodTiers
	}
	if o.ShippingWeightTiers != nil {
		toSerialize["shipping_weight_tiers"] = o.ShippingWeightTiers
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableShippingMethodDataRelationships struct {
	value *ShippingMethodDataRelationships
	isSet bool
}

func (v NullableShippingMethodDataRelationships) Get() *ShippingMethodDataRelationships {
	return v.value
}

func (v *NullableShippingMethodDataRelationships) Set(val *ShippingMethodDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingMethodDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingMethodDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingMethodDataRelationships(val *ShippingMethodDataRelationships) *NullableShippingMethodDataRelationships {
	return &NullableShippingMethodDataRelationships{value: val, isSet: true}
}

func (v NullableShippingMethodDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingMethodDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
