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

// checks if the GETCapturesCaptureId200ResponseDataRelationshipsReturnData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETCapturesCaptureId200ResponseDataRelationshipsReturnData{}

// GETCapturesCaptureId200ResponseDataRelationshipsReturnData struct for GETCapturesCaptureId200ResponseDataRelationshipsReturnData
type GETCapturesCaptureId200ResponseDataRelationshipsReturnData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewGETCapturesCaptureId200ResponseDataRelationshipsReturnData instantiates a new GETCapturesCaptureId200ResponseDataRelationshipsReturnData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCapturesCaptureId200ResponseDataRelationshipsReturnData() *GETCapturesCaptureId200ResponseDataRelationshipsReturnData {
	this := GETCapturesCaptureId200ResponseDataRelationshipsReturnData{}
	return &this
}

// NewGETCapturesCaptureId200ResponseDataRelationshipsReturnDataWithDefaults instantiates a new GETCapturesCaptureId200ResponseDataRelationshipsReturnData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCapturesCaptureId200ResponseDataRelationshipsReturnDataWithDefaults() *GETCapturesCaptureId200ResponseDataRelationshipsReturnData {
	this := GETCapturesCaptureId200ResponseDataRelationshipsReturnData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCapturesCaptureId200ResponseDataRelationshipsReturnData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCapturesCaptureId200ResponseDataRelationshipsReturnData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataRelationshipsReturnData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *GETCapturesCaptureId200ResponseDataRelationshipsReturnData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCapturesCaptureId200ResponseDataRelationshipsReturnData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCapturesCaptureId200ResponseDataRelationshipsReturnData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataRelationshipsReturnData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *GETCapturesCaptureId200ResponseDataRelationshipsReturnData) SetId(v interface{}) {
	o.Id = v
}

func (o GETCapturesCaptureId200ResponseDataRelationshipsReturnData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETCapturesCaptureId200ResponseDataRelationshipsReturnData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableGETCapturesCaptureId200ResponseDataRelationshipsReturnData struct {
	value *GETCapturesCaptureId200ResponseDataRelationshipsReturnData
	isSet bool
}

func (v NullableGETCapturesCaptureId200ResponseDataRelationshipsReturnData) Get() *GETCapturesCaptureId200ResponseDataRelationshipsReturnData {
	return v.value
}

func (v *NullableGETCapturesCaptureId200ResponseDataRelationshipsReturnData) Set(val *GETCapturesCaptureId200ResponseDataRelationshipsReturnData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCapturesCaptureId200ResponseDataRelationshipsReturnData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCapturesCaptureId200ResponseDataRelationshipsReturnData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCapturesCaptureId200ResponseDataRelationshipsReturnData(val *GETCapturesCaptureId200ResponseDataRelationshipsReturnData) *NullableGETCapturesCaptureId200ResponseDataRelationshipsReturnData {
	return &NullableGETCapturesCaptureId200ResponseDataRelationshipsReturnData{value: val, isSet: true}
}

func (v NullableGETCapturesCaptureId200ResponseDataRelationshipsReturnData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCapturesCaptureId200ResponseDataRelationshipsReturnData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
