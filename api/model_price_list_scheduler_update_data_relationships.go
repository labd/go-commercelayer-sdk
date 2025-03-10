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

// checks if the PriceListSchedulerUpdateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceListSchedulerUpdateDataRelationships{}

// PriceListSchedulerUpdateDataRelationships struct for PriceListSchedulerUpdateDataRelationships
type PriceListSchedulerUpdateDataRelationships struct {
	Market    *BundleCreateDataRelationshipsMarket    `json:"market,omitempty"`
	PriceList *MarketCreateDataRelationshipsPriceList `json:"price_list,omitempty"`
}

// NewPriceListSchedulerUpdateDataRelationships instantiates a new PriceListSchedulerUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceListSchedulerUpdateDataRelationships() *PriceListSchedulerUpdateDataRelationships {
	this := PriceListSchedulerUpdateDataRelationships{}
	return &this
}

// NewPriceListSchedulerUpdateDataRelationshipsWithDefaults instantiates a new PriceListSchedulerUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceListSchedulerUpdateDataRelationshipsWithDefaults() *PriceListSchedulerUpdateDataRelationships {
	this := PriceListSchedulerUpdateDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *PriceListSchedulerUpdateDataRelationships) GetMarket() BundleCreateDataRelationshipsMarket {
	if o == nil || IsNil(o.Market) {
		var ret BundleCreateDataRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceListSchedulerUpdateDataRelationships) GetMarketOk() (*BundleCreateDataRelationshipsMarket, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *PriceListSchedulerUpdateDataRelationships) HasMarket() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given BundleCreateDataRelationshipsMarket and assigns it to the Market field.
func (o *PriceListSchedulerUpdateDataRelationships) SetMarket(v BundleCreateDataRelationshipsMarket) {
	o.Market = &v
}

// GetPriceList returns the PriceList field value if set, zero value otherwise.
func (o *PriceListSchedulerUpdateDataRelationships) GetPriceList() MarketCreateDataRelationshipsPriceList {
	if o == nil || IsNil(o.PriceList) {
		var ret MarketCreateDataRelationshipsPriceList
		return ret
	}
	return *o.PriceList
}

// GetPriceListOk returns a tuple with the PriceList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceListSchedulerUpdateDataRelationships) GetPriceListOk() (*MarketCreateDataRelationshipsPriceList, bool) {
	if o == nil || IsNil(o.PriceList) {
		return nil, false
	}
	return o.PriceList, true
}

// HasPriceList returns a boolean if a field has been set.
func (o *PriceListSchedulerUpdateDataRelationships) HasPriceList() bool {
	if o != nil && !IsNil(o.PriceList) {
		return true
	}

	return false
}

// SetPriceList gets a reference to the given MarketCreateDataRelationshipsPriceList and assigns it to the PriceList field.
func (o *PriceListSchedulerUpdateDataRelationships) SetPriceList(v MarketCreateDataRelationshipsPriceList) {
	o.PriceList = &v
}

func (o PriceListSchedulerUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceListSchedulerUpdateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	if !IsNil(o.PriceList) {
		toSerialize["price_list"] = o.PriceList
	}
	return toSerialize, nil
}

type NullablePriceListSchedulerUpdateDataRelationships struct {
	value *PriceListSchedulerUpdateDataRelationships
	isSet bool
}

func (v NullablePriceListSchedulerUpdateDataRelationships) Get() *PriceListSchedulerUpdateDataRelationships {
	return v.value
}

func (v *NullablePriceListSchedulerUpdateDataRelationships) Set(val *PriceListSchedulerUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceListSchedulerUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceListSchedulerUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceListSchedulerUpdateDataRelationships(val *PriceListSchedulerUpdateDataRelationships) *NullablePriceListSchedulerUpdateDataRelationships {
	return &NullablePriceListSchedulerUpdateDataRelationships{value: val, isSet: true}
}

func (v NullablePriceListSchedulerUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceListSchedulerUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
