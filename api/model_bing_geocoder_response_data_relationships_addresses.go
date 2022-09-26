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

// BingGeocoderResponseDataRelationshipsAddresses struct for BingGeocoderResponseDataRelationshipsAddresses
type BingGeocoderResponseDataRelationshipsAddresses struct {
	Links *AddressResponseDataRelationshipsGeocoderLinks            `json:"links,omitempty"`
	Data  []BingGeocoderResponseDataRelationshipsAddressesDataInner `json:"data,omitempty"`
}

// NewBingGeocoderResponseDataRelationshipsAddresses instantiates a new BingGeocoderResponseDataRelationshipsAddresses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBingGeocoderResponseDataRelationshipsAddresses() *BingGeocoderResponseDataRelationshipsAddresses {
	this := BingGeocoderResponseDataRelationshipsAddresses{}
	return &this
}

// NewBingGeocoderResponseDataRelationshipsAddressesWithDefaults instantiates a new BingGeocoderResponseDataRelationshipsAddresses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBingGeocoderResponseDataRelationshipsAddressesWithDefaults() *BingGeocoderResponseDataRelationshipsAddresses {
	this := BingGeocoderResponseDataRelationshipsAddresses{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BingGeocoderResponseDataRelationshipsAddresses) GetLinks() AddressResponseDataRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret AddressResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BingGeocoderResponseDataRelationshipsAddresses) GetLinksOk() (*AddressResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BingGeocoderResponseDataRelationshipsAddresses) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AddressResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *BingGeocoderResponseDataRelationshipsAddresses) SetLinks(v AddressResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BingGeocoderResponseDataRelationshipsAddresses) GetData() []BingGeocoderResponseDataRelationshipsAddressesDataInner {
	if o == nil || o.Data == nil {
		var ret []BingGeocoderResponseDataRelationshipsAddressesDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BingGeocoderResponseDataRelationshipsAddresses) GetDataOk() ([]BingGeocoderResponseDataRelationshipsAddressesDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BingGeocoderResponseDataRelationshipsAddresses) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []BingGeocoderResponseDataRelationshipsAddressesDataInner and assigns it to the Data field.
func (o *BingGeocoderResponseDataRelationshipsAddresses) SetData(v []BingGeocoderResponseDataRelationshipsAddressesDataInner) {
	o.Data = v
}

func (o BingGeocoderResponseDataRelationshipsAddresses) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableBingGeocoderResponseDataRelationshipsAddresses struct {
	value *BingGeocoderResponseDataRelationshipsAddresses
	isSet bool
}

func (v NullableBingGeocoderResponseDataRelationshipsAddresses) Get() *BingGeocoderResponseDataRelationshipsAddresses {
	return v.value
}

func (v *NullableBingGeocoderResponseDataRelationshipsAddresses) Set(val *BingGeocoderResponseDataRelationshipsAddresses) {
	v.value = val
	v.isSet = true
}

func (v NullableBingGeocoderResponseDataRelationshipsAddresses) IsSet() bool {
	return v.isSet
}

func (v *NullableBingGeocoderResponseDataRelationshipsAddresses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBingGeocoderResponseDataRelationshipsAddresses(val *BingGeocoderResponseDataRelationshipsAddresses) *NullableBingGeocoderResponseDataRelationshipsAddresses {
	return &NullableBingGeocoderResponseDataRelationshipsAddresses{value: val, isSet: true}
}

func (v NullableBingGeocoderResponseDataRelationshipsAddresses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBingGeocoderResponseDataRelationshipsAddresses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
