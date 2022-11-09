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

// GETOrders200ResponseDataInnerRelationshipsShipments struct for GETOrders200ResponseDataInnerRelationshipsShipments
type GETOrders200ResponseDataInnerRelationshipsShipments struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks    `json:"links,omitempty"`
	Data  []GETOrders200ResponseDataInnerRelationshipsShipmentsDataInner `json:"data,omitempty"`
}

// NewGETOrders200ResponseDataInnerRelationshipsShipments instantiates a new GETOrders200ResponseDataInnerRelationshipsShipments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrders200ResponseDataInnerRelationshipsShipments() *GETOrders200ResponseDataInnerRelationshipsShipments {
	this := GETOrders200ResponseDataInnerRelationshipsShipments{}
	return &this
}

// NewGETOrders200ResponseDataInnerRelationshipsShipmentsWithDefaults instantiates a new GETOrders200ResponseDataInnerRelationshipsShipments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrders200ResponseDataInnerRelationshipsShipmentsWithDefaults() *GETOrders200ResponseDataInnerRelationshipsShipments {
	this := GETOrders200ResponseDataInnerRelationshipsShipments{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETOrders200ResponseDataInnerRelationshipsShipments) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsShipments) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsShipments) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETOrders200ResponseDataInnerRelationshipsShipments) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETOrders200ResponseDataInnerRelationshipsShipments) GetData() []GETOrders200ResponseDataInnerRelationshipsShipmentsDataInner {
	if o == nil || o.Data == nil {
		var ret []GETOrders200ResponseDataInnerRelationshipsShipmentsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsShipments) GetDataOk() ([]GETOrders200ResponseDataInnerRelationshipsShipmentsDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsShipments) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETOrders200ResponseDataInnerRelationshipsShipmentsDataInner and assigns it to the Data field.
func (o *GETOrders200ResponseDataInnerRelationshipsShipments) SetData(v []GETOrders200ResponseDataInnerRelationshipsShipmentsDataInner) {
	o.Data = v
}

func (o GETOrders200ResponseDataInnerRelationshipsShipments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETOrders200ResponseDataInnerRelationshipsShipments struct {
	value *GETOrders200ResponseDataInnerRelationshipsShipments
	isSet bool
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsShipments) Get() *GETOrders200ResponseDataInnerRelationshipsShipments {
	return v.value
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsShipments) Set(val *GETOrders200ResponseDataInnerRelationshipsShipments) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsShipments) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsShipments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrders200ResponseDataInnerRelationshipsShipments(val *GETOrders200ResponseDataInnerRelationshipsShipments) *NullableGETOrders200ResponseDataInnerRelationshipsShipments {
	return &NullableGETOrders200ResponseDataInnerRelationshipsShipments{value: val, isSet: true}
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsShipments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsShipments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
