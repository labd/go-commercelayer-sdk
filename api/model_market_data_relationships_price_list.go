/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// MarketDataRelationshipsPriceList struct for MarketDataRelationshipsPriceList
type MarketDataRelationshipsPriceList struct {
	Data *MarketDataRelationshipsPriceListData `json:"data,omitempty"`
}

// NewMarketDataRelationshipsPriceList instantiates a new MarketDataRelationshipsPriceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketDataRelationshipsPriceList() *MarketDataRelationshipsPriceList {
	this := MarketDataRelationshipsPriceList{}
	return &this
}

// NewMarketDataRelationshipsPriceListWithDefaults instantiates a new MarketDataRelationshipsPriceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketDataRelationshipsPriceListWithDefaults() *MarketDataRelationshipsPriceList {
	this := MarketDataRelationshipsPriceList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *MarketDataRelationshipsPriceList) GetData() MarketDataRelationshipsPriceListData {
	if o == nil || o.Data == nil {
		var ret MarketDataRelationshipsPriceListData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketDataRelationshipsPriceList) GetDataOk() (*MarketDataRelationshipsPriceListData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *MarketDataRelationshipsPriceList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given MarketDataRelationshipsPriceListData and assigns it to the Data field.
func (o *MarketDataRelationshipsPriceList) SetData(v MarketDataRelationshipsPriceListData) {
	o.Data = &v
}

func (o MarketDataRelationshipsPriceList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableMarketDataRelationshipsPriceList struct {
	value *MarketDataRelationshipsPriceList
	isSet bool
}

func (v NullableMarketDataRelationshipsPriceList) Get() *MarketDataRelationshipsPriceList {
	return v.value
}

func (v *NullableMarketDataRelationshipsPriceList) Set(val *MarketDataRelationshipsPriceList) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketDataRelationshipsPriceList) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketDataRelationshipsPriceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketDataRelationshipsPriceList(val *MarketDataRelationshipsPriceList) *NullableMarketDataRelationshipsPriceList {
	return &NullableMarketDataRelationshipsPriceList{value: val, isSet: true}
}

func (v NullableMarketDataRelationshipsPriceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketDataRelationshipsPriceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
