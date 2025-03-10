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

// checks if the GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData{}

// GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData struct for GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData
type GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewGETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData instantiates a new GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData() *GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData {
	this := GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData{}
	return &this
}

// NewGETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookDataWithDefaults instantiates a new GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookDataWithDefaults() *GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData {
	this := GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData) SetId(v interface{}) {
	o.Id = v
}

func (o GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableGETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData struct {
	value *GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData
	isSet bool
}

func (v NullableGETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData) Get() *GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData {
	return v.value
}

func (v *NullableGETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData) Set(val *GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData(val *GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData) *NullableGETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData {
	return &NullableGETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData{value: val, isSet: true}
}

func (v NullableGETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhookData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
