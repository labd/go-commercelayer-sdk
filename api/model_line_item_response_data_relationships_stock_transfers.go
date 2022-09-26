/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// LineItemResponseDataRelationshipsStockTransfers struct for LineItemResponseDataRelationshipsStockTransfers
type LineItemResponseDataRelationshipsStockTransfers struct {
	Links *AddressResponseDataRelationshipsGeocoderLinks             `json:"links,omitempty"`
	Data  []LineItemResponseDataRelationshipsStockTransfersDataInner `json:"data,omitempty"`
}

// NewLineItemResponseDataRelationshipsStockTransfers instantiates a new LineItemResponseDataRelationshipsStockTransfers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemResponseDataRelationshipsStockTransfers() *LineItemResponseDataRelationshipsStockTransfers {
	this := LineItemResponseDataRelationshipsStockTransfers{}
	return &this
}

// NewLineItemResponseDataRelationshipsStockTransfersWithDefaults instantiates a new LineItemResponseDataRelationshipsStockTransfers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemResponseDataRelationshipsStockTransfersWithDefaults() *LineItemResponseDataRelationshipsStockTransfers {
	this := LineItemResponseDataRelationshipsStockTransfers{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *LineItemResponseDataRelationshipsStockTransfers) GetLinks() AddressResponseDataRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret AddressResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemResponseDataRelationshipsStockTransfers) GetLinksOk() (*AddressResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *LineItemResponseDataRelationshipsStockTransfers) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AddressResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *LineItemResponseDataRelationshipsStockTransfers) SetLinks(v AddressResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LineItemResponseDataRelationshipsStockTransfers) GetData() []LineItemResponseDataRelationshipsStockTransfersDataInner {
	if o == nil || o.Data == nil {
		var ret []LineItemResponseDataRelationshipsStockTransfersDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemResponseDataRelationshipsStockTransfers) GetDataOk() ([]LineItemResponseDataRelationshipsStockTransfersDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LineItemResponseDataRelationshipsStockTransfers) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []LineItemResponseDataRelationshipsStockTransfersDataInner and assigns it to the Data field.
func (o *LineItemResponseDataRelationshipsStockTransfers) SetData(v []LineItemResponseDataRelationshipsStockTransfersDataInner) {
	o.Data = v
}

func (o LineItemResponseDataRelationshipsStockTransfers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableLineItemResponseDataRelationshipsStockTransfers struct {
	value *LineItemResponseDataRelationshipsStockTransfers
	isSet bool
}

func (v NullableLineItemResponseDataRelationshipsStockTransfers) Get() *LineItemResponseDataRelationshipsStockTransfers {
	return v.value
}

func (v *NullableLineItemResponseDataRelationshipsStockTransfers) Set(val *LineItemResponseDataRelationshipsStockTransfers) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemResponseDataRelationshipsStockTransfers) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemResponseDataRelationshipsStockTransfers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemResponseDataRelationshipsStockTransfers(val *LineItemResponseDataRelationshipsStockTransfers) *NullableLineItemResponseDataRelationshipsStockTransfers {
	return &NullableLineItemResponseDataRelationshipsStockTransfers{value: val, isSet: true}
}

func (v NullableLineItemResponseDataRelationshipsStockTransfers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemResponseDataRelationshipsStockTransfers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
