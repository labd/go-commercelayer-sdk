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

// AdyenGatewayResponseDataRelationshipsAdyenPayments struct for AdyenGatewayResponseDataRelationshipsAdyenPayments
type AdyenGatewayResponseDataRelationshipsAdyenPayments struct {
	Links *AddressResponseDataRelationshipsGeocoderLinks                `json:"links,omitempty"`
	Data  []AdyenGatewayResponseDataRelationshipsAdyenPaymentsDataInner `json:"data,omitempty"`
}

// NewAdyenGatewayResponseDataRelationshipsAdyenPayments instantiates a new AdyenGatewayResponseDataRelationshipsAdyenPayments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenGatewayResponseDataRelationshipsAdyenPayments() *AdyenGatewayResponseDataRelationshipsAdyenPayments {
	this := AdyenGatewayResponseDataRelationshipsAdyenPayments{}
	return &this
}

// NewAdyenGatewayResponseDataRelationshipsAdyenPaymentsWithDefaults instantiates a new AdyenGatewayResponseDataRelationshipsAdyenPayments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenGatewayResponseDataRelationshipsAdyenPaymentsWithDefaults() *AdyenGatewayResponseDataRelationshipsAdyenPayments {
	this := AdyenGatewayResponseDataRelationshipsAdyenPayments{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AdyenGatewayResponseDataRelationshipsAdyenPayments) GetLinks() AddressResponseDataRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret AddressResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenGatewayResponseDataRelationshipsAdyenPayments) GetLinksOk() (*AddressResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AdyenGatewayResponseDataRelationshipsAdyenPayments) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AddressResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *AdyenGatewayResponseDataRelationshipsAdyenPayments) SetLinks(v AddressResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AdyenGatewayResponseDataRelationshipsAdyenPayments) GetData() []AdyenGatewayResponseDataRelationshipsAdyenPaymentsDataInner {
	if o == nil || o.Data == nil {
		var ret []AdyenGatewayResponseDataRelationshipsAdyenPaymentsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenGatewayResponseDataRelationshipsAdyenPayments) GetDataOk() ([]AdyenGatewayResponseDataRelationshipsAdyenPaymentsDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AdyenGatewayResponseDataRelationshipsAdyenPayments) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []AdyenGatewayResponseDataRelationshipsAdyenPaymentsDataInner and assigns it to the Data field.
func (o *AdyenGatewayResponseDataRelationshipsAdyenPayments) SetData(v []AdyenGatewayResponseDataRelationshipsAdyenPaymentsDataInner) {
	o.Data = v
}

func (o AdyenGatewayResponseDataRelationshipsAdyenPayments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAdyenGatewayResponseDataRelationshipsAdyenPayments struct {
	value *AdyenGatewayResponseDataRelationshipsAdyenPayments
	isSet bool
}

func (v NullableAdyenGatewayResponseDataRelationshipsAdyenPayments) Get() *AdyenGatewayResponseDataRelationshipsAdyenPayments {
	return v.value
}

func (v *NullableAdyenGatewayResponseDataRelationshipsAdyenPayments) Set(val *AdyenGatewayResponseDataRelationshipsAdyenPayments) {
	v.value = val
	v.isSet = true
}

func (v NullableAdyenGatewayResponseDataRelationshipsAdyenPayments) IsSet() bool {
	return v.isSet
}

func (v *NullableAdyenGatewayResponseDataRelationshipsAdyenPayments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdyenGatewayResponseDataRelationshipsAdyenPayments(val *AdyenGatewayResponseDataRelationshipsAdyenPayments) *NullableAdyenGatewayResponseDataRelationshipsAdyenPayments {
	return &NullableAdyenGatewayResponseDataRelationshipsAdyenPayments{value: val, isSet: true}
}

func (v NullableAdyenGatewayResponseDataRelationshipsAdyenPayments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdyenGatewayResponseDataRelationshipsAdyenPayments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
