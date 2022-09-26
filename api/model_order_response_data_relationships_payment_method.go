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

// OrderResponseDataRelationshipsPaymentMethod struct for OrderResponseDataRelationshipsPaymentMethod
type OrderResponseDataRelationshipsPaymentMethod struct {
	Links *AddressResponseDataRelationshipsGeocoderLinks   `json:"links,omitempty"`
	Data  *OrderResponseDataRelationshipsPaymentMethodData `json:"data,omitempty"`
}

// NewOrderResponseDataRelationshipsPaymentMethod instantiates a new OrderResponseDataRelationshipsPaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderResponseDataRelationshipsPaymentMethod() *OrderResponseDataRelationshipsPaymentMethod {
	this := OrderResponseDataRelationshipsPaymentMethod{}
	return &this
}

// NewOrderResponseDataRelationshipsPaymentMethodWithDefaults instantiates a new OrderResponseDataRelationshipsPaymentMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderResponseDataRelationshipsPaymentMethodWithDefaults() *OrderResponseDataRelationshipsPaymentMethod {
	this := OrderResponseDataRelationshipsPaymentMethod{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrderResponseDataRelationshipsPaymentMethod) GetLinks() AddressResponseDataRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret AddressResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationshipsPaymentMethod) GetLinksOk() (*AddressResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrderResponseDataRelationshipsPaymentMethod) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AddressResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *OrderResponseDataRelationshipsPaymentMethod) SetLinks(v AddressResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OrderResponseDataRelationshipsPaymentMethod) GetData() OrderResponseDataRelationshipsPaymentMethodData {
	if o == nil || o.Data == nil {
		var ret OrderResponseDataRelationshipsPaymentMethodData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationshipsPaymentMethod) GetDataOk() (*OrderResponseDataRelationshipsPaymentMethodData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OrderResponseDataRelationshipsPaymentMethod) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given OrderResponseDataRelationshipsPaymentMethodData and assigns it to the Data field.
func (o *OrderResponseDataRelationshipsPaymentMethod) SetData(v OrderResponseDataRelationshipsPaymentMethodData) {
	o.Data = &v
}

func (o OrderResponseDataRelationshipsPaymentMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableOrderResponseDataRelationshipsPaymentMethod struct {
	value *OrderResponseDataRelationshipsPaymentMethod
	isSet bool
}

func (v NullableOrderResponseDataRelationshipsPaymentMethod) Get() *OrderResponseDataRelationshipsPaymentMethod {
	return v.value
}

func (v *NullableOrderResponseDataRelationshipsPaymentMethod) Set(val *OrderResponseDataRelationshipsPaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderResponseDataRelationshipsPaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderResponseDataRelationshipsPaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderResponseDataRelationshipsPaymentMethod(val *OrderResponseDataRelationshipsPaymentMethod) *NullableOrderResponseDataRelationshipsPaymentMethod {
	return &NullableOrderResponseDataRelationshipsPaymentMethod{value: val, isSet: true}
}

func (v NullableOrderResponseDataRelationshipsPaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderResponseDataRelationshipsPaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
