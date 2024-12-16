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

// checks if the ShipmentDataRelationshipsDeliveryLeadTime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentDataRelationshipsDeliveryLeadTime{}

// ShipmentDataRelationshipsDeliveryLeadTime struct for ShipmentDataRelationshipsDeliveryLeadTime
type ShipmentDataRelationshipsDeliveryLeadTime struct {
	Data *ShipmentDataRelationshipsDeliveryLeadTimeData `json:"data,omitempty"`
}

// NewShipmentDataRelationshipsDeliveryLeadTime instantiates a new ShipmentDataRelationshipsDeliveryLeadTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentDataRelationshipsDeliveryLeadTime() *ShipmentDataRelationshipsDeliveryLeadTime {
	this := ShipmentDataRelationshipsDeliveryLeadTime{}
	return &this
}

// NewShipmentDataRelationshipsDeliveryLeadTimeWithDefaults instantiates a new ShipmentDataRelationshipsDeliveryLeadTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentDataRelationshipsDeliveryLeadTimeWithDefaults() *ShipmentDataRelationshipsDeliveryLeadTime {
	this := ShipmentDataRelationshipsDeliveryLeadTime{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ShipmentDataRelationshipsDeliveryLeadTime) GetData() ShipmentDataRelationshipsDeliveryLeadTimeData {
	if o == nil || IsNil(o.Data) {
		var ret ShipmentDataRelationshipsDeliveryLeadTimeData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDataRelationshipsDeliveryLeadTime) GetDataOk() (*ShipmentDataRelationshipsDeliveryLeadTimeData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ShipmentDataRelationshipsDeliveryLeadTime) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ShipmentDataRelationshipsDeliveryLeadTimeData and assigns it to the Data field.
func (o *ShipmentDataRelationshipsDeliveryLeadTime) SetData(v ShipmentDataRelationshipsDeliveryLeadTimeData) {
	o.Data = &v
}

func (o ShipmentDataRelationshipsDeliveryLeadTime) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipmentDataRelationshipsDeliveryLeadTime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableShipmentDataRelationshipsDeliveryLeadTime struct {
	value *ShipmentDataRelationshipsDeliveryLeadTime
	isSet bool
}

func (v NullableShipmentDataRelationshipsDeliveryLeadTime) Get() *ShipmentDataRelationshipsDeliveryLeadTime {
	return v.value
}

func (v *NullableShipmentDataRelationshipsDeliveryLeadTime) Set(val *ShipmentDataRelationshipsDeliveryLeadTime) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentDataRelationshipsDeliveryLeadTime) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentDataRelationshipsDeliveryLeadTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentDataRelationshipsDeliveryLeadTime(val *ShipmentDataRelationshipsDeliveryLeadTime) *NullableShipmentDataRelationshipsDeliveryLeadTime {
	return &NullableShipmentDataRelationshipsDeliveryLeadTime{value: val, isSet: true}
}

func (v NullableShipmentDataRelationshipsDeliveryLeadTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentDataRelationshipsDeliveryLeadTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
