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

// GETMarkets200ResponseDataInnerRelationshipsPriceList struct for GETMarkets200ResponseDataInnerRelationshipsPriceList
type GETMarkets200ResponseDataInnerRelationshipsPriceList struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *GETMarkets200ResponseDataInnerRelationshipsPriceListData   `json:"data,omitempty"`
}

// NewGETMarkets200ResponseDataInnerRelationshipsPriceList instantiates a new GETMarkets200ResponseDataInnerRelationshipsPriceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETMarkets200ResponseDataInnerRelationshipsPriceList() *GETMarkets200ResponseDataInnerRelationshipsPriceList {
	this := GETMarkets200ResponseDataInnerRelationshipsPriceList{}
	return &this
}

// NewGETMarkets200ResponseDataInnerRelationshipsPriceListWithDefaults instantiates a new GETMarkets200ResponseDataInnerRelationshipsPriceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETMarkets200ResponseDataInnerRelationshipsPriceListWithDefaults() *GETMarkets200ResponseDataInnerRelationshipsPriceList {
	this := GETMarkets200ResponseDataInnerRelationshipsPriceList{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerRelationshipsPriceList) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerRelationshipsPriceList) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerRelationshipsPriceList) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETMarkets200ResponseDataInnerRelationshipsPriceList) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerRelationshipsPriceList) GetData() GETMarkets200ResponseDataInnerRelationshipsPriceListData {
	if o == nil || o.Data == nil {
		var ret GETMarkets200ResponseDataInnerRelationshipsPriceListData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerRelationshipsPriceList) GetDataOk() (*GETMarkets200ResponseDataInnerRelationshipsPriceListData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerRelationshipsPriceList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETMarkets200ResponseDataInnerRelationshipsPriceListData and assigns it to the Data field.
func (o *GETMarkets200ResponseDataInnerRelationshipsPriceList) SetData(v GETMarkets200ResponseDataInnerRelationshipsPriceListData) {
	o.Data = &v
}

func (o GETMarkets200ResponseDataInnerRelationshipsPriceList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETMarkets200ResponseDataInnerRelationshipsPriceList struct {
	value *GETMarkets200ResponseDataInnerRelationshipsPriceList
	isSet bool
}

func (v NullableGETMarkets200ResponseDataInnerRelationshipsPriceList) Get() *GETMarkets200ResponseDataInnerRelationshipsPriceList {
	return v.value
}

func (v *NullableGETMarkets200ResponseDataInnerRelationshipsPriceList) Set(val *GETMarkets200ResponseDataInnerRelationshipsPriceList) {
	v.value = val
	v.isSet = true
}

func (v NullableGETMarkets200ResponseDataInnerRelationshipsPriceList) IsSet() bool {
	return v.isSet
}

func (v *NullableGETMarkets200ResponseDataInnerRelationshipsPriceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETMarkets200ResponseDataInnerRelationshipsPriceList(val *GETMarkets200ResponseDataInnerRelationshipsPriceList) *NullableGETMarkets200ResponseDataInnerRelationshipsPriceList {
	return &NullableGETMarkets200ResponseDataInnerRelationshipsPriceList{value: val, isSet: true}
}

func (v NullableGETMarkets200ResponseDataInnerRelationshipsPriceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETMarkets200ResponseDataInnerRelationshipsPriceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}