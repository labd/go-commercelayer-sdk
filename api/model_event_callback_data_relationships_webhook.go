/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.5.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the EventCallbackDataRelationshipsWebhook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventCallbackDataRelationshipsWebhook{}

// EventCallbackDataRelationshipsWebhook struct for EventCallbackDataRelationshipsWebhook
type EventCallbackDataRelationshipsWebhook struct {
	Data *EventCallbackDataRelationshipsWebhookData `json:"data,omitempty"`
}

// NewEventCallbackDataRelationshipsWebhook instantiates a new EventCallbackDataRelationshipsWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventCallbackDataRelationshipsWebhook() *EventCallbackDataRelationshipsWebhook {
	this := EventCallbackDataRelationshipsWebhook{}
	return &this
}

// NewEventCallbackDataRelationshipsWebhookWithDefaults instantiates a new EventCallbackDataRelationshipsWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventCallbackDataRelationshipsWebhookWithDefaults() *EventCallbackDataRelationshipsWebhook {
	this := EventCallbackDataRelationshipsWebhook{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EventCallbackDataRelationshipsWebhook) GetData() EventCallbackDataRelationshipsWebhookData {
	if o == nil || IsNil(o.Data) {
		var ret EventCallbackDataRelationshipsWebhookData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCallbackDataRelationshipsWebhook) GetDataOk() (*EventCallbackDataRelationshipsWebhookData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EventCallbackDataRelationshipsWebhook) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given EventCallbackDataRelationshipsWebhookData and assigns it to the Data field.
func (o *EventCallbackDataRelationshipsWebhook) SetData(v EventCallbackDataRelationshipsWebhookData) {
	o.Data = &v
}

func (o EventCallbackDataRelationshipsWebhook) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventCallbackDataRelationshipsWebhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableEventCallbackDataRelationshipsWebhook struct {
	value *EventCallbackDataRelationshipsWebhook
	isSet bool
}

func (v NullableEventCallbackDataRelationshipsWebhook) Get() *EventCallbackDataRelationshipsWebhook {
	return v.value
}

func (v *NullableEventCallbackDataRelationshipsWebhook) Set(val *EventCallbackDataRelationshipsWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableEventCallbackDataRelationshipsWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableEventCallbackDataRelationshipsWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventCallbackDataRelationshipsWebhook(val *EventCallbackDataRelationshipsWebhook) *NullableEventCallbackDataRelationshipsWebhook {
	return &NullableEventCallbackDataRelationshipsWebhook{value: val, isSet: true}
}

func (v NullableEventCallbackDataRelationshipsWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventCallbackDataRelationshipsWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
