/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel struct for GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel
type GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks                     `json:"links,omitempty"`
	Data  *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData `json:"data,omitempty"`
}

// NewGETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel instantiates a new GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel() *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel {
	this := GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel{}
	return &this
}

// NewGETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelWithDefaults instantiates a new GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelWithDefaults() *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel {
	this := GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel) GetData() GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData {
	if o == nil || o.Data == nil {
		var ret GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel) GetDataOk() (*GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData and assigns it to the Data field.
func (o *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel) SetData(v GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModelData) {
	o.Data = &v
}

func (o GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel struct {
	value *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel
	isSet bool
}

func (v NullableGETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel) Get() *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel {
	return v.value
}

func (v *NullableGETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel) Set(val *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel(val *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel) *NullableGETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel {
	return &NullableGETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel{value: val, isSet: true}
}

func (v NullableGETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
