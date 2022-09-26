/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// WebhookDataRelationships struct for WebhookDataRelationships
type WebhookDataRelationships struct {
	LastEventCallbacks *EventDataRelationshipsLastEventCallbacks `json:"last_event_callbacks,omitempty"`
}

// NewWebhookDataRelationships instantiates a new WebhookDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookDataRelationships() *WebhookDataRelationships {
	this := WebhookDataRelationships{}
	return &this
}

// NewWebhookDataRelationshipsWithDefaults instantiates a new WebhookDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookDataRelationshipsWithDefaults() *WebhookDataRelationships {
	this := WebhookDataRelationships{}
	return &this
}

// GetLastEventCallbacks returns the LastEventCallbacks field value if set, zero value otherwise.
func (o *WebhookDataRelationships) GetLastEventCallbacks() EventDataRelationshipsLastEventCallbacks {
	if o == nil || o.LastEventCallbacks == nil {
		var ret EventDataRelationshipsLastEventCallbacks
		return ret
	}
	return *o.LastEventCallbacks
}

// GetLastEventCallbacksOk returns a tuple with the LastEventCallbacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookDataRelationships) GetLastEventCallbacksOk() (*EventDataRelationshipsLastEventCallbacks, bool) {
	if o == nil || o.LastEventCallbacks == nil {
		return nil, false
	}
	return o.LastEventCallbacks, true
}

// HasLastEventCallbacks returns a boolean if a field has been set.
func (o *WebhookDataRelationships) HasLastEventCallbacks() bool {
	if o != nil && o.LastEventCallbacks != nil {
		return true
	}

	return false
}

// SetLastEventCallbacks gets a reference to the given EventDataRelationshipsLastEventCallbacks and assigns it to the LastEventCallbacks field.
func (o *WebhookDataRelationships) SetLastEventCallbacks(v EventDataRelationshipsLastEventCallbacks) {
	o.LastEventCallbacks = &v
}

func (o WebhookDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastEventCallbacks != nil {
		toSerialize["last_event_callbacks"] = o.LastEventCallbacks
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookDataRelationships struct {
	value *WebhookDataRelationships
	isSet bool
}

func (v NullableWebhookDataRelationships) Get() *WebhookDataRelationships {
	return v.value
}

func (v *NullableWebhookDataRelationships) Set(val *WebhookDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookDataRelationships(val *WebhookDataRelationships) *NullableWebhookDataRelationships {
	return &NullableWebhookDataRelationships{value: val, isSet: true}
}

func (v NullableWebhookDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
