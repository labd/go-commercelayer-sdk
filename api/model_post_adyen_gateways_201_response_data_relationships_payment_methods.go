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

// checks if the POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods{}

// POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods struct for POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods
type POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks          `json:"links,omitempty"`
	Data  *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData `json:"data,omitempty"`
}

// NewPOSTAdyenGateways201ResponseDataRelationshipsPaymentMethods instantiates a new POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAdyenGateways201ResponseDataRelationshipsPaymentMethods() *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods {
	this := POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods{}
	return &this
}

// NewPOSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsWithDefaults instantiates a new POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsWithDefaults() *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods {
	this := POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods) GetData() POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData {
	if o == nil || IsNil(o.Data) {
		var ret POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods) GetDataOk() (*POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData and assigns it to the Data field.
func (o *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods) SetData(v POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData) {
	o.Data = &v
}

func (o POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTAdyenGateways201ResponseDataRelationshipsPaymentMethods struct {
	value *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods
	isSet bool
}

func (v NullablePOSTAdyenGateways201ResponseDataRelationshipsPaymentMethods) Get() *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods {
	return v.value
}

func (v *NullablePOSTAdyenGateways201ResponseDataRelationshipsPaymentMethods) Set(val *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAdyenGateways201ResponseDataRelationshipsPaymentMethods) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAdyenGateways201ResponseDataRelationshipsPaymentMethods) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAdyenGateways201ResponseDataRelationshipsPaymentMethods(val *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods) *NullablePOSTAdyenGateways201ResponseDataRelationshipsPaymentMethods {
	return &NullablePOSTAdyenGateways201ResponseDataRelationshipsPaymentMethods{value: val, isSet: true}
}

func (v NullablePOSTAdyenGateways201ResponseDataRelationshipsPaymentMethods) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAdyenGateways201ResponseDataRelationshipsPaymentMethods) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
