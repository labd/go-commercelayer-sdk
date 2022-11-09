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

// GETPrices200ResponseDataInnerRelationshipsPriceTiers struct for GETPrices200ResponseDataInnerRelationshipsPriceTiers
type GETPrices200ResponseDataInnerRelationshipsPriceTiers struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks     `json:"links,omitempty"`
	Data  []GETPrices200ResponseDataInnerRelationshipsPriceTiersDataInner `json:"data,omitempty"`
}

// NewGETPrices200ResponseDataInnerRelationshipsPriceTiers instantiates a new GETPrices200ResponseDataInnerRelationshipsPriceTiers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPrices200ResponseDataInnerRelationshipsPriceTiers() *GETPrices200ResponseDataInnerRelationshipsPriceTiers {
	this := GETPrices200ResponseDataInnerRelationshipsPriceTiers{}
	return &this
}

// NewGETPrices200ResponseDataInnerRelationshipsPriceTiersWithDefaults instantiates a new GETPrices200ResponseDataInnerRelationshipsPriceTiers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPrices200ResponseDataInnerRelationshipsPriceTiersWithDefaults() *GETPrices200ResponseDataInnerRelationshipsPriceTiers {
	this := GETPrices200ResponseDataInnerRelationshipsPriceTiers{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETPrices200ResponseDataInnerRelationshipsPriceTiers) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPrices200ResponseDataInnerRelationshipsPriceTiers) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETPrices200ResponseDataInnerRelationshipsPriceTiers) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETPrices200ResponseDataInnerRelationshipsPriceTiers) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETPrices200ResponseDataInnerRelationshipsPriceTiers) GetData() []GETPrices200ResponseDataInnerRelationshipsPriceTiersDataInner {
	if o == nil || o.Data == nil {
		var ret []GETPrices200ResponseDataInnerRelationshipsPriceTiersDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPrices200ResponseDataInnerRelationshipsPriceTiers) GetDataOk() ([]GETPrices200ResponseDataInnerRelationshipsPriceTiersDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETPrices200ResponseDataInnerRelationshipsPriceTiers) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETPrices200ResponseDataInnerRelationshipsPriceTiersDataInner and assigns it to the Data field.
func (o *GETPrices200ResponseDataInnerRelationshipsPriceTiers) SetData(v []GETPrices200ResponseDataInnerRelationshipsPriceTiersDataInner) {
	o.Data = v
}

func (o GETPrices200ResponseDataInnerRelationshipsPriceTiers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETPrices200ResponseDataInnerRelationshipsPriceTiers struct {
	value *GETPrices200ResponseDataInnerRelationshipsPriceTiers
	isSet bool
}

func (v NullableGETPrices200ResponseDataInnerRelationshipsPriceTiers) Get() *GETPrices200ResponseDataInnerRelationshipsPriceTiers {
	return v.value
}

func (v *NullableGETPrices200ResponseDataInnerRelationshipsPriceTiers) Set(val *GETPrices200ResponseDataInnerRelationshipsPriceTiers) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPrices200ResponseDataInnerRelationshipsPriceTiers) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPrices200ResponseDataInnerRelationshipsPriceTiers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPrices200ResponseDataInnerRelationshipsPriceTiers(val *GETPrices200ResponseDataInnerRelationshipsPriceTiers) *NullableGETPrices200ResponseDataInnerRelationshipsPriceTiers {
	return &NullableGETPrices200ResponseDataInnerRelationshipsPriceTiers{value: val, isSet: true}
}

func (v NullableGETPrices200ResponseDataInnerRelationshipsPriceTiers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPrices200ResponseDataInnerRelationshipsPriceTiers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}