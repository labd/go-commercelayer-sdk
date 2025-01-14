/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.6.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTLineItems201ResponseDataRelationshipsStockLineItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTLineItems201ResponseDataRelationshipsStockLineItems{}

// POSTLineItems201ResponseDataRelationshipsStockLineItems struct for POSTLineItems201ResponseDataRelationshipsStockLineItems
type POSTLineItems201ResponseDataRelationshipsStockLineItems struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks      `json:"links,omitempty"`
	Data  *POSTLineItems201ResponseDataRelationshipsStockLineItemsData `json:"data,omitempty"`
}

// NewPOSTLineItems201ResponseDataRelationshipsStockLineItems instantiates a new POSTLineItems201ResponseDataRelationshipsStockLineItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTLineItems201ResponseDataRelationshipsStockLineItems() *POSTLineItems201ResponseDataRelationshipsStockLineItems {
	this := POSTLineItems201ResponseDataRelationshipsStockLineItems{}
	return &this
}

// NewPOSTLineItems201ResponseDataRelationshipsStockLineItemsWithDefaults instantiates a new POSTLineItems201ResponseDataRelationshipsStockLineItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTLineItems201ResponseDataRelationshipsStockLineItemsWithDefaults() *POSTLineItems201ResponseDataRelationshipsStockLineItems {
	this := POSTLineItems201ResponseDataRelationshipsStockLineItems{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataRelationshipsStockLineItems) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataRelationshipsStockLineItems) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationshipsStockLineItems) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTLineItems201ResponseDataRelationshipsStockLineItems) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataRelationshipsStockLineItems) GetData() POSTLineItems201ResponseDataRelationshipsStockLineItemsData {
	if o == nil || IsNil(o.Data) {
		var ret POSTLineItems201ResponseDataRelationshipsStockLineItemsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataRelationshipsStockLineItems) GetDataOk() (*POSTLineItems201ResponseDataRelationshipsStockLineItemsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationshipsStockLineItems) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTLineItems201ResponseDataRelationshipsStockLineItemsData and assigns it to the Data field.
func (o *POSTLineItems201ResponseDataRelationshipsStockLineItems) SetData(v POSTLineItems201ResponseDataRelationshipsStockLineItemsData) {
	o.Data = &v
}

func (o POSTLineItems201ResponseDataRelationshipsStockLineItems) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTLineItems201ResponseDataRelationshipsStockLineItems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTLineItems201ResponseDataRelationshipsStockLineItems struct {
	value *POSTLineItems201ResponseDataRelationshipsStockLineItems
	isSet bool
}

func (v NullablePOSTLineItems201ResponseDataRelationshipsStockLineItems) Get() *POSTLineItems201ResponseDataRelationshipsStockLineItems {
	return v.value
}

func (v *NullablePOSTLineItems201ResponseDataRelationshipsStockLineItems) Set(val *POSTLineItems201ResponseDataRelationshipsStockLineItems) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTLineItems201ResponseDataRelationshipsStockLineItems) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTLineItems201ResponseDataRelationshipsStockLineItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTLineItems201ResponseDataRelationshipsStockLineItems(val *POSTLineItems201ResponseDataRelationshipsStockLineItems) *NullablePOSTLineItems201ResponseDataRelationshipsStockLineItems {
	return &NullablePOSTLineItems201ResponseDataRelationshipsStockLineItems{value: val, isSet: true}
}

func (v NullablePOSTLineItems201ResponseDataRelationshipsStockLineItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTLineItems201ResponseDataRelationshipsStockLineItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
