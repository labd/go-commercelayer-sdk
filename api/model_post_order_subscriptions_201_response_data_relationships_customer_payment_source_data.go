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

// checks if the POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData{}

// POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData struct for POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData
type POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData instantiates a new POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData() *POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData {
	this := POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData{}
	return &this
}

// NewPOSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceDataWithDefaults instantiates a new POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceDataWithDefaults() *POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData {
	this := POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData struct {
	value *POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData
	isSet bool
}

func (v NullablePOSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData) Get() *POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData {
	return v.value
}

func (v *NullablePOSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData) Set(val *POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData(val *POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData) *NullablePOSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData {
	return &NullablePOSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData{value: val, isSet: true}
}

func (v NullablePOSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSourceData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}