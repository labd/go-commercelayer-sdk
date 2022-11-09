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

// GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus struct for GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus
type GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks            `json:"links,omitempty"`
	Data  []GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner `json:"data,omitempty"`
}

// NewGETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus instantiates a new GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus() *GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus {
	this := GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus{}
	return &this
}

// NewGETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusWithDefaults instantiates a new GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusWithDefaults() *GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus {
	this := GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus) GetData() []GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner {
	if o == nil || o.Data == nil {
		var ret []GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus) GetDataOk() ([]GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner and assigns it to the Data field.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus) SetData(v []GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner) {
	o.Data = v
}

func (o GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus struct {
	value *GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus
	isSet bool
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus) Get() *GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus {
	return v.value
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus) Set(val *GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus(val *GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus) *NullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus {
	return &NullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus{value: val, isSet: true}
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeSkus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
