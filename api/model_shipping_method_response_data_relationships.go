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

// ShippingMethodResponseDataRelationships struct for ShippingMethodResponseDataRelationships
type ShippingMethodResponseDataRelationships struct {
	Market                      *BillingInfoValidationRuleResponseDataRelationshipsMarket           `json:"market,omitempty"`
	ShippingZone                *ShippingMethodResponseDataRelationshipsShippingZone                `json:"shipping_zone,omitempty"`
	ShippingCategory            *ShipmentResponseDataRelationshipsShippingCategory                  `json:"shipping_category,omitempty"`
	StockLocation               *DeliveryLeadTimeResponseDataRelationshipsStockLocation             `json:"stock_location,omitempty"`
	DeliveryLeadTimeForShipment *ShippingMethodResponseDataRelationshipsDeliveryLeadTimeForShipment `json:"delivery_lead_time_for_shipment,omitempty"`
	ShippingMethodTiers         *ShippingMethodResponseDataRelationshipsShippingMethodTiers         `json:"shipping_method_tiers,omitempty"`
	ShippingWeightTiers         *ShippingMethodResponseDataRelationshipsShippingWeightTiers         `json:"shipping_weight_tiers,omitempty"`
	Attachments                 *AvalaraAccountResponseDataRelationshipsAttachments                 `json:"attachments,omitempty"`
}

// NewShippingMethodResponseDataRelationships instantiates a new ShippingMethodResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingMethodResponseDataRelationships() *ShippingMethodResponseDataRelationships {
	this := ShippingMethodResponseDataRelationships{}
	return &this
}

// NewShippingMethodResponseDataRelationshipsWithDefaults instantiates a new ShippingMethodResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingMethodResponseDataRelationshipsWithDefaults() *ShippingMethodResponseDataRelationships {
	this := ShippingMethodResponseDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *ShippingMethodResponseDataRelationships) GetMarket() BillingInfoValidationRuleResponseDataRelationshipsMarket {
	if o == nil || o.Market == nil {
		var ret BillingInfoValidationRuleResponseDataRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodResponseDataRelationships) GetMarketOk() (*BillingInfoValidationRuleResponseDataRelationshipsMarket, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *ShippingMethodResponseDataRelationships) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given BillingInfoValidationRuleResponseDataRelationshipsMarket and assigns it to the Market field.
func (o *ShippingMethodResponseDataRelationships) SetMarket(v BillingInfoValidationRuleResponseDataRelationshipsMarket) {
	o.Market = &v
}

// GetShippingZone returns the ShippingZone field value if set, zero value otherwise.
func (o *ShippingMethodResponseDataRelationships) GetShippingZone() ShippingMethodResponseDataRelationshipsShippingZone {
	if o == nil || o.ShippingZone == nil {
		var ret ShippingMethodResponseDataRelationshipsShippingZone
		return ret
	}
	return *o.ShippingZone
}

// GetShippingZoneOk returns a tuple with the ShippingZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodResponseDataRelationships) GetShippingZoneOk() (*ShippingMethodResponseDataRelationshipsShippingZone, bool) {
	if o == nil || o.ShippingZone == nil {
		return nil, false
	}
	return o.ShippingZone, true
}

// HasShippingZone returns a boolean if a field has been set.
func (o *ShippingMethodResponseDataRelationships) HasShippingZone() bool {
	if o != nil && o.ShippingZone != nil {
		return true
	}

	return false
}

// SetShippingZone gets a reference to the given ShippingMethodResponseDataRelationshipsShippingZone and assigns it to the ShippingZone field.
func (o *ShippingMethodResponseDataRelationships) SetShippingZone(v ShippingMethodResponseDataRelationshipsShippingZone) {
	o.ShippingZone = &v
}

// GetShippingCategory returns the ShippingCategory field value if set, zero value otherwise.
func (o *ShippingMethodResponseDataRelationships) GetShippingCategory() ShipmentResponseDataRelationshipsShippingCategory {
	if o == nil || o.ShippingCategory == nil {
		var ret ShipmentResponseDataRelationshipsShippingCategory
		return ret
	}
	return *o.ShippingCategory
}

// GetShippingCategoryOk returns a tuple with the ShippingCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodResponseDataRelationships) GetShippingCategoryOk() (*ShipmentResponseDataRelationshipsShippingCategory, bool) {
	if o == nil || o.ShippingCategory == nil {
		return nil, false
	}
	return o.ShippingCategory, true
}

// HasShippingCategory returns a boolean if a field has been set.
func (o *ShippingMethodResponseDataRelationships) HasShippingCategory() bool {
	if o != nil && o.ShippingCategory != nil {
		return true
	}

	return false
}

// SetShippingCategory gets a reference to the given ShipmentResponseDataRelationshipsShippingCategory and assigns it to the ShippingCategory field.
func (o *ShippingMethodResponseDataRelationships) SetShippingCategory(v ShipmentResponseDataRelationshipsShippingCategory) {
	o.ShippingCategory = &v
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *ShippingMethodResponseDataRelationships) GetStockLocation() DeliveryLeadTimeResponseDataRelationshipsStockLocation {
	if o == nil || o.StockLocation == nil {
		var ret DeliveryLeadTimeResponseDataRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodResponseDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeResponseDataRelationshipsStockLocation, bool) {
	if o == nil || o.StockLocation == nil {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *ShippingMethodResponseDataRelationships) HasStockLocation() bool {
	if o != nil && o.StockLocation != nil {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given DeliveryLeadTimeResponseDataRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *ShippingMethodResponseDataRelationships) SetStockLocation(v DeliveryLeadTimeResponseDataRelationshipsStockLocation) {
	o.StockLocation = &v
}

// GetDeliveryLeadTimeForShipment returns the DeliveryLeadTimeForShipment field value if set, zero value otherwise.
func (o *ShippingMethodResponseDataRelationships) GetDeliveryLeadTimeForShipment() ShippingMethodResponseDataRelationshipsDeliveryLeadTimeForShipment {
	if o == nil || o.DeliveryLeadTimeForShipment == nil {
		var ret ShippingMethodResponseDataRelationshipsDeliveryLeadTimeForShipment
		return ret
	}
	return *o.DeliveryLeadTimeForShipment
}

// GetDeliveryLeadTimeForShipmentOk returns a tuple with the DeliveryLeadTimeForShipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodResponseDataRelationships) GetDeliveryLeadTimeForShipmentOk() (*ShippingMethodResponseDataRelationshipsDeliveryLeadTimeForShipment, bool) {
	if o == nil || o.DeliveryLeadTimeForShipment == nil {
		return nil, false
	}
	return o.DeliveryLeadTimeForShipment, true
}

// HasDeliveryLeadTimeForShipment returns a boolean if a field has been set.
func (o *ShippingMethodResponseDataRelationships) HasDeliveryLeadTimeForShipment() bool {
	if o != nil && o.DeliveryLeadTimeForShipment != nil {
		return true
	}

	return false
}

// SetDeliveryLeadTimeForShipment gets a reference to the given ShippingMethodResponseDataRelationshipsDeliveryLeadTimeForShipment and assigns it to the DeliveryLeadTimeForShipment field.
func (o *ShippingMethodResponseDataRelationships) SetDeliveryLeadTimeForShipment(v ShippingMethodResponseDataRelationshipsDeliveryLeadTimeForShipment) {
	o.DeliveryLeadTimeForShipment = &v
}

// GetShippingMethodTiers returns the ShippingMethodTiers field value if set, zero value otherwise.
func (o *ShippingMethodResponseDataRelationships) GetShippingMethodTiers() ShippingMethodResponseDataRelationshipsShippingMethodTiers {
	if o == nil || o.ShippingMethodTiers == nil {
		var ret ShippingMethodResponseDataRelationshipsShippingMethodTiers
		return ret
	}
	return *o.ShippingMethodTiers
}

// GetShippingMethodTiersOk returns a tuple with the ShippingMethodTiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodResponseDataRelationships) GetShippingMethodTiersOk() (*ShippingMethodResponseDataRelationshipsShippingMethodTiers, bool) {
	if o == nil || o.ShippingMethodTiers == nil {
		return nil, false
	}
	return o.ShippingMethodTiers, true
}

// HasShippingMethodTiers returns a boolean if a field has been set.
func (o *ShippingMethodResponseDataRelationships) HasShippingMethodTiers() bool {
	if o != nil && o.ShippingMethodTiers != nil {
		return true
	}

	return false
}

// SetShippingMethodTiers gets a reference to the given ShippingMethodResponseDataRelationshipsShippingMethodTiers and assigns it to the ShippingMethodTiers field.
func (o *ShippingMethodResponseDataRelationships) SetShippingMethodTiers(v ShippingMethodResponseDataRelationshipsShippingMethodTiers) {
	o.ShippingMethodTiers = &v
}

// GetShippingWeightTiers returns the ShippingWeightTiers field value if set, zero value otherwise.
func (o *ShippingMethodResponseDataRelationships) GetShippingWeightTiers() ShippingMethodResponseDataRelationshipsShippingWeightTiers {
	if o == nil || o.ShippingWeightTiers == nil {
		var ret ShippingMethodResponseDataRelationshipsShippingWeightTiers
		return ret
	}
	return *o.ShippingWeightTiers
}

// GetShippingWeightTiersOk returns a tuple with the ShippingWeightTiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodResponseDataRelationships) GetShippingWeightTiersOk() (*ShippingMethodResponseDataRelationshipsShippingWeightTiers, bool) {
	if o == nil || o.ShippingWeightTiers == nil {
		return nil, false
	}
	return o.ShippingWeightTiers, true
}

// HasShippingWeightTiers returns a boolean if a field has been set.
func (o *ShippingMethodResponseDataRelationships) HasShippingWeightTiers() bool {
	if o != nil && o.ShippingWeightTiers != nil {
		return true
	}

	return false
}

// SetShippingWeightTiers gets a reference to the given ShippingMethodResponseDataRelationshipsShippingWeightTiers and assigns it to the ShippingWeightTiers field.
func (o *ShippingMethodResponseDataRelationships) SetShippingWeightTiers(v ShippingMethodResponseDataRelationshipsShippingWeightTiers) {
	o.ShippingWeightTiers = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *ShippingMethodResponseDataRelationships) GetAttachments() AvalaraAccountResponseDataRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret AvalaraAccountResponseDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodResponseDataRelationships) GetAttachmentsOk() (*AvalaraAccountResponseDataRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *ShippingMethodResponseDataRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountResponseDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *ShippingMethodResponseDataRelationships) SetAttachments(v AvalaraAccountResponseDataRelationshipsAttachments) {
	o.Attachments = &v
}

func (o ShippingMethodResponseDataRelationships) MarshalJSON() ([]byte, error) {
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

type NullableShippingMethodResponseDataRelationships struct {
	value *ShippingMethodResponseDataRelationships
	isSet bool
}

func (v NullableShippingMethodResponseDataRelationships) Get() *ShippingMethodResponseDataRelationships {
	return v.value
}

func (v *NullableShippingMethodResponseDataRelationships) Set(val *ShippingMethodResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingMethodResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingMethodResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingMethodResponseDataRelationships(val *ShippingMethodResponseDataRelationships) *NullableShippingMethodResponseDataRelationships {
	return &NullableShippingMethodResponseDataRelationships{value: val, isSet: true}
}

func (v NullableShippingMethodResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingMethodResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
