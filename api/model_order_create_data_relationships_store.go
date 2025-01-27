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

// checks if the OrderCreateDataRelationshipsStore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreateDataRelationshipsStore{}

// OrderCreateDataRelationshipsStore struct for OrderCreateDataRelationshipsStore
type OrderCreateDataRelationshipsStore struct {
	Data MarketDataRelationshipsStoresData `json:"data"`
}

// NewOrderCreateDataRelationshipsStore instantiates a new OrderCreateDataRelationshipsStore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreateDataRelationshipsStore(data MarketDataRelationshipsStoresData) *OrderCreateDataRelationshipsStore {
	this := OrderCreateDataRelationshipsStore{}
	this.Data = data
	return &this
}

// NewOrderCreateDataRelationshipsStoreWithDefaults instantiates a new OrderCreateDataRelationshipsStore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreateDataRelationshipsStoreWithDefaults() *OrderCreateDataRelationshipsStore {
	this := OrderCreateDataRelationshipsStore{}
	return &this
}

// GetData returns the Data field value
func (o *OrderCreateDataRelationshipsStore) GetData() MarketDataRelationshipsStoresData {
	if o == nil {
		var ret MarketDataRelationshipsStoresData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OrderCreateDataRelationshipsStore) GetDataOk() (*MarketDataRelationshipsStoresData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *OrderCreateDataRelationshipsStore) SetData(v MarketDataRelationshipsStoresData) {
	o.Data = v
}

func (o OrderCreateDataRelationshipsStore) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreateDataRelationshipsStore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableOrderCreateDataRelationshipsStore struct {
	value *OrderCreateDataRelationshipsStore
	isSet bool
}

func (v NullableOrderCreateDataRelationshipsStore) Get() *OrderCreateDataRelationshipsStore {
	return v.value
}

func (v *NullableOrderCreateDataRelationshipsStore) Set(val *OrderCreateDataRelationshipsStore) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreateDataRelationshipsStore) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreateDataRelationshipsStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreateDataRelationshipsStore(val *OrderCreateDataRelationshipsStore) *NullableOrderCreateDataRelationshipsStore {
	return &NullableOrderCreateDataRelationshipsStore{value: val, isSet: true}
}

func (v NullableOrderCreateDataRelationshipsStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreateDataRelationshipsStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
