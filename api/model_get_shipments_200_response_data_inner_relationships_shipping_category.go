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

// GETShipments200ResponseDataInnerRelationshipsShippingCategory struct for GETShipments200ResponseDataInnerRelationshipsShippingCategory
type GETShipments200ResponseDataInnerRelationshipsShippingCategory struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks        `json:"links,omitempty"`
	Data  *GETShipments200ResponseDataInnerRelationshipsShippingCategoryData `json:"data,omitempty"`
}

// NewGETShipments200ResponseDataInnerRelationshipsShippingCategory instantiates a new GETShipments200ResponseDataInnerRelationshipsShippingCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETShipments200ResponseDataInnerRelationshipsShippingCategory() *GETShipments200ResponseDataInnerRelationshipsShippingCategory {
	this := GETShipments200ResponseDataInnerRelationshipsShippingCategory{}
	return &this
}

// NewGETShipments200ResponseDataInnerRelationshipsShippingCategoryWithDefaults instantiates a new GETShipments200ResponseDataInnerRelationshipsShippingCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETShipments200ResponseDataInnerRelationshipsShippingCategoryWithDefaults() *GETShipments200ResponseDataInnerRelationshipsShippingCategory {
	this := GETShipments200ResponseDataInnerRelationshipsShippingCategory{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerRelationshipsShippingCategory) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerRelationshipsShippingCategory) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerRelationshipsShippingCategory) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETShipments200ResponseDataInnerRelationshipsShippingCategory) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerRelationshipsShippingCategory) GetData() GETShipments200ResponseDataInnerRelationshipsShippingCategoryData {
	if o == nil || o.Data == nil {
		var ret GETShipments200ResponseDataInnerRelationshipsShippingCategoryData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerRelationshipsShippingCategory) GetDataOk() (*GETShipments200ResponseDataInnerRelationshipsShippingCategoryData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerRelationshipsShippingCategory) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETShipments200ResponseDataInnerRelationshipsShippingCategoryData and assigns it to the Data field.
func (o *GETShipments200ResponseDataInnerRelationshipsShippingCategory) SetData(v GETShipments200ResponseDataInnerRelationshipsShippingCategoryData) {
	o.Data = &v
}

func (o GETShipments200ResponseDataInnerRelationshipsShippingCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETShipments200ResponseDataInnerRelationshipsShippingCategory struct {
	value *GETShipments200ResponseDataInnerRelationshipsShippingCategory
	isSet bool
}

func (v NullableGETShipments200ResponseDataInnerRelationshipsShippingCategory) Get() *GETShipments200ResponseDataInnerRelationshipsShippingCategory {
	return v.value
}

func (v *NullableGETShipments200ResponseDataInnerRelationshipsShippingCategory) Set(val *GETShipments200ResponseDataInnerRelationshipsShippingCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableGETShipments200ResponseDataInnerRelationshipsShippingCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableGETShipments200ResponseDataInnerRelationshipsShippingCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETShipments200ResponseDataInnerRelationshipsShippingCategory(val *GETShipments200ResponseDataInnerRelationshipsShippingCategory) *NullableGETShipments200ResponseDataInnerRelationshipsShippingCategory {
	return &NullableGETShipments200ResponseDataInnerRelationshipsShippingCategory{value: val, isSet: true}
}

func (v NullableGETShipments200ResponseDataInnerRelationshipsShippingCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETShipments200ResponseDataInnerRelationshipsShippingCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}