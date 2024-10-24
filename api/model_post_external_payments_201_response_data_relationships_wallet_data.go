/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTExternalPayments201ResponseDataRelationshipsWalletData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTExternalPayments201ResponseDataRelationshipsWalletData{}

// POSTExternalPayments201ResponseDataRelationshipsWalletData struct for POSTExternalPayments201ResponseDataRelationshipsWalletData
type POSTExternalPayments201ResponseDataRelationshipsWalletData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTExternalPayments201ResponseDataRelationshipsWalletData instantiates a new POSTExternalPayments201ResponseDataRelationshipsWalletData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTExternalPayments201ResponseDataRelationshipsWalletData() *POSTExternalPayments201ResponseDataRelationshipsWalletData {
	this := POSTExternalPayments201ResponseDataRelationshipsWalletData{}
	return &this
}

// NewPOSTExternalPayments201ResponseDataRelationshipsWalletDataWithDefaults instantiates a new POSTExternalPayments201ResponseDataRelationshipsWalletData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTExternalPayments201ResponseDataRelationshipsWalletDataWithDefaults() *POSTExternalPayments201ResponseDataRelationshipsWalletData {
	this := POSTExternalPayments201ResponseDataRelationshipsWalletData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTExternalPayments201ResponseDataRelationshipsWalletData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTExternalPayments201ResponseDataRelationshipsWalletData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTExternalPayments201ResponseDataRelationshipsWalletData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTExternalPayments201ResponseDataRelationshipsWalletData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTExternalPayments201ResponseDataRelationshipsWalletData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTExternalPayments201ResponseDataRelationshipsWalletData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTExternalPayments201ResponseDataRelationshipsWalletData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTExternalPayments201ResponseDataRelationshipsWalletData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTExternalPayments201ResponseDataRelationshipsWalletData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTExternalPayments201ResponseDataRelationshipsWalletData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTExternalPayments201ResponseDataRelationshipsWalletData struct {
	value *POSTExternalPayments201ResponseDataRelationshipsWalletData
	isSet bool
}

func (v NullablePOSTExternalPayments201ResponseDataRelationshipsWalletData) Get() *POSTExternalPayments201ResponseDataRelationshipsWalletData {
	return v.value
}

func (v *NullablePOSTExternalPayments201ResponseDataRelationshipsWalletData) Set(val *POSTExternalPayments201ResponseDataRelationshipsWalletData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTExternalPayments201ResponseDataRelationshipsWalletData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTExternalPayments201ResponseDataRelationshipsWalletData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTExternalPayments201ResponseDataRelationshipsWalletData(val *POSTExternalPayments201ResponseDataRelationshipsWalletData) *NullablePOSTExternalPayments201ResponseDataRelationshipsWalletData {
	return &NullablePOSTExternalPayments201ResponseDataRelationshipsWalletData{value: val, isSet: true}
}

func (v NullablePOSTExternalPayments201ResponseDataRelationshipsWalletData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTExternalPayments201ResponseDataRelationshipsWalletData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}