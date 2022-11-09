/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.2
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// OrderDataRelationshipsShipments struct for OrderDataRelationshipsShipments
type OrderDataRelationshipsShipments struct {
	Data OrderDataRelationshipsShipmentsData `json:"data"`
}

// NewOrderDataRelationshipsShipments instantiates a new OrderDataRelationshipsShipments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDataRelationshipsShipments(data OrderDataRelationshipsShipmentsData) *OrderDataRelationshipsShipments {
	this := OrderDataRelationshipsShipments{}
	this.Data = data
	return &this
}

// NewOrderDataRelationshipsShipmentsWithDefaults instantiates a new OrderDataRelationshipsShipments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDataRelationshipsShipmentsWithDefaults() *OrderDataRelationshipsShipments {
	this := OrderDataRelationshipsShipments{}
	return &this
}

// GetData returns the Data field value
func (o *OrderDataRelationshipsShipments) GetData() OrderDataRelationshipsShipmentsData {
	if o == nil {
		var ret OrderDataRelationshipsShipmentsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OrderDataRelationshipsShipments) GetDataOk() (*OrderDataRelationshipsShipmentsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *OrderDataRelationshipsShipments) SetData(v OrderDataRelationshipsShipmentsData) {
	o.Data = v
}

func (o OrderDataRelationshipsShipments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableOrderDataRelationshipsShipments struct {
	value *OrderDataRelationshipsShipments
	isSet bool
}

func (v NullableOrderDataRelationshipsShipments) Get() *OrderDataRelationshipsShipments {
	return v.value
}

func (v *NullableOrderDataRelationshipsShipments) Set(val *OrderDataRelationshipsShipments) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDataRelationshipsShipments) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDataRelationshipsShipments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDataRelationshipsShipments(val *OrderDataRelationshipsShipments) *NullableOrderDataRelationshipsShipments {
	return &NullableOrderDataRelationshipsShipments{value: val, isSet: true}
}

func (v NullableOrderDataRelationshipsShipments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDataRelationshipsShipments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
