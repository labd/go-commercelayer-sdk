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

// GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner struct for GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner
type GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner instantiates a new GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner() *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner {
	this := GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner{}
	return &this
}

// NewGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInnerWithDefaults instantiates a new GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInnerWithDefaults() *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner {
	this := GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner) SetId(v string) {
	o.Id = &v
}

func (o GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner struct {
	value *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner
	isSet bool
}

func (v NullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner) Get() *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner {
	return v.value
}

func (v *NullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner) Set(val *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner(val *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner) *NullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner {
	return &NullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner{value: val, isSet: true}
}

func (v NullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
