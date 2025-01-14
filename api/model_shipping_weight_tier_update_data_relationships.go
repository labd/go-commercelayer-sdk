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

// checks if the ShippingWeightTierUpdateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingWeightTierUpdateDataRelationships{}

// ShippingWeightTierUpdateDataRelationships struct for ShippingWeightTierUpdateDataRelationships
type ShippingWeightTierUpdateDataRelationships struct {
	ShippingMethod *DeliveryLeadTimeCreateDataRelationshipsShippingMethod `json:"shipping_method,omitempty"`
}

// NewShippingWeightTierUpdateDataRelationships instantiates a new ShippingWeightTierUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingWeightTierUpdateDataRelationships() *ShippingWeightTierUpdateDataRelationships {
	this := ShippingWeightTierUpdateDataRelationships{}
	return &this
}

// NewShippingWeightTierUpdateDataRelationshipsWithDefaults instantiates a new ShippingWeightTierUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingWeightTierUpdateDataRelationshipsWithDefaults() *ShippingWeightTierUpdateDataRelationships {
	this := ShippingWeightTierUpdateDataRelationships{}
	return &this
}

// GetShippingMethod returns the ShippingMethod field value if set, zero value otherwise.
func (o *ShippingWeightTierUpdateDataRelationships) GetShippingMethod() DeliveryLeadTimeCreateDataRelationshipsShippingMethod {
	if o == nil || IsNil(o.ShippingMethod) {
		var ret DeliveryLeadTimeCreateDataRelationshipsShippingMethod
		return ret
	}
	return *o.ShippingMethod
}

// GetShippingMethodOk returns a tuple with the ShippingMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingWeightTierUpdateDataRelationships) GetShippingMethodOk() (*DeliveryLeadTimeCreateDataRelationshipsShippingMethod, bool) {
	if o == nil || IsNil(o.ShippingMethod) {
		return nil, false
	}
	return o.ShippingMethod, true
}

// HasShippingMethod returns a boolean if a field has been set.
func (o *ShippingWeightTierUpdateDataRelationships) HasShippingMethod() bool {
	if o != nil && !IsNil(o.ShippingMethod) {
		return true
	}

	return false
}

// SetShippingMethod gets a reference to the given DeliveryLeadTimeCreateDataRelationshipsShippingMethod and assigns it to the ShippingMethod field.
func (o *ShippingWeightTierUpdateDataRelationships) SetShippingMethod(v DeliveryLeadTimeCreateDataRelationshipsShippingMethod) {
	o.ShippingMethod = &v
}

func (o ShippingWeightTierUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShippingWeightTierUpdateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShippingMethod) {
		toSerialize["shipping_method"] = o.ShippingMethod
	}
	return toSerialize, nil
}

type NullableShippingWeightTierUpdateDataRelationships struct {
	value *ShippingWeightTierUpdateDataRelationships
	isSet bool
}

func (v NullableShippingWeightTierUpdateDataRelationships) Get() *ShippingWeightTierUpdateDataRelationships {
	return v.value
}

func (v *NullableShippingWeightTierUpdateDataRelationships) Set(val *ShippingWeightTierUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingWeightTierUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingWeightTierUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingWeightTierUpdateDataRelationships(val *ShippingWeightTierUpdateDataRelationships) *NullableShippingWeightTierUpdateDataRelationships {
	return &NullableShippingWeightTierUpdateDataRelationships{value: val, isSet: true}
}

func (v NullableShippingWeightTierUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingWeightTierUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
