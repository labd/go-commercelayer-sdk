/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETEventCallbacks200ResponseDataInnerRelationshipsWebhookData struct for GETEventCallbacks200ResponseDataInnerRelationshipsWebhookData
type GETEventCallbacks200ResponseDataInnerRelationshipsWebhookData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETEventCallbacks200ResponseDataInnerRelationshipsWebhookData instantiates a new GETEventCallbacks200ResponseDataInnerRelationshipsWebhookData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETEventCallbacks200ResponseDataInnerRelationshipsWebhookData() *GETEventCallbacks200ResponseDataInnerRelationshipsWebhookData {
	this := GETEventCallbacks200ResponseDataInnerRelationshipsWebhookData{}
	return &this
}

// NewGETEventCallbacks200ResponseDataInnerRelationshipsWebhookDataWithDefaults instantiates a new GETEventCallbacks200ResponseDataInnerRelationshipsWebhookData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETEventCallbacks200ResponseDataInnerRelationshipsWebhookDataWithDefaults() *GETEventCallbacks200ResponseDataInnerRelationshipsWebhookData {
	this := GETEventCallbacks200ResponseDataInnerRelationshipsWebhookData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETEventCallbacks200ResponseDataInnerRelationshipsWebhookData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETEventCallbacks200ResponseDataInnerRelationshipsWebhookData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETEventCallbacks200ResponseDataInnerRelationshipsWebhookData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETEventCallbacks200ResponseDataInnerRelationshipsWebhookData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETEventCallbacks200ResponseDataInnerRelationshipsWebhookData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETEventCallbacks200ResponseDataInnerRelationshipsWebhookData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETEventCallbacks200ResponseDataInnerRelationshipsWebhookData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETEventCallbacks200ResponseDataInnerRelationshipsWebhookData) SetId(v string) {
	o.Id = &v
}

func (o GETEventCallbacks200ResponseDataInnerRelationshipsWebhookData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETEventCallbacks200ResponseDataInnerRelationshipsWebhookData struct {
	value *GETEventCallbacks200ResponseDataInnerRelationshipsWebhookData
	isSet bool
}

func (v NullableGETEventCallbacks200ResponseDataInnerRelationshipsWebhookData) Get() *GETEventCallbacks200ResponseDataInnerRelationshipsWebhookData {
	return v.value
}

func (v *NullableGETEventCallbacks200ResponseDataInnerRelationshipsWebhookData) Set(val *GETEventCallbacks200ResponseDataInnerRelationshipsWebhookData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETEventCallbacks200ResponseDataInnerRelationshipsWebhookData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETEventCallbacks200ResponseDataInnerRelationshipsWebhookData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETEventCallbacks200ResponseDataInnerRelationshipsWebhookData(val *GETEventCallbacks200ResponseDataInnerRelationshipsWebhookData) *NullableGETEventCallbacks200ResponseDataInnerRelationshipsWebhookData {
	return &NullableGETEventCallbacks200ResponseDataInnerRelationshipsWebhookData{value: val, isSet: true}
}

func (v NullableGETEventCallbacks200ResponseDataInnerRelationshipsWebhookData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETEventCallbacks200ResponseDataInnerRelationshipsWebhookData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
