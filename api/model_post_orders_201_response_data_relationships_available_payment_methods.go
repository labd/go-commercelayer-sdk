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

// checks if the POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods{}

// POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods struct for POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods
type POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks            `json:"links,omitempty"`
	Data  *POSTOrders201ResponseDataRelationshipsAvailablePaymentMethodsData `json:"data,omitempty"`
}

// NewPOSTOrders201ResponseDataRelationshipsAvailablePaymentMethods instantiates a new POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrders201ResponseDataRelationshipsAvailablePaymentMethods() *POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods {
	this := POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods{}
	return &this
}

// NewPOSTOrders201ResponseDataRelationshipsAvailablePaymentMethodsWithDefaults instantiates a new POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrders201ResponseDataRelationshipsAvailablePaymentMethodsWithDefaults() *POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods {
	this := POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods) GetData() POSTOrders201ResponseDataRelationshipsAvailablePaymentMethodsData {
	if o == nil || IsNil(o.Data) {
		var ret POSTOrders201ResponseDataRelationshipsAvailablePaymentMethodsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods) GetDataOk() (*POSTOrders201ResponseDataRelationshipsAvailablePaymentMethodsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTOrders201ResponseDataRelationshipsAvailablePaymentMethodsData and assigns it to the Data field.
func (o *POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods) SetData(v POSTOrders201ResponseDataRelationshipsAvailablePaymentMethodsData) {
	o.Data = &v
}

func (o POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTOrders201ResponseDataRelationshipsAvailablePaymentMethods struct {
	value *POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods
	isSet bool
}

func (v NullablePOSTOrders201ResponseDataRelationshipsAvailablePaymentMethods) Get() *POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods {
	return v.value
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsAvailablePaymentMethods) Set(val *POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrders201ResponseDataRelationshipsAvailablePaymentMethods) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsAvailablePaymentMethods) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrders201ResponseDataRelationshipsAvailablePaymentMethods(val *POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods) *NullablePOSTOrders201ResponseDataRelationshipsAvailablePaymentMethods {
	return &NullablePOSTOrders201ResponseDataRelationshipsAvailablePaymentMethods{value: val, isSet: true}
}

func (v NullablePOSTOrders201ResponseDataRelationshipsAvailablePaymentMethods) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsAvailablePaymentMethods) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
