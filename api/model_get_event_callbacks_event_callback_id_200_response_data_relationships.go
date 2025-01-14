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

// checks if the GETEventCallbacksEventCallbackId200ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETEventCallbacksEventCallbackId200ResponseDataRelationships{}

// GETEventCallbacksEventCallbackId200ResponseDataRelationships struct for GETEventCallbacksEventCallbackId200ResponseDataRelationships
type GETEventCallbacksEventCallbackId200ResponseDataRelationships struct {
	Webhook *GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhook `json:"webhook,omitempty"`
}

// NewGETEventCallbacksEventCallbackId200ResponseDataRelationships instantiates a new GETEventCallbacksEventCallbackId200ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETEventCallbacksEventCallbackId200ResponseDataRelationships() *GETEventCallbacksEventCallbackId200ResponseDataRelationships {
	this := GETEventCallbacksEventCallbackId200ResponseDataRelationships{}
	return &this
}

// NewGETEventCallbacksEventCallbackId200ResponseDataRelationshipsWithDefaults instantiates a new GETEventCallbacksEventCallbackId200ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETEventCallbacksEventCallbackId200ResponseDataRelationshipsWithDefaults() *GETEventCallbacksEventCallbackId200ResponseDataRelationships {
	this := GETEventCallbacksEventCallbackId200ResponseDataRelationships{}
	return &this
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *GETEventCallbacksEventCallbackId200ResponseDataRelationships) GetWebhook() GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhook {
	if o == nil || IsNil(o.Webhook) {
		var ret GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhook
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETEventCallbacksEventCallbackId200ResponseDataRelationships) GetWebhookOk() (*GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhook, bool) {
	if o == nil || IsNil(o.Webhook) {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *GETEventCallbacksEventCallbackId200ResponseDataRelationships) HasWebhook() bool {
	if o != nil && !IsNil(o.Webhook) {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhook and assigns it to the Webhook field.
func (o *GETEventCallbacksEventCallbackId200ResponseDataRelationships) SetWebhook(v GETEventCallbacksEventCallbackId200ResponseDataRelationshipsWebhook) {
	o.Webhook = &v
}

func (o GETEventCallbacksEventCallbackId200ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETEventCallbacksEventCallbackId200ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Webhook) {
		toSerialize["webhook"] = o.Webhook
	}
	return toSerialize, nil
}

type NullableGETEventCallbacksEventCallbackId200ResponseDataRelationships struct {
	value *GETEventCallbacksEventCallbackId200ResponseDataRelationships
	isSet bool
}

func (v NullableGETEventCallbacksEventCallbackId200ResponseDataRelationships) Get() *GETEventCallbacksEventCallbackId200ResponseDataRelationships {
	return v.value
}

func (v *NullableGETEventCallbacksEventCallbackId200ResponseDataRelationships) Set(val *GETEventCallbacksEventCallbackId200ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETEventCallbacksEventCallbackId200ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETEventCallbacksEventCallbackId200ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETEventCallbacksEventCallbackId200ResponseDataRelationships(val *GETEventCallbacksEventCallbackId200ResponseDataRelationships) *NullableGETEventCallbacksEventCallbackId200ResponseDataRelationships {
	return &NullableGETEventCallbacksEventCallbackId200ResponseDataRelationships{value: val, isSet: true}
}

func (v NullableGETEventCallbacksEventCallbackId200ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETEventCallbacksEventCallbackId200ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
