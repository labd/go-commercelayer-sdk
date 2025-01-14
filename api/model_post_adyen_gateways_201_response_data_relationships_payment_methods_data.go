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

// checks if the POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData{}

// POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData struct for POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData
type POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData instantiates a new POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData() *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData {
	this := POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData{}
	return &this
}

// NewPOSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsDataWithDefaults instantiates a new POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsDataWithDefaults() *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData {
	this := POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData struct {
	value *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData
	isSet bool
}

func (v NullablePOSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData) Get() *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData {
	return v.value
}

func (v *NullablePOSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData) Set(val *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData(val *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData) *NullablePOSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData {
	return &NullablePOSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData{value: val, isSet: true}
}

func (v NullablePOSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAdyenGateways201ResponseDataRelationshipsPaymentMethodsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
