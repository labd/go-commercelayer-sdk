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

// checks if the PriceListSchedulerDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceListSchedulerDataRelationships{}

// PriceListSchedulerDataRelationships struct for PriceListSchedulerDataRelationships
type PriceListSchedulerDataRelationships struct {
	Market    *AvalaraAccountDataRelationshipsMarkets `json:"market,omitempty"`
	PriceList *MarketDataRelationshipsPriceList       `json:"price_list,omitempty"`
	Events    *AddressDataRelationshipsEvents         `json:"events,omitempty"`
	Versions  *AddressDataRelationshipsVersions       `json:"versions,omitempty"`
}

// NewPriceListSchedulerDataRelationships instantiates a new PriceListSchedulerDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceListSchedulerDataRelationships() *PriceListSchedulerDataRelationships {
	this := PriceListSchedulerDataRelationships{}
	return &this
}

// NewPriceListSchedulerDataRelationshipsWithDefaults instantiates a new PriceListSchedulerDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceListSchedulerDataRelationshipsWithDefaults() *PriceListSchedulerDataRelationships {
	this := PriceListSchedulerDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *PriceListSchedulerDataRelationships) GetMarket() AvalaraAccountDataRelationshipsMarkets {
	if o == nil || IsNil(o.Market) {
		var ret AvalaraAccountDataRelationshipsMarkets
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceListSchedulerDataRelationships) GetMarketOk() (*AvalaraAccountDataRelationshipsMarkets, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *PriceListSchedulerDataRelationships) HasMarket() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given AvalaraAccountDataRelationshipsMarkets and assigns it to the Market field.
func (o *PriceListSchedulerDataRelationships) SetMarket(v AvalaraAccountDataRelationshipsMarkets) {
	o.Market = &v
}

// GetPriceList returns the PriceList field value if set, zero value otherwise.
func (o *PriceListSchedulerDataRelationships) GetPriceList() MarketDataRelationshipsPriceList {
	if o == nil || IsNil(o.PriceList) {
		var ret MarketDataRelationshipsPriceList
		return ret
	}
	return *o.PriceList
}

// GetPriceListOk returns a tuple with the PriceList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceListSchedulerDataRelationships) GetPriceListOk() (*MarketDataRelationshipsPriceList, bool) {
	if o == nil || IsNil(o.PriceList) {
		return nil, false
	}
	return o.PriceList, true
}

// HasPriceList returns a boolean if a field has been set.
func (o *PriceListSchedulerDataRelationships) HasPriceList() bool {
	if o != nil && !IsNil(o.PriceList) {
		return true
	}

	return false
}

// SetPriceList gets a reference to the given MarketDataRelationshipsPriceList and assigns it to the PriceList field.
func (o *PriceListSchedulerDataRelationships) SetPriceList(v MarketDataRelationshipsPriceList) {
	o.PriceList = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *PriceListSchedulerDataRelationships) GetEvents() AddressDataRelationshipsEvents {
	if o == nil || IsNil(o.Events) {
		var ret AddressDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceListSchedulerDataRelationships) GetEventsOk() (*AddressDataRelationshipsEvents, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *PriceListSchedulerDataRelationships) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given AddressDataRelationshipsEvents and assigns it to the Events field.
func (o *PriceListSchedulerDataRelationships) SetEvents(v AddressDataRelationshipsEvents) {
	o.Events = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *PriceListSchedulerDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceListSchedulerDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *PriceListSchedulerDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *PriceListSchedulerDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

func (o PriceListSchedulerDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceListSchedulerDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	if !IsNil(o.PriceList) {
		toSerialize["price_list"] = o.PriceList
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullablePriceListSchedulerDataRelationships struct {
	value *PriceListSchedulerDataRelationships
	isSet bool
}

func (v NullablePriceListSchedulerDataRelationships) Get() *PriceListSchedulerDataRelationships {
	return v.value
}

func (v *NullablePriceListSchedulerDataRelationships) Set(val *PriceListSchedulerDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceListSchedulerDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceListSchedulerDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceListSchedulerDataRelationships(val *PriceListSchedulerDataRelationships) *NullablePriceListSchedulerDataRelationships {
	return &NullablePriceListSchedulerDataRelationships{value: val, isSet: true}
}

func (v NullablePriceListSchedulerDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceListSchedulerDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
