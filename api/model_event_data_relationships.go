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

// checks if the EventDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventDataRelationships{}

// EventDataRelationships struct for EventDataRelationships
type EventDataRelationships struct {
	Webhooks           *EventCallbackDataRelationshipsWebhook    `json:"webhooks,omitempty"`
	LastEventCallbacks *EventDataRelationshipsLastEventCallbacks `json:"last_event_callbacks,omitempty"`
}

// NewEventDataRelationships instantiates a new EventDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventDataRelationships() *EventDataRelationships {
	this := EventDataRelationships{}
	return &this
}

// NewEventDataRelationshipsWithDefaults instantiates a new EventDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventDataRelationshipsWithDefaults() *EventDataRelationships {
	this := EventDataRelationships{}
	return &this
}

// GetWebhooks returns the Webhooks field value if set, zero value otherwise.
func (o *EventDataRelationships) GetWebhooks() EventCallbackDataRelationshipsWebhook {
	if o == nil || IsNil(o.Webhooks) {
		var ret EventCallbackDataRelationshipsWebhook
		return ret
	}
	return *o.Webhooks
}

// GetWebhooksOk returns a tuple with the Webhooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDataRelationships) GetWebhooksOk() (*EventCallbackDataRelationshipsWebhook, bool) {
	if o == nil || IsNil(o.Webhooks) {
		return nil, false
	}
	return o.Webhooks, true
}

// HasWebhooks returns a boolean if a field has been set.
func (o *EventDataRelationships) HasWebhooks() bool {
	if o != nil && !IsNil(o.Webhooks) {
		return true
	}

	return false
}

// SetWebhooks gets a reference to the given EventCallbackDataRelationshipsWebhook and assigns it to the Webhooks field.
func (o *EventDataRelationships) SetWebhooks(v EventCallbackDataRelationshipsWebhook) {
	o.Webhooks = &v
}

// GetLastEventCallbacks returns the LastEventCallbacks field value if set, zero value otherwise.
func (o *EventDataRelationships) GetLastEventCallbacks() EventDataRelationshipsLastEventCallbacks {
	if o == nil || IsNil(o.LastEventCallbacks) {
		var ret EventDataRelationshipsLastEventCallbacks
		return ret
	}
	return *o.LastEventCallbacks
}

// GetLastEventCallbacksOk returns a tuple with the LastEventCallbacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDataRelationships) GetLastEventCallbacksOk() (*EventDataRelationshipsLastEventCallbacks, bool) {
	if o == nil || IsNil(o.LastEventCallbacks) {
		return nil, false
	}
	return o.LastEventCallbacks, true
}

// HasLastEventCallbacks returns a boolean if a field has been set.
func (o *EventDataRelationships) HasLastEventCallbacks() bool {
	if o != nil && !IsNil(o.LastEventCallbacks) {
		return true
	}

	return false
}

// SetLastEventCallbacks gets a reference to the given EventDataRelationshipsLastEventCallbacks and assigns it to the LastEventCallbacks field.
func (o *EventDataRelationships) SetLastEventCallbacks(v EventDataRelationshipsLastEventCallbacks) {
	o.LastEventCallbacks = &v
}

func (o EventDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Webhooks) {
		toSerialize["webhooks"] = o.Webhooks
	}
	if !IsNil(o.LastEventCallbacks) {
		toSerialize["last_event_callbacks"] = o.LastEventCallbacks
	}
	return toSerialize, nil
}

type NullableEventDataRelationships struct {
	value *EventDataRelationships
	isSet bool
}

func (v NullableEventDataRelationships) Get() *EventDataRelationships {
	return v.value
}

func (v *NullableEventDataRelationships) Set(val *EventDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableEventDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableEventDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventDataRelationships(val *EventDataRelationships) *NullableEventDataRelationships {
	return &NullableEventDataRelationships{value: val, isSet: true}
}

func (v NullableEventDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
