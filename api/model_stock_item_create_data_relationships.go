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

// checks if the StockItemCreateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StockItemCreateDataRelationships{}

// StockItemCreateDataRelationships struct for StockItemCreateDataRelationships
type StockItemCreateDataRelationships struct {
	StockLocation DeliveryLeadTimeCreateDataRelationshipsStockLocation `json:"stock_location"`
	Sku           *InStockSubscriptionCreateDataRelationshipsSku       `json:"sku,omitempty"`
}

// NewStockItemCreateDataRelationships instantiates a new StockItemCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockItemCreateDataRelationships(stockLocation DeliveryLeadTimeCreateDataRelationshipsStockLocation) *StockItemCreateDataRelationships {
	this := StockItemCreateDataRelationships{}
	this.StockLocation = stockLocation
	return &this
}

// NewStockItemCreateDataRelationshipsWithDefaults instantiates a new StockItemCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockItemCreateDataRelationshipsWithDefaults() *StockItemCreateDataRelationships {
	this := StockItemCreateDataRelationships{}
	return &this
}

// GetStockLocation returns the StockLocation field value
func (o *StockItemCreateDataRelationships) GetStockLocation() DeliveryLeadTimeCreateDataRelationshipsStockLocation {
	if o == nil {
		var ret DeliveryLeadTimeCreateDataRelationshipsStockLocation
		return ret
	}

	return o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value
// and a boolean to check if the value has been set.
func (o *StockItemCreateDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeCreateDataRelationshipsStockLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StockLocation, true
}

// SetStockLocation sets field value
func (o *StockItemCreateDataRelationships) SetStockLocation(v DeliveryLeadTimeCreateDataRelationshipsStockLocation) {
	o.StockLocation = v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *StockItemCreateDataRelationships) GetSku() InStockSubscriptionCreateDataRelationshipsSku {
	if o == nil || IsNil(o.Sku) {
		var ret InStockSubscriptionCreateDataRelationshipsSku
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockItemCreateDataRelationships) GetSkuOk() (*InStockSubscriptionCreateDataRelationshipsSku, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *StockItemCreateDataRelationships) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given InStockSubscriptionCreateDataRelationshipsSku and assigns it to the Sku field.
func (o *StockItemCreateDataRelationships) SetSku(v InStockSubscriptionCreateDataRelationshipsSku) {
	o.Sku = &v
}

func (o StockItemCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StockItemCreateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["stock_location"] = o.StockLocation
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	return toSerialize, nil
}

type NullableStockItemCreateDataRelationships struct {
	value *StockItemCreateDataRelationships
	isSet bool
}

func (v NullableStockItemCreateDataRelationships) Get() *StockItemCreateDataRelationships {
	return v.value
}

func (v *NullableStockItemCreateDataRelationships) Set(val *StockItemCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableStockItemCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableStockItemCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockItemCreateDataRelationships(val *StockItemCreateDataRelationships) *NullableStockItemCreateDataRelationships {
	return &NullableStockItemCreateDataRelationships{value: val, isSet: true}
}

func (v NullableStockItemCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockItemCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
