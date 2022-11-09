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

// GETMarkets200ResponseDataInnerRelationshipsMerchant struct for GETMarkets200ResponseDataInnerRelationshipsMerchant
type GETMarkets200ResponseDataInnerRelationshipsMerchant struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *GETMarkets200ResponseDataInnerRelationshipsMerchantData    `json:"data,omitempty"`
}

// NewGETMarkets200ResponseDataInnerRelationshipsMerchant instantiates a new GETMarkets200ResponseDataInnerRelationshipsMerchant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETMarkets200ResponseDataInnerRelationshipsMerchant() *GETMarkets200ResponseDataInnerRelationshipsMerchant {
	this := GETMarkets200ResponseDataInnerRelationshipsMerchant{}
	return &this
}

// NewGETMarkets200ResponseDataInnerRelationshipsMerchantWithDefaults instantiates a new GETMarkets200ResponseDataInnerRelationshipsMerchant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETMarkets200ResponseDataInnerRelationshipsMerchantWithDefaults() *GETMarkets200ResponseDataInnerRelationshipsMerchant {
	this := GETMarkets200ResponseDataInnerRelationshipsMerchant{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerRelationshipsMerchant) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerRelationshipsMerchant) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerRelationshipsMerchant) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETMarkets200ResponseDataInnerRelationshipsMerchant) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerRelationshipsMerchant) GetData() GETMarkets200ResponseDataInnerRelationshipsMerchantData {
	if o == nil || o.Data == nil {
		var ret GETMarkets200ResponseDataInnerRelationshipsMerchantData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerRelationshipsMerchant) GetDataOk() (*GETMarkets200ResponseDataInnerRelationshipsMerchantData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerRelationshipsMerchant) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETMarkets200ResponseDataInnerRelationshipsMerchantData and assigns it to the Data field.
func (o *GETMarkets200ResponseDataInnerRelationshipsMerchant) SetData(v GETMarkets200ResponseDataInnerRelationshipsMerchantData) {
	o.Data = &v
}

func (o GETMarkets200ResponseDataInnerRelationshipsMerchant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETMarkets200ResponseDataInnerRelationshipsMerchant struct {
	value *GETMarkets200ResponseDataInnerRelationshipsMerchant
	isSet bool
}

func (v NullableGETMarkets200ResponseDataInnerRelationshipsMerchant) Get() *GETMarkets200ResponseDataInnerRelationshipsMerchant {
	return v.value
}

func (v *NullableGETMarkets200ResponseDataInnerRelationshipsMerchant) Set(val *GETMarkets200ResponseDataInnerRelationshipsMerchant) {
	v.value = val
	v.isSet = true
}

func (v NullableGETMarkets200ResponseDataInnerRelationshipsMerchant) IsSet() bool {
	return v.isSet
}

func (v *NullableGETMarkets200ResponseDataInnerRelationshipsMerchant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETMarkets200ResponseDataInnerRelationshipsMerchant(val *GETMarkets200ResponseDataInnerRelationshipsMerchant) *NullableGETMarkets200ResponseDataInnerRelationshipsMerchant {
	return &NullableGETMarkets200ResponseDataInnerRelationshipsMerchant{value: val, isSet: true}
}

func (v NullableGETMarkets200ResponseDataInnerRelationshipsMerchant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETMarkets200ResponseDataInnerRelationshipsMerchant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}