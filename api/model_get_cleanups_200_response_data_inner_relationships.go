/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETCleanups200ResponseDataInnerRelationships struct for GETCleanups200ResponseDataInnerRelationships
type GETCleanups200ResponseDataInnerRelationships struct {
	Events *GETCleanups200ResponseDataInnerRelationshipsEvents `json:"events,omitempty"`
}

// NewGETCleanups200ResponseDataInnerRelationships instantiates a new GETCleanups200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCleanups200ResponseDataInnerRelationships() *GETCleanups200ResponseDataInnerRelationships {
	this := GETCleanups200ResponseDataInnerRelationships{}
	return &this
}

// NewGETCleanups200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETCleanups200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCleanups200ResponseDataInnerRelationshipsWithDefaults() *GETCleanups200ResponseDataInnerRelationships {
	this := GETCleanups200ResponseDataInnerRelationships{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *GETCleanups200ResponseDataInnerRelationships) GetEvents() GETCleanups200ResponseDataInnerRelationshipsEvents {
	if o == nil || o.Events == nil {
		var ret GETCleanups200ResponseDataInnerRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCleanups200ResponseDataInnerRelationships) GetEventsOk() (*GETCleanups200ResponseDataInnerRelationshipsEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *GETCleanups200ResponseDataInnerRelationships) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given GETCleanups200ResponseDataInnerRelationshipsEvents and assigns it to the Events field.
func (o *GETCleanups200ResponseDataInnerRelationships) SetEvents(v GETCleanups200ResponseDataInnerRelationshipsEvents) {
	o.Events = &v
}

func (o GETCleanups200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableGETCleanups200ResponseDataInnerRelationships struct {
	value *GETCleanups200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETCleanups200ResponseDataInnerRelationships) Get() *GETCleanups200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETCleanups200ResponseDataInnerRelationships) Set(val *GETCleanups200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCleanups200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCleanups200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCleanups200ResponseDataInnerRelationships(val *GETCleanups200ResponseDataInnerRelationships) *NullableGETCleanups200ResponseDataInnerRelationships {
	return &NullableGETCleanups200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETCleanups200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCleanups200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
