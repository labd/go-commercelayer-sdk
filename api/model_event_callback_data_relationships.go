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

// EventCallbackDataRelationships struct for EventCallbackDataRelationships
type EventCallbackDataRelationships struct {
	Webhook *EventCallbackDataRelationshipsWebhook `json:"webhook,omitempty"`
}

// NewEventCallbackDataRelationships instantiates a new EventCallbackDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventCallbackDataRelationships() *EventCallbackDataRelationships {
	this := EventCallbackDataRelationships{}
	return &this
}

// NewEventCallbackDataRelationshipsWithDefaults instantiates a new EventCallbackDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventCallbackDataRelationshipsWithDefaults() *EventCallbackDataRelationships {
	this := EventCallbackDataRelationships{}
	return &this
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *EventCallbackDataRelationships) GetWebhook() EventCallbackDataRelationshipsWebhook {
	if o == nil || o.Webhook == nil {
		var ret EventCallbackDataRelationshipsWebhook
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCallbackDataRelationships) GetWebhookOk() (*EventCallbackDataRelationshipsWebhook, bool) {
	if o == nil || o.Webhook == nil {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *EventCallbackDataRelationships) HasWebhook() bool {
	if o != nil && o.Webhook != nil {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given EventCallbackDataRelationshipsWebhook and assigns it to the Webhook field.
func (o *EventCallbackDataRelationships) SetWebhook(v EventCallbackDataRelationshipsWebhook) {
	o.Webhook = &v
}

func (o EventCallbackDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Webhook != nil {
		toSerialize["webhook"] = o.Webhook
	}
	return json.Marshal(toSerialize)
}

type NullableEventCallbackDataRelationships struct {
	value *EventCallbackDataRelationships
	isSet bool
}

func (v NullableEventCallbackDataRelationships) Get() *EventCallbackDataRelationships {
	return v.value
}

func (v *NullableEventCallbackDataRelationships) Set(val *EventCallbackDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableEventCallbackDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableEventCallbackDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventCallbackDataRelationships(val *EventCallbackDataRelationships) *NullableEventCallbackDataRelationships {
	return &NullableEventCallbackDataRelationships{value: val, isSet: true}
}

func (v NullableEventCallbackDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventCallbackDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
