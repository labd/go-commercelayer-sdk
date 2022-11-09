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

// GETOrders200ResponseDataInnerRelationshipsShippingAddress struct for GETOrders200ResponseDataInnerRelationshipsShippingAddress
type GETOrders200ResponseDataInnerRelationshipsShippingAddress struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks    `json:"links,omitempty"`
	Data  *GETOrders200ResponseDataInnerRelationshipsShippingAddressData `json:"data,omitempty"`
}

// NewGETOrders200ResponseDataInnerRelationshipsShippingAddress instantiates a new GETOrders200ResponseDataInnerRelationshipsShippingAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrders200ResponseDataInnerRelationshipsShippingAddress() *GETOrders200ResponseDataInnerRelationshipsShippingAddress {
	this := GETOrders200ResponseDataInnerRelationshipsShippingAddress{}
	return &this
}

// NewGETOrders200ResponseDataInnerRelationshipsShippingAddressWithDefaults instantiates a new GETOrders200ResponseDataInnerRelationshipsShippingAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrders200ResponseDataInnerRelationshipsShippingAddressWithDefaults() *GETOrders200ResponseDataInnerRelationshipsShippingAddress {
	this := GETOrders200ResponseDataInnerRelationshipsShippingAddress{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETOrders200ResponseDataInnerRelationshipsShippingAddress) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsShippingAddress) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsShippingAddress) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETOrders200ResponseDataInnerRelationshipsShippingAddress) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETOrders200ResponseDataInnerRelationshipsShippingAddress) GetData() GETOrders200ResponseDataInnerRelationshipsShippingAddressData {
	if o == nil || o.Data == nil {
		var ret GETOrders200ResponseDataInnerRelationshipsShippingAddressData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsShippingAddress) GetDataOk() (*GETOrders200ResponseDataInnerRelationshipsShippingAddressData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsShippingAddress) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETOrders200ResponseDataInnerRelationshipsShippingAddressData and assigns it to the Data field.
func (o *GETOrders200ResponseDataInnerRelationshipsShippingAddress) SetData(v GETOrders200ResponseDataInnerRelationshipsShippingAddressData) {
	o.Data = &v
}

func (o GETOrders200ResponseDataInnerRelationshipsShippingAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETOrders200ResponseDataInnerRelationshipsShippingAddress struct {
	value *GETOrders200ResponseDataInnerRelationshipsShippingAddress
	isSet bool
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsShippingAddress) Get() *GETOrders200ResponseDataInnerRelationshipsShippingAddress {
	return v.value
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsShippingAddress) Set(val *GETOrders200ResponseDataInnerRelationshipsShippingAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsShippingAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsShippingAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrders200ResponseDataInnerRelationshipsShippingAddress(val *GETOrders200ResponseDataInnerRelationshipsShippingAddress) *NullableGETOrders200ResponseDataInnerRelationshipsShippingAddress {
	return &NullableGETOrders200ResponseDataInnerRelationshipsShippingAddress{value: val, isSet: true}
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsShippingAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsShippingAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
