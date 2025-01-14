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

// checks if the POSTOrders201ResponseDataRelationshipsPaymentOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrders201ResponseDataRelationshipsPaymentOptions{}

// POSTOrders201ResponseDataRelationshipsPaymentOptions struct for POSTOrders201ResponseDataRelationshipsPaymentOptions
type POSTOrders201ResponseDataRelationshipsPaymentOptions struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks   `json:"links,omitempty"`
	Data  *POSTOrders201ResponseDataRelationshipsPaymentOptionsData `json:"data,omitempty"`
}

// NewPOSTOrders201ResponseDataRelationshipsPaymentOptions instantiates a new POSTOrders201ResponseDataRelationshipsPaymentOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrders201ResponseDataRelationshipsPaymentOptions() *POSTOrders201ResponseDataRelationshipsPaymentOptions {
	this := POSTOrders201ResponseDataRelationshipsPaymentOptions{}
	return &this
}

// NewPOSTOrders201ResponseDataRelationshipsPaymentOptionsWithDefaults instantiates a new POSTOrders201ResponseDataRelationshipsPaymentOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrders201ResponseDataRelationshipsPaymentOptionsWithDefaults() *POSTOrders201ResponseDataRelationshipsPaymentOptions {
	this := POSTOrders201ResponseDataRelationshipsPaymentOptions{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTOrders201ResponseDataRelationshipsPaymentOptions) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrders201ResponseDataRelationshipsPaymentOptions) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsPaymentOptions) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTOrders201ResponseDataRelationshipsPaymentOptions) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTOrders201ResponseDataRelationshipsPaymentOptions) GetData() POSTOrders201ResponseDataRelationshipsPaymentOptionsData {
	if o == nil || IsNil(o.Data) {
		var ret POSTOrders201ResponseDataRelationshipsPaymentOptionsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrders201ResponseDataRelationshipsPaymentOptions) GetDataOk() (*POSTOrders201ResponseDataRelationshipsPaymentOptionsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsPaymentOptions) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTOrders201ResponseDataRelationshipsPaymentOptionsData and assigns it to the Data field.
func (o *POSTOrders201ResponseDataRelationshipsPaymentOptions) SetData(v POSTOrders201ResponseDataRelationshipsPaymentOptionsData) {
	o.Data = &v
}

func (o POSTOrders201ResponseDataRelationshipsPaymentOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrders201ResponseDataRelationshipsPaymentOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTOrders201ResponseDataRelationshipsPaymentOptions struct {
	value *POSTOrders201ResponseDataRelationshipsPaymentOptions
	isSet bool
}

func (v NullablePOSTOrders201ResponseDataRelationshipsPaymentOptions) Get() *POSTOrders201ResponseDataRelationshipsPaymentOptions {
	return v.value
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsPaymentOptions) Set(val *POSTOrders201ResponseDataRelationshipsPaymentOptions) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrders201ResponseDataRelationshipsPaymentOptions) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsPaymentOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrders201ResponseDataRelationshipsPaymentOptions(val *POSTOrders201ResponseDataRelationshipsPaymentOptions) *NullablePOSTOrders201ResponseDataRelationshipsPaymentOptions {
	return &NullablePOSTOrders201ResponseDataRelationshipsPaymentOptions{value: val, isSet: true}
}

func (v NullablePOSTOrders201ResponseDataRelationshipsPaymentOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsPaymentOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
