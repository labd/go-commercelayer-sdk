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

// GETSkus200ResponseDataInnerRelationshipsStockItems struct for GETSkus200ResponseDataInnerRelationshipsStockItems
type GETSkus200ResponseDataInnerRelationshipsStockItems struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks   `json:"links,omitempty"`
	Data  []GETSkus200ResponseDataInnerRelationshipsStockItemsDataInner `json:"data,omitempty"`
}

// NewGETSkus200ResponseDataInnerRelationshipsStockItems instantiates a new GETSkus200ResponseDataInnerRelationshipsStockItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETSkus200ResponseDataInnerRelationshipsStockItems() *GETSkus200ResponseDataInnerRelationshipsStockItems {
	this := GETSkus200ResponseDataInnerRelationshipsStockItems{}
	return &this
}

// NewGETSkus200ResponseDataInnerRelationshipsStockItemsWithDefaults instantiates a new GETSkus200ResponseDataInnerRelationshipsStockItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETSkus200ResponseDataInnerRelationshipsStockItemsWithDefaults() *GETSkus200ResponseDataInnerRelationshipsStockItems {
	this := GETSkus200ResponseDataInnerRelationshipsStockItems{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerRelationshipsStockItems) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerRelationshipsStockItems) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerRelationshipsStockItems) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETSkus200ResponseDataInnerRelationshipsStockItems) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerRelationshipsStockItems) GetData() []GETSkus200ResponseDataInnerRelationshipsStockItemsDataInner {
	if o == nil || o.Data == nil {
		var ret []GETSkus200ResponseDataInnerRelationshipsStockItemsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerRelationshipsStockItems) GetDataOk() ([]GETSkus200ResponseDataInnerRelationshipsStockItemsDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerRelationshipsStockItems) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETSkus200ResponseDataInnerRelationshipsStockItemsDataInner and assigns it to the Data field.
func (o *GETSkus200ResponseDataInnerRelationshipsStockItems) SetData(v []GETSkus200ResponseDataInnerRelationshipsStockItemsDataInner) {
	o.Data = v
}

func (o GETSkus200ResponseDataInnerRelationshipsStockItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETSkus200ResponseDataInnerRelationshipsStockItems struct {
	value *GETSkus200ResponseDataInnerRelationshipsStockItems
	isSet bool
}

func (v NullableGETSkus200ResponseDataInnerRelationshipsStockItems) Get() *GETSkus200ResponseDataInnerRelationshipsStockItems {
	return v.value
}

func (v *NullableGETSkus200ResponseDataInnerRelationshipsStockItems) Set(val *GETSkus200ResponseDataInnerRelationshipsStockItems) {
	v.value = val
	v.isSet = true
}

func (v NullableGETSkus200ResponseDataInnerRelationshipsStockItems) IsSet() bool {
	return v.isSet
}

func (v *NullableGETSkus200ResponseDataInnerRelationshipsStockItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETSkus200ResponseDataInnerRelationshipsStockItems(val *GETSkus200ResponseDataInnerRelationshipsStockItems) *NullableGETSkus200ResponseDataInnerRelationshipsStockItems {
	return &NullableGETSkus200ResponseDataInnerRelationshipsStockItems{value: val, isSet: true}
}

func (v NullableGETSkus200ResponseDataInnerRelationshipsStockItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETSkus200ResponseDataInnerRelationshipsStockItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
