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

// checks if the POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources{}

// POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources struct for POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources
type POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks              `json:"links,omitempty"`
	Data  *POSTCustomers201ResponseDataRelationshipsCustomerPaymentSourcesData `json:"data,omitempty"`
}

// NewPOSTCustomers201ResponseDataRelationshipsCustomerPaymentSources instantiates a new POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCustomers201ResponseDataRelationshipsCustomerPaymentSources() *POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources {
	this := POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources{}
	return &this
}

// NewPOSTCustomers201ResponseDataRelationshipsCustomerPaymentSourcesWithDefaults instantiates a new POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCustomers201ResponseDataRelationshipsCustomerPaymentSourcesWithDefaults() *POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources {
	this := POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources) GetData() POSTCustomers201ResponseDataRelationshipsCustomerPaymentSourcesData {
	if o == nil || IsNil(o.Data) {
		var ret POSTCustomers201ResponseDataRelationshipsCustomerPaymentSourcesData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources) GetDataOk() (*POSTCustomers201ResponseDataRelationshipsCustomerPaymentSourcesData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTCustomers201ResponseDataRelationshipsCustomerPaymentSourcesData and assigns it to the Data field.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources) SetData(v POSTCustomers201ResponseDataRelationshipsCustomerPaymentSourcesData) {
	o.Data = &v
}

func (o POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTCustomers201ResponseDataRelationshipsCustomerPaymentSources struct {
	value *POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources
	isSet bool
}

func (v NullablePOSTCustomers201ResponseDataRelationshipsCustomerPaymentSources) Get() *POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources {
	return v.value
}

func (v *NullablePOSTCustomers201ResponseDataRelationshipsCustomerPaymentSources) Set(val *POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCustomers201ResponseDataRelationshipsCustomerPaymentSources) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCustomers201ResponseDataRelationshipsCustomerPaymentSources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCustomers201ResponseDataRelationshipsCustomerPaymentSources(val *POSTCustomers201ResponseDataRelationshipsCustomerPaymentSources) *NullablePOSTCustomers201ResponseDataRelationshipsCustomerPaymentSources {
	return &NullablePOSTCustomers201ResponseDataRelationshipsCustomerPaymentSources{value: val, isSet: true}
}

func (v NullablePOSTCustomers201ResponseDataRelationshipsCustomerPaymentSources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCustomers201ResponseDataRelationshipsCustomerPaymentSources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
