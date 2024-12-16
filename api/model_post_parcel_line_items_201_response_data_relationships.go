/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.5.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTParcelLineItems201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTParcelLineItems201ResponseDataRelationships{}

// POSTParcelLineItems201ResponseDataRelationships struct for POSTParcelLineItems201ResponseDataRelationships
type POSTParcelLineItems201ResponseDataRelationships struct {
	Parcel        *POSTParcelLineItems201ResponseDataRelationshipsParcel        `json:"parcel,omitempty"`
	StockLineItem *POSTParcelLineItems201ResponseDataRelationshipsStockLineItem `json:"stock_line_item,omitempty"`
	Versions      *POSTAddresses201ResponseDataRelationshipsVersions            `json:"versions,omitempty"`
}

// NewPOSTParcelLineItems201ResponseDataRelationships instantiates a new POSTParcelLineItems201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTParcelLineItems201ResponseDataRelationships() *POSTParcelLineItems201ResponseDataRelationships {
	this := POSTParcelLineItems201ResponseDataRelationships{}
	return &this
}

// NewPOSTParcelLineItems201ResponseDataRelationshipsWithDefaults instantiates a new POSTParcelLineItems201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTParcelLineItems201ResponseDataRelationshipsWithDefaults() *POSTParcelLineItems201ResponseDataRelationships {
	this := POSTParcelLineItems201ResponseDataRelationships{}
	return &this
}

// GetParcel returns the Parcel field value if set, zero value otherwise.
func (o *POSTParcelLineItems201ResponseDataRelationships) GetParcel() POSTParcelLineItems201ResponseDataRelationshipsParcel {
	if o == nil || IsNil(o.Parcel) {
		var ret POSTParcelLineItems201ResponseDataRelationshipsParcel
		return ret
	}
	return *o.Parcel
}

// GetParcelOk returns a tuple with the Parcel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTParcelLineItems201ResponseDataRelationships) GetParcelOk() (*POSTParcelLineItems201ResponseDataRelationshipsParcel, bool) {
	if o == nil || IsNil(o.Parcel) {
		return nil, false
	}
	return o.Parcel, true
}

// HasParcel returns a boolean if a field has been set.
func (o *POSTParcelLineItems201ResponseDataRelationships) HasParcel() bool {
	if o != nil && !IsNil(o.Parcel) {
		return true
	}

	return false
}

// SetParcel gets a reference to the given POSTParcelLineItems201ResponseDataRelationshipsParcel and assigns it to the Parcel field.
func (o *POSTParcelLineItems201ResponseDataRelationships) SetParcel(v POSTParcelLineItems201ResponseDataRelationshipsParcel) {
	o.Parcel = &v
}

// GetStockLineItem returns the StockLineItem field value if set, zero value otherwise.
func (o *POSTParcelLineItems201ResponseDataRelationships) GetStockLineItem() POSTParcelLineItems201ResponseDataRelationshipsStockLineItem {
	if o == nil || IsNil(o.StockLineItem) {
		var ret POSTParcelLineItems201ResponseDataRelationshipsStockLineItem
		return ret
	}
	return *o.StockLineItem
}

// GetStockLineItemOk returns a tuple with the StockLineItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTParcelLineItems201ResponseDataRelationships) GetStockLineItemOk() (*POSTParcelLineItems201ResponseDataRelationshipsStockLineItem, bool) {
	if o == nil || IsNil(o.StockLineItem) {
		return nil, false
	}
	return o.StockLineItem, true
}

// HasStockLineItem returns a boolean if a field has been set.
func (o *POSTParcelLineItems201ResponseDataRelationships) HasStockLineItem() bool {
	if o != nil && !IsNil(o.StockLineItem) {
		return true
	}

	return false
}

// SetStockLineItem gets a reference to the given POSTParcelLineItems201ResponseDataRelationshipsStockLineItem and assigns it to the StockLineItem field.
func (o *POSTParcelLineItems201ResponseDataRelationships) SetStockLineItem(v POSTParcelLineItems201ResponseDataRelationshipsStockLineItem) {
	o.StockLineItem = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *POSTParcelLineItems201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTParcelLineItems201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *POSTParcelLineItems201ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *POSTParcelLineItems201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

func (o POSTParcelLineItems201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTParcelLineItems201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Parcel) {
		toSerialize["parcel"] = o.Parcel
	}
	if !IsNil(o.StockLineItem) {
		toSerialize["stock_line_item"] = o.StockLineItem
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullablePOSTParcelLineItems201ResponseDataRelationships struct {
	value *POSTParcelLineItems201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTParcelLineItems201ResponseDataRelationships) Get() *POSTParcelLineItems201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTParcelLineItems201ResponseDataRelationships) Set(val *POSTParcelLineItems201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTParcelLineItems201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTParcelLineItems201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTParcelLineItems201ResponseDataRelationships(val *POSTParcelLineItems201ResponseDataRelationships) *NullablePOSTParcelLineItems201ResponseDataRelationships {
	return &NullablePOSTParcelLineItems201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTParcelLineItems201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTParcelLineItems201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
