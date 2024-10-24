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

// checks if the GETEventsEventId200ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETEventsEventId200ResponseDataRelationships{}

// GETEventsEventId200ResponseDataRelationships struct for GETEventsEventId200ResponseDataRelationships
type GETEventsEventId200ResponseDataRelationships struct {
	LastEventCallbacks *GETEventsEventId200ResponseDataRelationshipsLastEventCallbacks `json:"last_event_callbacks,omitempty"`
	Webhooks           *GETEventsEventId200ResponseDataRelationshipsWebhooks           `json:"webhooks,omitempty"`
}

// NewGETEventsEventId200ResponseDataRelationships instantiates a new GETEventsEventId200ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETEventsEventId200ResponseDataRelationships() *GETEventsEventId200ResponseDataRelationships {
	this := GETEventsEventId200ResponseDataRelationships{}
	return &this
}

// NewGETEventsEventId200ResponseDataRelationshipsWithDefaults instantiates a new GETEventsEventId200ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETEventsEventId200ResponseDataRelationshipsWithDefaults() *GETEventsEventId200ResponseDataRelationships {
	this := GETEventsEventId200ResponseDataRelationships{}
	return &this
}

// GetLastEventCallbacks returns the LastEventCallbacks field value if set, zero value otherwise.
func (o *GETEventsEventId200ResponseDataRelationships) GetLastEventCallbacks() GETEventsEventId200ResponseDataRelationshipsLastEventCallbacks {
	if o == nil || IsNil(o.LastEventCallbacks) {
		var ret GETEventsEventId200ResponseDataRelationshipsLastEventCallbacks
		return ret
	}
	return *o.LastEventCallbacks
}

// GetLastEventCallbacksOk returns a tuple with the LastEventCallbacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETEventsEventId200ResponseDataRelationships) GetLastEventCallbacksOk() (*GETEventsEventId200ResponseDataRelationshipsLastEventCallbacks, bool) {
	if o == nil || IsNil(o.LastEventCallbacks) {
		return nil, false
	}
	return o.LastEventCallbacks, true
}

// HasLastEventCallbacks returns a boolean if a field has been set.
func (o *GETEventsEventId200ResponseDataRelationships) HasLastEventCallbacks() bool {
	if o != nil && !IsNil(o.LastEventCallbacks) {
		return true
	}

	return false
}

// SetLastEventCallbacks gets a reference to the given GETEventsEventId200ResponseDataRelationshipsLastEventCallbacks and assigns it to the LastEventCallbacks field.
func (o *GETEventsEventId200ResponseDataRelationships) SetLastEventCallbacks(v GETEventsEventId200ResponseDataRelationshipsLastEventCallbacks) {
	o.LastEventCallbacks = &v
}

// GetWebhooks returns the Webhooks field value if set, zero value otherwise.
func (o *GETEventsEventId200ResponseDataRelationships) GetWebhooks() GETEventsEventId200ResponseDataRelationshipsWebhooks {
	if o == nil || IsNil(o.Webhooks) {
		var ret GETEventsEventId200ResponseDataRelationshipsWebhooks
		return ret
	}
	return *o.Webhooks
}

// GetWebhooksOk returns a tuple with the Webhooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETEventsEventId200ResponseDataRelationships) GetWebhooksOk() (*GETEventsEventId200ResponseDataRelationshipsWebhooks, bool) {
	if o == nil || IsNil(o.Webhooks) {
		return nil, false
	}
	return o.Webhooks, true
}

// HasWebhooks returns a boolean if a field has been set.
func (o *GETEventsEventId200ResponseDataRelationships) HasWebhooks() bool {
	if o != nil && !IsNil(o.Webhooks) {
		return true
	}

	return false
}

// SetWebhooks gets a reference to the given GETEventsEventId200ResponseDataRelationshipsWebhooks and assigns it to the Webhooks field.
func (o *GETEventsEventId200ResponseDataRelationships) SetWebhooks(v GETEventsEventId200ResponseDataRelationshipsWebhooks) {
	o.Webhooks = &v
}

func (o GETEventsEventId200ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETEventsEventId200ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LastEventCallbacks) {
		toSerialize["last_event_callbacks"] = o.LastEventCallbacks
	}
	if !IsNil(o.Webhooks) {
		toSerialize["webhooks"] = o.Webhooks
	}
	return toSerialize, nil
}

type NullableGETEventsEventId200ResponseDataRelationships struct {
	value *GETEventsEventId200ResponseDataRelationships
	isSet bool
}

func (v NullableGETEventsEventId200ResponseDataRelationships) Get() *GETEventsEventId200ResponseDataRelationships {
	return v.value
}

func (v *NullableGETEventsEventId200ResponseDataRelationships) Set(val *GETEventsEventId200ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETEventsEventId200ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETEventsEventId200ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETEventsEventId200ResponseDataRelationships(val *GETEventsEventId200ResponseDataRelationships) *NullableGETEventsEventId200ResponseDataRelationships {
	return &NullableGETEventsEventId200ResponseDataRelationships{value: val, isSet: true}
}

func (v NullableGETEventsEventId200ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETEventsEventId200ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}