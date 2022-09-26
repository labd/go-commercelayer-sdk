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

// ShippingWeightTierCreateDataRelationships struct for ShippingWeightTierCreateDataRelationships
type ShippingWeightTierCreateDataRelationships struct {
	ShippingMethod DeliveryLeadTimeDataRelationshipsShippingMethod `json:"shipping_method"`
}

// NewShippingWeightTierCreateDataRelationships instantiates a new ShippingWeightTierCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingWeightTierCreateDataRelationships(shippingMethod DeliveryLeadTimeDataRelationshipsShippingMethod) *ShippingWeightTierCreateDataRelationships {
	this := ShippingWeightTierCreateDataRelationships{}
	this.ShippingMethod = shippingMethod
	return &this
}

// NewShippingWeightTierCreateDataRelationshipsWithDefaults instantiates a new ShippingWeightTierCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingWeightTierCreateDataRelationshipsWithDefaults() *ShippingWeightTierCreateDataRelationships {
	this := ShippingWeightTierCreateDataRelationships{}
	return &this
}

// GetShippingMethod returns the ShippingMethod field value
func (o *ShippingWeightTierCreateDataRelationships) GetShippingMethod() DeliveryLeadTimeDataRelationshipsShippingMethod {
	if o == nil {
		var ret DeliveryLeadTimeDataRelationshipsShippingMethod
		return ret
	}

	return o.ShippingMethod
}

// GetShippingMethodOk returns a tuple with the ShippingMethod field value
// and a boolean to check if the value has been set.
func (o *ShippingWeightTierCreateDataRelationships) GetShippingMethodOk() (*DeliveryLeadTimeDataRelationshipsShippingMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShippingMethod, true
}

// SetShippingMethod sets field value
func (o *ShippingWeightTierCreateDataRelationships) SetShippingMethod(v DeliveryLeadTimeDataRelationshipsShippingMethod) {
	o.ShippingMethod = v
}

func (o ShippingWeightTierCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["shipping_method"] = o.ShippingMethod
	}
	return json.Marshal(toSerialize)
}

type NullableShippingWeightTierCreateDataRelationships struct {
	value *ShippingWeightTierCreateDataRelationships
	isSet bool
}

func (v NullableShippingWeightTierCreateDataRelationships) Get() *ShippingWeightTierCreateDataRelationships {
	return v.value
}

func (v *NullableShippingWeightTierCreateDataRelationships) Set(val *ShippingWeightTierCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingWeightTierCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingWeightTierCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingWeightTierCreateDataRelationships(val *ShippingWeightTierCreateDataRelationships) *NullableShippingWeightTierCreateDataRelationships {
	return &NullableShippingWeightTierCreateDataRelationships{value: val, isSet: true}
}

func (v NullableShippingWeightTierCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingWeightTierCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
