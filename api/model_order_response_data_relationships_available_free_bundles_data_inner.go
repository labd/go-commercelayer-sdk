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

// OrderResponseDataRelationshipsAvailableFreeBundlesDataInner struct for OrderResponseDataRelationshipsAvailableFreeBundlesDataInner
type OrderResponseDataRelationshipsAvailableFreeBundlesDataInner struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewOrderResponseDataRelationshipsAvailableFreeBundlesDataInner instantiates a new OrderResponseDataRelationshipsAvailableFreeBundlesDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderResponseDataRelationshipsAvailableFreeBundlesDataInner() *OrderResponseDataRelationshipsAvailableFreeBundlesDataInner {
	this := OrderResponseDataRelationshipsAvailableFreeBundlesDataInner{}
	return &this
}

// NewOrderResponseDataRelationshipsAvailableFreeBundlesDataInnerWithDefaults instantiates a new OrderResponseDataRelationshipsAvailableFreeBundlesDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderResponseDataRelationshipsAvailableFreeBundlesDataInnerWithDefaults() *OrderResponseDataRelationshipsAvailableFreeBundlesDataInner {
	this := OrderResponseDataRelationshipsAvailableFreeBundlesDataInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrderResponseDataRelationshipsAvailableFreeBundlesDataInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationshipsAvailableFreeBundlesDataInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrderResponseDataRelationshipsAvailableFreeBundlesDataInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrderResponseDataRelationshipsAvailableFreeBundlesDataInner) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderResponseDataRelationshipsAvailableFreeBundlesDataInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationshipsAvailableFreeBundlesDataInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderResponseDataRelationshipsAvailableFreeBundlesDataInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderResponseDataRelationshipsAvailableFreeBundlesDataInner) SetId(v string) {
	o.Id = &v
}

func (o OrderResponseDataRelationshipsAvailableFreeBundlesDataInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableOrderResponseDataRelationshipsAvailableFreeBundlesDataInner struct {
	value *OrderResponseDataRelationshipsAvailableFreeBundlesDataInner
	isSet bool
}

func (v NullableOrderResponseDataRelationshipsAvailableFreeBundlesDataInner) Get() *OrderResponseDataRelationshipsAvailableFreeBundlesDataInner {
	return v.value
}

func (v *NullableOrderResponseDataRelationshipsAvailableFreeBundlesDataInner) Set(val *OrderResponseDataRelationshipsAvailableFreeBundlesDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderResponseDataRelationshipsAvailableFreeBundlesDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderResponseDataRelationshipsAvailableFreeBundlesDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderResponseDataRelationshipsAvailableFreeBundlesDataInner(val *OrderResponseDataRelationshipsAvailableFreeBundlesDataInner) *NullableOrderResponseDataRelationshipsAvailableFreeBundlesDataInner {
	return &NullableOrderResponseDataRelationshipsAvailableFreeBundlesDataInner{value: val, isSet: true}
}

func (v NullableOrderResponseDataRelationshipsAvailableFreeBundlesDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderResponseDataRelationshipsAvailableFreeBundlesDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
