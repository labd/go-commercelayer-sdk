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

// GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems struct for GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems
type GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks         `json:"links,omitempty"`
	Data  *GETLineItems200ResponseDataInnerRelationshipsShipmentLineItemsData `json:"data,omitempty"`
}

// NewGETLineItems200ResponseDataInnerRelationshipsShipmentLineItems instantiates a new GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETLineItems200ResponseDataInnerRelationshipsShipmentLineItems() *GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems {
	this := GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems{}
	return &this
}

// NewGETLineItems200ResponseDataInnerRelationshipsShipmentLineItemsWithDefaults instantiates a new GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETLineItems200ResponseDataInnerRelationshipsShipmentLineItemsWithDefaults() *GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems {
	this := GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems) GetData() GETLineItems200ResponseDataInnerRelationshipsShipmentLineItemsData {
	if o == nil || o.Data == nil {
		var ret GETLineItems200ResponseDataInnerRelationshipsShipmentLineItemsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems) GetDataOk() (*GETLineItems200ResponseDataInnerRelationshipsShipmentLineItemsData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETLineItems200ResponseDataInnerRelationshipsShipmentLineItemsData and assigns it to the Data field.
func (o *GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems) SetData(v GETLineItems200ResponseDataInnerRelationshipsShipmentLineItemsData) {
	o.Data = &v
}

func (o GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETLineItems200ResponseDataInnerRelationshipsShipmentLineItems struct {
	value *GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems
	isSet bool
}

func (v NullableGETLineItems200ResponseDataInnerRelationshipsShipmentLineItems) Get() *GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems {
	return v.value
}

func (v *NullableGETLineItems200ResponseDataInnerRelationshipsShipmentLineItems) Set(val *GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems) {
	v.value = val
	v.isSet = true
}

func (v NullableGETLineItems200ResponseDataInnerRelationshipsShipmentLineItems) IsSet() bool {
	return v.isSet
}

func (v *NullableGETLineItems200ResponseDataInnerRelationshipsShipmentLineItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETLineItems200ResponseDataInnerRelationshipsShipmentLineItems(val *GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems) *NullableGETLineItems200ResponseDataInnerRelationshipsShipmentLineItems {
	return &NullableGETLineItems200ResponseDataInnerRelationshipsShipmentLineItems{value: val, isSet: true}
}

func (v NullableGETLineItems200ResponseDataInnerRelationshipsShipmentLineItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETLineItems200ResponseDataInnerRelationshipsShipmentLineItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
