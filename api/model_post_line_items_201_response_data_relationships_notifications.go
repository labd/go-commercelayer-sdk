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

// checks if the POSTLineItems201ResponseDataRelationshipsNotifications type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTLineItems201ResponseDataRelationshipsNotifications{}

// POSTLineItems201ResponseDataRelationshipsNotifications struct for POSTLineItems201ResponseDataRelationshipsNotifications
type POSTLineItems201ResponseDataRelationshipsNotifications struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks     `json:"links,omitempty"`
	Data  *POSTLineItems201ResponseDataRelationshipsNotificationsData `json:"data,omitempty"`
}

// NewPOSTLineItems201ResponseDataRelationshipsNotifications instantiates a new POSTLineItems201ResponseDataRelationshipsNotifications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTLineItems201ResponseDataRelationshipsNotifications() *POSTLineItems201ResponseDataRelationshipsNotifications {
	this := POSTLineItems201ResponseDataRelationshipsNotifications{}
	return &this
}

// NewPOSTLineItems201ResponseDataRelationshipsNotificationsWithDefaults instantiates a new POSTLineItems201ResponseDataRelationshipsNotifications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTLineItems201ResponseDataRelationshipsNotificationsWithDefaults() *POSTLineItems201ResponseDataRelationshipsNotifications {
	this := POSTLineItems201ResponseDataRelationshipsNotifications{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataRelationshipsNotifications) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataRelationshipsNotifications) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationshipsNotifications) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTLineItems201ResponseDataRelationshipsNotifications) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataRelationshipsNotifications) GetData() POSTLineItems201ResponseDataRelationshipsNotificationsData {
	if o == nil || IsNil(o.Data) {
		var ret POSTLineItems201ResponseDataRelationshipsNotificationsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataRelationshipsNotifications) GetDataOk() (*POSTLineItems201ResponseDataRelationshipsNotificationsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationshipsNotifications) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTLineItems201ResponseDataRelationshipsNotificationsData and assigns it to the Data field.
func (o *POSTLineItems201ResponseDataRelationshipsNotifications) SetData(v POSTLineItems201ResponseDataRelationshipsNotificationsData) {
	o.Data = &v
}

func (o POSTLineItems201ResponseDataRelationshipsNotifications) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTLineItems201ResponseDataRelationshipsNotifications) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTLineItems201ResponseDataRelationshipsNotifications struct {
	value *POSTLineItems201ResponseDataRelationshipsNotifications
	isSet bool
}

func (v NullablePOSTLineItems201ResponseDataRelationshipsNotifications) Get() *POSTLineItems201ResponseDataRelationshipsNotifications {
	return v.value
}

func (v *NullablePOSTLineItems201ResponseDataRelationshipsNotifications) Set(val *POSTLineItems201ResponseDataRelationshipsNotifications) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTLineItems201ResponseDataRelationshipsNotifications) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTLineItems201ResponseDataRelationshipsNotifications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTLineItems201ResponseDataRelationshipsNotifications(val *POSTLineItems201ResponseDataRelationshipsNotifications) *NullablePOSTLineItems201ResponseDataRelationshipsNotifications {
	return &NullablePOSTLineItems201ResponseDataRelationshipsNotifications{value: val, isSet: true}
}

func (v NullablePOSTLineItems201ResponseDataRelationshipsNotifications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTLineItems201ResponseDataRelationshipsNotifications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}