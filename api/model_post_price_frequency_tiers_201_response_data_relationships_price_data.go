/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.5.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData{}

// POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData struct for POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData
type POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData instantiates a new POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData() *POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData {
	this := POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData{}
	return &this
}

// NewPOSTPriceFrequencyTiers201ResponseDataRelationshipsPriceDataWithDefaults instantiates a new POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPriceFrequencyTiers201ResponseDataRelationshipsPriceDataWithDefaults() *POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData {
	this := POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData struct {
	value *POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData
	isSet bool
}

func (v NullablePOSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData) Get() *POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData {
	return v.value
}

func (v *NullablePOSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData) Set(val *POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData(val *POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData) *NullablePOSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData {
	return &NullablePOSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData{value: val, isSet: true}
}

func (v NullablePOSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
