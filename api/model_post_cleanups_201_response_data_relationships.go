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

// checks if the POSTCleanups201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCleanups201ResponseDataRelationships{}

// POSTCleanups201ResponseDataRelationships struct for POSTCleanups201ResponseDataRelationships
type POSTCleanups201ResponseDataRelationships struct {
	Events   *POSTAddresses201ResponseDataRelationshipsEvents   `json:"events,omitempty"`
	Versions *POSTAddresses201ResponseDataRelationshipsVersions `json:"versions,omitempty"`
}

// NewPOSTCleanups201ResponseDataRelationships instantiates a new POSTCleanups201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCleanups201ResponseDataRelationships() *POSTCleanups201ResponseDataRelationships {
	this := POSTCleanups201ResponseDataRelationships{}
	return &this
}

// NewPOSTCleanups201ResponseDataRelationshipsWithDefaults instantiates a new POSTCleanups201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCleanups201ResponseDataRelationshipsWithDefaults() *POSTCleanups201ResponseDataRelationships {
	this := POSTCleanups201ResponseDataRelationships{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *POSTCleanups201ResponseDataRelationships) GetEvents() POSTAddresses201ResponseDataRelationshipsEvents {
	if o == nil || IsNil(o.Events) {
		var ret POSTAddresses201ResponseDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCleanups201ResponseDataRelationships) GetEventsOk() (*POSTAddresses201ResponseDataRelationshipsEvents, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *POSTCleanups201ResponseDataRelationships) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given POSTAddresses201ResponseDataRelationshipsEvents and assigns it to the Events field.
func (o *POSTCleanups201ResponseDataRelationships) SetEvents(v POSTAddresses201ResponseDataRelationshipsEvents) {
	o.Events = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *POSTCleanups201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCleanups201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *POSTCleanups201ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *POSTCleanups201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

func (o POSTCleanups201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCleanups201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullablePOSTCleanups201ResponseDataRelationships struct {
	value *POSTCleanups201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTCleanups201ResponseDataRelationships) Get() *POSTCleanups201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTCleanups201ResponseDataRelationships) Set(val *POSTCleanups201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCleanups201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCleanups201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCleanups201ResponseDataRelationships(val *POSTCleanups201ResponseDataRelationships) *NullablePOSTCleanups201ResponseDataRelationships {
	return &NullablePOSTCleanups201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTCleanups201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCleanups201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
