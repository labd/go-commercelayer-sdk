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

// checks if the POSTAddresses201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTAddresses201ResponseDataRelationships{}

// POSTAddresses201ResponseDataRelationships struct for POSTAddresses201ResponseDataRelationships
type POSTAddresses201ResponseDataRelationships struct {
	Geocoder *POSTAddresses201ResponseDataRelationshipsGeocoder `json:"geocoder,omitempty"`
	Events   *POSTAddresses201ResponseDataRelationshipsEvents   `json:"events,omitempty"`
	Tags     *POSTAddresses201ResponseDataRelationshipsTags     `json:"tags,omitempty"`
	Versions *POSTAddresses201ResponseDataRelationshipsVersions `json:"versions,omitempty"`
}

// NewPOSTAddresses201ResponseDataRelationships instantiates a new POSTAddresses201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAddresses201ResponseDataRelationships() *POSTAddresses201ResponseDataRelationships {
	this := POSTAddresses201ResponseDataRelationships{}
	return &this
}

// NewPOSTAddresses201ResponseDataRelationshipsWithDefaults instantiates a new POSTAddresses201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAddresses201ResponseDataRelationshipsWithDefaults() *POSTAddresses201ResponseDataRelationships {
	this := POSTAddresses201ResponseDataRelationships{}
	return &this
}

// GetGeocoder returns the Geocoder field value if set, zero value otherwise.
func (o *POSTAddresses201ResponseDataRelationships) GetGeocoder() POSTAddresses201ResponseDataRelationshipsGeocoder {
	if o == nil || IsNil(o.Geocoder) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoder
		return ret
	}
	return *o.Geocoder
}

// GetGeocoderOk returns a tuple with the Geocoder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAddresses201ResponseDataRelationships) GetGeocoderOk() (*POSTAddresses201ResponseDataRelationshipsGeocoder, bool) {
	if o == nil || IsNil(o.Geocoder) {
		return nil, false
	}
	return o.Geocoder, true
}

// HasGeocoder returns a boolean if a field has been set.
func (o *POSTAddresses201ResponseDataRelationships) HasGeocoder() bool {
	if o != nil && !IsNil(o.Geocoder) {
		return true
	}

	return false
}

// SetGeocoder gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoder and assigns it to the Geocoder field.
func (o *POSTAddresses201ResponseDataRelationships) SetGeocoder(v POSTAddresses201ResponseDataRelationshipsGeocoder) {
	o.Geocoder = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *POSTAddresses201ResponseDataRelationships) GetEvents() POSTAddresses201ResponseDataRelationshipsEvents {
	if o == nil || IsNil(o.Events) {
		var ret POSTAddresses201ResponseDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAddresses201ResponseDataRelationships) GetEventsOk() (*POSTAddresses201ResponseDataRelationshipsEvents, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *POSTAddresses201ResponseDataRelationships) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given POSTAddresses201ResponseDataRelationshipsEvents and assigns it to the Events field.
func (o *POSTAddresses201ResponseDataRelationships) SetEvents(v POSTAddresses201ResponseDataRelationshipsEvents) {
	o.Events = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *POSTAddresses201ResponseDataRelationships) GetTags() POSTAddresses201ResponseDataRelationshipsTags {
	if o == nil || IsNil(o.Tags) {
		var ret POSTAddresses201ResponseDataRelationshipsTags
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAddresses201ResponseDataRelationships) GetTagsOk() (*POSTAddresses201ResponseDataRelationshipsTags, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *POSTAddresses201ResponseDataRelationships) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given POSTAddresses201ResponseDataRelationshipsTags and assigns it to the Tags field.
func (o *POSTAddresses201ResponseDataRelationships) SetTags(v POSTAddresses201ResponseDataRelationshipsTags) {
	o.Tags = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *POSTAddresses201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAddresses201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *POSTAddresses201ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *POSTAddresses201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

func (o POSTAddresses201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTAddresses201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Geocoder) {
		toSerialize["geocoder"] = o.Geocoder
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullablePOSTAddresses201ResponseDataRelationships struct {
	value *POSTAddresses201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTAddresses201ResponseDataRelationships) Get() *POSTAddresses201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTAddresses201ResponseDataRelationships) Set(val *POSTAddresses201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAddresses201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAddresses201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAddresses201ResponseDataRelationships(val *POSTAddresses201ResponseDataRelationships) *NullablePOSTAddresses201ResponseDataRelationships {
	return &NullablePOSTAddresses201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTAddresses201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAddresses201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
