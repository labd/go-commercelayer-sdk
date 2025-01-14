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

// checks if the POSTPrices201ResponseDataRelationshipsPriceTiersData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTPrices201ResponseDataRelationshipsPriceTiersData{}

// POSTPrices201ResponseDataRelationshipsPriceTiersData struct for POSTPrices201ResponseDataRelationshipsPriceTiersData
type POSTPrices201ResponseDataRelationshipsPriceTiersData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTPrices201ResponseDataRelationshipsPriceTiersData instantiates a new POSTPrices201ResponseDataRelationshipsPriceTiersData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPrices201ResponseDataRelationshipsPriceTiersData() *POSTPrices201ResponseDataRelationshipsPriceTiersData {
	this := POSTPrices201ResponseDataRelationshipsPriceTiersData{}
	return &this
}

// NewPOSTPrices201ResponseDataRelationshipsPriceTiersDataWithDefaults instantiates a new POSTPrices201ResponseDataRelationshipsPriceTiersData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPrices201ResponseDataRelationshipsPriceTiersDataWithDefaults() *POSTPrices201ResponseDataRelationshipsPriceTiersData {
	this := POSTPrices201ResponseDataRelationshipsPriceTiersData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPrices201ResponseDataRelationshipsPriceTiersData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPrices201ResponseDataRelationshipsPriceTiersData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTPrices201ResponseDataRelationshipsPriceTiersData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTPrices201ResponseDataRelationshipsPriceTiersData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPrices201ResponseDataRelationshipsPriceTiersData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPrices201ResponseDataRelationshipsPriceTiersData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTPrices201ResponseDataRelationshipsPriceTiersData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTPrices201ResponseDataRelationshipsPriceTiersData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTPrices201ResponseDataRelationshipsPriceTiersData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTPrices201ResponseDataRelationshipsPriceTiersData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTPrices201ResponseDataRelationshipsPriceTiersData struct {
	value *POSTPrices201ResponseDataRelationshipsPriceTiersData
	isSet bool
}

func (v NullablePOSTPrices201ResponseDataRelationshipsPriceTiersData) Get() *POSTPrices201ResponseDataRelationshipsPriceTiersData {
	return v.value
}

func (v *NullablePOSTPrices201ResponseDataRelationshipsPriceTiersData) Set(val *POSTPrices201ResponseDataRelationshipsPriceTiersData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPrices201ResponseDataRelationshipsPriceTiersData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPrices201ResponseDataRelationshipsPriceTiersData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPrices201ResponseDataRelationshipsPriceTiersData(val *POSTPrices201ResponseDataRelationshipsPriceTiersData) *NullablePOSTPrices201ResponseDataRelationshipsPriceTiersData {
	return &NullablePOSTPrices201ResponseDataRelationshipsPriceTiersData{value: val, isSet: true}
}

func (v NullablePOSTPrices201ResponseDataRelationshipsPriceTiersData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPrices201ResponseDataRelationshipsPriceTiersData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
