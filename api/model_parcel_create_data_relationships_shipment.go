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

// ParcelCreateDataRelationshipsShipment struct for ParcelCreateDataRelationshipsShipment
type ParcelCreateDataRelationshipsShipment struct {
	Data OrderDataRelationshipsShipmentsData `json:"data"`
}

// NewParcelCreateDataRelationshipsShipment instantiates a new ParcelCreateDataRelationshipsShipment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParcelCreateDataRelationshipsShipment(data OrderDataRelationshipsShipmentsData) *ParcelCreateDataRelationshipsShipment {
	this := ParcelCreateDataRelationshipsShipment{}
	this.Data = data
	return &this
}

// NewParcelCreateDataRelationshipsShipmentWithDefaults instantiates a new ParcelCreateDataRelationshipsShipment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParcelCreateDataRelationshipsShipmentWithDefaults() *ParcelCreateDataRelationshipsShipment {
	this := ParcelCreateDataRelationshipsShipment{}
	return &this
}

// GetData returns the Data field value
func (o *ParcelCreateDataRelationshipsShipment) GetData() OrderDataRelationshipsShipmentsData {
	if o == nil {
		var ret OrderDataRelationshipsShipmentsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ParcelCreateDataRelationshipsShipment) GetDataOk() (*OrderDataRelationshipsShipmentsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ParcelCreateDataRelationshipsShipment) SetData(v OrderDataRelationshipsShipmentsData) {
	o.Data = v
}

func (o ParcelCreateDataRelationshipsShipment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableParcelCreateDataRelationshipsShipment struct {
	value *ParcelCreateDataRelationshipsShipment
	isSet bool
}

func (v NullableParcelCreateDataRelationshipsShipment) Get() *ParcelCreateDataRelationshipsShipment {
	return v.value
}

func (v *NullableParcelCreateDataRelationshipsShipment) Set(val *ParcelCreateDataRelationshipsShipment) {
	v.value = val
	v.isSet = true
}

func (v NullableParcelCreateDataRelationshipsShipment) IsSet() bool {
	return v.isSet
}

func (v *NullableParcelCreateDataRelationshipsShipment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParcelCreateDataRelationshipsShipment(val *ParcelCreateDataRelationshipsShipment) *NullableParcelCreateDataRelationshipsShipment {
	return &NullableParcelCreateDataRelationshipsShipment{value: val, isSet: true}
}

func (v NullableParcelCreateDataRelationshipsShipment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParcelCreateDataRelationshipsShipment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
