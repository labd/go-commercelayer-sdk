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

// checks if the POSTNotifications201ResponseDataRelationshipsNotifiable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTNotifications201ResponseDataRelationshipsNotifiable{}

// POSTNotifications201ResponseDataRelationshipsNotifiable struct for POSTNotifications201ResponseDataRelationshipsNotifiable
type POSTNotifications201ResponseDataRelationshipsNotifiable struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks      `json:"links,omitempty"`
	Data  *POSTNotifications201ResponseDataRelationshipsNotifiableData `json:"data,omitempty"`
}

// NewPOSTNotifications201ResponseDataRelationshipsNotifiable instantiates a new POSTNotifications201ResponseDataRelationshipsNotifiable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTNotifications201ResponseDataRelationshipsNotifiable() *POSTNotifications201ResponseDataRelationshipsNotifiable {
	this := POSTNotifications201ResponseDataRelationshipsNotifiable{}
	return &this
}

// NewPOSTNotifications201ResponseDataRelationshipsNotifiableWithDefaults instantiates a new POSTNotifications201ResponseDataRelationshipsNotifiable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTNotifications201ResponseDataRelationshipsNotifiableWithDefaults() *POSTNotifications201ResponseDataRelationshipsNotifiable {
	this := POSTNotifications201ResponseDataRelationshipsNotifiable{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTNotifications201ResponseDataRelationshipsNotifiable) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTNotifications201ResponseDataRelationshipsNotifiable) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTNotifications201ResponseDataRelationshipsNotifiable) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTNotifications201ResponseDataRelationshipsNotifiable) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTNotifications201ResponseDataRelationshipsNotifiable) GetData() POSTNotifications201ResponseDataRelationshipsNotifiableData {
	if o == nil || IsNil(o.Data) {
		var ret POSTNotifications201ResponseDataRelationshipsNotifiableData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTNotifications201ResponseDataRelationshipsNotifiable) GetDataOk() (*POSTNotifications201ResponseDataRelationshipsNotifiableData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTNotifications201ResponseDataRelationshipsNotifiable) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTNotifications201ResponseDataRelationshipsNotifiableData and assigns it to the Data field.
func (o *POSTNotifications201ResponseDataRelationshipsNotifiable) SetData(v POSTNotifications201ResponseDataRelationshipsNotifiableData) {
	o.Data = &v
}

func (o POSTNotifications201ResponseDataRelationshipsNotifiable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTNotifications201ResponseDataRelationshipsNotifiable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTNotifications201ResponseDataRelationshipsNotifiable struct {
	value *POSTNotifications201ResponseDataRelationshipsNotifiable
	isSet bool
}

func (v NullablePOSTNotifications201ResponseDataRelationshipsNotifiable) Get() *POSTNotifications201ResponseDataRelationshipsNotifiable {
	return v.value
}

func (v *NullablePOSTNotifications201ResponseDataRelationshipsNotifiable) Set(val *POSTNotifications201ResponseDataRelationshipsNotifiable) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTNotifications201ResponseDataRelationshipsNotifiable) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTNotifications201ResponseDataRelationshipsNotifiable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTNotifications201ResponseDataRelationshipsNotifiable(val *POSTNotifications201ResponseDataRelationshipsNotifiable) *NullablePOSTNotifications201ResponseDataRelationshipsNotifiable {
	return &NullablePOSTNotifications201ResponseDataRelationshipsNotifiable{value: val, isSet: true}
}

func (v NullablePOSTNotifications201ResponseDataRelationshipsNotifiable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTNotifications201ResponseDataRelationshipsNotifiable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
