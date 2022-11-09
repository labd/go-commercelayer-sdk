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

// GETInStockSubscriptions200ResponseDataInnerRelationshipsSku struct for GETInStockSubscriptions200ResponseDataInnerRelationshipsSku
type GETInStockSubscriptions200ResponseDataInnerRelationshipsSku struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks      `json:"links,omitempty"`
	Data  *GETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData `json:"data,omitempty"`
}

// NewGETInStockSubscriptions200ResponseDataInnerRelationshipsSku instantiates a new GETInStockSubscriptions200ResponseDataInnerRelationshipsSku object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETInStockSubscriptions200ResponseDataInnerRelationshipsSku() *GETInStockSubscriptions200ResponseDataInnerRelationshipsSku {
	this := GETInStockSubscriptions200ResponseDataInnerRelationshipsSku{}
	return &this
}

// NewGETInStockSubscriptions200ResponseDataInnerRelationshipsSkuWithDefaults instantiates a new GETInStockSubscriptions200ResponseDataInnerRelationshipsSku object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETInStockSubscriptions200ResponseDataInnerRelationshipsSkuWithDefaults() *GETInStockSubscriptions200ResponseDataInnerRelationshipsSku {
	this := GETInStockSubscriptions200ResponseDataInnerRelationshipsSku{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationshipsSku) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationshipsSku) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationshipsSku) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationshipsSku) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationshipsSku) GetData() GETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData {
	if o == nil || o.Data == nil {
		var ret GETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationshipsSku) GetDataOk() (*GETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationshipsSku) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData and assigns it to the Data field.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationshipsSku) SetData(v GETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData) {
	o.Data = &v
}

func (o GETInStockSubscriptions200ResponseDataInnerRelationshipsSku) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETInStockSubscriptions200ResponseDataInnerRelationshipsSku struct {
	value *GETInStockSubscriptions200ResponseDataInnerRelationshipsSku
	isSet bool
}

func (v NullableGETInStockSubscriptions200ResponseDataInnerRelationshipsSku) Get() *GETInStockSubscriptions200ResponseDataInnerRelationshipsSku {
	return v.value
}

func (v *NullableGETInStockSubscriptions200ResponseDataInnerRelationshipsSku) Set(val *GETInStockSubscriptions200ResponseDataInnerRelationshipsSku) {
	v.value = val
	v.isSet = true
}

func (v NullableGETInStockSubscriptions200ResponseDataInnerRelationshipsSku) IsSet() bool {
	return v.isSet
}

func (v *NullableGETInStockSubscriptions200ResponseDataInnerRelationshipsSku) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETInStockSubscriptions200ResponseDataInnerRelationshipsSku(val *GETInStockSubscriptions200ResponseDataInnerRelationshipsSku) *NullableGETInStockSubscriptions200ResponseDataInnerRelationshipsSku {
	return &NullableGETInStockSubscriptions200ResponseDataInnerRelationshipsSku{value: val, isSet: true}
}

func (v NullableGETInStockSubscriptions200ResponseDataInnerRelationshipsSku) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETInStockSubscriptions200ResponseDataInnerRelationshipsSku) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}