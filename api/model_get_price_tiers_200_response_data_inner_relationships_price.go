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

// GETPriceTiers200ResponseDataInnerRelationshipsPrice struct for GETPriceTiers200ResponseDataInnerRelationshipsPrice
type GETPriceTiers200ResponseDataInnerRelationshipsPrice struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *GETPriceTiers200ResponseDataInnerRelationshipsPriceData    `json:"data,omitempty"`
}

// NewGETPriceTiers200ResponseDataInnerRelationshipsPrice instantiates a new GETPriceTiers200ResponseDataInnerRelationshipsPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPriceTiers200ResponseDataInnerRelationshipsPrice() *GETPriceTiers200ResponseDataInnerRelationshipsPrice {
	this := GETPriceTiers200ResponseDataInnerRelationshipsPrice{}
	return &this
}

// NewGETPriceTiers200ResponseDataInnerRelationshipsPriceWithDefaults instantiates a new GETPriceTiers200ResponseDataInnerRelationshipsPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPriceTiers200ResponseDataInnerRelationshipsPriceWithDefaults() *GETPriceTiers200ResponseDataInnerRelationshipsPrice {
	this := GETPriceTiers200ResponseDataInnerRelationshipsPrice{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETPriceTiers200ResponseDataInnerRelationshipsPrice) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPriceTiers200ResponseDataInnerRelationshipsPrice) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETPriceTiers200ResponseDataInnerRelationshipsPrice) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETPriceTiers200ResponseDataInnerRelationshipsPrice) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETPriceTiers200ResponseDataInnerRelationshipsPrice) GetData() GETPriceTiers200ResponseDataInnerRelationshipsPriceData {
	if o == nil || o.Data == nil {
		var ret GETPriceTiers200ResponseDataInnerRelationshipsPriceData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPriceTiers200ResponseDataInnerRelationshipsPrice) GetDataOk() (*GETPriceTiers200ResponseDataInnerRelationshipsPriceData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETPriceTiers200ResponseDataInnerRelationshipsPrice) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETPriceTiers200ResponseDataInnerRelationshipsPriceData and assigns it to the Data field.
func (o *GETPriceTiers200ResponseDataInnerRelationshipsPrice) SetData(v GETPriceTiers200ResponseDataInnerRelationshipsPriceData) {
	o.Data = &v
}

func (o GETPriceTiers200ResponseDataInnerRelationshipsPrice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETPriceTiers200ResponseDataInnerRelationshipsPrice struct {
	value *GETPriceTiers200ResponseDataInnerRelationshipsPrice
	isSet bool
}

func (v NullableGETPriceTiers200ResponseDataInnerRelationshipsPrice) Get() *GETPriceTiers200ResponseDataInnerRelationshipsPrice {
	return v.value
}

func (v *NullableGETPriceTiers200ResponseDataInnerRelationshipsPrice) Set(val *GETPriceTiers200ResponseDataInnerRelationshipsPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPriceTiers200ResponseDataInnerRelationshipsPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPriceTiers200ResponseDataInnerRelationshipsPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPriceTiers200ResponseDataInnerRelationshipsPrice(val *GETPriceTiers200ResponseDataInnerRelationshipsPrice) *NullableGETPriceTiers200ResponseDataInnerRelationshipsPrice {
	return &NullableGETPriceTiers200ResponseDataInnerRelationshipsPrice{value: val, isSet: true}
}

func (v NullableGETPriceTiers200ResponseDataInnerRelationshipsPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPriceTiers200ResponseDataInnerRelationshipsPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
