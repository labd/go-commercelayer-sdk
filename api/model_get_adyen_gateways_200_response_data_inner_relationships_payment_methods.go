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

// GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods struct for GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods
type GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks                `json:"links,omitempty"`
	Data  []GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner `json:"data,omitempty"`
}

// NewGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods instantiates a new GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods() *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods {
	this := GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods{}
	return &this
}

// NewGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsWithDefaults instantiates a new GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsWithDefaults() *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods {
	this := GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods) GetData() []GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner {
	if o == nil || o.Data == nil {
		var ret []GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods) GetDataOk() ([]GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner and assigns it to the Data field.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods) SetData(v []GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataInner) {
	o.Data = v
}

func (o GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods struct {
	value *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods
	isSet bool
}

func (v NullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods) Get() *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods {
	return v.value
}

func (v *NullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods) Set(val *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods(val *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods) *NullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods {
	return &NullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods{value: val, isSet: true}
}

func (v NullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}