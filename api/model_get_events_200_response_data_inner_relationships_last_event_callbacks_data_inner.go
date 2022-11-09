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

// GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner struct for GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner
type GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner instantiates a new GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner() *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner {
	this := GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner{}
	return &this
}

// NewGETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInnerWithDefaults instantiates a new GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInnerWithDefaults() *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner {
	this := GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner) SetId(v string) {
	o.Id = &v
}

func (o GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner struct {
	value *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner
	isSet bool
}

func (v NullableGETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner) Get() *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner {
	return v.value
}

func (v *NullableGETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner) Set(val *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner(val *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner) *NullableGETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner {
	return &NullableGETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner{value: val, isSet: true}
}

func (v NullableGETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
