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

// checks if the POSTCustomers201ResponseDataRelationshipsCustomerGroupData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCustomers201ResponseDataRelationshipsCustomerGroupData{}

// POSTCustomers201ResponseDataRelationshipsCustomerGroupData struct for POSTCustomers201ResponseDataRelationshipsCustomerGroupData
type POSTCustomers201ResponseDataRelationshipsCustomerGroupData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTCustomers201ResponseDataRelationshipsCustomerGroupData instantiates a new POSTCustomers201ResponseDataRelationshipsCustomerGroupData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCustomers201ResponseDataRelationshipsCustomerGroupData() *POSTCustomers201ResponseDataRelationshipsCustomerGroupData {
	this := POSTCustomers201ResponseDataRelationshipsCustomerGroupData{}
	return &this
}

// NewPOSTCustomers201ResponseDataRelationshipsCustomerGroupDataWithDefaults instantiates a new POSTCustomers201ResponseDataRelationshipsCustomerGroupData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCustomers201ResponseDataRelationshipsCustomerGroupDataWithDefaults() *POSTCustomers201ResponseDataRelationshipsCustomerGroupData {
	this := POSTCustomers201ResponseDataRelationshipsCustomerGroupData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTCustomers201ResponseDataRelationshipsCustomerGroupData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCustomers201ResponseDataRelationshipsCustomerGroupData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerGroupData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerGroupData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTCustomers201ResponseDataRelationshipsCustomerGroupData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCustomers201ResponseDataRelationshipsCustomerGroupData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerGroupData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerGroupData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTCustomers201ResponseDataRelationshipsCustomerGroupData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCustomers201ResponseDataRelationshipsCustomerGroupData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTCustomers201ResponseDataRelationshipsCustomerGroupData struct {
	value *POSTCustomers201ResponseDataRelationshipsCustomerGroupData
	isSet bool
}

func (v NullablePOSTCustomers201ResponseDataRelationshipsCustomerGroupData) Get() *POSTCustomers201ResponseDataRelationshipsCustomerGroupData {
	return v.value
}

func (v *NullablePOSTCustomers201ResponseDataRelationshipsCustomerGroupData) Set(val *POSTCustomers201ResponseDataRelationshipsCustomerGroupData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCustomers201ResponseDataRelationshipsCustomerGroupData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCustomers201ResponseDataRelationshipsCustomerGroupData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCustomers201ResponseDataRelationshipsCustomerGroupData(val *POSTCustomers201ResponseDataRelationshipsCustomerGroupData) *NullablePOSTCustomers201ResponseDataRelationshipsCustomerGroupData {
	return &NullablePOSTCustomers201ResponseDataRelationshipsCustomerGroupData{value: val, isSet: true}
}

func (v NullablePOSTCustomers201ResponseDataRelationshipsCustomerGroupData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCustomers201ResponseDataRelationshipsCustomerGroupData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
