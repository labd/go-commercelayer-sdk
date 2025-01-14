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

// checks if the AddressDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressDataRelationships{}

// AddressDataRelationships struct for AddressDataRelationships
type AddressDataRelationships struct {
	Geocoder *AddressDataRelationshipsGeocoder `json:"geocoder,omitempty"`
	Events   *AddressDataRelationshipsEvents   `json:"events,omitempty"`
	Tags     *AddressDataRelationshipsTags     `json:"tags,omitempty"`
	Versions *AddressDataRelationshipsVersions `json:"versions,omitempty"`
}

// NewAddressDataRelationships instantiates a new AddressDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressDataRelationships() *AddressDataRelationships {
	this := AddressDataRelationships{}
	return &this
}

// NewAddressDataRelationshipsWithDefaults instantiates a new AddressDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressDataRelationshipsWithDefaults() *AddressDataRelationships {
	this := AddressDataRelationships{}
	return &this
}

// GetGeocoder returns the Geocoder field value if set, zero value otherwise.
func (o *AddressDataRelationships) GetGeocoder() AddressDataRelationshipsGeocoder {
	if o == nil || IsNil(o.Geocoder) {
		var ret AddressDataRelationshipsGeocoder
		return ret
	}
	return *o.Geocoder
}

// GetGeocoderOk returns a tuple with the Geocoder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataRelationships) GetGeocoderOk() (*AddressDataRelationshipsGeocoder, bool) {
	if o == nil || IsNil(o.Geocoder) {
		return nil, false
	}
	return o.Geocoder, true
}

// HasGeocoder returns a boolean if a field has been set.
func (o *AddressDataRelationships) HasGeocoder() bool {
	if o != nil && !IsNil(o.Geocoder) {
		return true
	}

	return false
}

// SetGeocoder gets a reference to the given AddressDataRelationshipsGeocoder and assigns it to the Geocoder field.
func (o *AddressDataRelationships) SetGeocoder(v AddressDataRelationshipsGeocoder) {
	o.Geocoder = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *AddressDataRelationships) GetEvents() AddressDataRelationshipsEvents {
	if o == nil || IsNil(o.Events) {
		var ret AddressDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataRelationships) GetEventsOk() (*AddressDataRelationshipsEvents, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *AddressDataRelationships) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given AddressDataRelationshipsEvents and assigns it to the Events field.
func (o *AddressDataRelationships) SetEvents(v AddressDataRelationshipsEvents) {
	o.Events = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AddressDataRelationships) GetTags() AddressDataRelationshipsTags {
	if o == nil || IsNil(o.Tags) {
		var ret AddressDataRelationshipsTags
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataRelationships) GetTagsOk() (*AddressDataRelationshipsTags, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AddressDataRelationships) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given AddressDataRelationshipsTags and assigns it to the Tags field.
func (o *AddressDataRelationships) SetTags(v AddressDataRelationshipsTags) {
	o.Tags = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *AddressDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *AddressDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *AddressDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

func (o AddressDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressDataRelationships) ToMap() (map[string]interface{}, error) {
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

type NullableAddressDataRelationships struct {
	value *AddressDataRelationships
	isSet bool
}

func (v NullableAddressDataRelationships) Get() *AddressDataRelationships {
	return v.value
}

func (v *NullableAddressDataRelationships) Set(val *AddressDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressDataRelationships(val *AddressDataRelationships) *NullableAddressDataRelationships {
	return &NullableAddressDataRelationships{value: val, isSet: true}
}

func (v NullableAddressDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
