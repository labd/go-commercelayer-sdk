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

// checks if the BundleCreateDataRelationshipsMarket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleCreateDataRelationshipsMarket{}

// BundleCreateDataRelationshipsMarket struct for BundleCreateDataRelationshipsMarket
type BundleCreateDataRelationshipsMarket struct {
	Data AvalaraAccountDataRelationshipsMarketsData `json:"data"`
}

// NewBundleCreateDataRelationshipsMarket instantiates a new BundleCreateDataRelationshipsMarket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleCreateDataRelationshipsMarket(data AvalaraAccountDataRelationshipsMarketsData) *BundleCreateDataRelationshipsMarket {
	this := BundleCreateDataRelationshipsMarket{}
	this.Data = data
	return &this
}

// NewBundleCreateDataRelationshipsMarketWithDefaults instantiates a new BundleCreateDataRelationshipsMarket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleCreateDataRelationshipsMarketWithDefaults() *BundleCreateDataRelationshipsMarket {
	this := BundleCreateDataRelationshipsMarket{}
	return &this
}

// GetData returns the Data field value
func (o *BundleCreateDataRelationshipsMarket) GetData() AvalaraAccountDataRelationshipsMarketsData {
	if o == nil {
		var ret AvalaraAccountDataRelationshipsMarketsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BundleCreateDataRelationshipsMarket) GetDataOk() (*AvalaraAccountDataRelationshipsMarketsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BundleCreateDataRelationshipsMarket) SetData(v AvalaraAccountDataRelationshipsMarketsData) {
	o.Data = v
}

func (o BundleCreateDataRelationshipsMarket) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleCreateDataRelationshipsMarket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableBundleCreateDataRelationshipsMarket struct {
	value *BundleCreateDataRelationshipsMarket
	isSet bool
}

func (v NullableBundleCreateDataRelationshipsMarket) Get() *BundleCreateDataRelationshipsMarket {
	return v.value
}

func (v *NullableBundleCreateDataRelationshipsMarket) Set(val *BundleCreateDataRelationshipsMarket) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleCreateDataRelationshipsMarket) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleCreateDataRelationshipsMarket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleCreateDataRelationshipsMarket(val *BundleCreateDataRelationshipsMarket) *NullableBundleCreateDataRelationshipsMarket {
	return &NullableBundleCreateDataRelationshipsMarket{value: val, isSet: true}
}

func (v NullableBundleCreateDataRelationshipsMarket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleCreateDataRelationshipsMarket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
