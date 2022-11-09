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

// ShipmentUpdateDataRelationships struct for ShipmentUpdateDataRelationships
type ShipmentUpdateDataRelationships struct {
	ShippingMethod *DeliveryLeadTimeDataRelationshipsShippingMethod `json:"shipping_method,omitempty"`
}

// NewShipmentUpdateDataRelationships instantiates a new ShipmentUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentUpdateDataRelationships() *ShipmentUpdateDataRelationships {
	this := ShipmentUpdateDataRelationships{}
	return &this
}

// NewShipmentUpdateDataRelationshipsWithDefaults instantiates a new ShipmentUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentUpdateDataRelationshipsWithDefaults() *ShipmentUpdateDataRelationships {
	this := ShipmentUpdateDataRelationships{}
	return &this
}

// GetShippingMethod returns the ShippingMethod field value if set, zero value otherwise.
func (o *ShipmentUpdateDataRelationships) GetShippingMethod() DeliveryLeadTimeDataRelationshipsShippingMethod {
	if o == nil || o.ShippingMethod == nil {
		var ret DeliveryLeadTimeDataRelationshipsShippingMethod
		return ret
	}
	return *o.ShippingMethod
}

// GetShippingMethodOk returns a tuple with the ShippingMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentUpdateDataRelationships) GetShippingMethodOk() (*DeliveryLeadTimeDataRelationshipsShippingMethod, bool) {
	if o == nil || o.ShippingMethod == nil {
		return nil, false
	}
	return o.ShippingMethod, true
}

// HasShippingMethod returns a boolean if a field has been set.
func (o *ShipmentUpdateDataRelationships) HasShippingMethod() bool {
	if o != nil && o.ShippingMethod != nil {
		return true
	}

	return false
}

// SetShippingMethod gets a reference to the given DeliveryLeadTimeDataRelationshipsShippingMethod and assigns it to the ShippingMethod field.
func (o *ShipmentUpdateDataRelationships) SetShippingMethod(v DeliveryLeadTimeDataRelationshipsShippingMethod) {
	o.ShippingMethod = &v
}

func (o ShipmentUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ShippingMethod != nil {
		toSerialize["shipping_method"] = o.ShippingMethod
	}
	return json.Marshal(toSerialize)
}

type NullableShipmentUpdateDataRelationships struct {
	value *ShipmentUpdateDataRelationships
	isSet bool
}

func (v NullableShipmentUpdateDataRelationships) Get() *ShipmentUpdateDataRelationships {
	return v.value
}

func (v *NullableShipmentUpdateDataRelationships) Set(val *ShipmentUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentUpdateDataRelationships(val *ShipmentUpdateDataRelationships) *NullableShipmentUpdateDataRelationships {
	return &NullableShipmentUpdateDataRelationships{value: val, isSet: true}
}

func (v NullableShipmentUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
