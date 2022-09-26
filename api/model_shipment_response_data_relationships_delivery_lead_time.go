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

// ShipmentResponseDataRelationshipsDeliveryLeadTime struct for ShipmentResponseDataRelationshipsDeliveryLeadTime
type ShipmentResponseDataRelationshipsDeliveryLeadTime struct {
	Links *AddressResponseDataRelationshipsGeocoderLinks         `json:"links,omitempty"`
	Data  *ShipmentResponseDataRelationshipsDeliveryLeadTimeData `json:"data,omitempty"`
}

// NewShipmentResponseDataRelationshipsDeliveryLeadTime instantiates a new ShipmentResponseDataRelationshipsDeliveryLeadTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentResponseDataRelationshipsDeliveryLeadTime() *ShipmentResponseDataRelationshipsDeliveryLeadTime {
	this := ShipmentResponseDataRelationshipsDeliveryLeadTime{}
	return &this
}

// NewShipmentResponseDataRelationshipsDeliveryLeadTimeWithDefaults instantiates a new ShipmentResponseDataRelationshipsDeliveryLeadTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentResponseDataRelationshipsDeliveryLeadTimeWithDefaults() *ShipmentResponseDataRelationshipsDeliveryLeadTime {
	this := ShipmentResponseDataRelationshipsDeliveryLeadTime{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ShipmentResponseDataRelationshipsDeliveryLeadTime) GetLinks() AddressResponseDataRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret AddressResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentResponseDataRelationshipsDeliveryLeadTime) GetLinksOk() (*AddressResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ShipmentResponseDataRelationshipsDeliveryLeadTime) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AddressResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *ShipmentResponseDataRelationshipsDeliveryLeadTime) SetLinks(v AddressResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ShipmentResponseDataRelationshipsDeliveryLeadTime) GetData() ShipmentResponseDataRelationshipsDeliveryLeadTimeData {
	if o == nil || o.Data == nil {
		var ret ShipmentResponseDataRelationshipsDeliveryLeadTimeData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentResponseDataRelationshipsDeliveryLeadTime) GetDataOk() (*ShipmentResponseDataRelationshipsDeliveryLeadTimeData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ShipmentResponseDataRelationshipsDeliveryLeadTime) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ShipmentResponseDataRelationshipsDeliveryLeadTimeData and assigns it to the Data field.
func (o *ShipmentResponseDataRelationshipsDeliveryLeadTime) SetData(v ShipmentResponseDataRelationshipsDeliveryLeadTimeData) {
	o.Data = &v
}

func (o ShipmentResponseDataRelationshipsDeliveryLeadTime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableShipmentResponseDataRelationshipsDeliveryLeadTime struct {
	value *ShipmentResponseDataRelationshipsDeliveryLeadTime
	isSet bool
}

func (v NullableShipmentResponseDataRelationshipsDeliveryLeadTime) Get() *ShipmentResponseDataRelationshipsDeliveryLeadTime {
	return v.value
}

func (v *NullableShipmentResponseDataRelationshipsDeliveryLeadTime) Set(val *ShipmentResponseDataRelationshipsDeliveryLeadTime) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentResponseDataRelationshipsDeliveryLeadTime) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentResponseDataRelationshipsDeliveryLeadTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentResponseDataRelationshipsDeliveryLeadTime(val *ShipmentResponseDataRelationshipsDeliveryLeadTime) *NullableShipmentResponseDataRelationshipsDeliveryLeadTime {
	return &NullableShipmentResponseDataRelationshipsDeliveryLeadTime{value: val, isSet: true}
}

func (v NullableShipmentResponseDataRelationshipsDeliveryLeadTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentResponseDataRelationshipsDeliveryLeadTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
