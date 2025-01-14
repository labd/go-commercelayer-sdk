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

// checks if the POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment{}

// POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment struct for POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment
type POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks                         `json:"links,omitempty"`
	Data  *POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipmentData `json:"data,omitempty"`
}

// NewPOSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment instantiates a new POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment() *POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment {
	this := POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment{}
	return &this
}

// NewPOSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipmentWithDefaults instantiates a new POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipmentWithDefaults() *POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment {
	this := POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment) GetData() POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipmentData {
	if o == nil || IsNil(o.Data) {
		var ret POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipmentData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment) GetDataOk() (*POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipmentData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipmentData and assigns it to the Data field.
func (o *POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment) SetData(v POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipmentData) {
	o.Data = &v
}

func (o POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment struct {
	value *POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment
	isSet bool
}

func (v NullablePOSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment) Get() *POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment {
	return v.value
}

func (v *NullablePOSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment) Set(val *POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment(val *POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment) *NullablePOSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment {
	return &NullablePOSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment{value: val, isSet: true}
}

func (v NullablePOSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
