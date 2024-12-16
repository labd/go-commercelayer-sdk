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

// checks if the POSTShipments201ResponseDataRelationshipsDeliveryLeadTime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTShipments201ResponseDataRelationshipsDeliveryLeadTime{}

// POSTShipments201ResponseDataRelationshipsDeliveryLeadTime struct for POSTShipments201ResponseDataRelationshipsDeliveryLeadTime
type POSTShipments201ResponseDataRelationshipsDeliveryLeadTime struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks        `json:"links,omitempty"`
	Data  *POSTShipments201ResponseDataRelationshipsDeliveryLeadTimeData `json:"data,omitempty"`
}

// NewPOSTShipments201ResponseDataRelationshipsDeliveryLeadTime instantiates a new POSTShipments201ResponseDataRelationshipsDeliveryLeadTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTShipments201ResponseDataRelationshipsDeliveryLeadTime() *POSTShipments201ResponseDataRelationshipsDeliveryLeadTime {
	this := POSTShipments201ResponseDataRelationshipsDeliveryLeadTime{}
	return &this
}

// NewPOSTShipments201ResponseDataRelationshipsDeliveryLeadTimeWithDefaults instantiates a new POSTShipments201ResponseDataRelationshipsDeliveryLeadTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTShipments201ResponseDataRelationshipsDeliveryLeadTimeWithDefaults() *POSTShipments201ResponseDataRelationshipsDeliveryLeadTime {
	this := POSTShipments201ResponseDataRelationshipsDeliveryLeadTime{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTShipments201ResponseDataRelationshipsDeliveryLeadTime) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShipments201ResponseDataRelationshipsDeliveryLeadTime) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTShipments201ResponseDataRelationshipsDeliveryLeadTime) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTShipments201ResponseDataRelationshipsDeliveryLeadTime) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTShipments201ResponseDataRelationshipsDeliveryLeadTime) GetData() POSTShipments201ResponseDataRelationshipsDeliveryLeadTimeData {
	if o == nil || IsNil(o.Data) {
		var ret POSTShipments201ResponseDataRelationshipsDeliveryLeadTimeData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShipments201ResponseDataRelationshipsDeliveryLeadTime) GetDataOk() (*POSTShipments201ResponseDataRelationshipsDeliveryLeadTimeData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTShipments201ResponseDataRelationshipsDeliveryLeadTime) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTShipments201ResponseDataRelationshipsDeliveryLeadTimeData and assigns it to the Data field.
func (o *POSTShipments201ResponseDataRelationshipsDeliveryLeadTime) SetData(v POSTShipments201ResponseDataRelationshipsDeliveryLeadTimeData) {
	o.Data = &v
}

func (o POSTShipments201ResponseDataRelationshipsDeliveryLeadTime) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTShipments201ResponseDataRelationshipsDeliveryLeadTime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTShipments201ResponseDataRelationshipsDeliveryLeadTime struct {
	value *POSTShipments201ResponseDataRelationshipsDeliveryLeadTime
	isSet bool
}

func (v NullablePOSTShipments201ResponseDataRelationshipsDeliveryLeadTime) Get() *POSTShipments201ResponseDataRelationshipsDeliveryLeadTime {
	return v.value
}

func (v *NullablePOSTShipments201ResponseDataRelationshipsDeliveryLeadTime) Set(val *POSTShipments201ResponseDataRelationshipsDeliveryLeadTime) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTShipments201ResponseDataRelationshipsDeliveryLeadTime) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTShipments201ResponseDataRelationshipsDeliveryLeadTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTShipments201ResponseDataRelationshipsDeliveryLeadTime(val *POSTShipments201ResponseDataRelationshipsDeliveryLeadTime) *NullablePOSTShipments201ResponseDataRelationshipsDeliveryLeadTime {
	return &NullablePOSTShipments201ResponseDataRelationshipsDeliveryLeadTime{value: val, isSet: true}
}

func (v NullablePOSTShipments201ResponseDataRelationshipsDeliveryLeadTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTShipments201ResponseDataRelationshipsDeliveryLeadTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
