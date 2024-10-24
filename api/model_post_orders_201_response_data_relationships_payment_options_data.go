/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTOrders201ResponseDataRelationshipsPaymentOptionsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrders201ResponseDataRelationshipsPaymentOptionsData{}

// POSTOrders201ResponseDataRelationshipsPaymentOptionsData struct for POSTOrders201ResponseDataRelationshipsPaymentOptionsData
type POSTOrders201ResponseDataRelationshipsPaymentOptionsData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTOrders201ResponseDataRelationshipsPaymentOptionsData instantiates a new POSTOrders201ResponseDataRelationshipsPaymentOptionsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrders201ResponseDataRelationshipsPaymentOptionsData() *POSTOrders201ResponseDataRelationshipsPaymentOptionsData {
	this := POSTOrders201ResponseDataRelationshipsPaymentOptionsData{}
	return &this
}

// NewPOSTOrders201ResponseDataRelationshipsPaymentOptionsDataWithDefaults instantiates a new POSTOrders201ResponseDataRelationshipsPaymentOptionsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrders201ResponseDataRelationshipsPaymentOptionsDataWithDefaults() *POSTOrders201ResponseDataRelationshipsPaymentOptionsData {
	this := POSTOrders201ResponseDataRelationshipsPaymentOptionsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataRelationshipsPaymentOptionsData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataRelationshipsPaymentOptionsData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsPaymentOptionsData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTOrders201ResponseDataRelationshipsPaymentOptionsData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataRelationshipsPaymentOptionsData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataRelationshipsPaymentOptionsData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsPaymentOptionsData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTOrders201ResponseDataRelationshipsPaymentOptionsData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTOrders201ResponseDataRelationshipsPaymentOptionsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrders201ResponseDataRelationshipsPaymentOptionsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTOrders201ResponseDataRelationshipsPaymentOptionsData struct {
	value *POSTOrders201ResponseDataRelationshipsPaymentOptionsData
	isSet bool
}

func (v NullablePOSTOrders201ResponseDataRelationshipsPaymentOptionsData) Get() *POSTOrders201ResponseDataRelationshipsPaymentOptionsData {
	return v.value
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsPaymentOptionsData) Set(val *POSTOrders201ResponseDataRelationshipsPaymentOptionsData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrders201ResponseDataRelationshipsPaymentOptionsData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsPaymentOptionsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrders201ResponseDataRelationshipsPaymentOptionsData(val *POSTOrders201ResponseDataRelationshipsPaymentOptionsData) *NullablePOSTOrders201ResponseDataRelationshipsPaymentOptionsData {
	return &NullablePOSTOrders201ResponseDataRelationshipsPaymentOptionsData{value: val, isSet: true}
}

func (v NullablePOSTOrders201ResponseDataRelationshipsPaymentOptionsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsPaymentOptionsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}