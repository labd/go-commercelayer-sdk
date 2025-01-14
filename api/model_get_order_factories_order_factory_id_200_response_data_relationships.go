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

// checks if the GETOrderFactoriesOrderFactoryId200ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETOrderFactoriesOrderFactoryId200ResponseDataRelationships{}

// GETOrderFactoriesOrderFactoryId200ResponseDataRelationships struct for GETOrderFactoriesOrderFactoryId200ResponseDataRelationships
type GETOrderFactoriesOrderFactoryId200ResponseDataRelationships struct {
	SourceOrder *POSTOrderCopies201ResponseDataRelationshipsSourceOrder `json:"source_order,omitempty"`
	TargetOrder *POSTOrderCopies201ResponseDataRelationshipsTargetOrder `json:"target_order,omitempty"`
	Events      *POSTAddresses201ResponseDataRelationshipsEvents        `json:"events,omitempty"`
}

// NewGETOrderFactoriesOrderFactoryId200ResponseDataRelationships instantiates a new GETOrderFactoriesOrderFactoryId200ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrderFactoriesOrderFactoryId200ResponseDataRelationships() *GETOrderFactoriesOrderFactoryId200ResponseDataRelationships {
	this := GETOrderFactoriesOrderFactoryId200ResponseDataRelationships{}
	return &this
}

// NewGETOrderFactoriesOrderFactoryId200ResponseDataRelationshipsWithDefaults instantiates a new GETOrderFactoriesOrderFactoryId200ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrderFactoriesOrderFactoryId200ResponseDataRelationshipsWithDefaults() *GETOrderFactoriesOrderFactoryId200ResponseDataRelationships {
	this := GETOrderFactoriesOrderFactoryId200ResponseDataRelationships{}
	return &this
}

// GetSourceOrder returns the SourceOrder field value if set, zero value otherwise.
func (o *GETOrderFactoriesOrderFactoryId200ResponseDataRelationships) GetSourceOrder() POSTOrderCopies201ResponseDataRelationshipsSourceOrder {
	if o == nil || IsNil(o.SourceOrder) {
		var ret POSTOrderCopies201ResponseDataRelationshipsSourceOrder
		return ret
	}
	return *o.SourceOrder
}

// GetSourceOrderOk returns a tuple with the SourceOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderFactoriesOrderFactoryId200ResponseDataRelationships) GetSourceOrderOk() (*POSTOrderCopies201ResponseDataRelationshipsSourceOrder, bool) {
	if o == nil || IsNil(o.SourceOrder) {
		return nil, false
	}
	return o.SourceOrder, true
}

// HasSourceOrder returns a boolean if a field has been set.
func (o *GETOrderFactoriesOrderFactoryId200ResponseDataRelationships) HasSourceOrder() bool {
	if o != nil && !IsNil(o.SourceOrder) {
		return true
	}

	return false
}

// SetSourceOrder gets a reference to the given POSTOrderCopies201ResponseDataRelationshipsSourceOrder and assigns it to the SourceOrder field.
func (o *GETOrderFactoriesOrderFactoryId200ResponseDataRelationships) SetSourceOrder(v POSTOrderCopies201ResponseDataRelationshipsSourceOrder) {
	o.SourceOrder = &v
}

// GetTargetOrder returns the TargetOrder field value if set, zero value otherwise.
func (o *GETOrderFactoriesOrderFactoryId200ResponseDataRelationships) GetTargetOrder() POSTOrderCopies201ResponseDataRelationshipsTargetOrder {
	if o == nil || IsNil(o.TargetOrder) {
		var ret POSTOrderCopies201ResponseDataRelationshipsTargetOrder
		return ret
	}
	return *o.TargetOrder
}

// GetTargetOrderOk returns a tuple with the TargetOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderFactoriesOrderFactoryId200ResponseDataRelationships) GetTargetOrderOk() (*POSTOrderCopies201ResponseDataRelationshipsTargetOrder, bool) {
	if o == nil || IsNil(o.TargetOrder) {
		return nil, false
	}
	return o.TargetOrder, true
}

// HasTargetOrder returns a boolean if a field has been set.
func (o *GETOrderFactoriesOrderFactoryId200ResponseDataRelationships) HasTargetOrder() bool {
	if o != nil && !IsNil(o.TargetOrder) {
		return true
	}

	return false
}

// SetTargetOrder gets a reference to the given POSTOrderCopies201ResponseDataRelationshipsTargetOrder and assigns it to the TargetOrder field.
func (o *GETOrderFactoriesOrderFactoryId200ResponseDataRelationships) SetTargetOrder(v POSTOrderCopies201ResponseDataRelationshipsTargetOrder) {
	o.TargetOrder = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *GETOrderFactoriesOrderFactoryId200ResponseDataRelationships) GetEvents() POSTAddresses201ResponseDataRelationshipsEvents {
	if o == nil || IsNil(o.Events) {
		var ret POSTAddresses201ResponseDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderFactoriesOrderFactoryId200ResponseDataRelationships) GetEventsOk() (*POSTAddresses201ResponseDataRelationshipsEvents, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *GETOrderFactoriesOrderFactoryId200ResponseDataRelationships) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given POSTAddresses201ResponseDataRelationshipsEvents and assigns it to the Events field.
func (o *GETOrderFactoriesOrderFactoryId200ResponseDataRelationships) SetEvents(v POSTAddresses201ResponseDataRelationshipsEvents) {
	o.Events = &v
}

func (o GETOrderFactoriesOrderFactoryId200ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETOrderFactoriesOrderFactoryId200ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SourceOrder) {
		toSerialize["source_order"] = o.SourceOrder
	}
	if !IsNil(o.TargetOrder) {
		toSerialize["target_order"] = o.TargetOrder
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	return toSerialize, nil
}

type NullableGETOrderFactoriesOrderFactoryId200ResponseDataRelationships struct {
	value *GETOrderFactoriesOrderFactoryId200ResponseDataRelationships
	isSet bool
}

func (v NullableGETOrderFactoriesOrderFactoryId200ResponseDataRelationships) Get() *GETOrderFactoriesOrderFactoryId200ResponseDataRelationships {
	return v.value
}

func (v *NullableGETOrderFactoriesOrderFactoryId200ResponseDataRelationships) Set(val *GETOrderFactoriesOrderFactoryId200ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrderFactoriesOrderFactoryId200ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrderFactoriesOrderFactoryId200ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrderFactoriesOrderFactoryId200ResponseDataRelationships(val *GETOrderFactoriesOrderFactoryId200ResponseDataRelationships) *NullableGETOrderFactoriesOrderFactoryId200ResponseDataRelationships {
	return &NullableGETOrderFactoriesOrderFactoryId200ResponseDataRelationships{value: val, isSet: true}
}

func (v NullableGETOrderFactoriesOrderFactoryId200ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrderFactoriesOrderFactoryId200ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
